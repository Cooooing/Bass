package service

import (
	v1 "common/api/user/v1"
	"common/pkg/constant"
	"context"
	"user/internal/biz"
	"user/internal/biz/model"
	"user/internal/data/ent/gen"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type OssService struct {
	v1.UnimplementedUserOssServiceServer
	*BaseService
	ossObjectStorageDomain *biz.ObjectStorageDomain
}

func NewOssService(baseService *BaseService, ossObjectStorageDomain *biz.ObjectStorageDomain) *OssService {
	return &OssService{
		BaseService:            baseService,
		ossObjectStorageDomain: ossObjectStorageDomain,
	}
}

func (s *OssService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterUserOssServiceServer(gs, s)
}

func (s *OssService) RegisterHttp(hs *http.Server) {
	v1.RegisterUserOssServiceHTTPServer(hs, s)
}

func (s *OssService) UploadToken(ctx context.Context, req *v1.UploadTokenRequest) (*v1.UploadTokenReply, error) {
	tokens, err := s.ossObjectStorageDomain.UploadToken(ctx, int(req.Num))
	if err != nil {
		return nil, err
	}
	uploadTokens := make([]*v1.UploadTokenReply_UploadToken, 0, len(tokens))
	for _, token := range tokens {
		uploadTokens = append(uploadTokens, &v1.UploadTokenReply_UploadToken{
			Key:   token.Key,
			Token: token.Token,
		})
	}
	return &v1.UploadTokenReply{
		UploadToken: uploadTokens,
	}, nil
}

func (s *OssService) QiniuUploadCallback(ctx context.Context, req *v1.QiniuUploadCallbackRequest) (*v1.QiniuUploadCallbackReply, error) {
	err := s.ossObjectStorageDomain.QiniuUploadCallback(ctx, &model.ObjectStorage{ObjectStorage: &gen.ObjectStorage{
		Provider: constant.Qiniu.String(),
		Bucket:   req.Bucket,
		Key:      req.Key,
		MimeType: req.MimeType,
		Size:     req.Size,
		Hash:     req.Hash,
	}})
	return &v1.QiniuUploadCallbackReply{}, err
}
