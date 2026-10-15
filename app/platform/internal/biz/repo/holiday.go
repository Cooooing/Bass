package repo

import (
	"context"
	"time"

	"platform/internal/biz/model"
)

// HolidayClient 封装外部节假日数据源，返回与具体供应商无关的日期状态。
type HolidayClient interface {
	Fetch(ctx context.Context, date time.Time) (*model.HolidayInfo, error)
}

// HolidayRepo 管理按自然日持久化的节假日状态。
type HolidayRepo interface {
	GetByDate(ctx context.Context, date time.Time) (*model.HolidayInfo, error)
	Upsert(ctx context.Context, info *model.HolidayInfo) (*model.HolidayInfo, error)
}
