package biz

import (
	"common/pkg/util"
	"notify/internal/biz/base"
	"notify/internal/biz/domain"
	"notify/internal/biz/handler"

	"github.com/google/wire"
)

// BizProviderSet is biz providers.
var BizProviderSet = wire.NewSet(
	handler.HandlerSet,
	base.NewBaseDomain,
	util.NewTokenCache,
	handler.ProvideHandlers,
	domain.NewEventHandler,
	domain.NewNotificationMetaDomain,
	domain.NewNotificationRecordDomain,
	domain.NewNotificationTemplateDomain,
	domain.NewTencentSmsDomain,
	domain.NewEmailDomain,
)
