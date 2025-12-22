package oss

import (
	"user/internal/biz/repo"
	"user/internal/data/oss/minio"
	"user/internal/data/oss/qiniu"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewFactory,
	minio.NewMinio,
	qiniu.NewQiniu,
)

type Factory struct {
	providers map[string]repo.ObjectStorageProvider
}

func NewFactory(
	minio *minio.Minio,
	qiniu *qiniu.Qiniu,
) *Factory {
	return &Factory{
		providers: map[string]repo.ObjectStorageProvider{
			minio.Name(): minio,
			qiniu.Name(): qiniu,
		},
	}
}

func (f *Factory) Get(name string) repo.ObjectStorageProvider {
	return f.providers[name]
}
