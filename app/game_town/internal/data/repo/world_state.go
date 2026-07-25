package repo

import (
	"context"
	"fmt"
	"time"

	"common/pkg/apperror"
	utilent "common/pkg/util/ent"
	cerrors "common/proto/gen/common/errors"
	"game_town/internal/biz/model"
	bizrepo "game_town/internal/biz/repo"
	"game_town/internal/data/gen"
	"game_town/internal/data/gen/worldstate"

	"github.com/samber/lo"
)

var _ bizrepo.WorldStateRepo = (*WorldStateRepo)(nil)

type WorldStateRepo struct {
	db *gen.Client
}

func NewWorldStateRepo(
	db *gen.Client,
) bizrepo.WorldStateRepo {
	return &WorldStateRepo{
		db: db,
	}
}

func (r *WorldStateRepo) getClient(
	ctx context.Context,
) *gen.Client {
	if tx, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return tx
	}
	return r.db
}

func (r *WorldStateRepo) Save(
	ctx context.Context,
	row *model.WorldState,
) (*model.WorldState, error) {
	saved, err := r.getClient(ctx).WorldState.Create().
		SetWorldID(row.WorldID).
		SetVersion(row.Version).
		SetEventSequence(row.EventSequence).
		SetAgentCursor(row.AgentCursor).
		SetSummary(row.Summary).
		SetCurrentArc(row.CurrentArc).
		SetPublicChronicle(row.PublicChronicle).
		SetCurrentEra(row.CurrentEra).
		SetTimeScale(row.TimeScale).
		SetRuleVersion(row.RuleVersion).
		SetNillableNextDueAt(row.NextDueAt).
		SetNillableNextTickAt(row.NextTickAt).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.WorldState{
		ID:              saved.ID,
		WorldID:         saved.WorldID,
		Version:         saved.Version,
		EventSequence:   saved.EventSequence,
		AgentCursor:     saved.AgentCursor,
		Summary:         saved.Summary,
		CurrentArc:      saved.CurrentArc,
		PublicChronicle: saved.PublicChronicle,
		CurrentEra:      saved.CurrentEra,
		WorldTime:       saved.WorldTime,
		TimeAnchor:      saved.TimeAnchor,
		TimeScale:       saved.TimeScale,
		RuleVersion:     saved.RuleVersion,
		NextDueAt:       saved.NextDueAt,
		NextTickAt:      saved.NextTickAt,
		CreatedAt:       saved.CreatedAt,
		UpdatedAt:       saved.UpdatedAt,
	}, nil
}

func worldStateQuery(
	q *gen.WorldStateQuery,
	req *bizrepo.WorldStateQuery,
) *gen.WorldStateQuery {
	if req == nil {
		return q
	}
	if req.ID != nil {
		q = q.Where(worldstate.ID(*req.ID))
	}
	if len(req.IDs) > 0 {
		q = q.Where(worldstate.IDIn(req.IDs...))
	}
	if req.WorldID != nil {
		q = q.Where(worldstate.WorldID(*req.WorldID))
	}
	if req.TickDueBefore != nil {
		q = q.Where(worldstate.NextTickAtNotNil(), worldstate.NextTickAtLTE(*req.TickDueBefore))
	}
	if req.DueBefore != nil {
		q = q.Where(worldstate.NextDueAtNotNil(), worldstate.NextDueAtLTE(*req.DueBefore))
	}
	return q
}

func (r *WorldStateRepo) Get(
	ctx context.Context,
	req *bizrepo.WorldStateQuery,
) (*model.WorldState, error) {
	row, err := worldStateQuery(r.getClient(ctx).WorldState.Query(), req).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return &model.WorldState{
		ID:              row.ID,
		WorldID:         row.WorldID,
		Version:         row.Version,
		EventSequence:   row.EventSequence,
		AgentCursor:     row.AgentCursor,
		Summary:         row.Summary,
		CurrentArc:      row.CurrentArc,
		NextTickAt:      row.NextTickAt,
		CreatedAt:       row.CreatedAt,
		PublicChronicle: row.PublicChronicle,
		CurrentEra:      row.CurrentEra,
		WorldTime:       row.WorldTime,
		TimeAnchor:      row.TimeAnchor,
		TimeScale:       row.TimeScale,
		RuleVersion:     row.RuleVersion,
		NextDueAt:       row.NextDueAt,
		UpdatedAt:       row.UpdatedAt,
	}, nil
}

