package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/content/v1/enum"
)

// ContentActionStatus 表示需要异步确认的内容互动命令状态。
type ContentActionStatus string

const (
	ContentActionStatusPending   ContentActionStatus = "pending"   // ContentActionStatusPending 表示待确认。
	ContentActionStatusPublished ContentActionStatus = "published" // ContentActionStatusPublished 表示已生效。
	ContentActionStatusCanceled  ContentActionStatus = "canceled"  // ContentActionStatusCanceled 表示已取消。
)

// ContentActionStatusMap 维护互动命令状态映射。
var ContentActionStatusMap = commonenum.NewMapping[ContentActionStatus, v1.ContentActionStatus](map[ContentActionStatus]commonenum.Entry[ContentActionStatus, v1.ContentActionStatus]{
	ContentActionStatusPending:   {Proto: v1.ContentActionStatus_CONTENT_ACTION_STATUS_PENDING},
	ContentActionStatusPublished: {Proto: v1.ContentActionStatus_CONTENT_ACTION_STATUS_PUBLISHED},
	ContentActionStatusCanceled:  {Proto: v1.ContentActionStatus_CONTENT_ACTION_STATUS_CANCELED},
})

func (e ContentActionStatus) String() string { return string(e) }
