package repo

import (
	"context"
	"fmt"

	commonclient "common/pkg/client"
	"common/pkg/constant"
	utilent "common/pkg/util/ent"

	"im/internal/biz/model"
	bizrepo "im/internal/biz/repo"
	"im/internal/data/gen"
	"im/internal/data/gen/globalchattopic"
	"im/internal/enum"
)

var _ bizrepo.GlobalChatTopicRepo = (*GlobalChatTopicRepo)(nil)

// GlobalChatTopicRepo 是全局聊天室话题的 Ent 仓储。
type GlobalChatTopicRepo struct {
	db  *gen.Client
	sql *commonclient.SQLClient
}

func NewGlobalChatTopicRepo(db *gen.Client, sqlClient *commonclient.SQLClient) bizrepo.GlobalChatTopicRepo {
	return &GlobalChatTopicRepo{db: db, sql: sqlClient}
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
		ID:         row.ID,
		UserID:     row.UserID,
		Content:    row.Content,
		PointsCost: row.PointsCost,
		Status:     enum.GlobalChatTopicStatus(row.Status),
		CreatedAt:  row.CreatedAt,
	}, nil
}

func (r *GlobalChatTopicRepo) Get(ctx context.Context, id int64) (*model.GlobalChatTopic, error) {
	row, err := r.client(ctx).GlobalChatTopic.Get(ctx, id)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &model.GlobalChatTopic{
		ID:         row.ID,
		UserID:     row.UserID,
		Content:    row.Content,
		PointsCost: row.PointsCost,
		Status:     enum.GlobalChatTopicStatus(row.Status),
		CreatedAt:  row.CreatedAt,
	}, nil
}

func (r *GlobalChatTopicRepo) NextID(ctx context.Context) (int64, error) {
	var id int64
	err := r.sql.DB.QueryRowContext(ctx, "SELECT nextval(pg_get_serial_sequence($1, 'id'))", constant.TablePrefixIM.String()+"global_chat_topics").Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("allocate global chat topic ID: %w", err)
	}
	return id, nil
}

func (r *GlobalChatTopicRepo) CreatePending(ctx context.Context, item *model.GlobalChatTopic) (*model.GlobalChatTopic, error) {
	row, err := r.client(ctx).GlobalChatTopic.Create().
		SetID(item.ID).
		SetUserID(item.UserID).
		SetContent(item.Content).
		SetPointsCost(item.PointsCost).
		SetStatus(globalchattopic.Status(enum.GlobalChatTopicStatusPending)).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.GlobalChatTopic{
		ID:         row.ID,
		UserID:     row.UserID,
		Content:    row.Content,
		PointsCost: row.PointsCost,
		Status:     enum.GlobalChatTopicStatus(row.Status),
		CreatedAt:  row.CreatedAt,
	}, nil
}

func (r *GlobalChatTopicRepo) Publish(ctx context.Context, id int64) (*model.GlobalChatTopic, error) {
	client := r.client(ctx)
	row, err := client.GlobalChatTopic.Get(ctx, id)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	if row.Status != globalchattopic.Status(enum.GlobalChatTopicStatusPending) {
		return &model.GlobalChatTopic{
			ID:         row.ID,
			UserID:     row.UserID,
			Content:    row.Content,
			PointsCost: row.PointsCost,
			Status:     enum.GlobalChatTopicStatus(row.Status),
			CreatedAt:  row.CreatedAt,
		}, nil
	}
	updated, err := client.GlobalChatTopic.Update().
		Where(
			globalchattopic.IDEQ(row.ID),
			globalchattopic.StatusEQ(globalchattopic.Status(enum.GlobalChatTopicStatusPending)),
		).
		SetStatus(globalchattopic.Status(enum.GlobalChatTopicStatusPublished)).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	if updated == 0 {
		return r.Get(ctx, id)
	}
	row.Status = globalchattopic.Status(enum.GlobalChatTopicStatusPublished)
	return &model.GlobalChatTopic{
		ID:         row.ID,
		UserID:     row.UserID,
		Content:    row.Content,
		PointsCost: row.PointsCost,
		Status:     enum.GlobalChatTopicStatus(row.Status),
		CreatedAt:  row.CreatedAt,
	}, nil
}

func (r *GlobalChatTopicRepo) Cancel(ctx context.Context, id int64) error {
	_, err := r.client(ctx).GlobalChatTopic.Update().
		Where(globalchattopic.IDEQ(id), globalchattopic.StatusEQ(globalchattopic.Status(enum.GlobalChatTopicStatusPending))).
		SetStatus(globalchattopic.Status(enum.GlobalChatTopicStatusCanceled)).
		Save(ctx)
	return err
}
