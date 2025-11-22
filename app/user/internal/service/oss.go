package service

import (
	v1 "common/api/user/v1"
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/qiniu/go-sdk/v7/storagev2/credentials"
	"github.com/qiniu/go-sdk/v7/storagev2/uptoken"
)

type OssService struct {
	v1.UnimplementedUserOssServiceServer
	*BaseService
}

func NewOssService(baseService *BaseService) *OssService {
	return &OssService{
		BaseService: baseService,
	}
}

func (s *OssService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterUserOssServiceServer(gs, s)
}

func (s *OssService) RegisterHttp(hs *http.Server) {
	v1.RegisterUserOssServiceHTTPServer(hs, s)
}

func (s *OssService) UploadToken(ctx context.Context, req *v1.UploadTokenRequest) (*v1.UploadTokenReply, error) {
	mac := credentials.NewCredentials(s.conf.Oss.Qiniu.AccessKey, s.conf.Oss.Qiniu.SecretKey)
	putPolicy, err := uptoken.NewPutPolicy(s.conf.Oss.Qiniu.Bucket, time.Now().Add(s.conf.Oss.Qiniu.Timeout.AsDuration()))
	if err != nil {
		return nil, err
	}
	upToken, err := uptoken.NewSigner(putPolicy, mac).GetUpToken(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.UploadTokenReply{
		Token: upToken,
	}, nil
}
