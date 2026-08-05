package enum

import (
	"common/pkg/enum"
	v1 "common/proto/gen/content/v1/enum"
)

// ContentModerationTarget 表示内容管理处置目标类型。
type ContentModerationTarget string

const (
	// ContentModerationTargetArticle 表示处置目标为文章。
	ContentModerationTargetArticle ContentModerationTarget = "article"
	// ContentModerationTargetComment 表示处置目标为评论。
	ContentModerationTargetComment ContentModerationTarget = "comment"
	// ContentModerationTargetPostscript 表示处置目标为附言。
	ContentModerationTargetPostscript ContentModerationTarget = "postscript"
)

// ContentModerationTargetMap 维护内容管理目标内部枚举与 proto 枚举的映射。
var ContentModerationTargetMap = enum.NewMapping[ContentModerationTarget, v1.ContentModerationTarget](map[ContentModerationTarget]enum.Entry[ContentModerationTarget, v1.ContentModerationTarget]{
	ContentModerationTargetArticle:    {Proto: v1.ContentModerationTarget_CONTENT_MODERATION_TARGET_ARTICLE},
	ContentModerationTargetComment:    {Proto: v1.ContentModerationTarget_CONTENT_MODERATION_TARGET_COMMENT},
	ContentModerationTargetPostscript: {Proto: v1.ContentModerationTarget_CONTENT_MODERATION_TARGET_POSTSCRIPT},
})

// ContentModerationAction 表示内容管理处置动作。
type ContentModerationAction string

const (
	// ContentModerationActionHide 表示管理隐藏内容。
	ContentModerationActionHide ContentModerationAction = "hide"
	// ContentModerationActionUnhide 表示取消管理隐藏。
	ContentModerationActionUnhide ContentModerationAction = "unhide"
	// ContentModerationActionLock 表示管理锁定内容。
	ContentModerationActionLock ContentModerationAction = "lock"
	// ContentModerationActionUnlock 表示取消管理锁定。
	ContentModerationActionUnlock ContentModerationAction = "unlock"
	// ContentModerationActionDelete 表示管理删除内容。
	ContentModerationActionDelete ContentModerationAction = "delete"
	// ContentModerationActionRestore 表示恢复已删除内容。
	ContentModerationActionRestore ContentModerationAction = "restore"
)

// ContentModerationActionMap 维护内容管理动作内部枚举与 proto 枚举的映射。
var ContentModerationActionMap = enum.NewMapping[ContentModerationAction, v1.ContentModerationAction](map[ContentModerationAction]enum.Entry[ContentModerationAction, v1.ContentModerationAction]{
	ContentModerationActionHide:    {Proto: v1.ContentModerationAction_CONTENT_MODERATION_ACTION_HIDE},
	ContentModerationActionUnhide:  {Proto: v1.ContentModerationAction_CONTENT_MODERATION_ACTION_UNHIDE},
	ContentModerationActionLock:    {Proto: v1.ContentModerationAction_CONTENT_MODERATION_ACTION_LOCK},
	ContentModerationActionUnlock:  {Proto: v1.ContentModerationAction_CONTENT_MODERATION_ACTION_UNLOCK},
	ContentModerationActionDelete:  {Proto: v1.ContentModerationAction_CONTENT_MODERATION_ACTION_DELETE},
	ContentModerationActionRestore: {Proto: v1.ContentModerationAction_CONTENT_MODERATION_ACTION_RESTORE},
})

func (e ContentModerationAction) String() string {
	return string(e)
}

func (e ContentModerationTarget) String() string {
	return string(e)
}
