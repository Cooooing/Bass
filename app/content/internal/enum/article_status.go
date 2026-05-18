package enum

import (
	v1 "common/api/gen/content/v1"
	"common/pkg/enum"
)

type ArticleStatus string

const (
	ArticleStatusNormal  ArticleStatus = "normal"
	ArticleStatusHidden  ArticleStatus = "hidden"
	ArticleStatusLocked  ArticleStatus = "locked"
	ArticleStatusDrafts  ArticleStatus = "drafts"
	ArticleStatusDeleted ArticleStatus = "deleted"
)

var ArticleStatusMap = enum.NewMapping[ArticleStatus, v1.ArticleStatus](map[ArticleStatus]enum.Entry[ArticleStatus, v1.ArticleStatus]{
	ArticleStatusNormal:  {Proto: v1.ArticleStatus_ARTICLE_STATUS_NORMAL},
	ArticleStatusHidden:  {Proto: v1.ArticleStatus_ARTICLE_STATUS_HIDDEN},
	ArticleStatusLocked:  {Proto: v1.ArticleStatus_ARTICLE_STATUS_LOCKED},
	ArticleStatusDrafts:  {Proto: v1.ArticleStatus_ARTICLE_STATUS_DRAFTS},
	ArticleStatusDeleted: {Proto: v1.ArticleStatus_ARTICLE_STATUS_DELETED},
})
