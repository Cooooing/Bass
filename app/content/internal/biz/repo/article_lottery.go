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
	SaveDraft(ctx context.Context, articleID int64, draft *model.ArticleLotteryDraft) error
	DeleteDraft(ctx context.Context, articleID int64) error
	Get(ctx context.Context, articleID, userID int64) (*model.ArticleLottery, error)
	SetEntry(ctx context.Context, articleID, userID int64, joined bool, now time.Time) error
	Draw(ctx context.Context, articleID int64, now time.Time) ([]*ArticleLotteryWinner, error)
	ListParticipantUserIDs(ctx context.Context, articleID int64) ([]int64, error)
}

// ArticleLotteryWinner 是一次开奖产生的中奖记录。
type ArticleLotteryWinner struct {
	PrizeID int64
	UserID  int64
}
