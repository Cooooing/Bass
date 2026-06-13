package constant

type ServiceName string

const (
	UserServiceName        ServiceName = "user"
	ContentServiceName     ServiceName = "content"
	NotifyServiceName      ServiceName = "notify"
	IMServiceName          ServiceName = "im"
	IntegrationServiceName ServiceName = "integration"
	PushHubServiceName     ServiceName = "push_hub"
	PushNodeServiceName    ServiceName = "push_node"
	PlatformServiceName    ServiceName = "platform"
)

func (s ServiceName) String() string {
	return string(s)
}

type TablePrefix string

func (t TablePrefix) String() string {
	return string(t)
}

const (
	TablePrefixUser        TablePrefix = TablePrefix(UserServiceName + "_")
	TablePrefixContent     TablePrefix = TablePrefix(ContentServiceName + "_")
	TablePrefixNotify      TablePrefix = TablePrefix(NotifyServiceName + "_")
	TablePrefixIM          TablePrefix = TablePrefix(IMServiceName + "_")
	TablePrefixIntegration TablePrefix = TablePrefix(IntegrationServiceName + "_")
	TablePrefixPushHub     TablePrefix = TablePrefix(PushHubServiceName + "_")
	TablePrefixPushNode    TablePrefix = TablePrefix(PushNodeServiceName + "_")
	TablePrefixPlatform    TablePrefix = TablePrefix(PlatformServiceName + "_")
)
