package enum

import (
	commonenum "common/pkg/enum"
	realtimehubv1enum "common/proto/gen/realtime_hub/v1/enum"
)

// DeliveryKind 表示 node 可处理的实时下行类别。
type DeliveryKind string

const (
	DeliveryKindNotification DeliveryKind = "notification" // 站内信变更
)

var DeliveryKindMap = commonenum.NewMapping[DeliveryKind, realtimehubv1enum.DeliveryKind](map[DeliveryKind]commonenum.Entry[DeliveryKind, realtimehubv1enum.DeliveryKind]{
	DeliveryKindNotification: {Proto: realtimehubv1enum.DeliveryKind_DELIVERY_KIND_NOTIFICATION},
})
