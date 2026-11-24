package repo

import (
	"context"
	"time"

	"im/internal/biz/model"
)

// StatisticsRepo 批量读取即时消息领域的原始消息数量。
type StatisticsRepo interface {
	BatchCountGlobalChatMessages(ctx context.Context, userIDs []int64, start, end time.Time) ([]*model.GlobalChatMessageCount, error)
}
