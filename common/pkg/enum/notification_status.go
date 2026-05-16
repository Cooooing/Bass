package enum

type NotificationStatus string

const (
	NotificationStatusNormal NotificationStatus = "NOTIFICATION_STATUS_NORMAL"
	NotificationStatusCancel NotificationStatus = "NOTIFICATION_STATUS_CANCEL"
)

func (NotificationStatus) Values() []string {
	return []string{
		string(NotificationStatusNormal),
		string(NotificationStatusCancel),
	}
}
