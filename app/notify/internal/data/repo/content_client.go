package repo

import (
	"common/pkg/client/rpc"
	"common/proto/gen/common"
	contentv1 "common/proto/gen/content/v1"
	userv1 "common/proto/gen/user/v1"
	"context"
	"notify/internal/biz/model"
	bizrepo "notify/internal/biz/repo"
)

var _ bizrepo.ContentClient = (*ContentClient)(nil)

type ContentClient struct {
	contentClient *rpc.ContentClient
	userClient    *rpc.UserClient
}

func NewContentClient(contentClient *rpc.ContentClient, userClient *rpc.UserClient) bizrepo.ContentClient {
	return &ContentClient{contentClient: contentClient, userClient: userClient}
}

func (c *ContentClient) GetArticle(ctx context.Context, req *bizrepo.ContentGetArticleReq) (*bizrepo.ContentGetArticleResponse, error) {
	if req == nil || req.ArticleID == 0 {
		return &bizrepo.ContentGetArticleResponse{}, nil
	}
	reply, err := c.contentClient.Article.Get(ctx, &contentv1.GetArticle_Request{ArticleId: req.ArticleID})
	if err != nil {
		return nil, err
	}
	article := reply.GetArticle()
	if article == nil {
		return &bizrepo.ContentGetArticleResponse{}, nil
	}
	result := &model.ContentArticle{
		ID:       article.GetId(),
		Title:    article.GetTitle(),
		AuthorID: article.GetCreatedBy(),
	}
	accounts, err := c.mapAccounts(ctx, []int64{result.AuthorID})
	if err != nil {
		return nil, err
	}
	if author := accounts[result.AuthorID]; author != nil {
		result.AuthorName = author.Name
		result.AuthorNickname = author.Nickname
	}
	return &bizrepo.ContentGetArticleResponse{Article: result}, nil
}

func (c *ContentClient) GetComment(ctx context.Context, req *bizrepo.ContentGetCommentReq) (*bizrepo.ContentGetCommentResponse, error) {
	if req == nil || req.CommentID == 0 {
		return &bizrepo.ContentGetCommentResponse{}, nil
	}
	reply, err := c.contentClient.Comment.Page(ctx, &contentv1.PageComments_Request{
		Page: &common.PageRequest{Page: 1, Size: 1},
		Query: &contentv1.PageComments_Request_CommentQueryParams{
			CommentId: new(req.CommentID),
		},
	})
	if err != nil {
		return nil, err
	}
	if len(reply.GetRows()) == 0 {
		return &bizrepo.ContentGetCommentResponse{}, nil
	}
	comment := reply.GetRows()[0]
	result := &model.ContentComment{
		ID:          comment.GetId(),
		ArticleID:   comment.GetArticleId(),
		Content:     comment.GetContent(),
		UserID:      comment.GetCreatedBy(),
		ReplyUserID: comment.GetReplyUserId(),
	}
	accounts, err := c.mapAccounts(ctx, []int64{result.UserID, result.ReplyUserID})
	if err != nil {
		return nil, err
	}
	if user := accounts[result.UserID]; user != nil {
		result.UserName = user.Name
		result.UserNickname = user.Nickname
	}
	if replyUser := accounts[result.ReplyUserID]; replyUser != nil {
		result.ReplyUserName = replyUser.Name
	}
	if result.ArticleID != 0 {
		articleResponse, err := c.GetArticle(ctx, &bizrepo.ContentGetArticleReq{ArticleID: result.ArticleID})
		if err != nil {
			return nil, err
		}
		result.Article = articleResponse.Article
	}
	return &bizrepo.ContentGetCommentResponse{Comment: result}, nil
}

func (c *ContentClient) mapAccounts(ctx context.Context, userIDs []int64) (map[int64]*model.UserAccount, error) {
	ids := make([]int64, 0, len(userIDs))
	seen := make(map[int64]struct{}, len(userIDs))
	for _, userID := range userIDs {
		if userID == 0 {
			continue
		}
		if _, ok := seen[userID]; ok {
			continue
		}
		seen[userID] = struct{}{}
		ids = append(ids, userID)
	}
	if len(ids) == 0 {
		return map[int64]*model.UserAccount{}, nil
	}
	reply, err := c.userClient.Account.Map(ctx, &userv1.MapAccounts_Request{
		Query: &userv1.MapAccounts_Request_AccountQuery{UserIds: ids},
	})
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*model.UserAccount, len(reply.GetAccounts()))
	for userID, account := range reply.GetAccounts() {
		basic := account.GetBasic()
		if basic == nil {
			continue
		}
		result[userID] = &model.UserAccount{
			ID:       userID,
			Name:     basic.GetName(),
			Nickname: basic.GetNickname(),
		}
	}
	return result, nil
}
