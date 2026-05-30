package data

import (
	"bbs/internal/biz/repo"
	bbscontentv1 "common/api/gen/bbs/v1/content"
	contentv1 "common/api/gen/content/v1"
	"common/pkg/client/rpc"
	"context"
)

var _ repo.ContentTagRepo = (*ContentTagRepo)(nil)

type ContentTagRepo struct {
	contentClient *rpc.ContentClient
}

func NewContentTagRepo(contentClient *rpc.ContentClient) repo.ContentTagRepo {
	return &ContentTagRepo{contentClient: contentClient}
}

func (r *ContentTagRepo) ListTags(ctx context.Context, req *bbscontentv1.ListTags_Request) (*bbscontentv1.ListTags_Reply, error) {
	query := req.GetQuery()
	if query == nil {
		query = &bbscontentv1.TagQuery{}
	}
	contentQuery := &contentv1.TagQueryParams{
		Ids:         query.GetIds(),
		Name:        query.Name,
		Names:       query.GetNames(),
		Description: query.Description,
		DomainId:    query.DomainId,
	}
	if query.Status != nil {
		contentQuery.Status = new(contentv1.TagStatus(*query.Status))
	}
	reply, err := r.contentClient.Tag.List(ctx, &contentv1.ListTags_Request{
		Page:  req.Page,
		Query: contentQuery,
	})
	if err != nil {
		return nil, err
	}
	rows := make([]*bbscontentv1.Tag, 0, len(reply.GetRows()))
	for _, item := range reply.GetRows() {
		rows = append(rows, &bbscontentv1.Tag{
			Id:          item.GetId(),
			Name:        item.GetName(),
			Description: item.Description,
			DomainId:    item.DomainId,
			Status:      new(bbscontentv1.TagStatus(item.GetStatus())),
			CreatedBy:   item.CreatedBy,
			UpdatedBy:   item.UpdatedBy,
			CreatedAt:   formatProtoTime(item.GetCreatedAt()),
			UpdatedAt:   formatProtoTime(item.GetUpdatedAt()),
		})
	}
	return &bbscontentv1.ListTags_Reply{Page: reply.GetPage(), Rows: rows}, nil
}
