package biz

import (
	commonenum "common/pkg/enum"

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
	ProvideEventSubjects,

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
) map[commonenum.EventType][]usecase.EventHandler {
	eventHandlers := map[commonenum.EventType][]usecase.EventHandler{
		commonenum.EventTypeContentArticleCollect:       {articleCollectedHandler},
		commonenum.EventTypeContentArticleLike:          {articleLikedHandler},
		commonenum.EventTypeContentArticlePublish:       {articlePublishedHandler, articleMentionedHandler},
		commonenum.EventTypeContentArticleThank:         {articleThankedHandler},
		commonenum.EventTypeContentArticleReward:        {articleRewardedHandler},
		commonenum.EventTypeContentArticleWatch:         {articleWatchCreatedHandler},
		commonenum.EventTypeContentArticleAcceptAnswer:  {articleAcceptedAnswerHandler, articleWatchAnswerAcceptedHandler},
		commonenum.EventTypeContentArticleLotteryWon:    {articleLotteryWonHandler},
		commonenum.EventTypeContentArticleLotteryDrawn:  {articleWatchLotteryDrawnHandler},
		commonenum.EventTypeContentArticlePollEnded:     {articleWatchPollEndedHandler},
		commonenum.EventTypeContentArticlePostscriptAdd: {articleWatchPostscriptAddedHandler, postscriptMentionedHandler},
		commonenum.EventTypeContentCommentLike:          {commentLikedHandler},
		commonenum.EventTypeContentCommentThank:         {commentThankedHandler},
		commonenum.EventTypeContentCommentPublish:       {commentPublishedHandler, articleWatchCommentPublishedHandler, commentMentionedHandler},
		commonenum.EventTypeIMGlobalChatMessageCreated:  {globalChatMentionedHandler},
		commonenum.EventTypeUserFollow:                  {userFollowHandler},
		commonenum.EventTypeUserRegister:                {userRegisterHandler},
		commonenum.EventTypeUserEmailVerificationCode:   {userVerificationCodeHandler},
		commonenum.EventTypeUserPhoneVerificationCode:   {userVerificationCodeHandler},
	}

	return eventHandlers
}

func ProvideEventSubjects() []commonenum.EventSubject {
	return []commonenum.EventSubject{
		commonenum.EventSubjectContentArticleCollect,
		commonenum.EventSubjectContentArticleLike,
		commonenum.EventSubjectContentArticlePublish,
		commonenum.EventSubjectContentArticleThank,
		commonenum.EventSubjectContentArticleReward,
		commonenum.EventSubjectContentArticleWatch,
		commonenum.EventSubjectContentArticleAcceptAnswer,
		commonenum.EventSubjectContentArticlePostscriptAdd,
		commonenum.EventSubjectContentArticleLotteryWon,
		commonenum.EventSubjectContentArticleLotteryDrawn,
		commonenum.EventSubjectContentArticlePollEnded,
		commonenum.EventSubjectContentCommentLike,
		commonenum.EventSubjectContentCommentThank,
		commonenum.EventSubjectContentCommentPublish,
		commonenum.EventSubjectIMGlobalChatMessageCreated,
		commonenum.EventSubjectUserFollow,
		commonenum.EventSubjectUserRegister,
		commonenum.EventSubjectUserEmailVerificationCode,
		commonenum.EventSubjectUserPhoneVerificationCode,
	}
}
