package repo

import (
	"context"
	"math"
	"slices"

	"common/pkg/apperror"
	utilent "common/pkg/util/ent"
	cerrors "common/proto/gen/common/errors"
	"game_town/internal/biz/model"
	bizrepo "game_town/internal/biz/repo"
	"game_town/internal/data/gen"
	"game_town/internal/data/gen/npcmemory"
	"game_town/internal/enum"

	"entgo.io/ent/dialect/sql"
	"github.com/pgvector/pgvector-go"
	"github.com/samber/lo"
)

var _ bizrepo.NpcMemoryRepo = (*NpcMemoryRepo)(nil)

type NpcMemoryRepo struct {
	pageHelper
	db *gen.Client
}

func NewNpcMemoryRepo(
	db *gen.Client,
) bizrepo.NpcMemoryRepo {
	return &NpcMemoryRepo{
		db: db,
	}
}

func (r *NpcMemoryRepo) getClient(ctx context.Context) *gen.Client {
	if tx, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return tx
	}
	return r.db
}

func (r *NpcMemoryRepo) Save(ctx context.Context, row *model.NpcMemory) (*model.NpcMemory, error) {
	saved, err := r.getClient(ctx).NpcMemory.Create().
		SetWorldID(row.WorldID).
		SetNpcID(row.NpcID).
		SetNillableSourceEventID(row.SourceEventID).
		SetNillableSourceObservationID(row.SourceObservationID).
		SetKind(npcmemory.Kind(row.Kind)).
		SetContent(row.Content).
		SetImportance(row.Importance).
		SetOccurredWorldTime(row.OccurredWorldTime).
		SetNillableLastRecalledAt(row.LastRecalledAt).
		SetEmbeddingModel(row.EmbeddingModel).
		SetEmbeddingStatus(npcmemory.EmbeddingStatus(row.EmbeddingStatus)).
		SetEmbeddingError(row.EmbeddingError).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.npcMemory(saved), nil
}

func (r *NpcMemoryRepo) SetEmbedding(ctx context.Context, req *bizrepo.NpcMemoryEmbeddingReq) error {
	update := r.getClient(ctx).NpcMemory.UpdateOneID(req.ID).
		SetEmbeddingModel(req.Model).
		SetEmbeddingStatus(npcmemory.EmbeddingStatus(req.Status)).
		SetEmbeddingError(req.ErrorSummary)
	if len(req.Vector) > 0 {
		update.SetEmbedding(pgvector.NewVector(req.Vector))
	}
	return update.Exec(ctx)
}

func (r *NpcMemoryRepo) npcMemoryQuery(q *gen.NpcMemoryQuery, req *bizrepo.NpcMemoryQuery) *gen.NpcMemoryQuery {
	if req == nil {
		return q
	}
	if req.ID != nil {
		q = q.Where(npcmemory.ID(*req.ID))
	}
	if req.WorldID != nil {
		q = q.Where(npcmemory.WorldID(*req.WorldID))
	}
	if req.NpcID != nil {
		q = q.Where(npcmemory.NpcID(*req.NpcID))
	}
	if req.SourceEventID != nil {
		q = q.Where(npcmemory.SourceEventID(*req.SourceEventID))
	}
	if req.Status != nil {
		q = q.Where(npcmemory.EmbeddingStatusEQ(npcmemory.EmbeddingStatus(*req.Status)))
	}
	return q
}

func (r *NpcMemoryRepo) Get(ctx context.Context, req *bizrepo.NpcMemoryQuery) (*model.NpcMemory, error) {
	row, err := r.npcMemoryQuery(r.getClient(ctx).NpcMemory.Query(), req).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return r.npcMemory(row), nil
}

func (r *NpcMemoryRepo) List(ctx context.Context, req *bizrepo.NpcMemoryQuery) ([]*model.NpcMemory, error) {
	q := r.npcMemoryQuery(r.getClient(ctx).NpcMemory.Query(), req)
	if req != nil && req.RecentLimit > 0 {
		q = q.Order(npcmemory.ByOccurredWorldTime(sql.OrderDesc())).Limit(req.RecentLimit)
	} else {
		q = q.Order(npcmemory.ByOccurredWorldTime())
	}

	rows, err := q.All(ctx)
	if err != nil {
		return nil, err
	}
	if req != nil && req.RecentLimit > 0 {
		slices.Reverse(rows)
	}
	return lo.Map(rows, func(row *gen.NpcMemory, _ int) *model.NpcMemory {
		return r.npcMemory(row)
	}), nil
}

