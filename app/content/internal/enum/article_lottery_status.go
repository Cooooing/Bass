package enum

import (
	v1 "common/api/gen/content/v1"
	"common/pkg/enum"
)

type ArticleLotteryStatus string

const (
	ArticleLotteryStatusNotStarted ArticleLotteryStatus = "not_started"
	ArticleLotteryStatusInProgress ArticleLotteryStatus = "in_progress"
	ArticleLotteryStatusFinished   ArticleLotteryStatus = "finished"
)

var ArticleLotteryStatusMap = enum.NewMapping[ArticleLotteryStatus, v1.ArticleLotteryStatus](map[ArticleLotteryStatus]enum.Entry[ArticleLotteryStatus, v1.ArticleLotteryStatus]{
	ArticleLotteryStatusNotStarted: {Proto: v1.ArticleLotteryStatus_ARTICLE_LOTTERY_STATUS_NOT_STARTED},
	ArticleLotteryStatusInProgress: {Proto: v1.ArticleLotteryStatus_ARTICLE_LOTTERY_STATUS_IN_PROGRESS},
	ArticleLotteryStatusFinished:   {Proto: v1.ArticleLotteryStatus_ARTICLE_LOTTERY_STATUS_FINISHED},
})
