package enum

import (
	v1 "common/api/gen/content/v1"
	"common/pkg/enum"
)

type CommentOrder string

const (
	CommentOrderNewest  CommentOrder = "newest"
	CommentOrderHottest CommentOrder = "hottest"
)

var CommentOrderMap = enum.NewMapping[CommentOrder, v1.CommentOrder](map[CommentOrder]enum.Entry[CommentOrder, v1.CommentOrder]{
	CommentOrderNewest:  {Proto: v1.CommentOrder_COMMENT_ORDER_NEWEST},
	CommentOrderHottest: {Proto: v1.CommentOrder_COMMENT_ORDER_HOTTEST},
})
