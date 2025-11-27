package model

import (
	v1 "common/api/content/v1"
	userv1 "common/api/user/v1"
	"common/pkg/util"
	"common/pkg/util/collections/set"
	"content/internal/data/ent/gen"
	"fmt"
	"time"

	"github.com/88250/lute/ast"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Article struct {
	*gen.Article
	ContentRender       string `json:"content_render"`
	RewardContentRender string `json:"reward_content_render"`

	AuthorUser           *userv1.User `json:"author_user"`
	LastReplyCommentUser *userv1.User `json:"last_reply_user"`
	LastReplyCommentAt   *time.Time   `json:"last_replied_at"`
}

func NewArticle(model *gen.Article) *Article {
	a := &Article{Article: model}
	return a
}

// Summary 文章摘要
func (a *Article) Summary() {
	r := []rune(a.Content)
	if len(r) > 20 {
		a.Content = string(r[:20]) + "..."
	}
}

// FormatContent 格式化文章内容
func (a *Article) FormatContent() {
	a.Content = util.LuteEngine.FormatStr(fmt.Sprintf("%s_%d", "article_content", a.ID), a.Content)
}

// ParseContent 解析文章内容
func (a *Article) ParseContent() (atUserNames set.Set[string]) {
	atUserNames = set.New[string](0)
	util.LuteEngine.Md2HTMLRendererFuncs[ast.NodeLink] = func(n *ast.Node, entering bool) (string, ast.WalkStatus) {
		return util.ParseNodeLink(n, entering, atUserNames)
	}
	a.ContentRender = util.LuteEngine.MarkdownStr(fmt.Sprintf("%s_%d", "article_content", a.ID), a.Content)
	return atUserNames
}

// FormatRewardContent 格式化文章打赏区内容
func (a *Article) FormatRewardContent() {
	a.Content = util.LuteEngine.FormatStr(fmt.Sprintf("%s_%d", "article_reward_content", a.ID), a.Content)
}

// ParseRewardContent 解析文章打赏区内容
func (a *Article) ParseRewardContent() (atUserNames set.Set[string]) {
	atUserNames = set.New[string](0)
	if a.RewardContent != nil {
		util.LuteEngine.Md2HTMLRendererFuncs[ast.NodeLink] = func(n *ast.Node, entering bool) (string, ast.WalkStatus) {
			return util.ParseNodeLink(n, entering, atUserNames)
		}
		a.RewardContentRender = util.LuteEngine.MarkdownStr(fmt.Sprintf("%s_%d", "article_reward_content", a.ID), *a.RewardContent)
	}
	return atUserNames
}

// ConvertToRpc 转换为RPC返回格式
func (a *Article) ConvertToRpc() *v1.Article {
	a.ParseContent()
	a.ParseRewardContent()
	article := &v1.Article{
		CreatedAt:               timestamppb.New(*a.CreatedAt),
		UpdatedAt:               timestamppb.New(*a.UpdatedAt),
		CreatedBy:               a.CreatedBy,
		UpdatedBy:               a.UpdatedBy,
		Id:                      a.ID,
		Title:                   a.Title,
		Content:                 a.Content,
		ContentRender:           a.ContentRender,
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
		AuthorUser:              a.AuthorUser,
		LastReplyUser:           a.LastReplyCommentUser,
	}
	if a.LastReplyCommentAt != nil {
		article.LastReplyAt = timestamppb.New(*a.LastReplyCommentAt)
	}
	if len(a.Edges.Postscripts) > 0 {
		for _, postscript := range a.Edges.Postscripts {
			article.Postscripts = append(article.Postscripts, (&ArticlePostscript{ArticlePostscript: postscript}).ConvertToRpc())
		}
	}
	if len(a.Edges.Tags) > 0 {
		for _, tag := range a.Edges.Tags {
			article.Tags = append(article.Tags, (*Tag)(tag).ConvertToRpc())
		}
	}
	return article
}
