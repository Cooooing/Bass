package enum

import (
	"common/pkg/enum"
	v1 "common/proto/gen/content/v1/enum"
)

// TagStatus 表示标签状态。
type TagStatus string

const (
	// TagStatusEnabled 表示标签启用。
	TagStatusEnabled TagStatus = "enabled"
	// TagStatusDisabled 表示标签禁用。
	TagStatusDisabled TagStatus = "disabled"
)

// TagStatusMap 维护标签状态内部枚举与 proto 枚举的映射。
var TagStatusMap = enum.NewMapping[TagStatus, v1.TagStatus](map[TagStatus]enum.Entry[TagStatus, v1.TagStatus]{
	TagStatusEnabled:  {Proto: v1.TagStatus_TAG_STATUS_ENABLED},
	TagStatusDisabled: {Proto: v1.TagStatus_TAG_STATUS_DISABLED},
})
