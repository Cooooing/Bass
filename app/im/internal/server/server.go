package server

import (
	"github.com/google/wire"
)

// ServerProviderSet 是 server 层依赖集合。
var ServerProviderSet = wire.NewSet(
	NewGRPCServer,
	NewHTTPServer,
)
