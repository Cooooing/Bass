package enum

import (
	v1 "common/api/gen/content/v1"
	"common/pkg/enum"
)

type CommentStatus string

const (
	CommentStatusNormal CommentStatus = "normal"
	CommentStatusHidden CommentStatus = "hidden"
)

var CommentStatusMap = enum.NewMapping[CommentStatus, v1.CommentStatus](map[CommentStatus]enum.Entry[CommentStatus, v1.CommentStatus]{
	CommentStatusNormal: {Proto: v1.CommentStatus_COMMENT_STATUS_NORMAL},
	CommentStatusHidden: {Proto: v1.CommentStatus_COMMENT_STATUS_HIDDEN},
})
