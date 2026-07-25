package service

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"common/pkg/apperror"
	"common/proto/gen/common"
	cerrors "common/proto/gen/common/errors"
	v1 "common/proto/gen/game_town/v1"
	"game_town/internal/biz/base"
	"game_town/internal/biz/model"
	"game_town/internal/biz/usecase"
	"game_town/internal/enum"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	ggrpc "google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type EventService struct {
	v1.UnimplementedGameTownEventServiceServer
	usecase *usecase.EventUsecase
}

func NewEventService(
	usecase *usecase.EventUsecase,
) *EventService {
	return &EventService{
		usecase: usecase,
	}
}

func (s *EventService) RegisterGrpc(server *grpc.Server) {
	v1.RegisterGameTownEventServiceServer(server, s)
}

func (s *EventService) RegisterHttp(*http.Server) {
}

func (s *EventService) Page(ctx context.Context, req *v1.PageGameTownEvents_Request) (*v1.PageGameTownEvents_Resp, error) {
	if req.GetWorldId() <= 0 || req.GetPlayerId() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}

	page := base.PageRequest{}
	if req.GetPage() != nil {
		page.Page = int64(req.GetPage().GetPage())
		page.Size = int64(req.GetPage().GetSize())
	}

	var eventType *enum.EventType
	if req.Type != nil {
		value, ok := enum.EventTypeMap.ToEnum(*req.Type)
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		eventType = new(value)
	}

	resp, err := s.usecase.Page(ctx, &usecase.PageEventsReq{
		Page:          page,
		WorldID:       req.GetWorldId(),
		PlayerID:      req.GetPlayerId(),
		AfterSequence: req.GetAfterSequence(),
		Type:          eventType,
		SkipTotal:     req.GetAfterSequence() > 0,
	})
	if err != nil {
		return nil, err
	}

	reply := &v1.PageGameTownEvents_Resp{
		Page: &common.PageResp{
			Page:  uint32(resp.Page.Page),
			Size:  uint32(resp.Page.Size),
			Total: uint32(resp.Page.Total),
		},
		Rows: make([]*v1.PageGameTownEvents_Resp_Row, 0, len(resp.Rows)),
	}
	for _, value := range resp.Rows {
		row := value.Event
		observation := value.Observation
		if row == nil || observation == nil {
			continue
		}
		payload, err := newEventPayload(row.Payload)
		if err != nil {
			return nil, err
		}

		suggestions := make([]*v1.PageGameTownEvents_Resp_Row_SuggestedAction, 0)
		for _, raw := range sliceValues(row.Payload["suggested_actions"]) {
			item, ok := raw.(map[string]any)
			if !ok {
				continue
			}
			label, _ := item["label"].(string)
			content, _ := item["content"].(string)
			targets := make([]*v1.PageGameTownEvents_Resp_Row_SuggestedAction_EntityRef, 0)
			for _, rawTarget := range sliceValues(item["targets"]) {
				target, ok := rawTarget.(map[string]any)
				if !ok {
					continue
				}
				targetType, _ := target["type"].(string)
				targetID := int64Value(target["id"])
				typeValue := enum.EntityType(targetType)
				targets = append(targets, &v1.PageGameTownEvents_Resp_Row_SuggestedAction_EntityRef{
					Type: enum.EntityTypeMap.MustToProto(typeValue),
					Id:   targetID,
				})
			}
			suggestions = append(suggestions, &v1.PageGameTownEvents_Resp_Row_SuggestedAction{
				Label:   label,
				Content: content,
				Targets: targets,
			})
		}

		reply.Rows = append(reply.Rows, &v1.PageGameTownEvents_Resp_Row{
			Id:               row.ID,
			WorldId:          row.WorldID,
			Sequence:         row.Sequence,
			Type:             enum.EventTypeMap.MustToProto(row.Type),
			ActorPlayerId:    row.ActorPlayerID,
			NpcId:            row.NpcID,
			LocationId:       row.LocationID,
			CausationEventId: row.CausationEventID,
			Summary:          row.Summary,
			Content:          row.Content,
			Payload:          payload,
			OccurredAt:       timestamppb.New(row.OccurredAt),
			WorldTime:        timestamppb.New(row.WorldTime),
			Source:           enum.ObservationSourceMap.MustToProto(observation.Source),
			Certainty:        enum.KnowledgeCertaintyMap.MustToProto(observation.Certainty),
			SuggestedActions: suggestions,
		})
	}
	return reply, nil
}

