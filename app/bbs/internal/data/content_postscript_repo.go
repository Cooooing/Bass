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

func (r *ContentPostscriptClient) AddPostscript(ctx context.Context, req *repo.AddPostscriptReq) (*repo.ArticlePostscript, error) {
	reply, err := r.contentClient.Postscript.Add(ctx, &contentv1.AddPostscript_Req{
		ArticleId: req.ArticleID,
		Content:   req.Content,
		UserId:    req.UserID,
	})
	if err != nil {
		return nil, err
	}
	item := reply.GetPostscript()
	return &repo.ArticlePostscript{
		ID:            item.GetId(),
		ArticleID:     item.GetArticleId(),
		Content:       item.GetContent(),
		ContentRender: util.LuteEngine.MarkdownStr(fmt.Sprintf("article_postscript_%d", item.GetId()), item.GetContent()),
		Restriction:   int32(item.GetRestriction()),
		CreatedBy:     item.CreatedBy,
		UpdatedBy:     item.UpdatedBy,
		CreatedAt:     new(item.GetCreatedAt().AsTime()),
		UpdatedAt:     new(item.GetUpdatedAt().AsTime()),
	}, nil
}

func (r *ContentPostscriptClient) ListPostscripts(ctx context.Context, req *repo.ListPostscriptsReq) ([]*repo.ArticlePostscript, error) {
	reply, err := r.contentClient.Postscript.List(ctx, &contentv1.ListPostscripts_Req{
		ArticleId: req.ArticleID,
	})
	if err != nil {
		return nil, err
	}
	rows := make([]*repo.ArticlePostscript, 0, len(reply.GetRows()))
	for _, item := range reply.GetRows() {
		rows = append(rows, &repo.ArticlePostscript{
			ID:            item.GetId(),
			ArticleID:     item.GetArticleId(),
			Content:       item.GetContent(),
			ContentRender: util.LuteEngine.MarkdownStr(fmt.Sprintf("article_postscript_%d", item.GetId()), item.GetContent()),
			Restriction:   int32(item.GetRestriction()),
			CreatedBy:     item.CreatedBy,
			UpdatedBy:     item.UpdatedBy,
			CreatedAt:     new(item.GetCreatedAt().AsTime()),
			UpdatedAt:     new(item.GetUpdatedAt().AsTime()),
		})
	}
	return rows, nil
}
