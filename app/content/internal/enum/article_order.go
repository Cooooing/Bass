package enum

import (
	v1 "common/api/gen/content/v1"
	"common/pkg/enum"
)

type ArticleOrder string

const (
	ArticleOrderNewest  ArticleOrder = "newest"
	ArticleOrderHottest ArticleOrder = "hottest"
)

var ArticleOrderMap = enum.NewMapping[ArticleOrder, v1.ArticleOrder](map[ArticleOrder]enum.Entry[ArticleOrder, v1.ArticleOrder]{
	ArticleOrderNewest:  {Proto: v1.ArticleOrder_ARTICLE_ORDER_NEWEST},
	ArticleOrderHottest: {Proto: v1.ArticleOrder_ARTICLE_ORDER_HOTTEST},
})
