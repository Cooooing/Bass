package constant

type ServiceName string

const (
	GatewayServiceName ServiceName = "gateway"
	UserServiceName    ServiceName = "user"
	ContentServiceName ServiceName = "content"
	NotifyServiceName  ServiceName = "notify"
	IMServiceName      ServiceName = "im"
	SignalServiceName  ServiceName = "signal"
	InfraServiceName   ServiceName = "infra"
)

func (s ServiceName) String() string {
	return string(s)
}
