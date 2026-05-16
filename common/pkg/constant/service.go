package constant

type ServiceName string

const (
	GatewayServiceName ServiceName = "gateway"
	UserServiceName    ServiceName = "user"
	ContentServiceName ServiceName = "content"
	NotifyServiceName  ServiceName = "notify"
	IMServiceName      ServiceName = "im"
	SignalServiceName  ServiceName = "signal"
)

func (s ServiceName) String() string {
	return string(s)
}

type TablePrefix string

func (t TablePrefix) String() string {
	return string(t)
}

const (
	TablePrefixUser    TablePrefix = TablePrefix(UserServiceName + "_")
	TablePrefixContent TablePrefix = TablePrefix(ContentServiceName + "_")
	TablePrefixNotify  TablePrefix = TablePrefix(NotifyServiceName + "_")
	TablePrefixSignal  TablePrefix = TablePrefix(SignalServiceName + "_")
	TablePrefixIM      TablePrefix = TablePrefix(IMServiceName + "_")
)
