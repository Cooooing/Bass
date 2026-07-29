package oss

import (
	"fmt"
	"platform/internal/biz/repo"
	"platform/internal/config"
	"platform/internal/data/oss/minio"
	"platform/internal/data/oss/qiniu"
	"platform/internal/enum"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	ProvideObjectStorageClient,
)

func ProvideObjectStorageClient(conf *config.Bootstrap) (repo.ObjectStorageClient, error) {
	provider := conf.GetPlatform().GetOss().GetProvider()
	switch provider {
	case enum.ObjectStorageProviderMinio.String():
		return minio.NewMinio(conf)
	case enum.ObjectStorageProviderQiniu.String():
		return qiniu.NewQiniu(conf), nil
	default:
		return nil, fmt.Errorf("unsupported object storage provider: %s", provider)
	}
}
