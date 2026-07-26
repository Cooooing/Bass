package service

import (
	"bbs/internal/biz/usecase"
	bbscontentv1 "common/proto/gen/bbs/v1/content"
	bbscontentv1enum "common/proto/gen/bbs/v1/content/enum"
	"common/proto/gen/common"
	"context"

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

func (s *ContentDomainService) RegisterHttp(hs *http.Server) {
	bbscontentv1.RegisterDomainServiceHTTPServer(hs, s)
}

func (s *ContentDomainService) List(ctx context.Context, req *bbscontentv1.ListDomains_Req) (*bbscontentv1.ListDomains_Resp, error) {
	resp, err := s.contentDomainUsecase.ListDomains(ctx, &usecase.ListDomainsReq{
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
	rows := make([]*bbscontentv1.ListDomains_Resp_Domain, 0, len(resp.Rows))
	for _, row := range resp.Rows {
		if row == nil {
			rows = append(rows, nil)
			continue
		}
		rows = append(rows, &bbscontentv1.ListDomains_Resp_Domain{
			Id:          row.ID,
			Name:        row.Name,
			Description: row.Description,
			Status:      bbscontentv1enum.DomainStatus(row.Status),
			Url:         row.URL,
			Icon:        row.Icon,
			IsNav:       row.IsNav,
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
