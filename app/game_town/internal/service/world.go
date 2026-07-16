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

type WorldService struct {
	v1.UnimplementedGameTownWorldServiceServer
	gameUsecase *usecase.GameUsecase
}

func NewWorldService(gameUsecase *usecase.GameUsecase) *WorldService {
	return &WorldService{gameUsecase: gameUsecase}
}
func (s *WorldService) RegisterGrpc(gs *grpc.Server) { v1.RegisterGameTownWorldServiceServer(gs, s) }
func (s *WorldService) RegisterHttp(hs *http.Server) {}
func (s *WorldService) Create(ctx context.Context, req *v1.CreateGameTownWorld_Request) (*v1.CreateGameTownWorld_Response, error) {
	scale, ok := gameenum.WorldScaleMap.ToEnum(req.GetScale())
	if !ok {
		scale = gameenum.WorldScaleSmall
	}
	createResponse, err := s.gameUsecase.CreateWorld(ctx, &usecase.CreateWorldReq{PlayerID: req.GetCreatorPlayerId(), Description: req.GetDescription(), NpcCount: req.GetNpcCount(), LocationCount: req.GetLocationCount(), Scale: string(scale), Seed: req.Seed, Tags: req.GetStyleTags(), AgentConfigID: req.AgentConfigId})
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
	structValue := func(values map[string]any) *structpb.Struct {
		st, err := structpb.NewStruct(values)
		if err != nil {
			return &structpb.Struct{}
		}
		return st
	}
	result := createResponse.Result
	reply := &v1.CreateGameTownWorld_Response{}
	if row := result.World; row != nil {
		reply.World = &v1.CreateGameTownWorld_Response_GameTownWorld{
			CreatedAt:         timestamp(row.CreatedAt),
			UpdatedAt:         timestamp(row.UpdatedAt),
			Id:                row.ID,
			Code:              row.Code,
			Name:              row.Name,
			Description:       row.Description,
			Scale:             gameenum.WorldScaleMap.MustToProto(gameenum.WorldScale(row.Scale)),
			Status:            gameenum.WorldStatusMap.MustToProto(gameenum.WorldStatus(row.Status)),
			CreatorPlayerId:   row.CreatorPlayerID,
			DefaultLocationId: row.DefaultLocationID,
			Seed:              row.Seed,
			GenerationParams:  structValue(row.GenerationParams),
			GenerationSummary: row.GenerationSummary,
			AgentConfigId:     row.AgentConfigID,
		}
	}
	if row := result.DefaultLocation; row != nil {
		reply.DefaultLocation = &v1.CreateGameTownWorld_Response_GameTownLocation{
			CreatedAt:   timestamp(row.CreatedAt),
			UpdatedAt:   timestamp(row.UpdatedAt),
			Id:          row.ID,
			WorldId:     row.WorldID,
			Code:        row.Code,
			Name:        row.Name,
			Description: row.Description,
			Tags:        structValue(row.Tags),
			Sort:        row.Sort,
			Enabled:     row.Enabled,
		}
	}
	reply.Npcs = make([]*v1.CreateGameTownWorld_Response_GameTownNpc, 0, len(result.Npcs))
	for _, row := range result.Npcs {
		if row == nil {
			reply.Npcs = append(reply.Npcs, nil)
			continue
		}
		reply.Npcs = append(reply.Npcs, &v1.CreateGameTownWorld_Response_GameTownNpc{
			CreatedAt:         timestamp(row.CreatedAt),
			UpdatedAt:         timestamp(row.UpdatedAt),
			Id:                row.ID,
			WorldId:           row.WorldID,
			Code:              row.Code,
			Name:              row.Name,
			Role:              row.Role,
			Personality:       row.Personality,
			Goal:              row.Goal,
			Background:        row.Background,
			CurrentLocationId: row.CurrentLocationID,
			State:             gameenum.NpcStateMap.MustToProto(gameenum.NpcState(row.State)),
			SystemPrompt:      row.SystemPrompt,
			GeneratedProfile:  structValue(row.GeneratedProfile),
			Enabled:           row.Enabled,
		})
	}
	if row := result.State; row != nil {
		reply.State = &v1.CreateGameTownWorld_Response_GameTownWorldStateSnapshot{
			CreatedAt:     timestampValue(row.CreatedAt),
			Id:            row.ID,
			WorldId:       row.WorldID,
			TickCount:     row.TickCount,
			CurrentArc:    row.CurrentArc,
			Metrics:       structValue(row.Metrics),
			Summary:       row.Summary,
			ReasonEventId: row.ReasonEventID,
		}
	}
	reply.Events = make([]*v1.CreateGameTownWorld_Response_GameTownEvent, 0, len(result.Events))
	for _, row := range result.Events {
		if row == nil {
			reply.Events = append(reply.Events, nil)
			continue
		}
		reply.Events = append(reply.Events, &v1.CreateGameTownWorld_Response_GameTownEvent{
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
func (s *WorldService) Join(ctx context.Context, req *v1.JoinGameTownWorld_Request) (*v1.JoinGameTownWorld_Response, error) {
	joinResponse, err := s.gameUsecase.JoinWorld(ctx, &usecase.JoinWorldReq{PlayerID: req.GetPlayerId(), WorldCode: req.GetWorldCode()})
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
	structValue := func(values map[string]any) *structpb.Struct {
		st, err := structpb.NewStruct(values)
		if err != nil {
			return &structpb.Struct{}
		}
		return st
	}
	result := joinResponse.Result
	reply := &v1.JoinGameTownWorld_Response{}
	if row := result.World; row != nil {
		reply.World = &v1.JoinGameTownWorld_Response_GameTownWorld{
			CreatedAt:         timestamp(row.CreatedAt),
			UpdatedAt:         timestamp(row.UpdatedAt),
			Id:                row.ID,
			Code:              row.Code,
			Name:              row.Name,
			Description:       row.Description,
			Scale:             gameenum.WorldScaleMap.MustToProto(gameenum.WorldScale(row.Scale)),
			Status:            gameenum.WorldStatusMap.MustToProto(gameenum.WorldStatus(row.Status)),
			CreatorPlayerId:   row.CreatorPlayerID,
			DefaultLocationId: row.DefaultLocationID,
			Seed:              row.Seed,
			GenerationParams:  structValue(row.GenerationParams),
			GenerationSummary: row.GenerationSummary,
			AgentConfigId:     row.AgentConfigID,
		}
	}
	if row := result.Member; row != nil {
		reply.Member = &v1.JoinGameTownWorld_Response_GameTownWorldMember{
			CreatedAt:         timestamp(row.CreatedAt),
			UpdatedAt:         timestamp(row.UpdatedAt),
			Id:                row.ID,
			WorldId:           row.WorldID,
			PlayerId:          row.PlayerID,
			CurrentLocationId: row.CurrentLocationID,
			Role:              gameenum.WorldMemberRoleMap.MustToProto(gameenum.WorldMemberRole(row.Role)),
			JoinedAt:          timestampValue(row.JoinedAt),
			LastSeenAt:        timestampValue(row.LastSeenAt),
		}
	}
	if row := result.Location; row != nil {
		reply.Location = &v1.JoinGameTownWorld_Response_GameTownLocation{
			CreatedAt:   timestamp(row.CreatedAt),
			UpdatedAt:   timestamp(row.UpdatedAt),
			Id:          row.ID,
			WorldId:     row.WorldID,
			Code:        row.Code,
			Name:        row.Name,
			Description: row.Description,
			Tags:        structValue(row.Tags),
			Sort:        row.Sort,
			Enabled:     row.Enabled,
		}
	}
	return reply, nil
}
func (s *WorldService) Get(ctx context.Context, req *v1.GetGameTownWorld_Request) (*v1.GetGameTownWorld_Response, error) {
	getResponse, err := s.gameUsecase.GetWorld(ctx, &usecase.GetWorldReq{ID: req.GetId()})
	if err != nil {
		return nil, err
	}
	timestamp := func(t *time.Time) *timestamppb.Timestamp {
		if t == nil {
			return nil
		}
		return timestamppb.New(*t)
	}
	structValue := func(values map[string]any) *structpb.Struct {
		st, err := structpb.NewStruct(values)
		if err != nil {
			return &structpb.Struct{}
		}
		return st
	}
	reply := &v1.GetGameTownWorld_Response{}
	if row := getResponse.Row; row != nil {
		reply.Row = &v1.GetGameTownWorld_Response_GameTownWorld{
			CreatedAt:         timestamp(row.CreatedAt),
			UpdatedAt:         timestamp(row.UpdatedAt),
			Id:                row.ID,
			Code:              row.Code,
			Name:              row.Name,
			Description:       row.Description,
			Scale:             gameenum.WorldScaleMap.MustToProto(gameenum.WorldScale(row.Scale)),
			Status:            gameenum.WorldStatusMap.MustToProto(gameenum.WorldStatus(row.Status)),
			CreatorPlayerId:   row.CreatorPlayerID,
			DefaultLocationId: row.DefaultLocationID,
			Seed:              row.Seed,
			GenerationParams:  structValue(row.GenerationParams),
			GenerationSummary: row.GenerationSummary,
			AgentConfigId:     row.AgentConfigID,
		}
	}
	return reply, nil
}
func (s *WorldService) Page(ctx context.Context, req *v1.PageGameTownWorlds_Request) (*v1.PageGameTownWorlds_Response, error) {
	var status *string
	if req.Status != nil && *req.Status != v1.GameTownWorldStatus_GAME_TOWN_WORLD_STATUS_UNSPECIFIED {
		value, ok := gameenum.WorldStatusMap.ToEnum(*req.Status)
		if ok {
			statusValue := string(value)
			status = &statusValue
		}
	}
	pageResponse, err := s.gameUsecase.PageWorlds(ctx, &usecase.PageWorldsReq{Page: req.GetPage(), CreatorPlayerID: req.CreatorPlayerId, Status: status})
	if err != nil {
		return nil, err
	}
	timestamp := func(t *time.Time) *timestamppb.Timestamp {
		if t == nil {
			return nil
		}
		return timestamppb.New(*t)
	}
	structValue := func(values map[string]any) *structpb.Struct {
		st, err := structpb.NewStruct(values)
		if err != nil {
			return &structpb.Struct{}
		}
		return st
	}
	reply := &v1.PageGameTownWorlds_Response{Page: pageResponse.Page, Rows: make([]*v1.PageGameTownWorlds_Response_GameTownWorld, 0, len(pageResponse.Rows))}
	for _, row := range pageResponse.Rows {
		if row == nil {
			reply.Rows = append(reply.Rows, nil)
			continue
		}
		reply.Rows = append(reply.Rows, &v1.PageGameTownWorlds_Response_GameTownWorld{
			CreatedAt:         timestamp(row.CreatedAt),
			UpdatedAt:         timestamp(row.UpdatedAt),
			Id:                row.ID,
			Code:              row.Code,
			Name:              row.Name,
			Description:       row.Description,
			Scale:             gameenum.WorldScaleMap.MustToProto(gameenum.WorldScale(row.Scale)),
			Status:            gameenum.WorldStatusMap.MustToProto(gameenum.WorldStatus(row.Status)),
			CreatorPlayerId:   row.CreatorPlayerID,
			DefaultLocationId: row.DefaultLocationID,
			Seed:              row.Seed,
			GenerationParams:  structValue(row.GenerationParams),
			GenerationSummary: row.GenerationSummary,
			AgentConfigId:     row.AgentConfigID,
		})
	}
	return reply, nil
}
func (s *WorldService) GetState(ctx context.Context, req *v1.GetGameTownWorldState_Request) (*v1.GetGameTownWorldState_Response, error) {
	stateResponse, err := s.gameUsecase.GetState(ctx, &usecase.GetStateReq{WorldID: req.GetWorldId()})
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
	structValue := func(values map[string]any) *structpb.Struct {
		st, err := structpb.NewStruct(values)
		if err != nil {
			return &structpb.Struct{}
		}
		return st
	}
	reply := &v1.GetGameTownWorldState_Response{Metrics: make([]*v1.GetGameTownWorldState_Response_GameTownWorldMetricDefinition, 0, len(stateResponse.Metrics))}
	if row := stateResponse.State; row != nil {
		reply.Row = &v1.GetGameTownWorldState_Response_GameTownWorldStateSnapshot{
			CreatedAt:     timestampValue(row.CreatedAt),
			Id:            row.ID,
			WorldId:       row.WorldID,
			TickCount:     row.TickCount,
			CurrentArc:    row.CurrentArc,
			Metrics:       structValue(row.Metrics),
			Summary:       row.Summary,
			ReasonEventId: row.ReasonEventID,
		}
	}
	for _, row := range stateResponse.Metrics {
		if row == nil {
			reply.Metrics = append(reply.Metrics, nil)
			continue
		}
		reply.Metrics = append(reply.Metrics, &v1.GetGameTownWorldState_Response_GameTownWorldMetricDefinition{
			CreatedAt:    timestamp(row.CreatedAt),
			UpdatedAt:    timestamp(row.UpdatedAt),
			Id:           row.ID,
			WorldId:      row.WorldID,
			Key:          row.Key,
			Name:         row.Name,
			Description:  row.Description,
			MinValue:     row.MinValue,
			MaxValue:     row.MaxValue,
			InitialValue: row.InitialValue,
		})
	}
	return reply, nil
}
func (s *WorldService) Tick(ctx context.Context, req *v1.TickGameTownWorld_Request) (*v1.TickGameTownWorld_Response, error) {
	tickResponse, err := s.gameUsecase.Tick(ctx, &usecase.TickReq{WorldID: req.GetWorldId(), OperatorPlayerID: req.GetOperatorPlayerId(), Limit: req.GetRecentEventLimit()})
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
	reply := &v1.TickGameTownWorld_Response{Events: make([]*v1.TickGameTownWorld_Response_GameTownEvent, 0, len(tickResponse.Events))}
	if row := tickResponse.State; row != nil {
		reply.State = &v1.TickGameTownWorld_Response_GameTownWorldStateSnapshot{
			CreatedAt:     timestampValue(row.CreatedAt),
			Id:            row.ID,
			WorldId:       row.WorldID,
			TickCount:     row.TickCount,
			CurrentArc:    row.CurrentArc,
			Metrics:       structValue(row.Metrics),
			Summary:       row.Summary,
			ReasonEventId: row.ReasonEventID,
		}
	}
	for _, row := range tickResponse.Events {
		if row == nil {
			reply.Events = append(reply.Events, nil)
			continue
		}
		reply.Events = append(reply.Events, &v1.TickGameTownWorld_Response_GameTownEvent{
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
