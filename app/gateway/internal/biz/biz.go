package biz

import (
	"common/pkg/util/jwt"
	domainbase "gateway/internal/biz/base"
	"gateway/internal/biz/domain"

	"github.com/google/wire"
)

// BizProviderSet is biz providers.
var BizProviderSet = wire.NewSet(
	domainbase.NewBaseDomain,

	jwt.NewTokenCache,
	domain.NewIpDomain,
)
