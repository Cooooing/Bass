package enum

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
