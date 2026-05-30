package enum

import "common/api/gen/common/enums"

// EventSubject 定义跨服务事件的 MQ 主题。
type EventSubject string

const (
	EventSubjectUserRegister              EventSubject = "user.register"
	EventSubjectUserLogin                 EventSubject = "user.login"
	EventSubjectUserLogout                EventSubject = "user.logout"
	EventSubjectUserEmailVerificationCode EventSubject = "user.email.verification_code"
	EventSubjectUserPhoneVerificationCode EventSubject = "user.phone.verification_code"
	EventSubjectUserFollow                EventSubject = "user.follow"
	EventSubjectUserUnfollow              EventSubject = "user.unfollow"
	EventSubjectUserBlock                 EventSubject = "user.block"
	EventSubjectUserUnblock               EventSubject = "user.unblock"
	EventSubjectUserTotpEnable            EventSubject = "user.totp.enable"
	EventSubjectUserTotpDisable           EventSubject = "user.totp.disable"
	EventSubjectContentArticlePublish     EventSubject = "content.article.published"
	EventSubjectContentArticleLike        EventSubject = "content.article.liked"
	EventSubjectContentArticleThank       EventSubject = "content.article.thanked"
	EventSubjectContentArticleCollect     EventSubject = "content.article.collected"
	EventSubjectContentArticleWatch       EventSubject = "content.article.watched"
	EventSubjectContentCommentPublish     EventSubject = "content.comment.published"
	EventSubjectContentCommentLike        EventSubject = "content.comment.liked"
)

// EventSubjectMap 将内部事件主题映射到 proto 事件主题。
var EventSubjectMap = NewMapping[EventSubject, enums.EventSubject](map[EventSubject]Entry[EventSubject, enums.EventSubject]{
	EventSubjectUserRegister:              {Proto: enums.EventSubject_EVENT_SUBJECT_USER_REGISTER},
	EventSubjectUserLogin:                 {Proto: enums.EventSubject_EVENT_SUBJECT_USER_LOGIN},
	EventSubjectUserLogout:                {Proto: enums.EventSubject_EVENT_SUBJECT_USER_LOGOUT},
	EventSubjectUserEmailVerificationCode: {Proto: enums.EventSubject_EVENT_SUBJECT_USER_EMAIL_VERIFICATION_CODE},
	EventSubjectUserPhoneVerificationCode: {Proto: enums.EventSubject_EVENT_SUBJECT_USER_PHONE_VERIFICATION_CODE},
	EventSubjectUserFollow:                {Proto: enums.EventSubject_EVENT_SUBJECT_USER_FOLLOW},
	EventSubjectUserUnfollow:              {Proto: enums.EventSubject_EVENT_SUBJECT_USER_UNFOLLOW},
	EventSubjectUserBlock:                 {Proto: enums.EventSubject_EVENT_SUBJECT_USER_BLOCK},
	EventSubjectUserUnblock:               {Proto: enums.EventSubject_EVENT_SUBJECT_USER_UNBLOCK},
	EventSubjectUserTotpEnable:            {Proto: enums.EventSubject_EVENT_SUBJECT_USER_TOTP_ENABLE},
	EventSubjectUserTotpDisable:           {Proto: enums.EventSubject_EVENT_SUBJECT_USER_TOTP_DISABLE},
	EventSubjectContentArticlePublish:     {Proto: enums.EventSubject_EVENT_SUBJECT_ARTICLE_PUBLISHED},
	EventSubjectContentArticleLike:        {Proto: enums.EventSubject_EVENT_SUBJECT_ARTICLE_LIKED},
	EventSubjectContentArticleThank:       {Proto: enums.EventSubject_EVENT_SUBJECT_ARTICLE_THANKED},
	EventSubjectContentArticleCollect:     {Proto: enums.EventSubject_EVENT_SUBJECT_ARTICLE_COLLECTED},
	EventSubjectContentArticleWatch:       {Proto: enums.EventSubject_EVENT_SUBJECT_ARTICLE_WATCHED},
	EventSubjectContentCommentPublish:     {Proto: enums.EventSubject_EVENT_SUBJECT_COMMENT_PUBLISHED},
	EventSubjectContentCommentLike:        {Proto: enums.EventSubject_EVENT_SUBJECT_COMMENT_LIKED},
})

// Values 返回 Ent enum 允许写入的持久化值。
func (EventSubject) Values() []string {
	return EventSubjectMap.EnumValues()
}

// EventSubjectByEventType 根据事件类型返回对应的 MQ 主题。
func EventSubjectByEventType(eventType enums.EventType) (EventSubject, bool) {
	return EventSubjectMap.ToEnum(enums.EventSubject(eventType))
}
