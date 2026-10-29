package usecase

import (
	"context"
	"time"

	"common/pkg/apperror"
	commonclient "common/pkg/client"
	cerrors "common/proto/gen/common/errors"
	userv1 "common/proto/gen/user/v1"

	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/enum"
)

// TCCCheckinUsecase 执行自动签到的本地 TCC 分支。
type TCCCheckinUsecase struct {
	records repo.CheckinRecordRepo
	stats   repo.CheckinStatRepo
	branch  commonclient.TCCBranchExecutor
}

// NewTCCCheckinUsecase 创建自动签到 TCC 分支用例。
func NewTCCCheckinUsecase(records repo.CheckinRecordRepo, stats repo.CheckinStatRepo, branch commonclient.TCCBranchExecutor) *TCCCheckinUsecase {
	return &TCCCheckinUsecase{records: records, stats: stats, branch: branch}
}

// Try 记录待确认的签到事实。
func (u *TCCCheckinUsecase) Try(ctx context.Context, req *userv1.TCCCheckin_Req) error {
	return u.branch.Execute(ctx, func(ctx context.Context) error {
		if req.GetUserId() <= 0 || req.GetDate() == nil || req.GetCheckinPoints() <= 0 || req.GetStreakPoints() < 0 || req.GetCurrentStreak() <= 0 || req.GetLongestStreak() < req.GetCurrentStreak() {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		date := req.GetDate().AsTime()
		record, err := u.records.Get(ctx, &repo.CheckinRecordGetReq{UserID: new(req.UserId), Date: &date})
		if err != nil || (record != nil && record.Status != enum.CheckinStatusCanceled) {
			return err
		}
		checkinPoints := req.GetCheckinPoints()
		streakPoints := req.GetStreakPoints()
		_, err = u.records.UpsertRecord(ctx, &model.CheckinRecord{UserID: req.GetUserId(), Date: &date, CheckinPoints: &checkinPoints, StreakPoints: &streakPoints, Status: enum.CheckinStatusPending})
		return err
	})
}

// Confirm 标记签到完成并更新连续签到统计。
func (u *TCCCheckinUsecase) Confirm(ctx context.Context, req *userv1.TCCCheckin_Req) error {
	return u.branch.Execute(ctx, func(ctx context.Context) error {
		date := req.GetDate().AsTime()
		record, err := u.records.Get(ctx, &repo.CheckinRecordGetReq{UserID: new(req.UserId), Date: &date})
		if err != nil {
			return err
		}
		if record == nil || record.Status == enum.CheckinStatusCanceled {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_CONFLICT)
		}
		if record.Status == enum.CheckinStatusCompleted {
			return nil
		}
		checkinPoints := req.GetCheckinPoints()
		streakPoints := req.GetStreakPoints()
		now := time.Now()
		if _, err = u.records.UpsertRecord(ctx, &model.CheckinRecord{UserID: req.GetUserId(), Date: &date, CheckinPoints: &checkinPoints, StreakPoints: &streakPoints, Status: enum.CheckinStatusCompleted, CheckedAt: &now}); err != nil {
			return err
		}
		currentStreak := req.GetCurrentStreak()
		longestStreak := req.GetLongestStreak()
		_, err = u.stats.UpsertStat(ctx, &model.CheckinStat{UserID: req.GetUserId(), CurrentStreak: &currentStreak, LongestStreak: &longestStreak})
		return err
	})
}

// Cancel 标记签到取消，使后续任务重试可以重新编排事务。
func (u *TCCCheckinUsecase) Cancel(ctx context.Context, req *userv1.TCCCheckin_Req) error {
	return u.branch.Execute(ctx, func(ctx context.Context) error {
		date := req.GetDate().AsTime()
		record, err := u.records.Get(ctx, &repo.CheckinRecordGetReq{UserID: new(req.UserId), Date: &date})
		if err != nil {
			return err
		}
		if record == nil {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_CONFLICT)
		}
		if record.Status == enum.CheckinStatusCanceled || record.Status == enum.CheckinStatusCompleted {
			return nil
		}
		_, err = u.records.UpsertRecord(ctx, &model.CheckinRecord{UserID: req.GetUserId(), Date: &date, Status: enum.CheckinStatusCanceled})
		return err
	})
}
