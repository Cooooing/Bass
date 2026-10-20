package server

import (
	"common/pkg/client"
	"realtime_hub/internal/biz/usecase"

	"github.com/go-kratos/kratos/v3/transport"
	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"github.com/google/wire"
)

var ServerProviderSet = wire.NewSet(ProvideServers, NewGRPCServer, NewHTTPServer, NewNotificationDeliveryConsumer, NewControlConsumer)

func ProvideServers(
	notificationConsumer *NotificationDeliveryConsumer,
	controlConsumer *ControlConsumer,
	grpcServer *grpc.Server,
	httpServer *http.Server,
) []transport.Server {
	return []transport.Server{notificationConsumer, controlConsumer, grpcServer, httpServer}
}

func NewNotificationDeliveryConsumer(
	nats client.Subscriber,
	realtime *usecase.RealtimeUsecase,
) *NotificationDeliveryConsumer {
	return &NotificationDeliveryConsumer{nats: nats, realtime: realtime}
}
