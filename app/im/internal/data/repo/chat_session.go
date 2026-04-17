package repo

import (
	"common/api/gen/common"
	cerrors "common/api/gen/common/errors"
	"common/pkg/constant"
	"context"
	"fmt"
	"im/internal/biz/model"
	"im/internal/biz/repo"
	database "im/internal/data/base"
	"im/internal/data/ent/gen"
)

type ChatSessionRepo struct {
	*database.BaseData
}

func NewChatSessionRepo(BaseData *database.BaseData) repo.ChatSessionRepo {
	return &ChatSessionRepo{
		BaseData: BaseData,
	}
}

func (r *ChatSessionRepo) Save(ctx context.Context, tx *gen.Client, chatSession *model.ChatSession) (*model.ChatSession, error) {
	if (chatSession.ReceiverID == nil && chatSession.GroupID == nil) || (chatSession.ReceiverID != nil && chatSession.GroupID != nil) {
		return nil, fmt.Errorf("receiver_id and group_id cannot be both nil or not nil at the same time")
	}
	save, err := tx.ChatSession.Create().
		SetNillableReceiverID(chatSession.ReceiverID).
		SetNillableGroupID(chatSession.GroupID).
		Save(ctx)
	return &model.ChatSession{ChatSession: save}, err
}

func (r *ChatSessionRepo) UpdateLastReadMessage(ctx context.Context, tx *gen.Client, chatSessionId int64, messageId int64, operationReadCount int32) (*model.ChatSession, error) {
	tx.ChatMessage.Query()
	update, err := tx.ChatSession.UpdateOneID(chatSessionId).
		SetLastReadMessageID(messageId).
		AddReadCount(operationReadCount).
		Save(ctx)
	return &model.ChatSession{ChatSession: update}, err
}

func (r *ChatSessionRepo) UpdateMuted(ctx context.Context, tx *gen.Client, chatSessionId int64, muted bool) (*model.ChatSession, error) {
	update, err := tx.ChatSession.UpdateOneID(chatSessionId).
		SetIsMuted(muted).
		Save(ctx)
	return &model.ChatSession{ChatSession: update}, err
}

func (r *ChatSessionRepo) UpdatePinned(ctx context.Context, tx *gen.Client, chatSessionId int64, pinned bool) (*model.ChatSession, error) {
	update, err := tx.ChatSession.UpdateOneID(chatSessionId).
		SetIsPinned(pinned).
		Save(ctx)
	return &model.ChatSession{ChatSession: update}, err
}

func (r *ChatSessionRepo) GetOne(ctx context.Context, tx *gen.Client, req *repo.ChatSessionGetReq) (*model.ChatSession, error) {
	query := tx.ChatSession.Query()
	query = r.getQuery(query, req)
	t, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, cerrors.ErrorBadRequest("chatSession is not found")
	}
	return &model.ChatSession{ChatSession: t}, err
}

func (r *ChatSessionRepo) GetList(ctx context.Context, tx *gen.Client, req *repo.ChatSessionGetReq) ([]*model.ChatSession, error) {
	var (
		chatSessions []*model.ChatSession
		err          error
	)
	query := tx.ChatSession.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}

	for _, item := range list {
		chatSessions = append(chatSessions, &model.ChatSession{ChatSession: item})
	}
	return chatSessions, nil
}

func (r *ChatSessionRepo) GetPage(ctx context.Context, tx *gen.Client, page *common.PageRequest, req *repo.ChatSessionGetReq) ([]*model.ChatSession, *common.PageReply, error) {
	var (
		chatSessions []*model.ChatSession
		err          error
	)
	page = constant.PageValid(page)
	query := tx.ChatSession.Query()
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

	for _, item := range list {
		chatSessions = append(chatSessions, &model.ChatSession{ChatSession: item})
	}
	return chatSessions, &common.PageReply{
		Total: uint32(count),
		Size:  page.Size,
		Page:  page.Page,
	}, nil
}

func (r *ChatSessionRepo) getQuery(query *gen.ChatSessionQuery, req *repo.ChatSessionGetReq) *gen.ChatSessionQuery {
	query.WithGroup().WithLastMessageOfSession()
	return query
}
