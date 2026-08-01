package enum

import (
	"common/pkg/enum"
	v1 "common/proto/gen/content/v1/enum"
)

// ArticleAction 表示文章互动行为类型。
type ArticleAction string

const (
	// ArticleActionLike 表示点赞文章。
	ArticleActionLike ArticleAction = "like"
	// ArticleActionThank 表示感谢文章。
	ArticleActionThank ArticleAction = "thank"
	// ArticleActionCollect 表示收藏文章。
	ArticleActionCollect ArticleAction = "collect"
	// ArticleActionReward 表示打赏文章。
	ArticleActionReward ArticleAction = "reward"
	// ArticleActionReply 表示回复文章。
	ArticleActionReply ArticleAction = "reply"
)

// ArticleActionMap 维护文章行为内部枚举与 proto 枚举的映射。
var ArticleActionMap = enum.NewMapping[ArticleAction, v1.ArticleAction](map[ArticleAction]enum.Entry[ArticleAction, v1.ArticleAction]{
	ArticleActionLike:    {Proto: v1.ArticleAction_ARTICLE_ACTION_LIKE},
	ArticleActionThank:   {Proto: v1.ArticleAction_ARTICLE_ACTION_THANK},
	ArticleActionCollect: {Proto: v1.ArticleAction_ARTICLE_ACTION_COLLECT},
	ArticleActionReward:  {Proto: v1.ArticleAction_ARTICLE_ACTION_REWARD},
	ArticleActionReply:   {Proto: v1.ArticleAction_ARTICLE_ACTION_REPLY},
})
