package service

import (
	"common/pkg/constant"
	commonModel "common/pkg/model"
	"common/pkg/util"
	"common/pkg/util/server"
	"context"

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

func withUserID(ctx context.Context, userID int64) context.Context {
	return util.SetContextValue[*commonModel.User](ctx, constant.CtxUserInfo, &commonModel.User{ID: userID})
}
