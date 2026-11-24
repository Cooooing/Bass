package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/user/enum/v1"
)

// LivenessRewardStatus 表示每日活跃奖励的领取状态。
type LivenessRewardStatus string

const (
	LivenessRewardStatusUnclaimed LivenessRewardStatus = "unclaimed" // 尚未领取奖励
	LivenessRewardStatusPending   LivenessRewardStatus = "pending"   // 等待发放奖励
	LivenessRewardStatusCompleted LivenessRewardStatus = "completed" // 奖励已发放
	LivenessRewardStatusCanceled  LivenessRewardStatus = "canceled"  // 奖励已取消
)

var LivenessRewardStatusMap = commonenum.NewMapping[LivenessRewardStatus, v1.LivenessRewardStatus](map[LivenessRewardStatus]commonenum.Entry[LivenessRewardStatus, v1.LivenessRewardStatus]{
	LivenessRewardStatusUnclaimed: {Proto: v1.LivenessRewardStatus_LIVENESS_REWARD_STATUS_UNCLAIMED},
	LivenessRewardStatusPending:   {Proto: v1.LivenessRewardStatus_LIVENESS_REWARD_STATUS_PENDING},
	LivenessRewardStatusCompleted: {Proto: v1.LivenessRewardStatus_LIVENESS_REWARD_STATUS_COMPLETED},
	LivenessRewardStatusCanceled:  {Proto: v1.LivenessRewardStatus_LIVENESS_REWARD_STATUS_CANCELED},
})

func (e LivenessRewardStatus) String() string {
	return string(e)
}
