package repo

import (
	cerrors "common/proto/gen/common/errors"
	"context"
	"im/internal/biz/base"

	"common/pkg/apperror"
	utilent "common/pkg/util/ent"
	"im/internal/biz/model"
	"im/internal/biz/repo"
	"im/internal/data/gen"
	"im/internal/data/gen/chatsession"
	"im/internal/enum"
)

var _ repo.ChatSessionRepo = (*ChatSessionRepo)(nil)

type ChatSessionRepo struct {
	pageNormalizer
	db *gen.Client
}

func NewChatSessionRepo(
	db *gen.Client,
) repo.ChatSessionRepo {
	return &ChatSessionRepo{
		db: db,
	}
}

func (r *ChatSessionRepo) getClient(ctx context.Context) *gen.Client {
	if tx, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return tx
	}
	return r.db
}

func (r *ChatSessionRepo) Save(ctx context.Context, chatSession *model.ChatSession) (*model.ChatSession, error) {
	if (chatSession.ReceiverID == nil && chatSession.GroupID == nil) || (chatSession.ReceiverID != nil && chatSession.GroupID != nil) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_IM_CHAT_SESSION_INVALID)
	}
	save, err := r.getClient(ctx).ChatSession.Create().
		SetNillableReceiverID(chatSession.ReceiverID).
		SetNillableGroupID(chatSession.GroupID).
		SetNillableCreatedBy(chatSession.CreatedBy).
		SetNillableUpdatedBy(chatSession.UpdatedBy).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.toModel(save), nil
}

func (r *ChatSessionRepo) UpdateLastReadMessage(ctx context.Context, req *repo.ChatSessionUpdateLastReadMessageReq) (*model.ChatSession, error) {
	update, err := r.getClient(ctx).ChatSession.UpdateOneID(req.ChatSessionID).
		SetLastReadMessageID(req.MessageID).
		AddReadCount(req.OperationReadCount).
		SetUpdatedBy(req.UpdatedBy).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.toModel(update), nil
}

func (r *ChatSessionRepo) UpdateMuted(ctx context.Context, req *repo.ChatSessionUpdateMutedReq) (*model.ChatSession, error) {
	update, err := r.getClient(ctx).ChatSession.UpdateOneID(req.ChatSessionID).
		SetIsMuted(req.Muted).
		SetUpdatedBy(req.UpdatedBy).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.toModel(update), nil
}

func (r *ChatSessionRepo) UpdatePinned(ctx context.Context, req *repo.ChatSessionUpdatePinnedReq) (*model.ChatSession, error) {
	update, err := r.getClient(ctx).ChatSession.UpdateOneID(req.ChatSessionID).
		SetIsPinned(req.Pinned).
		SetUpdatedBy(req.UpdatedBy).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.toModel(update), nil
}

func (r *ChatSessionRepo) Get(ctx context.Context, req *repo.ChatSessionQuery) (*model.ChatSession, error) {
	query := r.getClient(ctx).ChatSession.Query()
	query = r.getQuery(query, req)
	t, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_IM_CHAT_SESSION_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return r.toModel(t), nil
}

func (r *ChatSessionRepo) List(ctx context.Context, req *repo.ChatSessionQuery) ([]*model.ChatSession, error) {
	query := r.getClient(ctx).ChatSession.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	chatSessions := make([]*model.ChatSession, 0, len(list))
	for _, item := range list {
		chatSessions = append(chatSessions, r.toModel(item))
	}
	return chatSessions, nil
}

func (r *ChatSessionRepo) Map(ctx context.Context, req *repo.ChatSessionQuery) (map[int64]*model.ChatSession, error) {
	listResp, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*model.ChatSession, len(listResp))
	for _, item := range listResp {
		result[item.ID] = item
	}
	return result, nil
}

func (r *ChatSessionRepo) Count(ctx context.Context, req *repo.ChatSessionQuery) (int, error) {
	query := r.getClient(ctx).ChatSession.Query()
	query = r.getQuery(query, req)
	count, err := query.Count(ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *ChatSessionRepo) Page(ctx context.Context, req *repo.ChatSessionQuery) (*repo.ChatSessionPageResp, error) {
	page := r.normalizePage(nil)
	if req != nil {
		page = r.normalizePage(req.Page)
	}
	query := r.getClient(ctx).ChatSession.Query()
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
	chatSessions := make([]*model.ChatSession, 0, len(list))
	for _, item := range list {
		chatSessions = append(chatSessions, r.toModel(item))
	}
	return &repo.ChatSessionPageResp{
		Rows: chatSessions,
		Page: &base.PageResp{
			Total: int64(count),
			Size:  page.Size,
			Page:  page.Page,
		},
	}, nil
}

func (r *ChatSessionRepo) getQuery(query *gen.ChatSessionQuery, req *repo.ChatSessionQuery) *gen.ChatSessionQuery {
	query = query.Where(chatsession.DeletedAtIsNil())
	query.WithGroup().WithLastMessageOfSession()
	if req == nil {
		return query
	}
	if len(req.IDs) > 0 {
		query = query.Where(chatsession.IDIn(req.IDs...))
	}
	if req.CreatedBy != nil {
		query = query.Where(chatsession.CreatedBy(*req.CreatedBy))
	}
	if req.GroupID != nil {
		query = query.Where(chatsession.GroupID(*req.GroupID))
	}
	if req.ReceiverID != nil {
		query = query.Where(chatsession.ReceiverID(*req.ReceiverID))
	}
	return query
}

func (r *ChatSessionRepo) toModel(t *gen.ChatSession) *model.ChatSession {
	item := &model.ChatSession{
		ID:                t.ID,
		ReceiverID:        t.ReceiverID,
		GroupID:           t.GroupID,
		IsMuted:           t.IsMuted,
		IsPinned:          t.IsPinned,
		LastReadMessageID: t.LastReadMessageID,
		ReadCount:         t.ReadCount,
		MessageCount:      t.MessageCount,
		LastMessageID:     t.LastMessageID,
		CreatedAt:         t.CreatedAt,
		UpdatedAt:         t.UpdatedAt,
		CreatedBy:         t.CreatedBy,
		UpdatedBy:         t.UpdatedBy,
	}
	if t.Edges.Group != nil {
		item.Group = &model.ChatGroup{
			ID:            t.Edges.Group.ID,
			Name:          t.Edges.Group.Name,
			AvatarAssetID: t.Edges.Group.AvatarAssetID,
			Introduction:  t.Edges.Group.Introduction,
			OwnerID:       t.Edges.Group.OwnerID,
			Status:        enum.ChatGroupStatus(t.Edges.Group.Status),
			MemberCount:   t.Edges.Group.MemberCount,
			MessageCount:  t.Edges.Group.MessageCount,
			LastMessageID: t.Edges.Group.LastMessageID,
			CreatedAt:     t.Edges.Group.CreatedAt,
			UpdatedAt:     t.Edges.Group.UpdatedAt,
			CreatedBy:     t.Edges.Group.CreatedBy,
			UpdatedBy:     t.Edges.Group.UpdatedBy,
		}
	}
	return item
}
