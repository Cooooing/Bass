package repo

import (
	"common/proto/gen/common"
	"context"
	"user/internal/biz/model"
)

type LocationRepo interface {
	Get(ctx context.Context, req *LocationGetReq) (*model.Location, error)
	List(ctx context.Context, req *LocationGetReq) ([]*model.Location, error)
	Map(ctx context.Context, req *LocationGetReq) (map[int64]*model.Location, error)
	Count(ctx context.Context, req *LocationGetReq) (int, error)
	Page(ctx context.Context, page *common.PageRequest, req *LocationGetReq) ([]*model.Location, *common.PageReply, error)
	UpsertByUserID(ctx context.Context, l *model.Location) (*model.Location, error)
	Update(ctx context.Context, l *model.Location) (*model.Location, error)
}

type LocationGetReq struct {
	ID      *int64
	IDs     []int64
	UserID  *int64
	UserIDs []int64
}
