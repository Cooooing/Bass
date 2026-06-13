package model

import (
	"fmt"
	"time"

	"common/pkg/util"
	"content/internal/enum"
)

type Comment struct {
	ID          int64
	ArticleID   int64
	Content     string
	Level       int32
	ParentID    *int64
	ReplyID     *int64
	Restriction enum.ContentRestriction
	ThankCount  int32
	LikeCount   int32
	ReplyCount  int32
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	CreatedBy   *int64
	UpdatedBy   *int64
	DeletedAt   *time.Time

	ReplyUserID *int64 `json:"reply_user_id"`
}

func (c *Comment) FormatContent() {
	c.Content = util.LuteEngine.FormatStr(fmt.Sprintf("comment_%d", c.ID), c.Content)
}
