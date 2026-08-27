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
	usecase.NewWebSocketUsecase,
	command.NewChatMessageSendHandler,
	command.NewActionAddHandler,
	command.NewActionClearHandler,
	command.NewActionMoveHandler,
	command.NewActionRemoveHandler,
	command.NewInitGetHandler,
	event.NewActionCompletedHandler,
	event.NewAbilityLeveledUpHandler,
	event.NewChatMessageHandler,
	event.NewCloseSessionHandler,
)

func ProvideWebSocketWorkerPool(c *config.Bootstrap) *common.WorkerPool {
	return c.GetWebsocket().GetWorkerPool()
}

func ProvideWebSocketEventHandlers(
	actionCompletedHandler *event.ActionCompletedHandler,
	abilityLeveledUpHandler *event.AbilityLeveledUpHandler,
	chatMessageHandler *event.ChatMessageHandler,
	closeSessionHandler *event.CloseSessionHandler,
) usecase.WebSocketEventHandlers {
	return usecase.WebSocketEventHandlers{
		actionCompletedHandler.Type():  actionCompletedHandler,
		abilityLeveledUpHandler.Type(): abilityLeveledUpHandler,
		chatMessageHandler.Type():      chatMessageHandler,
		closeSessionHandler.Type():     closeSessionHandler,
	}
}

func ProvideWebSocketCommandHandlers(
	chatMessageSendHandler *command.ChatMessageSendHandler,
	actionAddHandler *command.ActionAddHandler,
	actionClearHandler *command.ActionClearHandler,
	actionMoveHandler *command.ActionMoveHandler,
	actionRemoveHandler *command.ActionRemoveHandler,
	initGetHandler *command.InitGetHandler,
) usecase.WebSocketCommandHandlers {
	return usecase.WebSocketCommandHandlers{
		chatMessageSendHandler.Type(): chatMessageSendHandler,
		actionAddHandler.Type():       actionAddHandler,
		actionClearHandler.Type():     actionClearHandler,
		actionMoveHandler.Type():      actionMoveHandler,
		actionRemoveHandler.Type():    actionRemoveHandler,
		initGetHandler.Type():         initGetHandler,
	}
}
