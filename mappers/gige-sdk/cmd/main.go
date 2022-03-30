package main

import (
	"gitee.com/ascend/mapper-go-sdk/mapper-go-sdk/pkg/service"
	"gitee.com/ascend/mapper-go-sdk/mappers/gige-sdk/driver"
)

func main() {
	gd := &driver.GigEVisionDevice{}
	service.Bootstrap("GigECamera", gd)
}
