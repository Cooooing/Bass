package service

import (
	cv1 "common/api/common/v1"
	v1 "common/api/infra/v1"
	"common/pkg/constant"
	"common/pkg/cutil/base"
	commonModel "common/pkg/model"
	"context"
	"infra/internal/biz"
	"infra/internal/biz/model"
	"infra/internal/biz/repo"
	"infra/internal/data/ent/gen"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/protobuf/encoding/protojson"
)

type OssService struct {
	v1.UnimplementedInfraOssServiceServer
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
	v1.RegisterInfraOssServiceServer(gs, s)
}

func (s *OssService) RegisterHttp(hs *http.Server) {
	v1.RegisterInfraOssServiceHTTPServer(hs, s)
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

func (s *OssService) Audit(ctx context.Context, req *v1.AuditRequest) (*v1.AuditReply, error) {
	err := s.ossObjectStorageDomain.UpdateAudit(ctx, req.Key, req.Status, req.Reason)
	if err != nil {
		return nil, err
	}
	return &v1.AuditReply{}, nil
}

func (s *OssService) Page(ctx context.Context, req *v1.PageOssRequest) (*v1.PageOssReply, error) {
	req.Page = base.OrDefault(req.Page, &cv1.PageRequest{})
	req.Query = base.OrDefault(req.Query, &v1.OssQueryParams{})
	ObjectStorages, page, err := s.ossObjectStorageDomain.Page(ctx, req.Page, &repo.ObjectStorageGetReq{
		Provider:      req.Query.Provider,
		Bucket:        req.Query.Bucket,
		Key:           req.Query.Key,
		MimeType:      req.Query.MimeType,
		Size:          req.Query.Size,
		Blocked:       req.Query.Blocked,
		BlockedByName: req.Query.BlockedByName,
	})
	if err != nil {
		return nil, err
	}
	return &v1.PageOssReply{
		Page: page,
		Rows: commonModel.ConvertToRpcList(ObjectStorages),
	}, nil
}

func (s *OssService) QiniuUploadCallback(ctx context.Context, req *v1.QiniuUploadCallbackRequest) (*v1.QiniuUploadCallbackReply, error) {
	var err error

	opts := protojson.MarshalOptions{
		EmitUnpopulated: true,
	}
	bytes, err := opts.Marshal(req)
	if err != nil {
		return nil, err
	}
	s.log.Infof("qiniu upload callback: %s", string(bytes))

	err = s.ossObjectStorageDomain.QiniuUploadCallback(ctx, &model.ObjectStorage{ObjectStorage: &gen.ObjectStorage{
		Provider:     constant.Qiniu.String(),
		Bucket:       req.Bucket,
		Key:          req.Key,
		MimeType:     req.MimeType,
		Size:         req.Size,
		Hash:         req.Hash,
		UploadBy:     req.UploadBy,
		UploadByName: req.UploadByName,
	}})
	return &v1.QiniuUploadCallbackReply{}, err
}

func (s *OssService) QiniuIncrementAuditCallback(ctx context.Context, req *v1.QiniuIncrementAuditCallbackRequest) (*v1.QiniuIncrementAuditCallbackReply, error) {
	var err error
	opts := protojson.MarshalOptions{
		EmitUnpopulated: true,
	}
	bytes, err := opts.Marshal(req)
	if err != nil {
		return nil, err
	}
	s.log.Infof("qiniu increment audit callback: %s", string(bytes))

	suggestion := req.Items[0].Result.Result.Suggestion
	if suggestion == "block" {
		err := s.ossObjectStorageDomain.QiniuIncrementAuditCallback(ctx, req.InputKey, string(bytes), true)
		return &v1.QiniuIncrementAuditCallbackReply{}, err
	}
	return &v1.QiniuIncrementAuditCallbackReply{}, err
}
