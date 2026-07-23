package enum

import (
	"common/pkg/enum"
	v1 "common/proto/gen/content/v1/enum"
)

// ContentRestriction 表示内容的管理限制状态。
type ContentRestriction string

const (
	// ContentRestrictionNone 表示没有管理限制。
	ContentRestrictionNone ContentRestriction = "none"
	// ContentRestrictionHidden 表示被管理端隐藏。
	ContentRestrictionHidden ContentRestriction = "hidden"
	// ContentRestrictionLocked 表示被管理端锁定。
	ContentRestrictionLocked ContentRestriction = "locked"
)

// ContentRestrictionMap 维护内容管理限制内部枚举与 proto 枚举的映射。
var ContentRestrictionMap = enum.NewMapping[ContentRestriction, v1.ContentRestriction](map[ContentRestriction]enum.Entry[ContentRestriction, v1.ContentRestriction]{
	ContentRestrictionNone:   {Proto: v1.ContentRestriction_CONTENT_RESTRICTION_NONE},
	ContentRestrictionHidden: {Proto: v1.ContentRestriction_CONTENT_RESTRICTION_HIDDEN},
	ContentRestrictionLocked: {Proto: v1.ContentRestriction_CONTENT_RESTRICTION_LOCKED},
})
