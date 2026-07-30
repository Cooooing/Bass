package service

import (
	"common/pkg/apperror"
	"common/pkg/util"
	"common/proto/gen/common"
	cerrors "common/proto/gen/common/errors"
	v1 "common/proto/gen/content/v1"
	contentv1enum "common/proto/gen/content/v1/enum"
	"content/internal/biz/base"
	"content/internal/biz/model"
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

func (s *TagService) RegisterHttp(hs *http.Server) {
}

func NewTagService(
	tagUsecase *usecase.TagUsecase,
) *TagService {
	return &TagService{
		tagUsecase: tagUsecase,
	}
}

func (s *TagService) BatchCreate(ctx context.Context, req *v1.BatchCreateTags_Req) (*v1.BatchCreateTags_Resp, error) {
	if req.UserId <= 0 || len(req.Tags) == 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	tags := make([]*model.Tag, 0, len(req.Tags))
	for _, item := range req.Tags {
		tagStatus, ok := enum.TagStatusMap.ToEnum(util.DerefOrDefault(item.Status, contentv1enum.TagStatus_TAG_STATUS_ENABLED))
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_TAG_INVALID)
		}
		tags = append(tags, &model.Tag{
			Code:        item.GetCode(),
			Name:        item.GetName(),
			Description: item.Description,
			DomainID:    item.DomainId,
			Icon:        item.Icon,
			Sort:        item.GetSort(),
			Status:      tagStatus,
			CreatedBy:   &req.UserId,
			UpdatedBy:   &req.UserId,
		})
	}
	saves, err := s.tagUsecase.Saves(ctx, tags)
	if err != nil {
		return nil, err
	}
	rows := make([]*v1.BatchCreateTags_Resp_Tag, 0, len(saves))
	for _, save := range saves {
		row := &v1.BatchCreateTags_Resp_Tag{
			Id:           save.ID,
			Code:         save.Code,
			Name:         save.Name,
			Description:  save.Description,
			DomainId:     save.DomainID,
			Icon:         save.Icon,
			Sort:         save.Sort,
			ArticleCount: save.ArticleCount,
			Status:       new(enum.TagStatusMap.MustToProto(save.Status)),
			CreatedBy:    save.CreatedBy,
			UpdatedBy:    save.UpdatedBy,
		}
		if save.CreatedAt != nil {
			row.CreatedAt = timestamppb.New(*save.CreatedAt)
		}
		if save.UpdatedAt != nil {
			row.UpdatedAt = timestamppb.New(*save.UpdatedAt)
		}
		rows = append(rows, row)
	}
	return &v1.BatchCreateTags_Resp{Rows: rows}, nil
}

func (s *TagService) Update(ctx context.Context, req *v1.UpdateTag_Req) (*v1.UpdateTag_Resp, error) {
	if req.Tag == nil || req.TagId <= 0 || req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	tagStatus, ok := enum.TagStatusMap.ToEnum(util.DerefOrDefault(req.Tag.Status, contentv1enum.TagStatus_TAG_STATUS_ENABLED))
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_TAG_INVALID)
	}
	data, err := s.tagUsecase.Update(ctx, &model.Tag{
		ID:          req.TagId,
		Code:        req.Tag.GetCode(),
		Name:        req.Tag.GetName(),
		Description: req.Tag.Description,
		DomainID:    req.Tag.DomainId,
		Icon:        req.Tag.Icon,
		Sort:        req.Tag.GetSort(),
		Status:      tagStatus,
		UpdatedBy:   &req.UserId,
	})
	if err != nil {
		return nil, err
	}
	reply := &v1.UpdateTag_Resp_Tag{
		Id:           data.ID,
		Code:         data.Code,
		Name:         data.Name,
		Description:  data.Description,
		DomainId:     data.DomainID,
		Icon:         data.Icon,
		Sort:         data.Sort,
		ArticleCount: data.ArticleCount,
		Status:       new(enum.TagStatusMap.MustToProto(data.Status)),
		CreatedBy:    data.CreatedBy,
		UpdatedBy:    data.UpdatedBy,
	}
	if data.CreatedAt != nil {
		reply.CreatedAt = timestamppb.New(*data.CreatedAt)
	}
	if data.UpdatedAt != nil {
		reply.UpdatedAt = timestamppb.New(*data.UpdatedAt)
	}
	return &v1.UpdateTag_Resp{Tag: reply}, nil
}

func (s *TagService) BindArticle(ctx context.Context, req *v1.BindArticleTags_Req) (*v1.BindArticleTags_Resp, error) {
	if req.ArticleId <= 0 || req.UserId <= 0 || len(req.TagIds) == 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if err := s.tagUsecase.BindArticle(ctx, &usecase.ArticleTagReq{
		ArticleID: req.GetArticleId(),
		TagIDs:    req.GetTagIds(),
		UserID:    req.GetUserId(),
		Manager:   req.GetManager(),
	}); err != nil {
		return nil, err
	}
	return &v1.BindArticleTags_Resp{}, nil
}

func (s *TagService) UnbindArticle(ctx context.Context, req *v1.UnbindArticleTags_Req) (*v1.UnbindArticleTags_Resp, error) {
	if req.ArticleId <= 0 || req.UserId <= 0 || len(req.TagIds) == 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if err := s.tagUsecase.UnbindArticle(ctx, &usecase.ArticleTagReq{
		ArticleID: req.GetArticleId(),
		TagIDs:    req.GetTagIds(),
		UserID:    req.GetUserId(),
		Manager:   req.GetManager(),
	}); err != nil {
		return nil, err
	}
	return &v1.UnbindArticleTags_Resp{}, nil
}

