package service

import (
	"common/pkg/constant"
	"common/pkg/util"
	"common/proto/gen/common"
	v1 "common/proto/gen/platform/v1"
	"context"
	"platform/internal/biz/model"
	"platform/internal/biz/usecase"

	"log/slog"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type OssService struct {
	v1.UnimplementedPlatformOssServiceServer
	log                  *slog.Logger
	objectStorageUsecase *usecase.ObjectStorageUsecase
}

func NewOssService(
	logger *slog.Logger,
	objectStorageUsecase *usecase.ObjectStorageUsecase,
) *OssService {
	return &OssService{
		log:                  logger,
		objectStorageUsecase: objectStorageUsecase,
	}
}

func (s *OssService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterPlatformOssServiceServer(gs, s)
}

func (s *OssService) RegisterHttp(hs *http.Server) {
}

func (s *OssService) GetUploadToken(ctx context.Context, req *v1.GetUploadTokenOss_Request) (*v1.GetUploadTokenOss_Response, error) {
	uploadTokenResponse, err := s.objectStorageUsecase.UploadToken(ctx, &usecase.UploadTokenReq{
		Num:      int(req.Num),
		UserID:   req.GetUserId(),
		UserName: req.GetUserName(),
	})
	if err != nil {
		return nil, err
	}
	tokens := uploadTokenResponse.Rows
	uploadTokens := make([]*v1.GetUploadTokenOss_Response_UploadToken, 0, len(tokens))
	for _, token := range tokens {
		uploadTokens = append(uploadTokens, &v1.GetUploadTokenOss_Response_UploadToken{
			Key:   token.Key,
			Token: token.Token,
		})
	}
	return &v1.GetUploadTokenOss_Response{
		UploadToken: uploadTokens,
	}, nil
}

func (s *OssService) Audit(ctx context.Context, req *v1.AuditOss_Request) (*v1.AuditOss_Response, error) {
	err := s.objectStorageUsecase.UpdateAudit(ctx, &usecase.UpdateAuditReq{
		Key:      req.Key,
		Enable:   req.Status,
		Reason:   req.Reason,
		UserID:   req.GetUserId(),
		UserName: req.GetUserName(),
	})
	if err != nil {
		return nil, err
	}
	return &v1.AuditOss_Response{}, nil
}

func (s *OssService) List(ctx context.Context, req *v1.ListOss_Request) (*v1.ListOss_Response, error) {
	req.Page = util.OrDefault(req.Page, &common.PageRequest{})
	req.Query = util.OrDefault(req.Query, &v1.ListOss_Request_OssQueryParams{})
	pageResponse, err := s.objectStorageUsecase.Page(ctx, &usecase.ObjectStoragePageReq{
		Page:          req.Page,
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
	items := pageResponse.Rows
	rows := make([]*v1.ListOss_Response_Oss, 0, len(items))
	for _, item := range items {
		row := &v1.ListOss_Response_Oss{
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
	return &v1.ListOss_Response{
		Page: pageResponse.Page,
		Rows: rows,
	}, nil
}

func (s *OssService) QiniuUploadCallback(ctx context.Context, req *v1.QiniuUploadCallbackOss_Request) (*v1.QiniuUploadCallbackOss_Response, error) {
	s.log.Info("qiniu upload callback",
		"bucket", req.GetBucket(),
		"key", req.GetKey(),
		"mime_type", req.GetMimeType(),
		"size", req.GetSize(),
		"upload_by", req.GetUploadBy(),
	)

	err := s.objectStorageUsecase.QiniuUploadCallback(ctx, &usecase.QiniuUploadCallbackReq{
		ObjectStorage: &model.ObjectStorage{
			Provider:     constant.Qiniu.String(),
			Bucket:       req.Bucket,
			Key:          req.Key,
			MimeType:     req.MimeType,
			Size:         req.Size,
			Hash:         req.Hash,
			UploadBy:     req.UploadBy,
			UploadByName: req.UploadByName,
		},
	})
	return &v1.QiniuUploadCallbackOss_Response{}, err
}
func (s *OssService) QiniuIncrementAuditCallback(ctx context.Context, req *v1.QiniuIncrementAuditCallbackOss_Request) (*v1.QiniuIncrementAuditCallbackOss_Response, error) {
	opts := protojson.MarshalOptions{
		EmitUnpopulated: true,
	}
	bytes, err := opts.Marshal(req)
	if err != nil {
		return nil, err
	}
	itemsCount := len(req.GetItems())
	suggestion := ""
	if itemsCount > 0 && req.GetItems()[0].GetResult() != nil && req.GetItems()[0].GetResult().GetResult() != nil {
		suggestion = req.GetItems()[0].GetResult().GetResult().GetSuggestion()
	}
	s.log.Info("qiniu increment audit callback",
		"input_key", req.GetInputKey(),
		"items_count", itemsCount,
		"suggestion", suggestion,
	)

	if suggestion == "block" {
		err := s.objectStorageUsecase.QiniuIncrementAuditCallback(ctx, &usecase.QiniuIncrementAuditCallbackReq{
			Key:     req.InputKey,
			Reply:   string(bytes),
			Blocked: true,
		})
		return &v1.QiniuIncrementAuditCallbackOss_Response{}, err
	}
	return &v1.QiniuIncrementAuditCallbackOss_Response{}, err
}
