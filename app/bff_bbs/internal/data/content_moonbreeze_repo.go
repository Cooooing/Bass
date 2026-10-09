package data

import (
	"bff_bbs/internal/biz/model"
	"common/pkg/client/rpc"
	contentv1 "common/proto/gen/content/v1"
	"context"

	"bff_bbs/internal/biz/repo"
)

var _ repo.ContentMoonbreezeClient = (*ContentMoonbreezeClient)(nil)

type ContentMoonbreezeClient struct {
	contentClient *rpc.ContentClient
}

func NewContentMoonbreezeClient(contentClient *rpc.ContentClient) repo.ContentMoonbreezeClient {
	return &ContentMoonbreezeClient{contentClient: contentClient}
}

func (r *ContentMoonbreezeClient) CreateMoonbreeze(ctx context.Context, req *repo.CreateMoonbreezeReq) (*model.Moonbreeze, error) {
	reply, err := r.contentClient.Moonbreeze.Create(ctx, &contentv1.CreateMoonbreeze_Req{
		Content:  req.Content,
		AuthorId: req.AuthorID,
		City:     req.City,
	})
	if err != nil {
		return nil, err
	}
	return r.moonbreezeRepoModel(reply.GetMoonbreeze()), nil
}

func (r *ContentMoonbreezeClient) PageMoonbreezes(ctx context.Context, req *repo.PageMoonbreezesReq) (*model.PageMoonbreezesResp, error) {
	reply, err := r.contentClient.Moonbreeze.Page(ctx, &contentv1.PageMoonbreezes_Req{
		AuthorIds:         req.AuthorIDs,
		ExcludedAuthorIds: req.ExcludedAuthorIDs,
		Cursor:            req.Cursor,
		Size:              req.Size,
	})
	if err != nil {
		return nil, err
	}
	rows := make([]*model.Moonbreeze, 0, len(reply.GetRows()))
	for _, row := range reply.GetRows() {
		rows = append(rows, r.moonbreezeRepoModel(row))
	}
	return &model.PageMoonbreezesResp{Rows: rows, NextCursor: reply.NextCursor}, nil
}

func (*ContentMoonbreezeClient) moonbreezeRepoModel(row *contentv1.Moonbreeze) *model.Moonbreeze {
	if row == nil {
		return nil
	}
	result := &model.Moonbreeze{
		ID:       row.GetId(),
		Content:  row.GetContent(),
		AuthorID: row.GetAuthorId(),
		City:     row.City,
	}
	if row.GetCreatedAt() != nil {
		value := row.GetCreatedAt().AsTime()
		result.CreatedAt = &value
	}
	return result
}
