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

type WebSocketTicketService struct {
	v1.UnimplementedWebSocketServiceServer
	webSocketUsecase *usecase.WebSocketUsecase
	path             string
}

func NewWebSocketTicketService(
	webSocketUsecase *usecase.WebSocketUsecase,
) *WebSocketTicketService {
	return &WebSocketTicketService{
		webSocketUsecase: webSocketUsecase,
		path:             "/ws",
	}
}

func (s *WebSocketTicketService) RegisterGrpc(*grpc.Server) {
}

func (s *WebSocketTicketService) RegisterHttp(hs *http.Server) {
	v1.RegisterWebSocketServiceHTTPServer(hs, s)
}

func (s *WebSocketTicketService) CreateTicket(ctx context.Context, req *v1.CreateWebSocketTicket_Req) (*v1.CreateWebSocketTicket_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	ticket, err := s.webSocketUsecase.CreateTicket(ctx, &usecase.CreateWebSocketTicketReq{
		UserID:      user.ID,
		CharacterID: req.GetCharacterId(),
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateWebSocketTicket_Resp{
		CharacterId:      ticket.CharacterID,
		Ticket:           ticket.Ticket,
		ExpiresInSeconds: int64(ticket.RemainingDuration.Seconds()),
		Path:             s.path,
	}, nil
}
