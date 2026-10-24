package repo

import (
	"context"

	utilent "common/pkg/util/ent"

	"im/internal/biz/model"
	bizrepo "im/internal/biz/repo"
	"im/internal/data/gen"
	"im/internal/data/gen/globalchatbarrage"
	"im/internal/enum"
)

var _ bizrepo.GlobalChatBarrageRepo = (*GlobalChatBarrageRepo)(nil)

type GlobalChatBarrageRepo struct{ db *gen.Client }

func NewGlobalChatBarrageRepo(db *gen.Client) bizrepo.GlobalChatBarrageRepo {
	return &GlobalChatBarrageRepo{db: db}
}

func (r *GlobalChatBarrageRepo) client(ctx context.Context) *gen.Client {
	if tx, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return tx
	}
	return r.db
}

func (r *GlobalChatBarrageRepo) GetByClientRequestID(ctx context.Context, requestID string) (*model.GlobalChatBarrage, error) {
	row, err := r.client(ctx).GlobalChatBarrage.Query().Where(globalchatbarrage.ClientRequestIDEQ(requestID)).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return r.model(row), nil
}

func (r *GlobalChatBarrageRepo) CreatePending(ctx context.Context, item *model.GlobalChatBarrage) (*model.GlobalChatBarrage, error) {
	row, err := r.client(ctx).GlobalChatBarrage.Create().
		SetUserID(item.UserID).
		SetClientRequestID(item.ClientRequestID).
		SetContent(item.Content).
		SetPointsCost(item.PointsCost).
		SetStatus(globalchatbarrage.Status(enum.GlobalChatBarrageStatusPending)).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.model(row), nil
}

func (r *GlobalChatBarrageRepo) Publish(ctx context.Context, requestID string) (*model.GlobalChatBarrage, error) {
	client := r.client(ctx)
	row, err := client.GlobalChatBarrage.Query().Where(globalchatbarrage.ClientRequestIDEQ(requestID)).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	if row.Status != globalchatbarrage.Status(enum.GlobalChatBarrageStatusPending) {
		return r.model(row), nil
	}
	updated, err := client.GlobalChatBarrage.Update().
		Where(
			globalchatbarrage.IDEQ(row.ID),
			globalchatbarrage.StatusEQ(globalchatbarrage.Status(enum.GlobalChatBarrageStatusPending)),
		).
		SetStatus(globalchatbarrage.Status(enum.GlobalChatBarrageStatusPublished)).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	if updated == 0 {
		return r.GetByClientRequestID(ctx, requestID)
	}
	row.Status = globalchatbarrage.Status(enum.GlobalChatBarrageStatusPublished)
	return r.model(row), nil
}

func (r *GlobalChatBarrageRepo) Cancel(ctx context.Context, requestID string) error {
	_, err := r.client(ctx).GlobalChatBarrage.Update().
		Where(
			globalchatbarrage.ClientRequestIDEQ(requestID),
			globalchatbarrage.StatusEQ(globalchatbarrage.Status(enum.GlobalChatBarrageStatusPending)),
		).
		SetStatus(globalchatbarrage.Status(enum.GlobalChatBarrageStatusCanceled)).
		Save(ctx)
	return err
}

func (r *GlobalChatBarrageRepo) List(ctx context.Context, userID *int64, limit, offset int) ([]*model.GlobalChatBarrage, int64, error) {
	query := r.client(ctx).GlobalChatBarrage.Query().Where(globalchatbarrage.StatusEQ(globalchatbarrage.Status(enum.GlobalChatBarrageStatusPublished)))
	if userID != nil {
		query = query.Where(globalchatbarrage.UserIDEQ(*userID))
	}
	total, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	rows, err := query.Order(gen.Desc(globalchatbarrage.FieldID)).Limit(limit).Offset(offset).All(ctx)
	if err != nil {
		return nil, 0, err
	}
	result := make([]*model.GlobalChatBarrage, 0, len(rows))
	for _, row := range rows {
		result = append(result, r.model(row))
	}
	return result, int64(total), nil
}

func (*GlobalChatBarrageRepo) model(row *gen.GlobalChatBarrage) *model.GlobalChatBarrage {
	if row == nil {
		return nil
	}
	return &model.GlobalChatBarrage{
		ID:              row.ID,
		UserID:          row.UserID,
		ClientRequestID: row.ClientRequestID,
		Content:         row.Content,
		PointsCost:      row.PointsCost,
		Status:          enum.GlobalChatBarrageStatus(row.Status),
		CreatedAt:       row.CreatedAt,
	}
}
