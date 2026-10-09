package enum

import (
	"common/pkg/enum"
	v1 "common/proto/gen/content/v1/enum"
)

// ArticleVisibility 表示文章可见范围。
type ArticleVisibility string

const (
	ArticleVisibilityPublic  ArticleVisibility = "public"  // ArticleVisibilityPublic 表示公开可见。
	ArticleVisibilityPrivate ArticleVisibility = "private" // ArticleVisibilityPrivate 表示私有可见。
)

// ArticleVisibilityMap 维护文章可见范围内部枚举与 proto 枚举的映射。
var ArticleVisibilityMap = enum.NewMapping[ArticleVisibility, v1.ArticleVisibility](map[ArticleVisibility]enum.Entry[ArticleVisibility, v1.ArticleVisibility]{
	ArticleVisibilityPublic:  {Proto: v1.ArticleVisibility_ARTICLE_VISIBILITY_PUBLIC},
	ArticleVisibilityPrivate: {Proto: v1.ArticleVisibility_ARTICLE_VISIBILITY_PRIVATE},
})

func (e ArticleVisibility) String() string {
	return string(e)
}
