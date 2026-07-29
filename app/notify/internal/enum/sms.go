package enum

type SMSType string

const (
	SMSTypeTencent SMSType = "tencent"
)

func (e SMSType) String() string {
	return string(e)
}
