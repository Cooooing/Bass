package usecase

import (
	"context"
	"time"

	"bff_bbs/internal/biz/repo"
)

// LivenessUsecase 提供当前活跃度读取和昨日活跃奖励领取。
type LivenessUsecase struct{ client repo.LivenessClient }

func NewLivenessUsecase(client repo.LivenessClient) *LivenessUsecase {
	return &LivenessUsecase{client: client}
}

func (u *LivenessUsecase) GetLiveness(ctx context.Context, userID int64) (*repo.Liveness, error) {
	return u.client.Get(ctx, userID, time.Now())
}

func (u *LivenessUsecase) ClaimYesterdayReward(ctx context.Context, userID int64) (*repo.Liveness, error) {
	return u.client.Claim(ctx, userID, time.Now().AddDate(0, 0, -1))
}
