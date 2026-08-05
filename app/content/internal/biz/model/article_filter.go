package model

import (
	"content/internal/enum"
	"time"
)

type ArticleFilter struct {
	TagID           *int64
	DomainID        *int64
	ArticleID       *int64
	ArticleIDs      []int64
	PublishStatus   *enum.ArticlePublishStatus
	PublishStatuses []enum.ArticlePublishStatus
	Visibility      *enum.ArticleVisibility
	Visibilities    []enum.ArticleVisibility
	Restriction     *enum.ContentRestriction
	Restrictions    []enum.ContentRestriction
	AuthorID        *int64
	Order           *enum.ArticleOrder
	Type            *enum.ArticleType
	Keyword         *string
	PublishedAtEnd  *time.Time
}
