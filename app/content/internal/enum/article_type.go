package enum

import (
	v1 "common/api/gen/content/v1"
	"common/pkg/enum"
)

type ArticleType string

const (
	ArticleTypeNormal  ArticleType = "normal"
	ArticleTypeQA      ArticleType = "qa"
	ArticleTypeVote    ArticleType = "vote"
	ArticleTypeLottery ArticleType = "lottery"
)

var ArticleTypeMap = enum.NewMapping[ArticleType, v1.ArticleType](map[ArticleType]enum.Entry[ArticleType, v1.ArticleType]{
	ArticleTypeNormal:  {Proto: v1.ArticleType_ARTICLE_TYPE_NORMAL},
	ArticleTypeQA:      {Proto: v1.ArticleType_ARTICLE_TYPE_QA},
	ArticleTypeVote:    {Proto: v1.ArticleType_ARTICLE_TYPE_VOTE},
	ArticleTypeLottery: {Proto: v1.ArticleType_ARTICLE_TYPE_LOTTERY},
})
