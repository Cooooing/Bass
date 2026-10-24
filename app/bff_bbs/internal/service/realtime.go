package service

import (
	"common/pkg/apperror"
	"common/pkg/constant"
	commonmodel "common/pkg/model"
	"common/pkg/util"
	bffbbsv1 "common/proto/gen/bff_bbs/v1/realtime"
	cerrors "common/proto/gen/common/errors"
	"context"

	"bff_bbs/internal/biz/usecase"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

type RealtimeService struct {
	bffbbsv1.UnimplementedRealtimeServiceServer
	realtime *usecase.RealtimeUsecase
}

func NewRealtimeService(realtime *usecase.RealtimeUsecase) *RealtimeService {
	return &RealtimeService{realtime: realtime}
}
func (*RealtimeService) RegisterGrpc(*grpc.Server) {}
func (s *RealtimeService) RegisterHTTP(server *http.Server) {
	bffbbsv1.RegisterRealtimeServiceHTTPServer(server, s)
}
func (s *RealtimeService) CreateTicket(ctx context.Context, _ *bffbbsv1.CreateRealtimeTicket_Req) (*bffbbsv1.CreateRealtimeTicket_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil || user.ID <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	ticket, expires, err := s.realtime.CreateTicket(ctx, user.ID)
	if err != nil {
		return nil, err
	}
	return &bffbbsv1.CreateRealtimeTicket_Resp{Ticket: ticket, ExpiresInSeconds: expires, Path: "/v1/stream/bbs/sse"}, nil
}
