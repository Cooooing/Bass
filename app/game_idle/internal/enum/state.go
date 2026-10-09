package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/game_idle/v1/enum"
)

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

var StateCommandTypeMap = commonenum.NewMapping[StateCommandType, v1.StateCommandType](map[StateCommandType]commonenum.Entry[StateCommandType, v1.StateCommandType]{
	StateCommandTypeActionQueueAdd: {Proto: v1.StateCommandType_STATE_COMMAND_TYPE_ACTION_QUEUE_ADD}, StateCommandTypeActionQueueMove: {Proto: v1.StateCommandType_STATE_COMMAND_TYPE_ACTION_QUEUE_MOVE}, StateCommandTypeActionQueueRemove: {Proto: v1.StateCommandType_STATE_COMMAND_TYPE_ACTION_QUEUE_REMOVE}, StateCommandTypeActionQueueClear: {Proto: v1.StateCommandType_STATE_COMMAND_TYPE_ACTION_QUEUE_CLEAR}, StateCommandTypeActionQueueCompleteHead: {Proto: v1.StateCommandType_STATE_COMMAND_TYPE_ACTION_QUEUE_COMPLETE_HEAD}, StateCommandTypeActionSettlement: {Proto: v1.StateCommandType_STATE_COMMAND_TYPE_ACTION_SETTLEMENT},
})
