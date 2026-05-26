package enum

import (
	v1 "common/api/gen/content/v1"
	"common/pkg/enum"
)

type ArticleType string

const (
	ArticleTypeNormal ArticleType = "normal"
	ArticleTypeQA     ArticleType = "qa"
)

var ArticleTypeMap = enum.NewMapping[ArticleType, v1.ArticleType](map[ArticleType]enum.Entry[ArticleType, v1.ArticleType]{
	ArticleTypeNormal: {Proto: v1.ArticleType_ARTICLE_TYPE_NORMAL},
	ArticleTypeQA:     {Proto: v1.ArticleType_ARTICLE_TYPE_QA},
})
