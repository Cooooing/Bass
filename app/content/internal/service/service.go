package service

import (
	"common/pkg/server"

	"github.com/google/wire"
)

// ServiceProviderSet 是 service 层依赖集合。
var ServiceProviderSet = wire.NewSet(
	NewSystemService,
	NewArticleService,
	NewCommentService,
	NewDomainService,
	NewTagService,

	ProvideServices,
	ProvideHttpServices,
)

func ProvideServices(
	systemService *SystemService,
	articleService *ArticleService,
	domainService *DomainService,
	commentService *CommentService,
	tagService *TagService,
) []server.GrpcService {
	return []server.GrpcService{
		systemService,
		articleService,
		domainService,
		commentService,
		tagService,
	}
}

func ProvideHttpServices(
	systemService *SystemService,
	articleService *ArticleService,
	domainService *DomainService,
	commentService *CommentService,
	tagService *TagService,
) []server.HttpService {
	return []server.HttpService{
		systemService,
		articleService,
		domainService,
		commentService,
		tagService,
	}
}
