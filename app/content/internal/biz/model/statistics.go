package model

// UserBehaviorCount 是一个账号在内容领域时间窗口内的原始行为数量。
type UserBehaviorCount struct {
	UserID             int64
	ArticleCount       int32
	CommentCount       int32
	ArticleLikeCount   int32
	CommentLikeCount   int32
	ArticleThankCount  int32
	CommentThankCount  int32
	ArticleRewardCount int32
	ViewCount          int32
}
