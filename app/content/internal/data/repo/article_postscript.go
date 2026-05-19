package repo

import (
	commonClient "common/pkg/client"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/conf"
	"content/internal/data/gen"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

var _ repo.ArticlePostscriptRepo = (*ArticlePostscriptRepo)(nil)

type ArticlePostscriptRepo struct {
	conf   *conf.Bootstrap
	log    *log.Helper
	consul *commonClient.ConsulClient
	redis  *commonClient.RedisClient
	nats   *commonClient.NatsClient
}

func NewArticlePostscriptRepo(
	conf *conf.Bootstrap,
	logger log.Logger,
	consul *commonClient.ConsulClient,
	redis *commonClient.RedisClient,
	nats *commonClient.NatsClient,
) repo.ArticlePostscriptRepo {
	return &ArticlePostscriptRepo{
		conf:   conf,
		log:    log.NewHelper(logger),
		consul: consul,
		redis:  redis,
		nats:   nats,
	}
}

func (a ArticlePostscriptRepo) Save(ctx context.Context, client *gen.Client, articlePostscript *model.ArticlePostscript) (*model.ArticlePostscript, error) {
	save, err := client.ArticlePostscript.Create().
		SetArticleID(articlePostscript.ArticleID).
		SetContent(articlePostscript.Content).
		Save(ctx)
	return &model.ArticlePostscript{ArticlePostscript: save}, err
}
