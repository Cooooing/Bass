package data

import (
	"bff_bbs/internal/biz/repo"
	"common/pkg/client/rpc"
	contentv1 "common/proto/gen/content/v1"
	"context"
)

var _ repo.ContentBreezemoonClient = (*ContentBreezemoonClient)(nil)

type ContentBreezemoonClient struct {
	contentClient *rpc.ContentClient
}

func NewContentBreezemoonClient(contentClient *rpc.ContentClient) repo.ContentBreezemoonClient {
	return &ContentBreezemoonClient{contentClient: contentClient}
}

func (r *ContentBreezemoonClient) CreateBreezemoon(ctx context.Context, req *repo.CreateBreezemoonReq) (*repo.Breezemoon, error) {
	reply, err := r.contentClient.Breezemoon.Create(ctx, &contentv1.CreateBreezemoon_Req{
		Content:  req.Content,
		AuthorId: req.AuthorID,
		City:     req.City,
	})
	if err != nil {
		return nil, err
	}
	return r.breezemoonRepoModel(reply.GetBreezemoon()), nil
}

func (r *ContentBreezemoonClient) PageBreezemoons(ctx context.Context, req *repo.PageBreezemoonsReq) (*repo.PageBreezemoonsResp, error) {
	reply, err := r.contentClient.Breezemoon.Page(ctx, &contentv1.PageBreezemoons_Req{
		AuthorIds: req.AuthorIDs,
		Cursor:    req.Cursor,
		Size:      req.Size,
	})
	if err != nil {
		return nil, err
	}
	rows := make([]*repo.Breezemoon, 0, len(reply.GetRows()))
	for _, row := range reply.GetRows() {
		rows = append(rows, r.breezemoonRepoModel(row))
	}
	return &repo.PageBreezemoonsResp{Rows: rows, NextCursor: reply.NextCursor}, nil
}

func (*ContentBreezemoonClient) breezemoonRepoModel(row *contentv1.Breezemoon) *repo.Breezemoon {
	if row == nil {
		return nil
	}
	result := &repo.Breezemoon{
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
