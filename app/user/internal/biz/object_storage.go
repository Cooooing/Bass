package biz

import (
	"context"
	"encoding/json"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/data/ent"
	"user/internal/data/ent/gen"
	"user/internal/data/oss"

	"github.com/google/uuid"
)

type ObjectStorageDomain struct {
	*BaseDomain
	objectStorageRepo repo.ObjectStorageRepo
}

func NewObjectStorageDomain(base *BaseDomain, ossFactory *oss.Factory) *ObjectStorageDomain {
	return &ObjectStorageDomain{
		BaseDomain:        base,
		objectStorageRepo: ossFactory.Get(base.conf.Oss.Provider),
	}
}

func (d *ObjectStorageDomain) UploadToken(ctx context.Context, num int) ([]*model.UploadToken, error) {
	tokens := make([]*model.UploadToken, 0, num)
	for range num {
		key := uuid.New().String()
		tokens = append(tokens, &model.UploadToken{
			Key:   key,
			Token: d.objectStorageRepo.UploadToken(key),
		})
	}
	return tokens, nil
}

func (d *ObjectStorageDomain) QiniuUploadCallback(ctx context.Context, o *model.ObjectStorage) error {
	marshal, _ := json.Marshal(o)
	d.log.Infof("QiniuUploadCallback: %s", marshal)

	err := ent.WithTx(ctx, d.db, func(tx *gen.Client) error {
		_, err := d.objectStorageRepo.Save(ctx, tx, o)
		return err
	})
	return err
}
