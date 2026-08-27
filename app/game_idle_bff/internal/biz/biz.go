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
	usecase.NewBackpackUsecase,
	usecase.NewActionQueueUsecase,
	usecase.NewChatUsecase,
	usecase.NewWebSocketUsecase,
	command.NewBackpackGetHandler,
	command.NewChatMessageListHandler,
	command.NewChatMessageSendHandler,
	command.NewActionAddHandler,
	command.NewActionClearHandler,
	command.NewActionListHandler,
	command.NewActionMoveHandler,
	command.NewActionRemoveHandler,
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
	backpackGetHandler *command.BackpackGetHandler,
	chatMessageListHandler *command.ChatMessageListHandler,
	chatMessageSendHandler *command.ChatMessageSendHandler,
	actionAddHandler *command.ActionAddHandler,
	actionClearHandler *command.ActionClearHandler,
	actionListHandler *command.ActionListHandler,
	actionMoveHandler *command.ActionMoveHandler,
	actionRemoveHandler *command.ActionRemoveHandler,
) usecase.WebSocketCommandHandlers {
	return usecase.WebSocketCommandHandlers{
		backpackGetHandler.Type():     backpackGetHandler,
		chatMessageListHandler.Type(): chatMessageListHandler,
		chatMessageSendHandler.Type(): chatMessageSendHandler,
		actionAddHandler.Type():       actionAddHandler,
		actionClearHandler.Type():     actionClearHandler,
		actionListHandler.Type():      actionListHandler,
		actionMoveHandler.Type():      actionMoveHandler,
		actionRemoveHandler.Type():    actionRemoveHandler,
	}
}
