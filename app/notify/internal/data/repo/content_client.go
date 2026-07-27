package repo

import (
	"common/pkg/client/rpc"
	"common/proto/gen/common"
	contentv1 "common/proto/gen/content/v1"
	"context"
	"notify/internal/biz/model"
	bizrepo "notify/internal/biz/repo"
)

var _ bizrepo.ContentClient = (*ContentClient)(nil)

type ContentClient struct {
	contentClient *rpc.ContentClient
}

func NewContentClient(
	contentClient *rpc.ContentClient,
) bizrepo.ContentClient {
	return &ContentClient{
		contentClient: contentClient,
	}
}

func (c *ContentClient) GetArticle(ctx context.Context, articleID int64) (*model.ContentArticle, error) {
	if articleID == 0 {
		return nil, nil
	}
	reply, err := c.contentClient.Article.Get(ctx, &contentv1.GetArticle_Req{
		ArticleId: articleID,
	})
	if err != nil {
		return nil, err
	}
	article := reply.GetArticle()
	if article == nil {
		return nil, nil
	}
	result := &model.ContentArticle{
		ID:       article.GetId(),
		Title:    article.GetTitle(),
		AuthorID: article.GetCreatedBy(),
	}
	return result, nil
}

func (c *ContentClient) GetComment(ctx context.Context, commentID int64) (*model.ContentComment, error) {
	if commentID == 0 {
		return nil, nil
	}
	reply, err := c.contentClient.Comment.Page(ctx, &contentv1.PageComments_Req{
		Page: &common.PageReq{
			Page: 1,
			Size: 1,
		},
		Query: &contentv1.PageComments_Req_CommentQueryParams{
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
		ID:          comment.GetId(),
		ArticleID:   comment.GetArticleId(),
		Content:     comment.GetContent(),
		UserID:      comment.GetCreatedBy(),
		ReplyUserID: comment.GetReplyUserId(),
	}
	if result.ArticleID != 0 {
		articleResp, err := c.GetArticle(ctx, result.ArticleID)
		if err != nil {
			return nil, err
		}
		result.Article = articleResp
	}
	return result, nil
}
