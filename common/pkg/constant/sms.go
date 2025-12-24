package constant

type SMSType string

func (s SMSType) String() string {
	return string(s)
}

const (
	SMSTypeTencent SMSType = "tencent"
)
