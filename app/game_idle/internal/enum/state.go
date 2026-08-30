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

// StateEventType 表示状态机产生的领域事件类型。
type StateEventType string

const (
	StateEventTypeActionQueueChanged StateEventType = "action_queue.changed" // 行动队列变化
	StateEventTypeItemsChanged       StateEventType = "items.changed"        // 物品变化
	StateEventTypeAbilityExpGained   StateEventType = "ability.exp_gained"   // 能力经验变化
	StateEventTypeAbilityLeveledUp   StateEventType = "ability.leveled_up"   // 能力升级
)
