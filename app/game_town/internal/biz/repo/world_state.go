package repo

import (
	"context"
	"time"

	"game_town/internal/biz/base"
	"game_town/internal/biz/model"
)

type WorldStateRepo interface {
	Save(context.Context, *model.WorldState) (*model.WorldState, error)
	NextEventSequence(context.Context, int64) (uint64, error)
	AdvanceCursor(context.Context, int64, uint64) error
	UpdateNarrative(context.Context, *WorldStateUpdateNarrativeReq) (*model.WorldState, error)
	UpdateNextTick(context.Context, int64, time.Time) error
	UpdateNextDue(context.Context, int64, *time.Time) error
	AdvanceTime(context.Context, *WorldStateAdvanceTimeReq) (*model.WorldState, error)
	Get(context.Context, *WorldStateQuery) (*model.WorldState, error)
	List(context.Context, *WorldStateQuery) ([]*model.WorldState, error)
	Map(context.Context, *WorldStateQuery) (map[int64]*model.WorldState, error)
	Count(context.Context, *WorldStateQuery) (int, error)
	Page(context.Context, *WorldStatePageReq) (*WorldStatePageResp, error)
}

type WorldStateUpdateNarrativeReq struct {
	WorldID, Version                                 int64
	Summary, CurrentArc, PublicChronicle, CurrentEra string
	NextTickAt, NextDueAt                            *time.Time
}

type WorldStateAdvanceTimeReq struct {
	WorldID, Version      int64
	WorldTime, TimeAnchor time.Time
	NextDueAt             *time.Time
}

type WorldStateQuery struct {
	ID                       *int64
	IDs                      []int64
	WorldID                  *int64
	TickDueBefore, DueBefore *time.Time
}

type WorldStatePageReq struct {
	Page  base.PageRequest
	Query WorldStateQuery
}
type WorldStatePageResp struct {
	Rows []*model.WorldState
	Page base.PageResp
}
