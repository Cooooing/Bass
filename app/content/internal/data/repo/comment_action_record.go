package repo

import (
	v1 "common/api/gen/content/v1"
	commonClient "common/pkg/client"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/conf"
	"content/internal/data/gen"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

var _ repo.CommentActionRecordRepo = (*CommentActionRecordRepo)(nil)

type CommentActionRecordRepo struct {
	conf   *conf.Bootstrap
	log    *log.Helper
	consul *commonClient.ConsulClient
	redis  *commonClient.RedisClient
	nats   *commonClient.NatsClient
}

func NewCommentActionRecordRepo(
	conf *conf.Bootstrap,
	logger log.Logger,
	consul *commonClient.ConsulClient,
	redis *commonClient.RedisClient,
	nats *commonClient.NatsClient,
) repo.CommentActionRecordRepo {
	return &CommentActionRecordRepo{
		conf:   conf,
		log:    log.NewHelper(logger),
		consul: consul,
		redis:  redis,
		nats:   nats,
	}
}

func (a CommentActionRecordRepo) Save(ctx context.Context, client *gen.Client, record *model.CommentActionRecord) (*model.CommentActionRecord, error) {
	// TODO: 待实现。
	panic("implement me")
}

func (a CommentActionRecordRepo) Delete(ctx context.Context, client *gen.Client, commentId int64, userId int64, action v1.CommentAction) error {
	// TODO: 待实现。
	panic("implement me")
}
