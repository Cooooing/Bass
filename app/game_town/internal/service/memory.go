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

type MemoryService struct {
	v1.UnimplementedGameTownMemoryServiceServer
	gameUsecase *usecase.GameUsecase
}

func NewMemoryService(gameUsecase *usecase.GameUsecase) *MemoryService {
	return &MemoryService{gameUsecase: gameUsecase}
}
func (s *MemoryService) RegisterGrpc(gs *grpc.Server) { v1.RegisterGameTownMemoryServiceServer(gs, s) }
func (s *MemoryService) RegisterHttp(hs *http.Server) {}
func (s *MemoryService) List(ctx context.Context, req *v1.ListGameTownMemories_Request) (*v1.ListGameTownMemories_Response, error) {
	var typ *string
	if req.Type != nil && *req.Type != v1.GameTownMemoryType_GAME_TOWN_MEMORY_TYPE_UNSPECIFIED {
		value, ok := gameenum.MemoryTypeMap.ToEnum(*req.Type)
		if ok {
			typeValue := string(value)
			typ = &typeValue
		}
	}
	listResponse, err := s.gameUsecase.ListMemories(ctx, &usecase.ListMemoriesReq{WorldID: req.GetWorldId(), PlayerID: req.GetPlayerId(), NpcID: req.NpcId, Type: typ})
	if err != nil {
		return nil, err
	}
	timestamp := func(t *time.Time) *timestamppb.Timestamp {
		if t == nil {
			return nil
		}
		return timestamppb.New(*t)
	}
	reply := &v1.ListGameTownMemories_Response{Rows: make([]*v1.ListGameTownMemories_Response_GameTownMemory, 0, len(listResponse.Rows))}
	for _, row := range listResponse.Rows {
		if row == nil {
			reply.Rows = append(reply.Rows, nil)
			continue
		}
		reply.Rows = append(reply.Rows, &v1.ListGameTownMemories_Response_GameTownMemory{
			CreatedAt:      timestamp(row.CreatedAt),
			UpdatedAt:      timestamp(row.UpdatedAt),
			Id:             row.ID,
			WorldId:        row.WorldID,
			PlayerId:       row.PlayerID,
			NpcId:          row.NpcID,
			Type:           gameenum.MemoryTypeMap.MustToProto(gameenum.MemoryType(row.Type)),
			Content:        row.Content,
			Importance:     row.Importance,
			SourceEventId:  row.SourceEventID,
			LastRecalledAt: timestamp(row.LastRecalledAt),
			ExpiresAt:      timestamp(row.ExpiresAt),
		})
	}
	return reply, nil
}
