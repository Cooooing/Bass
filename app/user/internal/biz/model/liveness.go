package model

import (
	"time"
	"user/internal/enum"
)

// LivenessDetail 是冻结或实时聚合的各来源活跃次数。
type LivenessDetail struct {
	ArticleCount           int32
	CommentCount           int32
	LikeCount              int32
	ThankCount             int32
	RewardCount            int32
	ViewCount              int32
	GlobalChatMessageCount int32
}

// LivenessRecord 是某业务日的不可变活跃度结算及手动奖励事实。
type LivenessRecord struct {
	ID           int64
	UserID       int64
	Date         *time.Time
	Score        int32
	RewardStatus enum.LivenessRewardStatus
	ClaimedAt    *time.Time
}
