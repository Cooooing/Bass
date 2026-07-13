package repo

import (
	cerrors "common/proto/gen/common/errors"
	"context"

	"common/pkg/apperror"
	"common/pkg/server"
	utilent "common/pkg/util/ent"
	"common/proto/gen/common"
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
	return &model.ChatSession{
		ID:                save.ID,
		ReceiverID:        save.ReceiverID,
		GroupID:           save.GroupID,
		IsMuted:           save.IsMuted,
		IsPinned:          save.IsPinned,
		LastReadMessageID: save.LastReadMessageID,
		ReadCount:         save.ReadCount,
		MessageCount:      save.MessageCount,
		LastMessageID:     save.LastMessageID,
		CreatedAt:         save.CreatedAt,
		UpdatedAt:         save.UpdatedAt,
		CreatedBy:         save.CreatedBy,
		UpdatedBy:         save.UpdatedBy,
	}, nil
}

func (r *ChatSessionRepo) UpdateLastReadMessage(ctx context.Context, chatSessionId int64, messageId int64, operationReadCount int32, updatedBy int64) (*model.ChatSession, error) {
	update, err := r.getClient(ctx).ChatSession.UpdateOneID(chatSessionId).
		SetLastReadMessageID(messageId).
		AddReadCount(operationReadCount).
		SetUpdatedBy(updatedBy).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.ChatSession{
		ID:                update.ID,
		ReceiverID:        update.ReceiverID,
		GroupID:           update.GroupID,
		IsMuted:           update.IsMuted,
		IsPinned:          update.IsPinned,
		LastReadMessageID: update.LastReadMessageID,
		ReadCount:         update.ReadCount,
		MessageCount:      update.MessageCount,
		LastMessageID:     update.LastMessageID,
		CreatedAt:         update.CreatedAt,
		UpdatedAt:         update.UpdatedAt,
		CreatedBy:         update.CreatedBy,
		UpdatedBy:         update.UpdatedBy,
	}, nil
}

func (r *ChatSessionRepo) UpdateMuted(ctx context.Context, chatSessionId int64, muted bool, updatedBy int64) (*model.ChatSession, error) {
	update, err := r.getClient(ctx).ChatSession.UpdateOneID(chatSessionId).
		SetIsMuted(muted).
		SetUpdatedBy(updatedBy).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.ChatSession{
		ID:                update.ID,
		ReceiverID:        update.ReceiverID,
		GroupID:           update.GroupID,
		IsMuted:           update.IsMuted,
		IsPinned:          update.IsPinned,
		LastReadMessageID: update.LastReadMessageID,
		ReadCount:         update.ReadCount,
		MessageCount:      update.MessageCount,
		LastMessageID:     update.LastMessageID,
		CreatedAt:         update.CreatedAt,
		UpdatedAt:         update.UpdatedAt,
		CreatedBy:         update.CreatedBy,
		UpdatedBy:         update.UpdatedBy,
	}, nil
}

func (r *ChatSessionRepo) UpdatePinned(ctx context.Context, chatSessionId int64, pinned bool, updatedBy int64) (*model.ChatSession, error) {
	update, err := r.getClient(ctx).ChatSession.UpdateOneID(chatSessionId).
		SetIsPinned(pinned).
		SetUpdatedBy(updatedBy).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.ChatSession{
		ID:                update.ID,
		ReceiverID:        update.ReceiverID,
		GroupID:           update.GroupID,
		IsMuted:           update.IsMuted,
		IsPinned:          update.IsPinned,
		LastReadMessageID: update.LastReadMessageID,
		ReadCount:         update.ReadCount,
		MessageCount:      update.MessageCount,
		LastMessageID:     update.LastMessageID,
		CreatedAt:         update.CreatedAt,
		UpdatedAt:         update.UpdatedAt,
		CreatedBy:         update.CreatedBy,
		UpdatedBy:         update.UpdatedBy,
	}, nil
}

func (r *ChatSessionRepo) Get(ctx context.Context, req *repo.ChatSessionGetReq) (*model.ChatSession, error) {
	query := r.getClient(ctx).ChatSession.Query()
	query = r.getQuery(query, req)
	t, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_IM_CHAT_SESSION_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
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
	return item, nil
}

