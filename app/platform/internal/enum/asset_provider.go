package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/platform/v1/enum"
)

type AssetProvider string

const (
	AssetProviderMinio AssetProvider = "minio"
	AssetProviderQiniu AssetProvider = "qiniu"
)

var AssetProviderMap = commonenum.NewMapping[AssetProvider, v1.AssetProvider](map[AssetProvider]commonenum.Entry[AssetProvider, v1.AssetProvider]{
	AssetProviderMinio: {Proto: v1.AssetProvider_ASSET_PROVIDER_MINIO},
	AssetProviderQiniu: {Proto: v1.AssetProvider_ASSET_PROVIDER_QINIU},
})

func (e AssetProvider) String() string {
	return string(e)
}
