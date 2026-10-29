package service

import (
	"context"

	"bff_bbs/internal/biz/usecase"
	"common/pkg/apperror"
	"common/pkg/constant"
	commonmodel "common/pkg/model"
	"common/pkg/util"
	bbsuserv1 "common/proto/gen/bff_bbs/v1/user"
	cerrors "common/proto/gen/common/errors"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

// LivenessService 适配社区活跃度 HTTP 协议。
type LivenessService struct {
	bbsuserv1.UnimplementedLivenessServiceServer
	usecase *usecase.LivenessUsecase
}

func NewLivenessService(usecase *usecase.LivenessUsecase) *LivenessService {
	return &LivenessService{usecase: usecase}
}
func (*LivenessService) RegisterGrpc(*grpc.Server) {}
func (s *LivenessService) RegisterHTTP(server *http.Server) {
	bbsuserv1.RegisterLivenessServiceHTTPServer(server, s)
}

func livenessUser(ctx context.Context) (*commonmodel.User, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	return user, nil
}

func (s *LivenessService) GetLiveness(ctx context.Context, _ *bbsuserv1.GetLiveness_Req) (*bbsuserv1.GetLiveness_Resp, error) {
	user, err := livenessUser(ctx)
	if err != nil {
		return nil, err
	}
	row, err := s.usecase.GetLiveness(ctx, user.ID)
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.GetLiveness_Resp{Score: row.Score, MaxScore: row.MaxScore, CheckinScore: row.CheckinScore}, nil
}

func (s *LivenessService) ClaimYesterdayReward(ctx context.Context, _ *bbsuserv1.ClaimYesterdayLivenessReward_Req) (*bbsuserv1.ClaimYesterdayLivenessReward_Resp, error) {
	user, err := livenessUser(ctx)
	if err != nil {
		return nil, err
	}
	row, err := s.usecase.ClaimYesterdayReward(ctx, user.ID)
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.ClaimYesterdayLivenessReward_Resp{RewardPoints: int64(row.Score), Claimed: true}, nil
}
