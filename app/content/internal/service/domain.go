package service

import (
	cv1 "common/api/common/v1"
	v1 "common/api/content/v1"
	commonModel "common/pkg/model"
	"common/pkg/util/base"
	"content/internal/biz"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type DomainService struct {
	v1.UnimplementedContentDomainServiceServer
	*BaseService
	domainDomain *biz.DomainDomain
}

func (s *DomainService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterContentDomainServiceServer(gs, s)
}

func (s *DomainService) RegisterHttp(hs *http.Server) {
	v1.RegisterContentDomainServiceHTTPServer(hs, s)
}

func NewDomainService(baseService *BaseService, domainDomain *biz.DomainDomain) *DomainService {
	return &DomainService{
		BaseService:  baseService,
		domainDomain: domainDomain,
	}
}

func (s *DomainService) Adds(ctx context.Context, req *v1.AddDomainsRequest) (*v1.AddDomainsReply, error) {
	domains := make([]*model.Domain, len(req.Domains))
	for i, domain := range req.Domains {
		domains[i] = &model.Domain{
			Name:        domain.Name,
			Description: domain.Description,
			Status:      base.DerefOrDefault(domain.Status, int32(cv1.DomainStatus_DomainNormal)),
			URL:         domain.Url,
			Icon:        domain.Icon,
			IsNav:       domain.IsNav,
		}
	}
	_, err := s.domainDomain.Adds(ctx, domains)
	if err != nil {
		return nil, err
	}
	return &v1.AddDomainsReply{}, nil
}

func (s *DomainService) Update(ctx context.Context, req *v1.UpdateDomainRequest) (*v1.UpdateDomainReply, error) {
	data, err := s.domainDomain.Update(ctx, &model.Domain{
		Name:        req.Domain.Name,
		Description: req.Domain.Description,
		Status:      base.DerefOrDefault(req.Domain.Status, int32(cv1.DomainStatus_DomainNormal)),
		URL:         req.Domain.Url,
		Icon:        req.Domain.Icon,
		IsNav:       req.Domain.IsNav,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UpdateDomainReply{
		Data: &v1.DomainReply{
			CreatedAt:   timestamppb.New(*data.CreatedAt),
			UpdatedAt:   timestamppb.New(*data.UpdatedAt),
			Id:          data.ID,
			Name:        data.Name,
			Description: data.Description,
			Status:      data.Status,
			Url:         data.URL,
			Icon:        data.Icon,
			TagCount:    data.TagCount,
			IsNav:       data.IsNav,
		},
	}, err
}

func (s *DomainService) Page(ctx context.Context, req *v1.PageDomainRequest) (*v1.PageDomainReply, error) {
	getReq := &repo.DomainGetReq{
		Ids:         req.Ids,
		Name:        req.Name,
		Description: req.Description,
		Url:         req.Url,
		Icon:        req.Icon,
		IsNav:       req.IsNav,
	}
	if req.Status != nil {
		getReq.Status = base.Ptr(cv1.DomainStatus(*req.Status))
	}
	if req.TagCount != nil {
		getReq.TagCount = &commonModel.Range[int32]{}
		if req.TagCount.Start != nil {
			getReq.TagCount.Start = req.TagCount.Start
		}
		if req.TagCount.End != nil {
			getReq.TagCount.End = req.TagCount.End
		}
	}
	data, page, err := s.domainDomain.Page(ctx, req.Page, getReq)
	reply := make([]*v1.DomainReply, len(data))
	for i, datum := range data {
		replyTags := make([]*v1.TagReply, len(datum.Edges.Tags))
		for j, tag := range datum.Edges.Tags {
			replyTags[j] = &v1.TagReply{
				CreatedAt:    timestamppb.New(*tag.CreatedAt),
				UpdatedAt:    timestamppb.New(*tag.UpdatedAt),
				CreatedBy:    tag.CreatedBy,
				UpdatedBy:    tag.UpdatedBy,
				Id:           tag.ID,
				Name:         tag.Name,
				Description:  tag.Description,
				DomainId:     tag.DomainID,
				Status:       tag.Status,
				ArticleCount: tag.ArticleCount,
			}
		}
		reply[i] = &v1.DomainReply{
			CreatedAt:   timestamppb.New(*datum.CreatedAt),
			UpdatedAt:   timestamppb.New(*datum.UpdatedAt),
			Id:          datum.ID,
			Name:        datum.Name,
			Description: datum.Description,
			Status:      datum.Status,
			Url:         datum.URL,
			Icon:        datum.Icon,
			TagCount:    datum.TagCount,
			IsNav:       datum.IsNav,
			Tags:        replyTags,
		}
	}
	return &v1.PageDomainReply{
		Page:    page,
		Domains: reply,
	}, err
}
