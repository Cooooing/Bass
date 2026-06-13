package model

import (
	"fmt"
	"time"

	"common/pkg/util"
	"content/internal/enum"
)

type ArticlePostscript struct {
	ID          int64
	ArticleID   int64
	Content     string
	Restriction enum.ContentRestriction
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	CreatedBy   *int64
	UpdatedBy   *int64
}

func (p *ArticlePostscript) FormatContent() {
	p.Content = util.LuteEngine.FormatStr(fmt.Sprintf("article_postscript_%d", p.ID), p.Content)
}