func (s *EventService) Watch(req *v1.WatchGameTownEvents_Request, stream ggrpc.ServerStreamingServer[v1.WatchGameTownEvents_Resp]) error {
	if req.GetWorldId() <= 0 || req.GetPlayerId() <= 0 {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}

	events, unsubscribe := s.usecase.Watch(req.GetWorldId())
	defer unsubscribe()

	lastSequence := req.GetAfterSequence()
	send := func(row *model.Event, observation *model.Observation) error {
		if row == nil || observation == nil || row.Sequence <= lastSequence {
			return nil
		}

		payload, err := newEventPayload(row.Payload)
		if err != nil {
			return err
		}
		suggestions := make([]*v1.WatchGameTownEvents_Resp_SuggestedAction, 0)
		for _, raw := range sliceValues(row.Payload["suggested_actions"]) {
			item, ok := raw.(map[string]any)
			if !ok {
				continue
			}
			label, _ := item["label"].(string)
			content, _ := item["content"].(string)
			targets := make([]*v1.WatchGameTownEvents_Resp_SuggestedAction_EntityRef, 0)
			for _, rawTarget := range sliceValues(item["targets"]) {
				target, ok := rawTarget.(map[string]any)
				if !ok {
					continue
				}
				targetType, _ := target["type"].(string)
				targetID := int64Value(target["id"])
				typeValue := enum.EntityType(targetType)
				targets = append(targets, &v1.WatchGameTownEvents_Resp_SuggestedAction_EntityRef{
					Type: enum.EntityTypeMap.MustToProto(typeValue),
					Id:   targetID,
				})
			}
			suggestions = append(suggestions, &v1.WatchGameTownEvents_Resp_SuggestedAction{
				Label:   label,
				Content: content,
				Targets: targets,
			})
		}

		err = stream.Send(&v1.WatchGameTownEvents_Resp{
			Id:               row.ID,
			WorldId:          row.WorldID,
			Sequence:         row.Sequence,
			Type:             enum.EventTypeMap.MustToProto(row.Type),
			ActorPlayerId:    row.ActorPlayerID,
			NpcId:            row.NpcID,
			LocationId:       row.LocationID,
			CausationEventId: row.CausationEventID,
			Summary:          row.Summary,
			Content:          row.Content,
			Payload:          payload,
			OccurredAt:       timestamppb.New(row.OccurredAt),
			WorldTime:        timestamppb.New(row.WorldTime),
			Source:           enum.ObservationSourceMap.MustToProto(observation.Source),
			Certainty:        enum.KnowledgeCertaintyMap.MustToProto(observation.Certainty),
			SuggestedActions: suggestions,
		})
		if err != nil {
			return err
		}
		lastSequence = row.Sequence
		return nil
	}

	history, err := s.usecase.ListAfter(stream.Context(), &usecase.ListEventsAfterReq{
		WorldID:       req.GetWorldId(),
		PlayerID:      req.GetPlayerId(),
		AfterSequence: lastSequence,
	})
	if err != nil {
		return err
	}
	for _, value := range history {
		if err = send(value.Event, value.Observation); err != nil {
			return err
		}
	}

	for {
		select {
		case <-stream.Context().Done():
			return stream.Context().Err()
		case row, ok := <-events:
			if !ok || row == nil {
				return nil
			}
			observation, visible, err := s.usecase.VisibleToPlayer(
				stream.Context(),
				req.GetWorldId(),
				req.GetPlayerId(),
				row.ID,
			)
			if err != nil {
				return err
			}
			if visible {
				if err = send(row, observation); err != nil {
					return err
				}
			}
		}
	}
}

func newEventPayload(payload map[string]any) (*structpb.Struct, error) {
	if payload == nil {
		return structpb.NewStruct(map[string]any{})
	}
	normalized, ok := normalizeProtoValue(payload).(map[string]any)
	if !ok {
		return nil, fmt.Errorf("invalid event payload")
	}
	return structpb.NewStruct(normalized)
}

