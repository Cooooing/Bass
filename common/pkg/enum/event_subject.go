package enum

import "common/proto/gen/common/enums"

type EventSubject string

const (
	EventSubjectUserRegister                EventSubject = "user.register"
	EventSubjectUserLogin                   EventSubject = "user.login"
	EventSubjectUserLogout                  EventSubject = "user.logout"
	EventSubjectUserEmailVerificationCode   EventSubject = "user.email.verification_code"
	EventSubjectUserPhoneVerificationCode   EventSubject = "user.phone.verification_code"
	EventSubjectUserFollow                  EventSubject = "user.follow"
	EventSubjectUserUnfollow                EventSubject = "user.unfollow"
	EventSubjectUserBlock                   EventSubject = "user.block"
	EventSubjectUserUnblock                 EventSubject = "user.unblock"
	EventSubjectUserTotpEnable              EventSubject = "user.totp.enable"
	EventSubjectUserTotpDisable             EventSubject = "user.totp.disable"
	EventSubjectUserAccountCancelled        EventSubject = "user.account.cancelled"
	EventSubjectUserAccountBanned           EventSubject = "user.account.banned"
	EventSubjectUserAccountUnbanned         EventSubject = "user.account.unbanned"
	EventSubjectContentArticlePublish       EventSubject = "content.article.published"
	EventSubjectContentArticleLike          EventSubject = "content.article.liked"
	EventSubjectContentArticleThank         EventSubject = "content.article.thanked"
	EventSubjectContentArticleCollect       EventSubject = "content.article.collected"
	EventSubjectContentArticleWatch         EventSubject = "content.article.watched"
	EventSubjectContentArticleAcceptAnswer  EventSubject = "content.article.accepted_answer"
	EventSubjectContentArticlePostscriptAdd EventSubject = "content.article.postscript_added"
	EventSubjectContentArticleStatusUpdate  EventSubject = "content.article.status_updated"
	EventSubjectContentCommentPublish       EventSubject = "content.comment.published"
	EventSubjectContentCommentLike          EventSubject = "content.comment.liked"
	EventSubjectContentCommentThank         EventSubject = "content.comment.thanked"
	EventSubjectContentCommentStatusUpdate  EventSubject = "content.comment.status_updated"
	EventSubjectGameIdleChatWorldMessage    EventSubject = "game_idle.chat.world.message"
)

var EventSubjectMap = NewMapping[EventSubject, enums.EventSubject](map[EventSubject]Entry[EventSubject, enums.EventSubject]{
	EventSubjectUserRegister:                {Proto: enums.EventSubject_EVENT_SUBJECT_USER_REGISTER},
	EventSubjectUserLogin:                   {Proto: enums.EventSubject_EVENT_SUBJECT_USER_LOGIN},
	EventSubjectUserLogout:                  {Proto: enums.EventSubject_EVENT_SUBJECT_USER_LOGOUT},
	EventSubjectUserEmailVerificationCode:   {Proto: enums.EventSubject_EVENT_SUBJECT_USER_EMAIL_VERIFICATION_CODE},
	EventSubjectUserPhoneVerificationCode:   {Proto: enums.EventSubject_EVENT_SUBJECT_USER_PHONE_VERIFICATION_CODE},
	EventSubjectUserFollow:                  {Proto: enums.EventSubject_EVENT_SUBJECT_USER_FOLLOW},
	EventSubjectUserUnfollow:                {Proto: enums.EventSubject_EVENT_SUBJECT_USER_UNFOLLOW},
	EventSubjectUserBlock:                   {Proto: enums.EventSubject_EVENT_SUBJECT_USER_BLOCK},
	EventSubjectUserUnblock:                 {Proto: enums.EventSubject_EVENT_SUBJECT_USER_UNBLOCK},
	EventSubjectUserTotpEnable:              {Proto: enums.EventSubject_EVENT_SUBJECT_USER_TOTP_ENABLE},
	EventSubjectUserTotpDisable:             {Proto: enums.EventSubject_EVENT_SUBJECT_USER_TOTP_DISABLE},
	EventSubjectUserAccountCancelled:        {Proto: enums.EventSubject_EVENT_SUBJECT_USER_ACCOUNT_CANCELLED},
	EventSubjectUserAccountBanned:           {Proto: enums.EventSubject_EVENT_SUBJECT_USER_ACCOUNT_BANNED},
	EventSubjectUserAccountUnbanned:         {Proto: enums.EventSubject_EVENT_SUBJECT_USER_ACCOUNT_UNBANNED},
	EventSubjectContentArticlePublish:       {Proto: enums.EventSubject_EVENT_SUBJECT_ARTICLE_PUBLISHED},
	EventSubjectContentArticleLike:          {Proto: enums.EventSubject_EVENT_SUBJECT_ARTICLE_LIKED},
	EventSubjectContentArticleThank:         {Proto: enums.EventSubject_EVENT_SUBJECT_ARTICLE_THANKED},
	EventSubjectContentArticleCollect:       {Proto: enums.EventSubject_EVENT_SUBJECT_ARTICLE_COLLECTED},
	EventSubjectContentArticleWatch:         {Proto: enums.EventSubject_EVENT_SUBJECT_ARTICLE_WATCHED},
	EventSubjectContentArticleAcceptAnswer:  {Proto: enums.EventSubject_EVENT_SUBJECT_ARTICLE_ACCEPTED_ANSWER},
	EventSubjectContentArticlePostscriptAdd: {Proto: enums.EventSubject_EVENT_SUBJECT_ARTICLE_POSTSCRIPT_ADDED},
	EventSubjectContentArticleStatusUpdate:  {Proto: enums.EventSubject_EVENT_SUBJECT_ARTICLE_STATUS_UPDATED},
	EventSubjectContentCommentPublish:       {Proto: enums.EventSubject_EVENT_SUBJECT_COMMENT_PUBLISHED},
	EventSubjectContentCommentLike:          {Proto: enums.EventSubject_EVENT_SUBJECT_COMMENT_LIKED},
	EventSubjectContentCommentThank:         {Proto: enums.EventSubject_EVENT_SUBJECT_COMMENT_THANKED},
	EventSubjectContentCommentStatusUpdate:  {Proto: enums.EventSubject_EVENT_SUBJECT_COMMENT_STATUS_UPDATED},
	EventSubjectGameIdleChatWorldMessage:    {Proto: enums.EventSubject_EVENT_SUBJECT_GAME_IDLE_CHAT_WORLD_MESSAGE},
})

func (EventSubject) Values() []string {
	return nil
}

func (e EventSubject) String() string {
	return string(e)
}
