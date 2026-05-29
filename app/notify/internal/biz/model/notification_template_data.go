package model

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

type UserFollowTemplateData struct {
	Follower TemplateUser
	Followed TemplateUser
}

type ArticlePublishedTemplateData struct {
	Article TemplateArticle
}

type ArticleActorTemplateData struct {
	Article TemplateArticle
	Actor   TemplateUser
}

type CommentPublishedTemplateData struct {
	Comment TemplateComment
}

type CommentLikedTemplateData struct {
	Comment TemplateComment
	Actor   TemplateUser
}
