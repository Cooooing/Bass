package service

import (
	"common/pkg/server"

	"github.com/google/wire"
)

// ServiceProviderSet 是 service 层依赖集合。
var ServiceProviderSet = wire.NewSet(
	NewCommonSystemService,
	NewArticleService,
	NewCommentService,
	NewDomainService,
	NewTagService,

	ProvideServices,
	ProvideHttpServices,
)

func ProvideServices(
	commonSystemService *CommonSystemService,
	articleService *ArticleService,
	domainService *DomainService,
	commentService *CommentService,
	tagService *TagService,
) []server.GrpcService {
	return []server.GrpcService{
		commonSystemService,
		articleService,
		domainService,
		commentService,
		tagService,
	}
}

func ProvideHttpServices(
	commonSystemService *CommonSystemService,
	articleService *ArticleService,
	domainService *DomainService,
	commentService *CommentService,
	tagService *TagService,
) []server.HttpService {
	return []server.HttpService{
		commonSystemService,
		articleService,
		domainService,
		commentService,
		tagService,
	}
}
