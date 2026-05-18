package biz

import (
	"common/pkg/client/rpc"
	"common/pkg/util/jwt"
	"notify/internal/biz/domain"
	"notify/internal/biz/domain/consumer"
	"notify/internal/biz/domain/handler"
	"notify/internal/biz/domain/sender"

	"github.com/google/wire"
)

// BizProviderSet is biz providers.
var BizProviderSet = wire.NewSet(
	jwt.NewTokenCache,

	consumer.NewConsumer,
	handler.NewDispatcher,
	sender.ProvideRegistry,
	domain.NewNotifyService,
	domain.NewRPCUserResolver,
	wire.Bind(new(domain.UserResolver), new(*domain.RPCUserResolver)),

	sender.NewSmtpSender,
	sender.NewTencentSmsSender,

	handler.NewFollowHandler,
	handler.NewArticlePublishHandler,
	handler.NewArticleActionHandler,
	handler.NewCommentHandler,
	handler.NewCommentActionHandler,
	handler.NewDefaultHandler,

	domain.NewNotificationMetaDomain,
	domain.NewNotificationRecordDomain,
	domain.NewNotificationTemplateDomain,
	domain.NewObjectStorageDomain,

	rpc.ProvideUserClient,
	rpc.ProvideContentClient,
)
