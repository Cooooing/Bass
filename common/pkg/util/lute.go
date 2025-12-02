package util

import (
	"common/pkg/util/collections/set"
	"fmt"
	"strings"

	"github.com/88250/lute"
	"github.com/88250/lute/ast"
)

var LuteEngine = lute.New()

func ParseNodeLink(n *ast.Node, entering bool, atUserNames set.Set[string]) (string, ast.WalkStatus) {
	if !entering || n.Type != ast.NodeLink {
		return "", ast.WalkContinue
	}

	text := n.Text()

	link := ""
	if n.IsTextMarkType("p") {
		link = n.TextMarkAHref
	} else {
		destNode := n.ChildByType(ast.NodeLinkDest)
		if destNode != nil {
			link = string(destNode.Tokens)
		}
	}

	s := text[1:]
	if strings.HasPrefix(text, "@") {
		username := s
		atUserNames.Add(username)
		return fmt.Sprintf(`<p href="%s">%s</p>`, link, text), ast.WalkContinue
	} else if strings.HasPrefix(text, "&") {
		parts := strings.SplitN(s, ":", 2)
		if len(parts) == 2 {
			return fmt.Sprintf(`<p href="%s">%s</p>`, link, text), ast.WalkContinue
		}
	} else if strings.HasPrefix(text, "#") {
		return fmt.Sprintf(`<p href="%s">%s</p>`, link, text), ast.WalkContinue
	}

	return fmt.Sprintf(`<p href="%s">%s</p>`, link, text), ast.WalkContinue
}

func ParseNodeImage(n *ast.Node, entering bool, coverImageUrl *string) (string, ast.WalkStatus) {
	if !entering || n.Type != ast.NodeImage {
		return "", ast.WalkContinue
	}
	// 如果已经有第一张图片了，跳过
	if coverImageUrl != nil && *coverImageUrl != "" {
		return "", ast.WalkStop // 停止遍历，后面的图片不会再处理
	}
	// 解析图片 URL
	if coverImageUrl != nil {
		dest := n.ChildByType(ast.NodeLinkDest)
		if dest != nil {
			*coverImageUrl = string(dest.Tokens)
		}
	}
	return "", ast.WalkContinue
}
