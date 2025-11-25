package model

import (
	v1 "common/api/content/v1"
	userv1 "common/api/user/v1"
	"content/internal/data/ent/gen"
	"fmt"
	"strings"
	"time"

	"github.com/88250/lute"
	"github.com/88250/lute/ast"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Article gen.Article

var luteEngine = lute.New()

// Summary 文章摘要
func (a *Article) Summary() {
	r := []rune(a.Content)
	if len(r) > 20 {
		a.Content = string(r[:20]) + "..."
	}
}

func (a *Article) luteDocName() string {
	return fmt.Sprintf("article_%d_%d_%s", a.ID, a.CreatedBy, a.Title)
}

// FormatContent 格式化文章内容
func (a *Article) FormatContent() {
	a.Content = luteEngine.FormatStr(a.luteDocName(), a.Content)
}

// ParseContent 解析文章内容
func (a *Article) ParseContent() (renderContent string, atUserNames []string) {
	luteEngine.Md2HTMLRendererFuncs[ast.NodeLink] = func(n *ast.Node, entering bool) (string, ast.WalkStatus) {
		if !entering {
			return "", ast.WalkContinue
		}
		// 链接文本
		text := n.Text() // n.Text() 会递归获取 NodeLinkText 或 TextMarkTextContent
		// 链接地址
		link := ""
		if n.IsTextMarkType("a") {
			link = n.TextMarkAHref
		} else {
			// 或者尝试找 NodeLinkDest 子节点
			destNode := n.ChildByType(ast.NodeLinkDest)
			if destNode != nil {
				link = string(destNode.Tokens)
			}
		}

		s := text[1:]
		if strings.HasPrefix(text, "@") {
			// 解析 at 用户 [@username](user's home page link)
			username := s
			// 这里可以收集用户ID或者生成特定HTML
			atUserNames = append(atUserNames, username)
			return fmt.Sprintf(`<a href="%s">%s</a>`, link, text), ast.WalkContinue
		} else if strings.HasPrefix(text, "&") {
			// 解析引用文章 [&username:title](article's link)
			parts := strings.SplitN(s, ":", 2)
			if len(parts) == 2 {
				return fmt.Sprintf(`<a href="%s">%s</a>`, link, text), ast.WalkContinue
			}
		} else if strings.HasPrefix(text, "#") {
			// 解析引用标签或领域 [#tag/domain](tag's link)
			return fmt.Sprintf(`<a href="%s">%s</a>`, link, text), ast.WalkContinue
		}

		// 默认返回原生链接
		return fmt.Sprintf(`<a href="%s">%s</a>`, link, text), ast.WalkContinue
	}
	return luteEngine.MarkdownStr(a.luteDocName(), a.Content), atUserNames
}

func (a *Article) ConvertToRpc(authorUser *userv1.User, lastReplyUser *userv1.User, repliedAt *time.Time) *v1.Article {
	article := &v1.Article{
		CreatedAt:               timestamppb.New(*a.CreatedAt),
		UpdatedAt:               timestamppb.New(*a.UpdatedAt),
		CreatedBy:               a.CreatedBy,
		UpdatedBy:               a.UpdatedBy,
		Id:                      a.ID,
		Title:                   a.Title,
		Content:                 a.Content,
		HasPostscript:           a.HasPostscript,
		RewardContent:           a.RewardContent,
		RewardPoints:            a.RewardPoints,
		Status:                  a.Status,
		Type:                    a.Type,
		Commentable:             a.Commentable,
		Anonymous:               a.Anonymous,
		ThankCount:              a.ThankCount,
		LikeCount:               a.LikeCount,
		CollectCount:            a.CollectCount,
		WatchCount:              a.WatchCount,
		ReplyCount:              a.ReplyCount,
		BountyPoints:            a.BountyPoints,
		AcceptedAnswerId:        a.AcceptedAnswerID,
		VoteTotal:               a.VoteTotal,
		LotteryParticipantCount: a.LotteryParticipantCount,
		LotteryWinnerCount:      a.LotteryWinnerCount,
		AuthorUser:              authorUser,
		LastReplyUser:           lastReplyUser,
	}
	if repliedAt != nil {
		article.LastReplyAt = timestamppb.New(*repliedAt)
	}
	entArticle := (*gen.Article)(a)
	if len(entArticle.Edges.Postscripts) > 0 {
		for _, postscript := range entArticle.Edges.Postscripts {
			article.Postscripts = append(article.Postscripts, (*ArticlePostscript)(postscript).ConvertToRpc())
		}
	}
	if len(entArticle.Edges.Tags) > 0 {
		for _, tag := range entArticle.Edges.Tags {
			article.Tags = append(article.Tags, (*Tag)(tag).ConvertToRpc())
		}
	}
	return article
}
