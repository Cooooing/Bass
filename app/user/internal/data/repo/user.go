package repo

import (
	"common/api/gen/common"
	cerrors "common/api/gen/common/errors"
	v1 "common/api/gen/user/v1"
	"common/pkg/constant"
	"context"
	"fmt"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/data/base"
	"user/internal/data/ent/gen"
	"user/internal/data/ent/gen/user"

	utilent "common/pkg/util/ent"
)

type UserRepo struct {
	*base.BaseData
}

func NewUserRepo(repo *base.BaseData) repo.UserRepo {
	return &UserRepo{
		BaseData: repo,
	}
}

func (r *UserRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.Db
}

func toDomain(u *gen.User) *model.User {
	return &model.User{
		ID:            u.ID,
		Name:          u.Name,
		Nickname:      u.Nickname,
		Password:      u.Password,
		Email:         u.Email,
		Phone:         u.Phone,
		URL:           u.URL,
		AvatarURL:     u.AvatarURL,
		Introduction:  u.Introduction,
		Mbti:          u.Mbti,
		Status:        u.Status,
		GroupName:     u.GroupName,
		FollowCount:   u.FollowCount,
		FollowerCount: u.FollowerCount,
		BlockCount:    u.BlockCount,
		BlockedCount:  u.BlockedCount,
		LastLoginTime: u.LastLoginTime,
		LastLoginIP:   u.LastLoginIP,
		CreatedAt:     u.CreatedAt,
		UpdatedAt:     u.UpdatedAt,
	}
}

func (r *UserRepo) Save(ctx context.Context, u *model.User) (*model.User, error) {
	tx := r.getClient(ctx)
	created, err := tx.User.Create().
		SetName(u.Name).
		SetPassword(u.Password).
		SetNillableEmail(u.Email).
		SetNillablePhone(u.Phone).
		SetNillableNickname(u.Nickname).
		SetAvatarURL(fmt.Sprintf(r.Conf.Server.Avatar, u.Name)).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	// Create default sub-table records
	_, _ = tx.UserPreferences.Create().SetUserID(created.ID).Save(ctx)
	_, _ = tx.UserPrivacy.Create().SetUserID(created.ID).Save(ctx)
	return toDomain(created), nil
}

func (r *UserRepo) Update(ctx context.Context, u *model.User) (*model.User, error) {
	tx := r.getClient(ctx)
	updated, err := tx.User.UpdateOneID(u.ID).
		SetNillableAvatarURL(u.AvatarURL).
		SetNillableNickname(u.Nickname).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return toDomain(updated), nil
}

func (r *UserRepo) UpdateStat(ctx context.Context, userId int64, statType v1.UserStatType, num int32) (*model.User, error) {
	tx := r.getClient(ctx)
	updateOne := tx.User.UpdateOneID(userId)
	switch statType {
	case v1.UserStatType_USER_STAT_TYPE_FOLLOW:
		updateOne.AddFollowCount(num)
	case v1.UserStatType_USER_STAT_TYPE_FOLLOWER:
		updateOne.AddFollowerCount(num)
	case v1.UserStatType_USER_STAT_TYPE_BLOCK:
		updateOne.AddBlockCount(num)
	case v1.UserStatType_USER_STAT_TYPE_BLOCKED:
		updateOne.AddBlockedCount(num)
	default:
		return nil, fmt.Errorf("unknown statType")
	}
	saved, err := updateOne.Save(ctx)
	if err != nil {
		return nil, err
	}
	return toDomain(saved), nil
}

func (r *UserRepo) ConstantAccount(ctx context.Context, account string) (bool, error) {
	tx := r.getClient(ctx)
	return tx.User.Query().
		Where(user.Or(
			user.NameEQ(account),
			user.EmailEQ(account),
			user.PhoneEQ(account),
		)).
		Exist(ctx)
}

func (r *UserRepo) GetOne(ctx context.Context, req *repo.UserGetReq) (*model.User, error) {
	tx := r.getClient(ctx)
	query := tx.User.Query().
		WithPreferences().
		WithPrivacy().
		WithLocation().
		WithTfa().
		WithCheckinStat()
	query = r.getQuery(query, req)
	u, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, cerrors.ErrorBadRequest("user is not found")
	}
	if err != nil {
		return nil, err
	}
	return toDomain(u), nil
}

func (r *UserRepo) GetByAccount(ctx context.Context, account string) (*model.User, error) {
	tx := r.getClient(ctx)
	u, err := tx.User.Query().
		WithPreferences().
		WithPrivacy().
		WithLocation().
		WithTfa().
		WithCheckinStat().
		Where(user.Or(
			user.NameEQ(account),
			user.EmailEQ(account),
			user.PhoneEQ(account),
		)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return toDomain(u), nil
}

func (r *UserRepo) GetList(ctx context.Context, req *repo.UserGetReq) ([]*model.User, error) {
	tx := r.getClient(ctx)
	query := tx.User.Query().
		WithPreferences().
		WithPrivacy().
		WithLocation().
		WithTfa().
		WithCheckinStat()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.User, 0, len(list))
	for _, u := range list {
		result = append(result, toDomain(u))
	}
	return result, nil
}

func (r *UserRepo) GetPage(ctx context.Context, page *common.PageRequest, req *repo.UserGetReq) ([]*model.User, *common.PageReply, error) {
	tx := r.getClient(ctx)
	page = constant.PageValid(page)
	query := tx.User.Query().
		WithPreferences().
		WithPrivacy().
		WithLocation().
		WithTfa().
		WithCheckinStat()
	query = r.getQuery(query, req)

	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, nil, err
	}

	list, err := query.
		Limit(int(page.Size)).
		Offset(int((page.Page - 1) * page.Size)).
		All(ctx)
	if err != nil {
		return nil, nil, err
	}

	result := make([]*model.User, 0, len(list))
	for _, u := range list {
		result = append(result, toDomain(u))
	}
	return result, &common.PageReply{
		Total: uint32(total),
		Page:  page.Page,
		Size:  page.Size,
	}, nil
}

func (r *UserRepo) getQuery(query *gen.UserQuery, req *repo.UserGetReq) *gen.UserQuery {
	if req.UserId != nil {
		query = query.Where(user.ID(*req.UserId))
	}
	if len(req.UserIds) > 0 {
		query = query.Where(user.IDIn(req.UserIds...))
	}
	if req.Name != nil {
		query = query.Where(user.NameContains(*req.Name))
	}
	if len(req.Names) > 0 {
		query = query.Where(user.NameIn(req.Names...))
	}
	if req.Nickname != nil {
		query = query.Where(user.NicknameContains(*req.Nickname))
	}
	if len(req.Nicknames) > 0 {
		query = query.Where(user.NicknameIn(req.Nicknames...))
	}
	if req.Email != nil {
		query = query.Where(user.EmailContains(*req.Email))
	}
	if len(req.Emails) > 0 {
		query = query.Where(user.EmailIn(req.Emails...))
	}
	if req.Phone != nil {
		query = query.Where(user.PhoneContains(*req.Phone))
	}
	if len(req.Phones) > 0 {
		query = query.Where(user.PhoneIn(req.Phones...))
	}
	return query
}