func normalizeProtoValue(value any) any {
	switch typed := value.(type) {
	case nil:
		return nil
	case json.Number:
		return jsonNumberValue(typed)
	case fmt.Stringer:
		return typed.String()
	case map[string]any:
		result := make(map[string]any, len(typed))
		for key, item := range typed {
			result[key] = normalizeProtoValue(item)
		}
		return result
	case []any:
		result := make([]any, 0, len(typed))
		for _, item := range typed {
			result = append(result, normalizeProtoValue(item))
		}
		return result
	case []string:
		result := make([]any, 0, len(typed))
		for _, item := range typed {
			result = append(result, item)
		}
		return result
	case []int:
		result := make([]any, 0, len(typed))
		for _, item := range typed {
			result = append(result, float64(item))
		}
		return result
	case []int64:
		result := make([]any, 0, len(typed))
		for _, item := range typed {
			result = append(result, float64(item))
		}
		return result
	case []float64:
		result := make([]any, 0, len(typed))
		for _, item := range typed {
			result = append(result, item)
		}
		return result
	case []bool:
		result := make([]any, 0, len(typed))
		for _, item := range typed {
			result = append(result, item)
		}
		return result
	case bool, string, float64:
		return typed
	case int:
		return float64(typed)
	case int8:
		return float64(typed)
	case int16:
		return float64(typed)
	case int32:
		return float64(typed)
	case int64:
		return float64(typed)
	case uint:
		return float64(typed)
	case uint8:
		return float64(typed)
	case uint16:
		return float64(typed)
	case uint32:
		return float64(typed)
	case uint64:
		return float64(typed)
	case float32:
		return float64(typed)
	default:
		return normalizeReflectValue(value)
	}
}

func normalizeReflectValue(value any) any {
	reflected := reflect.ValueOf(value)
	if !reflected.IsValid() {
		return nil
	}
	if reflected.Kind() == reflect.Pointer || reflected.Kind() == reflect.Interface {
		if reflected.IsNil() {
			return nil
		}
		return normalizeProtoValue(reflected.Elem().Interface())
	}
	if reflected.Kind() == reflect.Map && reflected.Type().Key().Kind() == reflect.String {
		result := make(map[string]any, reflected.Len())
		iterator := reflected.MapRange()
		for iterator.Next() {
			result[iterator.Key().String()] = normalizeProtoValue(iterator.Value().Interface())
		}
		return result
	}
	if reflected.Kind() == reflect.Slice || reflected.Kind() == reflect.Array {
		result := make([]any, 0, reflected.Len())
		for index := 0; index < reflected.Len(); index++ {
			result = append(result, normalizeProtoValue(reflected.Index(index).Interface()))
		}
		return result
	}
	return fmt.Sprint(value)
}

func jsonNumberValue(value json.Number) any {
	if parsed, err := value.Int64(); err == nil {
		return float64(parsed)
	}
	if parsed, err := value.Float64(); err == nil {
		return parsed
	}
	return value.String()
}

func sliceValues(value any) []any {
	normalized, ok := normalizeProtoValue(value).([]any)
	if !ok {
		return nil
	}
	return normalized
}

func int64Value(value any) int64 {
	switch typed := value.(type) {
	case int:
		return int64(typed)
	case int8:
		return int64(typed)
	case int16:
		return int64(typed)
	case int32:
		return int64(typed)
	case int64:
		return typed
	case uint:
		return int64(typed)
	case uint8:
		return int64(typed)
	case uint16:
		return int64(typed)
	case uint32:
		return int64(typed)
	case uint64:
		if typed > uint64(^uint64(0)>>1) {
			return 0
		}
		return int64(typed)
	case float32:
		return int64(typed)
	case float64:
		return int64(typed)
	case string:
		return parseInt64String(typed)
	case fmt.Stringer:
		return parseInt64String(typed.String())
	default:
		return 0
	}
}

func parseInt64String(value string) int64 {
	value = strings.TrimSpace(value)
	if value == "" {
		return 0
	}
	result, err := strconv.ParseInt(value, 10, 64)
	if err == nil {
		return result
	}
	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0
	}
	return int64(floatValue)
}
