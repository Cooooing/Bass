package enum

// MessageHandleAction 定义调度消息处理结果。
type MessageHandleAction string

const (
	MessageHandleActionComplete MessageHandleAction = "complete"
	MessageHandleActionRetry    MessageHandleAction = "retry"
	MessageHandleActionDiscard  MessageHandleAction = "discard"
)

func (MessageHandleAction) Values() []string {
	return []string{
		MessageHandleActionComplete.String(),
		MessageHandleActionRetry.String(),
		MessageHandleActionDiscard.String(),
	}
}

func (e MessageHandleAction) String() string {
	return string(e)
}
