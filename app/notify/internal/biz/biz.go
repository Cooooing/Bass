package biz

import (
	"common/pkg/client/rpc"
	"common/pkg/util/jwt"
	domainbase "notify/internal/biz/base"
	"notify/internal/biz/domain"
	"notify/internal/biz/handler"

	"github.com/google/wire"
)

// BizProviderSet is biz providers.
var BizProviderSet = wire.NewSet(
	handler.HandlerSet,
	domainbase.NewBaseDomain,
	jwt.NewTokenCache,
	handler.ProvideHandlers,
	domain.NewEventHandler,
	domain.NewNotificationMetaDomain,
	domain.NewNotificationRecordDomain,
	domain.NewNotificationTemplateDomain,

	rpc.ProvideUserClient,
	rpc.ProvideInfraClient,
)
