package service

import (
	v1 "common/api/gen/notify/v1"
	"context"
	"notify/internal/biz/domain"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type EmailService struct {
	v1.UnimplementedNotifyEmailServiceServer
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
	v1.RegisterNotifyEmailServiceServer(gs, s)
}

func (s *EmailService) RegisterHttp(hs *http.Server) {
}

func (s *EmailService) Send(ctx context.Context, req *v1.SendEmail_Request) (rsp *v1.SendEmail_Reply, err error) {
	if len(req.To) > 0 {
		err = s.emailDomain.Send(ctx, req.To, req.Title, req.Content)
		if err != nil {
			return nil, err
		}
	}
	return &v1.SendEmail_Reply{}, nil
}
