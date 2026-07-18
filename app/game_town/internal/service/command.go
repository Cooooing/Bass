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

type CommandService struct {
	v1.UnimplementedGameTownCommandServiceServer
	gameUsecase *usecase.GameUsecase
}

func NewCommandService(gameUsecase *usecase.GameUsecase) *CommandService {
	return &CommandService{gameUsecase: gameUsecase}
}
func (s *CommandService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterGameTownCommandServiceServer(gs, s)
}
func (s *CommandService) RegisterHttp(hs *http.Server) {}
func (s *CommandService) Execute(ctx context.Context, req *v1.ExecuteGameTownCommand_Req) (*v1.ExecuteGameTownCommand_Resp, error) {
	result, err := s.gameUsecase.ExecuteCommand(ctx, &usecase.ExecuteCommandReq{SessionID: req.GetSessionId(), PlayerID: req.GetPlayerId(), Raw: req.GetRawText()})
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
	reply := &v1.ExecuteGameTownCommand_Resp{FeedbackLines: result.FeedbackLines}
	if row := result.Command; row != nil {
		reply.Command = &v1.ExecuteGameTownCommand_Resp_GameTownCommand{
			CreatedAt:     timestampValue(row.CreatedAt),
			HandledAt:     timestamp(row.HandledAt),
			Id:            row.ID,
			WorldId:       row.WorldID,
			SessionId:     row.SessionID,
			PlayerId:      row.PlayerID,
			RawText:       row.RawText,
			Type:          gameenum.CommandTypeMap.MustToProto(gameenum.CommandType(row.Type)),
			ParsedPayload: structValue(row.ParsedPayload),
			Status:        gameenum.CommandStatusMap.MustToProto(gameenum.CommandStatus(row.Status)),
			ErrorCode:     row.ErrorCode,
			ResultSummary: row.ResultSummary,
		}
	}
	if row := result.CurrentWorld; row != nil {
		reply.CurrentWorld = &v1.ExecuteGameTownCommand_Resp_GameTownWorld{
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
	if row := result.CurrentLocation; row != nil {
		reply.CurrentLocation = &v1.ExecuteGameTownCommand_Resp_GameTownLocation{
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
	reply.VisibleNpcs = make([]*v1.ExecuteGameTownCommand_Resp_GameTownNpc, 0, len(result.VisibleNpcs))
	for _, row := range result.VisibleNpcs {
		if row == nil {
			reply.VisibleNpcs = append(reply.VisibleNpcs, nil)
			continue
		}
		reply.VisibleNpcs = append(reply.VisibleNpcs, &v1.ExecuteGameTownCommand_Resp_GameTownNpc{
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
	if row := result.WorldState; row != nil {
		reply.WorldState = &v1.ExecuteGameTownCommand_Resp_GameTownWorldStateSnapshot{
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
	reply.Events = make([]*v1.ExecuteGameTownCommand_Resp_GameTownEvent, 0, len(result.Events))
	for _, row := range result.Events {
		if row == nil {
			reply.Events = append(reply.Events, nil)
			continue
		}
		reply.Events = append(reply.Events, &v1.ExecuteGameTownCommand_Resp_GameTownEvent{
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
func (s *CommandService) Page(ctx context.Context, req *v1.PageGameTownCommands_Req) (*v1.PageGameTownCommands_Resp, error) {
	pageResp, err := s.gameUsecase.PageCommands(ctx, &usecase.PageCommandsReq{Page: req.GetPage(), WorldID: req.WorldId, SessionID: req.SessionId, PlayerID: req.PlayerId})
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
	reply := &v1.PageGameTownCommands_Resp{Page: pageResp.Page, Rows: make([]*v1.PageGameTownCommands_Resp_GameTownCommand, 0, len(pageResp.Rows))}
	for _, row := range pageResp.Rows {
		if row == nil {
			reply.Rows = append(reply.Rows, nil)
			continue
		}
		reply.Rows = append(reply.Rows, &v1.PageGameTownCommands_Resp_GameTownCommand{
			CreatedAt:     timestampValue(row.CreatedAt),
			HandledAt:     timestamp(row.HandledAt),
			Id:            row.ID,
			WorldId:       row.WorldID,
			SessionId:     row.SessionID,
			PlayerId:      row.PlayerID,
			RawText:       row.RawText,
			Type:          gameenum.CommandTypeMap.MustToProto(gameenum.CommandType(row.Type)),
			ParsedPayload: structValue(row.ParsedPayload),
			Status:        gameenum.CommandStatusMap.MustToProto(gameenum.CommandStatus(row.Status)),
			ErrorCode:     row.ErrorCode,
			ResultSummary: row.ResultSummary,
		})
	}
	return reply, nil
}
func (s *CommandService) Replay(ctx context.Context, req *v1.ReplayGameTownCommands_Req) (*v1.ReplayGameTownCommands_Resp, error) {
	replayResp, err := s.gameUsecase.ReplayCommands(ctx, &usecase.ReplayCommandsReq{SessionID: req.GetSessionId(), PlayerID: req.GetPlayerId()})
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
	reply := &v1.ReplayGameTownCommands_Resp{
		Commands: make([]*v1.ReplayGameTownCommands_Resp_GameTownCommand, 0, len(replayResp.Commands)),
		Events:   make([]*v1.ReplayGameTownCommands_Resp_GameTownEvent, 0, len(replayResp.Events)),
	}
	for _, row := range replayResp.Commands {
		if row == nil {
			reply.Commands = append(reply.Commands, nil)
			continue
		}
		reply.Commands = append(reply.Commands, &v1.ReplayGameTownCommands_Resp_GameTownCommand{
			CreatedAt:     timestampValue(row.CreatedAt),
			HandledAt:     timestamp(row.HandledAt),
			Id:            row.ID,
			WorldId:       row.WorldID,
			SessionId:     row.SessionID,
			PlayerId:      row.PlayerID,
			RawText:       row.RawText,
			Type:          gameenum.CommandTypeMap.MustToProto(gameenum.CommandType(row.Type)),
			ParsedPayload: structValue(row.ParsedPayload),
			Status:        gameenum.CommandStatusMap.MustToProto(gameenum.CommandStatus(row.Status)),
			ErrorCode:     row.ErrorCode,
			ResultSummary: row.ResultSummary,
		})
	}
	for _, row := range replayResp.Events {
		if row == nil {
			reply.Events = append(reply.Events, nil)
			continue
		}
		reply.Events = append(reply.Events, &v1.ReplayGameTownCommands_Resp_GameTownEvent{
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
