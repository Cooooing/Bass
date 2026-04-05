package model

import (
	v1 "common/gen/content/v1"
	userv1 "common/gen/user/v1"
	"common/pkg/util"
	"content/internal/data/ent/gen"
	"fmt"
	"time"

	"github.com/88250/lute/ast"
	"github.com/88250/lute/parse"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Article struct {
	*gen.Article
	ContentRender       string  `json:"content_render"`
	RewardContentRender *string `json:"reward_content_render"`
	CoverImageUrl       *string `json:"cover_image_url"`

	AuthorUser           *userv1.User `json:"author_user"`
	LastReplyCommentUser *userv1.User `json:"last_reply_user"`
	LastReplyCommentAt   *time.Time   `json:"last_replied_at"`

	// option
	IsSummary bool `json:"-"`
}

// Summary 文章摘要
func (a *Article) Summary() {
	r := []rune(a.Content)
	if len(r) > 200 {
		a.Content = string(r[:200]) + "..."
	}
}

// FormatContent 格式化文章内容
func (a *Article) FormatContent() {
	a.Content = util.LuteEngine.FormatStr(fmt.Sprintf("%s_%d", "article_content", a.ID), a.Content)
}

// ParseContent 解析文章内容
func (a *Article) ParseContent() (atUserNames map[string]struct{}) {
	atUserNames = make(map[string]struct{})
	tree := parse.Parse(fmt.Sprintf("%s_%d", "article_content", a.ID), []byte(a.Content), parse.NewOptions())
	ast.Walk(tree.Root, func(n *ast.Node, entering bool) ast.WalkStatus {
		return util.ParseNodeLinkAtUsernames(n, entering, atUserNames)
	})
	ast.Walk(tree.Root, func(n *ast.Node, entering bool) ast.WalkStatus {
		coverImageUrl := new("")
		status := util.ParseNodeImageCoverImageUrl(n, entering, coverImageUrl)
		if *coverImageUrl != "" {
			a.CoverImageUrl = coverImageUrl
		}
		return status
	})
	a.ContentRender = util.LuteEngine.MarkdownStr(fmt.Sprintf("%s_%d", "article_content", a.ID), a.Content)
	return atUserNames
}

// FormatRewardContent 格式化文章打赏区内容
func (a *Article) FormatRewardContent() {
	a.Content = util.LuteEngine.FormatStr(fmt.Sprintf("%s_%d", "article_reward_content", a.ID), a.Content)
}

// ParseRewardContent 解析文章打赏区内容
func (a *Article) ParseRewardContent() (atUserNames map[string]struct{}) {
	atUserNames = make(map[string]struct{})
	if a.RewardContent != nil && len(a.Edges.ActionRecords) > 0 && a.HasRewarded() {
		tree := parse.Parse(fmt.Sprintf("%s_%d", "article_reward_content", a.ID), []byte(*a.RewardContent), parse.NewOptions())
		ast.Walk(tree.Root, func(n *ast.Node, entering bool) ast.WalkStatus {
			return util.ParseNodeLinkAtUsernames(n, entering, atUserNames)
		})
		a.RewardContentRender = new(util.LuteEngine.MarkdownStr(fmt.Sprintf("%s_%d", "article_reward_content", a.ID), *a.RewardContent))
	}
	return atUserNames
}

// HasRewarded 判断文章是否打赏过，需要查询时使用 WithActionRecords 并按 UserId 过滤
func (a *Article) HasRewarded() bool {
	for _, record := range a.Edges.ActionRecords {
		if record.Type == int32(v1.ArticleAction_ARTICLE_ACTION_REWARD) {
			return true
		}
	}
	return false
}

// ConvertToRpc 转换为RPC返回格式
func (a *Article) ConvertToRpc() *v1.Article {
	a.ParseContent()
	a.ParseRewardContent()
	if a.IsSummary {
		a.Summary()
	}
	article := &v1.Article{
		CreatedAt:               timestamppb.New(*a.CreatedAt),
		UpdatedAt:               timestamppb.New(*a.UpdatedAt),
		CreatedBy:               a.CreatedBy,
		UpdatedBy:               a.UpdatedBy,
		CreatedByName:           a.CreatedByName,
		UpdatedByName:           a.UpdatedByName,
		Id:                      a.ID,
		Title:                   a.Title,
		Content:                 a.Content,
		ContentRender:           a.ContentRender,
		RewardContent:           a.RewardContent,
		RewardContentRender:     a.RewardContentRender,
		HasPostscript:           a.HasPostscript,
		HasReward:               util.IsNotNil(a.RewardPoints),
		RewardPoints:            a.RewardPoints,
		Status:                  a.Status,
		Type:                    a.Type,
		Statement:               a.Statement,
		Commentable:             a.Commentable,
		Anonymous:               a.Anonymous,
		ViewCount:               a.ViewCount,
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
		CoverImageUrl:           a.CoverImageUrl,
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
			article.Tags = append(article.Tags, (&Tag{Tag: tag}).ConvertToRpc())
		}
	}
	return article
}
