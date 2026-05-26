package model

import (
	"fmt"
	"time"

	userv1 "common/api/gen/user/v1"
	"common/pkg/util"
	"content/internal/enum"

	"github.com/88250/lute/ast"
	"github.com/88250/lute/parse"
)

type Article struct {
	ID               int64
	Title            string
	Content          string
	HasPostscript    bool
	RewardContent    *string
	RewardPoints     *int32
	Status           enum.ArticleStatus
	Type             enum.ArticleType
	Statement        *string
	Commentable      bool
	Anonymous        bool
	Listable         bool
	ViewCount        int32
	ThankCount       int32
	LikeCount        int32
	CollectCount     int32
	WatchCount       int32
	ReplyCount       int32
	BountyPoints     *int32
	AcceptedAnswerID *int64
	CreatedAt        *time.Time
	UpdatedAt        *time.Time
	CreatedBy        *int64
	UpdatedBy        *int64

	Postscripts   []*ArticlePostscript
	Tags          []*Tag
	ActionRecords []*ArticleActionRecord

	ContentRender       string  `json:"content_render"`
	RewardContentRender *string `json:"reward_content_render"`
	CoverImageUrl       *string `json:"cover_image_url"`

	AuthorUser           *userv1.AccountBasic `json:"author_user"`
	LastReplyCommentUser *userv1.AccountBasic `json:"last_reply_user"`
	LastReplyCommentAt   *time.Time           `json:"last_replied_at"`

	IsSummary bool `json:"-"`
}

func (a *Article) Summary() {
	r := []rune(a.Content)
	if len(r) > 200 {
		a.Content = string(r[:200]) + "..."
	}
}

func (a *Article) FormatContent() {
	a.Content = util.LuteEngine.FormatStr(fmt.Sprintf("%s_%d", "article_content", a.ID), a.Content)
}

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

func (a *Article) FormatRewardContent() {
	a.Content = util.LuteEngine.FormatStr(fmt.Sprintf("%s_%d", "article_reward_content", a.ID), a.Content)
}

func (a *Article) ParseRewardContent() (atUserNames map[string]struct{}) {
	atUserNames = make(map[string]struct{})
	if a.RewardContent != nil && len(a.ActionRecords) > 0 && a.HasRewarded() {
		tree := parse.Parse(fmt.Sprintf("%s_%d", "article_reward_content", a.ID), []byte(*a.RewardContent), parse.NewOptions())
		ast.Walk(tree.Root, func(n *ast.Node, entering bool) ast.WalkStatus {
			return util.ParseNodeLinkAtUsernames(n, entering, atUserNames)
		})
		a.RewardContentRender = new(util.LuteEngine.MarkdownStr(fmt.Sprintf("%s_%d", "article_reward_content", a.ID), *a.RewardContent))
	}
	return atUserNames
}

func (a *Article) HasRewarded() bool {
	for _, record := range a.ActionRecords {
		if record.Type == enum.ArticleActionReward {
			return true
		}
	}
	return false
}
