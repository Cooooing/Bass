package minio

import (
	"common/pkg/constant"
	"context"
	"platform/internal/biz/model"
	"platform/internal/biz/repo"
	"platform/internal/data/gen"
)

type Minio struct {
}

func NewMinio() *Minio {
	return &Minio{}
}

func (m *Minio) Name() string {
	return constant.Minio.String()
}

func (m *Minio) Save(ctx context.Context, tx *gen.Client, o *model.ObjectStorage) (*model.ObjectStorage, error) {
	panic("implement me")
}

func (m *Minio) UploadToken(ctx context.Context, req *repo.ObjectStorageUploadTokenReq) (*repo.ObjectStorageUploadTokenResponse, error) {
	panic("implement me")
}

func (m *Minio) Status(ctx context.Context, req *repo.ObjectStorageStatusReq) (*repo.ObjectStorageStatusResponse, error) {
	panic("implement me")
}
