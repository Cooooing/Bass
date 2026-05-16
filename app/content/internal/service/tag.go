package service

import (
	v1 "common/api/gen/content/v1"
	commonModel "common/pkg/model"
	"common/pkg/util"
	"content/internal/biz/domain"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/ent/gen"
	"context"
	"errors"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type TagService struct {
	v1.UnimplementedContentTagServiceServer
	domainTag *domain.TagDomain
}

func (s *TagService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterContentTagServiceServer(gs, s)
}

func (s *TagService) RegisterHttp(hs *http.Server) {
	v1.RegisterContentTagServiceHTTPServer(hs, s)
}

func NewTagService(domainTag *domain.TagDomain) *TagService {
	return &TagService{
		domainTag: domainTag,
	}
}

func (s *TagService) Adds(ctx context.Context, req *v1.AddsTag_Request) (*v1.AddsTag_Reply, error) {
	tags := make([]*model.Tag, 0, len(req.Tags))
	for i, tagSave := range req.Tags {
		tags[i] = &model.Tag{Tag: &gen.Tag{
			Name:        tagSave.Name,
			Description: tagSave.Description,
			DomainID:    tagSave.DomainId,
			Status:      int32(util.DerefOrDefault(tagSave.Status, v1.TagStatus_TAG_STATUS_NORMAL)),
		}}
	}
	saves, err := s.domainTag.Saves(ctx, tags)
	if err != nil {
		return nil, err
	}
	reply := make([]*v1.Tag, 0, len(saves))
	for i, save := range saves {
		reply[i] = save.ConvertToRpc()
	}
	return &v1.AddsTag_Reply{
		Tags: reply,
	}, err
}

func (s *TagService) Update(ctx context.Context, req *v1.UpdateTag_Request) (*v1.UpdateTag_Reply, error) {
	if req.Tag.Id == nil {
		return nil, errors.New("tag id is nil")
	}
	update, err := s.domainTag.Update(ctx, &model.Tag{Tag: &gen.Tag{
		ID:          *req.Tag.Id,
		Name:        req.Tag.Name,
		Description: req.Tag.Description,
		DomainID:    req.Tag.DomainId,
		Status:      int32(util.DerefOrDefault(req.Tag.Status, v1.TagStatus_TAG_STATUS_NORMAL)),
	}})
	if err != nil {
		return nil, err
	}
	return &v1.UpdateTag_Reply{
		Tag: update.ConvertToRpc(),
	}, nil
}

func (s *TagService) Page(ctx context.Context, req *v1.PageTag_Request) (*v1.PageTag_Reply, error) {
	req.Query = util.OrDefault(req.Query, &v1.TagQueryParams{})
	getReq := &repo.TagGetReq{
		TagIds:      req.Query.Ids,
		UserId:      req.Query.UserId,
		Name:        req.Query.Name,
		Names:       req.Query.Names,
		Description: req.Query.Description,
		Status:      new(v1.TagStatus_TAG_STATUS_NORMAL),
		DomainId:    req.Query.DomainId,
	}
	if req.Query.ArticleCount != nil {
		getReq.ArticleCount = &commonModel.Range[int32]{Start: req.Query.ArticleCount.Start, End: req.Query.ArticleCount.End}
	}
	data, page, err := s.domainTag.Page(ctx, req.Page, getReq)
	reply := make([]*v1.Tag, len(data))
	for i, datum := range data {
		reply[i] = datum.ConvertToRpc()
	}
	return &v1.PageTag_Reply{
		Page: page,
		Rows: reply,
	}, err
}
