package service

import (
	v1 "common/api/content/v1"
	"content/internal/biz"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TagService struct {
	v1.UnimplementedContentTagServiceServer
	*BaseService
	domainTag *biz.TagDomain
}

func (s *TagService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterContentTagServiceServer(gs, s)
}

func (s *TagService) RegisterHttp(hs *http.Server) {
	v1.RegisterContentTagServiceHTTPServer(hs, s)
}

func NewTagService(baseService *BaseService, domainTag *biz.TagDomain) *TagService {
	return &TagService{
		BaseService: baseService,
		domainTag:   domainTag,
	}
}

func (s *TagService) Adds(ctx context.Context, req *v1.AddTagsRequest) (*v1.AddTagsReply, error) {
	tags := make([]*model.Tag, 0, len(req.Tags))
	for _, tag := range req.Tags {
		tags = append(tags, &model.Tag{
			Name:        tag.Name,
			Description: tag.Description,
			DomainID:    tag.DomainId,
			Status:      tag.Status,
		})
	}
	saves, err := s.domainTag.Saves(ctx, tags)
	if err != nil {
		return nil, err
	}
	reply := make([]*v1.TagReply, 0, len(saves))
	for _, save := range saves {
		reply = append(reply, &v1.TagReply{
			CreatedAt:    timestamppb.New(*save.CreatedAt),
			UpdatedAt:    timestamppb.New(*save.UpdatedAt),
			CreatedBy:    save.CreatedBy,
			UpdatedBy:    save.UpdatedBy,
			Id:           save.ID,
			Name:         save.Name,
			Description:  save.Description,
			DomainId:     save.DomainID,
			Status:       save.Status,
			ArticleCount: save.ArticleCount,
		})
	}
	return &v1.AddTagsReply{
		Tags: reply,
	}, err
}

func (s *TagService) Update(ctx context.Context, req *v1.UpdateTagRequest) (*v1.UpdateTagReply, error) {
	update, err := s.domainTag.Update(ctx, &model.Tag{
		ID:          req.Tag.Id,
		Name:        req.Tag.Name,
		Description: req.Tag.Description,
		DomainID:    req.Tag.DomainId,
		Status:      req.Tag.Status,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UpdateTagReply{
		Tag: &v1.TagReply{
			CreatedAt: timestamppb.New(*update.CreatedAt),
			UpdatedAt: timestamppb.New(*update.UpdatedAt),
			CreatedBy: update.CreatedBy,
			UpdatedBy: update.UpdatedBy,
			Id:        update.ID,
			Name:      update.Name,
		},
	}, nil
}

func (s *TagService) Page(ctx context.Context, req *v1.PageTagRequest) (*v1.PageTagReply, error) {
	data, page, err := s.domainTag.Page(ctx, req.Page, &repo.TagGetReq{})
	reply := make([]*v1.TagReply, len(data))
	for i, datum := range data {
		reply[i] = &v1.TagReply{
			CreatedAt:    timestamppb.New(*datum.CreatedAt),
			UpdatedAt:    timestamppb.New(*datum.UpdatedAt),
			CreatedBy:    datum.CreatedBy,
			UpdatedBy:    datum.UpdatedBy,
			Id:           datum.ID,
			Name:         datum.Name,
			Description:  datum.Description,
			DomainId:     datum.DomainID,
			Status:       datum.Status,
			ArticleCount: datum.ArticleCount,
		}
	}
	return &v1.PageTagReply{
		Page: page,
		Tags: reply,
	}, err
}
