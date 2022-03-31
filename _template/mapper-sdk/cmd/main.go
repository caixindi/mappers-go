package main

import (
	"gitee.com/ascend/mapper-go-sdk/mapper-go-sdk/pkg/service"
	"gitee.com/ascend/mapper-go-sdk/mappers/Template/driver"
)

// main Virtual device program entry
func main() {
	d := &driver.Template{}
	service.Bootstrap("Template", d)
}
