package domain

import (
	cv1 "common/api/common/v1"
	"common/pkg/constant"
	commonModel "common/pkg/model"
	"common/pkg/util"

	"context"

	doaminbase "infra/internal/biz/base"
	"infra/internal/biz/model"
	"infra/internal/biz/repo"
	"infra/internal/data/ent"
	"infra/internal/data/ent/gen"
	"infra/internal/data/oss"
	"time"

	"github.com/google/uuid"
)

type ObjectStorageDomain struct {
	*doaminbase.BaseDomain
	objectStorageRepo     repo.ObjectStorageRepo
	objectStorageProvider repo.ObjectStorageProvider
}

func NewObjectStorageDomain(base *doaminbase.BaseDomain, objectStorageRepo repo.ObjectStorageRepo, ossFactory *oss.Factory) *ObjectStorageDomain {
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
	var err error
	infra, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return cv1.ErrorUnauthorized("infra not login")
	}

	err = d.objectStorageProvider.Status(ctx, key, enable)
	if err != nil {
		return err
	}
	err = ent.WithTx(ctx, d.Db, func(tx *gen.Client) error {
		err := d.objectStorageRepo.UpdateAudit(ctx, tx, &model.ObjectStorage{ObjectStorage: &gen.ObjectStorage{
			Key:           key,
			Blocked:       enable,
			BlockedReason: reason,
			BlockedAt:     util.Ptr(time.Now()),
			BlockedBy:     util.Ptr(infra.ID),
			BlockedByName: util.Ptr(infra.Name),
		}})
		return err
	})
	return err
}

func (d *ObjectStorageDomain) Page(ctx context.Context, page *cv1.PageRequest, req *repo.ObjectStorageGetReq) ([]*model.ObjectStorage, *cv1.PageReply, error) {
	return d.objectStorageRepo.GetPage(ctx, d.Db, page, req)
}

func (d *ObjectStorageDomain) QiniuUploadCallback(ctx context.Context, o *model.ObjectStorage) error {
	err := ent.WithTx(ctx, d.Db, func(tx *gen.Client) error {
		_, err := d.objectStorageRepo.Save(ctx, tx, o)
		return err
	})
	return err
}

func (d *ObjectStorageDomain) QiniuIncrementAuditCallback(ctx context.Context, key string, reply string, blocked bool) error {
	err := ent.WithTx(ctx, d.Db, func(tx *gen.Client) error {
		err := d.objectStorageRepo.UpdateAudit(ctx, tx, &model.ObjectStorage{ObjectStorage: &gen.ObjectStorage{
			Key:                key,
			AuditCallbackReply: util.Ptr(reply),
			Blocked:            blocked,
		}})
		return err
	})
	return err
}
