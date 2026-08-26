package enum

// ActionStopReason 表示行动结算后是否继续执行队首行动。
type ActionStopReason string

const (
	ActionStopReasonNone              ActionStopReason = "none"               // 继续执行
	ActionStopReasonInsufficientItems ActionStopReason = "insufficient_items" // 物品不足
)
