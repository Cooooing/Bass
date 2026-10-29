package service

import (
	"common/pkg/server"

	"github.com/google/wire"
)

var ServiceProviderSet = wire.NewSet(
	ProvideServices,
	NewCommonSystemService,
	NewArticleService,
	NewStatisticsService,
	NewArticleSearchService,
	NewPostscriptService,
	NewCommentService,
	NewDomainService,
	NewOutboxService,
	NewTagService,
	NewMoonbreezeService,
	NewTCCArticleThankService,
	NewTCCArticleRewardService,
	NewTCCArticlePublishService,
	NewTCCCommentThankService,
)

func ProvideServices(
	commonSystemService *CommonSystemService,
	articleService *ArticleService,
	statisticsService *StatisticsService,
	articleSearchService *ArticleSearchService,
	postscriptService *PostscriptService,
	domainService *DomainService,
	commentService *CommentService,
	tagService *TagService,
	outboxService *OutboxService,
	moonbreezeService *MoonbreezeService,
	tccArticleThankService *TCCArticleThankService,
	tccArticleRewardService *TCCArticleRewardService,
	tccArticlePublishService *TCCArticlePublishService,
	tccCommentThankService *TCCCommentThankService,
) []server.Service {
	return []server.Service{
		commonSystemService,
		articleService,
		statisticsService,
		articleSearchService,
		postscriptService,
		domainService,
		commentService,
		tagService,
		outboxService,
		moonbreezeService,
		tccArticleThankService,
		tccArticleRewardService,
		tccArticlePublishService,
		tccCommentThankService,
	}
}
