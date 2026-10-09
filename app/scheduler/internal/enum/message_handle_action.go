package enum

import (
	commonenum "common/pkg/enum"
	schedulerv1 "common/proto/gen/scheduler/v1/enum"
)

// MessageHandleAction 定义调度消息处理结果。
type MessageHandleAction string

const (
	MessageHandleActionComplete MessageHandleAction = "complete" // complete
	MessageHandleActionRetry    MessageHandleAction = "retry"    // retry
	MessageHandleActionDiscard  MessageHandleAction = "discard"  // discard
)

var MessageHandleActionMap = commonenum.NewMapping[MessageHandleAction, schedulerv1.SchedulerMessageHandleAction](map[MessageHandleAction]commonenum.Entry[MessageHandleAction, schedulerv1.SchedulerMessageHandleAction]{
	MessageHandleActionComplete: {Proto: schedulerv1.SchedulerMessageHandleAction_SCHEDULER_MESSAGE_HANDLE_ACTION_COMPLETE}, MessageHandleActionRetry: {Proto: schedulerv1.SchedulerMessageHandleAction_SCHEDULER_MESSAGE_HANDLE_ACTION_RETRY}, MessageHandleActionDiscard: {Proto: schedulerv1.SchedulerMessageHandleAction_SCHEDULER_MESSAGE_HANDLE_ACTION_DISCARD},
})

func (e MessageHandleAction) String() string {
	return string(e)
}
