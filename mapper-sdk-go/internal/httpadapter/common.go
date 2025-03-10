package httpadapter

import (
	"gitee.com/ascend/mapper-go-sdk/mapper-sdk-go/internal/common"
	"net/http"
	"time"
)

// Ping handles the requests to /ping endpoint. Is used to test if the service is working
// It returns a response as specified by the V1 API swagger in openapi/common
func (c *RestController) Ping(writer http.ResponseWriter, request *http.Request) {
	response := "This is api " + common.APIVersion + ". Now is " + time.Now().Format(time.UnixDate)
	c.sendResponse(writer, request, common.APIPingRoute, response, http.StatusOK)
}
