package usecase

import (
	"common/proto/gen/common"
	"context"
	"fmt"
	"io"

	"platform/internal/biz/base"
	"platform/internal/biz/model"
	"platform/internal/biz/repo"
	"platform/internal/config"
	"platform/internal/enum"
	"time"

	"github.com/google/uuid"
)

type ObjectStorageUsecase struct {
	conf                *config.Bootstrap
	tx                  base.Tx
	objectStorageRepo   repo.ObjectStorageRepo
	objectStorageClient repo.ObjectStorageClient
}

func NewObjectStorageUsecase(
	conf *config.Bootstrap,
	tx base.Tx,
	objectStorageRepo repo.ObjectStorageRepo,
	objectStorageClient repo.ObjectStorageClient,
) *ObjectStorageUsecase {
	return &ObjectStorageUsecase{
		conf:                conf,
		tx:                  tx,
		objectStorageRepo:   objectStorageRepo,
		objectStorageClient: objectStorageClient,
	}
}

func (d *ObjectStorageUsecase) CreateBucket(ctx context.Context, bucket string) error {
	return d.objectStorageClient.CreateBucket(ctx, bucket)
}

func (d *ObjectStorageUsecase) DeleteBucket(ctx context.Context, bucket string) error {
	return d.objectStorageClient.DeleteBucket(ctx, bucket)
}

type UploadReq struct {
	UserID   int64
	Key      string
	FileName string
	MimeType string
	Content  []byte
}

func (d *ObjectStorageUsecase) Upload(ctx context.Context, req *UploadReq) (*model.ObjectStorage, error) {
	if req == nil {
		req = &UploadReq{}
	}
	key := req.Key
	if key == "" {
		key = uuid.New().String()
	}
	uploadResp, err := d.objectStorageClient.Upload(ctx, &repo.ObjectStorageUploadReq{
		Key:        key,
		FileName:   req.FileName,
		MimeType:   req.MimeType,
		Content:    req.Content,
		UploaderID: req.UserID,
	})
	if err != nil {
		return nil, err
	}
	row := &model.ObjectStorage{
		Provider: uploadResp.Provider,
		Bucket:   uploadResp.Bucket,
		Key:      uploadResp.Key,
		MimeType: uploadResp.MimeType,
		Size:     uploadResp.Size,
		Hash:     uploadResp.Hash,
		UploadBy: req.UserID,
	}
	err = d.tx(ctx, func(ctx context.Context) error {
		saved, err := d.objectStorageRepo.Save(ctx, row)
		if err != nil {
			return err
		}
		row = saved
		return nil
	})
	if err != nil {
		return nil, err
	}
	return row, nil
}

type StreamUploadReq struct {
	UserID   int64
	Key      string
	FileName string
	MimeType string
	Size     int64
	Body     io.Reader
}

func (d *ObjectStorageUsecase) StreamUpload(ctx context.Context, req *StreamUploadReq) (*model.ObjectStorage, error) {
	if req == nil {
		req = &StreamUploadReq{}
	}
	if req.Body == nil {
		return nil, fmt.Errorf("stream upload body is nil")
	}
	key := req.Key
	if key == "" {
		key = uuid.New().String()
	}
	uploadResp, err := d.objectStorageClient.StreamUpload(ctx, &repo.ObjectStorageStreamUploadReq{
		Key:        key,
		FileName:   req.FileName,
		MimeType:   req.MimeType,
		Size:       req.Size,
		Body:       req.Body,
		UploaderID: req.UserID,
	})
	if err != nil {
		return nil, err
	}
	row := &model.ObjectStorage{
		Provider: uploadResp.Provider,
		Bucket:   uploadResp.Bucket,
		Key:      uploadResp.Key,
		MimeType: uploadResp.MimeType,
		Size:     uploadResp.Size,
		Hash:     uploadResp.Hash,
		UploadBy: req.UserID,
	}
	err = d.tx(ctx, func(ctx context.Context) error {
		saved, err := d.objectStorageRepo.Save(ctx, row)
		if err != nil {
			return err
		}
		row = saved
		return nil
	})
	if err != nil {
		return nil, err
	}
	return row, nil
}

type DownloadResp struct {
	Key      string
	MimeType string
	Size     int64
	Content  []byte
}

func (d *ObjectStorageUsecase) Download(ctx context.Context, key string) (*DownloadResp, error) {
	downloadResp, err := d.objectStorageClient.Download(ctx, key)
	if err != nil {
		return nil, err
	}
	return &DownloadResp{
		Key:      downloadResp.Key,
		MimeType: downloadResp.MimeType,
		Size:     downloadResp.Size,
		Content:  downloadResp.Content,
	}, nil
}

