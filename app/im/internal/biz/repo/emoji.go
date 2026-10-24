package repo

import (
	"context"
	"time"

	"im/internal/biz/model"
)

// EmojiGroupRepo 存取表情分组及其表情项事实。
type EmojiGroupRepo interface {
	List(context.Context, int64, string) ([]*model.EmojiGroup, error)
	Get(context.Context, int64, int64) (*model.EmojiGroup, error)
	GetDefault(context.Context, int64) (*model.EmojiGroup, error)
	Create(context.Context, *model.EmojiGroup) (*model.EmojiGroup, error)
	Rename(context.Context, int64, int64, string) (*model.EmojiGroup, error)
	Delete(context.Context, int64, int64) error
	CreateItem(context.Context, *model.EmojiItem) (*model.EmojiItem, error)
	RenameItem(context.Context, int64, int64, string) (*model.EmojiItem, error)
	ReorderGroups(context.Context, int64, []int64) error
	ReorderItems(context.Context, int64, *int64, []int64) error
	DeleteItem(context.Context, int64, int64) error
	MoveItem(context.Context, int64, int64, int64) (*model.EmojiItem, error)
}

// EmojiShareCache 保存短期分享快照，不能作为表情领域事实来源。
type EmojiShareCache interface {
	Save(context.Context, string, *model.EmojiShare, time.Duration) error
	Get(context.Context, string) (*model.EmojiShare, error)
}

// EmojiAssetClient 校验分组引用的平台资产仍可用。
type EmojiAssetClient interface {
	ValidateAvailable(context.Context, int64) error
}
