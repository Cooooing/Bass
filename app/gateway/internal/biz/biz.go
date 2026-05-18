package biz

import (
	"common/pkg/util/jwt"
	"gateway/internal/biz/domain"

	"github.com/google/wire"
)

// BizProviderSet is biz providers.
var BizProviderSet = wire.NewSet(
	jwt.NewTokenCache,
	domain.NewIpDomain,
)
