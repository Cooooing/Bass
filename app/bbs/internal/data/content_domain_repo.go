package data

import (
	"bbs/internal/biz/repo"
	bbscontentv1 "common/api/gen/bbs/v1/content"
	contentv1 "common/api/gen/content/v1"
	"common/pkg/client/rpc"
	"context"
)

var _ repo.ContentDomainRepo = (*ContentDomainRepo)(nil)

type ContentDomainRepo struct {
	contentClient *rpc.ContentClient
}

func NewContentDomainRepo(contentClient *rpc.ContentClient) repo.ContentDomainRepo {
	return &ContentDomainRepo{contentClient: contentClient}
}

func (r *ContentDomainRepo) ListDomains(ctx context.Context, req *bbscontentv1.ListDomains_Request) (*bbscontentv1.ListDomains_Reply, error) {
	query := req.GetQuery()
	if query == nil {
		query = &bbscontentv1.DomainQuery{}
	}
	contentQuery := &contentv1.DomainQueryParams{
		Ids:         query.GetIds(),
		Name:        query.Name,
		Description: query.Description,
		Url:         query.Url,
		Icon:        query.Icon,
		IsNav:       query.IsNav,
	}
	if query.Status != nil {
		contentQuery.Status = new(contentv1.DomainStatus(*query.Status))
	}
	reply, err := r.contentClient.Domain.List(ctx, &contentv1.ListDomains_Request{
		Page:  req.Page,
		Query: contentQuery,
	})
	if err != nil {
		return nil, err
	}
	rows := make([]*bbscontentv1.Domain, 0, len(reply.GetRows()))
	for _, item := range reply.GetRows() {
		row := &bbscontentv1.Domain{
			Id:          item.GetId(),
			Name:        item.GetName(),
			Description: item.Description,
			Status:      bbscontentv1.DomainStatus(item.GetStatus()),
			Url:         item.Url,
			Icon:        item.Icon,
			IsNav:       item.GetIsNav(),
			CreatedBy:   item.CreatedBy,
			UpdatedBy:   item.UpdatedBy,
			CreatedAt:   formatProtoTime(item.GetCreatedAt()),
			UpdatedAt:   formatProtoTime(item.GetUpdatedAt()),
		}
		rows = append(rows, row)
	}
	return &bbscontentv1.ListDomains_Reply{Page: reply.GetPage(), Rows: rows}, nil
}
