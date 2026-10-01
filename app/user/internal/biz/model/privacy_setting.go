package model

type PrivacySetting struct {
	// ID 是隐私设置记录 ID。
	ID int64
	// UserID 是归属账号 ID。
	UserID int64
	// PublicPoints 控制积分是否公开。
	PublicPoints *bool
	// PublicFollowerList 控制粉丝列表是否公开。
	PublicFollowerList *bool
	// PublicFollowingList 控制关注列表是否公开。
	PublicFollowingList *bool
	// PublicArticleList 控制帖子列表是否公开。
	PublicArticleList *bool
	// PublicCommentList 控制回复列表是否公开。
	PublicCommentList *bool
	// PublicOnlineStatus 控制在线状态是否公开。
	PublicOnlineStatus *bool
	// PublicLocation 控制位置是否公开。
	PublicLocation *bool
	// PublicMoonbreezeList 控制清风明月列表是否公开。
	PublicMoonbreezeList *bool
}
