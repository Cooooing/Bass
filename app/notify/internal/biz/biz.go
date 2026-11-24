package biz

import (
	event "common/proto/gen/common/enums/event/v1"

	"notify/internal/biz/usecase"
	channelusecase "notify/internal/biz/usecase/channel"
	"notify/internal/biz/usecase/handler"
	notifyenum "notify/internal/enum"

	"github.com/google/wire"
)

// BizProviderSet is the dependency provider set for biz.
var BizProviderSet = wire.NewSet(

	handler.NewArticleCollectedHandler,
	handler.NewArticleLikedHandler,
	handler.NewArticlePublishedHandler,
	handler.NewArticleMentionedHandler,
	handler.NewArticleThankedHandler,
	handler.NewArticleRewardedHandler,
	handler.NewArticleAcceptedAnswerHandler,
	handler.NewArticleLotteryWonHandler,
	handler.NewArticleWatchCreatedHandler,
	handler.NewArticleWatchCommentPublishedHandler,
	handler.NewArticleWatchPostscriptAddedHandler,
	handler.NewArticleWatchAnswerAcceptedHandler,
	handler.NewArticleWatchPollEndedHandler,
	handler.NewArticleWatchLotteryDrawnHandler,
	handler.NewCommentLikedHandler,
	handler.NewCommentThankedHandler,
	handler.NewCommentPublishedHandler,
	handler.NewCommentMentionedHandler,
	handler.NewPostscriptMentionedHandler,
	handler.NewGlobalChatMentionedHandler,
	handler.NewUserRegisterHandler,
	handler.NewUserFollowHandler,
	handler.NewUserVerificationCodeHandler,
	ProvideEventHandlers,

	usecase.NewConsumerUsecase,
	usecase.NewEventUsecase,
	usecase.NewNotifyUsecase,
	usecase.NewRateLimitUsecase,
	usecase.NewStationMessageUsecase,
	usecase.NewTemplateUsecase,
	channelusecase.NewStationUsecase,
	channelusecase.NewEmailUsecase,
	channelusecase.NewTencentSMSUsecase,
	channelusecase.NewLarkWebhookUsecase,
)

