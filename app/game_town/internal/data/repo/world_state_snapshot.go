package repo

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	"context"
	"game_town/internal/biz/model"
	bizrepo "game_town/internal/biz/repo"
	"game_town/internal/data/gen"
	"game_town/internal/data/gen/worldstatesnapshot"
	"time"
)

type WorldStateSnapshotRepo struct{ *baseRepo }

func NewWorldStateSnapshotRepo(db *gen.Client) bizrepo.WorldStateSnapshotRepo {
	return &WorldStateSnapshotRepo{baseRepo: &baseRepo{db: db}}
}

func (r *WorldStateSnapshotRepo) GetLatestWorldState(ctx context.Context, worldID int64) (*model.WorldStateSnapshot, error) {
	row, err := r.db.WorldStateSnapshot.Query().Where(worldstatesnapshot.WorldID(worldID)).Order(gen.Desc(worldstatesnapshot.FieldTickCount)).First(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_TOWN_WORLD_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return r.state(row), nil
}

func (r *WorldStateSnapshotRepo) CreateState(ctx context.Context, row *model.WorldStateSnapshot) (*model.WorldStateSnapshot, error) {
	now := time.Now()
	created, err := r.db.WorldStateSnapshot.Create().SetWorldID(row.WorldID).SetTickCount(row.TickCount).SetCurrentArc(row.CurrentArc).SetMetrics(row.Metrics).SetSummary(row.Summary).SetNillableReasonEventID(row.ReasonEventID).SetCreatedAt(now).Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.state(created), nil
}
