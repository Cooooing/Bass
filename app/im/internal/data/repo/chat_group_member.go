package repo

import (
	cerrors "common/proto/gen/common/errors"
	"context"
	"im/internal/biz/base"
	"time"

	"common/pkg/apperror"
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

func (r *ChatGroupMemberRepo) Save(ctx context.Context, req *repo.ChatGroupMemberSaveReq) (*repo.ChatGroupMemberSaveResponse, error) {
	chatGroupMember := req.ChatGroupMember
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
	return &repo.ChatGroupMemberSaveResponse{ChatGroupMember: r.toModel(save)}, nil
}

func (r *ChatGroupMemberRepo) UpdateMuteEndAt(ctx context.Context, req *repo.ChatGroupMemberUpdateMuteEndAtReq) (*repo.ChatGroupMemberUpdateMuteEndAtResponse, error) {
	var muteEndAtTime *time.Time
	if req.MuteEndAt > 0 {
		t := time.Now().Add(req.MuteEndAt)
		muteEndAtTime = &t
	}
	query := r.getClient(ctx).ChatGroupMember.Update().
		Where(
			chatgroupmember.GroupID(req.GroupID),
			chatgroupmember.UserID(req.UserID),
		).
		SetUpdatedBy(req.UpdatedBy)
	if muteEndAtTime != nil {
		query = query.SetMuteEndAt(*muteEndAtTime)
	} else {
		query = query.ClearMuteEndAt()
	}
	err := query.Exec(ctx)
	if err != nil {
		return nil, err
	}
	t, err := r.getClient(ctx).ChatGroupMember.Query().
		Where(chatgroupmember.GroupID(req.GroupID), chatgroupmember.UserID(req.UserID)).
		First(ctx)
	if err != nil {
		return nil, err
	}
	return &repo.ChatGroupMemberUpdateMuteEndAtResponse{ChatGroupMember: r.toModel(t)}, nil
}

func (r *ChatGroupMemberRepo) Get(ctx context.Context, req *repo.ChatGroupMemberGetReq) (*repo.ChatGroupMemberGetResponse, error) {
	query := r.getClient(ctx).ChatGroupMember.Query()
	var queryReq *repo.ChatGroupMemberQuery
	if req != nil {
		queryReq = &req.ChatGroupMemberQuery
	}
	query = r.getQuery(query, queryReq)
	t, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_IM_CHAT_GROUP_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return &repo.ChatGroupMemberGetResponse{ChatGroupMember: r.toModel(t)}, nil
}

func (r *ChatGroupMemberRepo) List(ctx context.Context, req *repo.ChatGroupMemberListReq) (*repo.ChatGroupMemberListResponse, error) {
	query := r.getClient(ctx).ChatGroupMember.Query()
	var queryReq *repo.ChatGroupMemberQuery
	if req != nil {
		queryReq = &req.ChatGroupMemberQuery
	}
	query = r.getQuery(query, queryReq)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.ChatGroupMember, 0, len(list))
	for _, item := range list {
		result = append(result, r.toModel(item))
	}
	return &repo.ChatGroupMemberListResponse{Rows: result}, nil
}

func (r *ChatGroupMemberRepo) Map(ctx context.Context, req *repo.ChatGroupMemberMapReq) (*repo.ChatGroupMemberMapResponse, error) {
	listReq := &repo.ChatGroupMemberListReq{}
	if req != nil {
		listReq.ChatGroupMemberQuery = req.ChatGroupMemberQuery
	}
	listResp, err := r.List(ctx, listReq)
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*model.ChatGroupMember, len(listResp.Rows))
	for _, item := range listResp.Rows {
		result[item.ID] = item
	}
	return &repo.ChatGroupMemberMapResponse{Rows: result}, nil
}

func (r *ChatGroupMemberRepo) Count(ctx context.Context, req *repo.ChatGroupMemberCountReq) (*repo.ChatGroupMemberCountResponse, error) {
	query := r.getClient(ctx).ChatGroupMember.Query()
	var queryReq *repo.ChatGroupMemberQuery
	if req != nil {
		queryReq = &req.ChatGroupMemberQuery
	}
	query = r.getQuery(query, queryReq)
	count, err := query.Count(ctx)
	if err != nil {
		return nil, err
	}
	return &repo.ChatGroupMemberCountResponse{Count: count}, nil
}

func (r *ChatGroupMemberRepo) Page(ctx context.Context, req *repo.ChatGroupMemberPageReq) (*repo.ChatGroupMemberPageResponse, error) {
	page := normalizePage(nil)
	if req != nil {
		page = normalizePage(req.Page)
	}
	query := r.getClient(ctx).ChatGroupMember.Query()
	var queryReq *repo.ChatGroupMemberQuery
	if req != nil {
		queryReq = &req.ChatGroupMemberQuery
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
	result := make([]*model.ChatGroupMember, 0, len(list))
	for _, item := range list {
		result = append(result, r.toModel(item))
	}
	return &repo.ChatGroupMemberPageResponse{
		Rows: result,
		Page: &base.PageResponse{
			Total: int64(count),
			Size:  page.Size,
			Page:  page.Page,
		},
	}, nil
}

func (r *ChatGroupMemberRepo) getQuery(query *gen.ChatGroupMemberQuery, req *repo.ChatGroupMemberQuery) *gen.ChatGroupMemberQuery {
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
