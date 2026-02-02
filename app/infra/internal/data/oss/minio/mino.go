package minio

import (
	"common/pkg/constant"
	"context"
	"infra/internal/biz/model"
	"infra/internal/data/ent/gen"
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
	// TODO implement me
	panic("implement me")
}

func (m *Minio) UploadToken(ctx context.Context, key string) (string, error) {
	// TODO implement me
	panic("implement me")
}

func (m *Minio) Status(ctx context.Context, key string, enable bool) error {
	// TODO implement me
	panic("implement me")
}
