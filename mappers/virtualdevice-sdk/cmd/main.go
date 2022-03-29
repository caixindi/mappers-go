package main

import (
	"gitee.com/ascend/mapper-go-sdk/mapper-go-sdk/pkg/service"
	"gitee.com/ascend/mapper-go-sdk/mappers/virtualdevice-sdk/driver"
)

// main Virtual device program entry
func main() {
	vd := &driver.VirtualDevice{}
	service.Bootstrap("RandomNumber", vd)
}
