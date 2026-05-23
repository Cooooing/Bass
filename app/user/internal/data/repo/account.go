package repo

import (
	"common/api/gen/common"
	cerrors "common/api/gen/common/errors"
	"common/pkg/constant"
	"context"
	"fmt"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/conf"
	"user/internal/data/gen"
	"user/internal/data/gen/account"
	"user/internal/enum"

	commonClient "common/pkg/client"
	utilent "common/pkg/util/ent"

	"github.com/go-kratos/kratos/v2/log"
)

var _ repo.AccountRepo = (*AccountRepo)(nil)

type AccountRepo struct {
	conf   *conf.Bootstrap
	log    *log.Helper
	db     *gen.Client
	consul *commonClient.ConsulClient
	redis  *commonClient.RedisClient
	nats   *commonClient.NatsClient
}

func NewAccountRepo(
	conf *conf.Bootstrap,
	logger log.Logger,
	db *gen.Client,
	consul *commonClient.ConsulClient,
	redis *commonClient.RedisClient,
	nats *commonClient.NatsClient,
) repo.AccountRepo {
	return &AccountRepo{
		conf:   conf,
		log:    log.NewHelper(logger),
		db:     db,
		consul: consul,
		redis:  redis,
		nats:   nats,
	}
}

func (r *AccountRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func toDomain(u *gen.Account) *model.Account {
	return &model.Account{
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
		Status:        (*enum.AccountStatus)(u.Status),
		GroupName:     u.GroupName,
		FollowCount:   u.FollowCount,
		FollowerCount: u.FollowerCount,
		BlockCount:    u.BlockCount,
		BlockedCount:  u.BlockedCount,
		CreatedAt:     u.CreatedAt,
		UpdatedAt:     u.UpdatedAt,
	}
}

func (r *AccountRepo) Create(ctx context.Context, u *model.Account) (*model.Account, error) {
	tx := r.getClient(ctx)
	created, err := tx.Account.Create().
		SetName(u.Name).
		SetPassword(u.Password).
		SetNillableEmail(u.Email).
		SetNillablePhone(u.Phone).
		SetNillableNickname(u.Nickname).
		SetAvatarURL(fmt.Sprintf(r.conf.Server.Avatar, u.Name)).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return toDomain(created), nil
}

func (r *AccountRepo) Update(ctx context.Context, u *model.Account) (*model.Account, error) {
	tx := r.getClient(ctx)
	updated, err := tx.Account.UpdateOneID(u.ID).
		SetNillableAvatarURL(u.AvatarURL).
		SetNillableNickname(u.Nickname).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return toDomain(updated), nil
}

func (r *AccountRepo) AddStat(ctx context.Context, userId int64, statType enum.AccountStatType, num int32) (*model.Account, error) {
	if _, ok := enum.AccountStatTypeMap.ToProto(statType); !ok {
		return nil, fmt.Errorf("unknown statType")
	}

	tx := r.getClient(ctx)
	updateOne := tx.Account.UpdateOneID(userId)
	switch statType {
	case enum.AccountStatTypeFollow:
		updateOne.AddFollowCount(num)
	case enum.AccountStatTypeFollower:
		updateOne.AddFollowerCount(num)
	case enum.AccountStatTypeBlock:
		updateOne.AddBlockCount(num)
	case enum.AccountStatTypeBlocked:
		updateOne.AddBlockedCount(num)
	}
	saved, err := updateOne.Save(ctx)
	if err != nil {
		return nil, err
	}
	return toDomain(saved), nil
}

func (r *AccountRepo) ExistsByAccount(ctx context.Context, accountValue string) (bool, error) {
	tx := r.getClient(ctx)
	return tx.Account.Query().
		Where(account.Or(
			account.NameEQ(accountValue),
			account.EmailEQ(accountValue),
			account.PhoneEQ(accountValue),
		)).
		Exist(ctx)
}

func (r *AccountRepo) Get(ctx context.Context, req *repo.AccountGetReq) (*model.Account, error) {
	tx := r.getClient(ctx)
	query := tx.Account.Query()
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

func (r *AccountRepo) GetByAccount(ctx context.Context, accountValue string) (*model.Account, error) {
	tx := r.getClient(ctx)
	u, err := tx.Account.Query().
		Where(account.Or(
			account.NameEQ(accountValue),
			account.EmailEQ(accountValue),
			account.PhoneEQ(accountValue),
		)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return toDomain(u), nil
}

func (r *AccountRepo) List(ctx context.Context, req *repo.AccountGetReq) ([]*model.Account, error) {
	tx := r.getClient(ctx)
	query := tx.Account.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.Account, 0, len(list))
	for _, u := range list {
		result = append(result, toDomain(u))
	}
	return result, nil
}

func (r *AccountRepo) Page(ctx context.Context, page *common.PageRequest, req *repo.AccountGetReq) ([]*model.Account, *common.PageReply, error) {
	tx := r.getClient(ctx)
	page = constant.PageValid(page)
	query := tx.Account.Query()
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

	result := make([]*model.Account, 0, len(list))
	for _, u := range list {
		result = append(result, toDomain(u))
	}
	return result, &common.PageReply{
		Total: uint32(total),
		Page:  page.Page,
		Size:  page.Size,
	}, nil
}

func (r *AccountRepo) getQuery(query *gen.AccountQuery, req *repo.AccountGetReq) *gen.AccountQuery {
	if req.UserID != nil {
		query = query.Where(account.ID(*req.UserID))
	}
	if len(req.UserIds) > 0 {
		query = query.Where(account.IDIn(req.UserIds...))
	}
	if req.Name != nil {
		query = query.Where(account.NameContains(*req.Name))
	}
	if len(req.Names) > 0 {
		query = query.Where(account.NameIn(req.Names...))
	}
	if req.Nickname != nil {
		query = query.Where(account.NicknameContains(*req.Nickname))
	}
	if len(req.Nicknames) > 0 {
		query = query.Where(account.NicknameIn(req.Nicknames...))
	}
	if req.Email != nil {
		query = query.Where(account.EmailContains(*req.Email))
	}
	if len(req.Emails) > 0 {
		query = query.Where(account.EmailIn(req.Emails...))
	}
	if req.Phone != nil {
		query = query.Where(account.PhoneContains(*req.Phone))
	}
	if len(req.Phones) > 0 {
		query = query.Where(account.PhoneIn(req.Phones...))
	}
	return query
}