func (r *NpcMemoryRepo) Search(ctx context.Context, req *bizrepo.NpcMemorySearchReq) ([]*model.NpcMemory, error) {
	if req.WorldID <= 0 || req.NpcID <= 0 || len(req.Vector) == 0 {
		return nil, nil
	}
	limit := req.CandidateLimit
	if limit <= 0 {
		limit = 12
	}
	resultLimit := req.ResultLimit
	if resultLimit <= 0 {
		resultLimit = 5
	}

	rows, err := r.searchVectorCandidates(ctx, req, limit)
	if err != nil {
		status := enum.EmbeddingStatusReady
		fallbackLimit := resultLimit
		if fallbackLimit <= 0 {
			fallbackLimit = 5
		}
		return r.List(ctx, &bizrepo.NpcMemoryQuery{
			WorldID:     new(req.WorldID),
			NpcID:       new(req.NpcID),
			Status:      new(status),
			RecentLimit: fallbackLimit,
		})
	}

	values := make([]scoredMemory, 0, len(rows))
	for index, row := range rows {
		ageDays := math.Max(0, req.Now.Sub(row.OccurredWorldTime).Hours()/24)
		recency := 1 / (1 + ageDays/30)
		rank := 1 / float64(index+1)
		values = append(values, scoredMemory{
			row:   row,
			score: 0.55*rank + 0.25*row.Importance + 0.20*recency,
		})
	}
	slices.SortFunc(values, func(a scoredMemory, b scoredMemory) int {
		if a.score > b.score {
			return -1
		}
		if a.score < b.score {
			return 1
		}
		return 0
	})
	if len(values) > resultLimit {
		values = values[:resultLimit]
	}

	out := make([]*model.NpcMemory, 0, len(values))
	for _, value := range values {
		out = append(out, r.npcMemory(value.row))
	}
	return out, nil
}

func (r *NpcMemoryRepo) searchVectorCandidates(ctx context.Context, req *bizrepo.NpcMemorySearchReq, limit int) ([]*gen.NpcMemory, error) {
	return r.getClient(ctx).NpcMemory.Query().
		Where(
			npcmemory.WorldID(req.WorldID),
			npcmemory.NpcID(req.NpcID),
			npcmemory.EmbeddingStatusEQ(npcmemory.EmbeddingStatus(enum.EmbeddingStatusReady)),
			npcmemory.EmbeddingNotNil(),
		).
		Modify(func(selector *sql.Selector) {
			vector := pgvector.NewVector(req.Vector)
			selector.OrderExpr(sql.ExprFunc(func(builder *sql.Builder) {
				builder.Ident(npcmemory.FieldEmbedding).WriteString(" <=> ").Arg(vector)
			}))
		}).
		Limit(limit).
		All(ctx)
}

type scoredMemory struct {
	row   *gen.NpcMemory
	score float64
}

func (r *NpcMemoryRepo) Map(ctx context.Context, req *bizrepo.NpcMemoryQuery) (map[int64]*model.NpcMemory, error) {
	rows, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	out := make(map[int64]*model.NpcMemory, len(rows))
	for _, row := range rows {
		out[row.ID] = row
	}
	return out, nil
}

func (r *NpcMemoryRepo) Count(ctx context.Context, req *bizrepo.NpcMemoryQuery) (int, error) {
	return r.npcMemoryQuery(r.getClient(ctx).NpcMemory.Query(), req).Count(ctx)
}

func (r *NpcMemoryRepo) Page(ctx context.Context, req *bizrepo.NpcMemoryPageReq) (*bizrepo.NpcMemoryPageResp, error) {
	p := r.page(req.Page)
	q := r.npcMemoryQuery(r.getClient(ctx).NpcMemory.Query(), &req.Query)
	total, err := q.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := q.Order(npcmemory.ByOccurredWorldTime(sql.OrderDesc())).
		Offset(r.pageOffset(p)).
		Limit(r.pageLimit(p)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return &bizrepo.NpcMemoryPageResp{
		Rows: lo.Map(rows, func(row *gen.NpcMemory, _ int) *model.NpcMemory {
			return r.npcMemory(row)
		}),
		Page: r.basePage(total, p),
	}, nil
}

func (r *NpcMemoryRepo) npcMemory(row *gen.NpcMemory) *model.NpcMemory {
	return &model.NpcMemory{
		ID:                  row.ID,
		WorldID:             row.WorldID,
		NpcID:               row.NpcID,
		SourceEventID:       row.SourceEventID,
		SourceObservationID: row.SourceObservationID,
		Kind:                enum.MemoryKind(row.Kind),
		Content:             row.Content,
		Importance:          row.Importance,
		OccurredWorldTime:   row.OccurredWorldTime,
		LastRecalledAt:      row.LastRecalledAt,
		EmbeddingModel:      row.EmbeddingModel,
		EmbeddingStatus:     enum.EmbeddingStatus(row.EmbeddingStatus),
		EmbeddingError:      row.EmbeddingError,
		CreatedAt:           row.CreatedAt,
		UpdatedAt:           row.UpdatedAt,
	}
}
