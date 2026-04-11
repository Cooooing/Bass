package repo

import (
	cv1 "common/api/gen/common/v1"
	v1 "common/api/gen/user/v1"
	"common/pkg/constant"
	"context"
	"fmt"
	"time"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/data/base"
	"user/internal/data/ent/gen"
	"user/internal/data/ent/gen/user"
)

type UserRepo struct {
	*base.BaseData
}

func NewUserRepo(repo *base.BaseData) repo.UserRepo {
	return &UserRepo{
		BaseData: repo,
	}
}

func (r *UserRepo) Save(ctx context.Context, tx *gen.Client, u *model.User) (*model.User, error) {
	userCreate := tx.User.Create().
		SetName(u.Name).
		SetPassword(u.Password).
		SetNillableEmail(u.Email).
		SetNillablePhone(u.Phone).
		SetNillableNickname(u.Nickname).
		SetAvatarURL(fmt.Sprintf(r.Conf.Server.Avatar, u.Name))
	createdUser, err := userCreate.Save(ctx)
	return &model.User{User: createdUser}, err
}

func (r *UserRepo) Update(ctx context.Context, tx *gen.Client, u *model.User) (*model.User, error) {
	update := tx.User.UpdateOneID(u.ID)
	update.SetNillableAvatarURL(u.AvatarURL)
	update.SetNillableNickname(u.Nickname)
	update.SetNillableLanguage(u.Language)
	update.SetNillableTimezone(u.Timezone)
	update.SetNillableTheme(u.Theme)
	update.SetNillableMobileTheme(u.MobileTheme)
	update.SetNillableEnableWebNotify(u.EnableWebNotify)
	update.SetNillableEnableEmailSubscribe(u.EnableEmailSubscribe)
	update.SetNillablePublicPoints(u.PublicPoints)
	update.SetNillablePublicFollowers(u.PublicFollowers)
	update.SetNillablePublicArticles(u.PublicArticles)
	update.SetNillablePublicComments(u.PublicComments)
	update.SetNillablePublicOnlineStatus(u.PublicOnlineStatus)
	save, err := update.Save(ctx)
	return &model.User{User: save}, err
}

func (r *UserRepo) UpdateStat(ctx context.Context, tx *gen.Client, userId int64, statType v1.UserStatType, num int32) (*model.User, error) {
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
	save, err := updateOne.Save(ctx)
	return &model.User{User: save}, err
}

func (r *UserRepo) EnableTwoFactorAuthentication(ctx context.Context, tx *gen.Client, name string, secret string) (int, error) {
	return tx.User.Update().
		Where(user.NameEQ(name)).
		SetTwofaEnable(true).
		SetTwofaEnableTime(time.Now()).
		SetTwofaSecret(secret).
		Save(ctx)
}

func (r *UserRepo) DisableTwoFactorAuthentication(ctx context.Context, tx *gen.Client, name string) (int, error) {
	return tx.User.Update().
		Where(user.NameEQ(name)).
		SetTwofaEnable(false).
		SetTwofaSecret("").
		Save(ctx)
}

func (r *UserRepo) GetOne(ctx context.Context, tx *gen.Client, req *repo.UserGetReq) (*model.User, error) {
	query := tx.User.Query()
	query = r.getQuery(query, req)
	u, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, cv1.ErrorBadRequest("user is not found")
	}
	return &model.User{User: u}, err
}

func (r *UserRepo) GetByAccount(ctx context.Context, tx *gen.Client, account string) (*model.User, error) {
	queryUser, err := tx.User.Query().Where(user.Or(user.NameEQ(account), user.EmailEQ(account), user.PhoneEQ(account))).Only(ctx)
	return &model.User{User: queryUser}, err
}

func (r *UserRepo) ConstantAccount(ctx context.Context, tx *gen.Client, account string) (bool, error) {
	return tx.User.Query().Where(user.Or(user.NameEQ(account), user.EmailEQ(account), user.PhoneEQ(account))).Exist(ctx)
}

func (r *UserRepo) GetList(ctx context.Context, tx *gen.Client, req *repo.UserGetReq) ([]*model.User, error) {
	res := make([]*model.User, 0)
	query := tx.User.Query()
	query = r.getQuery(query, req)
	users, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	for _, u := range users {
		res = append(res, &model.User{User: u})
	}
	return res, nil
}
func (r *UserRepo) GetPage(ctx context.Context, tx *gen.Client, page *cv1.PageRequest, req *repo.UserGetReq) ([]*model.User, *cv1.PageReply, error) {
	var (
		users []*model.User
		err   error
		total int
	)
	page = constant.PageValid(page)
	query := tx.User.Query()
	query = r.getQuery(query, req)
	countQuery := query.Clone()
	total, err = countQuery.Count(ctx)
	if err != nil {
		return nil, nil, err
	}
	list, err := query.Limit(int(page.Size)).Offset(int((page.Page - 1) * page.Size)).All(ctx)
	if err != nil {
		return nil, nil, err
	}
	for _, u := range list {
		users = append(users, &model.User{User: u})
	}
	return users, &cv1.PageReply{
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
