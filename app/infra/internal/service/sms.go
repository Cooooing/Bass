package service

import (
	v1 "common/gen/infra/v1"
	"context"
	"infra/internal/biz/domain"

	"github.com/go-kratos/kratos/v2/transport/grpc"
)

type SmsService struct {
	v1.UnimplementedInfraSmsServiceServer
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
	v1.RegisterInfraSmsServiceServer(gs, s)
}

func (s *SmsService) Send(ctx context.Context, req *v1.SendSmsRequest) (rsp *v1.SendSmsReply, err error) {
	if len(req.Phone) > 0 {
		err = s.smsDomain.Send(ctx, req.Phone, req.Params)
		if err != nil {
			return nil, err
		}
	}
	return &v1.SendSmsReply{}, nil
}
