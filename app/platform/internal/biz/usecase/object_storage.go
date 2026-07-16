package usecase

import (
	"common/proto/gen/common"
	"context"

	base "platform/internal/biz/base"
	"platform/internal/biz/model"
	"platform/internal/biz/repo"
	"platform/internal/config"
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

type UploadTokenReq struct {
	Num      int
	UserID   int64
	UserName string
}

type UploadTokenResponse struct {
	Rows []*model.UploadToken
}

func (d *ObjectStorageUsecase) UploadToken(ctx context.Context, req *UploadTokenReq) (*UploadTokenResponse, error) {
	if req == nil {
		req = &UploadTokenReq{}
	}
	tokens := make([]*model.UploadToken, 0, req.Num)
	for range req.Num {
		key := uuid.New().String()
		tokenResp, err := d.objectStorageClient.UploadToken(ctx, &repo.ObjectStorageUploadTokenReq{Key: key, UploaderID: req.UserID, UploaderName: req.UserName})
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, &model.UploadToken{
			Key:   key,
			Token: tokenResp.Token,
		})
	}
	return &UploadTokenResponse{Rows: tokens}, nil
}

type UpdateAuditReq struct {
	Key      string
	Enable   bool
	Reason   *string
	UserID   int64
	UserName string
}

func (d *ObjectStorageUsecase) UpdateAudit(ctx context.Context, req *UpdateAuditReq) error {
	if req == nil {
		req = &UpdateAuditReq{}
	}
	_, err := d.objectStorageClient.Status(ctx, &repo.ObjectStorageStatusReq{Key: req.Key, Enable: req.Enable})
	if err != nil {
		return err
	}
	return d.tx(ctx, func(ctx context.Context) error {
		_, err := d.objectStorageRepo.UpdateAudit(ctx, &repo.ObjectStorageUpdateAuditReq{Row: &model.ObjectStorage{
			Key:           req.Key,
			Blocked:       req.Enable,
			BlockedReason: req.Reason,
			BlockedAt:     new(time.Now()),
			BlockedBy:     new(req.UserID),
			BlockedByName: new(req.UserName),
		}})
		return err
	})
}

type ObjectStoragePageReq struct {
	Page          *common.PageRequest
	Provider      *string
	Bucket        *string
	Key           *string
	MimeType      *string
	Size          *common.Int64Range
	Blocked       *bool
	BlockedByName *string
}

type ObjectStoragePageResponse struct {
	Rows []*model.ObjectStorage
	Page *common.PageResponse
}

func (d *ObjectStorageUsecase) Page(ctx context.Context, req *ObjectStoragePageReq) (*ObjectStoragePageResponse, error) {
	if req == nil {
		req = &ObjectStoragePageReq{}
	}
	var (
		rows         []*model.ObjectStorage
		pageResponse *common.PageResponse
	)
	err := d.tx(ctx, func(ctx context.Context) error {
		pageResp, err := d.objectStorageRepo.Page(ctx, &repo.ObjectStoragePageReq{
			Page: req.Page,
			ObjectStorageGetReq: repo.ObjectStorageGetReq{
				Provider:      req.Provider,
				Bucket:        req.Bucket,
				Key:           req.Key,
				MimeType:      req.MimeType,
				Size:          req.Size,
				Blocked:       req.Blocked,
				BlockedByName: req.BlockedByName,
			},
		})
		if err != nil {
			return err
		}
		rows = pageResp.Rows
		pageResponse = pageResp.Page
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &ObjectStoragePageResponse{Rows: rows, Page: pageResponse}, nil
}

type QiniuUploadCallbackReq struct {
	ObjectStorage *model.ObjectStorage
}

func (d *ObjectStorageUsecase) QiniuUploadCallback(ctx context.Context, req *QiniuUploadCallbackReq) error {
	if req == nil {
		req = &QiniuUploadCallbackReq{}
	}
	return d.tx(ctx, func(ctx context.Context) error {
		_, err := d.objectStorageRepo.Save(ctx, &repo.ObjectStorageSaveReq{Row: req.ObjectStorage})
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
		_, err := d.objectStorageRepo.UpdateAudit(ctx, &repo.ObjectStorageUpdateAuditReq{Row: &model.ObjectStorage{
			Key:                req.Key,
			AuditCallbackReply: new(req.Reply),
			Blocked:            req.Blocked,
		}})
		return err
	})
}
