package enum

import (
	v1 "common/api/gen/content/v1"
	"common/pkg/enum"
)

type DomainStatus string

const (
	DomainStatusNormal   DomainStatus = "normal"
	DomainStatusDisabled DomainStatus = "disabled"
)

var DomainStatusMap = enum.NewMapping[DomainStatus, v1.DomainStatus](map[DomainStatus]enum.Entry[DomainStatus, v1.DomainStatus]{
	DomainStatusNormal:   {Proto: v1.DomainStatus_DOMAIN_STATUS_NORMAL},
	DomainStatusDisabled: {Proto: v1.DomainStatus_DOMAIN_STATUS_DISABLED},
})
