package biz

import (
	"common/pkg/util"
	doaminbase "user/internal/biz/base"
	"user/internal/biz/doamin"

	"github.com/google/wire"
)

// BizProviderSet is biz providers.
var BizProviderSet = wire.NewSet(
	doaminbase.NewBaseDomain,

	util.NewTokenCache,
	doamin.NewTokenService,
	util.NewEventPool,

	doamin.NewAuthenticationDomain,
	doamin.NewUserDomain,
	doamin.NewUserRelationDomain,
	doamin.NewTwoFactorAuthenticationDomain,
)
