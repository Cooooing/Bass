package repo

import (
	"context"
	"slices"

	"common/pkg/apperror"
	utilent "common/pkg/util/ent"
	cerrors "common/proto/gen/common/errors"
	"game_town/internal/biz/model"
	bizrepo "game_town/internal/biz/repo"
	"game_town/internal/data/gen"
	"game_town/internal/data/gen/event"
	"game_town/internal/enum"

	"entgo.io/ent/dialect/sql"
	"github.com/samber/lo"
)

var _ bizrepo.EventRepo = (*EventRepo)(nil)

type EventRepo struct {
	pageHelper
	db *gen.Client
}

func NewEventRepo(
	db *gen.Client,
) bizrepo.EventRepo {
	return &EventRepo{
		db: db,
	}
}

func (r *EventRepo) getClient(ctx context.Context) *gen.Client {
	if tx, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return tx
	}
	return r.db
}

func (r *EventRepo) Save(ctx context.Context, row *model.Event) (*model.Event, error) {
	saved, err := r.getClient(ctx).Event.Create().
		SetWorldID(row.WorldID).
		SetSequence(row.Sequence).
		SetType(event.Type(row.Type)).
		SetNillableActorPlayerID(row.ActorPlayerID).
		SetNillableNpcID(row.NpcID).
		SetNillableLocationID(row.LocationID).
		SetNillableCausationEventID(row.CausationEventID).
		SetSummary(row.Summary).
		SetContent(row.Content).
		SetPayload(row.Payload).
		SetWorldTime(row.WorldTime).
		SetOccurredAt(row.OccurredAt).
		SetCreatedAt(row.CreatedAt).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.Event{
		ID:               saved.ID,
		WorldID:          saved.WorldID,
		Sequence:         saved.Sequence,
		Type:             enum.EventType(saved.Type),
		ActorPlayerID:    saved.ActorPlayerID,
		NpcID:            saved.NpcID,
		LocationID:       saved.LocationID,
		CausationEventID: saved.CausationEventID,
		Summary:          saved.Summary,
		Content:          saved.Content,
		Payload:          saved.Payload,
		WorldTime:        saved.WorldTime,
		OccurredAt:       saved.OccurredAt,
		CreatedAt:        lo.FromPtr(saved.CreatedAt),
	}, nil
}

func (r *EventRepo) eventQuery(q *gen.EventQuery, req *bizrepo.EventQuery) *gen.EventQuery {
	if req == nil {
		return q
	}
	if req.ID != nil {
		q = q.Where(event.ID(*req.ID))
	}
	if len(req.IDs) > 0 {
		q = q.Where(event.IDIn(req.IDs...))
	}
	if req.WorldID != nil {
		q = q.Where(event.WorldID(*req.WorldID))
	}
	if req.AfterSequence != nil {
		q = q.Where(event.SequenceGT(*req.AfterSequence))
	}
	if req.Type != nil {
		q = q.Where(event.TypeEQ(event.Type(*req.Type)))
	}
	if req.ActorPlayerID != nil {
		q = q.Where(event.ActorPlayerID(*req.ActorPlayerID))
	}
	if req.NpcID != nil {
		q = q.Where(event.NpcID(*req.NpcID))
	}
	if req.CausationEventID != nil {
		q = q.Where(event.CausationEventID(*req.CausationEventID))
	}
	return q
}

func (r *EventRepo) Get(ctx context.Context, req *bizrepo.EventQuery) (*model.Event, error) {
	row, err := r.eventQuery(r.getClient(ctx).Event.Query(), req).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return &model.Event{
		ID:               row.ID,
		WorldID:          row.WorldID,
		Sequence:         row.Sequence,
		Type:             enum.EventType(row.Type),
		ActorPlayerID:    row.ActorPlayerID,
		NpcID:            row.NpcID,
		LocationID:       row.LocationID,
		CausationEventID: row.CausationEventID,
		Summary:          row.Summary,
		Content:          row.Content,
		Payload:          row.Payload,
		WorldTime:        row.WorldTime,
		OccurredAt:       row.OccurredAt,
		CreatedAt:        lo.FromPtr(row.CreatedAt),
	}, nil
}

func (r *EventRepo) List(ctx context.Context, req *bizrepo.EventQuery) ([]*model.Event, error) {
	query := r.eventQuery(r.getClient(ctx).Event.Query(), req)
	var rows []*gen.Event
	var err error
	if req != nil && req.RecentLimit > 0 {
		rows, err = query.Order(event.BySequence(sql.OrderDesc())).Limit(req.RecentLimit).All(ctx)
		slices.Reverse(rows)
	} else if req != nil && req.Limit > 0 {
		rows, err = query.Order(event.BySequence()).Limit(req.Limit).All(ctx)
	} else {
		rows, err = query.Order(event.BySequence()).All(ctx)
	}
	if err != nil {
		return nil, err
	}
	out := lo.Map(rows, func(row *gen.Event, _ int) *model.Event {
		return &model.Event{
			ID:               row.ID,
			WorldID:          row.WorldID,
			Sequence:         row.Sequence,
			Type:             enum.EventType(row.Type),
			ActorPlayerID:    row.ActorPlayerID,
			NpcID:            row.NpcID,
			LocationID:       row.LocationID,
			CausationEventID: row.CausationEventID,
			Summary:          row.Summary,
			Content:          row.Content,
			Payload:          row.Payload,
			WorldTime:        row.WorldTime,
			OccurredAt:       row.OccurredAt,
			CreatedAt:        lo.FromPtr(row.CreatedAt),
		}
	})
	return out, nil
}

func (r *EventRepo) Map(ctx context.Context, req *bizrepo.EventQuery) (map[int64]*model.Event, error) {
	rows, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	out := make(map[int64]*model.Event, len(rows))
	for _, row := range rows {
		out[row.ID] = row
	}
	return out, nil
}

func (r *EventRepo) Count(ctx context.Context, req *bizrepo.EventQuery) (int, error) {
	return r.eventQuery(r.getClient(ctx).Event.Query(), req).Count(ctx)
}

func (r *EventRepo) Page(ctx context.Context, req *bizrepo.EventPageReq) (*bizrepo.EventPageResp, error) {
	p := r.page(req.Page)
	q := r.eventQuery(r.getClient(ctx).Event.Query(), &req.Query)
	total, err := q.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := q.Order(event.BySequence()).Offset(r.pageOffset(p)).Limit(r.pageLimit(p)).All(ctx)
	if err != nil {
		return nil, err
	}
	out := lo.Map(rows, func(row *gen.Event, _ int) *model.Event {
		return &model.Event{
			ID:               row.ID,
			WorldID:          row.WorldID,
			Sequence:         row.Sequence,
			Type:             enum.EventType(row.Type),
			ActorPlayerID:    row.ActorPlayerID,
			NpcID:            row.NpcID,
			LocationID:       row.LocationID,
			CausationEventID: row.CausationEventID,
			Summary:          row.Summary,
			Content:          row.Content,
			Payload:          row.Payload,
			WorldTime:        row.WorldTime,
			OccurredAt:       row.OccurredAt,
			CreatedAt:        lo.FromPtr(row.CreatedAt),
		}
	})
	return &bizrepo.EventPageResp{
		Rows: out,
		Page: r.basePage(total, p),
	}, nil
}
