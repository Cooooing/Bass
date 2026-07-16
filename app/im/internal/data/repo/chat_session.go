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
	db *gen.Client
}

func NewChatSessionRepo(db *gen.Client) repo.ChatSessionRepo {
	return &ChatSessionRepo{db: db}
}

func (r *ChatSessionRepo) getClient(ctx context.Context) *gen.Client {
	if tx, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return tx
	}
	return r.db
}

func (r *ChatSessionRepo) Save(ctx context.Context, req *repo.ChatSessionSaveReq) (*repo.ChatSessionSaveResponse, error) {
	chatSession := req.ChatSession
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
	return &repo.ChatSessionSaveResponse{ChatSession: r.toModel(save)}, nil
}

func (r *ChatSessionRepo) UpdateLastReadMessage(ctx context.Context, req *repo.ChatSessionUpdateLastReadMessageReq) (*repo.ChatSessionUpdateLastReadMessageResponse, error) {
	update, err := r.getClient(ctx).ChatSession.UpdateOneID(req.ChatSessionID).
		SetLastReadMessageID(req.MessageID).
		AddReadCount(req.OperationReadCount).
		SetUpdatedBy(req.UpdatedBy).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &repo.ChatSessionUpdateLastReadMessageResponse{ChatSession: r.toModel(update)}, nil
}

func (r *ChatSessionRepo) UpdateMuted(ctx context.Context, req *repo.ChatSessionUpdateMutedReq) (*repo.ChatSessionUpdateMutedResponse, error) {
	update, err := r.getClient(ctx).ChatSession.UpdateOneID(req.ChatSessionID).
		SetIsMuted(req.Muted).
		SetUpdatedBy(req.UpdatedBy).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &repo.ChatSessionUpdateMutedResponse{ChatSession: r.toModel(update)}, nil
}

func (r *ChatSessionRepo) UpdatePinned(ctx context.Context, req *repo.ChatSessionUpdatePinnedReq) (*repo.ChatSessionUpdatePinnedResponse, error) {
	update, err := r.getClient(ctx).ChatSession.UpdateOneID(req.ChatSessionID).
		SetIsPinned(req.Pinned).
		SetUpdatedBy(req.UpdatedBy).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &repo.ChatSessionUpdatePinnedResponse{ChatSession: r.toModel(update)}, nil
}

func (r *ChatSessionRepo) Get(ctx context.Context, req *repo.ChatSessionGetReq) (*repo.ChatSessionGetResponse, error) {
	query := r.getClient(ctx).ChatSession.Query()
	var queryReq *repo.ChatSessionQuery
	if req != nil {
		queryReq = &req.ChatSessionQuery
	}
	query = r.getQuery(query, queryReq)
	t, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_IM_CHAT_SESSION_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return &repo.ChatSessionGetResponse{ChatSession: r.toModel(t)}, nil
}

func (r *ChatSessionRepo) List(ctx context.Context, req *repo.ChatSessionListReq) (*repo.ChatSessionListResponse, error) {
	query := r.getClient(ctx).ChatSession.Query()
	var queryReq *repo.ChatSessionQuery
	if req != nil {
		queryReq = &req.ChatSessionQuery
	}
	query = r.getQuery(query, queryReq)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	chatSessions := make([]*model.ChatSession, 0, len(list))
	for _, item := range list {
		chatSessions = append(chatSessions, r.toModel(item))
	}
	return &repo.ChatSessionListResponse{Rows: chatSessions}, nil
}

func (r *ChatSessionRepo) Map(ctx context.Context, req *repo.ChatSessionMapReq) (*repo.ChatSessionMapResponse, error) {
	listReq := &repo.ChatSessionListReq{}
	if req != nil {
		listReq.ChatSessionQuery = req.ChatSessionQuery
	}
	listResp, err := r.List(ctx, listReq)
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*model.ChatSession, len(listResp.Rows))
	for _, item := range listResp.Rows {
		result[item.ID] = item
	}
	return &repo.ChatSessionMapResponse{Rows: result}, nil
}

func (r *ChatSessionRepo) Count(ctx context.Context, req *repo.ChatSessionCountReq) (*repo.ChatSessionCountResponse, error) {
	query := r.getClient(ctx).ChatSession.Query()
	var queryReq *repo.ChatSessionQuery
	if req != nil {
		queryReq = &req.ChatSessionQuery
	}
	query = r.getQuery(query, queryReq)
	count, err := query.Count(ctx)
	if err != nil {
		return nil, err
	}
	return &repo.ChatSessionCountResponse{Count: count}, nil
}

func (r *ChatSessionRepo) Page(ctx context.Context, req *repo.ChatSessionPageReq) (*repo.ChatSessionPageResponse, error) {
	page := normalizePage(nil)
	if req != nil {
		page = normalizePage(req.Page)
	}
	query := r.getClient(ctx).ChatSession.Query()
	var queryReq *repo.ChatSessionQuery
	if req != nil {
		queryReq = &req.ChatSessionQuery
	}
	query = r.getQuery(query, queryReq)
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
	return &repo.ChatSessionPageResponse{
		Rows: chatSessions,
		Page: &base.PageResponse{
			Total: int64(count),
			Size:  page.Size,
			Page:  page.Page,
		},
	}, nil
}

func (r *ChatSessionRepo) getQuery(query *gen.ChatSessionQuery, req *repo.ChatSessionQuery) *gen.ChatSessionQuery {
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
			Avatar:        t.Edges.Group.Avatar,
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
