package model

import "content/internal/enum"

type CommentScopeFilter struct {
	CreatedBy            *int64
	Restriction          *enum.ContentRestriction
	Restrictions         []enum.ContentRestriction
	ArticlePublicVisible bool
}
