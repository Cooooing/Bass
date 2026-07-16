package repo

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	"context"
	bizrepo "game_town/internal/biz/repo"
	"game_town/internal/data/gen"
	"game_town/internal/data/gen/worldstatesnapshot"
	"time"
)

type WorldStateSnapshotRepo struct{ *baseRepo }

func NewWorldStateSnapshotRepo(db *gen.Client) bizrepo.WorldStateSnapshotRepo {
	return &WorldStateSnapshotRepo{baseRepo: &baseRepo{db: db}}
}

func (r *WorldStateSnapshotRepo) GetLatestWorldState(ctx context.Context, req *bizrepo.GetLatestWorldStateReq) (*bizrepo.GetLatestWorldStateResponse, error) {
	row, err := r.db.WorldStateSnapshot.Query().Where(worldstatesnapshot.WorldID(req.WorldID)).Order(gen.Desc(worldstatesnapshot.FieldTickCount)).First(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_TOWN_WORLD_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return &bizrepo.GetLatestWorldStateResponse{Row: r.state(row)}, nil
}

func (r *WorldStateSnapshotRepo) CreateState(ctx context.Context, req *bizrepo.CreateStateReq) (*bizrepo.CreateStateResponse, error) {
	now := time.Now()
	row := req.Row
	created, err := r.db.WorldStateSnapshot.Create().SetWorldID(row.WorldID).SetTickCount(row.TickCount).SetCurrentArc(row.CurrentArc).SetMetrics(row.Metrics).SetSummary(row.Summary).SetNillableReasonEventID(row.ReasonEventID).SetCreatedAt(now).Save(ctx)
	if err != nil {
		return nil, err
	}
	return &bizrepo.CreateStateResponse{Row: r.state(created)}, nil
}