type StreamDownloadResp struct {
	Key      string
	MimeType string
	Size     int64
	Body     io.ReadCloser
}

func (d *ObjectStorageUsecase) StreamDownload(ctx context.Context, key string) (*StreamDownloadResp, error) {
	downloadResp, err := d.objectStorageClient.StreamDownload(ctx, key)
	if err != nil {
		return nil, err
	}
	return &StreamDownloadResp{
		Key:      downloadResp.Key,
		MimeType: downloadResp.MimeType,
		Size:     downloadResp.Size,
		Body:     downloadResp.Body,
	}, nil
}

type UploadTokenReq struct {
	Num    int
	UserID int64
}

func (d *ObjectStorageUsecase) UploadToken(ctx context.Context, req *UploadTokenReq) ([]*model.UploadToken, error) {
	if req == nil {
		req = &UploadTokenReq{}
	}
	tokens := make([]*model.UploadToken, 0, req.Num)
	for range req.Num {
		key := uuid.New().String()
		token, err := d.objectStorageClient.UploadToken(ctx, &repo.ObjectStorageUploadTokenReq{
			Key:        key,
			UploaderID: req.UserID,
		})
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, &model.UploadToken{
			Key:   key,
			Token: token,
		})
	}
	return tokens, nil
}

type UpdateAuditReq struct {
	Key    string
	Enable bool
	Reason *string
	UserID int64
}

func (d *ObjectStorageUsecase) UpdateAudit(ctx context.Context, req *UpdateAuditReq) error {
	if req == nil {
		req = &UpdateAuditReq{}
	}
	if err := d.objectStorageClient.Status(ctx, &repo.ObjectStorageStatusReq{
		Key:    req.Key,
		Enable: req.Enable,
	}); err != nil {
		return err
	}
	return d.tx(ctx, func(ctx context.Context) error {
		err := d.objectStorageRepo.UpdateAudit(ctx, &model.ObjectStorage{
			Key:           req.Key,
			Blocked:       req.Enable,
			BlockedReason: req.Reason,
			BlockedAt:     new(time.Now()),
			BlockedBy:     new(req.UserID),
		})
		return err
	})
}

type ObjectStoragePageReq struct {
	Page     *common.PageReq
	Provider *enum.ObjectStorageProvider
	Bucket   *string
	Key      *string
	MimeType *string
	Size     *common.Int64Range
	Blocked  *bool
}

type ObjectStoragePageResp struct {
	Rows []*model.ObjectStorage
	Page *common.PageResp
}

func (d *ObjectStorageUsecase) Page(ctx context.Context, req *ObjectStoragePageReq) (*ObjectStoragePageResp, error) {
	if req == nil {
		req = &ObjectStoragePageReq{}
	}
	var (
		rows     []*model.ObjectStorage
		pageResp *common.PageResp
	)
	err := d.tx(ctx, func(ctx context.Context) error {
		resp, err := d.objectStorageRepo.Page(ctx, &repo.ObjectStoragePageReq{
			Page: req.Page,
			ObjectStorageGetReq: repo.ObjectStorageGetReq{
				Provider: req.Provider,
				Bucket:   req.Bucket,
				Key:      req.Key,
				MimeType: req.MimeType,
				Size:     req.Size,
				Blocked:  req.Blocked,
			},
		})
		if err != nil {
			return err
		}
		rows = resp.Rows
		pageResp = resp.Page
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &ObjectStoragePageResp{
		Rows: rows,
		Page: pageResp,
	}, nil
}

func (d *ObjectStorageUsecase) QiniuUploadCallback(ctx context.Context, row *model.ObjectStorage) error {
	return d.tx(ctx, func(ctx context.Context) error {
		_, err := d.objectStorageRepo.Save(ctx, row)
		return err
	})
}

type QiniuIncrementAuditCallbackReq struct {
	Key     string
	Reply   string
	Blocked bool
}

func (d *ObjectStorageUsecase) QiniuIncrementAuditCallback(ctx context.Context, req *QiniuIncrementAuditCallbackReq) error {
	if req == nil {
		req = &QiniuIncrementAuditCallbackReq{}
	}
	return d.tx(ctx, func(ctx context.Context) error {
		err := d.objectStorageRepo.UpdateAudit(ctx, &model.ObjectStorage{
			Key:                req.Key,
			AuditCallbackReply: new(req.Reply),
			Blocked:            req.Blocked,
		})
		return err
	})
}
