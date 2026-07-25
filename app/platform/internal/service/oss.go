package service

import (
	"common/pkg/util"
	"common/proto/gen/common"
	v1 "common/proto/gen/platform/v1"
	"context"
	"errors"
	"io"
	"platform/internal/biz/model"
	"platform/internal/biz/usecase"
	"platform/internal/enum"

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

func (s *OssService) RegisterGrpc(
	gs *grpc.Server,
) {
	v1.RegisterPlatformOssServiceServer(gs, s)
}

func (s *OssService) RegisterHttp(
	hs *http.Server,
) {
}

func (s *OssService) Upload(
	ctx context.Context,
	req *v1.UploadOss_Req,
) (*v1.UploadOss_Resp, error) {
	row, err := s.objectStorageUsecase.Upload(ctx, &usecase.UploadReq{
		UserID:   req.GetUserId(),
		Key:      req.GetKey(),
		FileName: req.GetFileName(),
		MimeType: req.GetMimeType(),
		Content:  req.GetContent(),
	})
	if err != nil {
		return nil, err
	}
	return &v1.UploadOss_Resp{
		Id:       row.ID,
		Provider: enum.ObjectStorageProviderMap.MustToProto(row.Provider),
		Bucket:   row.Bucket,
		Key:      row.Key,
		MimeType: row.MimeType,
		Size:     row.Size,
		Hash:     row.Hash,
	}, nil
}

func (s *OssService) StreamUpload(
	stream v1.PlatformOssService_StreamUploadServer,
) error {
	first, err := stream.Recv()
	if err != nil {
		return err
	}
	meta := first.GetMeta()
	if meta == nil {
		return errors.New("stream upload meta is required")
	}
	reader, writer := io.Pipe()
	recvErr := make(chan error, 1)
	go func() {
		for {
			req, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				recvErr <- writer.Close()
				return
			}
			if err != nil {
				recvErr <- writer.CloseWithError(err)
				return
			}
			if req.GetMeta() != nil {
				recvErr <- writer.CloseWithError(errors.New("stream upload meta can only be sent once"))
				return
			}
			chunk := req.GetChunk()
			if len(chunk) == 0 {
				continue
			}
			if _, err := writer.Write(chunk); err != nil {
				recvErr <- err
				return
			}
		}
	}()
	row, err := s.objectStorageUsecase.StreamUpload(stream.Context(), &usecase.StreamUploadReq{
		UserID:   meta.GetUserId(),
		Key:      meta.GetKey(),
		FileName: meta.GetFileName(),
		MimeType: meta.GetMimeType(),
		Size:     meta.GetSize(),
		Body:     reader,
	})
	if err != nil {
		_ = reader.CloseWithError(err)
		return err
	}
	if err := <-recvErr; err != nil {
		return err
	}
	return stream.SendAndClose(&v1.StreamUploadOss_Resp{
		Id:       row.ID,
		Provider: enum.ObjectStorageProviderMap.MustToProto(row.Provider),
		Bucket:   row.Bucket,
		Key:      row.Key,
		MimeType: row.MimeType,
		Size:     row.Size,
		Hash:     row.Hash,
	})
}

func (s *OssService) Download(
	ctx context.Context,
	req *v1.DownloadOss_Req,
) (*v1.DownloadOss_Resp, error) {
	downloadResp, err := s.objectStorageUsecase.Download(ctx, req.GetKey())
	if err != nil {
		return nil, err
	}
	return &v1.DownloadOss_Resp{
		Key:      downloadResp.Key,
		MimeType: downloadResp.MimeType,
		Size:     downloadResp.Size,
		Content:  downloadResp.Content,
	}, nil
}

func (s *OssService) StreamDownload(
	req *v1.StreamDownloadOss_Req,
	stream v1.PlatformOssService_StreamDownloadServer,
) error {
	downloadResp, err := s.objectStorageUsecase.StreamDownload(stream.Context(), req.GetKey())
	if err != nil {
		return err
	}
	defer downloadResp.Body.Close()
	buf := make([]byte, 256*1024)
	for {
		n, err := downloadResp.Body.Read(buf)
		if n > 0 {
			if sendErr := stream.Send(&v1.StreamDownloadOss_Resp{
				Key:      downloadResp.Key,
				MimeType: downloadResp.MimeType,
				Size:     downloadResp.Size,
				Chunk:    append([]byte(nil), buf[:n]...),
			}); sendErr != nil {
				return sendErr
			}
		}
		if errors.Is(err, io.EOF) {
			return nil
		}
		if err != nil {
			return err
		}
	}
}

