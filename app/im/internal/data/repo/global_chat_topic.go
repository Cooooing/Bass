package repo

import (
	"context"

	utilent "common/pkg/util/ent"

	"im/internal/biz/model"
	bizrepo "im/internal/biz/repo"
	"im/internal/data/gen"
	"im/internal/data/gen/globalchattopic"
	"im/internal/enum"
)

var _ bizrepo.GlobalChatTopicRepo = (*GlobalChatTopicRepo)(nil)

// GlobalChatTopicRepo 是全局聊天室话题的 Ent 仓储。
type GlobalChatTopicRepo struct{ db *gen.Client }

func NewGlobalChatTopicRepo(db *gen.Client) bizrepo.GlobalChatTopicRepo {
	return &GlobalChatTopicRepo{db: db}
}

func (r *GlobalChatTopicRepo) client(ctx context.Context) *gen.Client {
	if tx, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return tx
	}
	return r.db
}

func (r *GlobalChatTopicRepo) GetCurrent(ctx context.Context) (*model.GlobalChatTopic, error) {
	row, err := r.client(ctx).GlobalChatTopic.Query().
		Where(globalchattopic.StatusEQ(globalchattopic.Status(enum.GlobalChatTopicStatusPublished))).
		Order(gen.Desc(globalchattopic.FieldID)).
		First(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &model.GlobalChatTopic{
		ID:              row.ID,
		UserID:          row.UserID,
		ClientRequestID: row.ClientRequestID,
		Content:         row.Content,
		PointsCost:      row.PointsCost,
		Status:          enum.GlobalChatTopicStatus(row.Status),
		CreatedAt:       row.CreatedAt,
	}, nil
}

func (r *GlobalChatTopicRepo) GetByClientRequestID(ctx context.Context, clientRequestID string) (*model.GlobalChatTopic, error) {
	row, err := r.client(ctx).GlobalChatTopic.Query().Where(globalchattopic.ClientRequestIDEQ(clientRequestID)).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &model.GlobalChatTopic{
		ID:              row.ID,
		UserID:          row.UserID,
		ClientRequestID: row.ClientRequestID,
		Content:         row.Content,
		PointsCost:      row.PointsCost,
		Status:          enum.GlobalChatTopicStatus(row.Status),
		CreatedAt:       row.CreatedAt,
	}, nil
}

func (r *GlobalChatTopicRepo) CreatePending(ctx context.Context, item *model.GlobalChatTopic) (*model.GlobalChatTopic, error) {
	row, err := r.client(ctx).GlobalChatTopic.Create().
		SetUserID(item.UserID).
		SetClientRequestID(item.ClientRequestID).
		SetContent(item.Content).
		SetPointsCost(item.PointsCost).
		SetStatus(globalchattopic.Status(enum.GlobalChatTopicStatusPending)).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.GlobalChatTopic{
		ID:              row.ID,
		UserID:          row.UserID,
		ClientRequestID: row.ClientRequestID,
		Content:         row.Content,
		PointsCost:      row.PointsCost,
		Status:          enum.GlobalChatTopicStatus(row.Status),
		CreatedAt:       row.CreatedAt,
	}, nil
}

func (r *GlobalChatTopicRepo) Publish(ctx context.Context, clientRequestID string) (*model.GlobalChatTopic, error) {
	client := r.client(ctx)
	row, err := client.GlobalChatTopic.Query().Where(globalchattopic.ClientRequestIDEQ(clientRequestID)).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	if row.Status != globalchattopic.Status(enum.GlobalChatTopicStatusPending) {
		return &model.GlobalChatTopic{
			ID:              row.ID,
			UserID:          row.UserID,
			ClientRequestID: row.ClientRequestID,
			Content:         row.Content,
			PointsCost:      row.PointsCost,
			Status:          enum.GlobalChatTopicStatus(row.Status),
			CreatedAt:       row.CreatedAt,
		}, nil
	}
	row, err = client.GlobalChatTopic.UpdateOneID(row.ID).SetStatus(globalchattopic.Status(enum.GlobalChatTopicStatusPublished)).Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.GlobalChatTopic{
		ID:              row.ID,
		UserID:          row.UserID,
		ClientRequestID: row.ClientRequestID,
		Content:         row.Content,
		PointsCost:      row.PointsCost,
		Status:          enum.GlobalChatTopicStatus(row.Status),
		CreatedAt:       row.CreatedAt,
	}, nil
}

func (r *GlobalChatTopicRepo) Cancel(ctx context.Context, clientRequestID string) error {
	_, err := r.client(ctx).GlobalChatTopic.Update().
		Where(globalchattopic.ClientRequestIDEQ(clientRequestID), globalchattopic.StatusEQ(globalchattopic.Status(enum.GlobalChatTopicStatusPending))).
		SetStatus(globalchattopic.Status(enum.GlobalChatTopicStatusCanceled)).
		Save(ctx)
	return err
}
