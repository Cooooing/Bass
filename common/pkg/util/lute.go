package util

import (
	"strings"

	"github.com/88250/lute"
	"github.com/88250/lute/ast"
)

var LuteEngine = lute.New()

func ParseNodeLinkAtUsernames(
	n *ast.Node,
	entering bool,
	atUsernames map[string]struct{},
) ast.WalkStatus {
	if !entering || n.Type != ast.NodeLink {
		return ast.WalkContinue
	}

	text := n.Text()
	if len(text) == 0 {
		return ast.WalkContinue
	}

	/*
		匹配用户链接：
		@username
		[@username](用户主页链接)

		匹配文章链接：
		&username:title
		[&username:title](文章链接)

		匹配标签链接：
		#tag/domain
		[#tag/domain](标签链接)
	*/
	if strings.HasPrefix(text, "@") {
		atUsernames[text[1:]] = struct{}{}
	} else if strings.HasPrefix(text, "&") {
		parts := strings.SplitN(text[1:], ":", 2)
		if len(parts) == 2 {
			// 解析 username:title。
		}
	}

	return ast.WalkContinue
}

func ParseNodeImageCoverImageUrl(
	n *ast.Node,
	entering bool,
	coverImageUrl *string,
) ast.WalkStatus {
	if !entering || n.Type != ast.NodeImage {
		return ast.WalkContinue
	}
	// 已经找到第一张图片时直接停止遍历。
	if coverImageUrl != nil && *coverImageUrl != "" {
		return ast.WalkStop
	}
	// 解析图片 URL。
	if coverImageUrl != nil {
		dest := n.ChildByType(ast.NodeLinkDest)
		if dest != nil {
			*coverImageUrl = string(dest.Tokens)
		}
	}
	return ast.WalkContinue
}
