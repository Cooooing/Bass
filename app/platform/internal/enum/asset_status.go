package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/platform/enum/v1"
)

type AssetStatus string

const (
	AssetStatusAvailable AssetStatus = "available" // 资源可正常访问
	AssetStatusBlocked   AssetStatus = "blocked"   // 资源已禁止访问
)

var AssetStatusMap = commonenum.NewMapping[AssetStatus, v1.AssetStatus](map[AssetStatus]commonenum.Entry[AssetStatus, v1.AssetStatus]{
	AssetStatusAvailable: {Proto: v1.AssetStatus_ASSET_STATUS_AVAILABLE},
	AssetStatusBlocked:   {Proto: v1.AssetStatus_ASSET_STATUS_BLOCKED},
})

func (e AssetStatus) String() string {
	return string(e)
}
