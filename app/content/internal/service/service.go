package service

import (
	"common/pkg/server"

	"github.com/google/wire"
)

var ServiceProviderSet = wire.NewSet(
	ProvideServices,
	NewCommonSystemService,
	NewArticleService,
	NewPostscriptService,
	NewCommentService,
	NewDomainService,
	NewOutboxService,
	NewTagService,
)

func ProvideServices(commonSystemService *CommonSystemService, articleService *ArticleService, postscriptService *PostscriptService, domainService *DomainService, commentService *CommentService, tagService *TagService, outboxService *OutboxService) []server.Service {
	return []server.Service{
		commonSystemService,
		articleService,
		postscriptService,
		domainService,
		commentService,
		tagService,
		outboxService,
	}
}
