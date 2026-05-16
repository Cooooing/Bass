package enum

type NotificationChannel string

const (
	NotificationChannelEmail   NotificationChannel = "NOTIFICATION_CHANNEL_EMAIL"
	NotificationChannelSMS     NotificationChannel = "NOTIFICATION_CHANNEL_SMS"
	NotificationChannelStation NotificationChannel = "NOTIFICATION_CHANNEL_STATION"
)

func (NotificationChannel) Values() []string {
	return []string{
		string(NotificationChannelEmail),
		string(NotificationChannelSMS),
		string(NotificationChannelStation),
	}
}
