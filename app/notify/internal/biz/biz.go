package biz

import (
	"common/pkg/client/rpc"
	"common/pkg/util/jwt"
	domainbase "notify/internal/biz/base"
	"notify/internal/biz/domain"

	"github.com/google/wire"
)

// BizProviderSet is biz providers.
var BizProviderSet = wire.NewSet(
	domainbase.NewBaseDomain,
	jwt.NewTokenCache,

	domain.NewEventHandler,
	domain.NewNotificationMetaDomain,
	domain.NewNotificationRecordDomain,
	domain.NewNotificationTemplateDomain,
	domain.NewEmailDomain,
	domain.NewTencentSmsDomain,
	domain.NewObjectStorageDomain,

	rpc.ProvideUserClient,
	rpc.ProvideContentClient,
)
