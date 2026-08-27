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
	command.NewChatListHandler,
	command.NewChatSendHandler,
	command.NewQueueAddHandler,
	command.NewQueueClearHandler,
	command.NewQueueListHandler,
	command.NewQueueMoveHandler,
	command.NewQueueRemoveHandler,
	event.NewActionCompletedHandler,
	event.NewChatMessageHandler,
	event.NewCloseSessionHandler,
)

func ProvideWebSocketWorkerPool(c *config.Bootstrap) *common.WorkerPool {
	return c.GetWebsocket().GetWorkerPool()
}

func ProvideWebSocketEventHandlers(
	actionCompletedHandler *event.ActionCompletedHandler,
	chatMessageHandler *event.ChatMessageHandler,
	closeSessionHandler *event.CloseSessionHandler,
) usecase.WebSocketEventHandlers {
	return usecase.WebSocketEventHandlers{
		actionCompletedHandler.Type(): actionCompletedHandler,
		chatMessageHandler.Type():     chatMessageHandler,
		closeSessionHandler.Type():    closeSessionHandler,
	}
}

func ProvideWebSocketCommandHandlers(
	backpackGetHandler *command.BackpackGetHandler,
	chatListHandler *command.ChatListHandler,
	chatSendHandler *command.ChatSendHandler,
	queueAddHandler *command.QueueAddHandler,
	queueClearHandler *command.QueueClearHandler,
	queueListHandler *command.QueueListHandler,
	queueMoveHandler *command.QueueMoveHandler,
	queueRemoveHandler *command.QueueRemoveHandler,
) usecase.WebSocketCommandHandlers {
	return usecase.WebSocketCommandHandlers{
		backpackGetHandler.Type(): backpackGetHandler,
		chatListHandler.Type():    chatListHandler,
		chatSendHandler.Type():    chatSendHandler,
		queueAddHandler.Type():    queueAddHandler,
		queueClearHandler.Type():  queueClearHandler,
		queueListHandler.Type():   queueListHandler,
		queueMoveHandler.Type():   queueMoveHandler,
		queueRemoveHandler.Type(): queueRemoveHandler,
	}
}
