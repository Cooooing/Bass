package template_data

type VerificationCode struct {
	Code           string
	ExpiresSeconds int64
}

func (VerificationCode) notificationTemplateData() {}

func (d VerificationCode) ExpiresMinutes() int64 {
	expiresMinutes := d.ExpiresSeconds / 60
	if expiresMinutes <= 0 {
		return 1
	}
	return expiresMinutes
}
