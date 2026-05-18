package service

import (
	"common/pkg/util/server"

	"github.com/google/wire"
)

// ServiceProviderSet is service providers.
var ServiceProviderSet = wire.NewSet(
	NewSystemService,
	NewCallbackService,
	ProvideServices,
)

func ProvideServices(
	systemService *SystemService,
	callbackService *CallbackService,
) []server.GrpcService {
	return []server.GrpcService{
		systemService,
		callbackService,
	}
}
