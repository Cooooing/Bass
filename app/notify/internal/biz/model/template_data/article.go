package template_data

type ArticlePublished struct {
	Article Article
}

func (ArticlePublished) notificationTemplateData() {}

type ArticleActor struct {
	Article Article
	Actor   User
}

func (ArticleActor) notificationTemplateData() {}
