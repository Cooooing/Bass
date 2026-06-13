package repo

import (
	cerrors "common/proto/gen/common/errors"
	"context"

	"common/proto/gen/common"
	"common/pkg/apperror"
	"common/pkg/server"
	utilent "common/pkg/util/ent"
	"im/internal/biz/model"
	"im/internal/biz/repo"
	"im/internal/data/gen"
	"im/internal/data/gen/chatgroup"
	"im/internal/enum"
)

var _ repo.ChatGroupRepo = (*ChatGroupRepo)(nil)

type ChatGroupRepo struct {
	db *gen.Client
}

func NewChatGroupRepo(db *gen.Client) repo.ChatGroupRepo {
	return &ChatGroupRepo{db: db}
}

func (r *ChatGroupRepo) getClient(ctx context.Context) *gen.Client {
	if tx, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return tx
	}
	return r.db
}

func (r *ChatGroupRepo) Save(ctx context.Context, chatGroup *model.ChatGroup) (*model.ChatGroup, error) {
	save, err := r.getClient(ctx).ChatGroup.Create().
		SetName(chatGroup.Name).
		SetNillableAvatar(chatGroup.Avatar).
		SetNillableIntroduction(chatGroup.Introduction).
		SetOwnerID(chatGroup.OwnerID).
		SetStatus(chatgroup.Status(enum.ChatGroupStatusNormal)).
		SetMemberCount(chatGroup.MemberCount).
		SetNillableCreatedBy(chatGroup.CreatedBy).
		SetNillableUpdatedBy(chatGroup.UpdatedBy).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.ChatGroup{
		ID:            save.ID,
		Name:          save.Name,
		Avatar:        save.Avatar,
		Introduction:  save.Introduction,
		OwnerID:       save.OwnerID,
		Status:        enum.ChatGroupStatus(save.Status),
		MemberCount:   save.MemberCount,
		MessageCount:  save.MessageCount,
		LastMessageID: save.LastMessageID,
		CreatedAt:     save.CreatedAt,
		UpdatedAt:     save.UpdatedAt,
		CreatedBy:     save.CreatedBy,
		UpdatedBy:     save.UpdatedBy,
	}, nil
}

func (r *ChatGroupRepo) UpdateAvatar(ctx context.Context, chatGroupId int64, avatar string, updatedBy int64) (*model.ChatGroup, error) {
	update, err := r.getClient(ctx).ChatGroup.UpdateOneID(chatGroupId).
		SetAvatar(avatar).
		SetUpdatedBy(updatedBy).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.ChatGroup{
		ID:            update.ID,
		Name:          update.Name,
		Avatar:        update.Avatar,
		Introduction:  update.Introduction,
		OwnerID:       update.OwnerID,
		Status:        enum.ChatGroupStatus(update.Status),
		MemberCount:   update.MemberCount,
		MessageCount:  update.MessageCount,
		LastMessageID: update.LastMessageID,
		CreatedAt:     update.CreatedAt,
		UpdatedAt:     update.UpdatedAt,
		CreatedBy:     update.CreatedBy,
		UpdatedBy:     update.UpdatedBy,
	}, nil
}

func (r *ChatGroupRepo) UpdateOwner(ctx context.Context, chatGroupId int64, ownerId int64, updatedBy int64) (*model.ChatGroup, error) {
	update, err := r.getClient(ctx).ChatGroup.UpdateOneID(chatGroupId).
		SetOwnerID(ownerId).
		SetUpdatedBy(updatedBy).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.ChatGroup{
		ID:            update.ID,
		Name:          update.Name,
		Avatar:        update.Avatar,
		Introduction:  update.Introduction,
		OwnerID:       update.OwnerID,
		Status:        enum.ChatGroupStatus(update.Status),
		MemberCount:   update.MemberCount,
		MessageCount:  update.MessageCount,
		LastMessageID: update.LastMessageID,
		CreatedAt:     update.CreatedAt,
		UpdatedAt:     update.UpdatedAt,
		CreatedBy:     update.CreatedBy,
		UpdatedBy:     update.UpdatedBy,
	}, nil
}

func (r *ChatGroupRepo) UpdateLastMessage(ctx context.Context, message *model.ChatMessage, updatedBy int64) (*model.ChatGroup, error) {
	update, err := r.getClient(ctx).ChatGroup.UpdateOneID(*message.GroupID).
		AddMessageCount(1).
		SetLastMessageID(message.ID).
		SetUpdatedBy(updatedBy).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.ChatGroup{
		ID:            update.ID,
		Name:          update.Name,
		Avatar:        update.Avatar,
		Introduction:  update.Introduction,
		OwnerID:       update.OwnerID,
		Status:        enum.ChatGroupStatus(update.Status),
		MemberCount:   update.MemberCount,
		MessageCount:  update.MessageCount,
		LastMessageID: update.LastMessageID,
		CreatedAt:     update.CreatedAt,
		UpdatedAt:     update.UpdatedAt,
		CreatedBy:     update.CreatedBy,
		UpdatedBy:     update.UpdatedBy,
	}, nil
}

