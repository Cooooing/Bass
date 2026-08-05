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
	// ArticleTypeLottery 表示抽奖文章。
	ArticleTypeLottery ArticleType = "lottery"
	// ArticleTypePoll 表示投票文章。
	ArticleTypePoll ArticleType = "poll"
	// ArticleTypeColumn 表示专栏文章。
	ArticleTypeColumn ArticleType = "column"
)

// ArticleTypeMap 维护文章类型内部枚举与 proto 枚举的映射。
var ArticleTypeMap = enum.NewMapping[ArticleType, v1.ArticleType](map[ArticleType]enum.Entry[ArticleType, v1.ArticleType]{
	ArticleTypeNormal:  {Proto: v1.ArticleType_ARTICLE_TYPE_NORMAL},
	ArticleTypeQA:      {Proto: v1.ArticleType_ARTICLE_TYPE_QA},
	ArticleTypeLottery: {Proto: v1.ArticleType_ARTICLE_TYPE_LOTTERY},
	ArticleTypePoll:    {Proto: v1.ArticleType_ARTICLE_TYPE_POLL},
	ArticleTypeColumn:  {Proto: v1.ArticleType_ARTICLE_TYPE_COLUMN},
})

func (e ArticleType) String() string {
	return string(e)
}
