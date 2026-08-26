package constant

type ServiceName string

const (
	UserServiceName      ServiceName = "user"
	ContentServiceName   ServiceName = "content"
	EconomyServiceName   ServiceName = "economy"
	NotifyServiceName    ServiceName = "notify"
	IMServiceName        ServiceName = "im"
	PushHubServiceName   ServiceName = "push_hub"
	PushNodeServiceName  ServiceName = "push_node"
	PlatformServiceName  ServiceName = "platform"
	SchedulerServiceName ServiceName = "scheduler"
	GameTownServiceName  ServiceName = "game_town"
	GameIdleServiceName  ServiceName = "game_idle"
)

func (s ServiceName) String() string {
	return string(s)
}

type TablePrefix string

func (t TablePrefix) String() string {
	return string(t)
}

const (
	TablePrefixUser      TablePrefix = TablePrefix(UserServiceName + "_")
	TablePrefixContent   TablePrefix = TablePrefix(ContentServiceName + "_")
	TablePrefixEconomy   TablePrefix = TablePrefix(EconomyServiceName + "_")
	TablePrefixNotify    TablePrefix = TablePrefix(NotifyServiceName + "_")
	TablePrefixIM        TablePrefix = TablePrefix(IMServiceName + "_")
	TablePrefixPushHub   TablePrefix = TablePrefix(PushHubServiceName + "_")
	TablePrefixPushNode  TablePrefix = TablePrefix(PushNodeServiceName + "_")
	TablePrefixPlatform  TablePrefix = TablePrefix(PlatformServiceName + "_")
	TablePrefixScheduler TablePrefix = TablePrefix(SchedulerServiceName + "_")
	TablePrefixGameTown  TablePrefix = TablePrefix(GameTownServiceName + "_")
	TablePrefixGameIdle  TablePrefix = TablePrefix(GameIdleServiceName + "_")
)

const (
	PushNodeSubjectPrefix = "push.node."
)

func GetPushNodeSubject(nodeID string) string {
	return PushNodeSubjectPrefix + nodeID
}
