package biz

import (
	event "common/proto/gen/common/enums/event"

	"notify/internal/biz/usecase"
	channelusecase "notify/internal/biz/usecase/channel"
	"notify/internal/biz/usecase/handler"

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
) map[event.EventSubject][]usecase.EventHandler {
	eventHandlers := map[event.EventSubject][]usecase.EventHandler{
		event.EventSubject_EVENT_SUBJECT_ARTICLE_COLLECTED:              {articleCollectedHandler},
		event.EventSubject_EVENT_SUBJECT_ARTICLE_LIKED:                  {articleLikedHandler},
		event.EventSubject_EVENT_SUBJECT_ARTICLE_PUBLISHED:              {articlePublishedHandler, articleMentionedHandler},
		event.EventSubject_EVENT_SUBJECT_ARTICLE_THANKED:                {articleThankedHandler},
		event.EventSubject_EVENT_SUBJECT_ARTICLE_REWARDED:               {articleRewardedHandler},
		event.EventSubject_EVENT_SUBJECT_ARTICLE_WATCHED:                {articleWatchCreatedHandler},
		event.EventSubject_EVENT_SUBJECT_ARTICLE_ACCEPTED_ANSWER:        {articleAcceptedAnswerHandler, articleWatchAnswerAcceptedHandler},
		event.EventSubject_EVENT_SUBJECT_ARTICLE_LOTTERY_WON:            {articleLotteryWonHandler},
		event.EventSubject_EVENT_SUBJECT_ARTICLE_LOTTERY_DRAWN:          {articleWatchLotteryDrawnHandler},
		event.EventSubject_EVENT_SUBJECT_ARTICLE_POLL_ENDED:             {articleWatchPollEndedHandler},
		event.EventSubject_EVENT_SUBJECT_ARTICLE_POSTSCRIPT_ADDED:       {articleWatchPostscriptAddedHandler, postscriptMentionedHandler},
		event.EventSubject_EVENT_SUBJECT_COMMENT_LIKED:                  {commentLikedHandler},
		event.EventSubject_EVENT_SUBJECT_COMMENT_THANKED:                {commentThankedHandler},
		event.EventSubject_EVENT_SUBJECT_COMMENT_PUBLISHED:              {commentPublishedHandler, articleWatchCommentPublishedHandler, commentMentionedHandler},
		event.EventSubject_EVENT_SUBJECT_IM_GLOBAL_CHAT_MESSAGE_CREATED: {globalChatMentionedHandler},
		event.EventSubject_EVENT_SUBJECT_USER_FOLLOW:                    {userFollowHandler},
		event.EventSubject_EVENT_SUBJECT_USER_REGISTER:                  {userRegisterHandler},
		event.EventSubject_EVENT_SUBJECT_USER_EMAIL_VERIFICATION_CODE:   {userVerificationCodeHandler},
		event.EventSubject_EVENT_SUBJECT_USER_PHONE_VERIFICATION_CODE:   {userVerificationCodeHandler},
	}

	return eventHandlers
}
