package data

import (
	"common/pkg/client/rpc"
	userv1 "common/proto/gen/user/v1"
	"context"
	"fmt"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"bff_bbs/internal/biz/repo"
)

var _ repo.LivenessClient = (*LivenessClient)(nil)

type LivenessClient struct{ user *rpc.UserClient }

func NewLivenessClient(user *rpc.UserClient) repo.LivenessClient { return &LivenessClient{user: user} }
func (c *LivenessClient) Get(ctx context.Context, userID int64, date time.Time) (*repo.Liveness, error) {
	res, err := c.user.Liveness.GetLiveness(ctx, &userv1.GetUserLiveness_Req{UserId: userID, Date: timestamppb.New(date)})
	if err != nil {
		return nil, err
	}
	return &repo.Liveness{Score: res.GetScore(), MaxScore: res.GetMaxScore(), CheckinScore: res.GetCheckinScore(), Settled: res.GetSettled(), RewardStatus: res.GetRewardStatus()}, nil
}
func (c *LivenessClient) Claim(ctx context.Context, userID int64, date time.Time) (*repo.Liveness, error) {
	res, err := c.user.Liveness.ClaimLivenessReward(ctx, &userv1.ClaimUserLivenessReward_Req{UserId: userID, Date: timestamppb.New(date)})
	if err != nil {
		return nil, err
	}
	points := res.GetRewardPoints()
	const minInt32, maxInt32 = int64(-1 << 31), int64(1<<31 - 1)
	if points < minInt32 || points > maxInt32 {
		return nil, fmt.Errorf("liveness reward points %d exceed int32 range", points)
	}
	return &repo.Liveness{Score: int32(points), Settled: true, RewardStatus: res.GetRewardStatus()}, nil
}
