package service

import (
	"common/api/gen/common"
	v1 "common/api/gen/notify/v1"
	"common/pkg/constant"
	commonModel "common/pkg/model"
	"common/pkg/util"
	"context"
	"notify/internal/biz/domain"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	"notify/internal/data/ent/gen"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/protobuf/encoding/protojson"
)

type OssService struct {
	v1.UnimplementedNotifyOssServiceServer
	log                    *log.Helper
	ossObjectStorageDomain *domain.ObjectStorageDomain
}

func NewOssService(
	logger log.Logger,
	ossObjectStorageDomain *domain.ObjectStorageDomain,
) *OssService {
	return &OssService{
		log:                    log.NewHelper(logger),
		ossObjectStorageDomain: ossObjectStorageDomain,
	}
}

func (s *OssService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterNotifyOssServiceServer(gs, s)
}

func (s *OssService) RegisterHttp(hs *http.Server) {
	v1.RegisterNotifyOssServiceHTTPServer(hs, s)
}

func (s *OssService) GetUploadToken(ctx context.Context, req *v1.GetUploadTokenOss_Request) (*v1.GetUploadTokenOss_Reply, error) {
	tokens, err := s.ossObjectStorageDomain.UploadToken(ctx, int(req.Num))
	if err != nil {
		return nil, err
	}
	uploadTokens := make([]*v1.UploadToken, 0, len(tokens))
	for _, token := range tokens {
		uploadTokens = append(uploadTokens, &v1.UploadToken{
			Key:   token.Key,
			Token: token.Token,
		})
	}
	return &v1.GetUploadTokenOss_Reply{
		UploadToken: uploadTokens,
	}, nil
}

func (s *OssService) Audit(ctx context.Context, req *v1.AuditOss_Request) (*v1.AuditOss_Reply, error) {
	err := s.ossObjectStorageDomain.UpdateAudit(ctx, req.Key, req.Status, req.Reason)
	if err != nil {
		return nil, err
	}
	return &v1.AuditOss_Reply{}, nil
}

func (s *OssService) Page(ctx context.Context, req *v1.PageOssOss_Request) (*v1.PageOssOss_Reply, error) {
	req.Page = util.OrDefault(req.Page, &common.PageRequest{})
	req.Query = util.OrDefault(req.Query, &v1.OssQueryParams{})
	items, page, err := s.ossObjectStorageDomain.Page(ctx, req.Page, &repo.ObjectStorageGetReq{
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
	return &v1.PageOssOss_Reply{
		Page: page,
		Rows: commonModel.ConvertToRpcList(items),
	}, nil
}

func (s *OssService) QiniuUploadCallback(ctx context.Context, req *v1.QiniuUploadCallbackOss_Request) (*v1.QiniuUploadCallbackOss_Reply, error) {
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
	return &v1.QiniuUploadCallbackOss_Reply{}, err
}

func (s *OssService) QiniuIncrementAuditCallback(ctx context.Context, req *v1.QiniuIncrementAuditCallbackOss_Request) (*v1.QiniuIncrementAuditCallbackOss_Reply, error) {
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
		return &v1.QiniuIncrementAuditCallbackOss_Reply{}, err
	}
	return &v1.QiniuIncrementAuditCallbackOss_Reply{}, err
}
