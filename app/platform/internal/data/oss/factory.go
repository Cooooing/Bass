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
	ProvideStorageClient,
)

func ProvideStorageClient(conf *config.Bootstrap) (repo.StorageClient, error) {
	provider := conf.GetPlatform().GetOss().GetProvider()
	switch provider {
	case enum.AssetProviderMinio.String():
		return minio.NewMinio(conf)
	case enum.AssetProviderQiniu.String():
		return qiniu.NewQiniu(conf), nil
	default:
		return nil, fmt.Errorf("unsupported object storage provider: %s", provider)
	}
}
