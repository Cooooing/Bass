package repo

import (
	"context"

	"common/pkg/apperror"
	utilent "common/pkg/util/ent"
	cerrors "common/proto/gen/common/errors"
	"game_town/internal/biz/model"
	bizrepo "game_town/internal/biz/repo"
	"game_town/internal/data/gen"
	"game_town/internal/data/gen/event"
	"game_town/internal/data/gen/observation"
	"game_town/internal/enum"

	"entgo.io/ent/dialect/sql"
	"github.com/samber/lo"
)

var _ bizrepo.ObservationRepo = (*ObservationRepo)(nil)

type ObservationRepo struct {
	pageHelper
	db *gen.Client
}

func NewObservationRepo(
	db *gen.Client,
) bizrepo.ObservationRepo {
	return &ObservationRepo{
		db: db,
	}
}

func (r *ObservationRepo) getClient(ctx context.Context) *gen.Client {
	if tx, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return tx
	}
	return r.db
}

func (r *ObservationRepo) Save(ctx context.Context, row *model.Observation) (*model.Observation, error) {
	saved, err := r.getClient(ctx).Observation.Create().
		SetWorldID(row.WorldID).
		SetEventID(row.EventID).
		SetNillableNpcID(row.NpcID).
		SetNillablePlayerID(row.PlayerID).
		SetSource(observation.Source(row.Source)).
		SetCertainty(observation.Certainty(row.Certainty)).
		SetSummary(row.Summary).
		SetSalience(row.Salience).
		SetObservedAt(row.ObservedAt).
		SetWorldTime(row.WorldTime).
		SetCreatedAt(row.ObservedAt).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.observation(saved), nil
}

func (r *ObservationRepo) observationQuery(q *gen.ObservationQuery, req *bizrepo.ObservationQuery) *gen.ObservationQuery {
	if req == nil {
		return q
	}
	if req.ID != nil {
		q = q.Where(observation.ID(*req.ID))
	}
	if req.WorldID != nil {
		q = q.Where(observation.WorldID(*req.WorldID))
	}
	if req.EventID != nil {
		q = q.Where(observation.EventID(*req.EventID))
	}
	if len(req.EventIDs) > 0 {
		q = q.Where(observation.EventIDIn(req.EventIDs...))
	}
	if req.NpcID != nil {
		q = q.Where(observation.NpcID(*req.NpcID))
	}
	if req.PlayerID != nil {
		q = q.Where(observation.PlayerID(*req.PlayerID))
	}
	if req.AfterEventID != nil {
		q = q.Where(observation.EventIDGT(*req.AfterEventID))
	}
	if req.AfterWorldTime != nil {
		q = q.Where(observation.WorldTimeGT(*req.AfterWorldTime))
	}
	if req.AfterEventSequence != nil || req.EventType != nil {
		q.Modify(func(selector *sql.Selector) {
			events := sql.Table(event.Table).As("visible_events")
			selector.Join(events).On(selector.C(observation.FieldEventID), events.C(event.FieldID))
			if req.AfterEventSequence != nil {
				selector.Where(sql.GT(events.C(event.FieldSequence), *req.AfterEventSequence))
			}
			if req.EventType != nil {
				selector.Where(sql.EQ(events.C(event.FieldType), string(*req.EventType)))
			}
		})
	}
	return q
}

func (r *ObservationRepo) Get(ctx context.Context, req *bizrepo.ObservationQuery) (*model.Observation, error) {
	row, err := r.observationQuery(r.getClient(ctx).Observation.Query(), req).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return r.observation(row), nil
}

func (r *ObservationRepo) List(ctx context.Context, req *bizrepo.ObservationQuery) ([]*model.Observation, error) {
	rows, err := r.observationQuery(r.getClient(ctx).Observation.Query(), req).
		Order(observation.ByEventID()).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return lo.Map(rows, func(row *gen.Observation, _ int) *model.Observation {
		return r.observation(row)
	}), nil
}

func (r *ObservationRepo) Map(ctx context.Context, req *bizrepo.ObservationQuery) (map[int64]*model.Observation, error) {
	rows, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	out := make(map[int64]*model.Observation, len(rows))
	for _, row := range rows {
		out[row.ID] = row
	}
	return out, nil
}

func (r *ObservationRepo) Count(ctx context.Context, req *bizrepo.ObservationQuery) (int, error) {
	return r.observationQuery(r.getClient(ctx).Observation.Query(), req).Count(ctx)
}

func (r *ObservationRepo) Page(ctx context.Context, req *bizrepo.ObservationPageReq) (*bizrepo.ObservationPageResp, error) {
	p := r.page(req.Page)
	q := r.observationQuery(r.getClient(ctx).Observation.Query(), &req.Query)
	total := 0
	if !req.SkipTotal {
		var err error
		total, err = q.Clone().Count(ctx)
		if err != nil {
			return nil, err
		}
	}
	q = r.orderObservations(q, &req.Query)
	rows, err := q.
		Offset(r.pageOffset(p)).
		Limit(r.pageLimit(p)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return &bizrepo.ObservationPageResp{
		Rows: lo.Map(rows, func(row *gen.Observation, _ int) *model.Observation {
			return r.observation(row)
		}),
		Page: r.basePage(total, p),
	}, nil
}

func (r *ObservationRepo) orderObservations(q *gen.ObservationQuery, req *bizrepo.ObservationQuery) *gen.ObservationQuery {
	if req != nil && (req.AfterEventSequence != nil || req.EventType != nil) {
		q.Modify(func(selector *sql.Selector) {
			events := sql.Table(event.Table).As("visible_events")
			selector.OrderBy(events.C(event.FieldSequence))
		})
		return q
	}
	return q.Order(observation.ByEventID())
}

func (r *ObservationRepo) observation(row *gen.Observation) *model.Observation {
	return &model.Observation{
		ID:         row.ID,
		WorldID:    row.WorldID,
		EventID:    row.EventID,
		NpcID:      row.NpcID,
		PlayerID:   row.PlayerID,
		Source:     enum.ObservationSource(row.Source),
		Certainty:  enum.KnowledgeCertainty(row.Certainty),
		Summary:    row.Summary,
		Salience:   row.Salience,
		ObservedAt: row.ObservedAt,
		WorldTime:  row.WorldTime,
		CreatedAt:  row.CreatedAt,
	}
}