func (s *TagService) ListArticleTags(ctx context.Context, req *v1.ListArticleTags_Req) (*v1.ListArticleTags_Resp, error) {
	if req.ArticleId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	rows, err := s.tagUsecase.ListArticleTags(ctx, req.GetArticleId())
	if err != nil {
		return nil, err
	}
	reply := make([]*v1.ListArticleTags_Resp_Tag, 0, len(rows))
	for _, item := range rows {
		row := &v1.ListArticleTags_Resp_Tag{
			Id:           item.ID,
			Code:         item.Code,
			Name:         item.Name,
			Description:  item.Description,
			DomainId:     item.DomainID,
			Icon:         item.Icon,
			Sort:         item.Sort,
			ArticleCount: item.ArticleCount,
			Status:       new(enum.TagStatusMap.MustToProto(item.Status)),
			CreatedBy:    item.CreatedBy,
			UpdatedBy:    item.UpdatedBy,
		}
		if item.CreatedAt != nil {
			row.CreatedAt = timestamppb.New(*item.CreatedAt)
		}
		if item.UpdatedAt != nil {
			row.UpdatedAt = timestamppb.New(*item.UpdatedAt)
		}
		reply = append(reply, row)
	}
	return &v1.ListArticleTags_Resp{Rows: reply}, nil
}

func (s *TagService) List(ctx context.Context, req *v1.ListTags_Req) (*v1.ListTags_Resp, error) {
	query := req.GetQuery()
	if query == nil {
		query = &v1.ListTags_Req_Query{}
	}
	var tagStatus *enum.TagStatus
	if query.Status != nil {
		status, ok := enum.TagStatusMap.ToEnum(*query.Status)
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_TAG_INVALID)
		}
		tagStatus = &status
	}
	pageResp, err := s.tagUsecase.Page(ctx, &usecase.TagPageReq{
		Page: &base.PageRequest{
			Page: 1,
			Size: 1000,
		},
		TagIDs:      query.GetIds(),
		Code:        query.Code,
		Name:        query.Name,
		Names:       query.GetNames(),
		Description: query.Description,
		Status:      tagStatus,
		DomainID:    query.DomainId,
	})
	if err != nil {
		return nil, err
	}
	reply := make([]*v1.ListTags_Resp_Tag, 0, len(pageResp.Rows))
	for _, item := range pageResp.Rows {
		row := &v1.ListTags_Resp_Tag{
			Id:           item.ID,
			Code:         item.Code,
			Name:         item.Name,
			Description:  item.Description,
			DomainId:     item.DomainID,
			Icon:         item.Icon,
			Sort:         item.Sort,
			ArticleCount: item.ArticleCount,
			Status:       new(enum.TagStatusMap.MustToProto(item.Status)),
			CreatedBy:    item.CreatedBy,
			UpdatedBy:    item.UpdatedBy,
		}
		if item.CreatedAt != nil {
			row.CreatedAt = timestamppb.New(*item.CreatedAt)
		}
		if item.UpdatedAt != nil {
			row.UpdatedAt = timestamppb.New(*item.UpdatedAt)
		}
		reply = append(reply, row)
	}
	return &v1.ListTags_Resp{Rows: reply}, nil
}

func (s *TagService) Page(ctx context.Context, req *v1.PageTags_Req) (*v1.PageTags_Resp, error) {
	query := req.GetQuery()
	if query == nil {
		query = &v1.PageTags_Req_Query{}
	}
	var tagStatus *enum.TagStatus
	if query.Status != nil {
		status, ok := enum.TagStatusMap.ToEnum(*query.Status)
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_TAG_INVALID)
		}
		tagStatus = &status
	}
	pageResp, err := s.tagUsecase.Page(ctx, &usecase.TagPageReq{
		Page: &base.PageRequest{
			Page: int64(req.GetPage().GetPage()),
			Size: int64(req.GetPage().GetSize()),
		},
		TagIDs:      query.GetIds(),
		Code:        query.Code,
		Name:        query.Name,
		Names:       query.GetNames(),
		Description: query.Description,
		Status:      tagStatus,
		DomainID:    query.DomainId,
	})
	if err != nil {
		return nil, err
	}
	reply := make([]*v1.PageTags_Resp_Tag, 0, len(pageResp.Rows))
	for _, item := range pageResp.Rows {
		row := &v1.PageTags_Resp_Tag{
			Id:           item.ID,
			Code:         item.Code,
			Name:         item.Name,
			Description:  item.Description,
			DomainId:     item.DomainID,
			Icon:         item.Icon,
			Sort:         item.Sort,
			ArticleCount: item.ArticleCount,
			Status:       new(enum.TagStatusMap.MustToProto(item.Status)),
			CreatedBy:    item.CreatedBy,
			UpdatedBy:    item.UpdatedBy,
		}
		if item.CreatedAt != nil {
			row.CreatedAt = timestamppb.New(*item.CreatedAt)
		}
		if item.UpdatedAt != nil {
			row.UpdatedAt = timestamppb.New(*item.UpdatedAt)
		}
		reply = append(reply, row)
	}
	return &v1.PageTags_Resp{
		Page: &common.PageResp{
			Page:  uint32(pageResp.Page.Page),
			Size:  uint32(pageResp.Page.Size),
			Total: uint32(pageResp.Page.Total),
		},
		Rows: reply,
	}, nil
}
