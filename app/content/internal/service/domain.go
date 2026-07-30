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

type DomainService struct {
	v1.UnimplementedContentDomainServiceServer
	contentUsecase *usecase.ContentUsecase
}

func (s *DomainService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterContentDomainServiceServer(gs, s)
}

func (s *DomainService) RegisterHttp(hs *http.Server) {
}

func NewDomainService(
	contentUsecase *usecase.ContentUsecase,
) *DomainService {
	return &DomainService{
		contentUsecase: contentUsecase,
	}
}

func (s *DomainService) BatchCreate(ctx context.Context, req *v1.BatchCreateDomains_Req) (*v1.BatchCreateDomains_Resp, error) {
	if req.UserId <= 0 || len(req.Domains) == 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	domains := make([]*model.Domain, 0, len(req.Domains))
	for _, item := range req.Domains {
		domainStatus, ok := enum.DomainStatusMap.ToEnum(util.DerefOrDefault(item.Status, contentv1enum.DomainStatus_DOMAIN_STATUS_ENABLED))
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_DOMAIN_INVALID)
		}
		domains = append(domains, &model.Domain{
			Code:        item.GetCode(),
			Name:        item.GetName(),
			Description: item.Description,
			Status:      domainStatus,
			URL:         item.Url,
			Icon:        item.Icon,
			IsNav:       item.GetIsNav(),
			Sort:        item.GetSort(),
			CreatedBy:   &req.UserId,
			UpdatedBy:   &req.UserId,
		})
	}
	saves, err := s.contentUsecase.Adds(ctx, domains)
	if err != nil {
		return nil, err
	}
	rows := make([]*v1.BatchCreateDomains_Resp_Domain, 0, len(saves))
	for _, save := range saves {
		row := &v1.BatchCreateDomains_Resp_Domain{
			Id:          save.ID,
			Code:        save.Code,
			Name:        save.Name,
			Description: save.Description,
			Status:      enum.DomainStatusMap.MustToProto(save.Status),
			Url:         save.URL,
			Icon:        save.Icon,
			IsNav:       save.IsNav,
			Sort:        save.Sort,
			CreatedBy:   save.CreatedBy,
			UpdatedBy:   save.UpdatedBy,
		}
		if save.CreatedAt != nil {
			row.CreatedAt = timestamppb.New(*save.CreatedAt)
		}
		if save.UpdatedAt != nil {
			row.UpdatedAt = timestamppb.New(*save.UpdatedAt)
		}
		rows = append(rows, row)
	}
	return &v1.BatchCreateDomains_Resp{Rows: rows}, nil
}

func (s *DomainService) Update(ctx context.Context, req *v1.UpdateDomain_Req) (*v1.UpdateDomain_Resp, error) {
	if req.Domain == nil || req.UserId <= 0 || req.Domain.Id <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	domainStatus, ok := enum.DomainStatusMap.ToEnum(util.DerefOrDefault(req.Domain.Status, contentv1enum.DomainStatus_DOMAIN_STATUS_ENABLED))
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_DOMAIN_INVALID)
	}
	data, err := s.contentUsecase.Update(ctx, &model.Domain{
		ID:          req.Domain.GetId(),
		Code:        req.Domain.GetCode(),
		Name:        req.Domain.GetName(),
		Description: req.Domain.Description,
		Status:      domainStatus,
		URL:         req.Domain.Url,
		Icon:        req.Domain.Icon,
		IsNav:       req.Domain.GetIsNav(),
		Sort:        req.Domain.GetSort(),
		UpdatedBy:   &req.UserId,
	})
	if err != nil {
		return nil, err
	}
	reply := &v1.UpdateDomain_Resp_Domain{
		Id:          data.ID,
		Code:        data.Code,
		Name:        data.Name,
		Description: data.Description,
		Status:      enum.DomainStatusMap.MustToProto(data.Status),
		Url:         data.URL,
		Icon:        data.Icon,
		IsNav:       data.IsNav,
		Sort:        data.Sort,
		CreatedBy:   data.CreatedBy,
		UpdatedBy:   data.UpdatedBy,
	}
	if data.CreatedAt != nil {
		reply.CreatedAt = timestamppb.New(*data.CreatedAt)
	}
	if data.UpdatedAt != nil {
		reply.UpdatedAt = timestamppb.New(*data.UpdatedAt)
	}
	return &v1.UpdateDomain_Resp{Domain: reply}, nil
}

