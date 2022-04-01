// Package requests used to call define add device request's struct
package requests

import "gitee.com/ascend/mapper-go-sdk/mapper-sdk-go/internal/configmap"

type AddDeviceRequest struct {
	DeviceInstance *configmap.DeviceInstance `json:"deviceInstances"`
	DeviceModels   []*configmap.DeviceModel  `json:"deviceModels"`
	Protocol       []*configmap.Protocol     `json:"protocols"`
}
