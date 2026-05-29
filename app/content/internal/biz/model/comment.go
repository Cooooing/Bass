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

type Comment struct {
	ID         int64
	ArticleID  int64
	Content    string
	Level      int32
	ParentID   *int64
	ReplyID    *int64
	Status     enum.CommentStatus
	ThankCount int32
	LikeCount  int32
	ReplyCount int32
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
	CreatedBy  *int64
	UpdatedBy  *int64

	Article *Article
	Reply   *Comment

	ContentRender string `json:"content_render"`

	User      *userv1.AccountBasic `json:"user"`
	ReplyUser *userv1.AccountBasic `json:"reply_user"`
}

func (c *Comment) FormatContent() {
	c.Content = util.LuteEngine.MarkdownStr(fmt.Sprintf("comment_%d", c.ID), c.Content)
}

func (c *Comment) ParseContent() (atUserNames map[string]struct{}) {
	atUserNames = make(map[string]struct{})
	tree := parse.Parse(fmt.Sprintf("article_postscript_%d", c.ID), []byte(c.Content), parse.NewOptions())
	ast.Walk(tree.Root, func(n *ast.Node, entering bool) ast.WalkStatus {
		return util.ParseNodeLinkAtUsernames(n, entering, atUserNames)
	})
	c.ContentRender = util.LuteEngine.MarkdownStr(fmt.Sprintf("article_postscript_%d", c.ID), c.Content)
	return atUserNames
}