func (r *WorldStateRepo) List(
	ctx context.Context,
	req *bizrepo.WorldStateQuery,
) ([]*model.WorldState, error) {
	rows, err := worldStateQuery(r.getClient(ctx).WorldState.Query(), req).Order(worldstate.ByID()).All(ctx)
	if err != nil {
		return nil, err
	}
	out := lo.Map(rows, func(row *gen.WorldState, _ int) *model.WorldState {
		return &model.WorldState{
			ID:              row.ID,
			WorldID:         row.WorldID,
			Version:         row.Version,
			EventSequence:   row.EventSequence,
			AgentCursor:     row.AgentCursor,
			Summary:         row.Summary,
			CurrentArc:      row.CurrentArc,
			NextTickAt:      row.NextTickAt,
			CreatedAt:       row.CreatedAt,
			UpdatedAt:       row.UpdatedAt,
			PublicChronicle: row.PublicChronicle,
			CurrentEra:      row.CurrentEra,
			WorldTime:       row.WorldTime,
			TimeAnchor:      row.TimeAnchor,
			TimeScale:       row.TimeScale,
			RuleVersion:     row.RuleVersion,
			NextDueAt:       row.NextDueAt,
		}
	})
	return out, nil
}

func (r *WorldStateRepo) Map(
	ctx context.Context,
	req *bizrepo.WorldStateQuery,
) (map[int64]*model.WorldState, error) {
	rows, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	out := make(map[int64]*model.WorldState, len(rows))
	for _, row := range rows {
		out[row.ID] = row
	}
	return out, nil
}

func (r *WorldStateRepo) Count(
	ctx context.Context,
	req *bizrepo.WorldStateQuery,
) (int, error) {
	return worldStateQuery(r.getClient(ctx).WorldState.Query(), req).Count(ctx)
}

func (r *WorldStateRepo) Page(
	ctx context.Context,
	req *bizrepo.WorldStatePageReq,
) (*bizrepo.WorldStatePageResp, error) {
	p := page(req.Page)
	q := worldStateQuery(r.getClient(ctx).WorldState.Query(), &req.Query)
	total, err := q.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := q.Order(worldstate.ByID()).Offset(pageOffset(p)).Limit(pageLimit(p)).All(ctx)
	if err != nil {
		return nil, err
	}
	out := lo.Map(rows, func(row *gen.WorldState, _ int) *model.WorldState {
		return &model.WorldState{
			ID:              row.ID,
			WorldID:         row.WorldID,
			Version:         row.Version,
			EventSequence:   row.EventSequence,
			AgentCursor:     row.AgentCursor,
			Summary:         row.Summary,
			CurrentArc:      row.CurrentArc,
			NextTickAt:      row.NextTickAt,
			CreatedAt:       row.CreatedAt,
			UpdatedAt:       row.UpdatedAt,
			PublicChronicle: row.PublicChronicle,
			CurrentEra:      row.CurrentEra,
			WorldTime:       row.WorldTime,
			TimeAnchor:      row.TimeAnchor,
			TimeScale:       row.TimeScale,
			RuleVersion:     row.RuleVersion,
			NextDueAt:       row.NextDueAt,
		}
	})
	return &bizrepo.WorldStatePageResp{
		Rows: out,
		Page: basePage(total, p),
	}, nil
}

func (r *WorldStateRepo) NextEventSequence(
	ctx context.Context,
	worldID int64,
) (uint64, error) {
	count, err := r.getClient(ctx).WorldState.Update().Where(worldstate.WorldID(worldID)).AddEventSequence(1).Save(ctx)
	if err != nil {
		return 0, err
	}
	if count == 0 {
		return 0, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_NOT_FOUND)
	}
	row, err := r.Get(ctx, &bizrepo.WorldStateQuery{
		WorldID: new(worldID),
	})
	if err != nil {
		return 0, err
	}
	return row.EventSequence, nil
}

func (r *WorldStateRepo) AdvanceCursor(
	ctx context.Context,
	worldID int64,
	sequence uint64,
) error {
	_, err := r.getClient(ctx).WorldState.Update().Where(worldstate.WorldID(worldID), worldstate.AgentCursorLT(sequence)).SetAgentCursor(sequence).Save(ctx)
	return err
}

func (r *WorldStateRepo) UpdateNextTick(
	ctx context.Context,
	worldID int64,
	next time.Time,
) error {
	count, err := r.getClient(ctx).WorldState.Update().Where(worldstate.WorldID(worldID)).SetNextTickAt(next).SetNextDueAt(next).Save(ctx)
	if err != nil {
		return err
	}
	if count == 0 {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_NOT_FOUND)
	}
	return nil
}

func (r *WorldStateRepo) UpdateNarrative(
	ctx context.Context,
	req *bizrepo.WorldStateUpdateNarrativeReq,
) (*model.WorldState, error) {
	count, err := r.getClient(ctx).WorldState.Update().
		Where(worldstate.WorldID(req.WorldID), worldstate.Version(req.Version)).
		SetSummary(req.Summary).
		SetCurrentArc(req.CurrentArc).
		SetPublicChronicle(req.PublicChronicle).
		SetCurrentEra(req.CurrentEra).
		SetNillableNextTickAt(req.NextTickAt).
		SetNillableNextDueAt(req.NextDueAt).
		AddVersion(1).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, fmt.Errorf("world version conflict")
	}
	return r.Get(ctx, &bizrepo.WorldStateQuery{
		WorldID: new(req.WorldID),
	})
}
