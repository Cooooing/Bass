package service

import (
	"bbs/internal/biz/usecase"
	bbscontentv1 "common/proto/gen/bbs/v1/content"
	"common/proto/gen/common"
	"context"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

type ContentDomainService struct {
	bbscontentv1.UnimplementedDomainServiceServer
	contentDomainUsecase *usecase.ContentDomainUsecase
}

func NewContentDomainService(contentDomainUsecase *usecase.ContentDomainUsecase) *ContentDomainService {
	return &ContentDomainService{contentDomainUsecase: contentDomainUsecase}
}

func (s *ContentDomainService) RegisterGrpc(gs *grpc.Server) {}

func (s *ContentDomainService) RegisterHttp(hs *http.Server) {
	bbscontentv1.RegisterDomainServiceHTTPServer(hs, s)
}

func (s *ContentDomainService) List(ctx context.Context, req *bbscontentv1.ListDomains_Request) (*bbscontentv1.ListDomains_Response, error) {
	response, err := s.contentDomainUsecase.ListDomains(ctx, &usecase.ListDomainsReq{Page: req.GetPage(), Query: req.GetQuery()})
	if err != nil {
		return nil, err
	}
	var page *common.PageResponse
	if response.Page != nil {
		page = &common.PageResponse{Page: response.Page.Page, Size: response.Page.Size, Total: response.Page.Total}
	}
	rows := make([]*bbscontentv1.ListDomains_Response_Domain, 0, len(response.Rows))
	for _, row := range response.Rows {
		if row == nil {
			rows = append(rows, nil)
			continue
		}
		rows = append(rows, &bbscontentv1.ListDomains_Response_Domain{
			Id:          row.ID,
			Name:        row.Name,
			Description: row.Description,
			Status:      bbscontentv1.DomainStatus(row.Status),
			Url:         row.URL,
			Icon:        row.Icon,
			IsNav:       row.IsNav,
			CreatedBy:   row.CreatedBy,
			UpdatedBy:   row.UpdatedBy,
			CreatedAt:   row.CreatedAt,
			UpdatedAt:   row.UpdatedAt,
		})
	}
	return &bbscontentv1.ListDomains_Response{Page: page, Rows: rows}, nil
}
