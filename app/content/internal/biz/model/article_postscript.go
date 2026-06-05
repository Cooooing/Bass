package model

import (
	"fmt"
	"time"

	"common/pkg/util"
	"content/internal/enum"

	"github.com/88250/lute/ast"
	"github.com/88250/lute/parse"
)

type ArticlePostscript struct {
	ID        int64
	ArticleID int64
	Content   string
	Status    enum.ArticlePostscriptStatus
	CreatedAt *time.Time
	UpdatedAt *time.Time
	CreatedBy *int64
	UpdatedBy *int64

	ContentRender string `json:"content_render"`
}

func (p *ArticlePostscript) FormatContent() {
	p.Content = util.LuteEngine.FormatStr(fmt.Sprintf("article_postscript_%d", p.ID), p.Content)
}

func (p *ArticlePostscript) ParseContent() (atUserNames map[string]struct{}) {
	atUserNames = make(map[string]struct{})
	tree := parse.Parse(fmt.Sprintf("article_postscript_%d", p.ID), []byte(p.Content), parse.NewOptions())
	ast.Walk(tree.Root, func(n *ast.Node, entering bool) ast.WalkStatus {
		return util.ParseNodeLinkAtUsernames(n, entering, atUserNames)
	})
	p.ContentRender = util.LuteEngine.MarkdownStr(fmt.Sprintf("article_postscript_%d", p.ID), p.Content)
	return atUserNames
}
