package usecase

import (
	"common/pkg/apperror"
	commonclient "common/pkg/client"
	"common/pkg/constant"
	commonmodel "common/pkg/model"
	cerrors "common/proto/gen/common/errors"
	economyv1 "common/proto/gen/economy/v1"
	economyenum "common/proto/gen/economy/v1/enum"
	userv1 "common/proto/gen/user/v1"
	"context"
	"fmt"
	"time"

	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/config"
	"user/internal/enum"

	"github.com/dtm-labs/client/dtmgrpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// LivenessUsecase 聚合活跃度并结算每日快照。
type LivenessUsecase struct {
	source   repo.LivenessSource
	records  repo.LivenessRecordRepo
	accounts repo.AccountRepo
	checkin  *CheckinUsecase
	dtm      *commonclient.DTMClient
	branches *commonclient.TCCBranchCaller
	business *config.Business
}

// NewLivenessUsecase 组装活跃度查询、日结和奖励领取所需的依赖。
func NewLivenessUsecase(source repo.LivenessSource, records repo.LivenessRecordRepo, accounts repo.AccountRepo, checkin *CheckinUsecase, dtm *commonclient.DTMClient, branches *commonclient.TCCBranchCaller, business *config.Business) *LivenessUsecase {
	return &LivenessUsecase{source: source, records: records, accounts: accounts, checkin: checkin, dtm: dtm, branches: branches, business: business}
}

// livenessWindow 将任意时间规整为本机时区下的业务日半开区间 [start, end)。
func livenessWindow(date time.Time) (time.Time, time.Time) {
	local := date.In(time.Local)
	start := time.Date(local.Year(), local.Month(), local.Day(), 0, 0, 0, 0, time.Local)
	return start, start.AddDate(0, 0, 1)
}

// livenessRule 返回已校验的活跃度规则。
func (u *LivenessUsecase) livenessRule() (*config.Business_Liveness, error) {
	rule := u.business.GetLiveness()
	if rule == nil || rule.GetScore() == nil || rule.GetCheckin() == nil || rule.GetScore().GetMaximum() <= 0 || rule.GetCheckin().GetMinimumScore() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	for _, item := range []*config.Business_Liveness_Rule{rule.GetScore().GetArticle(), rule.GetScore().GetComment(), rule.GetScore().GetLike(), rule.GetScore().GetThank(), rule.GetScore().GetReward(), rule.GetScore().GetView(), rule.GetScore().GetGlobalChatMessage()} {
		if item == nil || item.GetLimit() < 0 || item.GetCountPerUnit() <= 0 || item.GetPointsPerUnit() < 0 {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
	}
	return rule, nil
}

func livenessScore(rule *config.Business_Liveness_Rule, count int32) int32 {
	if rule.GetLimit() > 0 {
		count = min(count, rule.GetLimit())
	}
	return count / rule.GetCountPerUnit() * rule.GetPointsPerUnit()
}

// score 按配置将原始行为计数转换为单日活跃度。
func (u *LivenessUsecase) score(detail *model.LivenessDetail) (int32, error) {
	rule, err := u.livenessRule()
	if err != nil {
		return 0, err
	}
	score := rule.GetScore()
	article := livenessScore(score.GetArticle(), detail.ArticleCount)
	comment := livenessScore(score.GetComment(), detail.CommentCount)
	like := livenessScore(score.GetLike(), detail.LikeCount)
	thank := livenessScore(score.GetThank(), detail.ThankCount)
	reward := livenessScore(score.GetReward(), detail.RewardCount)
	view := livenessScore(score.GetView(), detail.ViewCount)
	chat := livenessScore(score.GetGlobalChatMessage(), detail.GlobalChatMessageCount)
	return min(article+comment+like+thank+reward+view+chat, score.GetMaximum()), nil
}

// CheckinScore 返回当前自动签到所需的最低活跃度。
func (u *LivenessUsecase) CheckinScore() (int32, error) {
	rule, err := u.livenessRule()
	if err != nil {
		return 0, err
	}
	return rule.GetCheckin().GetMinimumScore(), nil
}

// MaximumScore 返回当前业务日允许累积的活跃度上限。
func (u *LivenessUsecase) MaximumScore() (int32, error) {
	rule, err := u.livenessRule()
	if err != nil {
		return 0, err
	}
	return rule.GetScore().GetMaximum(), nil
}

// Get 查询指定账号和业务日的活跃度。
func (u *LivenessUsecase) Get(ctx context.Context, userID int64, date time.Time) (int32, *model.LivenessDetail, *model.LivenessRecord, error) {
	if userID <= 0 {
		return 0, nil, nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	start, end := livenessWindow(date)
	record, err := u.records.Get(ctx, userID, start)
	if err != nil {
		return 0, nil, nil, err
	}
	if record != nil {
		return record.Score, nil, record, nil
	}
	details, err := u.source.BatchGet(ctx, []int64{userID}, start, end)
	if err != nil {
		return 0, nil, nil, err
	}
	detail := details[userID]
	score, err := u.score(detail)
	if err != nil {
		return 0, nil, nil, err
	}
	return score, detail, nil, nil
}

// Claim 领取指定业务日的活跃奖励。
func (u *LivenessUsecase) Claim(ctx context.Context, userID int64, date time.Time) (*model.LivenessRecord, error) {
	start, _ := livenessWindow(date)
	row, err := u.records.Get(ctx, userID, start)
	if err != nil {
		return nil, err
	}
	if row == nil || row.Score <= 0 || (row.RewardStatus != enum.LivenessRewardStatusUnclaimed && row.RewardStatus != enum.LivenessRewardStatusCanceled) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_CONFLICT)
	}
	request := &userv1.TCCLivenessReward_Req{UserId: userID, Date: timestamppb.New(start)}
	reward := &economyv1.TCCPoints_Req{UserId: userID, Amount: int64(row.Score), RecordType: economyenum.EconomyRecordType_ECONOMY_RECORD_TYPE_LIVENESS_REWARD, Remark: new(fmt.Sprintf("liveness:%s", start.Format(time.DateOnly)))}
	_, err = u.dtm.InTCC(ctx, func(tcc *dtmgrpc.TccGrpc) error {
		if err := u.branches.Call(ctx, tcc, constant.UserServiceName.String(), request, userv1.TCCLivenessRewardService_Try_FullMethodName, userv1.TCCLivenessRewardService_Confirm_FullMethodName, userv1.TCCLivenessRewardService_Cancel_FullMethodName, &userv1.TCCLivenessReward_Resp{}); err != nil {
			return err
		}
		return u.branches.Call(ctx, tcc, constant.EconomyServiceName.String(), reward, economyv1.TCCPointsService_Try_FullMethodName, economyv1.TCCPointsService_Confirm_FullMethodName, economyv1.TCCPointsService_Cancel_FullMethodName, &economyv1.TCCPoints_Resp{})
	})
	if err != nil {
		return nil, err
	}
	return u.records.Get(ctx, userID, start)
}

// Settle 按页顺序结算指定业务日的活跃度。
func (u *LivenessUsecase) Settle(ctx context.Context, date time.Time, pageSize int32) (int32, int32, error) {
	if pageSize <= 0 || pageSize > 1000 {
		return 0, 0, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	start, end := livenessWindow(date)
	normal := enum.AccountStatusNormal
	page := int64(1)
	rule, err := u.livenessRule()
	if err != nil {
		return 0, 0, err
	}
	previousCheckinDate := start.AddDate(0, 0, -1)
	var settled, checked int32
	for {
		// todo 用户后续可以使用其他方式减少查询数量
		rows, err := u.accounts.Page(ctx, &repo.AccountPageReq{Page: commonmodel.PageReq{Page: page, Size: int64(pageSize)}, Query: repo.AccountGetReq{Status: &normal}})
		if err != nil {
			return 0, 0, err
		}

		userIDs := make([]int64, 0, len(rows.Rows))
		for _, account := range rows.Rows {
			userIDs = append(userIDs, account.ID)
		}
		if len(userIDs) == 0 {
			break
		}
		details, err := u.source.BatchGet(ctx, userIDs, start, end)
		if err != nil {
			return 0, 0, err
		}
		for _, account := range rows.Rows {
			record, err := u.records.Get(ctx, account.ID, start)
			if err != nil {
				return 0, 0, err
			}
			if record != nil {
				continue
			}
			score, err := u.score(details[account.ID])
			if err != nil {
				return 0, 0, err
			}
			if score == 0 {
				continue
			}
			if _, err = u.records.Save(ctx, &model.LivenessRecord{UserID: account.ID, Date: &start, Score: score, RewardStatus: enum.LivenessRewardStatusUnclaimed}); err != nil {
				return 0, 0, err
			}
			settled++
			if score >= rule.GetCheckin().GetMinimumScore() {
				if err = u.checkin.Auto(ctx, account.ID, start, previousCheckinDate); err != nil {
					return 0, 0, err
				}
				checked++
			}
		}
		if len(rows.Rows) < int(pageSize) {
			break
		}
		page++
	}
	return settled, checked, nil
}
