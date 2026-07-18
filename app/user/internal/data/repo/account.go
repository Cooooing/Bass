package repo

import (
	"common/pkg/apperror"
	"common/proto/gen/common"
	cerrors "common/proto/gen/common/errors"
	"context"
	"fmt"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/config"
	"user/internal/data/gen"
	"user/internal/data/gen/account"
	"user/internal/enum"

	"common/pkg/server"
	utilent "common/pkg/util/ent"
)

var _ repo.AccountRepo = (*AccountRepo)(nil)

type AccountRepo struct {
	conf *config.Bootstrap
	db   *gen.Client
}

func NewAccountRepo(
	conf *config.Bootstrap,
	db *gen.Client,
) repo.AccountRepo {
	return &AccountRepo{
		conf: conf,
		db:   db,
	}
}

func (r *AccountRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func (r *AccountRepo) Create(ctx context.Context, account *model.Account) (*model.Account, error) {
	return r.create(ctx, account)
}

func (r *AccountRepo) Update(ctx context.Context, account *model.Account) (*model.Account, error) {
	return r.update(ctx, account)
}

func (r *AccountRepo) UpdateProfile(ctx context.Context, profile *model.AccountProfileUpdate) (*model.Account, error) {
	return r.updateProfile(ctx, profile)
}

func (r *AccountRepo) AddStat(ctx context.Context, req *repo.AccountAddStatReq) (*model.Account, error) {
	return r.addStat(ctx, req.UserID, req.StatType, req.Num)
}

func (r *AccountRepo) ExistsByAccount(ctx context.Context, account string) (bool, error) {
	return r.existsByAccount(ctx, account)
}

func (r *AccountRepo) Get(ctx context.Context, req *repo.AccountGetReq) (*model.Account, error) {
	return r.get(ctx, req)
}

func (r *AccountRepo) List(ctx context.Context, req *repo.AccountGetReq) ([]*model.Account, error) {
	return r.list(ctx, req)
}

func (r *AccountRepo) Map(ctx context.Context, req *repo.AccountGetReq) (map[int64]*model.Account, error) {
	return r.mapRows(ctx, req)
}

func (r *AccountRepo) Count(ctx context.Context, req *repo.AccountGetReq) (int, error) {
	return r.count(ctx, req)
}

func (r *AccountRepo) Page(ctx context.Context, req *repo.AccountPageReq) (*repo.AccountPageResp, error) {
	rows, page, err := r.page(ctx, &common.PageReq{Page: req.Page.Page, Size: req.Page.Size}, &req.Query)
	if err != nil {
		return nil, err
	}
	resp := repo.PageResp{}
	if page != nil {
		resp = repo.PageResp{Total: page.GetTotal(), Page: page.GetPage(), Size: page.GetSize()}
	}
	return &repo.AccountPageResp{Rows: rows, Page: resp}, nil
}
func (r *AccountRepo) create(ctx context.Context, u *model.Account) (*model.Account, error) {
	tx := r.getClient(ctx)
	created, err := tx.Account.Create().
		SetName(u.Name).
		SetPassword(u.Password).
		SetNillableEmail(u.Email).
		SetNillablePhone(u.Phone).
		SetNillableNickname(u.Nickname).
		SetAvatarURL(fmt.Sprintf(r.conf.Business.Avatar, u.Name)).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.Account{
		ID:            created.ID,
		Name:          created.Name,
		Nickname:      created.Nickname,
		Password:      created.Password,
		Email:         created.Email,
		Phone:         created.Phone,
		URL:           created.URL,
		AvatarURL:     created.AvatarURL,
		Introduction:  created.Introduction,
		Mbti:          (*enum.MBTI)(created.Mbti),
		Status:        new(enum.AccountStatus(created.Status)),
		FollowCount:   new(created.FollowCount),
		FollowerCount: new(created.FollowerCount),
		CreatedAt:     created.CreatedAt,
		UpdatedAt:     created.UpdatedAt,
	}, nil
}

func (r *AccountRepo) update(ctx context.Context, u *model.Account) (*model.Account, error) {
	tx := r.getClient(ctx)
	updated, err := tx.Account.UpdateOneID(u.ID).
		SetNillableAvatarURL(u.AvatarURL).
		SetNillableNickname(u.Nickname).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.Account{
		ID:            updated.ID,
		Name:          updated.Name,
		Nickname:      updated.Nickname,
		Password:      updated.Password,
		Email:         updated.Email,
		Phone:         updated.Phone,
		URL:           updated.URL,
		AvatarURL:     updated.AvatarURL,
		Introduction:  updated.Introduction,
		Mbti:          (*enum.MBTI)(updated.Mbti),
		Status:        new(enum.AccountStatus(updated.Status)),
		FollowCount:   new(updated.FollowCount),
		FollowerCount: new(updated.FollowerCount),
		CreatedAt:     updated.CreatedAt,
		UpdatedAt:     updated.UpdatedAt,
	}, nil
}

func (r *AccountRepo) updateProfile(ctx context.Context, req *model.AccountProfileUpdate) (*model.Account, error) {
	if req.AvatarURL == nil && req.Nickname == nil && req.URL == nil && req.Introduction == nil && req.Mbti == nil && !req.ClearMBTI {
		return r.get(ctx, &repo.AccountGetReq{UserID: &req.UserID})
	}

	tx := r.getClient(ctx)
	update := tx.Account.UpdateOneID(req.UserID)
	if req.AvatarURL != nil {
		if *req.AvatarURL == "" {
			update.ClearAvatarURL()
		} else {
			update.SetAvatarURL(*req.AvatarURL)
		}
	}
	if req.Nickname != nil {
		if *req.Nickname == "" {
			update.ClearNickname()
		} else {
			update.SetNickname(*req.Nickname)
		}
	}
	if req.URL != nil {
		if *req.URL == "" {
			update.ClearURL()
		} else {
			update.SetURL(*req.URL)
		}
	}
	if req.Introduction != nil {
		if *req.Introduction == "" {
			update.ClearIntroduction()
		} else {
			update.SetIntroduction(*req.Introduction)
		}
	}
	if req.ClearMBTI {
		update.ClearMbti()
	} else if req.Mbti != nil {
		update.SetMbti(account.Mbti(*req.Mbti))
	}

	updated, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.Account{
		ID:            updated.ID,
		Name:          updated.Name,
		Nickname:      updated.Nickname,
		Password:      updated.Password,
		Email:         updated.Email,
		Phone:         updated.Phone,
		URL:           updated.URL,
		AvatarURL:     updated.AvatarURL,
		Introduction:  updated.Introduction,
		Mbti:          (*enum.MBTI)(updated.Mbti),
		Status:        new(enum.AccountStatus(updated.Status)),
		FollowCount:   new(updated.FollowCount),
		FollowerCount: new(updated.FollowerCount),
		CreatedAt:     updated.CreatedAt,
		UpdatedAt:     updated.UpdatedAt,
	}, nil
}

func (r *AccountRepo) addStat(ctx context.Context, userId int64, statType enum.AccountStatType, num int32) (*model.Account, error) {
	_ = enum.AccountStatTypeMap.MustToProto(statType)

	tx := r.getClient(ctx)
	updateOne := tx.Account.UpdateOneID(userId)
	switch statType {
	case enum.AccountStatTypeFollow:
		updateOne.AddFollowCount(num)
	case enum.AccountStatTypeFollower:
		updateOne.AddFollowerCount(num)
	}
	saved, err := updateOne.Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.Account{
		ID:            saved.ID,
		Name:          saved.Name,
		Nickname:      saved.Nickname,
		Password:      saved.Password,
		Email:         saved.Email,
		Phone:         saved.Phone,
		URL:           saved.URL,
		AvatarURL:     saved.AvatarURL,
		Introduction:  saved.Introduction,
		Mbti:          (*enum.MBTI)(saved.Mbti),
		Status:        new(enum.AccountStatus(saved.Status)),
		FollowCount:   new(saved.FollowCount),
		FollowerCount: new(saved.FollowerCount),
		CreatedAt:     saved.CreatedAt,
		UpdatedAt:     saved.UpdatedAt,
	}, nil
}

func (r *AccountRepo) existsByAccount(ctx context.Context, accountValue string) (bool, error) {
	tx := r.getClient(ctx)
	return tx.Account.Query().
		Where(account.Or(
			account.NameEQ(accountValue),
			account.EmailEQ(accountValue),
			account.PhoneEQ(accountValue),
		)).
		Exist(ctx)
}

func (r *AccountRepo) get(ctx context.Context, req *repo.AccountGetReq) (*model.Account, error) {
	tx := r.getClient(ctx)
	query := tx.Account.Query()
	query = r.getQuery(query, req)
	u, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_ACCOUNT_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
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
		Mbti:          (*enum.MBTI)(u.Mbti),
		Status:        new(enum.AccountStatus(u.Status)),
		FollowCount:   new(u.FollowCount),
		FollowerCount: new(u.FollowerCount),
		CreatedAt:     u.CreatedAt,
		UpdatedAt:     u.UpdatedAt,
	}, nil
}

