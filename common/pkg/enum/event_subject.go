package enum

import (
	"fmt"
	"strings"

	event "common/proto/gen/common/enums/event"
)

type EventSubject string

const (
	EventSubjectUserRegister                 EventSubject = "user.register"
	EventSubjectUserLogin                    EventSubject = "user.login"
	EventSubjectUserLogout                   EventSubject = "user.logout"
	EventSubjectUserEmailVerificationCode    EventSubject = "user.email.verification_code"
	EventSubjectUserPhoneVerificationCode    EventSubject = "user.phone.verification_code"
	EventSubjectUserFollow                   EventSubject = "user.follow"
	EventSubjectUserUnfollow                 EventSubject = "user.unfollow"
	EventSubjectUserBlock                    EventSubject = "user.block"
	EventSubjectUserUnblock                  EventSubject = "user.unblock"
	EventSubjectUserTotpEnable               EventSubject = "user.totp.enable"
	EventSubjectUserTotpDisable              EventSubject = "user.totp.disable"
	EventSubjectUserAccountCancelled         EventSubject = "user.account.cancelled"
	EventSubjectUserAccountBanned            EventSubject = "user.account.banned"
	EventSubjectUserAccountUnbanned          EventSubject = "user.account.unbanned"
	EventSubjectUserCheckinCompleted         EventSubject = "user.checkin.completed"
	EventSubjectContentArticlePublish        EventSubject = "content.article.published"
	EventSubjectContentArticleLike           EventSubject = "content.article.liked"
	EventSubjectContentArticleThank          EventSubject = "content.article.thanked"
	EventSubjectContentArticleReward         EventSubject = "content.article.rewarded"
	EventSubjectContentArticleCollect        EventSubject = "content.article.collected"
	EventSubjectContentArticleWatch          EventSubject = "content.article.watched"
	EventSubjectContentArticleAcceptAnswer   EventSubject = "content.article.accepted_answer"
	EventSubjectContentArticlePostscriptAdd  EventSubject = "content.article.postscript_added"
	EventSubjectContentArticleStatusUpdate   EventSubject = "content.article.status_updated"
	EventSubjectContentArticleLotteryWon     EventSubject = "content.article.lottery_won"
	EventSubjectContentArticlePollEnded      EventSubject = "content.article.poll_ended"
	EventSubjectContentArticleLotteryDrawn   EventSubject = "content.article.lottery_drawn"
	EventSubjectContentCommentPublish        EventSubject = "content.comment.published"
	EventSubjectContentCommentLike           EventSubject = "content.comment.liked"
	EventSubjectContentCommentThank          EventSubject = "content.comment.thanked"
	EventSubjectContentCommentStatusUpdate   EventSubject = "content.comment.status_updated"
	EventSubjectIMPrivateChatMessageCreated  EventSubject = "im.private_chat.message_created"
	EventSubjectIMGlobalChatMessageCreated   EventSubject = "im.global_chat.message_created"
	EventSubjectIMGlobalChatBarrageCreated   EventSubject = "im.global_chat.barrage_created"
	EventSubjectIMGlobalChatTopicUpdated     EventSubject = "im.global_chat.topic_updated"
	EventSubjectGameIdle                     EventSubject = "game_idle"
	EventSubjectGameIdleChatMessage          EventSubject = "game_idle.chat_message"
	EventSubjectGameIdleCloseSession         EventSubject = "game_idle.close_session"
	EventSubjectGameIdleActionCompleted      EventSubject = "game_idle.action_completed"
	EventSubjectGameIdleAbilityLeveledUp     EventSubject = "game_idle.ability_leveled_up"
	EventSubjectGameIdleActionQueueUpdated   EventSubject = "game_idle.action_queue_updated"
	EventSubjectNotifyStationMessagesCreated EventSubject = "notify.station_messages.created"
)

