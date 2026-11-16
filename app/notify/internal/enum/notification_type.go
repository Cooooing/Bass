package enum

import (
	commonenum "common/pkg/enum"
	notifyproto "common/proto/gen/notify/v1/enum"
)

// NotificationType identifies the notification template and delivery kind.
type NotificationType string

const (
	NotificationTypeUserFollow                            NotificationType = "user_follow"                               // 关注通知
	NotificationTypeUserRegister                          NotificationType = "user_register"                             // 注册通知
	NotificationTypeUserEmailVerificationCode             NotificationType = "user_email_verification_code"              // 邮箱验证码通知
	NotificationTypeUserPhoneVerificationCode             NotificationType = "user_phone_verification_code"              // 手机验证码通知
	NotificationTypeContentArticlePublish                 NotificationType = "content_article_publish"                   // 文章发布通知
	NotificationTypeContentArticleLike                    NotificationType = "content_article_like"                      // 文章点赞通知
	NotificationTypeContentArticleThank                   NotificationType = "content_article_thank"                     // 文章感谢通知
	NotificationTypeContentArticleCollect                 NotificationType = "content_article_collect"                   // 文章收藏通知
	NotificationTypeContentArticleWatch                   NotificationType = "content_article_watch"                     // 文章关注通知
	NotificationTypeContentArticleAcceptAnswer            NotificationType = "content_article_accept_answer"             // 回答采纳通知
	NotificationTypeContentArticleLotteryWon              NotificationType = "content_article_lottery_won"               // 抽奖中奖通知
	NotificationTypeContentArticleWatchedCommentPublished NotificationType = "content_article_watched_comment_published" // 关注文章评论通知
	NotificationTypeContentArticleWatchedPostscriptAdded  NotificationType = "content_article_watched_postscript_added"  // 关注文章附言通知
	NotificationTypeContentArticleWatchedAnswerAccepted   NotificationType = "content_article_watched_answer_accepted"   // 关注文章采纳通知
	NotificationTypeContentArticleWatchedPollEnded        NotificationType = "content_article_watched_poll_ended"        // 关注文章投票结束通知
	NotificationTypeContentArticleWatchedLotteryDrawn     NotificationType = "content_article_watched_lottery_drawn"     // 关注文章开奖通知
	NotificationTypeContentArticleMention                 NotificationType = "content_article_mention"                   // 文章提及通知
	NotificationTypeContentPostscriptMention              NotificationType = "content_postscript_mention"                // 附言提及通知
	NotificationTypeContentArticleReward                  NotificationType = "content_article_reward"                    // 文章打赏通知
	NotificationTypeContentCommentPublish                 NotificationType = "content_comment_publish"                   // 评论发布通知
	NotificationTypeContentCommentLike                    NotificationType = "content_comment_like"                      // 评论点赞通知
	NotificationTypeContentCommentThank                   NotificationType = "content_comment_thank"                     // 评论感谢通知
	NotificationTypeContentCommentMention                 NotificationType = "content_comment_mention"                   // 评论提及通知
	NotificationTypeIMGlobalChatMention                   NotificationType = "im_global_chat_mention"                    // 全局聊天提及通知
)

func (e NotificationType) String() string {
	return string(e)
}

var NotificationTypeMap = commonenum.NewMapping[NotificationType, notifyproto.NotificationType](map[NotificationType]commonenum.Entry[NotificationType, notifyproto.NotificationType]{
	NotificationTypeUserFollow:                            {Proto: notifyproto.NotificationType_NOTIFICATION_TYPE_USER_FOLLOW},
	NotificationTypeUserRegister:                          {Proto: notifyproto.NotificationType_NOTIFICATION_TYPE_USER_REGISTER},
	NotificationTypeUserEmailVerificationCode:             {Proto: notifyproto.NotificationType_NOTIFICATION_TYPE_USER_EMAIL_VERIFICATION_CODE},
	NotificationTypeUserPhoneVerificationCode:             {Proto: notifyproto.NotificationType_NOTIFICATION_TYPE_USER_PHONE_VERIFICATION_CODE},
	NotificationTypeContentArticlePublish:                 {Proto: notifyproto.NotificationType_NOTIFICATION_TYPE_ARTICLE_PUBLISHED},
	NotificationTypeContentArticleLike:                    {Proto: notifyproto.NotificationType_NOTIFICATION_TYPE_ARTICLE_LIKED},
	NotificationTypeContentArticleThank:                   {Proto: notifyproto.NotificationType_NOTIFICATION_TYPE_ARTICLE_THANKED},
	NotificationTypeContentArticleCollect:                 {Proto: notifyproto.NotificationType_NOTIFICATION_TYPE_ARTICLE_COLLECTED},
	NotificationTypeContentArticleWatch:                   {Proto: notifyproto.NotificationType_NOTIFICATION_TYPE_ARTICLE_WATCHED},
	NotificationTypeContentArticleAcceptAnswer:            {Proto: notifyproto.NotificationType_NOTIFICATION_TYPE_ARTICLE_ACCEPTED_ANSWER},
	NotificationTypeContentArticleLotteryWon:              {Proto: notifyproto.NotificationType_NOTIFICATION_TYPE_ARTICLE_LOTTERY_WON},
	NotificationTypeContentArticleWatchedCommentPublished: {Proto: notifyproto.NotificationType_NOTIFICATION_TYPE_ARTICLE_WATCHED_COMMENT_PUBLISHED},
	NotificationTypeContentArticleWatchedPostscriptAdded:  {Proto: notifyproto.NotificationType_NOTIFICATION_TYPE_ARTICLE_WATCHED_POSTSCRIPT_ADDED},
	NotificationTypeContentArticleWatchedAnswerAccepted:   {Proto: notifyproto.NotificationType_NOTIFICATION_TYPE_ARTICLE_WATCHED_ANSWER_ACCEPTED},
	NotificationTypeContentArticleWatchedPollEnded:        {Proto: notifyproto.NotificationType_NOTIFICATION_TYPE_ARTICLE_WATCHED_POLL_ENDED},
	NotificationTypeContentArticleWatchedLotteryDrawn:     {Proto: notifyproto.NotificationType_NOTIFICATION_TYPE_ARTICLE_WATCHED_LOTTERY_DRAWN},
	NotificationTypeContentArticleMention:                 {Proto: notifyproto.NotificationType_NOTIFICATION_TYPE_ARTICLE_MENTIONED},
	NotificationTypeContentPostscriptMention:              {Proto: notifyproto.NotificationType_NOTIFICATION_TYPE_POSTSCRIPT_MENTIONED},
	NotificationTypeContentArticleReward:                  {Proto: notifyproto.NotificationType_NOTIFICATION_TYPE_ARTICLE_REWARDED},
	NotificationTypeContentCommentPublish:                 {Proto: notifyproto.NotificationType_NOTIFICATION_TYPE_COMMENT_PUBLISHED},
	NotificationTypeContentCommentLike:                    {Proto: notifyproto.NotificationType_NOTIFICATION_TYPE_COMMENT_LIKED},
	NotificationTypeContentCommentThank:                   {Proto: notifyproto.NotificationType_NOTIFICATION_TYPE_COMMENT_THANKED},
	NotificationTypeContentCommentMention:                 {Proto: notifyproto.NotificationType_NOTIFICATION_TYPE_COMMENT_MENTIONED},
	NotificationTypeIMGlobalChatMention:                   {Proto: notifyproto.NotificationType_NOTIFICATION_TYPE_IM_GLOBAL_CHAT_MENTIONED},
})
