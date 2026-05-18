package enum

import (
	v1 "common/api/gen/content/v1"
	"common/pkg/enum"
)

type ArticlePostscriptStatus string

const (
	ArticlePostscriptStatusNormal ArticlePostscriptStatus = "normal"
	ArticlePostscriptStatusHidden ArticlePostscriptStatus = "hidden"
)

var ArticlePostscriptStatusMap = enum.NewMapping[ArticlePostscriptStatus, v1.ArticlePostscriptStatus](map[ArticlePostscriptStatus]enum.Entry[ArticlePostscriptStatus, v1.ArticlePostscriptStatus]{
	ArticlePostscriptStatusNormal: {Proto: v1.ArticlePostscriptStatus_ARTICLE_POSTSCRIPT_STATUS_NORMAL},
	ArticlePostscriptStatusHidden: {Proto: v1.ArticlePostscriptStatus_ARTICLE_POSTSCRIPT_STATUS_HIDDEN},
})
