package minio

import (
	"common/pkg/constant"
	"context"
	"user/internal/biz/model"
	"user/internal/data/ent/gen"
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
	//TODO implement me
	panic("implement me")
}

func (m *Minio) UploadToken(key string) string {
	//TODO implement me
	panic("implement me")
}

func (m *Minio) Status(ctx context.Context, key string, enable bool) error {
	//TODO implement me
	panic("implement me")
}

func (m *Minio) VerifyCallback(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}
