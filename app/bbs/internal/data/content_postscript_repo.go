package data

import (
	"bbs/internal/biz/repo"
	"common/pkg/client/rpc"
	"common/pkg/util"
	contentv1 "common/proto/gen/content/v1"
	"context"
	"fmt"
)

var _ repo.ContentPostscriptClient = (*ContentPostscriptClient)(nil)

type ContentPostscriptClient struct {
	contentClient *rpc.ContentClient
}

func NewContentPostscriptClient(
	contentClient *rpc.ContentClient,
) repo.ContentPostscriptClient {
	return &ContentPostscriptClient{
		contentClient: contentClient,
	}
}

func (r *ContentPostscriptClient) AddPostscript(
	ctx context.Context,
	req *repo.AddPostscriptReq,
) (*repo.ArticlePostscript, error) {
	reply, err := r.contentClient.Article.AddPostscript(ctx, &contentv1.AddPostscriptArticle_Req{
		ArticleId: req.ArticleID,
		Content:   req.Content,
		UserId:    req.UserID,
	})
	if err != nil {
		return nil, err
	}
	item := reply.GetArticlePostscript()
	return &repo.ArticlePostscript{
		ID:            item.GetId(),
		ArticleID:     item.GetArticleId(),
		Content:       item.GetContent(),
		ContentRender: util.LuteEngine.MarkdownStr(fmt.Sprintf("article_postscript_%d", item.GetId()), item.GetContent()),
		CreatedBy:     item.CreatedBy,
		UpdatedBy:     item.UpdatedBy,
		CreatedAt:     formatProtoTime(item.GetCreatedAt()),
		UpdatedAt:     formatProtoTime(item.GetUpdatedAt()),
	}, nil
}
