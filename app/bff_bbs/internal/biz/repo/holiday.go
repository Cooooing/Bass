package repo

import (
	"context"

	"bff_bbs/internal/biz/model"
)

// HolidayClient 从 Platform 读取节假日状态。
type HolidayClient interface {
	GetCurrent(ctx context.Context) (*model.HolidayInfo, error)
}
