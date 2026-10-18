package constant

type ServiceName string

const (
	UserServiceName      ServiceName = "user"
	ContentServiceName   ServiceName = "content"
	EconomyServiceName   ServiceName = "economy"
	NotifyServiceName    ServiceName = "notify"
	IMServiceName        ServiceName = "im"
	PlatformServiceName  ServiceName = "platform"
	SchedulerServiceName ServiceName = "scheduler"
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
	TablePrefixPlatform  TablePrefix = TablePrefix(PlatformServiceName + "_")
	TablePrefixScheduler TablePrefix = TablePrefix(SchedulerServiceName + "_")
	TablePrefixGameIdle  TablePrefix = TablePrefix(GameIdleServiceName + "_")
)

const (
	RealtimeNodeSubjectPrefix = "realtime.node."
)

func GetRealtimeNodeSubject(realm, nodeID string) string {
	return RealtimeNodeSubjectPrefix + realm + "." + nodeID
}
