package data

import (
	cv1 "common/api/common/v1"
	"context"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/data/ent/gen"
	"user/internal/data/ent/gen/user"
)

type UserRepo struct {
	*BaseRepo
	client *gen.Client
}

func NewUserRepo(repo *BaseRepo, client *gen.Client) repo.UserRepo {
	return &UserRepo{
		BaseRepo: repo,
		client:   client,
	}
}

func (r *UserRepo) Save(ctx context.Context, client *gen.Client, u *model.User) (*model.User, error) {
	userCreate := client.User.Create().
		SetName(u.Name).
		SetPassword(u.Password)
	if u.Email != "" {
		userCreate.SetEmail(u.Email)
	}
	if u.Phone != "" {
		userCreate.SetPhone(u.Phone)
	}
	if u.Nickname != "" {
		userCreate.SetNickname(u.Nickname)
	}
	createdUser, err := userCreate.Save(ctx)
	return (*model.User)(createdUser), err
}

func (r *UserRepo) GetUserById(ctx context.Context, client *gen.Client, id int64) (*model.User, error) {
	u, err := client.User.Query().Where(user.IDEQ(id)).Only(ctx)
	return (*model.User)(u), err
}

func (r *UserRepo) GetUserByAccount(ctx context.Context, client *gen.Client, account string) (*model.User, error) {
	queryUser, err := client.User.Query().Where(user.Or(user.NameEQ(account), user.EmailEQ(account), user.PhoneEQ(account))).Only(ctx)
	return (*model.User)(queryUser), err
}

func (r *UserRepo) ConstantAccount(ctx context.Context, client *gen.Client, account string) (bool, error) {
	return client.User.Query().Where(user.Or(user.NameEQ(account), user.EmailEQ(account), user.PhoneEQ(account))).Exist(ctx)
}

func (r *UserRepo) GetUserList(ctx context.Context, client *gen.Client, page *cv1.PageRequest, ids []int64) ([]*model.User, error) {
	res := make([]*model.User, 0)
	query := client.User.Query()
	if len(ids) > 0 {
		query.Where(user.IDIn(ids...))
	}
	users, err := query.Limit(int(page.Size)).Offset(int((page.Page - 1) * page.Size)).All(ctx)
	if err != nil {
		return nil, err
	}
	for _, u := range users {
		res = append(res, (*model.User)(u))
	}
	return res, nil
}
