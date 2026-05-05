package service

import (
	v1 "common/api/gen/notify/v1"
	"context"
	"notify/internal/biz/domain"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type SmsService struct {
	v1.UnimplementedNotifySmsServiceServer
	*BaseService
	smsDomain *domain.TencentSmsDomain
}

func NewSmsService(baseService *BaseService, smsDomain *domain.TencentSmsDomain) *SmsService {
	return &SmsService{
		BaseService: baseService,
		smsDomain:   smsDomain,
	}
}

func (s *SmsService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterNotifySmsServiceServer(gs, s)
}

func (s *SmsService) RegisterHttp(hs *http.Server) {
}

func (s *SmsService) Send(ctx context.Context, req *v1.SendSms_Request) (rsp *v1.SendSms_Reply, err error) {
	if len(req.Phone) > 0 {
		err = s.smsDomain.Send(ctx, req.Phone, req.Params)
		if err != nil {
			return nil, err
		}
	}
	return &v1.SendSms_Reply{}, nil
}