func ProvideEventHandlers(
	articleCollectedHandler *handler.ArticleCollectedHandler,
	articleLikedHandler *handler.ArticleLikedHandler,
	articlePublishedHandler *handler.ArticlePublishedHandler,
	articleMentionedHandler *handler.ArticleMentionedHandler,
	articleThankedHandler *handler.ArticleThankedHandler,
	articleRewardedHandler *handler.ArticleRewardedHandler,
	articleAcceptedAnswerHandler *handler.ArticleAcceptedAnswerHandler,
	articleLotteryWonHandler *handler.ArticleLotteryWonHandler,
	articleWatchCreatedHandler *handler.ArticleWatchCreatedHandler,
	articleWatchCommentPublishedHandler *handler.ArticleWatchCommentPublishedHandler,
	articleWatchPostscriptAddedHandler *handler.ArticleWatchPostscriptAddedHandler,
	articleWatchAnswerAcceptedHandler *handler.ArticleWatchAnswerAcceptedHandler,
	articleWatchPollEndedHandler *handler.ArticleWatchPollEndedHandler,
	articleWatchLotteryDrawnHandler *handler.ArticleWatchLotteryDrawnHandler,
	commentLikedHandler *handler.CommentLikedHandler,
	commentThankedHandler *handler.CommentThankedHandler,
	commentPublishedHandler *handler.CommentPublishedHandler,
	commentMentionedHandler *handler.CommentMentionedHandler,
	postscriptMentionedHandler *handler.PostscriptMentionedHandler,
	globalChatMentionedHandler *handler.GlobalChatMentionedHandler,
	userFollowHandler *handler.UserFollowHandler,
	userRegisterHandler *handler.UserRegisterHandler,
	userVerificationCodeHandler *handler.UserVerificationCodeHandler,
) map[event.EventSubject][]usecase.EventHandlerBinding {
	eventHandlers := map[event.EventSubject][]usecase.EventHandlerBinding{
		event.EventSubject_EVENT_SUBJECT_ARTICLE_COLLECTED: {{Handler: articleCollectedHandler, NotificationType: notifyenum.NotificationTypeContentArticleCollect}},
		event.EventSubject_EVENT_SUBJECT_ARTICLE_LIKED:     {{Handler: articleLikedHandler, NotificationType: notifyenum.NotificationTypeContentArticleLike}},
		event.EventSubject_EVENT_SUBJECT_ARTICLE_PUBLISHED: {
			{Handler: articlePublishedHandler, NotificationType: notifyenum.NotificationTypeContentArticlePublish},
			{Handler: articleMentionedHandler, NotificationType: notifyenum.NotificationTypeContentArticleMention},
		},
		event.EventSubject_EVENT_SUBJECT_ARTICLE_THANKED:  {{Handler: articleThankedHandler, NotificationType: notifyenum.NotificationTypeContentArticleThank}},
		event.EventSubject_EVENT_SUBJECT_ARTICLE_REWARDED: {{Handler: articleRewardedHandler, NotificationType: notifyenum.NotificationTypeContentArticleReward}},
		event.EventSubject_EVENT_SUBJECT_ARTICLE_WATCHED:  {{Handler: articleWatchCreatedHandler, NotificationType: notifyenum.NotificationTypeContentArticleWatch}},
		event.EventSubject_EVENT_SUBJECT_ARTICLE_ACCEPTED_ANSWER: {
			{Handler: articleAcceptedAnswerHandler, NotificationType: notifyenum.NotificationTypeContentArticleAcceptAnswer},
			{Handler: articleWatchAnswerAcceptedHandler, NotificationType: notifyenum.NotificationTypeContentArticleWatchedAnswerAccepted},
		},
		event.EventSubject_EVENT_SUBJECT_ARTICLE_LOTTERY_WON:   {{Handler: articleLotteryWonHandler, NotificationType: notifyenum.NotificationTypeContentArticleLotteryWon}},
		event.EventSubject_EVENT_SUBJECT_ARTICLE_LOTTERY_DRAWN: {{Handler: articleWatchLotteryDrawnHandler, NotificationType: notifyenum.NotificationTypeContentArticleWatchedLotteryDrawn}},
		event.EventSubject_EVENT_SUBJECT_ARTICLE_POLL_ENDED:    {{Handler: articleWatchPollEndedHandler, NotificationType: notifyenum.NotificationTypeContentArticleWatchedPollEnded}},
		event.EventSubject_EVENT_SUBJECT_ARTICLE_POSTSCRIPT_ADDED: {
			{Handler: articleWatchPostscriptAddedHandler, NotificationType: notifyenum.NotificationTypeContentArticleWatchedPostscriptAdded},
			{Handler: postscriptMentionedHandler, NotificationType: notifyenum.NotificationTypeContentPostscriptMention},
		},
		event.EventSubject_EVENT_SUBJECT_COMMENT_LIKED:   {{Handler: commentLikedHandler, NotificationType: notifyenum.NotificationTypeContentCommentLike}},
		event.EventSubject_EVENT_SUBJECT_COMMENT_THANKED: {{Handler: commentThankedHandler, NotificationType: notifyenum.NotificationTypeContentCommentThank}},
		event.EventSubject_EVENT_SUBJECT_COMMENT_PUBLISHED: {
			{Handler: commentPublishedHandler, NotificationType: notifyenum.NotificationTypeContentCommentPublish},
			{Handler: articleWatchCommentPublishedHandler, NotificationType: notifyenum.NotificationTypeContentArticleWatchedCommentPublished},
			{Handler: commentMentionedHandler, NotificationType: notifyenum.NotificationTypeContentCommentMention},
		},
		event.EventSubject_EVENT_SUBJECT_IM_GLOBAL_CHAT_MESSAGE_CREATED: {{Handler: globalChatMentionedHandler, NotificationType: notifyenum.NotificationTypeIMGlobalChatMention}},
		event.EventSubject_EVENT_SUBJECT_USER_FOLLOW:                    {{Handler: userFollowHandler, NotificationType: notifyenum.NotificationTypeUserFollow}},
		event.EventSubject_EVENT_SUBJECT_USER_REGISTER:                  {{Handler: userRegisterHandler, NotificationType: notifyenum.NotificationTypeUserRegister}},
		event.EventSubject_EVENT_SUBJECT_USER_EMAIL_VERIFICATION_CODE:   {{Handler: userVerificationCodeHandler, NotificationType: notifyenum.NotificationTypeUserEmailVerificationCode}},
		event.EventSubject_EVENT_SUBJECT_USER_PHONE_VERIFICATION_CODE:   {{Handler: userVerificationCodeHandler, NotificationType: notifyenum.NotificationTypeUserPhoneVerificationCode}},
	}

	return eventHandlers
}