func (s *DomainService) List(ctx context.Context, req *v1.ListDomains_Req) (*v1.ListDomains_Resp, error) {
	query := req.GetQuery()
	if query == nil {
		query = &v1.ListDomains_Req_Query{}
	}
	var domainStatus *enum.DomainStatus
	if query.Status != nil {
		status, ok := enum.DomainStatusMap.ToEnum(*query.Status)
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_DOMAIN_INVALID)
		}
		domainStatus = &status
	}
	pageResp, err := s.contentUsecase.Page(ctx, &usecase.DomainPageReq{
		Page: &base.PageRequest{
			Page: 1,
			Size: 1000,
		},
		DomainIDs:   query.GetIds(),
		Code:        query.Code,
		Name:        query.Name,
		Description: query.Description,
		Status:      domainStatus,
		URL:         query.Url,
		Icon:        query.Icon,
		IsNav:       query.IsNav,
	})
	if err != nil {
		return nil, err
	}
	rows := make([]*v1.ListDomains_Resp_Domain, 0, len(pageResp.Rows))
	for _, item := range pageResp.Rows {
		row := &v1.ListDomains_Resp_Domain{
			Id:          item.ID,
			Code:        item.Code,
			Name:        item.Name,
			Description: item.Description,
			Status:      enum.DomainStatusMap.MustToProto(item.Status),
			Url:         item.URL,
			Icon:        item.Icon,
			IsNav:       item.IsNav,
			Sort:        item.Sort,
			CreatedBy:   item.CreatedBy,
			UpdatedBy:   item.UpdatedBy,
		}
		if item.CreatedAt != nil {
			row.CreatedAt = timestamppb.New(*item.CreatedAt)
		}
		if item.UpdatedAt != nil {
			row.UpdatedAt = timestamppb.New(*item.UpdatedAt)
		}
		rows = append(rows, row)
	}
	return &v1.ListDomains_Resp{Rows: rows}, nil
}

func (s *DomainService) Page(ctx context.Context, req *v1.PageDomains_Req) (*v1.PageDomains_Resp, error) {
	query := req.GetQuery()
	if query == nil {
		query = &v1.PageDomains_Req_Query{}
	}
	var domainStatus *enum.DomainStatus
	if query.Status != nil {
		status, ok := enum.DomainStatusMap.ToEnum(*query.Status)
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_DOMAIN_INVALID)
		}
		domainStatus = &status
	}
	pageResp, err := s.contentUsecase.Page(ctx, &usecase.DomainPageReq{
		Page: &base.PageRequest{
			Page: int64(req.GetPage().GetPage()),
			Size: int64(req.GetPage().GetSize()),
		},
		DomainIDs:   query.GetIds(),
		Code:        query.Code,
		Name:        query.Name,
		Description: query.Description,
		Status:      domainStatus,
		URL:         query.Url,
		Icon:        query.Icon,
		IsNav:       query.IsNav,
	})
	if err != nil {
		return nil, err
	}
	rows := make([]*v1.PageDomains_Resp_Domain, 0, len(pageResp.Rows))
	for _, item := range pageResp.Rows {
		row := &v1.PageDomains_Resp_Domain{
			Id:          item.ID,
			Code:        item.Code,
			Name:        item.Name,
			Description: item.Description,
			Status:      enum.DomainStatusMap.MustToProto(item.Status),
			Url:         item.URL,
			Icon:        item.Icon,
			IsNav:       item.IsNav,
			Sort:        item.Sort,
			CreatedBy:   item.CreatedBy,
			UpdatedBy:   item.UpdatedBy,
		}
		if item.CreatedAt != nil {
			row.CreatedAt = timestamppb.New(*item.CreatedAt)
		}
		if item.UpdatedAt != nil {
			row.UpdatedAt = timestamppb.New(*item.UpdatedAt)
		}
		rows = append(rows, row)
	}
	return &v1.PageDomains_Resp{
		Page: &common.PageResp{
			Page:  uint32(pageResp.Page.Page),
			Size:  uint32(pageResp.Page.Size),
			Total: uint32(pageResp.Page.Total),
		},
		Rows: rows,
	}, nil
}
