package enum

import "common/proto/gen/common/enums"

type EventType string

const (
	EventTypeUserRegister                          EventType = "user_register"
	EventTypeUserLogin                             EventType = "user_login"
	EventTypeUserLogout                            EventType = "user_logout"
	EventTypeUserEmailVerificationCode             EventType = "user_email_verification_code"
	EventTypeUserPhoneVerificationCode             EventType = "user_phone_verification_code"
	EventTypeUserFollow                            EventType = "user_follow"
	EventTypeUserUnfollow                          EventType = "user_unfollow"
	EventTypeUserBlock                             EventType = "user_block"
	EventTypeUserUnblock                           EventType = "user_unblock"
	EventTypeUserTotpEnable                        EventType = "user_totp_enable"
	EventTypeUserTotpDisable                       EventType = "user_totp_disable"
	EventTypeUserAccountCancelled                  EventType = "user_account_cancelled"
	EventTypeUserAccountBanned                     EventType = "user_account_banned"
	EventTypeUserAccountUnbanned                   EventType = "user_account_unbanned"
	EventTypeContentArticlePublish                 EventType = "content_article_publish"
	EventTypeContentArticleLike                    EventType = "content_article_like"
	EventTypeContentArticleThank                   EventType = "content_article_thank"
	EventTypeContentArticleCollect                 EventType = "content_article_collect"
	EventTypeContentArticleWatch                   EventType = "content_article_watch"
	EventTypeContentArticleAcceptAnswer            EventType = "content_article_accept_answer"
	EventTypeContentArticlePostscriptAdd           EventType = "content_article_postscript_add"
	EventTypeContentArticleStatusUpdate            EventType = "content_article_status_update"
	EventTypeContentArticleLotteryWon              EventType = "content_article_lottery_won"
	EventTypeContentArticlePollEnded               EventType = "content_article_poll_ended"
	EventTypeContentArticleLotteryDrawn            EventType = "content_article_lottery_drawn"
	EventTypeContentArticleMention                 EventType = "content_article_mention"
	EventTypeContentPostscriptMention              EventType = "content_postscript_mention"
	EventTypeContentArticleWatchedCommentPublished EventType = "content_article_watched_comment_published"
	EventTypeContentArticleWatchedPostscriptAdded  EventType = "content_article_watched_postscript_added"
	EventTypeContentArticleWatchedAnswerAccepted   EventType = "content_article_watched_answer_accepted"
	EventTypeContentArticleWatchedPollEnded        EventType = "content_article_watched_poll_ended"
	EventTypeContentArticleWatchedLotteryDrawn     EventType = "content_article_watched_lottery_drawn"
	EventTypeContentCommentPublish                 EventType = "content_comment_publish"
	EventTypeContentCommentLike                    EventType = "content_comment_like"
	EventTypeContentCommentThank                   EventType = "content_comment_thank"
	EventTypeContentCommentStatusUpdate            EventType = "content_comment_status_update"
	EventTypeContentCommentMention                 EventType = "content_comment_mention"
	EventTypeIMGlobalChatMessageCreated            EventType = "im_global_chat_message_created"
	EventTypeIMGlobalChatBarrageCreated            EventType = "im_global_chat_barrage_created"
	EventTypeIMGlobalChatMention                   EventType = "im_global_chat_mention"
	EventTypeGameIdleChatMessage                   EventType = "game_idle_chat_message"
	EventTypeGameIdleCloseSession                  EventType = "game_idle_close_session"
	EventTypeGameIdleActionCompleted               EventType = "game_idle_action_completed"
	EventTypeGameIdleAbilityLeveledUp              EventType = "game_idle_ability_leveled_up"
	EventTypeGameIdleActionQueueUpdated            EventType = "game_idle_action_queue_updated"
	EventTypeNotifyStationMessagesCreated          EventType = "notify_station_messages_created"
)

