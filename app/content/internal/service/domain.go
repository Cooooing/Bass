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

type DomainService struct {
	v1.UnimplementedContentDomainServiceServer
	contentUsecase *usecase.ContentUsecase
}

func (s *DomainService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterContentDomainServiceServer(gs, s)
}

func (s *DomainService) RegisterHttp(hs *http.Server) {}

func NewDomainService(
	contentUsecase *usecase.ContentUsecase,
) *DomainService {
	return &DomainService{
		contentUsecase: contentUsecase,
	}
}

func (s *DomainService) BatchCreate(ctx context.Context, req *v1.BatchCreateDomains_Request) (*v1.BatchCreateDomains_Reply, error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	domains := make([]*model.Domain, len(req.Domains))
	for i, d := range req.Domains {
		domainStatus, ok := enum.DomainStatusMap.ToEnum(util.DerefOrDefault(d.Status, v1.DomainStatus_DOMAIN_STATUS_ENABLED))
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_DOMAIN_INVALID)
		}
		domains[i] = &model.Domain{
			Name:        d.Name,
			Description: d.Description,
			Status:      domainStatus,
			URL:         d.Url,
			Icon:        d.Icon,
			IsNav:       d.IsNav,
			CreatedBy:   new(req.UserId),
			UpdatedBy:   new(req.UserId),
		}
	}
	saves, err := s.contentUsecase.Adds(ctx, domains)
	if err != nil {
		return nil, err
	}
	rows := make([]*v1.Domain, 0, len(saves))
	for _, save := range saves {
		row := &v1.Domain{
			CreatedAt:   timestamppb.New(*save.CreatedAt),
			UpdatedAt:   timestamppb.New(*save.UpdatedAt),
			CreatedBy:   save.CreatedBy,
			UpdatedBy:   save.UpdatedBy,
			Id:          save.ID,
			Name:        save.Name,
			Description: save.Description,
			Status:      enum.DomainStatusMap.MustToProto(save.Status),
			Url:         save.URL,
			Icon:        save.Icon,
			IsNav:       save.IsNav,
		}
		rows = append(rows, row)
	}
	return &v1.BatchCreateDomains_Reply{
		Rows: rows,
	}, nil
}

func (s *DomainService) Update(ctx context.Context, req *v1.UpdateDomain_Request) (*v1.UpdateDomain_Reply, error) {
	if req.Domain == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_DOMAIN_INVALID)
	}
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	domainStatus, ok := enum.DomainStatusMap.ToEnum(util.DerefOrDefault(req.Domain.Status, v1.DomainStatus_DOMAIN_STATUS_ENABLED))
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_DOMAIN_INVALID)
	}
	data, err := s.contentUsecase.Update(ctx, &model.Domain{
		ID:          req.Domain.Id,
		Name:        req.Domain.Name,
		Description: req.Domain.Description,
		Status:      domainStatus,
		URL:         req.Domain.Url,
		Icon:        req.Domain.Icon,
		IsNav:       req.Domain.IsNav,
		UpdatedBy:   new(req.UserId),
	})
	if err != nil {
		return nil, err
	}
	reply := &v1.Domain{
		CreatedAt:   timestamppb.New(*data.CreatedAt),
		UpdatedAt:   timestamppb.New(*data.UpdatedAt),
		CreatedBy:   data.CreatedBy,
		UpdatedBy:   data.UpdatedBy,
		Id:          data.ID,
		Name:        data.Name,
		Description: data.Description,
		Status:      enum.DomainStatusMap.MustToProto(data.Status),
		Url:         data.URL,
		Icon:        data.Icon,
		IsNav:       data.IsNav,
	}
	return &v1.UpdateDomain_Reply{
		Domain: reply,
	}, err
}

func (s *DomainService) List(ctx context.Context, req *v1.ListDomains_Request) (*v1.ListDomains_Reply, error) {
	req.Query = util.OrDefault(req.Query, &v1.DomainQueryParams{})
	var domainStatus *enum.DomainStatus
	if req.Query.Status != nil {
		status, ok := enum.DomainStatusMap.ToEnum(*req.Query.Status)
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_DOMAIN_INVALID)
		}
		domainStatus = new(status)
	}
	getReq := &repo.DomainGetReq{
		DomainIds:   req.Query.Ids,
		Name:        req.Query.Name,
		Description: req.Query.Description,
		Status:      domainStatus,
		Url:         req.Query.Url,
		Icon:        req.Query.Icon,
		IsNav:       req.Query.IsNav,
	}
	data, _, err := s.contentUsecase.Page(ctx, server.GetPageMax(), getReq)
	reply := make([]*v1.Domain, 0, len(data))
	for _, datum := range data {
		row := &v1.Domain{
			CreatedAt:   timestamppb.New(*datum.CreatedAt),
			UpdatedAt:   timestamppb.New(*datum.UpdatedAt),
			CreatedBy:   datum.CreatedBy,
			UpdatedBy:   datum.UpdatedBy,
			Id:          datum.ID,
			Name:        datum.Name,
			Description: datum.Description,
			Status:      enum.DomainStatusMap.MustToProto(datum.Status),
			Url:         datum.URL,
			Icon:        datum.Icon,
			IsNav:       datum.IsNav,
		}
		reply = append(reply, row)
	}
	return &v1.ListDomains_Reply{
		Rows: reply,
	}, err
}

func (s *DomainService) Page(ctx context.Context, req *v1.PageDomains_Request) (*v1.PageDomains_Reply, error) {
	req.Query = util.OrDefault(req.Query, &v1.DomainQueryParams{})
	var domainStatus *enum.DomainStatus
	if req.Query.Status != nil {
		status, ok := enum.DomainStatusMap.ToEnum(*req.Query.Status)
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_DOMAIN_INVALID)
		}
		domainStatus = new(status)
	}
	getReq := &repo.DomainGetReq{
		DomainIds:   req.Query.Ids,
		Name:        req.Query.Name,
		Description: req.Query.Description,
		Status:      domainStatus,
		Url:         req.Query.Url,
		Icon:        req.Query.Icon,
		IsNav:       req.Query.IsNav,
	}
	data, page, err := s.contentUsecase.Page(ctx, req.Page, getReq)
	reply := make([]*v1.Domain, 0, len(data))
	for _, datum := range data {
		row := &v1.Domain{
			CreatedAt:   timestamppb.New(*datum.CreatedAt),
			UpdatedAt:   timestamppb.New(*datum.UpdatedAt),
			CreatedBy:   datum.CreatedBy,
			UpdatedBy:   datum.UpdatedBy,
			Id:          datum.ID,
			Name:        datum.Name,
			Description: datum.Description,
			Status:      enum.DomainStatusMap.MustToProto(datum.Status),
			Url:         datum.URL,
			Icon:        datum.Icon,
			IsNav:       datum.IsNav,
		}
		reply = append(reply, row)
	}
	return &v1.PageDomains_Reply{
		Page: page,
		Rows: reply,
	}, err
}
