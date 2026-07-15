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

func (d *ObjectStorageUsecase) UploadToken(ctx context.Context, num int, userID int64, userName string) ([]*model.UploadToken, error) {
	tokens := make([]*model.UploadToken, 0, num)
	for range num {
		key := uuid.New().String()
		token, err := d.objectStorageClient.UploadToken(ctx, key, userID, userName)
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

func (d *ObjectStorageUsecase) UpdateAudit(ctx context.Context, key string, enable bool, reason *string, userID int64, userName string) error {
	err := d.objectStorageClient.Status(ctx, key, enable)
	if err != nil {
		return err
	}
	return d.tx(ctx, func(ctx context.Context) error {
		return d.objectStorageRepo.UpdateAudit(ctx, &model.ObjectStorage{
			Key:           key,
			Blocked:       enable,
			BlockedReason: reason,
			BlockedAt:     new(time.Now()),
			BlockedBy:     new(userID),
			BlockedByName: new(userName),
		})
	})
}

func (d *ObjectStorageUsecase) Page(ctx context.Context, page *common.PageRequest, req *repo.ObjectStorageGetReq) ([]*model.ObjectStorage, *common.PageReply, error) {
	var (
		rows      []*model.ObjectStorage
		pageReply *common.PageReply
	)
	err := d.tx(ctx, func(ctx context.Context) error {
		var err error
		rows, pageReply, err = d.objectStorageRepo.Page(ctx, page, req)
		return err
	})
	return rows, pageReply, err
}

func (d *ObjectStorageUsecase) QiniuUploadCallback(ctx context.Context, o *model.ObjectStorage) error {
	return d.tx(ctx, func(ctx context.Context) error {
		_, err := d.objectStorageRepo.Save(ctx, o)
		return err
	})
}

func (d *ObjectStorageUsecase) QiniuIncrementAuditCallback(ctx context.Context, key string, reply string, blocked bool) error {
	return d.tx(ctx, func(ctx context.Context) error {
		return d.objectStorageRepo.UpdateAudit(ctx, &model.ObjectStorage{
			Key:                key,
			AuditCallbackReply: new(reply),
			Blocked:            blocked,
		})
	})
}
