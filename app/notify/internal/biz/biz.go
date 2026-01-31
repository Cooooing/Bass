package biz

import (
	"common/pkg/util"
	"notify/internal/biz/base"
	"notify/internal/biz/handler"
	"notify/internal/biz/infra"

	"github.com/google/wire"
)

// BizProviderSet is biz providers.
var BizProviderSet = wire.NewSet(
	handler.HandlerSet,
	base.NewBaseDomain,
	util.NewTokenCache,
	handler.ProvideHandlers,
	NewEventHandler,
	NewNotificationMetaDomain,
	NewNotificationRecordDomain,
	NewNotificationTemplateDomain,
	infra.NewTencentSmsDomain,
	infra.NewEmailDomain,
)
