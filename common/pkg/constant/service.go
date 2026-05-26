package constant

type ServiceName string

const (
	UserServiceName    ServiceName = "user"
	ContentServiceName ServiceName = "content"
	NotifyServiceName  ServiceName = "notify"
	IMServiceName      ServiceName = "im"
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
	TablePrefixIM      TablePrefix = TablePrefix(IMServiceName + "_")
)
