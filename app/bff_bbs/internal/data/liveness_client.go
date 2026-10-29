package data

import (
	"bff_bbs/internal/biz/repo"
	"common/pkg/client/rpc"
	userv1 "common/proto/gen/user/v1"
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
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
	return &repo.Liveness{Score: int32(res.GetRewardPoints()), Settled: true, RewardStatus: res.GetRewardStatus()}, nil
}
