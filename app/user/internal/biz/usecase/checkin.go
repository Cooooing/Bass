package usecase

import (
	"context"
	"fmt"
	"math/rand/v2"
	"time"

	"common/pkg/apperror"
	commonclient "common/pkg/client"
	"common/pkg/constant"
	cerrors "common/proto/gen/common/errors"
	economyv1 "common/proto/gen/economy/v1"
	economyenum "common/proto/gen/economy/v1/enum"
	userv1 "common/proto/gen/user/v1"

	"user/internal/biz/repo"
	"user/internal/config"
	"user/internal/enum"

	"github.com/dtm-labs/client/dtmgrpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// CheckinUsecase 维护自动签到、连签状态及其积分发放。
type CheckinUsecase struct {
	records  repo.CheckinRecordRepo
	stats    repo.CheckinStatRepo
	dtm      *commonclient.DTMClient
	branches *commonclient.TCCBranchCaller
	business *config.Business
}

// NewCheckinUsecase 组装自动签到所需的仓储与积分发放依赖。
func NewCheckinUsecase(records repo.CheckinRecordRepo, stats repo.CheckinStatRepo, dtm *commonclient.DTMClient, branches *commonclient.TCCBranchCaller, business *config.Business) *CheckinUsecase {
	return &CheckinUsecase{records: records, stats: stats, dtm: dtm, branches: branches, business: business}
}

// Auto 为达到活跃阈值的账号执行指定业务日的自动签到。
func (u *CheckinUsecase) Auto(ctx context.Context, userID int64, date, previousDate time.Time) error {
	rule := u.business.GetLiveness().GetCheckin()
	if rule == nil || rule.GetRewardMinimum() <= 0 || rule.GetRewardMaximum() < rule.GetRewardMinimum() || rule.GetStreakInterval() <= 0 || rule.GetStreakReward() < 0 {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	points := rule.GetRewardMinimum() + rand.Int64N(rule.GetRewardMaximum()-rule.GetRewardMinimum()+1)
	record, err := u.records.Get(ctx, &repo.CheckinRecordGetReq{UserID: &userID, Date: &date})
	if err != nil || (record != nil && record.Status != enum.CheckinStatusCanceled) {
		return err
	}
	if record != nil && record.CheckinPoints != nil {
		points = *record.CheckinPoints
	}
	previous, err := u.records.Get(ctx, &repo.CheckinRecordGetReq{UserID: &userID, Date: &previousDate})
	if err != nil {
		return err
	}
	stat, err := u.stats.Get(ctx, &repo.CheckinStatGetReq{UserID: &userID})
	if err != nil {
		return err
	}
	streak, longest := int32(1), int32(1)
	if stat != nil {
		longest = *stat.LongestStreak
		if previous != nil && previous.Status == enum.CheckinStatusCompleted {
			streak = *stat.CurrentStreak + 1
		}
	}
	longest = max(streak, longest)
	streakPoints := int64(0)
	if streak%rule.GetStreakInterval() == 0 {
		streakPoints = rule.GetStreakReward()
	}
	request := &userv1.TCCCheckin_Req{UserId: userID, Date: timestamppb.New(date), CheckinPoints: points, StreakPoints: streakPoints, CurrentStreak: streak, LongestStreak: longest}
	_, err = u.dtm.InTCC(ctx, func(tcc *dtmgrpc.TccGrpc) error {
		if err := u.branches.Call(ctx, tcc, constant.UserServiceName.String(), request, userv1.TCCCheckinService_Try_FullMethodName, userv1.TCCCheckinService_Confirm_FullMethodName, userv1.TCCCheckinService_Cancel_FullMethodName, &userv1.TCCCheckin_Resp{}); err != nil {
			return err
		}
		checkinReward := &economyv1.TCCPoints_Req{UserId: userID, Amount: points, RecordType: economyenum.EconomyRecordType_ECONOMY_RECORD_TYPE_SIGN_IN_REWARD, IdempotencyKey: new(fmt.Sprintf("checkin:%s", date.Format(time.DateOnly)))}
		if err := u.branches.Call(ctx, tcc, constant.EconomyServiceName.String(), checkinReward, economyv1.TCCPointsService_Try_FullMethodName, economyv1.TCCPointsService_Confirm_FullMethodName, economyv1.TCCPointsService_Cancel_FullMethodName, &economyv1.TCCPoints_Resp{}); err != nil {
			return err
		}
		if streakPoints == 0 {
			return nil
		}
		streakReward := &economyv1.TCCPoints_Req{UserId: userID, Amount: streakPoints, RecordType: economyenum.EconomyRecordType_ECONOMY_RECORD_TYPE_CHECKIN_STREAK_REWARD, IdempotencyKey: new(fmt.Sprintf("checkin_streak:%s", date.Format(time.DateOnly)))}
		return u.branches.Call(ctx, tcc, constant.EconomyServiceName.String(), streakReward, economyv1.TCCPointsService_Try_FullMethodName, economyv1.TCCPointsService_Confirm_FullMethodName, economyv1.TCCPointsService_Cancel_FullMethodName, &economyv1.TCCPoints_Resp{})
	})
	return err
}
