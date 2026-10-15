package repo

import (
	"context"
	"time"

	"content/internal/biz/model"
)

// ArticleLotteryClosedError 表示开奖已结束或报名窗口关闭。
type ArticleLotteryClosedError struct{}

func (*ArticleLotteryClosedError) Error() string {
	return "article lottery is closed"
}

// ArticleLotteryRepo 保存抽奖配置、报名和开奖结果。
type ArticleLotteryRepo interface {
	SaveDraft(context.Context, int64, *model.ArticleLotteryDraft) error
	DeleteDraft(context.Context, int64) error
	Get(context.Context, int64, int64) (*model.ArticleLottery, error)
	SetEntry(context.Context, int64, int64, bool, time.Time) error
	Draw(context.Context, int64, time.Time) ([]*ArticleLotteryWinner, error)
}

// ArticleLotteryWinner 是一次开奖产生的中奖记录。
type ArticleLotteryWinner struct {
	PrizeID int64
	UserID  int64
}
