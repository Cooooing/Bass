package data

import (
	cv1 "common/api/common/v1"
	"common/pkg/constant"
	"common/pkg/util/base"
	"context"
	"fmt"
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

func (r *UserRepo) Save(ctx context.Context, tx *gen.Client, u *model.User) (*model.User, error) {
	userCreate := tx.User.Create().
		SetName(u.Name).
		SetPassword(u.Password).
		SetEmail(u.Email).
		SetNillablePhone(u.Phone).
		SetNickname(u.Nickname).
		SetAvatarURL(fmt.Sprintf(r.conf.Oss.Local.Avatar, u.Name))
	createdUser, err := userCreate.Save(ctx)
	return (*model.User)(createdUser), err
}

func (r *UserRepo) Update(ctx context.Context, tx *gen.Client, u *model.User) (*model.User, error) {
	update := tx.User.UpdateOneID(u.ID)
	update.SetNillableAvatarURL(u.AvatarURL)
	update.SetNillableLanguage(u.Language)
	update.SetNillableNickname(u.Timezone)
	update.SetNillableEmail(u.Theme)
	update.SetNillableMobileTheme(u.MobileTheme)
	update.SetNillableEnableWebNotify(u.EnableWebNotify)
	update.SetNillableEnableEmailSubscribe(u.EnableEmailSubscribe)
	update.SetNillablePublicPoints(u.PublicPoints)
	update.SetNillablePublicFollowers(u.PublicFollowers)
	update.SetNillablePublicArticles(u.PublicArticles)
	update.SetNillablePublicComments(u.PublicComments)
	update.SetNillablePublicOnlineStatus(u.PublicOnlineStatus)
	save, err := update.Save(ctx)
	return (*model.User)(save), err
}

func (r *UserRepo) GetById(ctx context.Context, tx *gen.Client, id int64) (*model.User, error) {
	u, err := tx.User.Query().Where(user.IDEQ(id)).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, cv1.ErrorBadRequest("user is not found")
	}
	return (*model.User)(u), err
}

func (r *UserRepo) GetByAccount(ctx context.Context, tx *gen.Client, account string) (*model.User, error) {
	queryUser, err := tx.User.Query().Where(user.Or(user.NameEQ(account), user.EmailEQ(account), user.PhoneEQ(account))).Only(ctx)
	return (*model.User)(queryUser), err
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
		res = append(res, (*model.User)(u))
	}
	return res, nil
}
func (r *UserRepo) GetPage(ctx context.Context, tx *gen.Client, page *cv1.PageRequest, req *repo.UserGetReq) ([]*model.User, *cv1.PageReply, error) {
	var (
		users []*model.User
		err   error
		total int
	)
	page = base.OrDefault(page, constant.GetPageDefault())
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
		users = append(users, (*model.User)(u))
	}
	return users, &cv1.PageReply{
		Total: uint32(total),
		Page:  page.Page,
		Size:  page.Size,
	}, nil
}

func (r *UserRepo) getQuery(query *gen.UserQuery, req *repo.UserGetReq) *gen.UserQuery {
	if len(req.ArticleIds) > 0 {
		query = query.Where(user.IDIn(req.ArticleIds...))
	}
	return query
}
