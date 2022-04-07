// Package common used to store constants, data conversion functions, timers, etc
package common

// joint the topic like topic := fmt.Sprintf(TopicTwinUpdateDelta, deviceID)
const (
	TopicTwinUpdateDelta = "$hw/events/device/%s/twin/update/delta"
	TopicTwinUpdate      = "$hw/events/device/%s/twin/update"
	TopicStateUpdate     = "$hw/events/device/%s/state/update"
	TopicDataUpdate      = "$ke/events/device/%s/data/update"
	TopicDeviceUpdate    = "$hw/events/node/#"
)

// Device status definition.
const (
	DEVSTOK      = "OK"
	DEVSTDISCONN = "DISCONNECTED"
)

// joint x joint the instancepool like driverName :=  common.DriverPrefix+instanceID+twin.PropertyName
const (
	DriverPrefix = "Driver"
)

const (
	CorrelationHeader = "X-Correlation-ID"
)

const (
	APIVersion = "v1"
	APIBase    = "/api/v1"

	APIDeviceRoute                 = APIBase + "/device"
	APIDeviceWriteCommandByIDRoute = APIDeviceRoute + "/" + ID + "/{" + IDAndCommand + "}"
	APIDeviceReadCommandByIDRoute  = APIDeviceRoute + "/" + ID + "/{" + ID + "}" + "/{" + Command + "}"
	APIDeviceCallbackRoute         = APIBase + "/callback/device"
	APIDeviceCallbackIDRoute       = APIBase + "/callback/device/id/{id}"

	APIPingRoute = APIBase + "/ping"
)

const (
	ID           = "id"
	Command      = "command"
	IDAndCommand = "IdAndCommand"
)

// Constants related to the possible content types supported by the APIs
const (
	ContentType     = "Content-Type"
	ContentTypeJSON = "application/json"
)

type ErrKind string

// Constant Kind identifiers which can be used to label and group errors.
const (
	KindEntityDoesNotExist  ErrKind = "NotFound"
	KindServerError         ErrKind = "UnexpectedServerError"
	KindDuplicateName       ErrKind = "DuplicateName"
	KindInvalidID           ErrKind = "InvalidId"
	KindServiceUnavailable  ErrKind = "ServiceUnavailable"
	KindNotAllowed          ErrKind = "NotAllowed"
	KindServiceLocked       ErrKind = "ServiceLocked"
	KindNotImplemented      ErrKind = "NotImplemented"
	KindRangeNotSatisfiable ErrKind = "RangeNotSatisfiable"
	KindOverflowError       ErrKind = "OverflowError"
	KindNaNError            ErrKind = "NaNError"
)
