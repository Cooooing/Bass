package repo

import (
	"common/api/gen/common"
	"context"
	"user/internal/biz/model"
	"user/internal/enum"
)

type AccountRepo interface {
	Create(ctx context.Context, u *model.Account) (*model.Account, error)

	Update(ctx context.Context, u *model.Account) (*model.Account, error)
	UpdateProfile(ctx context.Context, patch *AccountProfilePatch) (*model.Account, error)
	AddStat(ctx context.Context, userId int64, statType enum.AccountStatType, num int32) (*model.Account, error)

	ExistsByAccount(ctx context.Context, account string) (bool, error)
	Get(ctx context.Context, req *AccountGetReq) (*model.Account, error)
	GetByAccount(ctx context.Context, account string) (*model.Account, error)
	List(ctx context.Context, req *AccountGetReq) ([]*model.Account, error)
	Page(ctx context.Context, page *common.PageRequest, req *AccountGetReq) ([]*model.Account, *common.PageReply, error)
}

type AccountGetReq struct {
	UserID    *int64
	UserIds   []int64
	Name      *string
	Names     []string
	Nickname  *string
	Nicknames []string
	Email     *string
	Emails    []string
	Phone     *string
	Phones    []string
}

type StringPatch struct {
	Set   bool
	Value string
}

func NewStringPatch(value *string) StringPatch {
	if value == nil {
		return StringPatch{}
	}
	return StringPatch{Set: true, Value: *value}
}

type MBTIPatch struct {
	Set   bool
	Clear bool
	Value enum.MBTI
}

type AccountProfilePatch struct {
	UserID       int64
	AvatarURL    StringPatch
	Nickname     StringPatch
	URL          StringPatch
	Introduction StringPatch
	Mbti         MBTIPatch
}

func (p *AccountProfilePatch) HasChanges() bool {
	return p.AvatarURL.Set ||
		p.Nickname.Set ||
		p.URL.Set ||
		p.Introduction.Set ||
		p.Mbti.Set
}
