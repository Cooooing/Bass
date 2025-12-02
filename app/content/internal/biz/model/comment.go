package model

import (
	v1 "common/api/content/v1"
	userv1 "common/api/user/v1"
	"common/pkg/util"
	"common/pkg/util/collections/set"
	"content/internal/data/ent/gen"
	"fmt"

	"github.com/88250/lute/ast"
	"github.com/88250/lute/parse"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Comment struct {
	*gen.Comment
	ContentRender string `json:"content_render"`

	User      *userv1.User `json:"user"`
	ReplyUser *userv1.User `json:"reply_user"`
}

// FormatContent 格式化评论内容
func (c *Comment) FormatContent() {
	c.Content = util.LuteEngine.MarkdownStr(fmt.Sprintf("comment_%d", c.ID), c.Content)
}

// ParseContent 解析评论内容
func (c *Comment) ParseContent() (atUserNames set.Set[string]) {
	atUserNames = set.New[string](0)
	tree := parse.Parse(fmt.Sprintf("article_postscript_%d", c.ID), []byte(c.Content), parse.NewOptions())
	ast.Walk(tree.Root, func(n *ast.Node, entering bool) ast.WalkStatus {
		return util.ParseNodeLinkAtUsernames(n, entering, atUserNames)
	})
	c.ContentRender = util.LuteEngine.MarkdownStr(fmt.Sprintf("article_postscript_%d", c.ID), c.Content)
	return atUserNames
}

func (c *Comment) ConvertToRpc() *v1.Comment {
	c.ParseContent()
	comment := &v1.Comment{
		CreatedAt:     timestamppb.New(*c.CreatedAt),
		UpdatedAt:     timestamppb.New(*c.UpdatedAt),
		CreatedBy:     c.CreatedBy,
		UpdatedBy:     c.UpdatedBy,
		Id:            c.ID,
		ArticleId:     c.ArticleID,
		Content:       c.Content,
		ContentRender: c.ContentRender,
		Level:         c.Level,
		ParentId:      c.ParentID,
		ReplyId:       c.ReplyID,
		Status:        c.Status,
		ThankCount:    c.ThankCount,
		LikeCount:     c.LikeCount,
		CollectCount:  c.CollectCount,
		ReplyCount:    c.ReplyCount,
		User:          c.User,
		ReplyUser:     c.ReplyUser,
	}
	return comment
}
