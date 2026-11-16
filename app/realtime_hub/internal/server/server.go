package server

import (
	nats "common/pkg/client/nats"

	"realtime_hub/internal/biz/usecase"

	"github.com/go-kratos/kratos/v3/transport"
	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"github.com/google/wire"
)

var ServerProviderSet = wire.NewSet(ProvideServers, NewGRPCServer, NewHTTPServer, NewNotificationDeliveryConsumer, NewPrivateChatDeliveryConsumer, NewGlobalChatDeliveryConsumer, NewControlConsumer)

func ProvideServers(
	notificationConsumer *NotificationDeliveryConsumer,
	privateChatConsumer *PrivateChatDeliveryConsumer,
	globalChatConsumer *GlobalChatDeliveryConsumer,
	controlConsumer *ControlConsumer,
	grpcServer *grpc.Server,
	httpServer *http.Server,
) []transport.Server {
	return []transport.Server{notificationConsumer, privateChatConsumer, globalChatConsumer, controlConsumer, grpcServer, httpServer}
}

func NewPrivateChatDeliveryConsumer(nats *nats.Client, realtime *usecase.RealtimeUsecase) *PrivateChatDeliveryConsumer {
	return &PrivateChatDeliveryConsumer{nats: nats, realtime: realtime}
}

func NewGlobalChatDeliveryConsumer(nats *nats.Client, realtime *usecase.RealtimeUsecase) *GlobalChatDeliveryConsumer {
	return &GlobalChatDeliveryConsumer{nats: nats, realtime: realtime}
}

func NewNotificationDeliveryConsumer(
	nats *nats.Client,
	realtime *usecase.RealtimeUsecase,
) *NotificationDeliveryConsumer {
	return &NotificationDeliveryConsumer{nats: nats, realtime: realtime}
}

func NewControlConsumer(nats *nats.Client, realtime *usecase.RealtimeUsecase) *ControlConsumer {
	return &ControlConsumer{nats: nats, realtime: realtime}
}