var EventTypeMap = NewMapping[EventType, enums.EventType](map[EventType]Entry[EventType, enums.EventType]{
	EventTypeUserRegister:                          {Proto: enums.EventType_EVENT_TYPE_USER_REGISTER},
	EventTypeUserLogin:                             {Proto: enums.EventType_EVENT_TYPE_USER_LOGIN},
	EventTypeUserLogout:                            {Proto: enums.EventType_EVENT_TYPE_USER_LOGOUT},
	EventTypeUserEmailVerificationCode:             {Proto: enums.EventType_EVENT_TYPE_USER_EMAIL_VERIFICATION_CODE},
	EventTypeUserPhoneVerificationCode:             {Proto: enums.EventType_EVENT_TYPE_USER_PHONE_VERIFICATION_CODE},
	EventTypeUserFollow:                            {Proto: enums.EventType_EVENT_TYPE_USER_FOLLOW},
	EventTypeUserUnfollow:                          {Proto: enums.EventType_EVENT_TYPE_USER_UNFOLLOW},
	EventTypeUserBlock:                             {Proto: enums.EventType_EVENT_TYPE_USER_BLOCK},
	EventTypeUserUnblock:                           {Proto: enums.EventType_EVENT_TYPE_USER_UNBLOCK},
	EventTypeUserTotpEnable:                        {Proto: enums.EventType_EVENT_TYPE_USER_TOTP_ENABLE},
	EventTypeUserTotpDisable:                       {Proto: enums.EventType_EVENT_TYPE_USER_TOTP_DISABLE},
	EventTypeUserAccountCancelled:                  {Proto: enums.EventType_EVENT_TYPE_USER_ACCOUNT_CANCELLED},
	EventTypeUserAccountBanned:                     {Proto: enums.EventType_EVENT_TYPE_USER_ACCOUNT_BANNED},
	EventTypeUserAccountUnbanned:                   {Proto: enums.EventType_EVENT_TYPE_USER_ACCOUNT_UNBANNED},
	EventTypeContentArticlePublish:                 {Proto: enums.EventType_EVENT_TYPE_ARTICLE_PUBLISHED},
	EventTypeContentArticleLike:                    {Proto: enums.EventType_EVENT_TYPE_ARTICLE_LIKED},
	EventTypeContentArticleThank:                   {Proto: enums.EventType_EVENT_TYPE_ARTICLE_THANKED},
	EventTypeContentArticleCollect:                 {Proto: enums.EventType_EVENT_TYPE_ARTICLE_COLLECTED},
	EventTypeContentArticleWatch:                   {Proto: enums.EventType_EVENT_TYPE_ARTICLE_WATCHED},
	EventTypeContentArticleAcceptAnswer:            {Proto: enums.EventType_EVENT_TYPE_ARTICLE_ACCEPTED_ANSWER},
	EventTypeContentArticlePostscriptAdd:           {Proto: enums.EventType_EVENT_TYPE_ARTICLE_POSTSCRIPT_ADDED},
	EventTypeContentArticleStatusUpdate:            {Proto: enums.EventType_EVENT_TYPE_ARTICLE_STATUS_UPDATED},
	EventTypeContentArticleLotteryWon:              {Proto: enums.EventType_EVENT_TYPE_ARTICLE_LOTTERY_WON},
	EventTypeContentArticlePollEnded:               {Proto: enums.EventType_EVENT_TYPE_ARTICLE_POLL_ENDED},
	EventTypeContentArticleLotteryDrawn:            {Proto: enums.EventType_EVENT_TYPE_ARTICLE_LOTTERY_DRAWN},
	EventTypeContentArticleMention:                 {Proto: enums.EventType_EVENT_TYPE_ARTICLE_MENTIONED},
	EventTypeContentPostscriptMention:              {Proto: enums.EventType_EVENT_TYPE_POSTSCRIPT_MENTIONED},
	EventTypeContentArticleWatchedCommentPublished: {Proto: enums.EventType_EVENT_TYPE_ARTICLE_WATCHED_COMMENT_PUBLISHED},
	EventTypeContentArticleWatchedPostscriptAdded:  {Proto: enums.EventType_EVENT_TYPE_ARTICLE_WATCHED_POSTSCRIPT_ADDED},
	EventTypeContentArticleWatchedAnswerAccepted:   {Proto: enums.EventType_EVENT_TYPE_ARTICLE_WATCHED_ANSWER_ACCEPTED},
	EventTypeContentArticleWatchedPollEnded:        {Proto: enums.EventType_EVENT_TYPE_ARTICLE_WATCHED_POLL_ENDED},
	EventTypeContentArticleWatchedLotteryDrawn:     {Proto: enums.EventType_EVENT_TYPE_ARTICLE_WATCHED_LOTTERY_DRAWN},
	EventTypeContentCommentPublish:                 {Proto: enums.EventType_EVENT_TYPE_COMMENT_PUBLISHED},
	EventTypeContentCommentLike:                    {Proto: enums.EventType_EVENT_TYPE_COMMENT_LIKED},
	EventTypeContentCommentThank:                   {Proto: enums.EventType_EVENT_TYPE_COMMENT_THANKED},
	EventTypeContentCommentStatusUpdate:            {Proto: enums.EventType_EVENT_TYPE_COMMENT_STATUS_UPDATED},
	EventTypeContentCommentMention:                 {Proto: enums.EventType_EVENT_TYPE_COMMENT_MENTIONED},
	EventTypeIMGlobalChatMessageCreated:            {Proto: enums.EventType_EVENT_TYPE_IM_GLOBAL_CHAT_MESSAGE_CREATED},
	EventTypeIMGlobalChatBarrageCreated:            {Proto: enums.EventType_EVENT_TYPE_IM_GLOBAL_CHAT_BARRAGE_CREATED},
	EventTypeIMGlobalChatMention:                   {Proto: enums.EventType_EVENT_TYPE_IM_GLOBAL_CHAT_MENTIONED},
	EventTypeGameIdleChatMessage:                   {Proto: enums.EventType_EVENT_TYPE_GAME_IDLE_CHAT_MESSAGE},
	EventTypeGameIdleCloseSession:                  {Proto: enums.EventType_EVENT_TYPE_GAME_IDLE_CLOSE_SESSION},
	EventTypeGameIdleActionCompleted:               {Proto: enums.EventType_EVENT_TYPE_GAME_IDLE_ACTION_COMPLETED},
	EventTypeGameIdleAbilityLeveledUp:              {Proto: enums.EventType_EVENT_TYPE_GAME_IDLE_ABILITY_LEVELED_UP},
	EventTypeGameIdleActionQueueUpdated:            {Proto: enums.EventType_EVENT_TYPE_GAME_IDLE_ACTION_QUEUE_UPDATED},
	EventTypeNotifyStationMessagesCreated:          {Proto: enums.EventType_EVENT_TYPE_NOTIFY_STATION_MESSAGES_CREATED},
})

func (e EventType) String() string {
	return string(e)
}
