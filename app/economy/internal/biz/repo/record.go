package repo

import (
	"context"
	"economy/internal/biz/base"
	"economy/internal/biz/model"
	"economy/internal/enum"
)

type RecordRepo interface {
	Save(ctx context.Context, record *model.Record) (*model.Record, error)
	Get(ctx context.Context, req *RecordGetReq) (*model.Record, error)
	List(ctx context.Context, req *RecordGetReq) ([]*model.Record, error)
	Map(ctx context.Context, req *RecordGetReq) (map[int64]*model.Record, error)
	Count(ctx context.Context, req *RecordGetReq) (int, error)
	Page(ctx context.Context, req *RecordGetReq) (*RecordPageResp, error)
}

type RecordGetReq struct {
	Page           *base.PageRequest
	ID             *int64
	IDs            []int64
	TransactionNo  *string
	UserID         *int64
	UserIDs        []int64
	RecordType     *enum.EconomyRecordType
	Direction      *enum.EconomyRecordDirection
	SourceID       *string
	IdempotencyKey *string
}

type RecordPageResp struct {
	Rows []*model.Record
	Page *base.PageResp
}
