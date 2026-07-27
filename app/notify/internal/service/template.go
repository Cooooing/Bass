package service

import (
	v1 "common/proto/gen/notify/v1"
	"context"
	"notify/internal/biz/usecase"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

type TemplateService struct {
	v1.UnimplementedNotifyTemplateServiceServer
	templateUsecase *usecase.TemplateUsecase
}

func NewTemplateService(
	templateUsecase *usecase.TemplateUsecase,
) *TemplateService {
	return &TemplateService{
		templateUsecase: templateUsecase,
	}
}

func (s *TemplateService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterNotifyTemplateServiceServer(gs, s)
}

func (s *TemplateService) RegisterHttp(hs *http.Server) {
}

func (s *TemplateService) Preview(ctx context.Context, req *v1.PreviewNotificationTemplate_Req) (*v1.PreviewNotificationTemplate_Resp, error) {
	resp, err := s.templateUsecase.Preview(ctx, &usecase.TemplatePreviewReq{
		Template:         req.GetTemplate(),
		TemplateDataJSON: req.GetTemplateDataJson(),
	})
	if err != nil {
		return nil, err
	}
	return &v1.PreviewNotificationTemplate_Resp{
		Rendered: resp.Rendered,
	}, nil
}
