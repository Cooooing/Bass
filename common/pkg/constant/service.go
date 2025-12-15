package constant

type ServiceName string

const (
	UserServiceName    ServiceName = "user-service"
	ContentServiceName ServiceName = "content-service"
	NotifyServiceName  ServiceName = "notify-service"
)

func (s ServiceName) String() string {
	return string(s)
}
