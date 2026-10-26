package enum

import (
	"common/pkg/enum"
	v1 "common/proto/gen/content/v1/enum"
)

// ArticlePublishStatus 表示文章发布生命周期状态。
type ArticlePublishStatus string

const (
	ArticlePublishStatusDraft     ArticlePublishStatus = "draft"     // ArticlePublishStatusDraft 表示草稿。
	ArticlePublishStatusPublished ArticlePublishStatus = "published" // ArticlePublishStatusPublished 表示已发布。
	ArticlePublishStatusArchived  ArticlePublishStatus = "archived"  // ArticlePublishStatusArchived 表示已归档。
	ArticlePublishStatusPending   ArticlePublishStatus = "pending"   // ArticlePublishStatusPending 表示正在确认发布事务。
)

// ArticlePublishStatusMap 维护文章发布状态内部枚举与 proto 枚举的映射。
var ArticlePublishStatusMap = enum.NewMapping[ArticlePublishStatus, v1.ArticlePublishStatus](map[ArticlePublishStatus]enum.Entry[ArticlePublishStatus, v1.ArticlePublishStatus]{
	ArticlePublishStatusDraft:     {Proto: v1.ArticlePublishStatus_ARTICLE_PUBLISH_STATUS_DRAFT},
	ArticlePublishStatusPublished: {Proto: v1.ArticlePublishStatus_ARTICLE_PUBLISH_STATUS_PUBLISHED},
	ArticlePublishStatusArchived:  {Proto: v1.ArticlePublishStatus_ARTICLE_PUBLISH_STATUS_ARCHIVED},
	ArticlePublishStatusPending:   {Proto: v1.ArticlePublishStatus_ARTICLE_PUBLISH_STATUS_PENDING},
})

func (e ArticlePublishStatus) String() string {
	return string(e)
}
