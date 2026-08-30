package usecase

import (
	"context"
	"fmt"
)

const maxStateEventDispatchCount = 1024

// StateCommandType 表示角色状态机可处理的命令类型。
type StateCommandType string

const (
	StateCommandTypeActionQueueAdd          StateCommandType = "action_queue.add"           // 队列添加行动
	StateCommandTypeActionQueueMove         StateCommandType = "action_queue.move"          // 队列移动行动
	StateCommandTypeActionQueueRemove       StateCommandType = "action_queue.remove"        // 队列移除行动
	StateCommandTypeActionQueueClear        StateCommandType = "action_queue.clear"         // 队列清空
	StateCommandTypeActionQueueCompleteHead StateCommandType = "action_queue.complete_head" // 队首完成
	StateCommandTypeActionSettlement        StateCommandType = "action.settlement"          // 行动结算
)

// StateCommand 是进入角色状态机的业务命令。
type StateCommand interface {
	StateCommandType() StateCommandType
}

// StateCommandHandler 同步完成核心状态变更，必须在返回前保证状态已经落到 Redis 热状态。
type StateCommandHandler interface {
	Apply(ctx context.Context, command StateCommand) (*StateChangeSet, error)
}

// StateEngine 是 game_idle 的状态机入口，负责分发命令并展开领域事件。
type StateEngine struct {
	commandHandlers map[StateCommandType]StateCommandHandler
	handlers        map[StateEventType][]StateEventHandler
}

func NewStateEngine(
	commandHandlers map[StateCommandType]StateCommandHandler,
	handlers map[StateEventType][]StateEventHandler,
) *StateEngine {
	return &StateEngine{
		commandHandlers: commandHandlers,
		handlers:        handlers,
	}
}

func (e *StateEngine) Apply(ctx context.Context, command StateCommand) (*StateChangeSet, error) {
	commandHandler := e.commandHandlers[command.StateCommandType()]
	if commandHandler == nil {
		return nil, fmt.Errorf("game idle state command handler is required: %s", command.StateCommandType())
	}
	changeSet, err := commandHandler.Apply(ctx, command)
	if err != nil || changeSet == nil {
		return changeSet, err
	}
	events, err := e.dispatch(ctx, changeSet.Events)
	if err != nil {
		return changeSet, err
	}
	changeSet.Events = events
	return changeSet, nil
}

// dispatch 按队列展开事件级联，避免 handler 直接递归造成不可控调用树。
func (e *StateEngine) dispatch(ctx context.Context, initialEvents []*StateEvent) ([]*StateEvent, error) {
	events := append([]*StateEvent(nil), initialEvents...)
	for index := 0; index < len(events); index++ {
		if len(events) > maxStateEventDispatchCount {
			return events, fmt.Errorf("game idle state event dispatch overflow")
		}
		for _, handler := range e.handlers[events[index].Type] {
			nextEvents, err := handler.Handle(ctx, events[index])
			if err != nil {
				return events, err
			}
			events = append(events, nextEvents...)
		}
	}
	return events, nil
}
