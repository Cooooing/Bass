package oss

import (
	"notify/internal/biz/repo"
	"notify/internal/conf"
	"notify/internal/data/oss/minio"
	"notify/internal/data/oss/qiniu"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewFactory,
	ProvideObjectStorageProvider,
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

func ProvideObjectStorageProvider(conf *conf.Bootstrap, minio *minio.Minio, qiniu *qiniu.Qiniu) repo.ObjectStorageProvider {
	return NewFactory(minio, qiniu).Get(conf.Server.Oss.Provider)
}