func (r *ChatSessionRepo) List(ctx context.Context, req *repo.ChatSessionGetReq) ([]*model.ChatSession, error) {
	query := r.getClient(ctx).ChatSession.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}

	chatSessions := make([]*model.ChatSession, 0, len(list))
	for _, item := range list {
		session := &model.ChatSession{
			ID:                item.ID,
			ReceiverID:        item.ReceiverID,
			GroupID:           item.GroupID,
			IsMuted:           item.IsMuted,
			IsPinned:          item.IsPinned,
			LastReadMessageID: item.LastReadMessageID,
			ReadCount:         item.ReadCount,
			MessageCount:      item.MessageCount,
			LastMessageID:     item.LastMessageID,
			CreatedAt:         item.CreatedAt,
			UpdatedAt:         item.UpdatedAt,
			CreatedBy:         item.CreatedBy,
			UpdatedBy:         item.UpdatedBy,
		}
		if item.Edges.Group != nil {
			session.Group = &model.ChatGroup{
				ID:            item.Edges.Group.ID,
				Name:          item.Edges.Group.Name,
				Avatar:        item.Edges.Group.Avatar,
				Introduction:  item.Edges.Group.Introduction,
				OwnerID:       item.Edges.Group.OwnerID,
				Status:        enum.ChatGroupStatus(item.Edges.Group.Status),
				MemberCount:   item.Edges.Group.MemberCount,
				MessageCount:  item.Edges.Group.MessageCount,
				LastMessageID: item.Edges.Group.LastMessageID,
				CreatedAt:     item.Edges.Group.CreatedAt,
				UpdatedAt:     item.Edges.Group.UpdatedAt,
				CreatedBy:     item.Edges.Group.CreatedBy,
				UpdatedBy:     item.Edges.Group.UpdatedBy,
			}
		}
		chatSessions = append(chatSessions, session)
	}
	return chatSessions, nil
}

func (r *ChatSessionRepo) Map(ctx context.Context, req *repo.ChatSessionGetReq) (map[int64]*model.ChatSession, error) {
	list, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*model.ChatSession, len(list))
	for _, item := range list {
		result[item.ID] = item
	}
	return result, nil
}

func (r *ChatSessionRepo) Count(ctx context.Context, req *repo.ChatSessionGetReq) (int, error) {
	query := r.getClient(ctx).ChatSession.Query()
	query = r.getQuery(query, req)
	return query.Count(ctx)
}

func (r *ChatSessionRepo) Page(ctx context.Context, page *common.PageRequest, req *repo.ChatSessionGetReq) ([]*model.ChatSession, *common.PageReply, error) {
	page = server.PageValid(page)
	query := r.getClient(ctx).ChatSession.Query()
	query = r.getQuery(query, req)
	countQuery := query.Clone()
	count, err := countQuery.Count(ctx)
	if err != nil {
		return nil, nil, err
	}
	list, err := query.Limit(int(page.Size)).Offset(int((page.Page - 1) * page.Size)).All(ctx)
	if err != nil {
		return nil, nil, err
	}

	chatSessions := make([]*model.ChatSession, 0, len(list))
	for _, item := range list {
		session := &model.ChatSession{
			ID:                item.ID,
			ReceiverID:        item.ReceiverID,
			GroupID:           item.GroupID,
			IsMuted:           item.IsMuted,
			IsPinned:          item.IsPinned,
			LastReadMessageID: item.LastReadMessageID,
			ReadCount:         item.ReadCount,
			MessageCount:      item.MessageCount,
			LastMessageID:     item.LastMessageID,
			CreatedAt:         item.CreatedAt,
			UpdatedAt:         item.UpdatedAt,
			CreatedBy:         item.CreatedBy,
			UpdatedBy:         item.UpdatedBy,
		}
		if item.Edges.Group != nil {
			session.Group = &model.ChatGroup{
				ID:            item.Edges.Group.ID,
				Name:          item.Edges.Group.Name,
				Avatar:        item.Edges.Group.Avatar,
				Introduction:  item.Edges.Group.Introduction,
				OwnerID:       item.Edges.Group.OwnerID,
				Status:        enum.ChatGroupStatus(item.Edges.Group.Status),
				MemberCount:   item.Edges.Group.MemberCount,
				MessageCount:  item.Edges.Group.MessageCount,
				LastMessageID: item.Edges.Group.LastMessageID,
				CreatedAt:     item.Edges.Group.CreatedAt,
				UpdatedAt:     item.Edges.Group.UpdatedAt,
				CreatedBy:     item.Edges.Group.CreatedBy,
				UpdatedBy:     item.Edges.Group.UpdatedBy,
			}
		}
		chatSessions = append(chatSessions, session)
	}
	return chatSessions, &common.PageReply{
		Total: uint32(count),
		Size:  page.Size,
		Page:  page.Page,
	}, nil
}

func (r *ChatSessionRepo) getQuery(query *gen.ChatSessionQuery, req *repo.ChatSessionGetReq) *gen.ChatSessionQuery {
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
