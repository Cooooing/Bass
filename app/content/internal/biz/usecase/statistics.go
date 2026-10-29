package usecase

import (
	"context"
	"time"

	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"

	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/enum"
)

// StatisticsUsecase 提供内容领域的内部批量事实统计。
type StatisticsUsecase struct{ repo repo.StatisticsRepo }

// NewStatisticsUsecase 创建内容统计用例。
func NewStatisticsUsecase(repo repo.StatisticsRepo) *StatisticsUsecase {
	return &StatisticsUsecase{repo: repo}
}

// BatchCountUserBehaviors 返回指定账号在时间窗口内的内容原始行为数量。
func (u *StatisticsUsecase) BatchCountUserBehaviors(ctx context.Context, userIDs []int64, start, end time.Time, actionStatuses []enum.ContentActionStatus) ([]*model.UserBehaviorCount, error) {
	if len(userIDs) == 0 || !end.After(start) || len(actionStatuses) == 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	return u.repo.BatchCountUserBehaviors(ctx, userIDs, start, end, actionStatuses)
}
