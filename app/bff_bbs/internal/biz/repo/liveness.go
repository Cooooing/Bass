package repo

import (
	"context"
	"time"

	userenum "common/proto/gen/user/v1/enum"
)

type Liveness struct {
	Score        int32
	MaxScore     int32
	CheckinScore int32
	Settled      bool
	RewardStatus userenum.LivenessRewardStatus
}
type LivenessClient interface {
	Get(context.Context, int64, time.Time) (*Liveness, error)
	Claim(context.Context, int64, time.Time) (*Liveness, error)
}
