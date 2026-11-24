package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/platform/enum/v1"
)

type AssetProvider string

const (
	AssetProviderMinio AssetProvider = "minio" // MinIO 对象存储
	AssetProviderQiniu AssetProvider = "qiniu" // 七牛云对象存储
)

var AssetProviderMap = commonenum.NewMapping[AssetProvider, v1.AssetProvider](map[AssetProvider]commonenum.Entry[AssetProvider, v1.AssetProvider]{
	AssetProviderMinio: {Proto: v1.AssetProvider_ASSET_PROVIDER_MINIO},
	AssetProviderQiniu: {Proto: v1.AssetProvider_ASSET_PROVIDER_QINIU},
})

func (e AssetProvider) String() string {
	return string(e)
}
