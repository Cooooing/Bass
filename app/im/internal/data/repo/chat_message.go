package repo

import (
	"context"
	"im/internal/biz/base"

	utilent "common/pkg/util/ent"
	"im/internal/biz/model"
	"im/internal/biz/repo"
	"im/internal/data/gen"
	"im/internal/data/gen/chatmessage"
	"im/internal/enum"
)

var _ repo.ChatMessageRepo = (*ChatMessageRepo)(nil)

type ChatMessageRepo struct {
	db *gen.Client
}

func NewChatMessageRepo(
	db *gen.Client,
) repo.ChatMessageRepo {
	return &ChatMessageRepo{
		db: db,
	}
}

func (r *ChatMessageRepo) getClient(
	ctx context.Context,
) *gen.Client {
	if tx, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return tx
	}
	return r.db
}

func (r *ChatMessageRepo) Save(
	ctx context.Context,
	chatMessage *model.ChatMessage,
) (*model.ChatMessage, error) {
	save, err := r.getClient(ctx).ChatMessage.Create().
		SetSenderID(chatMessage.SenderID).
		SetNillableReceiverID(chatMessage.ReceiverID).
		SetNillableGroupID(chatMessage.GroupID).
		SetNillableSessionID(chatMessage.SessionID).
		SetType(chatmessage.Type(chatMessage.Type)).
		SetContent(chatMessage.Content).
		SetStatus(chatmessage.Status(chatMessage.Status)).
		SetNillableCreatedBy(chatMessage.CreatedBy).
		SetNillableUpdatedBy(chatMessage.UpdatedBy).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.toModel(save), nil
}

func (r *ChatMessageRepo) UpdateStatus(
	ctx context.Context,
	req *repo.ChatMessageUpdateStatusReq,
) (*model.ChatMessage, error) {
	update, err := r.getClient(ctx).ChatMessage.UpdateOneID(req.ChatMessageID).
		SetStatus(chatmessage.Status(req.Status)).
		SetUpdatedBy(req.UpdatedBy).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.toModel(update), nil
}

func (r *ChatMessageRepo) Get(
	ctx context.Context,
	req *repo.ChatMessageQuery,
) (*model.ChatMessage, error) {
	query := r.getClient(ctx).ChatMessage.Query()
	query = r.getQuery(query, req)
	t, err := query.First(ctx)
	if err != nil {
		return nil, err
	}
	return r.toModel(t), nil
}

func (r *ChatMessageRepo) List(
	ctx context.Context,
	req *repo.ChatMessageQuery,
) ([]*model.ChatMessage, error) {
	query := r.getClient(ctx).ChatMessage.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.ChatMessage, 0, len(list))
	for _, item := range list {
		result = append(result, r.toModel(item))
	}
	return result, nil
}

func (r *ChatMessageRepo) Map(
	ctx context.Context,
	req *repo.ChatMessageQuery,
) (map[int64]*model.ChatMessage, error) {
	listResp, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*model.ChatMessage, len(listResp))
	for _, item := range listResp {
		result[item.ID] = item
	}
	return result, nil
}

func (r *ChatMessageRepo) Count(
	ctx context.Context,
	req *repo.ChatMessageQuery,
) (int, error) {
	query := r.getClient(ctx).ChatMessage.Query()
	query = r.getQuery(query, req)
	count, err := query.Count(ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *ChatMessageRepo) Page(
	ctx context.Context,
	req *repo.ChatMessageQuery,
) (*repo.ChatMessagePageResp, error) {
	page := normalizePage(nil)
	if req != nil {
		page = normalizePage(req.Page)
	}
	query := r.getClient(ctx).ChatMessage.Query()
	query = r.getQuery(query, req)
	countQuery := query.Clone()
	count, err := countQuery.Count(ctx)
	if err != nil {
		return nil, err
	}
	list, err := query.Limit(int(page.Size)).Offset(int((page.Page - 1) * page.Size)).All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.ChatMessage, 0, len(list))
	for _, item := range list {
		result = append(result, r.toModel(item))
	}
	return &repo.ChatMessagePageResp{
		Rows: result,
		Page: &base.PageResp{
			Total: int64(count),
			Size:  page.Size,
			Page:  page.Page,
		},
	}, nil
}

func (r *ChatMessageRepo) getQuery(
	query *gen.ChatMessageQuery,
	req *repo.ChatMessageQuery,
) *gen.ChatMessageQuery {
	if req == nil {
		return query
	}
	if req.IDs != nil {
		query = query.Where(chatmessage.IDIn(req.IDs...))
	}
	if req.SessionID != nil {
		query = query.Where(chatmessage.SessionID(*req.SessionID))
	}
	if req.SenderID != nil {
		query = query.Where(chatmessage.SenderID(*req.SenderID))
	}
	return query
}

func (r *ChatMessageRepo) toModel(
	t *gen.ChatMessage,
) *model.ChatMessage {
	return &model.ChatMessage{
		ID:         t.ID,
		SenderID:   t.SenderID,
		ReceiverID: t.ReceiverID,
		GroupID:    t.GroupID,
		SessionID:  t.SessionID,
		Type:       enum.MessageType(t.Type),
		Content:    t.Content,
		Status:     enum.MessageStatus(t.Status),
		CreatedAt:  t.CreatedAt,
		UpdatedAt:  t.UpdatedAt,
		CreatedBy:  t.CreatedBy,
		UpdatedBy:  t.UpdatedBy,
	}
}