func (r *AccountRepo) list(ctx context.Context, req *repo.AccountGetReq) ([]*model.Account, error) {
	tx := r.getClient(ctx)
	query := tx.Account.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.Account, 0, len(list))
	for _, u := range list {
		result = append(result, &model.Account{
			ID:            u.ID,
			Name:          u.Name,
			Nickname:      u.Nickname,
			Password:      u.Password,
			Email:         u.Email,
			Phone:         u.Phone,
			URL:           u.URL,
			AvatarURL:     u.AvatarURL,
			Introduction:  u.Introduction,
			Mbti:          (*enum.MBTI)(u.Mbti),
			Status:        new(enum.AccountStatus(u.Status)),
			FollowCount:   new(u.FollowCount),
			FollowerCount: new(u.FollowerCount),
			CreatedAt:     u.CreatedAt,
			UpdatedAt:     u.UpdatedAt,
		})
	}
	return result, nil
}

func (r *AccountRepo) mapRows(ctx context.Context, req *repo.AccountGetReq) (map[int64]*model.Account, error) {
	list, err := r.list(ctx, req)
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*model.Account, len(list))
	for _, item := range list {
		result[item.ID] = item
	}
	return result, nil
}

func (r *AccountRepo) count(ctx context.Context, req *repo.AccountGetReq) (int, error) {
	tx := r.getClient(ctx)
	query := tx.Account.Query()
	query = r.getQuery(query, req)
	return query.Count(ctx)
}

