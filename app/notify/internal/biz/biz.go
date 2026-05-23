package biz

import (
	"common/pkg/client/rpc"
	"common/pkg/util/jwt"
	"notify/internal/biz/usecase"
	"notify/internal/biz/usecase/consumer"
	"notify/internal/biz/usecase/handler"
	"notify/internal/biz/usecase/sender"

	"github.com/google/wire"
)

// BizProviderSet 是 biz 层依赖集合。
var BizProviderSet = wire.NewSet(
	jwt.NewTokenCache,

	consumer.NewConsumer,
	handler.NewDispatcher,
	sender.ProvideRegistry,
	usecase.NewNotifyUsecase,
	usecase.NewRPCUserResolver,
	wire.Bind(new(usecase.UserResolver), new(*usecase.RPCUserResolver)),

	sender.NewSmtpSender,
	sender.NewTencentSmsSender,

	handler.NewFollowHandler,
	handler.NewArticlePublishHandler,
	handler.NewArticleActionHandler,
	handler.NewCommentHandler,
	handler.NewCommentActionHandler,
	handler.NewDefaultHandler,

	usecase.NewNotificationMetaUsecase,
	usecase.NewNotificationRecordUsecase,
	usecase.NewNotificationTemplateUsecase,
	usecase.NewObjectStorageUsecase,

	rpc.ProvideUserClient,
	rpc.ProvideContentClient,
)
