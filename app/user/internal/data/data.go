package data

import (
	"user/internal/conf"
	"user/internal/data/client"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// DataProviderSet is data providers.
var DataProviderSet = wire.NewSet(
	NewBaseRepo,

	client.NewEtcdClient,
	client.NewDataBaseClient,
	client.NewDefault,
	client.NewRedisClient,
	client.NewRabbitMQClient,

	NewUserRepo,
	NewTokenRepo,
)

type BaseRepo struct {
	conf     *conf.Bootstrap
	log      *log.Helper
	etcd     *client.EtcdClient
	db       *client.DatabaseClient
	redis    *client.RedisClient
	rabbitmq *client.RabbitMQClient
}

func NewBaseRepo(conf *conf.Bootstrap, log *log.Helper, etcd *client.EtcdClient, db *client.DatabaseClient, redis *client.RedisClient, rabbitmq *client.RabbitMQClient) *BaseRepo {
	return &BaseRepo{
		conf:     conf,
		log:      log,
		etcd:     etcd,
		db:       db,
		redis:    redis,
		rabbitmq: rabbitmq,
	}
}
