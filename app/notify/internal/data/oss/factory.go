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
	ProvideObjectStorageClient,
	minio.NewMinio,
	qiniu.NewQiniu,
)

type Factory struct {
	clients map[string]repo.ObjectStorageClient
}

func NewFactory(
	minio *minio.Minio,
	qiniu *qiniu.Qiniu,
) *Factory {
	return &Factory{
		clients: map[string]repo.ObjectStorageClient{
			minio.Name(): minio,
			qiniu.Name(): qiniu,
		},
	}
}

func (f *Factory) Get(name string) repo.ObjectStorageClient {
	return f.clients[name]
}

func ProvideObjectStorageClient(conf *conf.Bootstrap, minio *minio.Minio, qiniu *qiniu.Qiniu) repo.ObjectStorageClient {
	return NewFactory(minio, qiniu).Get(conf.Server.Oss.Provider)
}
