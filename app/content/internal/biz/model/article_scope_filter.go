package model

import "content/internal/enum"

type ArticleScopeFilter struct {
	AuthorID          *int64
	PublishStatus     *enum.ArticlePublishStatus
	PublishStatuses   []enum.ArticlePublishStatus
	Visibility        *enum.ArticleVisibility
	Visibilities      []enum.ArticleVisibility
	Restriction       *enum.ContentRestriction
	Restrictions      []enum.ContentRestriction
	PublicVisibleOnly bool
}
