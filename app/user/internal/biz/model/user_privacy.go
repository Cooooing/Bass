package model

type UserPrivacy struct {
	// 用户隐私设置ID
	ID int64
	// 用户ID
	UserID int64
	// 是否公开积分榜
	PublicPoints *bool
	// 是否公开粉丝列表
	PublicFollowers *bool
	// 是否公开帖子列表
	PublicArticles *bool
	// 是否公开评论列表
	PublicComments *bool
	// 是否公开在线状态
	PublicOnlineStatus *bool
	// 是否公开地理位置
	PublicLocation *bool
}
