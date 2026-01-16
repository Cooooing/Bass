package constant

type ServiceName string

const (
	GatewayServiceName ServiceName = "gateway-service"
	UserServiceName    ServiceName = "user-service"
	ContentServiceName ServiceName = "content-service"
	NotifyServiceName  ServiceName = "notify-service"
	IMServiceName      ServiceName = "im-service"
	SignalServiceName  ServiceName = "signal-service"
)

func (s ServiceName) String() string {
	return string(s)
}
