package repo

import (
	"context"
	"time"
	"user/internal/biz/model"
)

type CheckinRecordRepo interface {
	Get(ctx context.Context, req *CheckinRecordGetReq) (*CheckinRecordGetResponse, error)
	List(ctx context.Context, req *CheckinRecordGetReq) (*CheckinRecordListResponse, error)
	Map(ctx context.Context, req *CheckinRecordGetReq) (*CheckinRecordMapResponse, error)
	Count(ctx context.Context, req *CheckinRecordGetReq) (*CheckinRecordCountResponse, error)
	Page(ctx context.Context, req *CheckinRecordPageReq) (*CheckinRecordPageResponse, error)
	UpsertRecord(ctx context.Context, req *CheckinRecordUpsertReq) (*CheckinRecordUpsertResponse, error)
}

type CheckinStatRepo interface {
	Get(ctx context.Context, req *CheckinStatGetReq) (*CheckinStatGetResponse, error)
	List(ctx context.Context, req *CheckinStatGetReq) (*CheckinStatListResponse, error)
	Map(ctx context.Context, req *CheckinStatGetReq) (*CheckinStatMapResponse, error)
	Count(ctx context.Context, req *CheckinStatGetReq) (*CheckinStatCountResponse, error)
	Page(ctx context.Context, req *CheckinStatPageReq) (*CheckinStatPageResponse, error)
	UpsertStat(ctx context.Context, req *CheckinStatUpsertReq) (*CheckinStatUpsertResponse, error)
}

type CheckinRecordGetReq struct {
	ID      *int64
	IDs     []int64
	UserID  *int64
	UserIDs []int64
	Date    *time.Time
}

type CheckinRecordGetResponse struct {
	Record *model.CheckinRecord
}

type CheckinRecordListResponse struct {
	Rows []*model.CheckinRecord
}

type CheckinRecordMapResponse struct {
	Rows map[int64]*model.CheckinRecord
}

type CheckinRecordCountResponse struct {
	Count int
}

type CheckinRecordPageReq struct {
	Page  PageReq
	Query CheckinRecordGetReq
}

type CheckinRecordPageResponse struct {
	Rows []*model.CheckinRecord
	Page PageResponse
}

type CheckinRecordUpsertReq struct {
	Record *model.CheckinRecord
}

type CheckinRecordUpsertResponse struct {
	Record *model.CheckinRecord
}

type CheckinStatGetReq struct {
	ID      *int64
	IDs     []int64
	UserID  *int64
	UserIDs []int64
}

type CheckinStatGetResponse struct {
	Stat *model.CheckinStat
}

type CheckinStatListResponse struct {
	Rows []*model.CheckinStat
}

type CheckinStatMapResponse struct {
	Rows map[int64]*model.CheckinStat
}

type CheckinStatCountResponse struct {
	Count int
}

type CheckinStatPageReq struct {
	Page  PageReq
	Query CheckinStatGetReq
}

type CheckinStatPageResponse struct {
	Rows []*model.CheckinStat
	Page PageResponse
}

type CheckinStatUpsertReq struct {
	Stat *model.CheckinStat
}

type CheckinStatUpsertResponse struct {
	Stat *model.CheckinStat
}
