package repo

import (
	"common/proto/gen/common"
	"context"
	"user/internal/biz/model"
)

type PreferencesRepo interface {
	Get(ctx context.Context, req *PreferencesGetReq) (*model.Preferences, error)
	List(ctx context.Context, req *PreferencesGetReq) ([]*model.Preferences, error)
	Map(ctx context.Context, req *PreferencesGetReq) (map[int64]*model.Preferences, error)
	Count(ctx context.Context, req *PreferencesGetReq) (int, error)
	Page(ctx context.Context, page *common.PageRequest, req *PreferencesGetReq) ([]*model.Preferences, *common.PageReply, error)
	UpsertByUserID(ctx context.Context, p *model.Preferences) (*model.Preferences, error)
	Update(ctx context.Context, p *model.Preferences) (*model.Preferences, error)
}

type PreferencesGetReq struct {
	ID      *int64
	IDs     []int64
	UserID  *int64
	UserIDs []int64
}
