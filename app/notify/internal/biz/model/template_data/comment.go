package template_data

type CommentPublished struct {
	Comment Comment
}

type CommentMention struct {
	Comment Comment
	Actor   User
}

func (CommentMention) notificationTemplateData() {}

type GlobalChatMention struct {
	Actor User
}

func (GlobalChatMention) notificationTemplateData() {}

func (CommentPublished) notificationTemplateData() {}

type CommentLiked struct {
	Comment Comment
	Actor   User
}

func (CommentLiked) notificationTemplateData() {}

type ArticleAnswerAccepted struct {
	Comment Comment
	Author  User
}

func (ArticleAnswerAccepted) notificationTemplateData() {}
