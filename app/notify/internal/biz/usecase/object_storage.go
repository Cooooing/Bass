package usecase

import (
	"common/api/gen/common"
	cerrors "common/api/gen/common/errors"
	"common/pkg/constant"
	commonModel "common/pkg/model"
	"common/pkg/util"
	utilent "common/pkg/util/ent"
	"context"
	"errors"
	"notify/internal/data/oss"

	base "notify/internal/biz/base"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	"notify/internal/conf"
	"notify/internal/data/gen"
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
	ossFactory *oss.Factory,
) *ObjectStorageUsecase {
	return &ObjectStorageUsecase{
		conf:                  conf,
		tx:                    tx,
		objectStorageRepo:     objectStorageRepo,
		objectStorageProvider: ossFactory.Get(conf.Server.Oss.Provider),
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
		c, ok := utilent.ClientFromCtx[*gen.Client](ctx)
		if !ok {
			return errors.New("no transaction in context")
		}
		return d.objectStorageRepo.UpdateAudit(ctx, c, &model.ObjectStorage{ObjectStorage: &gen.ObjectStorage{
			Key:           key,
			Blocked:       enable,
			BlockedReason: reason,
			BlockedAt:     new(time.Now()),
			BlockedBy:     new(user.ID),
			BlockedByName: new(user.Name),
		}})
	})
}

func (d *ObjectStorageUsecase) Page(ctx context.Context, page *common.PageRequest, req *repo.ObjectStorageGetReq) ([]*model.ObjectStorage, *common.PageReply, error) {
	c, ok := utilent.ClientFromCtx[*gen.Client](ctx)
	if !ok {
		return nil, nil, errors.New("no client in context")
	}
	return d.objectStorageRepo.GetPage(ctx, c, page, req)
}

func (d *ObjectStorageUsecase) QiniuUploadCallback(ctx context.Context, o *model.ObjectStorage) error {
	return d.tx(ctx, func(ctx context.Context) error {
		c, ok := utilent.ClientFromCtx[*gen.Client](ctx)
		if !ok {
			return errors.New("no transaction in context")
		}
		_, err := d.objectStorageRepo.Save(ctx, c, o)
		return err
	})
}

func (d *ObjectStorageUsecase) QiniuIncrementAuditCallback(ctx context.Context, key string, reply string, blocked bool) error {
	return d.tx(ctx, func(ctx context.Context) error {
		c, ok := utilent.ClientFromCtx[*gen.Client](ctx)
		if !ok {
			return errors.New("no transaction in context")
		}
		return d.objectStorageRepo.UpdateAudit(ctx, c, &model.ObjectStorage{ObjectStorage: &gen.ObjectStorage{
			Key:                key,
			AuditCallbackReply: new(reply),
			Blocked:            blocked,
		}})
	})
}
