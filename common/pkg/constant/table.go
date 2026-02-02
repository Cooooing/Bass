package constant

type TablePrefix string

func (t TablePrefix) String() string {
	return string(t)
}

const (
	TablePrefixTemplate TablePrefix = "template_"
	TablePrefixUser     TablePrefix = "user_"
	TablePrefixContent  TablePrefix = "content_"
	TablePrefixNotify   TablePrefix = "notify_"
	TablePrefixSignal   TablePrefix = "signal_"
	TablePrefixIM       TablePrefix = "im_"
	TablePrefixInfra    TablePrefix = "infra_"
)
