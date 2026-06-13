package model

import (
	"fmt"
	"time"

	"common/pkg/util"
	"content/internal/enum"
)

type Article struct {
	ID               int64
	Title            string
	Content          string
	HasPostscript    bool
	RewardContent    *string
	RewardPoints     *int32
	PublishStatus    enum.ArticlePublishStatus
	Visibility       enum.ArticleVisibility
	Restriction      enum.ContentRestriction
	Type             enum.ArticleType
	Statement        *string
	Commentable      bool
	Anonymous        bool
	PublishedAt      *time.Time
	EditedAt         *time.Time
	ViewCount        int32
	ThankCount       int32
	LikeCount        int32
	CollectCount     int32
	WatchCount       int32
	ReplyCount       int32
	BountyPoints     *int32
	AcceptedAnswerID *int64
	CreatedAt        *time.Time
	UpdatedAt        *time.Time
	CreatedBy        *int64
	UpdatedBy        *int64
	DeletedAt        *time.Time
}

func (a *Article) FormatContent() {
	a.Content = util.LuteEngine.FormatStr(fmt.Sprintf("%s_%d", "article_content", a.ID), a.Content)
}
