package template_data

type NotificationTemplateData interface {
	notificationTemplateData()
}

type User struct {
	ID       int64
	Name     string
	Nickname string
}

type Article struct {
	ID     int64
	Title  string
	Author User
}

type Comment struct {
	ID        int64
	ArticleID int64
	Content   string
	User      User
	ReplyUser User
	Article   Article
}
