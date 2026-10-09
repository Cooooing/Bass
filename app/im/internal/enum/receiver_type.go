package enum

import (
	"common/pkg/enum"
	v1 "common/proto/gen/im/v1/enum"
)

type ReceiverType string

const (
	ReceiverTypeUser  ReceiverType = "user"  // user
	ReceiverTypeGroup ReceiverType = "group" // group
)

var ReceiverTypeMap = enum.NewMapping[ReceiverType, v1.ReceiverType](map[ReceiverType]enum.Entry[ReceiverType, v1.ReceiverType]{
	ReceiverTypeUser:  {Proto: v1.ReceiverType_RECEIVER_TYPE_USER},
	ReceiverTypeGroup: {Proto: v1.ReceiverType_RECEIVER_TYPE_GROUP},
})

func (e ReceiverType) String() string { return string(e) }
