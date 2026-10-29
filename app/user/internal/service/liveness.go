package service

import (
	v1 "common/proto/gen/user/v1"
	"context"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"user/internal/biz/usecase"
	"user/internal/enum"
)

type LivenessService struct {
	v1.UnimplementedLivenessServiceServer
	livenessUsecase *usecase.LivenessUsecase
}

func NewLivenessService(livenessUsecase *usecase.LivenessUsecase) *LivenessService {
	return &LivenessService{livenessUsecase: livenessUsecase}
}

func (s *LivenessService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterLivenessServiceServer(gs, s)
}

func (s *LivenessService) RegisterHTTP(hs *http.Server) {
}

func (s *LivenessService) GetLiveness(ctx context.Context, req *v1.GetUserLiveness_Req) (*v1.GetUserLiveness_Resp, error) {
	score, detail, record, err := s.livenessUsecase.Get(ctx, req.GetUserId(), req.GetDate().AsTime())
	if err != nil {
		return nil, err
	}
	checkinScore, err := s.livenessUsecase.CheckinScore()
	if err != nil {
		return nil, err
	}
	maximumScore, err := s.livenessUsecase.MaximumScore()
	if err != nil {
		return nil, err
	}
	response := &v1.GetUserLiveness_Resp{Date: req.GetDate(), Score: score, MaxScore: maximumScore, CheckinScore: checkinScore, Detail: &v1.UserLivenessDetail{}}
	if detail != nil {
		response.Detail = &v1.UserLivenessDetail{ArticleCount: detail.ArticleCount, CommentCount: detail.CommentCount, LikeCount: detail.LikeCount, ThankCount: detail.ThankCount, RewardCount: detail.RewardCount, ViewCount: detail.ViewCount, GlobalChatMessageCount: detail.GlobalChatMessageCount}
	}
	if record != nil {
		response.Settled = true
		response.RewardPoints = int64(record.Score)
		response.RewardStatus = enum.LivenessRewardStatusMap.MustToProto(record.RewardStatus)
	}
	return response, nil
}

func (s *LivenessService) ClaimLivenessReward(ctx context.Context, req *v1.ClaimUserLivenessReward_Req) (*v1.ClaimUserLivenessReward_Resp, error) {
	row, err := s.livenessUsecase.Claim(ctx, req.GetUserId(), req.GetDate().AsTime())
	if err != nil {
		return nil, err
	}
	return &v1.ClaimUserLivenessReward_Resp{RewardPoints: int64(row.Score), RewardStatus: enum.LivenessRewardStatusMap.MustToProto(row.RewardStatus)}, nil
}

func (s *LivenessService) SettleLiveness(ctx context.Context, req *v1.SettleUserLiveness_Req) (*v1.SettleUserLiveness_Resp, error) {
	settled, checked, err := s.livenessUsecase.Settle(ctx, req.GetDate().AsTime(), req.GetPageSize())
	if err != nil {
		return nil, err
	}
	return &v1.SettleUserLiveness_Resp{SettledUsers: settled, CheckedInUsers: checked}, nil
}
