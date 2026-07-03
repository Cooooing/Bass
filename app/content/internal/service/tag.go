package service

import (
	"common/pkg/apperror"
	"common/pkg/server"
	"common/pkg/util"
	cerrors "common/proto/gen/common/errors"
	v1 "common/proto/gen/content/v1"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/biz/usecase"
	"content/internal/enum"
	"context"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TagService struct {
	v1.UnimplementedContentTagServiceServer
	tagUsecase *usecase.TagUsecase
}

func (s *TagService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterContentTagServiceServer(gs, s)
}

func (s *TagService) RegisterHttp(hs *http.Server) {}

func NewTagService(tagUsecase *usecase.TagUsecase) *TagService {
	return &TagService{
		tagUsecase: tagUsecase,
	}
}

func (s *TagService) BatchCreate(ctx context.Context, req *v1.BatchCreateTags_Request) (*v1.BatchCreateTags_Reply, error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	tags := make([]*model.Tag, 0, len(req.Tags))
	for _, tagSave := range req.Tags {
		tagStatus, ok := enum.TagStatusMap.ToEnum(util.DerefOrDefault(tagSave.Status, v1.TagStatus_TAG_STATUS_ENABLED))
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_TAG_INVALID)
		}
		tags = append(tags, &model.Tag{
			Name:        tagSave.Name,
			Description: tagSave.Description,
			DomainID:    tagSave.DomainId,
			Status:      tagStatus,
			CreatedBy:   new(req.UserId),
			UpdatedBy:   new(req.UserId),
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
		Rows: reply,
	}, err
}

func (s *TagService) Update(ctx context.Context, req *v1.UpdateTag_Request) (*v1.UpdateTag_Reply, error) {
	if req.Tag == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_TAG_INVALID)
	}
	if req.TagId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_TAG_INVALID)
	}
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	updateTagStatus, ok := enum.TagStatusMap.ToEnum(util.DerefOrDefault(req.Tag.Status, v1.TagStatus_TAG_STATUS_ENABLED))
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_TAG_INVALID)
	}
	update, err := s.tagUsecase.Update(ctx, &model.Tag{
		ID:          req.TagId,
		Name:        req.Tag.Name,
		Description: req.Tag.Description,
		DomainID:    req.Tag.DomainId,
		Status:      updateTagStatus,
		UpdatedBy:   new(req.UserId),
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
	var tagStatus *enum.TagStatus
	if req.Query.Status != nil {
		status, ok := enum.TagStatusMap.ToEnum(*req.Query.Status)
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_TAG_INVALID)
		}
		tagStatus = new(status)
	}
	getReq := &repo.TagGetReq{
		TagIds:      req.Query.Ids,
		Name:        req.Query.Name,
		Names:       req.Query.Names,
		Description: req.Query.Description,
		Status:      tagStatus,
		DomainId:    req.Query.DomainId,
	}
	data, _, err := s.tagUsecase.Page(ctx, server.GetPageMax(), getReq)
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
		Rows: reply,
	}, err
}

func (s *TagService) Page(ctx context.Context, req *v1.PageTags_Request) (*v1.PageTags_Reply, error) {
	req.Query = util.OrDefault(req.Query, &v1.TagQueryParams{})
	var tagStatus *enum.TagStatus
	if req.Query.Status != nil {
		status, ok := enum.TagStatusMap.ToEnum(*req.Query.Status)
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_TAG_INVALID)
		}
		tagStatus = new(status)
	}
	getReq := &repo.TagGetReq{
		TagIds:      req.Query.Ids,
		Name:        req.Query.Name,
		Names:       req.Query.Names,
		Description: req.Query.Description,
		Status:      tagStatus,
		DomainId:    req.Query.DomainId,
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
	return &v1.PageTags_Reply{
		Page: page,
		Rows: reply,
	}, err
}
