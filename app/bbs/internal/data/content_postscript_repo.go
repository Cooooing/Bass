package data

import (
	"bbs/internal/biz/repo"
	"common/pkg/client/rpc"
	"common/pkg/util"
	bbscontentv1 "common/proto/gen/bbs/v1/content"
	contentv1 "common/proto/gen/content/v1"
	"context"
	"fmt"
)

var _ repo.ContentPostscriptClient = (*ContentPostscriptClient)(nil)

type ContentPostscriptClient struct {
	contentClient *rpc.ContentClient
}

func NewContentPostscriptClient(contentClient *rpc.ContentClient) repo.ContentPostscriptClient {
	return &ContentPostscriptClient{contentClient: contentClient}
}

func (r *ContentPostscriptClient) AddPostscript(ctx context.Context, req *bbscontentv1.AddPostscript_Request) (*bbscontentv1.AddPostscript_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	reply, err := r.contentClient.Article.AddPostscript(ctx, &contentv1.AddPostscriptArticle_Request{ArticleId: req.GetArticleId(), Content: req.GetContent(), UserId: userID})
	if err != nil {
		return nil, err
	}
	item := reply.GetArticlePostscript()
	return &bbscontentv1.AddPostscript_Reply{Postscript: &bbscontentv1.ArticlePostscript{
		Id:            item.GetId(),
		ArticleId:     item.GetArticleId(),
		Content:       item.GetContent(),
		ContentRender: util.LuteEngine.MarkdownStr(fmt.Sprintf("article_postscript_%d", item.GetId()), item.GetContent()),
		CreatedBy:     item.CreatedBy,
		UpdatedBy:     item.UpdatedBy,
		CreatedAt:     formatProtoTime(item.GetCreatedAt()),
		UpdatedAt:     formatProtoTime(item.GetUpdatedAt()),
	}}, nil
}
