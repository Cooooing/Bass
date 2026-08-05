package enum

import (
	"common/pkg/enum"
	v1 "common/proto/gen/content/v1/enum"
)

// CommentAction 表示评论互动行为类型。
type CommentAction string

const (
	// CommentActionLike 表示点赞评论。
	CommentActionLike CommentAction = "like"
	// CommentActionReply 表示回复评论。
	CommentActionReply CommentAction = "reply"
	// CommentActionThank 表示感谢评论。
	CommentActionThank CommentAction = "thank"
)

// CommentActionMap 维护评论行为内部枚举与 proto 枚举的映射。
var CommentActionMap = enum.NewMapping[CommentAction, v1.CommentAction](map[CommentAction]enum.Entry[CommentAction, v1.CommentAction]{
	CommentActionLike:  {Proto: v1.CommentAction_COMMENT_ACTION_LIKE},
	CommentActionReply: {Proto: v1.CommentAction_COMMENT_ACTION_REPLY},
	CommentActionThank: {Proto: v1.CommentAction_COMMENT_ACTION_THANK},
})

func (e CommentAction) String() string {
	return string(e)
}
