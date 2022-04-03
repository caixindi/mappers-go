package main

import (
	"gitee.com/ascend/mapper-go-sdk/mapper-sdk-go/pkg/service"
	"gitee.com/ascend/mapper-go-sdk/mappers/idmvs-sdk/driver"
)

// main IDMVS device program entry
func main() {
	d := &driver.IDMVS{}
	service.Bootstrap("IDMVS", d)
}
