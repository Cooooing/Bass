package service

import (
	"common/pkg/apperror"
	"common/pkg/constant"
	commonmodel "common/pkg/model"
	"common/pkg/util"
	bbscontentv1enum "common/proto/gen/bff_bbs/content/enum/v1"
	bbscontentv1 "common/proto/gen/bff_bbs/content/v1"
	cerrors "common/proto/gen/common/errors/v1"
	"common/proto/gen/common/v1"
	"context"

	"bff_bbs/internal/biz/usecase"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ContentDomainService struct {
	bbscontentv1.UnimplementedDomainServiceServer
	contentDomainUsecase *usecase.ContentDomainUsecase
}

func NewContentDomainService(
	contentDomainUsecase *usecase.ContentDomainUsecase,
) *ContentDomainService {
	return &ContentDomainService{
		contentDomainUsecase: contentDomainUsecase,
	}
}

func (s *ContentDomainService) RegisterGrpc(gs *grpc.Server) {
}

func (s *ContentDomainService) RegisterHTTP(hs *http.Server) {
	bbscontentv1.RegisterDomainServiceHTTPServer(hs, s)
}

func (s *ContentDomainService) CreateDomain(ctx context.Context, req *bbscontentv1.CreateDomain_Req) (*bbscontentv1.CreateDomain_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	var domainSave *usecase.ContentDomainSave
	if domain := req.GetDomain(); domain != nil {
		domainSave = &usecase.ContentDomainSave{
			Code:        domain.GetCode(),
			Name:        domain.GetName(),
			Description: domain.Description,
			Status:      domain.Status,
			URL:         domain.Url,
			Icon:        domain.Icon,
			IsNav:       domain.GetIsNav(),
			Sort:        domain.GetSort(),
		}
	}
	resp, err := s.contentDomainUsecase.CreateDomain(ctx, &usecase.CreateDomainReq{
		UserID: user.ID,
		Domain: domainSave,
	})
	if err != nil {
		return nil, err
	}
	var domain *bbscontentv1.CreateDomain_Resp_Domain
	if resp != nil {
		domain = &bbscontentv1.CreateDomain_Resp_Domain{
			Id:          resp.ID,
			Code:        resp.Code,
			Name:        resp.Name,
			Description: resp.Description,
			Status:      bbscontentv1enum.DomainStatus(resp.Status),
			Url:         resp.URL,
			Icon:        resp.Icon,
			IsNav:       resp.IsNav,
			Sort:        resp.Sort,
			CreatedBy:   resp.CreatedBy,
			UpdatedBy:   resp.UpdatedBy,
		}
		if resp.CreatedAt != nil {
			domain.CreatedAt = timestamppb.New(*resp.CreatedAt)
		}
		if resp.UpdatedAt != nil {
			domain.UpdatedAt = timestamppb.New(*resp.UpdatedAt)
		}
	}
	return &bbscontentv1.CreateDomain_Resp{
		Domain: domain,
	}, nil
}

func (s *ContentDomainService) UpdateDomain(ctx context.Context, req *bbscontentv1.UpdateDomain_Req) (*bbscontentv1.UpdateDomain_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	var domainSave *usecase.ContentDomainSave
	if domain := req.GetDomain(); domain != nil {
		domainSave = &usecase.ContentDomainSave{
			Code:        domain.GetCode(),
			Name:        domain.GetName(),
			Description: domain.Description,
			Status:      domain.Status,
			URL:         domain.Url,
			Icon:        domain.Icon,
			IsNav:       domain.GetIsNav(),
			Sort:        domain.GetSort(),
		}
	}
	resp, err := s.contentDomainUsecase.UpdateDomain(ctx, &usecase.UpdateDomainReq{
		UserID:   user.ID,
		DomainID: req.GetDomainId(),
		Domain:   domainSave,
	})
	if err != nil {
		return nil, err
	}
	var domain *bbscontentv1.UpdateDomain_Resp_Domain
	if resp != nil {
		domain = &bbscontentv1.UpdateDomain_Resp_Domain{
			Id:          resp.ID,
			Code:        resp.Code,
			Name:        resp.Name,
			Description: resp.Description,
			Status:      bbscontentv1enum.DomainStatus(resp.Status),
			Url:         resp.URL,
			Icon:        resp.Icon,
			IsNav:       resp.IsNav,
			Sort:        resp.Sort,
			CreatedBy:   resp.CreatedBy,
			UpdatedBy:   resp.UpdatedBy,
		}
		if resp.CreatedAt != nil {
			domain.CreatedAt = timestamppb.New(*resp.CreatedAt)
		}
		if resp.UpdatedAt != nil {
			domain.UpdatedAt = timestamppb.New(*resp.UpdatedAt)
		}
	}
	return &bbscontentv1.UpdateDomain_Resp{
		Domain: domain,
	}, nil
}

func (s *ContentDomainService) ListDomains(ctx context.Context, req *bbscontentv1.ListDomains_Req) (*bbscontentv1.ListDomains_Resp, error) {
	resp, err := s.contentDomainUsecase.ListDomains(ctx, &usecase.ListDomainsReq{
		Page:  commonmodel.PageReqFromProto(req.GetPage()),
		Query: req.GetQuery(),
	})
	if err != nil {
		return nil, err
	}
	var page *common.PageResp
	if resp.Page != nil {
		page = resp.Page.ToProto()
	}
	rows := make([]*bbscontentv1.ListDomains_Resp_Domain, 0, len(resp.Rows))
	for _, row := range resp.Rows {
		if row == nil {
			rows = append(rows, nil)
			continue
		}
		rows = append(rows, &bbscontentv1.ListDomains_Resp_Domain{
			Id:          row.ID,
			Code:        row.Code,
			Name:        row.Name,
			Description: row.Description,
			Status:      bbscontentv1enum.DomainStatus(row.Status),
			Url:         row.URL,
			Icon:        row.Icon,
			IsNav:       row.IsNav,
			Sort:        row.Sort,
			CreatedBy:   row.CreatedBy,
			UpdatedBy:   row.UpdatedBy,
			CreatedAt:   timestamppb.New(*row.CreatedAt),
			UpdatedAt:   timestamppb.New(*row.UpdatedAt),
		})
	}
	return &bbscontentv1.ListDomains_Resp{
		Page: page,
		Rows: rows,
	}, nil
}
