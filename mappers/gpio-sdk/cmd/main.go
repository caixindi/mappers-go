package main

import (
	"gitee.com/ascend/mapper-go-sdk/mapper-sdk-go/pkg/service"
	"gitee.com/ascend/mapper-go-sdk/mappers/gpio-sdk/driver"
)

// main Virtual device program entry
func main() {
	gpio := &driver.GPIO{}
	service.Bootstrap("GPIO", gpio)
}
