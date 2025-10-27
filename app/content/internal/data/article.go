package data

import (
	cv1 "common/api/common/v1"
	v1 "common/api/content/v1"
	"common/pkg/constant"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/ent/gen"
	"content/internal/data/ent/gen/article"
	"context"
	"encoding/json"

	"github.com/jinzhu/copier"
)

type ArticleRepo struct {
	*BaseRepo
	client *gen.Client
}

func NewArticleRepo(baseRepo *BaseRepo, client *gen.Client) repo.ArticleRepo {
	return &ArticleRepo{
		BaseRepo: baseRepo,
		client:   client,
	}
}

func (r *ArticleRepo) Save(ctx context.Context, client *gen.Client, article *model.Article) (*model.Article, error) {
	save, err := client.Article.Create().
		SetTitle(article.Title).
		SetContent(article.Content).
		SetAcceptedAnswerID(article.AcceptedAnswerID).
		Save(ctx)
	return (*model.Article)(save), err
}

func (r *ArticleRepo) UpdateContent(ctx context.Context, client *gen.Client, articleId int, content string) error {
	return client.Article.UpdateOneID(articleId).
		SetContent(content).
		Exec(ctx)
}
func (r *ArticleRepo) UpdateStatus(ctx context.Context, client *gen.Client, articleId int, status cv1.ArticleStatus) error {
	return client.Article.UpdateOneID(articleId).
		SetStatus(int(status)).
		Exec(ctx)
}

func (r *ArticleRepo) UpdateHasPostscript(ctx context.Context, client *gen.Client, articleId int, hasPostscript bool) error {
	return client.Article.UpdateOneID(articleId).
		SetHasPostscript(hasPostscript).
		Exec(ctx)
}

func (r *ArticleRepo) UpdateStat(ctx context.Context, client *gen.Client, articleId int, action cv1.ArticleAction, num int) error {
	updateOne := client.Article.UpdateOneID(articleId)
	switch action {
	case cv1.ArticleAction_ArticleActionLike:
		updateOne.AddLikeCount(num)
	case cv1.ArticleAction_ArticleActionThank:
		updateOne.AddThankCount(num)
	case cv1.ArticleAction_ArticleActionCollect:
		updateOne.AddCollectCount(num)
	case cv1.ArticleAction_ArticleActionWatch:
		updateOne.AddWatchCount(num)
	}
	return updateOne.Exec(ctx)
}

func (r *ArticleRepo) Delete(ctx context.Context, articleId int) error {
	// TODO implement me
	panic("implement me")
}

func (r *ArticleRepo) GetArticleById(ctx context.Context, client *gen.Client, id int) (*model.Article, error) {
	query, err := client.Article.Query().
		Where(article.IDEQ(id)).
		WithPostscripts().
		WithTags().
		WithComments().
		First(ctx)
	return (*model.Article)(query), err
}

func (r *ArticleRepo) Publish(ctx context.Context, client *gen.Client, articleId int) error {
	first, err := r.GetArticleById(ctx, client, articleId)
	if err != nil {
		return err
	}
	publish := &v1.ArticleEventPublish{}
	err = copier.Copy(&publish, first)
	if err != nil {
		return err
	}
	marshal, err := json.Marshal(publish)
	if err != nil {
		return err
	}
	err = r.rabbitmq.Publish(constant.ExchangeContent.String(), constant.RoutingKeyArticleCreate.String(), marshal)
	return err
}
