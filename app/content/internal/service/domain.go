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

type DomainService struct {
	v1.UnimplementedContentDomainServiceServer
	contentUsecase *usecase.ContentUsecase
}

func (s *DomainService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterContentDomainServiceServer(gs, s)
}

func NewDomainService(
	contentUsecase *usecase.ContentUsecase,
) *DomainService {
	return &DomainService{
		contentUsecase: contentUsecase,
	}
}

func (s *DomainService) BatchCreate(ctx context.Context, req *v1.BatchCreateDomains_Request) (*v1.BatchCreateDomains_Reply, error) {
	domains := make([]*model.Domain, len(req.Domains))
	for i, d := range req.Domains {
		domainStatus, ok := enum.DomainStatusMap.ToEnum(util.DerefOrDefault(d.Status, v1.DomainStatus_DOMAIN_STATUS_NORMAL))
		if !ok {
			return nil, cerrors.ErrorBadRequest("invalid domain status")
		}
		domains[i] = &model.Domain{
			Name:        d.Name,
			Description: d.Description,
			Status:      domainStatus,
			URL:         d.Url,
			Icon:        d.Icon,
			IsNav:       d.IsNav,
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
		for _, tag := range save.Tags {
			row.Tags = append(row.Tags, &v1.Tag{
				CreatedAt:   timestamppb.New(*tag.CreatedAt),
				UpdatedAt:   timestamppb.New(*tag.UpdatedAt),
				CreatedBy:   tag.CreatedBy,
				UpdatedBy:   tag.UpdatedBy,
				Id:          tag.ID,
				Name:        tag.Name,
				Description: tag.Description,
				DomainId:    tag.DomainID,
				Status:      new(enum.TagStatusMap.MustToProto(tag.Status)),
			})
		}
		rows = append(rows, row)
	}
	return &v1.BatchCreateDomains_Reply{
		Rows: rows,
	}, nil
}

func (s *DomainService) Update(ctx context.Context, req *v1.UpdateDomain_Request) (*v1.UpdateDomain_Reply, error) {
	if req.Domain == nil {
		return nil, cerrors.ErrorBadRequest("domain is required")
	}
	domainStatus, ok := enum.DomainStatusMap.ToEnum(util.DerefOrDefault(req.Domain.Status, v1.DomainStatus_DOMAIN_STATUS_NORMAL))
	if !ok {
		return nil, cerrors.ErrorBadRequest("invalid domain status")
	}
	data, err := s.contentUsecase.Update(ctx, &model.Domain{
		ID:          req.Domain.Id,
		Name:        req.Domain.Name,
		Description: req.Domain.Description,
		Status:      domainStatus,
		URL:         req.Domain.Url,
		Icon:        req.Domain.Icon,
		IsNav:       req.Domain.IsNav,
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
	for _, tag := range data.Tags {
		reply.Tags = append(reply.Tags, &v1.Tag{
			CreatedAt:   timestamppb.New(*tag.CreatedAt),
			UpdatedAt:   timestamppb.New(*tag.UpdatedAt),
			CreatedBy:   tag.CreatedBy,
			UpdatedBy:   tag.UpdatedBy,
			Id:          tag.ID,
			Name:        tag.Name,
			Description: tag.Description,
			DomainId:    tag.DomainID,
			Status:      new(enum.TagStatusMap.MustToProto(tag.Status)),
		})
	}
	return &v1.UpdateDomain_Reply{
		Data: reply,
	}, err
}

func (s *DomainService) List(ctx context.Context, req *v1.ListDomains_Request) (*v1.ListDomains_Reply, error) {
	req.Query = util.OrDefault(req.Query, &v1.DomainQueryParams{})
	getReq := &repo.DomainGetReq{
		DomainIds:   req.Query.Ids,
		Name:        req.Query.Name,
		Description: req.Query.Description,
		Status:      new(v1.DomainStatus_DOMAIN_STATUS_NORMAL),
		Url:         req.Query.Url,
		Icon:        req.Query.Icon,
		IsNav:       req.Query.IsNav,
	}
	if req.Query.Status != nil {
		if _, ok := enum.DomainStatusMap.ToEnum(*req.Query.Status); !ok {
			return nil, cerrors.ErrorBadRequest("invalid domain status")
		}
		getReq.Status = new(*req.Query.Status)
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
		for _, tag := range datum.Tags {
			row.Tags = append(row.Tags, &v1.Tag{
				CreatedAt:   timestamppb.New(*tag.CreatedAt),
				UpdatedAt:   timestamppb.New(*tag.UpdatedAt),
				CreatedBy:   tag.CreatedBy,
				UpdatedBy:   tag.UpdatedBy,
				Id:          tag.ID,
				Name:        tag.Name,
				Description: tag.Description,
				DomainId:    tag.DomainID,
				Status:      new(enum.TagStatusMap.MustToProto(tag.Status)),
			})
		}
		reply = append(reply, row)
	}
	return &v1.ListDomains_Reply{
		Page: page,
		Rows: reply,
	}, err
}
