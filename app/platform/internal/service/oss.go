package service

import (
	"common/pkg/constant"
	"common/pkg/util"
	"common/proto/gen/common"
	v1 "common/proto/gen/notify/v1"
	"context"
	"platform/internal/biz/model"
	"platform/internal/biz/repo"
	"platform/internal/biz/usecase"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type OssService struct {
	v1.UnimplementedNotifyOssServiceServer
	log                  *log.Helper
	objectStorageUsecase *usecase.ObjectStorageUsecase
}

func NewOssService(
	logger log.Logger,
	objectStorageUsecase *usecase.ObjectStorageUsecase,
) *OssService {
	return &OssService{
		log:                  log.NewHelper(logger),
		objectStorageUsecase: objectStorageUsecase,
	}
}

func (s *OssService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterNotifyOssServiceServer(gs, s)
}

func (s *OssService) RegisterHttp(hs *http.Server) {
}

func (s *OssService) GetUploadToken(ctx context.Context, req *v1.GetUploadTokenOss_Request) (*v1.GetUploadTokenOss_Reply, error) {
	tokens, err := s.objectStorageUsecase.UploadToken(ctx, int(req.Num), req.GetUserId(), req.GetUserName())
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
	err := s.objectStorageUsecase.UpdateAudit(ctx, req.Key, req.Status, req.Reason, req.GetUserId(), req.GetUserName())
	if err != nil {
		return nil, err
	}
	return &v1.AuditOss_Reply{}, nil
}

func (s *OssService) List(ctx context.Context, req *v1.ListOss_Request) (*v1.ListOss_Reply, error) {
	req.Page = util.OrDefault(req.Page, &common.PageRequest{})
	req.Query = util.OrDefault(req.Query, &v1.OssQueryParams{})
	items, page, err := s.objectStorageUsecase.Page(ctx, req.Page, &repo.ObjectStorageGetReq{
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
	rows := make([]*v1.Oss, 0, len(items))
	for _, item := range items {
		row := &v1.Oss{
			CreatedAt:          timestamppb.New(*item.CreatedAt),
			UpdatedAt:          timestamppb.New(*item.UpdatedAt),
			Id:                 item.ID,
			Provider:           item.Provider,
			Bucket:             item.Bucket,
			Key:                item.Key,
			MimeType:           item.MimeType,
			Size:               item.Size,
			Hash:               item.Hash,
			AuditCallbackReply: item.AuditCallbackReply,
			Blocked:            item.Blocked,
			BlockedReason:      item.BlockedReason,
			BlockedBy:          item.BlockedBy,
			BlockedByName:      item.BlockedByName,
		}
		if item.BlockedAt != nil {
			row.BlockedAt = timestamppb.New(*item.BlockedAt)
		}
		rows = append(rows, row)
	}
	return &v1.ListOss_Reply{
		Page: page,
		Rows: rows,
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

	err = s.objectStorageUsecase.QiniuUploadCallback(ctx, &model.ObjectStorage{
		Provider:     constant.Qiniu.String(),
		Bucket:       req.Bucket,
		Key:          req.Key,
		MimeType:     req.MimeType,
		Size:         req.Size,
		Hash:         req.Hash,
		UploadBy:     req.UploadBy,
		UploadByName: req.UploadByName,
	})
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
		err := s.objectStorageUsecase.QiniuIncrementAuditCallback(ctx, req.InputKey, string(bytes), true)
		return &v1.QiniuIncrementAuditCallbackOss_Reply{}, err
	}
	return &v1.QiniuIncrementAuditCallbackOss_Reply{}, err
}
