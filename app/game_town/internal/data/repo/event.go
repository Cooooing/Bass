package repo

import (
	"common/pkg/server"
	"common/proto/gen/common"
	"context"
	"game_town/internal/biz/model"
	bizrepo "game_town/internal/biz/repo"
	"game_town/internal/data/gen"
	"game_town/internal/data/gen/event"
	"time"
)

type EventRepo struct{ *baseRepo }

func NewEventRepo(db *gen.Client) bizrepo.EventRepo {
	return &EventRepo{baseRepo: &baseRepo{db: db}}
}

func (r *EventRepo) CreateEvent(ctx context.Context, req *bizrepo.CreateEventReq) (*bizrepo.CreateEventResponse, error) {
	now := time.Now()
	row := req.Row
	if row.OccurredAt.IsZero() {
		row.OccurredAt = now
	}
	created, err := r.db.Event.Create().SetWorldID(row.WorldID).SetType(row.Type).SetNillableActorPlayerID(row.ActorPlayerID).SetNillableTargetNpcID(row.TargetNpcID).SetNillableLocationID(row.LocationID).SetNillableCommandID(row.CommandID).SetSummary(row.Summary).SetContent(row.Content).SetEffects(row.Effects).SetMetadata(row.Metadata).SetOccurredAt(row.OccurredAt).SetCreatedAt(now).Save(ctx)
	if err != nil {
		return nil, err
	}
	return &bizrepo.CreateEventResponse{Row: r.event(created)}, nil
}

func (r *EventRepo) Page(ctx context.Context, req *bizrepo.EventPageReq) (*bizrepo.EventPageResponse, error) {
	pageReq := server.PageValid(req.Page)
	queryReq := req.Query
	query := r.db.Event.Query().Where(event.WorldID(queryReq.WorldID))
	if queryReq.ActorPlayerID != nil {
		query = query.Where(event.ActorPlayerID(*queryReq.ActorPlayerID))
	}
	if queryReq.TargetNpcID != nil {
		query = query.Where(event.TargetNpcID(*queryReq.TargetNpcID))
	}
	if queryReq.Type != nil {
		query = query.Where(event.Type(*queryReq.Type))
	}
	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := query.Order(gen.Desc(event.FieldOccurredAt)).Limit(int(pageReq.Size)).Offset(int((pageReq.Page - 1) * pageReq.Size)).All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.Event, 0, len(rows))
	for _, row := range rows {
		result = append(result, r.event(row))
	}
	return &bizrepo.EventPageResponse{Rows: result, Page: &common.PageResponse{Total: uint32(total), Page: pageReq.Page, Size: pageReq.Size}}, nil
}

func (r *EventRepo) ListRecentEvents(ctx context.Context, req *bizrepo.ListRecentEventsReq) (*bizrepo.ListRecentEventsResponse, error) {
	limit := req.Limit
	if limit <= 0 {
		limit = 20
	}
	rows, err := r.db.Event.Query().Where(event.WorldID(req.WorldID)).Order(gen.Desc(event.FieldOccurredAt)).Limit(limit).All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.Event, 0, len(rows))
	for _, row := range rows {
		result = append(result, r.event(row))
	}
	return &bizrepo.ListRecentEventsResponse{Rows: result}, nil
}
