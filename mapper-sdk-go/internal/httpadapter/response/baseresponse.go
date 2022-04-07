// Package response used to implement the responses structure
package response

import "gitee.com/ascend/mapper-go-sdk/mapper-sdk-go/internal/common"

type BaseResponse struct {
	Version    string
	RequestID  string `json:"requestId,omitempty"`
	Message    string `json:"message,omitempty"`
	StatusCode int    `json:"statusCode"`
}

type ReadCommandResponse struct {
	BaseResponse
	DeviceID     string
	PropertyName string
	Value        string
}

type WriteCommandResponse struct {
	BaseResponse
	DeviceID     string
	PropertyName string
	Status       string
}

type UpdateDeviceResponse struct {
	BaseResponse
	DeviceID  string
	Operation string
	Status    string
}

func NewBaseResponse(requestID string, message string, statusCode int) BaseResponse {
	return BaseResponse{
		Version:    common.APIVersion,
		RequestID:  requestID,
		Message:    message,
		StatusCode: statusCode,
	}
}

func NewReadCommandResponse(response BaseResponse, deviceID, propertyName, value string) ReadCommandResponse {
	return ReadCommandResponse{
		response,
		deviceID,
		propertyName,
		value,
	}
}

func NewWriteCommandResponse(response BaseResponse, deviceID, propertyName, status string) WriteCommandResponse {
	return WriteCommandResponse{
		response,
		deviceID,
		propertyName,
		status,
	}
}

func NewUpdateDeviceResponse(response BaseResponse, deviceID, operation, status string) UpdateDeviceResponse {
	return UpdateDeviceResponse{
		response,
		deviceID,
		operation,
		status,
	}
}
