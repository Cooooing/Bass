package service

import (
	v1 "common/api/infra/v1"
	"context"
	"infra/internal/biz/domain"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type EmailService struct {
	v1.UnimplementedInfraEmailServiceServer
	*BaseService
	emailDomain *domain.EmailDomain
}

func NewEmailService(baseService *BaseService, emailDomain *domain.EmailDomain) *EmailService {
	return &EmailService{
		BaseService: baseService,
		emailDomain: emailDomain,
	}
}

func (s *EmailService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterInfraEmailServiceServer(gs, s)
}

func (s *EmailService) RegisterHttp(hs *http.Server) {
	v1.RegisterInfraEmailServiceHTTPServer(hs, s)
}

func (s *EmailService) Send(ctx context.Context, req *v1.SendEmailRequest) (rsp *v1.SendEmailReply, err error) {
	if len(req.To) > 0 {
		err = s.emailDomain.Send(ctx, req.To, req.Title, req.Content)
		if err != nil {
			return nil, err
		}
	}
	return &v1.SendEmailReply{}, nil
}
