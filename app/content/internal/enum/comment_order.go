package enum

import (
	"common/pkg/enum"
	v1 "common/proto/gen/content/v1/enum"
)

// CommentOrder 表示评论列表排序方式。
type CommentOrder string

const (
	// CommentOrderNewest 表示按最新发布时间排序。
	CommentOrderNewest CommentOrder = "newest"
	// CommentOrderHottest 表示按热度排序。
	CommentOrderHottest CommentOrder = "hottest"
	// CommentOrderOldest 表示按最早发布时间排序。
	CommentOrderOldest CommentOrder = "oldest"
)

// CommentOrderMap 维护评论排序内部枚举与 proto 枚举的映射。
var CommentOrderMap = enum.NewMapping[CommentOrder, v1.CommentOrder](map[CommentOrder]enum.Entry[CommentOrder, v1.CommentOrder]{
	CommentOrderNewest:  {Proto: v1.CommentOrder_COMMENT_ORDER_NEWEST},
	CommentOrderHottest: {Proto: v1.CommentOrder_COMMENT_ORDER_HOTTEST},
	CommentOrderOldest:  {Proto: v1.CommentOrder_COMMENT_ORDER_OLDEST},
})
