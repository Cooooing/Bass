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

func (r *ChatGroupRepo) Save(ctx context.Context, req *repo.ChatGroupSaveReq) (*repo.ChatGroupSaveResponse, error) {
	chatGroup := req.ChatGroup
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
	return &repo.ChatGroupSaveResponse{ChatGroup: r.toModel(save)}, nil
}

func (r *ChatGroupRepo) UpdateAvatar(ctx context.Context, req *repo.ChatGroupUpdateAvatarReq) (*repo.ChatGroupUpdateAvatarResponse, error) {
	update, err := r.getClient(ctx).ChatGroup.UpdateOneID(req.ChatGroupID).
		SetAvatar(req.Avatar).
		SetUpdatedBy(req.UpdatedBy).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &repo.ChatGroupUpdateAvatarResponse{ChatGroup: r.toModel(update)}, nil
}

func (r *ChatGroupRepo) UpdateOwner(ctx context.Context, req *repo.ChatGroupUpdateOwnerReq) (*repo.ChatGroupUpdateOwnerResponse, error) {
	update, err := r.getClient(ctx).ChatGroup.UpdateOneID(req.ChatGroupID).
		SetOwnerID(req.OwnerID).
		SetUpdatedBy(req.UpdatedBy).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &repo.ChatGroupUpdateOwnerResponse{ChatGroup: r.toModel(update)}, nil
}

func (r *ChatGroupRepo) UpdateLastMessage(ctx context.Context, req *repo.ChatGroupUpdateLastMessageReq) (*repo.ChatGroupUpdateLastMessageResponse, error) {
	update, err := r.getClient(ctx).ChatGroup.UpdateOneID(*req.Message.GroupID).
		AddMessageCount(1).
		SetLastMessageID(req.Message.ID).
		SetUpdatedBy(req.UpdatedBy).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &repo.ChatGroupUpdateLastMessageResponse{ChatGroup: r.toModel(update)}, nil
}

func (r *ChatGroupRepo) UpdateMemberCount(ctx context.Context, req *repo.ChatGroupUpdateMemberCountReq) (*repo.ChatGroupUpdateMemberCountResponse, error) {
	update, err := r.getClient(ctx).ChatGroup.UpdateOneID(req.ChatGroupID).
		AddMemberCount(req.OperationMemberCount).
		SetUpdatedBy(req.UpdatedBy).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &repo.ChatGroupUpdateMemberCountResponse{ChatGroup: r.toModel(update)}, nil
}

func (r *ChatGroupRepo) UpdateStatus(ctx context.Context, req *repo.ChatGroupUpdateStatusReq) (*repo.ChatGroupUpdateStatusResponse, error) {
	update, err := r.getClient(ctx).ChatGroup.UpdateOneID(req.ChatGroupID).
		SetStatus(chatgroup.Status(req.Status)).
		SetUpdatedBy(req.UpdatedBy).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &repo.ChatGroupUpdateStatusResponse{ChatGroup: r.toModel(update)}, nil
}

func (r *ChatGroupRepo) Get(ctx context.Context, req *repo.ChatGroupGetReq) (*repo.ChatGroupGetResponse, error) {
	query := r.getClient(ctx).ChatGroup.Query()
	var queryReq *repo.ChatGroupQuery
	if req != nil {
		queryReq = &req.ChatGroupQuery
	}
	query = r.getQuery(query, queryReq)
	t, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_IM_CHAT_GROUP_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return &repo.ChatGroupGetResponse{ChatGroup: r.toModel(t)}, nil
}

func (r *ChatGroupRepo) List(ctx context.Context, req *repo.ChatGroupListReq) (*repo.ChatGroupListResponse, error) {
	query := r.getClient(ctx).ChatGroup.Query()
	var queryReq *repo.ChatGroupQuery
	if req != nil {
		queryReq = &req.ChatGroupQuery
	}
	query = r.getQuery(query, queryReq)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	chatGroups := make([]*model.ChatGroup, 0, len(list))
	for _, item := range list {
		chatGroups = append(chatGroups, r.toModel(item))
	}
	return &repo.ChatGroupListResponse{Rows: chatGroups}, nil
}

func (r *ChatGroupRepo) Map(ctx context.Context, req *repo.ChatGroupMapReq) (*repo.ChatGroupMapResponse, error) {
	listReq := &repo.ChatGroupListReq{}
	if req != nil {
		listReq.ChatGroupQuery = req.ChatGroupQuery
	}
	listResp, err := r.List(ctx, listReq)
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*model.ChatGroup, len(listResp.Rows))
	for _, item := range listResp.Rows {
		result[item.ID] = item
	}
	return &repo.ChatGroupMapResponse{Rows: result}, nil
}

func (r *ChatGroupRepo) Count(ctx context.Context, req *repo.ChatGroupCountReq) (*repo.ChatGroupCountResponse, error) {
	query := r.getClient(ctx).ChatGroup.Query()
	var queryReq *repo.ChatGroupQuery
	if req != nil {
		queryReq = &req.ChatGroupQuery
	}
	query = r.getQuery(query, queryReq)
	count, err := query.Count(ctx)
	if err != nil {
		return nil, err
	}
	return &repo.ChatGroupCountResponse{Count: count}, nil
}

func (r *ChatGroupRepo) Page(ctx context.Context, req *repo.ChatGroupPageReq) (*repo.ChatGroupPageResponse, error) {
	page := normalizePage(nil)
	if req != nil {
		page = normalizePage(req.Page)
	}
	query := r.getClient(ctx).ChatGroup.Query()
	var queryReq *repo.ChatGroupQuery
	if req != nil {
		queryReq = &req.ChatGroupQuery
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
	chatGroups := make([]*model.ChatGroup, 0, len(list))
	for _, item := range list {
		chatGroups = append(chatGroups, r.toModel(item))
	}
	return &repo.ChatGroupPageResponse{
		Rows: chatGroups,
		Page: &base.PageResponse{
			Total: int64(count),
			Size:  page.Size,
			Page:  page.Page,
		},
	}, nil
}

func (r *ChatGroupRepo) getQuery(query *gen.ChatGroupQuery, req *repo.ChatGroupQuery) *gen.ChatGroupQuery {
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
	}
}
