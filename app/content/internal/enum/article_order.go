package enum

import (
	"common/pkg/enum"
	v1 "common/proto/gen/content/v1/enum"
)

// ArticleOrder 表示文章列表排序方式。
type ArticleOrder string

const (
	// ArticleOrderNewest 表示按最新发布时间排序。
	ArticleOrderNewest ArticleOrder = "newest"
	// ArticleOrderHottest 表示按热度排序。
	ArticleOrderHottest ArticleOrder = "hottest"
)

// ArticleOrderMap 维护文章排序内部枚举与 proto 枚举的映射。
var ArticleOrderMap = enum.NewMapping[ArticleOrder, v1.ArticleOrder](map[ArticleOrder]enum.Entry[ArticleOrder, v1.ArticleOrder]{
	ArticleOrderNewest:  {Proto: v1.ArticleOrder_ARTICLE_ORDER_NEWEST},
	ArticleOrderHottest: {Proto: v1.ArticleOrder_ARTICLE_ORDER_HOTTEST},
})

func (e ArticleOrder) String() string {
	return string(e)
}
