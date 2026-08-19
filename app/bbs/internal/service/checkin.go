package service

import (
	"bbs/internal/biz/usecase"
	"common/pkg/apperror"
	"common/pkg/constant"
	commonmodel "common/pkg/model"
	"common/pkg/util"
	bbsuserv1 "common/proto/gen/bbs/v1/user"
	cerrors "common/proto/gen/common/errors"
	"context"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

type CheckinService struct {
	bbsuserv1.UnimplementedCheckinServiceServer
	checkinUsecase *usecase.CheckinUsecase
}

func NewCheckinService(checkinUsecase *usecase.CheckinUsecase) *CheckinService {
	return &CheckinService{checkinUsecase: checkinUsecase}
}

func (s *CheckinService) RegisterGrpc(gs *grpc.Server) {
}

func (s *CheckinService) RegisterHttp(hs *http.Server) {
	bbsuserv1.RegisterCheckinServiceHTTPServer(hs, s)
}

func (s *CheckinService) CheckIn(ctx context.Context, req *bbsuserv1.CheckIn_Req) (*bbsuserv1.CheckIn_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	return s.checkinUsecase.CheckIn(ctx, user.ID)
}

func (s *CheckinService) GetOverview(ctx context.Context, req *bbsuserv1.GetCheckinOverview_Req) (*bbsuserv1.GetCheckinOverview_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	return s.checkinUsecase.GetOverview(ctx, user.ID, req.GetMonth())
}
