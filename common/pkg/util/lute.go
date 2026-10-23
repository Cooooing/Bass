package util

import (
	"strings"

	"github.com/88250/lute"
	"github.com/88250/lute/ast"
	"github.com/88250/lute/parse"
)

var LuteEngine = lute.New()

func ParseAtUsernames(markdown string) []string {
	if markdown == "" {
		return nil
	}
	usernames := make(map[string]struct{})
	tree := parse.Parse("at_usernames", []byte(markdown), parse.NewOptions())
	ast.Walk(tree.Root, func(n *ast.Node, entering bool) ast.WalkStatus {
		if !entering || n.Type != ast.NodeLink || usernames == nil {
			return ast.WalkContinue
		}
		text := n.Text()
		if !strings.HasPrefix(text, "@") || len(text) == 1 {
			return ast.WalkContinue
		}
		usernames[text[1:]] = struct{}{}
		return ast.WalkContinue
	})
	result := make([]string, 0, len(usernames))
	for username := range usernames {
		result = append(result, username)
	}
	return result
}

func ParseNodeImageCoverImageUrl(n *ast.Node, entering bool, coverImageUrl *string) ast.WalkStatus {
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
