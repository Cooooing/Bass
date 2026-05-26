package usecase

import (
	"common/api/gen/common"
	cerrors "common/api/gen/common/errors"
	"common/pkg/constant"
	commonModel "common/pkg/model"
	"common/pkg/util"
	"context"

	base "notify/internal/biz/base"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	"notify/internal/conf"
	"time"

	"github.com/google/uuid"
)

type ObjectStorageUsecase struct {
	conf                  *conf.Bootstrap
	tx                    base.Tx
	objectStorageRepo     repo.ObjectStorageRepo
	objectStorageProvider repo.ObjectStorageProvider
}

func NewObjectStorageUsecase(
	conf *conf.Bootstrap,
	tx base.Tx,
	objectStorageRepo repo.ObjectStorageRepo,
	objectStorageProvider repo.ObjectStorageProvider,
) *ObjectStorageUsecase {
	return &ObjectStorageUsecase{
		conf:                  conf,
		tx:                    tx,
		objectStorageRepo:     objectStorageRepo,
		objectStorageProvider: objectStorageProvider,
	}
}

func (d *ObjectStorageUsecase) UploadToken(ctx context.Context, num int) ([]*model.UploadToken, error) {
	tokens := make([]*model.UploadToken, 0, num)
	for range num {
		key := uuid.New().String()
		token, err := d.objectStorageProvider.UploadToken(ctx, key)
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

func (d *ObjectStorageUsecase) UpdateAudit(ctx context.Context, key string, enable bool, reason *string) error {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return cerrors.ErrorUnauthorized("user not login")
	}

	err := d.objectStorageProvider.Status(ctx, key, enable)
	if err != nil {
		return err
	}
	return d.tx(ctx, func(ctx context.Context) error {
		return d.objectStorageRepo.UpdateAudit(ctx, &model.ObjectStorage{
			Key:           key,
			Blocked:       enable,
			BlockedReason: reason,
			BlockedAt:     new(time.Now()),
			BlockedBy:     new(user.ID),
			BlockedByName: new(user.Name),
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
		rows, pageReply, err = d.objectStorageRepo.GetPage(ctx, page, req)
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
