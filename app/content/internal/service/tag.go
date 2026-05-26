package service

import (
	cerrors "common/api/gen/common/errors"
	v1 "common/api/gen/content/v1"
	"common/pkg/util"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/biz/usecase"
	"content/internal/enum"
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TagService struct {
	v1.UnimplementedContentTagServiceServer
	tagUsecase *usecase.TagUsecase
}

func (s *TagService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterContentTagServiceServer(gs, s)
}

func NewTagService(tagUsecase *usecase.TagUsecase) *TagService {
	return &TagService{
		tagUsecase: tagUsecase,
	}
}

func (s *TagService) BatchCreate(ctx context.Context, req *v1.BatchCreateTags_Request) (*v1.BatchCreateTags_Reply, error) {
	tags := make([]*model.Tag, 0, len(req.Tags))
	for _, tagSave := range req.Tags {
		tagStatus, ok := enum.TagStatusMap.ToEnum(util.DerefOrDefault(tagSave.Status, v1.TagStatus_TAG_STATUS_NORMAL))
		if !ok {
			return nil, cerrors.ErrorBadRequest("invalid tag status")
		}
		tags = append(tags, &model.Tag{
			Name:        tagSave.Name,
			Description: tagSave.Description,
			DomainID:    tagSave.DomainId,
			Status:      tagStatus,
		})
	}
	saves, err := s.tagUsecase.Saves(ctx, tags)
	if err != nil {
		return nil, err
	}
	reply := make([]*v1.Tag, len(saves))
	for i, save := range saves {
		reply[i] = &v1.Tag{
			CreatedAt:   timestamppb.New(*save.CreatedAt),
			UpdatedAt:   timestamppb.New(*save.UpdatedAt),
			CreatedBy:   save.CreatedBy,
			UpdatedBy:   save.UpdatedBy,
			Id:          save.ID,
			Name:        save.Name,
			Description: save.Description,
			DomainId:    save.DomainID,
			Status:      new(enum.TagStatusMap.MustToProto(save.Status)),
		}
	}
	return &v1.BatchCreateTags_Reply{
		Tags: reply,
	}, err
}

func (s *TagService) Update(ctx context.Context, req *v1.UpdateTag_Request) (*v1.UpdateTag_Reply, error) {
	if req.Tag == nil {
		return nil, cerrors.ErrorBadRequest("tag is required")
	}
	if req.Tag.Id == nil {
		return nil, cerrors.ErrorBadRequest("tag id is required")
	}
	updateTagStatus, ok := enum.TagStatusMap.ToEnum(util.DerefOrDefault(req.Tag.Status, v1.TagStatus_TAG_STATUS_NORMAL))
	if !ok {
		return nil, cerrors.ErrorBadRequest("invalid tag status")
	}
	update, err := s.tagUsecase.Update(ctx, &model.Tag{
		ID:          *req.Tag.Id,
		Name:        req.Tag.Name,
		Description: req.Tag.Description,
		DomainID:    req.Tag.DomainId,
		Status:      updateTagStatus,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UpdateTag_Reply{
		Tag: &v1.Tag{
			CreatedAt:   timestamppb.New(*update.CreatedAt),
			UpdatedAt:   timestamppb.New(*update.UpdatedAt),
			CreatedBy:   update.CreatedBy,
			UpdatedBy:   update.UpdatedBy,
			Id:          update.ID,
			Name:        update.Name,
			Description: update.Description,
			DomainId:    update.DomainID,
			Status:      new(enum.TagStatusMap.MustToProto(update.Status)),
		},
	}, nil
}

func (s *TagService) List(ctx context.Context, req *v1.ListTags_Request) (*v1.ListTags_Reply, error) {
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
	if req.Query.Status != nil {
		if _, ok := enum.TagStatusMap.ToEnum(*req.Query.Status); !ok {
			return nil, cerrors.ErrorBadRequest("invalid tag status")
		}
		getReq.Status = new(*req.Query.Status)
	}
	data, page, err := s.tagUsecase.Page(ctx, req.Page, getReq)
	reply := make([]*v1.Tag, len(data))
	for i, datum := range data {
		reply[i] = &v1.Tag{
			CreatedAt:   timestamppb.New(*datum.CreatedAt),
			UpdatedAt:   timestamppb.New(*datum.UpdatedAt),
			CreatedBy:   datum.CreatedBy,
			UpdatedBy:   datum.UpdatedBy,
			Id:          datum.ID,
			Name:        datum.Name,
			Description: datum.Description,
			DomainId:    datum.DomainID,
			Status:      new(enum.TagStatusMap.MustToProto(datum.Status)),
		}
	}
	return &v1.ListTags_Reply{
		Page: page,
		Rows: reply,
	}, err
}