func (r *AccountRepo) page(ctx context.Context, page *common.PageReq, req *repo.AccountGetReq) ([]*model.Account, *common.PageResp, error) {
	tx := r.getClient(ctx)
	page = server.PageValid(page)
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
		result = append(result, &model.Account{
			ID:            u.ID,
			Name:          u.Name,
			Nickname:      u.Nickname,
			Password:      u.Password,
			Email:         u.Email,
			Phone:         u.Phone,
			URL:           u.URL,
			AvatarURL:     u.AvatarURL,
			Introduction:  u.Introduction,
			Mbti:          (*enum.MBTI)(u.Mbti),
			Status:        new(enum.AccountStatus(u.Status)),
			FollowCount:   new(u.FollowCount),
			FollowerCount: new(u.FollowerCount),
			CreatedAt:     u.CreatedAt,
			UpdatedAt:     u.UpdatedAt,
		})
	}
	return result, &common.PageResp{
		Total: uint32(total),
		Page:  page.Page,
		Size:  page.Size,
	}, nil
}

func (r *AccountRepo) getQuery(query *gen.AccountQuery, req *repo.AccountGetReq) *gen.AccountQuery {
	if req == nil {
		return query
	}
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
	if req.Account != nil {
		query = query.Where(account.Or(
			account.NameEQ(*req.Account),
			account.EmailEQ(*req.Account),
			account.PhoneEQ(*req.Account),
		))
	}
	return query
}
