package repo

import (
	"context"
	"time"

	userenum "common/proto/gen/user/enum/v1"
)

type Liveness struct {
	Score        int32
	MaxScore     int32
	CheckinScore int32
	Settled      bool
	RewardStatus userenum.LivenessRewardStatus
}
type LivenessClient interface {
	Get(ctx context.Context, userID int64, date time.Time) (*Liveness, error)
	Claim(ctx context.Context, userID int64, date time.Time) (*Liveness, error)
}
