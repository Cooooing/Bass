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
	"im/internal/data/gen/chatgroup"
	"im/internal/enum"
)

var _ repo.ChatGroupRepo = (*ChatGroupRepo)(nil)

type ChatGroupRepo struct {
	pageNormalizer
	db *gen.Client
}

func NewChatGroupRepo(
	db *gen.Client,
) repo.ChatGroupRepo {
	return &ChatGroupRepo{
		db: db,
	}
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
		SetNillableAvatarAssetID(chatGroup.AvatarAssetID).
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
	return r.toModel(save), nil
}

func (r *ChatGroupRepo) UpdateAvatar(ctx context.Context, req *repo.ChatGroupUpdateAvatarReq) (*model.ChatGroup, error) {
	update, err := r.getClient(ctx).ChatGroup.UpdateOneID(req.ChatGroupID).
		SetAvatarAssetID(req.AvatarAssetID).
		SetUpdatedBy(req.UpdatedBy).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.toModel(update), nil
}

func (r *ChatGroupRepo) UpdateOwner(ctx context.Context, req *repo.ChatGroupUpdateOwnerReq) (*model.ChatGroup, error) {
	update, err := r.getClient(ctx).ChatGroup.UpdateOneID(req.ChatGroupID).
		SetOwnerID(req.OwnerID).
		SetUpdatedBy(req.UpdatedBy).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.toModel(update), nil
}

func (r *ChatGroupRepo) UpdateLastMessage(ctx context.Context, req *repo.ChatGroupUpdateLastMessageReq) (*model.ChatGroup, error) {
	update, err := r.getClient(ctx).ChatGroup.UpdateOneID(*req.Message.GroupID).
		AddMessageCount(1).
		SetLastMessageID(req.Message.ID).
		SetUpdatedBy(req.UpdatedBy).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.toModel(update), nil
}

func (r *ChatGroupRepo) UpdateMemberCount(ctx context.Context, req *repo.ChatGroupUpdateMemberCountReq) (*model.ChatGroup, error) {
	update, err := r.getClient(ctx).ChatGroup.UpdateOneID(req.ChatGroupID).
		AddMemberCount(req.OperationMemberCount).
		SetUpdatedBy(req.UpdatedBy).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.toModel(update), nil
}

func (r *ChatGroupRepo) UpdateStatus(ctx context.Context, req *repo.ChatGroupUpdateStatusReq) (*model.ChatGroup, error) {
	update, err := r.getClient(ctx).ChatGroup.UpdateOneID(req.ChatGroupID).
		SetStatus(chatgroup.Status(req.Status)).
		SetUpdatedBy(req.UpdatedBy).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.toModel(update), nil
}

func (r *ChatGroupRepo) Get(ctx context.Context, req *repo.ChatGroupQuery) (*model.ChatGroup, error) {
	query := r.getClient(ctx).ChatGroup.Query()
	query = r.getQuery(query, req)
	t, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_IM_CHAT_GROUP_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return r.toModel(t), nil
}

func (r *ChatGroupRepo) List(ctx context.Context, req *repo.ChatGroupQuery) ([]*model.ChatGroup, error) {
	query := r.getClient(ctx).ChatGroup.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	chatGroups := make([]*model.ChatGroup, 0, len(list))
	for _, item := range list {
		chatGroups = append(chatGroups, r.toModel(item))
	}
	return chatGroups, nil
}

func (r *ChatGroupRepo) Map(ctx context.Context, req *repo.ChatGroupQuery) (map[int64]*model.ChatGroup, error) {
	listResp, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*model.ChatGroup, len(listResp))
	for _, item := range listResp {
		result[item.ID] = item
	}
	return result, nil
}

func (r *ChatGroupRepo) Count(ctx context.Context, req *repo.ChatGroupQuery) (int, error) {
	query := r.getClient(ctx).ChatGroup.Query()
	query = r.getQuery(query, req)
	count, err := query.Count(ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *ChatGroupRepo) Page(ctx context.Context, req *repo.ChatGroupQuery) (*repo.ChatGroupPageResp, error) {
	page := r.normalizePage(nil)
	if req != nil {
		page = r.normalizePage(req.Page)
	}
	query := r.getClient(ctx).ChatGroup.Query()
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
	chatGroups := make([]*model.ChatGroup, 0, len(list))
	for _, item := range list {
		chatGroups = append(chatGroups, r.toModel(item))
	}
	return &repo.ChatGroupPageResp{
		Rows: chatGroups,
		Page: &base.PageResp{
			Total: int64(count),
			Size:  page.Size,
			Page:  page.Page,
		},
	}, nil
}

func (r *ChatGroupRepo) getQuery(query *gen.ChatGroupQuery, req *repo.ChatGroupQuery) *gen.ChatGroupQuery {
	query = query.Where(chatgroup.DeletedAtIsNil())
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

func (r *ChatGroupRepo) toModel(t *gen.ChatGroup) *model.ChatGroup {
	return &model.ChatGroup{
		ID:            t.ID,
		Name:          t.Name,
		AvatarAssetID: t.AvatarAssetID,
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
	}
}
