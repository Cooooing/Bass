package model

type PrivacySetting struct {
	// ID 是隐私设置记录 ID。
	ID int64
	// UserID 是归属账号 ID。
	UserID int64
	// PublicPoints 控制积分是否公开。
	PublicPoints *bool
	// PublicFollowers 控制粉丝列表是否公开。
	PublicFollowers *bool
	// PublicArticles 控制文章列表是否公开。
	PublicArticles *bool
	// PublicComments 控制评论列表是否公开。
	PublicComments *bool
	// PublicOnlineStatus 控制在线状态是否公开。
	PublicOnlineStatus *bool
	// PublicLocation 控制位置是否公开。
	PublicLocation *bool
}
