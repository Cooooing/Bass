package template_data

type ArticlePublished struct {
	Article Article
}

func (ArticlePublished) notificationTemplateData() {}

type ArticleActor struct {
	Article Article
	Actor   User
}

type ArticleMention struct {
	Article Article
	Actor   User
}

func (ArticleMention) notificationTemplateData() {}

type PostscriptMention struct {
	Article Article
	Actor   User
}

func (PostscriptMention) notificationTemplateData() {}

func (ArticleActor) notificationTemplateData() {}

type ArticleLotteryWon struct {
	Article Article
	Author  User
	Prize   string
}

func (ArticleLotteryWon) notificationTemplateData() {}

type ArticleWatchActivity struct {
	Article  Article
	Actor    User
	Activity string
}

func (ArticleWatchActivity) notificationTemplateData() {}
