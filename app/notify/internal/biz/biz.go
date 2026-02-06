package biz

import (
	"common/pkg/util"
	domainbase "notify/internal/biz/base"
	"notify/internal/biz/domain"
	"notify/internal/biz/handler"

	"github.com/google/wire"
)

// BizProviderSet is biz providers.
var BizProviderSet = wire.NewSet(
	handler.HandlerSet,
	domainbase.NewBaseDomain,
	util.NewTokenCache,
	handler.ProvideHandlers,
	domain.NewEventHandler,
	domain.NewNotificationMetaDomain,
	domain.NewNotificationRecordDomain,
	domain.NewNotificationTemplateDomain,
)
