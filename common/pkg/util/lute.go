package util

import (
	"common/pkg/cutil/collections/set"
	"strings"

	"github.com/88250/lute"
	"github.com/88250/lute/ast"
)

var LuteEngine = lute.New()

func ParseNodeLinkAtUsernames(n *ast.Node, entering bool, atUsernames set.Set[string]) ast.WalkStatus {
	if !entering || n.Type != ast.NodeLink {
		return ast.WalkContinue
	}

	text := n.Text()

	/*
		@username
		[@username](user's home page link)

		&username:title
		[&username:title](article's link)

		#tag/domain
		[#tag/domain](tag's link)
	*/
	s := text[1:]
	if strings.HasPrefix(text, "@") {
		username := s
		atUsernames.Add(username)
		return ast.WalkContinue
	} else if strings.HasPrefix(text, "&") {
		parts := strings.SplitN(s, ":", 2)
		if len(parts) == 2 {
			return ast.WalkContinue
		}
	} else if strings.HasPrefix(text, "#") {
		return ast.WalkContinue
	}

	return ast.WalkContinue
}

func ParseNodeImageCoverImageUrl(n *ast.Node, entering bool, coverImageUrl *string) ast.WalkStatus {
	if !entering || n.Type != ast.NodeImage {
		return ast.WalkContinue
	}
	// 如果已经有第一张图片了，跳过
	if coverImageUrl != nil && *coverImageUrl != "" {
		return ast.WalkStop // 停止遍历，后面的图片不会再处理
	}
	// 解析图片 URL
	if coverImageUrl != nil {
		dest := n.ChildByType(ast.NodeLinkDest)
		if dest != nil {
			*coverImageUrl = string(dest.Tokens)
		}
	}
	return ast.WalkContinue
}
