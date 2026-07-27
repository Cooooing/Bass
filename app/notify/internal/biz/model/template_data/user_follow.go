package template_data

type UserFollow struct {
	Follower User
	Followed User
}

func (UserFollow) notificationTemplateData() {}
