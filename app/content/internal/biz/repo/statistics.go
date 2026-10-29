package repo

import (
	"context"
	"time"

	"content/internal/biz/model"
	"content/internal/enum"
)

// StatisticsRepo 批量读取内容领域的原始行为数量。
type StatisticsRepo interface {
	BatchCountUserBehaviors(context.Context, []int64, time.Time, time.Time, []enum.ContentActionStatus) ([]*model.UserBehaviorCount, error)
}