func (r *ChatGroupRepo) UpdateMemberCount(ctx context.Context, chatGroupId int64, operationMemberCount int32, updatedBy int64) (*model.ChatGroup, error) {
	update, err := r.getClient(ctx).ChatGroup.UpdateOneID(chatGroupId).
		AddMemberCount(operationMemberCount).
		SetUpdatedBy(updatedBy).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.ChatGroup{
		ID:            update.ID,
		Name:          update.Name,
		Avatar:        update.Avatar,
		Introduction:  update.Introduction,
		OwnerID:       update.OwnerID,
		Status:        enum.ChatGroupStatus(update.Status),
		MemberCount:   update.MemberCount,
		MessageCount:  update.MessageCount,
		LastMessageID: update.LastMessageID,
		CreatedAt:     update.CreatedAt,
		UpdatedAt:     update.UpdatedAt,
		CreatedBy:     update.CreatedBy,
		UpdatedBy:     update.UpdatedBy,
	}, nil
}

func (r *ChatGroupRepo) UpdateStatus(ctx context.Context, chatGroupId int64, status enum.ChatGroupStatus, updatedBy int64) (*model.ChatGroup, error) {
	update, err := r.getClient(ctx).ChatGroup.UpdateOneID(chatGroupId).
		SetStatus(chatgroup.Status(status)).
		SetUpdatedBy(updatedBy).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.ChatGroup{
		ID:            update.ID,
		Name:          update.Name,
		Avatar:        update.Avatar,
		Introduction:  update.Introduction,
		OwnerID:       update.OwnerID,
		Status:        enum.ChatGroupStatus(update.Status),
		MemberCount:   update.MemberCount,
		MessageCount:  update.MessageCount,
		LastMessageID: update.LastMessageID,
		CreatedAt:     update.CreatedAt,
		UpdatedAt:     update.UpdatedAt,
		CreatedBy:     update.CreatedBy,
		UpdatedBy:     update.UpdatedBy,
	}, nil
}

func (r *ChatGroupRepo) Get(ctx context.Context, req *repo.ChatGroupGetReq) (*model.ChatGroup, error) {
	query := r.getClient(ctx).ChatGroup.Query()
	query = r.getQuery(query, req)
	t, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_IM_CHAT_GROUP_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return &model.ChatGroup{
		ID:            t.ID,
		Name:          t.Name,
		Avatar:        t.Avatar,
		Introduction:  t.Introduction,
		OwnerID:       t.OwnerID,
		Status:        enum.ChatGroupStatus(t.Status),
		MemberCount:   t.MemberCount,
		MessageCount:  t.MessageCount,
		LastMessageID: t.LastMessageID,
		CreatedAt:     t.CreatedAt,
		UpdatedAt:     t.UpdatedAt,
		CreatedBy:     t.CreatedBy,
		UpdatedBy:     t.UpdatedBy,
	}, nil
}

func (r *ChatGroupRepo) List(ctx context.Context, req *repo.ChatGroupGetReq) ([]*model.ChatGroup, error) {
	query := r.getClient(ctx).ChatGroup.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}

	chatGroups := make([]*model.ChatGroup, 0, len(list))
	for _, item := range list {
		chatGroups = append(chatGroups, &model.ChatGroup{
			ID:            item.ID,
			Name:          item.Name,
			Avatar:        item.Avatar,
			Introduction:  item.Introduction,
			OwnerID:       item.OwnerID,
			Status:        enum.ChatGroupStatus(item.Status),
			MemberCount:   item.MemberCount,
			MessageCount:  item.MessageCount,
			LastMessageID: item.LastMessageID,
			CreatedAt:     item.CreatedAt,
			UpdatedAt:     item.UpdatedAt,
			CreatedBy:     item.CreatedBy,
			UpdatedBy:     item.UpdatedBy,
		})
	}
	return chatGroups, nil
}

func (r *ChatGroupRepo) Map(ctx context.Context, req *repo.ChatGroupGetReq) (map[int64]*model.ChatGroup, error) {
	list, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*model.ChatGroup, len(list))
	for _, item := range list {
		result[item.ID] = item
	}
	return result, nil
}

func (r *ChatGroupRepo) Count(ctx context.Context, req *repo.ChatGroupGetReq) (int, error) {
	query := r.getClient(ctx).ChatGroup.Query()
	query = r.getQuery(query, req)
	return query.Count(ctx)
}

func (r *ChatGroupRepo) Page(ctx context.Context, page *common.PageRequest, req *repo.ChatGroupGetReq) ([]*model.ChatGroup, *common.PageReply, error) {
	page = server.PageValid(page)
	query := r.getClient(ctx).ChatGroup.Query()
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

	chatGroups := make([]*model.ChatGroup, 0, len(list))
	for _, item := range list {
		chatGroups = append(chatGroups, &model.ChatGroup{
			ID:            item.ID,
			Name:          item.Name,
			Avatar:        item.Avatar,
			Introduction:  item.Introduction,
			OwnerID:       item.OwnerID,
			Status:        enum.ChatGroupStatus(item.Status),
			MemberCount:   item.MemberCount,
			MessageCount:  item.MessageCount,
			LastMessageID: item.LastMessageID,
			CreatedAt:     item.CreatedAt,
			UpdatedAt:     item.UpdatedAt,
			CreatedBy:     item.CreatedBy,
			UpdatedBy:     item.UpdatedBy,
		})
	}
	return chatGroups, &common.PageReply{
		Total: uint32(count),
		Size:  page.Size,
		Page:  page.Page,
	}, nil
}

func (r *ChatGroupRepo) getQuery(query *gen.ChatGroupQuery, req *repo.ChatGroupGetReq) *gen.ChatGroupQuery {
	if req == nil {
		return query
	}
	if len(req.IDs) > 0 {
		query = query.Where(chatgroup.IDIn(req.IDs...))
	}
	if req.Status != nil {
		query = query.Where(chatgroup.StatusEQ(chatgroup.Status(*req.Status)))
	}
	return query
}