var EventSubjectMap = NewMapping[EventSubject, event.EventSubject](map[EventSubject]Entry[EventSubject, event.EventSubject]{
	EventSubjectUserRegister:                 {Proto: event.EventSubject_EVENT_SUBJECT_USER_REGISTER},
	EventSubjectUserLogin:                    {Proto: event.EventSubject_EVENT_SUBJECT_USER_LOGIN},
	EventSubjectUserLogout:                   {Proto: event.EventSubject_EVENT_SUBJECT_USER_LOGOUT},
	EventSubjectUserEmailVerificationCode:    {Proto: event.EventSubject_EVENT_SUBJECT_USER_EMAIL_VERIFICATION_CODE},
	EventSubjectUserPhoneVerificationCode:    {Proto: event.EventSubject_EVENT_SUBJECT_USER_PHONE_VERIFICATION_CODE},
	EventSubjectUserFollow:                   {Proto: event.EventSubject_EVENT_SUBJECT_USER_FOLLOW},
	EventSubjectUserUnfollow:                 {Proto: event.EventSubject_EVENT_SUBJECT_USER_UNFOLLOW},
	EventSubjectUserBlock:                    {Proto: event.EventSubject_EVENT_SUBJECT_USER_BLOCK},
	EventSubjectUserUnblock:                  {Proto: event.EventSubject_EVENT_SUBJECT_USER_UNBLOCK},
	EventSubjectUserTotpEnable:               {Proto: event.EventSubject_EVENT_SUBJECT_USER_TOTP_ENABLE},
	EventSubjectUserTotpDisable:              {Proto: event.EventSubject_EVENT_SUBJECT_USER_TOTP_DISABLE},
	EventSubjectUserAccountCancelled:         {Proto: event.EventSubject_EVENT_SUBJECT_USER_ACCOUNT_CANCELLED},
	EventSubjectUserAccountBanned:            {Proto: event.EventSubject_EVENT_SUBJECT_USER_ACCOUNT_BANNED},
	EventSubjectUserAccountUnbanned:          {Proto: event.EventSubject_EVENT_SUBJECT_USER_ACCOUNT_UNBANNED},
	EventSubjectUserCheckinCompleted:         {Proto: event.EventSubject_EVENT_SUBJECT_USER_CHECKIN_COMPLETED},
	EventSubjectContentArticlePublish:        {Proto: event.EventSubject_EVENT_SUBJECT_ARTICLE_PUBLISHED},
	EventSubjectContentArticleLike:           {Proto: event.EventSubject_EVENT_SUBJECT_ARTICLE_LIKED},
	EventSubjectContentArticleThank:          {Proto: event.EventSubject_EVENT_SUBJECT_ARTICLE_THANKED},
	EventSubjectContentArticleReward:         {Proto: event.EventSubject_EVENT_SUBJECT_ARTICLE_REWARDED},
	EventSubjectContentArticleCollect:        {Proto: event.EventSubject_EVENT_SUBJECT_ARTICLE_COLLECTED},
	EventSubjectContentArticleWatch:          {Proto: event.EventSubject_EVENT_SUBJECT_ARTICLE_WATCHED},
	EventSubjectContentArticleAcceptAnswer:   {Proto: event.EventSubject_EVENT_SUBJECT_ARTICLE_ACCEPTED_ANSWER},
	EventSubjectContentArticlePostscriptAdd:  {Proto: event.EventSubject_EVENT_SUBJECT_ARTICLE_POSTSCRIPT_ADDED},
	EventSubjectContentArticleStatusUpdate:   {Proto: event.EventSubject_EVENT_SUBJECT_ARTICLE_STATUS_UPDATED},
	EventSubjectContentArticleLotteryWon:     {Proto: event.EventSubject_EVENT_SUBJECT_ARTICLE_LOTTERY_WON},
	EventSubjectContentArticlePollEnded:      {Proto: event.EventSubject_EVENT_SUBJECT_ARTICLE_POLL_ENDED},
	EventSubjectContentArticleLotteryDrawn:   {Proto: event.EventSubject_EVENT_SUBJECT_ARTICLE_LOTTERY_DRAWN},
	EventSubjectContentCommentPublish:        {Proto: event.EventSubject_EVENT_SUBJECT_COMMENT_PUBLISHED},
	EventSubjectContentCommentLike:           {Proto: event.EventSubject_EVENT_SUBJECT_COMMENT_LIKED},
	EventSubjectContentCommentThank:          {Proto: event.EventSubject_EVENT_SUBJECT_COMMENT_THANKED},
	EventSubjectContentCommentStatusUpdate:   {Proto: event.EventSubject_EVENT_SUBJECT_COMMENT_STATUS_UPDATED},
	EventSubjectIMPrivateChatMessageCreated:  {Proto: event.EventSubject_EVENT_SUBJECT_IM_PRIVATE_CHAT_MESSAGE_CREATED},
	EventSubjectIMGlobalChatMessageCreated:   {Proto: event.EventSubject_EVENT_SUBJECT_IM_GLOBAL_CHAT_MESSAGE_CREATED},
	EventSubjectIMGlobalChatBarrageCreated:   {Proto: event.EventSubject_EVENT_SUBJECT_IM_GLOBAL_CHAT_BARRAGE_CREATED},
	EventSubjectIMGlobalChatTopicUpdated:     {Proto: event.EventSubject_EVENT_SUBJECT_IM_GLOBAL_CHAT_TOPIC_UPDATED},
	EventSubjectGameIdle:                     {Proto: event.EventSubject_EVENT_SUBJECT_GAME_IDLE},
	EventSubjectGameIdleChatMessage:          {Proto: event.EventSubject_EVENT_SUBJECT_GAME_IDLE_CHAT_MESSAGE},
	EventSubjectGameIdleCloseSession:         {Proto: event.EventSubject_EVENT_SUBJECT_GAME_IDLE_CLOSE_SESSION},
	EventSubjectGameIdleActionCompleted:      {Proto: event.EventSubject_EVENT_SUBJECT_GAME_IDLE_ACTION_COMPLETED},
	EventSubjectGameIdleAbilityLeveledUp:     {Proto: event.EventSubject_EVENT_SUBJECT_GAME_IDLE_ABILITY_LEVELED_UP},
	EventSubjectGameIdleActionQueueUpdated:   {Proto: event.EventSubject_EVENT_SUBJECT_GAME_IDLE_ACTION_QUEUE_UPDATED},
	EventSubjectNotifyStationMessagesCreated: {Proto: event.EventSubject_EVENT_SUBJECT_NOTIFY_STATION_MESSAGES_CREATED},
})

func (e EventSubject) String() string {
	return string(e)
}

func EventSubjectCheck(column string, subjects []string) string {
	quoted := make([]string, 0, len(subjects))
	for _, value := range subjects {
		quoted = append(quoted, fmt.Sprintf("'%s'", value))
	}
	return fmt.Sprintf("%s IN (%s)", column, strings.Join(quoted, ", "))
}
