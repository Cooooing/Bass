package model

import "time"

// EmojiGroup 是表情分组的领域展示模型。
type EmojiGroup struct {
	ID        int64
	UserID    int64
	Name      string
	Sort      int64
	CreatedAt *time.Time
	Items     []*EmojiItem
}

// EmojiItem 是表情分组中的平台资产引用。
type EmojiItem struct {
	ID      int64
	UserID  int64
	GroupID *int64
	AssetID int64
	Name    string
	Sort    int64
}

// EmojiShare 是可短期导入的表情分组快照。
type EmojiShare struct {
	OwnerUserID int64        `json:"owner_user_id"`
	GroupName   string       `json:"group_name"`
	Items       []*EmojiItem `json:"items"`
}
