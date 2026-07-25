package biz

import (
	commonenum "common/pkg/enum"
	"notify/internal/biz/usecase"
	"notify/internal/biz/usecase/consumer"
	"notify/internal/biz/usecase/handler"

	"github.com/google/wire"
)

// BizProviderSet 是 biz 层依赖集合。
var BizProviderSet = wire.NewSet(

	handler.NewArticleCollectedHandler,
	handler.NewArticleLikedHandler,
	handler.NewArticlePublishedHandler,
	handler.NewArticleThankedHandler,
	handler.NewArticleWatchedHandler,
	handler.NewCommentLikedHandler,
	handler.NewCommentPublishedHandler,
	handler.NewUserRegisterHandler,
	handler.NewUserFollowHandler,
	handler.NewUserVerificationCodeHandler,
	ProvideEventHandlers,
	ProvideEventSubjects,

	consumer.NewConsumer,
	usecase.NewEventUsecase,
	usecase.NewInboxDeadLetterScanner,
	usecase.NewNotifyUsecase,
	usecase.NewRateLimitUsecase,
	usecase.NewStationMessageUsecase,
)

func ProvideEventHandlers(
	articleCollectedHandler *handler.ArticleCollectedHandler,
	articleLikedHandler *handler.ArticleLikedHandler,
	articlePublishedHandler *handler.ArticlePublishedHandler,
	articleThankedHandler *handler.ArticleThankedHandler,
	articleWatchedHandler *handler.ArticleWatchedHandler,
	commentLikedHandler *handler.CommentLikedHandler,
	commentPublishedHandler *handler.CommentPublishedHandler,
	userFollowHandler *handler.UserFollowHandler,
	userRegisterHandler *handler.UserRegisterHandler,
	userVerificationCodeHandler *handler.UserVerificationCodeHandler,
) usecase.EventHandlers {
	eventHandlers := usecase.EventHandlers{
		commonenum.EventTypeContentArticleCollect:     articleCollectedHandler,
		commonenum.EventTypeContentArticleLike:        articleLikedHandler,
		commonenum.EventTypeContentArticlePublish:     articlePublishedHandler,
		commonenum.EventTypeContentArticleThank:       articleThankedHandler,
		commonenum.EventTypeContentArticleWatch:       articleWatchedHandler,
		commonenum.EventTypeContentCommentLike:        commentLikedHandler,
		commonenum.EventTypeContentCommentPublish:     commentPublishedHandler,
		commonenum.EventTypeUserFollow:                userFollowHandler,
		commonenum.EventTypeUserRegister:              userRegisterHandler,
		commonenum.EventTypeUserEmailVerificationCode: userVerificationCodeHandler,
		commonenum.EventTypeUserPhoneVerificationCode: userVerificationCodeHandler,
	}

	return eventHandlers
}

func ProvideEventSubjects() usecase.EventSubjects {
	return usecase.EventSubjects{
		commonenum.EventSubjectContentArticleCollect,
		commonenum.EventSubjectContentArticleLike,
		commonenum.EventSubjectContentArticlePublish,
		commonenum.EventSubjectContentArticleThank,
		commonenum.EventSubjectContentArticleWatch,
		commonenum.EventSubjectContentCommentLike,
		commonenum.EventSubjectContentCommentPublish,
		commonenum.EventSubjectUserFollow,
		commonenum.EventSubjectUserRegister,
		commonenum.EventSubjectUserEmailVerificationCode,
		commonenum.EventSubjectUserPhoneVerificationCode,
	}
}
