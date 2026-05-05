package domain

import (
	"common/api/gen/common"
	cerrors "common/api/gen/common/errors"
	"common/pkg/constant"
	commonModel "common/pkg/model"
	"common/pkg/util"
	"context"

	domainbase "notify/internal/biz/base"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	"notify/internal/data/ent"
	"notify/internal/data/ent/gen"
	"notify/internal/data/oss"
	"time"

	"github.com/google/uuid"
)

type ObjectStorageDomain struct {
	*domainbase.BaseDomain
	objectStorageRepo     repo.ObjectStorageRepo
	objectStorageProvider repo.ObjectStorageProvider
}

func NewObjectStorageDomain(base *domainbase.BaseDomain, objectStorageRepo repo.ObjectStorageRepo, ossFactory *oss.Factory) *ObjectStorageDomain {
	return &ObjectStorageDomain{
		BaseDomain:            base,
		objectStorageRepo:     objectStorageRepo,
		objectStorageProvider: ossFactory.Get(base.Conf.Server.Oss.Provider),
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
	return ent.WithTx(ctx, d.Db, func(tx *gen.Client) error {
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
	return d.objectStorageRepo.GetPage(ctx, d.Db, page, req)
}

func (d *ObjectStorageDomain) QiniuUploadCallback(ctx context.Context, o *model.ObjectStorage) error {
	return ent.WithTx(ctx, d.Db, func(tx *gen.Client) error {
		_, err := d.objectStorageRepo.Save(ctx, tx, o)
		return err
	})
}

func (d *ObjectStorageDomain) QiniuIncrementAuditCallback(ctx context.Context, key string, reply string, blocked bool) error {
	return ent.WithTx(ctx, d.Db, func(tx *gen.Client) error {
		return d.objectStorageRepo.UpdateAudit(ctx, tx, &model.ObjectStorage{ObjectStorage: &gen.ObjectStorage{
			Key:                key,
			AuditCallbackReply: new(reply),
			Blocked:            blocked,
		}})
	})
}
