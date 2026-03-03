package biz

import (
	"common/pkg/util"
	"common/pkg/util/jwt"
	doaminbase "user/internal/biz/base"
	"user/internal/biz/doamin"

	"github.com/google/wire"
)

// BizProviderSet is biz providers.
var BizProviderSet = wire.NewSet(
	doaminbase.NewBaseDomain,

	jwt.NewTokenCache,
	doamin.NewTokenService,
	util.NewEventPool,

	doamin.NewAuthenticationDomain,
	doamin.NewUserDomain,
	doamin.NewUserRelationDomain,
	doamin.NewTwoFactorAuthenticationDomain,
)
