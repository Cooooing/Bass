package service

import (
	"common/pkg/util/server"

	"github.com/google/wire"
)

// ServiceProviderSet is service providers.
var ServiceProviderSet = wire.NewSet(
	NewSystemService,
	NewArticleService,
	NewCommentService,
	NewDomainService,
	NewTagService,

	ProvideServices,
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
