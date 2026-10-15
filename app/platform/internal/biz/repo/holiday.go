package repo

import (
	"context"
	"time"

	"platform/internal/biz/model"
)

// HolidayClient 封装外部节假日数据源，返回与具体供应商无关的整年节假日配置。
type HolidayClient interface {
	FetchYear(ctx context.Context, year int) (map[string]*model.HolidayEntry, error)
}

// HolidayRepo 管理按自然日持久化的节假日状态。
type HolidayRepo interface {
	GetByDate(ctx context.Context, date time.Time) (*model.HolidayInfo, error)
	Upsert(ctx context.Context, info *model.HolidayInfo) (*model.HolidayInfo, error)
}

// HolidayConfigCacheRepo 将整年节假日配置缓存到外部存储，跨实例共享一份。
type HolidayConfigCacheRepo interface {
	Get(ctx context.Context, year int) (*model.HolidayConfigCache, error)
	// SetIfFresher 只在缓存缺失或已有数据更旧时写入。
	SetIfFresher(ctx context.Context, year int, cache *model.HolidayConfigCache) error
}
