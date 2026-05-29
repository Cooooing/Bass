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

	utilent "common/pkg/util/ent"
)

var _ repo.AccountRepo = (*AccountRepo)(nil)

type AccountRepo struct {
	conf *conf.Bootstrap
	db   *gen.Client
}

func NewAccountRepo(
	conf *conf.Bootstrap,
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
		Status:        (*enum.AccountStatus)(created.Status),
		FollowCount:   created.FollowCount,
		FollowerCount: created.FollowerCount,
		CreatedAt:     created.CreatedAt,
		UpdatedAt:     created.UpdatedAt,
	}, nil
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
		Status:        (*enum.AccountStatus)(updated.Status),
		FollowCount:   updated.FollowCount,
		FollowerCount: updated.FollowerCount,
		CreatedAt:     updated.CreatedAt,
		UpdatedAt:     updated.UpdatedAt,
	}, nil
}

func (r *AccountRepo) UpdateProfile(ctx context.Context, userID int64, avatarURL *string, nickname *string, url *string, introduction *string, mbti *enum.MBTI, clearMBTI bool) (*model.Account, error) {
	if avatarURL == nil && nickname == nil && url == nil && introduction == nil && mbti == nil && !clearMBTI {
		return r.Get(ctx, &repo.AccountGetReq{UserID: &userID})
	}

	tx := r.getClient(ctx)
	update := tx.Account.UpdateOneID(userID)
	if avatarURL != nil {
		if *avatarURL == "" {
			update.ClearAvatarURL()
		} else {
			update.SetAvatarURL(*avatarURL)
		}
	}
	if nickname != nil {
		if *nickname == "" {
			update.ClearNickname()
		} else {
			update.SetNickname(*nickname)
		}
	}
	if url != nil {
		if *url == "" {
			update.ClearURL()
		} else {
			update.SetURL(*url)
		}
	}
	if introduction != nil {
		if *introduction == "" {
			update.ClearIntroduction()
		} else {
			update.SetIntroduction(*introduction)
		}
	}
	if clearMBTI {
		update.ClearMbti()
	} else if mbti != nil {
		update.SetMbti(account.Mbti(*mbti))
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
		Status:        (*enum.AccountStatus)(updated.Status),
		FollowCount:   updated.FollowCount,
		FollowerCount: updated.FollowerCount,
		CreatedAt:     updated.CreatedAt,
		UpdatedAt:     updated.UpdatedAt,
	}, nil
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
		Status:        (*enum.AccountStatus)(saved.Status),
		FollowCount:   saved.FollowCount,
		FollowerCount: saved.FollowerCount,
		CreatedAt:     saved.CreatedAt,
		UpdatedAt:     saved.UpdatedAt,
	}, nil
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
		Status:        (*enum.AccountStatus)(u.Status),
		FollowCount:   u.FollowCount,
		FollowerCount: u.FollowerCount,
		CreatedAt:     u.CreatedAt,
		UpdatedAt:     u.UpdatedAt,
	}, nil
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
		Status:        (*enum.AccountStatus)(u.Status),
		FollowCount:   u.FollowCount,
		FollowerCount: u.FollowerCount,
		CreatedAt:     u.CreatedAt,
		UpdatedAt:     u.UpdatedAt,
	}, nil
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
			Status:        (*enum.AccountStatus)(u.Status),
			FollowCount:   u.FollowCount,
			FollowerCount: u.FollowerCount,
			CreatedAt:     u.CreatedAt,
			UpdatedAt:     u.UpdatedAt,
		})
	}
	return result, nil
}

func (r *AccountRepo) Map(ctx context.Context, req *repo.AccountGetReq) (map[int64]*model.Account, error) {
	list, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*model.Account, len(list))
	for _, item := range list {
		result[item.ID] = item
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
			Status:        (*enum.AccountStatus)(u.Status),
			FollowCount:   u.FollowCount,
			FollowerCount: u.FollowerCount,
			CreatedAt:     u.CreatedAt,
			UpdatedAt:     u.UpdatedAt,
		})
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
