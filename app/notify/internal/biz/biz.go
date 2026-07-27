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
	handler.NewArticleThankedHandler,
	handler.NewArticleWatchedHandler,
	handler.NewCommentLikedHandler,
	handler.NewCommentPublishedHandler,
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
	articleThankedHandler *handler.ArticleThankedHandler,
	articleWatchedHandler *handler.ArticleWatchedHandler,
	commentLikedHandler *handler.CommentLikedHandler,
	commentPublishedHandler *handler.CommentPublishedHandler,
	userFollowHandler *handler.UserFollowHandler,
	userRegisterHandler *handler.UserRegisterHandler,
	userVerificationCodeHandler *handler.UserVerificationCodeHandler,
) map[commonenum.EventType]usecase.EventHandler {
	eventHandlers := map[commonenum.EventType]usecase.EventHandler{
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

func ProvideEventSubjects() []commonenum.EventSubject {
	return []commonenum.EventSubject{
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
