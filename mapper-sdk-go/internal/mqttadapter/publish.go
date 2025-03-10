package mqttadapter

import (
	"context"
	"errors"
	"fmt"
	"gitee.com/ascend/mapper-go-sdk/mapper-sdk-go/internal/clients/mqttclient"
	"gitee.com/ascend/mapper-go-sdk/mapper-sdk-go/internal/common"
	"gitee.com/ascend/mapper-go-sdk/mapper-sdk-go/internal/configmap"
	"gitee.com/ascend/mapper-go-sdk/mapper-sdk-go/internal/controller"
	"gitee.com/ascend/mapper-go-sdk/mapper-sdk-go/pkg/di"
	"gitee.com/ascend/mapper-go-sdk/mapper-sdk-go/pkg/models"
	"sync"
	"time"
)

// SendTwin send twin to EdgeCore according to timer
func SendTwin(ctx context.Context, id string, instance *configmap.DeviceInstance, drivers models.ProtocolDriver, mqttClient mqttclient.MqttClient, wg *sync.WaitGroup, dic *di.Container, mutex *common.Lock) error {
	for _, twinV := range instance.Twins {
		// ---------------setVisitor---------------
		err := controller.SetVisitor(id, twinV, drivers, mutex, dic)
		if err != nil {
			err = errors.New("Set device config error:" + err.Error())
			return err
		}
		// ---------------setVisitor---------------
		// ---------------Send Data by MQTT---------------
		collectCycle := time.Duration(twinV.PVisitor.CollectCycle)
		wg.Add(1)
		if collectCycle == -1 {
			go func() {
				defer wg.Done()
				<-ctx.Done()
			}()
		} else {
			// If the collect cycle is not set, set it to 1 second.
			if collectCycle == 0 {
				collectCycle = 1 * time.Second
			}
			twinData := TwinData{
				Name:       twinV.PropertyName,
				Type:       twinV.Desired.Metadatas.Type,
				Topic:      fmt.Sprintf(common.TopicTwinUpdate, id),
				MqttClient: mqttClient,
				driverUnit: DriverUnit{
					instanceID: id,
					twin:       twinV,
					drivers:    drivers,
					mutex:      mutex,
					dic:        dic,
				},
			}
			timer := common.Timer{Function: twinData.Run, Duration: collectCycle, Times: 0}
			go func() {
				timer.Start()
			}()
			go func() {
				defer wg.Done()
				<-ctx.Done()
				timer.Stop()
			}()
		}

		// ---------------Send Data by MQTT---------------
	}
	return nil
}

// SendData send twin to third-part application according to timer
func SendData(ctx context.Context, id string, instance *configmap.DeviceInstance, drivers models.ProtocolDriver, mqttClient mqttclient.MqttClient, wg *sync.WaitGroup, dic *di.Container, mutex *common.Lock) error {
	for _, twinV := range instance.Twins {
		// ---------------Send Data by MQTT---------------
		collectCycle := time.Duration(twinV.PVisitor.CollectCycle)
		// If the collect cycle is not set, set it to 1 second.
		if collectCycle == -1 {
			wg.Add(1)
			go func() {
				defer wg.Done()
				<-ctx.Done()
			}()
		} else {
			if collectCycle == 0 {
				collectCycle = 1 * time.Second
			}
			twinData := TwinData{
				Name:       twinV.PropertyName,
				Type:       twinV.Desired.Metadatas.Type,
				Topic:      fmt.Sprintf(common.TopicDataUpdate, id),
				MqttClient: mqttClient,
				driverUnit: DriverUnit{
					instanceID: id,
					twin:       twinV,
					drivers:    drivers,
					mutex:      mutex,
					dic:        dic,
				},
			}
			timer := common.Timer{Function: twinData.Run, Duration: collectCycle, Times: 0}
			wg.Add(1)
			go func() {
				timer.Start()
			}()
			go func() {
				defer wg.Done()
				<-ctx.Done()
				timer.Stop()
			}()
		}
		// ---------------Send Data by MQTT---------------
	}
	return nil
}

// SendDeviceState send device's state to EdgeCore according to timer
func SendDeviceState(ctx context.Context, id string, instance *configmap.DeviceInstance, drivers models.ProtocolDriver, mqttClient mqttclient.MqttClient, wg *sync.WaitGroup, dic *di.Container, mutex *common.Lock) error {
	var statusData StatusData
	var collectCycle time.Duration
	for _, twinV := range instance.Twins {
		// ---------------Send Data by MQTT---------------
		collectCycle = time.Duration(twinV.PVisitor.CollectCycle)
		if collectCycle == -1 {
			wg.Add(1)
			go func() {
				defer wg.Done()
				<-ctx.Done()
			}()
		} else {
			// If the collect cycle is not set, set it to 1 second.
			if collectCycle == 0 {
				collectCycle = 1 * time.Second
			}
			statusData = StatusData{
				topic:      fmt.Sprintf(common.TopicStateUpdate, id),
				MqttClient: mqttClient,
				driverUnit: DriverUnit{
					instanceID: id,
					twin:       twinV,
					drivers:    drivers,
					mutex:      mutex,
					dic:        dic,
				},
			}
			timer := common.Timer{Function: statusData.Run, Duration: collectCycle, Times: 0}
			wg.Add(1)
			go func() {
				timer.Start()
			}()
			go func() {
				defer wg.Done()
				<-ctx.Done()
				timer.Stop()
			}()
		}
	}
	return nil
}
