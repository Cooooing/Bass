package enum

import (
	v1 "common/api/gen/content/v1"
	"common/pkg/enum"
)

type ArticleAction string

const (
	ArticleActionLike    ArticleAction = "like"
	ArticleActionThank   ArticleAction = "thank"
	ArticleActionCollect ArticleAction = "collect"
	ArticleActionWatch   ArticleAction = "watch"
	ArticleActionReward  ArticleAction = "reward"
	ArticleActionReply   ArticleAction = "reply"
)

var ArticleActionMap = enum.NewMapping[ArticleAction, v1.ArticleAction](map[ArticleAction]enum.Entry[ArticleAction, v1.ArticleAction]{
	ArticleActionLike:    {Proto: v1.ArticleAction_ARTICLE_ACTION_LIKE},
	ArticleActionThank:   {Proto: v1.ArticleAction_ARTICLE_ACTION_THANK},
	ArticleActionCollect: {Proto: v1.ArticleAction_ARTICLE_ACTION_COLLECT},
	ArticleActionWatch:   {Proto: v1.ArticleAction_ARTICLE_ACTION_WATCH},
	ArticleActionReward:  {Proto: v1.ArticleAction_ARTICLE_ACTION_REWARD},
	ArticleActionReply:   {Proto: v1.ArticleAction_ARTICLE_ACTION_REPLY},
})
