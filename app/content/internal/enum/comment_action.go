package enum

import (
	v1 "common/api/gen/content/v1"
	"common/pkg/enum"
)

type CommentAction string

const (
	CommentActionLike  CommentAction = "like"
	CommentActionReply CommentAction = "reply"
	CommentActionThank CommentAction = "thank"
)

var CommentActionMap = enum.NewMapping[CommentAction, v1.CommentAction](map[CommentAction]enum.Entry[CommentAction, v1.CommentAction]{
	CommentActionLike:  {Proto: v1.CommentAction_COMMENT_ACTION_LIKE},
	CommentActionReply: {Proto: v1.CommentAction_COMMENT_ACTION_REPLY},
	CommentActionThank: {Proto: v1.CommentAction_COMMENT_ACTION_THANK},
})
