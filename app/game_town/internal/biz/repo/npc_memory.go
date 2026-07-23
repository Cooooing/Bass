package repo

import (
	"context"
	"time"

	"game_town/internal/biz/base"
	"game_town/internal/biz/model"
	"game_town/internal/enum"
)

type NpcMemoryRepo interface {
	Save(context.Context, *model.NpcMemory) (*model.NpcMemory, error)
	SetEmbedding(context.Context, *NpcMemoryEmbeddingReq) error
	Search(context.Context, *NpcMemorySearchReq) ([]*model.NpcMemory, error)
	Get(context.Context, *NpcMemoryQuery) (*model.NpcMemory, error)
	List(context.Context, *NpcMemoryQuery) ([]*model.NpcMemory, error)
	Map(context.Context, *NpcMemoryQuery) (map[int64]*model.NpcMemory, error)
	Count(context.Context, *NpcMemoryQuery) (int, error)
	Page(context.Context, *NpcMemoryPageReq) (*NpcMemoryPageResp, error)
}

type NpcMemoryEmbeddingReq struct {
	ID           int64
	Vector       []float32
	Model        string
	Status       enum.EmbeddingStatus
	ErrorSummary string
}
type NpcMemorySearchReq struct {
	WorldID, NpcID              int64
	Vector                      []float32
	CandidateLimit, ResultLimit int
	Now                         time.Time
}
type NpcMemoryQuery struct {
	ID, WorldID, NpcID, SourceEventID *int64
	Status                            *enum.EmbeddingStatus
	RecentLimit                       int
}
type NpcMemoryPageReq struct {
	Page  base.PageRequest
	Query NpcMemoryQuery
}
type NpcMemoryPageResp struct {
	Rows []*model.NpcMemory
	Page base.PageResp
}
