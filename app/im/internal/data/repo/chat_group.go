package repo

import (
	cv1 "common/api/gen/common/v1"
	v1 "common/api/gen/im/v1"
	"common/pkg/constant"
	"context"
	"im/internal/biz/model"
	"im/internal/biz/repo"
	database "im/internal/data/base"
	"im/internal/data/ent/gen"
)

type ChatGroupRepo struct {
	*database.BaseData
}

func NewChatGroupRepo(BaseData *database.BaseData) repo.ChatGroupRepo {
	return &ChatGroupRepo{
		BaseData: BaseData,
	}
}

func (r *ChatGroupRepo) Save(ctx context.Context, tx *gen.Client, chatGroup *model.ChatGroup) (*model.ChatGroup, error) {
	save, err := tx.ChatGroup.Create().
		SetName(chatGroup.Name).
		SetNillableAvatar(chatGroup.Avatar).
		SetNillableIntroduction(chatGroup.Introduction).
		SetOwnerID(chatGroup.OwnerID).
		SetStatus(int32(v1.ChatGroupStatus_CHAT_GROUP_STATUS_NORMAL)).
		SetMemberCount(chatGroup.MemberCount).
		Save(ctx)
	return &model.ChatGroup{ChatGroup: save}, err
}

func (r *ChatGroupRepo) UpdateAvatar(ctx context.Context, tx *gen.Client, chatGroupId int64, avatar string) (*model.ChatGroup, error) {
	update, err := tx.ChatGroup.UpdateOneID(chatGroupId).
		SetAvatar(avatar).
		Save(ctx)
	return &model.ChatGroup{ChatGroup: update}, err
}

func (r *ChatGroupRepo) UpdateOwner(ctx context.Context, tx *gen.Client, chatGroupId int64, ownerId int64) (*model.ChatGroup, error) {
	update, err := tx.ChatGroup.UpdateOneID(chatGroupId).
		SetOwnerID(ownerId).
		Save(ctx)
	return &model.ChatGroup{ChatGroup: update}, err
}

func (r *ChatGroupRepo) UpdateLastMessage(ctx context.Context, tx *gen.Client, message *model.ChatMessage) (*model.ChatGroup, error) {
	update, err := tx.ChatGroup.UpdateOneID(*message.GroupID).
		AddMessageCount(1).
		SetLastMessageID(message.ID).
		Save(ctx)
	return &model.ChatGroup{ChatGroup: update}, err
}

func (r *ChatGroupRepo) UpdateMemberCount(ctx context.Context, tx *gen.Client, chatGroupId int64, operationMemberCount int32) (*model.ChatGroup, error) {
	update, err := tx.ChatGroup.UpdateOneID(chatGroupId).
		AddMemberCount(operationMemberCount).
		Save(ctx)
	return &model.ChatGroup{ChatGroup: update}, err
}

func (r *ChatGroupRepo) UpdateStatus(ctx context.Context, tx *gen.Client, chatGroupId int64, status v1.ChatGroupStatus) (*model.ChatGroup, error) {
	update, err := tx.ChatGroup.UpdateOneID(chatGroupId).
		SetStatus(int32(status)).
		Save(ctx)
	return &model.ChatGroup{ChatGroup: update}, err
}

func (r *ChatGroupRepo) GetOne(ctx context.Context, tx *gen.Client, req *repo.ChatGroupGetReq) (*model.ChatGroup, error) {
	query := tx.ChatGroup.Query()
	query = r.getQuery(query, req)
	t, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, cv1.ErrorBadRequest("chatGroup is not found")
	}
	return &model.ChatGroup{ChatGroup: t}, err
}

func (r *ChatGroupRepo) GetList(ctx context.Context, tx *gen.Client, req *repo.ChatGroupGetReq) ([]*model.ChatGroup, error) {
	var (
		chatGroups []*model.ChatGroup
		err        error
	)
	query := tx.ChatGroup.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}

	for _, item := range list {
		chatGroups = append(chatGroups, &model.ChatGroup{ChatGroup: item})
	}
	return chatGroups, nil
}

func (r *ChatGroupRepo) GetPage(ctx context.Context, tx *gen.Client, page *cv1.PageRequest, req *repo.ChatGroupGetReq) ([]*model.ChatGroup, *cv1.PageReply, error) {
	var (
		chatGroups []*model.ChatGroup
		err        error
	)
	page = constant.PageValid(page)
	query := tx.ChatGroup.Query()
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
		chatGroups = append(chatGroups, &model.ChatGroup{ChatGroup: item})
	}
	return chatGroups, &cv1.PageReply{
		Total: uint32(count),
		Size:  page.Size,
		Page:  page.Page,
	}, nil
}

func (r *ChatGroupRepo) getQuery(query *gen.ChatGroupQuery, req *repo.ChatGroupGetReq) *gen.ChatGroupQuery {
	return query
}
