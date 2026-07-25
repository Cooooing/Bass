package model

type NotificationTemplateData interface {
	notificationTemplateData()
}

type TemplateUser struct {
	ID       int64
	Name     string
	Nickname string
}

type TemplateArticle struct {
	ID     int64
	Title  string
	Author TemplateUser
}

type TemplateComment struct {
	ID        int64
	ArticleID int64
	Content   string
	User      TemplateUser
	ReplyUser TemplateUser
	Article   TemplateArticle
}

type UserRegisterTemplateData struct {
	User TemplateUser
}

func (UserRegisterTemplateData) notificationTemplateData() {}

type VerificationCodeTemplateData struct {
	Code           string
	ExpiresSeconds int64
}

func (VerificationCodeTemplateData) notificationTemplateData() {}

func (d VerificationCodeTemplateData) ExpiresMinutes() int64 {
	expiresMinutes := d.ExpiresSeconds / 60
	if expiresMinutes <= 0 {
		return 1
	}
	return expiresMinutes
}

type UserFollowTemplateData struct {
	Follower TemplateUser
	Followed TemplateUser
}

func (UserFollowTemplateData) notificationTemplateData() {}

type ArticlePublishedTemplateData struct {
	Article TemplateArticle
}

func (ArticlePublishedTemplateData) notificationTemplateData() {}

type ArticleActorTemplateData struct {
	Article TemplateArticle
	Actor   TemplateUser
}

func (ArticleActorTemplateData) notificationTemplateData() {}

type CommentPublishedTemplateData struct {
	Comment TemplateComment
}

func (CommentPublishedTemplateData) notificationTemplateData() {}

type CommentLikedTemplateData struct {
	Comment TemplateComment
	Actor   TemplateUser
}

func (CommentLikedTemplateData) notificationTemplateData() {}
