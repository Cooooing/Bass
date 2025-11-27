package model

import (
	v1 "common/api/content/v1"
	"common/pkg/util"
	"common/pkg/util/collections/set"
	"content/internal/data/ent/gen"
	"fmt"

	"github.com/88250/lute/ast"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ArticlePostscript struct {
	*gen.ArticlePostscript
	ContentRender string `json:"content_render"`
}

// FormatContent 格式化文章内容
func (p *ArticlePostscript) FormatContent() {
	p.Content = util.LuteEngine.FormatStr(fmt.Sprintf("article_postscript_%d", p.ID), p.Content)
}

// ParseContent 解析文章内容
func (p *ArticlePostscript) ParseContent() (atUserNames set.Set[string]) {
	atUserNames = set.New[string](0)
	util.LuteEngine.Md2HTMLRendererFuncs[ast.NodeLink] = func(n *ast.Node, entering bool) (string, ast.WalkStatus) {
		return util.ParseNodeLink(n, entering, atUserNames)
	}
	p.ContentRender = util.LuteEngine.MarkdownStr(fmt.Sprintf("article_postscript_%d", p.ID), p.Content)
	return atUserNames
}

// ConvertToRpc 转换为RPC返回格式
func (p *ArticlePostscript) ConvertToRpc() *v1.ArticlePostscript {
	p.ParseContent()
	return &v1.ArticlePostscript{
		CreatedAt:     timestamppb.New(*p.CreatedAt),
		UpdatedAt:     timestamppb.New(*p.UpdatedAt),
		CreatedBy:     p.CreatedBy,
		UpdatedBy:     p.UpdatedBy,
		Id:            p.ID,
		ArticleId:     p.ArticleID,
		Content:       p.Content,
		ContentRender: p.ContentRender,
	}
}
