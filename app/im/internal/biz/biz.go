package biz

import (
	"common/pkg/util"
	domainbase "im/internal/biz/base"
	"im/internal/biz/domain"

	"github.com/google/wire"
)

// BizProviderSet is biz providers.
var BizProviderSet = wire.NewSet(
	domainbase.NewBaseDomain,
	util.NewTokenCache,

	domain.NewChatGroupDomain,
	domain.NewChatSessionDomain,
	domain.NewChatMessageDomain,
)
