package model

import (
	"strings"
	"time"

	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
)

// ArticleLotteryDraft 是文章草稿期可编辑的抽奖配置。
type ArticleLotteryDraft struct {
	Prizes []*ArticleLotteryPrizeDraft
	DrawAt time.Time
}

// ArticleLotteryPrizeDraft 是一个奖项的草稿配置。
type ArticleLotteryPrizeDraft struct {
	Name        string
	Description *string
	WinnerLimit int32
}

// ArticleLottery 是文章详情中的抽奖事实和当前读者报名状态。
type ArticleLottery struct {
	ID         int64
	DrawAt     time.Time
	DrawnAt    *time.Time
	EntryCount int64
	Entered    bool
	Prizes     []*ArticleLotteryPrize
}

// ArticleLotteryPrize 是已发布抽奖的一档奖项及其中奖者。
type ArticleLotteryPrize struct {
	ID            int64
	Name          string
	Description   *string
	WinnerLimit   int32
	WinnerUserIDs []int64
}

func (d *ArticleLotteryDraft) Normalize(now time.Time) error {
	if d == nil || len(d.Prizes) == 0 || len(d.Prizes) > 10 {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if !d.DrawAt.After(now.Add(5*time.Minute)) || d.DrawAt.After(now.AddDate(0, 3, 0)) {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_SCHEDULED_AT_OUT_OF_RANGE)
	}
	for _, prize := range d.Prizes {
		if prize == nil || prize.WinnerLimit < 1 || prize.WinnerLimit > 100 {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		prize.Name = strings.TrimSpace(prize.Name)
		if prize.Name == "" {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		if prize.Description != nil {
			description := strings.TrimSpace(*prize.Description)
			prize.Description = &description
			if description == "" {
				prize.Description = nil
			}
		}
	}
	d.DrawAt = d.DrawAt.UTC().Truncate(time.Second)
	return nil
}
