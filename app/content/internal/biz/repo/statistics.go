package repo

import (
	"context"
	"time"

	"content/internal/biz/model"
	"content/internal/enum"
)

// StatisticsRepo 批量读取内容领域的原始行为数量。
type StatisticsRepo interface {
	BatchCountUserBehaviors(ctx context.Context, userIDs []int64, start, end time.Time, statuses []enum.ContentActionStatus) ([]*model.UserBehaviorCount, error)
}
