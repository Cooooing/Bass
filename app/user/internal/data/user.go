package data

import (
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

func (r *UserRepo) Save(ctx context.Context, u *model.User) (*model.User, error) {
	userCreate := r.client.User.Create().
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

func (r *UserRepo) GetUserById(ctx context.Context, id int) (*model.User, error) {
	u, err := r.client.User.Query().Where(user.IDEQ(id)).Only(ctx)
	return (*model.User)(u), err
}

func (r *UserRepo) GetUserByAccount(ctx context.Context, account string) (*model.User, error) {
	queryUser, err := r.client.User.Query().Where(user.Or(user.NameEQ(account), user.EmailEQ(account), user.PhoneEQ(account))).Only(ctx)
	return (*model.User)(queryUser), err
}

func (r *UserRepo) ConstantAccount(ctx context.Context, account string) (bool, error) {
	return r.client.User.Query().Where(user.Or(user.NameEQ(account), user.EmailEQ(account), user.PhoneEQ(account))).Exist(ctx)
}
