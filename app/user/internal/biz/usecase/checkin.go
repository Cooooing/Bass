package usecase

import (
	"common/pkg/apperror"
	commonenums "common/proto/gen/common/enums"
	cerrors "common/proto/gen/common/errors"
	"context"
	"time"
	"user/internal/biz/base"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/enum"
)

type CheckinUsecase struct {
	tx              base.Tx
	accountRepo     repo.AccountRepo
	preferencesRepo repo.PreferencesRepo
	recordRepo      repo.CheckinRecordRepo
	statRepo        repo.CheckinStatRepo
	outboxRepo      repo.OutboxEventRepo
	outboxUsecase   *OutboxUsecase
}

func NewCheckinUsecase(
	tx base.Tx,
	accountRepo repo.AccountRepo,
	preferencesRepo repo.PreferencesRepo,
	recordRepo repo.CheckinRecordRepo,
	statRepo repo.CheckinStatRepo,
	outboxRepo repo.OutboxEventRepo,
	outboxUsecase *OutboxUsecase,
) *CheckinUsecase {
	return &CheckinUsecase{
		tx:              tx,
		accountRepo:     accountRepo,
		preferencesRepo: preferencesRepo,
		recordRepo:      recordRepo,
		statRepo:        statRepo,
		outboxRepo:      outboxRepo,
		outboxUsecase:   outboxUsecase,
	}
}

type CheckInReq struct {
	UserID int64
}

type CheckInResp struct {
	Record *model.CheckinRecord
	Stat   *model.CheckinStat
}

func (u *CheckinUsecase) CheckIn(ctx context.Context, req *CheckInReq) (*CheckInResp, error) {
	if req.UserID <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	preferences, err := u.preferencesRepo.Get(ctx, &repo.PreferencesGetReq{UserID: &req.UserID})
	if err != nil {
		return nil, err
	}
	location := time.FixedZone("Asia/Shanghai", 8*60*60)
	if preferences != nil && preferences.Timezone != nil {
		if value, loadErr := time.LoadLocation(*preferences.Timezone); loadErr == nil {
			location = value
		}
	}
	now := time.Now().In(location)
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	yesterday := today.AddDate(0, 0, -1)
	var record *model.CheckinRecord
	var stat *model.CheckinStat
	var outboxEvent *model.OutboxEvent
	err = u.tx(ctx, func(ctx context.Context) error {
		account, err := u.accountRepo.Get(ctx, &repo.AccountGetReq{UserID: &req.UserID})
		if err != nil {
			return err
		}
		if account == nil || account.Status == nil || *account.Status != enum.AccountStatusNormal {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_ACCOUNT_NOT_FOUND)
		}
		existing, err := u.recordRepo.Get(ctx, &repo.CheckinRecordGetReq{
			UserID: &req.UserID,
			Date:   &today,
		})
		if err != nil {
			return err
		}
		if existing != nil && existing.Checked {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_CHECKIN_ALREADY_COMPLETED)
		}
		record, err = u.recordRepo.UpsertRecord(ctx, &model.CheckinRecord{
			UserID:  req.UserID,
			Date:    &today,
			Checked: true,
		})
		if err != nil {
			return err
		}
		previous, err := u.recordRepo.Get(ctx, &repo.CheckinRecordGetReq{UserID: &req.UserID, Date: &yesterday})
		if err != nil {
			return err
		}
		stat, err = u.statRepo.Get(ctx, &repo.CheckinStatGetReq{UserID: &req.UserID})
		if err != nil {
			return err
		}
		currentStreak := int32(1)
		longestStreak := int32(1)
		if stat != nil {
			if stat.LongestStreak != nil {
				longestStreak = *stat.LongestStreak
			}
			if previous != nil && previous.Checked && stat.CurrentStreak != nil {
				currentStreak = *stat.CurrentStreak + 1
			}
		}
		if currentStreak > longestStreak {
			longestStreak = currentStreak
		}
		stat, err = u.statRepo.UpsertStat(ctx, &model.CheckinStat{
			UserID:        req.UserID,
			CurrentStreak: &currentStreak,
			LongestStreak: &longestStreak,
		})
		if err != nil {
			return err
		}
		outboxEvent, err = u.outboxRepo.Save(ctx, &repo.OutboxEventSave{
			Event: &commonenums.Event{
				Type:    commonenums.EventType_EVENT_TYPE_USER_CHECKIN_COMPLETED,
				Subject: commonenums.EventSubject_EVENT_SUBJECT_USER_CHECKIN_COMPLETED,
				Payload: &commonenums.Event_UserCheckinCompleted{
					UserCheckinCompleted: &commonenums.UserCheckinCompletedPayload{
						UserId:          req.UserID,
						CheckinRecordId: record.ID,
						LocalDate:       now.Format(time.DateOnly),
						CurrentStreak:   currentStreak,
					},
				},
			},
		})
		return err
	})
	if err != nil {
		return nil, err
	}
	if outboxEvent != nil {
		_, _ = u.outboxUsecase.Publish(ctx, &PublishOutboxEventReq{ID: outboxEvent.ID})
	}
	return &CheckInResp{Record: record, Stat: stat}, nil
}

type GetCheckinOverviewReq struct {
	UserID int64
	Month  string
}

type GetCheckinOverviewResp struct {
	Records []*model.CheckinRecord
	Stat    *model.CheckinStat
}

func (u *CheckinUsecase) GetOverview(ctx context.Context, req *GetCheckinOverviewReq) (*GetCheckinOverviewResp, error) {
	if req.UserID <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	month, err := time.Parse("2006-01", req.Month)
	if err != nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	records, err := u.recordRepo.List(ctx, &repo.CheckinRecordGetReq{UserID: &req.UserID})
	if err != nil {
		return nil, err
	}
	filtered := make([]*model.CheckinRecord, 0, len(records))
	for _, record := range records {
		if record.Date != nil && record.Date.Year() == month.Year() && record.Date.Month() == month.Month() {
			filtered = append(filtered, record)
		}
	}
	stat, err := u.statRepo.Get(ctx, &repo.CheckinStatGetReq{UserID: &req.UserID})
	if err != nil {
		return nil, err
	}
	if stat == nil {
		zero := int32(0)
		stat = &model.CheckinStat{UserID: req.UserID, CurrentStreak: &zero, LongestStreak: &zero}
	}
	return &GetCheckinOverviewResp{Records: filtered, Stat: stat}, nil
}
