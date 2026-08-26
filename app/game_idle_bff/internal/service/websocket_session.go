package service

import (
	"common/pkg/apperror"
	"common/pkg/constant"
	commonmodel "common/pkg/model"
	"common/pkg/util"
	cerrors "common/proto/gen/common/errors"
	v1 "common/proto/gen/game_idle_bff/v1"
	"context"
	"game_idle_bff/internal/biz/usecase"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

type WebSocketSessionService struct {
	v1.UnimplementedGameIdleBffWebSocketServiceServer
	webSocketUsecase *usecase.WebSocketUsecase
}

func NewWebSocketSessionService(
	webSocketUsecase *usecase.WebSocketUsecase,
) *WebSocketSessionService {
	return &WebSocketSessionService{
		webSocketUsecase: webSocketUsecase,
	}
}

func (s *WebSocketSessionService) RegisterGrpc(*grpc.Server) {
}

func (s *WebSocketSessionService) RegisterHttp(hs *http.Server) {
	v1.RegisterGameIdleBffWebSocketServiceHTTPServer(hs, s)
}

func (s *WebSocketSessionService) CreateSession(ctx context.Context, req *v1.CreateGameIdleBffWebSocketSession_Req) (*v1.CreateGameIdleBffWebSocketSession_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	session, err := s.webSocketUsecase.CreateSession(ctx, &usecase.CreateWebSocketSessionReq{
		UserID:      user.ID,
		CharacterID: req.GetCharacterId(),
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateGameIdleBffWebSocketSession_Resp{
		SessionId:        session.SessionID,
		ExpiresInSeconds: int64(session.RemainingDuration.Seconds()),
		Path:             webSocketPath,
	}, nil
}
