package constant

type TablePrefix string

func (t TablePrefix) String() string {
	return string(t)
}

const (
	TablePrefixUser    TablePrefix = "user_"
	TablePrefixContent TablePrefix = "content_"
)
