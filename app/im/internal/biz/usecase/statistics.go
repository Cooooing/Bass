package usecase

import (
	"context"
	"time"

	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"

	"im/internal/biz/model"
	"im/internal/biz/repo"
)

// StatisticsUsecase 提供即时消息领域的内部批量事实统计。
type StatisticsUsecase struct{ repo repo.StatisticsRepo }

// NewStatisticsUsecase 创建即时消息统计用例。
func NewStatisticsUsecase(repo repo.StatisticsRepo) *StatisticsUsecase {
	return &StatisticsUsecase{repo: repo}
}

// BatchCountGlobalChatMessages 返回指定账号在时间窗口内的有效聊天室文本消息数量。
func (u *StatisticsUsecase) BatchCountGlobalChatMessages(ctx context.Context, userIDs []int64, start, end time.Time) ([]*model.GlobalChatMessageCount, error) {
	if len(userIDs) == 0 || !end.After(start) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	return u.repo.BatchCountGlobalChatMessages(ctx, userIDs, start, end)
}
