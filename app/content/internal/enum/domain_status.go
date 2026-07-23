package enum

import (
	"common/pkg/enum"
	v1 "common/proto/gen/content/v1/enum"
)

// DomainStatus 表示领域状态。
type DomainStatus string

const (
	// DomainStatusEnabled 表示领域启用。
	DomainStatusEnabled DomainStatus = "enabled"
	// DomainStatusDisabled 表示领域禁用。
	DomainStatusDisabled DomainStatus = "disabled"
)

// DomainStatusMap 维护领域状态内部枚举与 proto 枚举的映射。
var DomainStatusMap = enum.NewMapping[DomainStatus, v1.DomainStatus](map[DomainStatus]enum.Entry[DomainStatus, v1.DomainStatus]{
	DomainStatusEnabled:  {Proto: v1.DomainStatus_DOMAIN_STATUS_ENABLED},
	DomainStatusDisabled: {Proto: v1.DomainStatus_DOMAIN_STATUS_DISABLED},
})
