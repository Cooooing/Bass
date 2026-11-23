package repo

import (
	"context"
	"time"

	"user/internal/biz/model"
)

// LivenessSource 从其他领域读取指定账号在业务日内的有效活跃事实。
type LivenessSource interface {
	BatchGet(ctx context.Context, userIDs []int64, start, end time.Time) (map[int64]*model.LivenessDetail, error)
}
