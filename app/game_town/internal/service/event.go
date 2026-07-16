package service

import (
	v1 "common/proto/gen/game_town/v1"
	"context"
	"game_town/internal/biz/usecase"
	gameenum "game_town/internal/enum"
	"time"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type EventService struct {
	v1.UnimplementedGameTownEventServiceServer
	gameUsecase *usecase.GameUsecase
}

func NewEventService(gameUsecase *usecase.GameUsecase) *EventService {
	return &EventService{gameUsecase: gameUsecase}
}
func (s *EventService) RegisterGrpc(gs *grpc.Server) { v1.RegisterGameTownEventServiceServer(gs, s) }
func (s *EventService) RegisterHttp(hs *http.Server) {}
func (s *EventService) Page(ctx context.Context, req *v1.PageGameTownEvents_Request) (*v1.PageGameTownEvents_Response, error) {
	var typ *string
	if req.Type != nil && *req.Type != v1.GameTownEventType_GAME_TOWN_EVENT_TYPE_UNSPECIFIED {
		value, ok := gameenum.EventTypeMap.ToEnum(*req.Type)
		if ok {
			typeValue := string(value)
			typ = &typeValue
		}
	}
	pageResponse, err := s.gameUsecase.PageEvents(ctx, &usecase.PageEventsReq{Page: req.GetPage(), WorldID: req.GetWorldId(), ActorPlayerID: req.ActorPlayerId, TargetNpcID: req.TargetNpcId, Type: typ})
	if err != nil {
		return nil, err
	}
	timestampValue := func(t time.Time) *timestamppb.Timestamp {
		if t.IsZero() {
			return nil
		}
		return timestamppb.New(t)
	}
	structValue := func(values map[string]any) *structpb.Struct {
		st, err := structpb.NewStruct(values)
		if err != nil {
			return &structpb.Struct{}
		}
		return st
	}
	reply := &v1.PageGameTownEvents_Response{Page: pageResponse.Page, Rows: make([]*v1.PageGameTownEvents_Response_GameTownEvent, 0, len(pageResponse.Rows))}
	for _, row := range pageResponse.Rows {
		if row == nil {
			reply.Rows = append(reply.Rows, nil)
			continue
		}
		reply.Rows = append(reply.Rows, &v1.PageGameTownEvents_Response_GameTownEvent{
			CreatedAt:     timestampValue(row.CreatedAt),
			Id:            row.ID,
			WorldId:       row.WorldID,
			Type:          gameenum.EventTypeMap.MustToProto(gameenum.EventType(row.Type)),
			ActorPlayerId: row.ActorPlayerID,
			TargetNpcId:   row.TargetNpcID,
			LocationId:    row.LocationID,
			CommandId:     row.CommandID,
			Summary:       row.Summary,
			Content:       row.Content,
			Effects:       structValue(row.Effects),
			Metadata:      structValue(row.Metadata),
			OccurredAt:    timestampValue(row.OccurredAt),
		})
	}
	return reply, nil
}
