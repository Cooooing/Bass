package domain

import (
	"common/api/gen/common"
	cerrors "common/api/gen/common/errors"
	"common/pkg/constant"
	commonModel "common/pkg/model"
	"common/pkg/util"
	"context"
	"notify/internal/data/client"
	"notify/internal/data/oss"

	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	"notify/internal/conf"
	"notify/internal/data/gen"
	"time"

	"github.com/google/uuid"
)

type ObjectStorageDomain struct {
	conf                  *conf.Bootstrap
	db                    *gen.Client
	objectStorageRepo     repo.ObjectStorageRepo
	objectStorageProvider repo.ObjectStorageProvider
}

func NewObjectStorageDomain(
	conf *conf.Bootstrap,
	db *gen.Client,
	objectStorageRepo repo.ObjectStorageRepo,
	ossFactory *oss.Factory,
) *ObjectStorageDomain {
	return &ObjectStorageDomain{
		conf:                  conf,
		db:                    db,
		objectStorageRepo:     objectStorageRepo,
		objectStorageProvider: ossFactory.Get(conf.Server.Oss.Provider),
	}
}

func (d *ObjectStorageDomain) UploadToken(ctx context.Context, num int) ([]*model.UploadToken, error) {
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

func (d *ObjectStorageDomain) UpdateAudit(ctx context.Context, key string, enable bool, reason *string) error {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return cerrors.ErrorUnauthorized("user not login")
	}

	err := d.objectStorageProvider.Status(ctx, key, enable)
	if err != nil {
		return err
	}
	return client.WithTx(ctx, d.db, func(tx *gen.Client) error {
		return d.objectStorageRepo.UpdateAudit(ctx, tx, &model.ObjectStorage{ObjectStorage: &gen.ObjectStorage{
			Key:           key,
			Blocked:       enable,
			BlockedReason: reason,
			BlockedAt:     new(time.Now()),
			BlockedBy:     new(user.ID),
			BlockedByName: new(user.Name),
		}})
	})
}

func (d *ObjectStorageDomain) Page(ctx context.Context, page *common.PageRequest, req *repo.ObjectStorageGetReq) ([]*model.ObjectStorage, *common.PageReply, error) {
	return d.objectStorageRepo.GetPage(ctx, d.db, page, req)
}

func (d *ObjectStorageDomain) QiniuUploadCallback(ctx context.Context, o *model.ObjectStorage) error {
	return client.WithTx(ctx, d.db, func(tx *gen.Client) error {
		_, err := d.objectStorageRepo.Save(ctx, tx, o)
		return err
	})
}

func (d *ObjectStorageDomain) QiniuIncrementAuditCallback(ctx context.Context, key string, reply string, blocked bool) error {
	return client.WithTx(ctx, d.db, func(tx *gen.Client) error {
		return d.objectStorageRepo.UpdateAudit(ctx, tx, &model.ObjectStorage{ObjectStorage: &gen.ObjectStorage{
			Key:                key,
			AuditCallbackReply: new(reply),
			Blocked:            blocked,
		}})
	})
}
