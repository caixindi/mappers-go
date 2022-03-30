package driver

/*
#include <dlfcn.h>
int open_device(unsigned int** device,char* deviceId,char** error);
int get_value (unsigned int* device, char* feature, char** value,char** error);
void close_device (unsigned int* device);
#cgo LDFLAGS: -ldl
*/
import "C"
import (
	"encoding/json"
	"fmt"
	"sync"
)

type GigEVisionDeviceProtocolConfig struct {
}

type GigEVisionDeviceProtocolCommonConfig struct {
	DeviceId string `json:"deviceId"`
}

type GigEVisionDeviceVisitorConfig struct {
	FeatureName string `json:"FeatureName"`
}

type GigEVisionDevice struct {
	mutex                sync.RWMutex
	GigEprotocolConfig   GigEVisionDeviceProtocolConfig
	ProtocolCommonConfig GigEVisionDeviceProtocolCommonConfig
	visitorConfig        GigEVisionDeviceVisitorConfig
	dev                  *C.uint
}

func (geClient *GigEVisionDevice) InitDevice(protocolCommon []byte) (err error) {
	if protocolCommon != nil {
		if err = json.Unmarshal(protocolCommon, &geClient.ProtocolCommonConfig); err != nil {
			fmt.Printf("Unmarshal ProtocolCommonConfig error: %v\n", err)
			return err
		}
	}
	err = geClient.NewClient()
	if err != nil {
		fmt.Printf("NewClient error: %v\n", err)
		return err
	}
	return nil
}

func (geClient *GigEVisionDevice) SetConfig(protocolCommon, visitor, protocol []byte) (err error) {
	//geClient.NewClient()
	if protocolCommon != nil {
		if err = json.Unmarshal(protocolCommon, &geClient.ProtocolCommonConfig); err != nil {
			fmt.Printf("Unmarshal ProtocolCommonConfig error: %v\n", err)
			return err
		}
		//klog.V(0).Info(geClient.GigEprotocolConfig)
	}
	if visitor != nil {
		if err = json.Unmarshal(visitor, &geClient.visitorConfig); err != nil {
			fmt.Printf("Unmarshal visitorConfig error: %v\n", err)
			return err
		}

	}

	if protocol != nil {
		if err = json.Unmarshal(protocol, &geClient.GigEprotocolConfig); err != nil {
			fmt.Printf("Unmarshal ProtocolConfig error: %v\n", err)
			return err
		}
		//klog.V(0).Info(geClient.ProtocolCommonConfig)
	}

	return nil
}

// ReadDeviceData  is an interface that reads data from a specific device, data is a type of string
func (geClient *GigEVisionDevice) ReadDeviceData(protocolCommon, visitor, protocol []byte) (data interface{}, err error) {
	geClient.NewClient()
	err = geClient.SetConfig(protocolCommon, visitor, protocol)
	if err != nil {
		return nil, err
	}
	data, err = geClient.Get(geClient.visitorConfig.FeatureName)
	if err != nil {
		fmt.Println("GigE Get error", err.Error())
		return nil, err
	}
	return data, err
}

// WriteDeviceData is an interface that write data to a specific device, data's DataType is Consistent with configmap
func (geClient *GigEVisionDevice) WriteDeviceData(data interface{}, protocolCommon, visitor, protocol []byte) (err error) {
	geClient.NewClient()
	err = geClient.SetConfig(protocolCommon, visitor, protocol)
	if err != nil {
		return err
	}
	//geClient.mutex.Lock()
	//defer geClient.mutex.Unlock()
	err = geClient.Set(geClient.visitorConfig.FeatureName, data)
	if err != nil {
		fmt.Println("Set Err", err.Error())
		return err
	}
	return nil
}

// StopDevice is an interface to disconnect a specific device
func (geClient *GigEVisionDevice) StopDevice() (err error) {
	geClient.mutex.Lock()
	defer geClient.mutex.Unlock()
	C.close_device(geClient.dev)
	geClient.dev = nil
	fmt.Println("----------Stop GigE Device Successful----------")
	return nil
}

// GetDeviceStatus is an interface to get the device status true is OK , false is DISCONNECTED
func (geClient *GigEVisionDevice) GetDeviceStatus(protocolCommon, visitor, protocol []byte) (status bool) {
	var msg *C.char
	var value *C.char
	err := geClient.SetConfig(protocolCommon, visitor, protocol)
	if err != nil {
		return false
	}
	signal := C.get_value(geClient.dev, C.CString(geClient.visitorConfig.FeatureName), &value, &msg)
	if signal != 0 {
		fmt.Println("error:", C.GoString(msg))
		return false
	}
	return true
}
