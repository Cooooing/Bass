package biz

import (
	"common/pkg/util"
	"notify/internal/biz/base"
	"notify/internal/biz/handler"

	"github.com/google/wire"
)

// BizProviderSet is biz providers.
var BizProviderSet = wire.NewSet(
	handler.HandlerSet,
	base.NewBaseDomain,
	util.NewTokenRepo,
	handler.ProvideHandlers,
	NewEventHandler,
	NewNotificationMetaDomain,
	NewNotificationRecordDomain,
	NewNotificationTemplateDomain,
	base.NewEmailDomain,
)
