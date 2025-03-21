// Package service responsible for interacting with developers
package service

import (
	"context"
	"fmt"
	"gitee.com/ascend/mapper-go-sdk/mapper-sdk-go/internal/common"
	"gitee.com/ascend/mapper-go-sdk/mapper-sdk-go/internal/config"
	"gitee.com/ascend/mapper-go-sdk/mapper-sdk-go/internal/configmap"
	"gitee.com/ascend/mapper-go-sdk/mapper-sdk-go/internal/controller"
	"gitee.com/ascend/mapper-go-sdk/mapper-sdk-go/internal/mqttadapter"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"k8s.io/klog/v2"
	"os"
)

// Bootstrap the entrance to mapper
func Bootstrap(serviceName string, deviceInterface interface{}) {
	var err error
	var c config.Config
	klog.InitFlags(nil)
	defer klog.Flush()
	ms = &MapperService{}
	ms.InitMapperService(serviceName, c, deviceInterface)
	klog.V(1).Info("MapperService Init Successful......")

	err = controller.InitDeviceConfig(ms.driver, ms.dic)
	if err != nil {
		klog.Errorf("Failed to init device, please check your interface:%v", err)
		os.Exit(1)
	}
	for id, instance := range ms.deviceInstances {
		ms.wg.Add(1)
		go publishMqtt(id, instance)
	}
	err = initSubscribeMqtt()
	if err != nil {
		klog.Errorf("Failed to subscribe mqtt topic : %v\n", err)
		os.Exit(1)
	}
	ms.wg.Wait()
	klog.V(1).Info("All devices have been deleted.Mapper exit")
}

func publishMqtt(id string, instance *configmap.DeviceInstance) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	err := mqttadapter.SendTwin(ctx, id, instance, ms.driver, ms.mqttClient, ms.wg, ms.dic, ms.deviceMutex[id])
	if err != nil {
		klog.Errorf("Failed to get %s %s:%v\n", id, "twin", err)
	} else {
		err = mqttadapter.SendData(ctx, id, instance, ms.driver, ms.mqttClient, ms.wg, ms.dic, ms.deviceMutex[id])
		if err != nil {
			klog.Errorf("Failed to get %s %s:%v\n", id, "data", err)
		}
		err = mqttadapter.SendDeviceState(ctx, id, instance, ms.driver, ms.mqttClient, ms.wg, ms.dic, ms.deviceMutex[id])
		if err != nil {
			klog.Errorf("Failed to get %s %s:%v\n", id, "state", err)
		}
	}
	ms.stopFunctions[id] = cancelFunc
	ms.wg.Done()
}

func initSubscribeMqtt() error {
	for k := range ms.deviceInstances {
		topic := fmt.Sprintf(common.TopicTwinUpdateDelta, k)
		onMessage := func(client mqtt.Client, message mqtt.Message) {
			mqttadapter.SyncInfo(ms.dic, message)
		}
		err := ms.mqttClient.Subscribe(topic, onMessage)
		if err != nil {
			return err
		}
		updateDevice := func(client mqtt.Client, message mqtt.Message) {
			mqttadapter.UpdateDevice(ms.dic, message)
		}
		err = ms.mqttClient.Subscribe(common.TopicDeviceUpdate, updateDevice)
		if err != nil {
			return err
		}
		klog.V(1).Infof("Event %s is Listening....\n", k)
	}
	return nil
}
