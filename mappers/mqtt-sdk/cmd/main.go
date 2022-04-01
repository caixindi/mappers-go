package main

import (
	"gitee.com/ascend/mapper-go-sdk/mapper-sdk-go/pkg/service"
	"gitee.com/ascend/mapper-go-sdk/mappers/mqtt-sdk/driver"
)

// main Virtual device program entry
func main() {
	mqtt := &driver.MQTT{}
	service.Bootstrap("MQTT", mqtt)
}