func (s *OssService) GetUploadToken(
	ctx context.Context,
	req *v1.GetUploadTokenOss_Req,
) (*v1.GetUploadTokenOss_Resp, error) {
	tokens, err := s.objectStorageUsecase.UploadToken(ctx, &usecase.UploadTokenReq{
		Num:    int(req.Num),
		UserID: req.GetUserId(),
	})
	if err != nil {
		return nil, err
	}
	uploadTokens := make([]*v1.GetUploadTokenOss_Resp_UploadToken, 0, len(tokens))
	for _, token := range tokens {
		uploadTokens = append(uploadTokens, &v1.GetUploadTokenOss_Resp_UploadToken{
			Key:   token.Key,
			Token: token.Token,
		})
	}
	return &v1.GetUploadTokenOss_Resp{
		UploadToken: uploadTokens,
	}, nil
}

func (s *OssService) Audit(
	ctx context.Context,
	req *v1.AuditOss_Req,
) (*v1.AuditOss_Resp, error) {
	err := s.objectStorageUsecase.UpdateAudit(ctx, &usecase.UpdateAuditReq{
		Key:    req.Key,
		Enable: req.Status,
		Reason: req.Reason,
		UserID: req.GetUserId(),
	})
	if err != nil {
		return nil, err
	}
	return &v1.AuditOss_Resp{}, nil
}

func (s *OssService) List(
	ctx context.Context,
	req *v1.ListOss_Req,
) (*v1.ListOss_Resp, error) {
	req.Page = util.OrDefault(req.Page, &common.PageReq{})
	req.Query = util.OrDefault(req.Query, &v1.ListOss_Req_OssQueryParams{})
	var provider *enum.ObjectStorageProvider
	if req.Query.Provider != nil {
		item, ok := enum.ObjectStorageProviderMap.ToEnum(*req.Query.Provider)
		if !ok {
			return nil, errors.New("invalid object storage provider")
		}
		provider = &item
	}
	pageResp, err := s.objectStorageUsecase.Page(ctx, &usecase.ObjectStoragePageReq{
		Page:     req.Page,
		Provider: provider,
		Bucket:   req.Query.Bucket,
		Key:      req.Query.Key,
		MimeType: req.Query.MimeType,
		Size:     req.Query.Size,
		Blocked:  req.Query.Blocked,
	})
	if err != nil {
		return nil, err
	}
	items := pageResp.Rows
	rows := make([]*v1.ListOss_Resp_Oss, 0, len(items))
	for _, item := range items {
		row := &v1.ListOss_Resp_Oss{
			CreatedAt:          timestamppb.New(*item.CreatedAt),
			UpdatedAt:          timestamppb.New(*item.UpdatedAt),
			Id:                 item.ID,
			Provider:           enum.ObjectStorageProviderMap.MustToProto(item.Provider),
			Bucket:             item.Bucket,
			Key:                item.Key,
			MimeType:           item.MimeType,
			Size:               item.Size,
			Hash:               item.Hash,
			AuditCallbackReply: item.AuditCallbackReply,
			Blocked:            item.Blocked,
			BlockedReason:      item.BlockedReason,
			BlockedBy:          item.BlockedBy,
		}
		if item.BlockedAt != nil {
			row.BlockedAt = timestamppb.New(*item.BlockedAt)
		}
		rows = append(rows, row)
	}
	return &v1.ListOss_Resp{
		Page: pageResp.Page,
		Rows: rows,
	}, nil
}

func (s *OssService) QiniuUploadCallback(
	ctx context.Context,
	req *v1.QiniuUploadCallbackOss_Req,
) (*v1.QiniuUploadCallbackOss_Resp, error) {
	s.log.Info("qiniu upload callback",
		"bucket", req.GetBucket(),
		"key", req.GetKey(),
		"mime_type", req.GetMimeType(),
		"size", req.GetSize(),
		"upload_by", req.GetUploadBy(),
	)

	err := s.objectStorageUsecase.QiniuUploadCallback(ctx, &model.ObjectStorage{
		Provider: enum.ObjectStorageProviderQiniu,
		Bucket:   req.Bucket,
		Key:      req.Key,
		MimeType: req.MimeType,
		Size:     req.Size,
		Hash:     req.Hash,
		UploadBy: req.UploadBy,
	})
	return &v1.QiniuUploadCallbackOss_Resp{}, err
}

func (s *OssService) QiniuIncrementAuditCallback(
	ctx context.Context,
	req *v1.QiniuIncrementAuditCallbackOss_Req,
) (*v1.QiniuIncrementAuditCallbackOss_Resp, error) {
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
		return &v1.QiniuIncrementAuditCallbackOss_Resp{}, err
	}
	return &v1.QiniuIncrementAuditCallbackOss_Resp{}, err
}
