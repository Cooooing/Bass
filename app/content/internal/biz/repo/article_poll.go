package repo

import (
	"context"

	"content/internal/biz/model"
)

// ArticlePollRepo 保存投票配置、选项及不可修改的投票记录。
type ArticlePollRepo interface {
	SaveDraft(ctx context.Context, articleID int64, draft *model.ArticlePollDraft) error
	DeleteDraft(ctx context.Context, articleID int64) error
	Get(ctx context.Context, articleID, userID int64) (*model.ArticlePoll, error)
	Vote(ctx context.Context, articleID, userID int64, optionIDs []int64) error
	ListParticipantUserIDs(ctx context.Context, articleID int64) ([]int64, error)
}
