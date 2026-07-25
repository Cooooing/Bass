package repo

import (
	"context"
	"fmt"
	"time"

	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	"game_town/internal/biz/model"
	bizrepo "game_town/internal/biz/repo"
	"game_town/internal/data/gen/worldstate"
)

func (r *WorldStateRepo) UpdateNextDue(
	ctx context.Context,
	worldID int64,
	nextDueAt *time.Time,
) error {
	update := r.getClient(ctx).WorldState.Update().Where(worldstate.WorldID(worldID))
	if nextDueAt == nil {
		update.ClearNextDueAt()
	} else {
		update.SetNextDueAt(*nextDueAt)
	}
	count, err := update.Save(ctx)
	if err != nil {
		return err
	}
	if count == 0 {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_NOT_FOUND)
	}
	return nil
}

func (r *WorldStateRepo) AdvanceTime(
	ctx context.Context,
	req *bizrepo.WorldStateAdvanceTimeReq,
) (*model.WorldState, error) {
	update := r.getClient(ctx).WorldState.Update().Where(worldstate.WorldID(req.WorldID), worldstate.Version(req.Version)).
		SetWorldTime(req.WorldTime).
		SetTimeAnchor(req.TimeAnchor).
		AddVersion(1)
	if req.NextDueAt == nil {
		update.ClearNextDueAt()
	} else {
		update.SetNextDueAt(*req.NextDueAt)
	}
	count, err := update.Save(ctx)
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
