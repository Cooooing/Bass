package enum

import (
	"common/pkg/enum"
	v1 "common/proto/gen/content/v1/enum"
)

// ContentAccessScope 表示内容访问身份视角
type ContentAccessScope string

const (
	// ContentAccessScopeGuest 表示游客视角
	ContentAccessScopeGuest ContentAccessScope = "guest"
	// ContentAccessScopeUser 表示普通用户视角
	ContentAccessScopeUser ContentAccessScope = "user"
	// ContentAccessScopeAuthor 表示作者视角
	ContentAccessScopeAuthor ContentAccessScope = "author"
	// ContentAccessScopeAdmin 表示后台管理视角
	ContentAccessScopeAdmin ContentAccessScope = "admin"
	// ContentAccessScopeInternalTask 表示内部任务视角
	ContentAccessScopeInternalTask ContentAccessScope = "internal_task"
)

// ContentAccessScopeMap 维护访问视角内部枚举与 proto 枚举的映射
var ContentAccessScopeMap = enum.NewMapping[ContentAccessScope, v1.ContentAccessScope](map[ContentAccessScope]enum.Entry[ContentAccessScope, v1.ContentAccessScope]{
	ContentAccessScopeGuest:        {Proto: v1.ContentAccessScope_CONTENT_ACCESS_SCOPE_GUEST},
	ContentAccessScopeUser:         {Proto: v1.ContentAccessScope_CONTENT_ACCESS_SCOPE_USER},
	ContentAccessScopeAuthor:       {Proto: v1.ContentAccessScope_CONTENT_ACCESS_SCOPE_AUTHOR},
	ContentAccessScopeAdmin:        {Proto: v1.ContentAccessScope_CONTENT_ACCESS_SCOPE_ADMIN},
	ContentAccessScopeInternalTask: {Proto: v1.ContentAccessScope_CONTENT_ACCESS_SCOPE_INTERNAL_TASK},
})

func (e ContentAccessScope) String() string {
	return string(e)
}
