package enum

import (
	v1 "common/api/gen/content/v1"
	"common/pkg/enum"
)

type CommentAction string

const (
	CommentActionLike    CommentAction = "like"
	CommentActionCollect CommentAction = "collect"
	CommentActionReply   CommentAction = "reply"
)

var CommentActionMap = enum.NewMapping[CommentAction, v1.CommentAction](map[CommentAction]enum.Entry[CommentAction, v1.CommentAction]{
	CommentActionLike:    {Proto: v1.CommentAction_COMMENT_ACTION_LIKE},
	CommentActionCollect: {Proto: v1.CommentAction_COMMENT_ACTION_COLLECT},
	CommentActionReply:   {Proto: v1.CommentAction_COMMENT_ACTION_REPLY},
})
