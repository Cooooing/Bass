package repo

import (
	"context"

	"content/internal/biz/model"
)

// ArticlePollRepo 保存投票配置、选项及不可修改的投票记录。
type ArticlePollRepo interface {
	SaveDraft(context.Context, int64, *model.ArticlePollDraft) error
	DeleteDraft(context.Context, int64) error
	Get(context.Context, int64, int64) (*model.ArticlePoll, error)
	Vote(context.Context, int64, int64, []int64) error
}
