package enum

import (
	"common/pkg/enum"
	v1 "common/proto/gen/content/v1/enum"
)

// ArticlePublishStatus 表示文章发布生命周期状态。
type ArticlePublishStatus string

const (
	// ArticlePublishStatusDraft 表示草稿。
	ArticlePublishStatusDraft ArticlePublishStatus = "draft"
	// ArticlePublishStatusPublished 表示已发布。
	ArticlePublishStatusPublished ArticlePublishStatus = "published"
	// ArticlePublishStatusArchived 表示已归档。
	ArticlePublishStatusArchived ArticlePublishStatus = "archived"
	// ArticlePublishStatusScheduled 表示已定时发布。
	ArticlePublishStatusScheduled ArticlePublishStatus = "scheduled"
)

// ArticlePublishStatusMap 维护文章发布状态内部枚举与 proto 枚举的映射。
var ArticlePublishStatusMap = enum.NewMapping[ArticlePublishStatus, v1.ArticlePublishStatus](map[ArticlePublishStatus]enum.Entry[ArticlePublishStatus, v1.ArticlePublishStatus]{
	ArticlePublishStatusDraft:     {Proto: v1.ArticlePublishStatus_ARTICLE_PUBLISH_STATUS_DRAFT},
	ArticlePublishStatusPublished: {Proto: v1.ArticlePublishStatus_ARTICLE_PUBLISH_STATUS_PUBLISHED},
	ArticlePublishStatusArchived:  {Proto: v1.ArticlePublishStatus_ARTICLE_PUBLISH_STATUS_ARCHIVED},
	ArticlePublishStatusScheduled: {Proto: v1.ArticlePublishStatus_ARTICLE_PUBLISH_STATUS_SCHEDULED},
})

// ArticleVisibility 表示文章可见范围。
type ArticleVisibility string

const (
	// ArticleVisibilityPublic 表示公开可见。
	ArticleVisibilityPublic ArticleVisibility = "public"
	// ArticleVisibilityPrivate 表示私有可见。
	ArticleVisibilityPrivate ArticleVisibility = "private"
)

// ArticleVisibilityMap 维护文章可见范围内部枚举与 proto 枚举的映射。
var ArticleVisibilityMap = enum.NewMapping[ArticleVisibility, v1.ArticleVisibility](map[ArticleVisibility]enum.Entry[ArticleVisibility, v1.ArticleVisibility]{
	ArticleVisibilityPublic:  {Proto: v1.ArticleVisibility_ARTICLE_VISIBILITY_PUBLIC},
	ArticleVisibilityPrivate: {Proto: v1.ArticleVisibility_ARTICLE_VISIBILITY_PRIVATE},
})

func (e ArticlePublishStatus) String() string {
	return string(e)
}

func (e ArticleVisibility) String() string {
	return string(e)
}
