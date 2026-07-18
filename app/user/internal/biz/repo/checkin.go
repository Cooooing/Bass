package repo

import (
	"context"
	"time"
	"user/internal/biz/model"
)

type CheckinRecordRepo interface {
	Get(ctx context.Context, req *CheckinRecordGetReq) (*model.CheckinRecord, error)
	List(ctx context.Context, req *CheckinRecordGetReq) ([]*model.CheckinRecord, error)
	Map(ctx context.Context, req *CheckinRecordGetReq) (map[int64]*model.CheckinRecord, error)
	Count(ctx context.Context, req *CheckinRecordGetReq) (int, error)
	Page(ctx context.Context, req *CheckinRecordPageReq) (*CheckinRecordPageResp, error)
	UpsertRecord(ctx context.Context, record *model.CheckinRecord) (*model.CheckinRecord, error)
}

type CheckinStatRepo interface {
	Get(ctx context.Context, req *CheckinStatGetReq) (*model.CheckinStat, error)
	List(ctx context.Context, req *CheckinStatGetReq) ([]*model.CheckinStat, error)
	Map(ctx context.Context, req *CheckinStatGetReq) (map[int64]*model.CheckinStat, error)
	Count(ctx context.Context, req *CheckinStatGetReq) (int, error)
	Page(ctx context.Context, req *CheckinStatPageReq) (*CheckinStatPageResp, error)
	UpsertStat(ctx context.Context, stat *model.CheckinStat) (*model.CheckinStat, error)
}

type CheckinRecordGetReq struct {
	ID      *int64
	IDs     []int64
	UserID  *int64
	UserIDs []int64
	Date    *time.Time
}

type CheckinRecordPageReq struct {
	Page  PageReq
	Query CheckinRecordGetReq
}

type CheckinRecordPageResp struct {
	Rows []*model.CheckinRecord
	Page PageResp
}

type CheckinStatGetReq struct {
	ID      *int64
	IDs     []int64
	UserID  *int64
	UserIDs []int64
}

type CheckinStatPageReq struct {
	Page  PageReq
	Query CheckinStatGetReq
}

type CheckinStatPageResp struct {
	Rows []*model.CheckinStat
	Page PageResp
}
