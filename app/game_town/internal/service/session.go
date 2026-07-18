package service

import (
	v1 "common/proto/gen/game_town/v1"
	"context"
	"game_town/internal/biz/usecase"
	gameenum "game_town/internal/enum"
	"time"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SessionService struct {
	v1.UnimplementedGameTownSessionServiceServer
	gameUsecase *usecase.GameUsecase
}

func NewSessionService(gameUsecase *usecase.GameUsecase) *SessionService {
	return &SessionService{gameUsecase: gameUsecase}
}
func (s *SessionService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterGameTownSessionServiceServer(gs, s)
}
func (s *SessionService) RegisterHttp(hs *http.Server) {}
func (s *SessionService) Start(ctx context.Context, req *v1.StartGameTownSession_Req) (*v1.StartGameTownSession_Resp, error) {
	clientType := ""
	if value, ok := gameenum.SessionClientTypeMap.ToEnum(req.GetClientType()); ok {
		clientType = string(value)
	}
	row, err := s.gameUsecase.StartSession(ctx, &usecase.StartSessionReq{PlayerID: req.PlayerId, ClientType: clientType})
	if err != nil {
		return nil, err
	}
	timestamp := func(t *time.Time) *timestamppb.Timestamp {
		if t == nil {
			return nil
		}
		return timestamppb.New(*t)
	}
	timestampValue := func(t time.Time) *timestamppb.Timestamp {
		if t.IsZero() {
			return nil
		}
		return timestamppb.New(t)
	}
	reply := &v1.StartGameTownSession_Resp{}
	if row != nil {
		reply.Row = &v1.StartGameTownSession_Resp_GameTownSession{
			CreatedAt:      timestamp(row.CreatedAt),
			UpdatedAt:      timestamp(row.UpdatedAt),
			Id:             row.ID,
			PlayerId:       row.PlayerID,
			CurrentWorldId: row.CurrentWorldID,
			ClientType:     gameenum.SessionClientTypeMap.MustToProto(gameenum.SessionClientType(row.ClientType)),
			StartedAt:      timestampValue(row.StartedAt),
			LastSeenAt:     timestampValue(row.LastSeenAt),
			EndedAt:        timestamp(row.EndedAt),
		}
	}
	return reply, nil
}
func (s *SessionService) End(ctx context.Context, req *v1.EndGameTownSession_Req) (*v1.EndGameTownSession_Resp, error) {
	row, err := s.gameUsecase.EndSession(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	timestamp := func(t *time.Time) *timestamppb.Timestamp {
		if t == nil {
			return nil
		}
		return timestamppb.New(*t)
	}
	timestampValue := func(t time.Time) *timestamppb.Timestamp {
		if t.IsZero() {
			return nil
		}
		return timestamppb.New(t)
	}
	reply := &v1.EndGameTownSession_Resp{}
	if row != nil {
		reply.Row = &v1.EndGameTownSession_Resp_GameTownSession{
			CreatedAt:      timestamp(row.CreatedAt),
			UpdatedAt:      timestamp(row.UpdatedAt),
			Id:             row.ID,
			PlayerId:       row.PlayerID,
			CurrentWorldId: row.CurrentWorldID,
			ClientType:     gameenum.SessionClientTypeMap.MustToProto(gameenum.SessionClientType(row.ClientType)),
			StartedAt:      timestampValue(row.StartedAt),
			LastSeenAt:     timestampValue(row.LastSeenAt),
			EndedAt:        timestamp(row.EndedAt),
		}
	}
	return reply, nil
}
