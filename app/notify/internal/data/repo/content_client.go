package repo

import (
	"common/api/gen/common"
	contentv1 "common/api/gen/content/v1"
	"common/pkg/client/rpc"
	"context"
	"notify/internal/biz/model"
	bizrepo "notify/internal/biz/repo"
)

var _ bizrepo.ContentClient = (*ContentClient)(nil)

type ContentClient struct {
	contentClient *rpc.ContentClient
}

func NewContentClient(contentClient *rpc.ContentClient) bizrepo.ContentClient {
	return &ContentClient{contentClient: contentClient}
}

func (c *ContentClient) GetArticle(ctx context.Context, articleID int64) (*model.ContentArticle, error) {
	reply, err := c.contentClient.Article.Get(ctx, &contentv1.GetArticle_Request{ArticleId: articleID})
	if err != nil {
		return nil, err
	}
	article := reply.GetArticle()
	if article == nil {
		return nil, nil
	}
	author := article.GetAuthorUser()
	result := &model.ContentArticle{
		ID:    article.GetId(),
		Title: article.GetTitle(),
	}
	if author != nil {
		result.AuthorID = author.GetId()
		result.AuthorName = author.GetName()
		result.AuthorNickname = author.GetNickname()
	}
	if result.AuthorID == 0 {
		result.AuthorID = article.GetCreatedBy()
	}
	return result, nil
}

func (c *ContentClient) GetComment(ctx context.Context, commentID int64) (*model.ContentComment, error) {
	reply, err := c.contentClient.Comment.List(ctx, &contentv1.ListComments_Request{
		Page: &common.PageRequest{Page: 1, Size: 1},
		Query: &contentv1.CommentQueryParams{
			CommentId: new(commentID),
		},
	})
	if err != nil {
		return nil, err
	}
	if len(reply.GetRows()) == 0 {
		return nil, nil
	}
	comment := reply.GetRows()[0]
	result := &model.ContentComment{
		ID:        comment.GetId(),
		ArticleID: comment.GetArticleId(),
		Content:   comment.GetContent(),
	}
	if user := comment.GetUser(); user != nil {
		result.UserID = user.GetId()
		result.UserName = user.GetName()
		result.UserNickname = user.GetNickname()
	}
	if result.UserID == 0 {
		result.UserID = comment.GetCreatedBy()
	}
	if replyUser := comment.GetReplyUser(); replyUser != nil {
		result.ReplyUserID = replyUser.GetId()
		result.ReplyUserName = replyUser.GetName()
	}
	if result.ArticleID != 0 {
		article, err := c.GetArticle(ctx, result.ArticleID)
		if err != nil {
			return nil, err
		}
		result.Article = article
	}
	return result, nil
}
