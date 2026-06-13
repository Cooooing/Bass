package repo

import (
	"common/proto/gen/common"
	cerrors "common/proto/gen/common/errors"
	"context"
	"time"

	"common/pkg/apperror"
	"common/pkg/server"
	utilent "common/pkg/util/ent"
	"im/internal/biz/model"
	"im/internal/biz/repo"
	"im/internal/data/gen"
	"im/internal/data/gen/chatgroupmember"
	"im/internal/enum"
)

var _ repo.ChatGroupMemberRepo = (*ChatGroupMemberRepo)(nil)

type ChatGroupMemberRepo struct {
	db *gen.Client
}

func NewChatGroupMemberRepo(db *gen.Client) repo.ChatGroupMemberRepo {
	return &ChatGroupMemberRepo{db: db}
}

func (r *ChatGroupMemberRepo) getClient(ctx context.Context) *gen.Client {
	if tx, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return tx
	}
	return r.db
}

func (r *ChatGroupMemberRepo) Save(ctx context.Context, chatGroupMember *model.ChatGroupMember) (*model.ChatGroupMember, error) {
	save, err := r.getClient(ctx).ChatGroupMember.Create().
		SetGroupID(chatGroupMember.GroupID).
		SetUserID(chatGroupMember.UserID).
		SetNillableNickname(chatGroupMember.Nickname).
		SetRole(chatgroupmember.Role(chatGroupMember.Role)).
		SetNillableMuteEndAt(chatGroupMember.MuteEndAt).
		SetNillableCreatedBy(chatGroupMember.CreatedBy).
		SetNillableUpdatedBy(chatGroupMember.UpdatedBy).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.toModel(save), nil
}

func (r *ChatGroupMemberRepo) UpdateMuteEndAt(ctx context.Context, groupId int64, userId int64, muteEndAt time.Duration, updatedBy int64) (*model.ChatGroupMember, error) {
	var muteEndAtTime *time.Time
	if muteEndAt > 0 {
		t := time.Now().Add(muteEndAt)
		muteEndAtTime = &t
	}
	query := r.getClient(ctx).ChatGroupMember.Update().
		Where(
			chatgroupmember.GroupID(groupId),
			chatgroupmember.UserID(userId),
		).
		SetUpdatedBy(updatedBy)
	if muteEndAtTime != nil {
		query = query.SetMuteEndAt(*muteEndAtTime)
	} else {
		query = query.ClearMuteEndAt()
	}
	err := query.Exec(ctx)
	if err != nil {
		return nil, err
	}
	// 查询更新后的记录返回
	t, err := r.getClient(ctx).ChatGroupMember.Query().
		Where(chatgroupmember.GroupID(groupId), chatgroupmember.UserID(userId)).
		First(ctx)
	if err != nil {
		return nil, err
	}
	return r.toModel(t), nil
}

func (r *ChatGroupMemberRepo) Get(ctx context.Context, req *repo.ChatGroupMemberGetReq) (*model.ChatGroupMember, error) {
	query := r.getClient(ctx).ChatGroupMember.Query()
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

func (r *ChatGroupMemberRepo) List(ctx context.Context, req *repo.ChatGroupMemberGetReq) ([]*model.ChatGroupMember, error) {
	query := r.getClient(ctx).ChatGroupMember.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.ChatGroupMember, 0, len(list))
	for _, item := range list {
		result = append(result, r.toModel(item))
	}
	return result, nil
}

func (r *ChatGroupMemberRepo) Map(ctx context.Context, req *repo.ChatGroupMemberGetReq) (map[int64]*model.ChatGroupMember, error) {
	list, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*model.ChatGroupMember, len(list))
	for _, item := range list {
		result[item.ID] = item
	}
	return result, nil
}

func (r *ChatGroupMemberRepo) Count(ctx context.Context, req *repo.ChatGroupMemberGetReq) (int, error) {
	query := r.getClient(ctx).ChatGroupMember.Query()
	query = r.getQuery(query, req)
	return query.Count(ctx)
}

func (r *ChatGroupMemberRepo) Page(ctx context.Context, page *common.PageRequest, req *repo.ChatGroupMemberGetReq) ([]*model.ChatGroupMember, *common.PageReply, error) {
	page = server.PageValid(page)
	query := r.getClient(ctx).ChatGroupMember.Query()
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
	result := make([]*model.ChatGroupMember, 0, len(list))
	for _, item := range list {
		result = append(result, r.toModel(item))
	}
	return result, &common.PageReply{
		Total: uint32(count),
		Size:  page.Size,
		Page:  page.Page,
	}, nil
}

func (r *ChatGroupMemberRepo) getQuery(query *gen.ChatGroupMemberQuery, req *repo.ChatGroupMemberGetReq) *gen.ChatGroupMemberQuery {
	if req == nil {
		return query
	}
	if req.IDs != nil {
		query = query.Where(chatgroupmember.IDIn(req.IDs...))
	}
	if req.GroupID != nil {
		query = query.Where(chatgroupmember.GroupID(*req.GroupID))
	}
	if req.UserID != nil {
		query = query.Where(chatgroupmember.UserID(*req.UserID))
	}
	return query
}

func (r *ChatGroupMemberRepo) toModel(t *gen.ChatGroupMember) *model.ChatGroupMember {
	return &model.ChatGroupMember{
		ID:        t.ID,
		GroupID:   t.GroupID,
		UserID:    t.UserID,
		Nickname:  t.Nickname,
		Role:      enum.ChatGroupMemberRole(t.Role),
		MuteEndAt: t.MuteEndAt,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
		CreatedBy: t.CreatedBy,
		UpdatedBy: t.UpdatedBy,
	}
}
