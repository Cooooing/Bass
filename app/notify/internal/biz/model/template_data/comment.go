package template_data

type CommentPublished struct {
	Comment Comment
}

func (CommentPublished) notificationTemplateData() {}

type CommentLiked struct {
	Comment Comment
	Actor   User
}

func (CommentLiked) notificationTemplateData() {}
