package data

import (
	"content/internal/data/client"

	"github.com/google/wire"
)

// DataProviderSet is data providers.
var DataProviderSet = wire.NewSet(

	client.NewEtcdClient,
	client.NewDataBaseClient,

	NewGreeterRepo,
)
