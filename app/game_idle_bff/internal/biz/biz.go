package biz

import (
	commonclient "common/pkg/client"
	"common/proto/gen/common"
	"game_idle_bff/internal/biz/usecase"
	"game_idle_bff/internal/biz/usecase/command"
	"game_idle_bff/internal/biz/usecase/event"
	"game_idle_bff/internal/config"

	"github.com/google/wire"
)

var BizProviderSet = wire.NewSet(
	ProvideWebSocketEventHandlers,
	ProvideWebSocketCommandHandlers,
	commonclient.NewWorkerPool,
	ProvideWebSocketWorkerPool,
	usecase.NewAuthUsecase,
	usecase.NewCharacterUsecase,
	usecase.NewCharacterAbilityUsecase,
	usecase.NewBackpackUsecase,
	usecase.NewActionQueueUsecase,
	usecase.NewChatUsecase,
	usecase.NewConfigUsecase,
	usecase.NewWebSocketUsecase,
	command.NewChatMessageSendHandler,
	command.NewActionAddHandler,
	command.NewActionClearHandler,
	command.NewActionMoveHandler,
	command.NewActionRemoveHandler,
	command.NewInitGetHandler,
	command.NewConfigGetHandler,
	command.NewConfigVersionHandler,
	command.NewActionDetailGetHandler,
	event.NewActionCompletedHandler,
	event.NewActionQueueUpdatedHandler,
	event.NewAbilityLeveledUpHandler,
	event.NewChatMessageHandler,
	event.NewCloseSessionHandler,
)

func ProvideWebSocketWorkerPool(c *config.Bootstrap) *common.WorkerPool {
	return c.GetWebsocket().GetWorkerPool()
}

func ProvideWebSocketEventHandlers(
	actionCompletedHandler *event.ActionCompletedHandler,
	actionQueueUpdatedHandler *event.ActionQueueUpdatedHandler,
	abilityLeveledUpHandler *event.AbilityLeveledUpHandler,
	chatMessageHandler *event.ChatMessageHandler,
	closeSessionHandler *event.CloseSessionHandler,
) usecase.WebSocketEventHandlers {
	return usecase.WebSocketEventHandlers{
		actionCompletedHandler.Type():    actionCompletedHandler,
		actionQueueUpdatedHandler.Type(): actionQueueUpdatedHandler,
		abilityLeveledUpHandler.Type():   abilityLeveledUpHandler,
		chatMessageHandler.Type():        chatMessageHandler,
		closeSessionHandler.Type():       closeSessionHandler,
	}
}

func ProvideWebSocketCommandHandlers(
	chatMessageSendHandler *command.ChatMessageSendHandler,
	actionAddHandler *command.ActionAddHandler,
	actionClearHandler *command.ActionClearHandler,
	actionMoveHandler *command.ActionMoveHandler,
	actionRemoveHandler *command.ActionRemoveHandler,
	initGetHandler *command.InitGetHandler,
	configGetHandler *command.ConfigGetHandler,
	configVersionHandler *command.ConfigVersionHandler,
	actionDetailGetHandler *command.ActionDetailGetHandler,
) usecase.WebSocketCommandHandlers {
	return usecase.WebSocketCommandHandlers{
		chatMessageSendHandler.Type(): chatMessageSendHandler,
		actionAddHandler.Type():       actionAddHandler,
		actionClearHandler.Type():     actionClearHandler,
		actionMoveHandler.Type():      actionMoveHandler,
		actionRemoveHandler.Type():    actionRemoveHandler,
		initGetHandler.Type():         initGetHandler,
		configGetHandler.Type():       configGetHandler,
		configVersionHandler.Type():   configVersionHandler,
		actionDetailGetHandler.Type(): actionDetailGetHandler,
	}
}
