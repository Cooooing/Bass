package data

import (
	"bbs/internal/biz/repo"
	"common/pkg/client/rpc"
	bbscontentv1 "common/proto/gen/bbs/v1/content"
	contentv1 "common/proto/gen/content/v1"
	"context"
)

var _ repo.ContentTagClient = (*ContentTagClient)(nil)

type ContentTagClient struct {
	contentClient *rpc.ContentClient
}

func NewContentTagClient(contentClient *rpc.ContentClient) repo.ContentTagClient {
	return &ContentTagClient{contentClient: contentClient}
}

func (r *ContentTagClient) CreateTag(ctx context.Context, req *bbscontentv1.CreateTag_Request) (*bbscontentv1.CreateTag_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	tag := req.GetTag()
	var status *contentv1.TagStatus
	if tag.Status != nil {
		status = new(contentv1.TagStatus(tag.GetStatus()))
	}
	reply, err := r.contentClient.Tag.BatchCreate(ctx, &contentv1.BatchCreateTags_Request{
		UserId: userID,
		Tags: []*contentv1.BatchCreateTags_Request_Tag{
			{
				Name:        tag.GetName(),
				Description: tag.Description,
				DomainId:    tag.DomainId,
				Status:      status,
			},
		},
	})
	if err != nil {
		return nil, err
	}
	rows := reply.GetRows()
	if len(rows) == 0 {
		return &bbscontentv1.CreateTag_Reply{}, nil
	}
	item := rows[0]
	return &bbscontentv1.CreateTag_Reply{Tag: &bbscontentv1.Tag{
		Id:          item.GetId(),
		Name:        item.GetName(),
		Description: item.Description,
		DomainId:    item.DomainId,
		Status:      new(bbscontentv1.TagStatus(item.GetStatus())),
		CreatedBy:   item.CreatedBy,
		UpdatedBy:   item.UpdatedBy,
		CreatedAt:   formatProtoTime(item.GetCreatedAt()),
		UpdatedAt:   formatProtoTime(item.GetUpdatedAt()),
	}}, nil
}

func (r *ContentTagClient) UpdateTag(ctx context.Context, req *bbscontentv1.UpdateTag_Request) (*bbscontentv1.UpdateTag_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	tag := req.GetTag()
	var status *contentv1.TagStatus
	if tag.Status != nil {
		status = new(contentv1.TagStatus(tag.GetStatus()))
	}
	reply, err := r.contentClient.Tag.Update(ctx, &contentv1.UpdateTag_Request{
		TagId:  req.GetTagId(),
		UserId: userID,
		Tag: &contentv1.UpdateTag_Request_Tag{
			Name:        tag.GetName(),
			Description: tag.Description,
			DomainId:    tag.DomainId,
			Status:      status,
		},
	})
	if err != nil {
		return nil, err
	}
	item := reply.GetTag()
	return &bbscontentv1.UpdateTag_Reply{Tag: &bbscontentv1.Tag{
		Id:          item.GetId(),
		Name:        item.GetName(),
		Description: item.Description,
		DomainId:    item.DomainId,
		Status:      new(bbscontentv1.TagStatus(item.GetStatus())),
		CreatedBy:   item.CreatedBy,
		UpdatedBy:   item.UpdatedBy,
		CreatedAt:   formatProtoTime(item.GetCreatedAt()),
		UpdatedAt:   formatProtoTime(item.GetUpdatedAt()),
	}}, nil
}

func (r *ContentTagClient) ListTags(ctx context.Context, req *bbscontentv1.ListTags_Request) (*bbscontentv1.ListTags_Reply, error) {
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
	reply, err := r.contentClient.Tag.Page(ctx, &contentv1.PageTags_Request{
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
