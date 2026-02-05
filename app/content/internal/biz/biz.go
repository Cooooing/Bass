package biz

import (
	"common/pkg/util"
	"content/internal/biz/base"
	"content/internal/biz/domain"

	"github.com/google/wire"
)

// BizProviderSet is biz providers.
var BizProviderSet = wire.NewSet(
	base.NewBaseDomain,
	util.NewEventPool,

	domain.NewArticleDomain,
	domain.NewCommentDomain,
	domain.NewDomainDomain,
	domain.NewTagDomain,
)
