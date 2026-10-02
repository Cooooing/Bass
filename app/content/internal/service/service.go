package service

import (
	"common/pkg/server"

	"github.com/google/wire"
)

var ServiceProviderSet = wire.NewSet(
	ProvideServices,
	NewCommonSystemService,
	NewArticleService,
	NewArticleSearchService,
	NewPostscriptService,
	NewCommentService,
	NewDomainService,
	NewOutboxService,
	NewTagService,
	NewMoonbreezeService,
)

func ProvideServices(
	commonSystemService *CommonSystemService,
	articleService *ArticleService,
	articleSearchService *ArticleSearchService,
	postscriptService *PostscriptService,
	domainService *DomainService,
	commentService *CommentService,
	tagService *TagService,
	outboxService *OutboxService,
	moonbreezeService *MoonbreezeService,
) []server.Service {
	return []server.Service{
		commonSystemService,
		articleService,
		articleSearchService,
		postscriptService,
		domainService,
		commentService,
		tagService,
		outboxService,
		moonbreezeService,
	}
}
