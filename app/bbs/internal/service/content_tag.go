package service

import (
	"bbs/internal/biz/usecase"
	"common/pkg/apperror"
	"common/pkg/constant"
	commonmodel "common/pkg/model"
	"common/pkg/util"
	bbscontentv1 "common/proto/gen/bbs/v1/content"
	bbscontentv1enum "common/proto/gen/bbs/v1/content/enum"
	"common/proto/gen/common"
	cerrors "common/proto/gen/common/errors"
	"context"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ContentTagService struct {
	bbscontentv1.UnimplementedTagServiceServer
	contentTagUsecase *usecase.ContentTagUsecase
}

func NewContentTagService(
	contentTagUsecase *usecase.ContentTagUsecase,
) *ContentTagService {
	return &ContentTagService{
		contentTagUsecase: contentTagUsecase,
	}
}

func (s *ContentTagService) RegisterGrpc(gs *grpc.Server) {
}

func (s *ContentTagService) RegisterHttp(hs *http.Server) {
	bbscontentv1.RegisterTagServiceHTTPServer(hs, s)
}

func (s *ContentTagService) Create(ctx context.Context, req *bbscontentv1.CreateTag_Req) (*bbscontentv1.CreateTag_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	var tagSave *usecase.ContentTagSave
	if tag := req.GetTag(); tag != nil {
		tagSave = &usecase.ContentTagSave{
			Name:        tag.GetName(),
			Description: tag.Description,
			DomainID:    tag.DomainId,
			Status:      tag.Status,
		}
	}
	resp, err := s.contentTagUsecase.CreateTag(ctx, &usecase.CreateTagReq{
		UserID: user.ID,
		Tag:    tagSave,
	})
	if err != nil {
		return nil, err
	}
	var tag *bbscontentv1.CreateTag_Resp_Tag
	if resp != nil {
		tag = &bbscontentv1.CreateTag_Resp_Tag{
			Id:          resp.ID,
			Name:        resp.Name,
			Description: resp.Description,
			DomainId:    resp.DomainID,
			CreatedBy:   resp.CreatedBy,
			UpdatedBy:   resp.UpdatedBy,
			CreatedAt:   timestamppb.New(*resp.CreatedAt),
			UpdatedAt:   timestamppb.New(*resp.UpdatedAt),
		}
		if resp.Status != nil {
			tag.Status = new(bbscontentv1enum.TagStatus(*resp.Status))
		}
	}
	return &bbscontentv1.CreateTag_Resp{
		Tag: tag,
	}, nil
}

func (s *ContentTagService) Update(ctx context.Context, req *bbscontentv1.UpdateTag_Req) (*bbscontentv1.UpdateTag_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	var tagSave *usecase.ContentTagSave
	if tag := req.GetTag(); tag != nil {
		tagSave = &usecase.ContentTagSave{
			Name:        tag.GetName(),
			Description: tag.Description,
			DomainID:    tag.DomainId,
			Status:      tag.Status,
		}
	}
	resp, err := s.contentTagUsecase.UpdateTag(ctx, &usecase.UpdateTagReq{
		UserID: user.ID,
		TagID:  req.GetTagId(),
		Tag:    tagSave,
	})
	if err != nil {
		return nil, err
	}
	var tag *bbscontentv1.UpdateTag_Resp_Tag
	if resp != nil {
		tag = &bbscontentv1.UpdateTag_Resp_Tag{
			Id:          resp.ID,
			Name:        resp.Name,
			Description: resp.Description,
			DomainId:    resp.DomainID,
			CreatedBy:   resp.CreatedBy,
			UpdatedBy:   resp.UpdatedBy,
			CreatedAt:   timestamppb.New(*resp.CreatedAt),
			UpdatedAt:   timestamppb.New(*resp.UpdatedAt),
		}
		if resp.Status != nil {
			tag.Status = new(bbscontentv1enum.TagStatus(*resp.Status))
		}
	}
	return &bbscontentv1.UpdateTag_Resp{
		Tag: tag,
	}, nil
}

func (s *ContentTagService) List(ctx context.Context, req *bbscontentv1.ListTags_Req) (*bbscontentv1.ListTags_Resp, error) {
	resp, err := s.contentTagUsecase.ListTags(ctx, &usecase.ListTagsReq{
		Page:  req.GetPage(),
		Query: req.GetQuery(),
	})
	if err != nil {
		return nil, err
	}
	var page *common.PageResp
	if resp.Page != nil {
		page = &common.PageResp{
			Page:  resp.Page.Page,
			Size:  resp.Page.Size,
			Total: resp.Page.Total,
		}
	}
	rows := make([]*bbscontentv1.ListTags_Resp_Tag, 0, len(resp.Rows))
	for _, row := range resp.Rows {
		if row == nil {
			rows = append(rows, nil)
			continue
		}
		tag := &bbscontentv1.ListTags_Resp_Tag{
			Id:          row.ID,
			Name:        row.Name,
			Description: row.Description,
			DomainId:    row.DomainID,
			CreatedBy:   row.CreatedBy,
			UpdatedBy:   row.UpdatedBy,
			CreatedAt:   timestamppb.New(*row.CreatedAt),
			UpdatedAt:   timestamppb.New(*row.UpdatedAt),
		}
		if row.Status != nil {
			tag.Status = new(bbscontentv1enum.TagStatus(*row.Status))
		}
		rows = append(rows, tag)
	}
	return &bbscontentv1.ListTags_Resp{
		Page: page,
		Rows: rows,
	}, nil
}
