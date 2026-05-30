package data

import (
	"bbs/internal/biz/repo"
	bbscontentv1 "common/api/gen/bbs/v1/content"
	contentv1 "common/api/gen/content/v1"
	"common/pkg/client/rpc"
	"context"
)

var _ repo.ContentPostscriptRepo = (*ContentPostscriptRepo)(nil)

type ContentPostscriptRepo struct {
	contentClient *rpc.ContentClient
}

func NewContentPostscriptRepo(contentClient *rpc.ContentClient) repo.ContentPostscriptRepo {
	return &ContentPostscriptRepo{contentClient: contentClient}
}

func (r *ContentPostscriptRepo) AddPostscript(ctx context.Context, req *bbscontentv1.AddPostscript_Request) (*bbscontentv1.AddPostscript_Reply, error) {
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
		ContentRender: item.GetContentRender(),
		CreatedBy:     item.CreatedBy,
		UpdatedBy:     item.UpdatedBy,
		CreatedAt:     formatProtoTime(item.GetCreatedAt()),
		UpdatedAt:     formatProtoTime(item.GetUpdatedAt()),
	}}, nil
}
