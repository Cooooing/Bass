package service

import (
	"common/pkg/server"

	"github.com/google/wire"
)

var ServiceProviderSet = wire.NewSet(
	ProvideServices,
	NewCommonSystemService,
	NewArticleService,
	NewCommentService,
	NewDomainService,
	NewTagService,
)

func ProvideServices(commonSystemService *CommonSystemService, articleService *ArticleService, domainService *DomainService, commentService *CommentService, tagService *TagService) []server.Service {
	return []server.Service{
		commonSystemService,
		articleService,
		domainService,
		commentService,
		tagService,
	}
}
