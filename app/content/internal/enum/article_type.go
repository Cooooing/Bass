package enum

import (
	"common/pkg/enum"
	v1 "common/proto/gen/content/v1/enum"
)

// ArticleType 表示文章业务类型。
type ArticleType string

const (
	// ArticleTypeNormal 表示普通文章。
	ArticleTypeNormal ArticleType = "normal"
	// ArticleTypeQA 表示问答文章。
	ArticleTypeQA ArticleType = "qa"
)

// ArticleTypeMap 维护文章类型内部枚举与 proto 枚举的映射。
var ArticleTypeMap = enum.NewMapping[ArticleType, v1.ArticleType](map[ArticleType]enum.Entry[ArticleType, v1.ArticleType]{
	ArticleTypeNormal: {Proto: v1.ArticleType_ARTICLE_TYPE_NORMAL},
	ArticleTypeQA:     {Proto: v1.ArticleType_ARTICLE_TYPE_QA},
})
