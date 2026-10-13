package model

import (
	"time"

	"content/internal/enum"
)

type ArticleFilter struct {
	TagID             *int64
	DomainID          *int64
	ArticleID         *int64
	ArticleIDs        []int64
	PublishStatus     *enum.ArticlePublishStatus
	PublishStatuses   []enum.ArticlePublishStatus
	Visibility        *enum.ArticleVisibility
	Visibilities      []enum.ArticleVisibility
	Restriction       *enum.ContentRestriction
	Restrictions      []enum.ContentRestriction
	AuthorID          *int64
	AuthorIDs         []int64
	ExcludedAuthorIDs []int64
	PublishedCity     *string
	Order             *enum.ArticleOrder
	Type              *enum.ArticleType
	PublishedAtEnd    *time.Time
	Scheduled         *bool
}
