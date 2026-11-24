package repo

import (
	"context"
	"time"

	"im/internal/biz/model"
)

// EmojiGroupRepo 存取表情分组及其表情项事实。
type EmojiGroupRepo interface {
	List(ctx context.Context, userID int64, keyword string) ([]*model.EmojiGroup, error)
	Get(ctx context.Context, userID, groupID int64) (*model.EmojiGroup, error)
	GetDefault(ctx context.Context, userID int64) (*model.EmojiGroup, error)
	Create(ctx context.Context, group *model.EmojiGroup) (*model.EmojiGroup, error)
	Rename(ctx context.Context, userID, groupID int64, name string) (*model.EmojiGroup, error)
	Delete(ctx context.Context, userID, groupID int64) error
	CreateItem(ctx context.Context, item *model.EmojiItem) (*model.EmojiItem, error)
	RenameItem(ctx context.Context, userID, itemID int64, name string) (*model.EmojiItem, error)
	ReorderGroups(ctx context.Context, userID int64, groupIDs []int64) error
	ReorderItems(ctx context.Context, userID int64, groupID *int64, itemIDs []int64) error
	DeleteItem(ctx context.Context, userID, itemID int64) error
	MoveItem(ctx context.Context, userID, itemID, groupID int64) (*model.EmojiItem, error)
}

// EmojiShareCache 保存短期分享快照，不能作为表情领域事实来源。
type EmojiShareCache interface {
	Save(ctx context.Context, token string, share *model.EmojiShare, ttl time.Duration) error
	Get(ctx context.Context, token string) (*model.EmojiShare, error)
}

// EmojiAssetClient 校验分组引用的平台资产仍可用。
type EmojiAssetClient interface {
	ValidateAvailable(ctx context.Context, assetID int64) error
}
