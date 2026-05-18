package enum

import (
	v1 "common/api/gen/content/v1"
	"common/pkg/enum"
)

type TagStatus string

const (
	TagStatusNormal   TagStatus = "normal"
	TagStatusDisabled TagStatus = "disabled"
)

var TagStatusMap = enum.NewMapping[TagStatus, v1.TagStatus](map[TagStatus]enum.Entry[TagStatus, v1.TagStatus]{
	TagStatusNormal:   {Proto: v1.TagStatus_TAG_STATUS_NORMAL},
	TagStatusDisabled: {Proto: v1.TagStatus_TAG_STATUS_DISABLED},
})
