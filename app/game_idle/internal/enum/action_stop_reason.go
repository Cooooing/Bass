package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/game_idle/v1/enum"
)

// ActionStopReason 表示行动结算后是否继续执行队首行动。
type ActionStopReason string

const (
	ActionStopReasonNone              ActionStopReason = "none"               // 继续执行
	ActionStopReasonInsufficientItems ActionStopReason = "insufficient_items" // 物品不足
)

var ActionStopReasonMap = commonenum.NewMapping[ActionStopReason, v1.ActionStopReason](map[ActionStopReason]commonenum.Entry[ActionStopReason, v1.ActionStopReason]{
	ActionStopReasonNone: {Proto: v1.ActionStopReason_ACTION_STOP_REASON_NONE}, ActionStopReasonInsufficientItems: {Proto: v1.ActionStopReason_ACTION_STOP_REASON_INSUFFICIENT_ITEMS},
})
