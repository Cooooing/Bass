package data

import (
	commonClient "common/pkg/client"
	commonModel "common/pkg/model"
	"common/pkg/util"
	"content/internal/conf"
	"content/internal/data/client"
	"content/internal/data/ent/gen"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/jinzhu/copier"
)

// DataProviderSet is data providers.
var DataProviderSet = wire.NewSet(
	NewBaseRepo,

	client.NewDataBaseClient,
	NewEtcdClient,
	NewRedisClient,
	NewRabbitMQClient,

	NewArticleRepo,
	NewArticlePostscriptRepo,
	NewArticleActionRecordRepo,
	NewDomainRepo,

	util.NewTokenRepo,
)

type BaseRepo struct {
	conf     *conf.Bootstrap
	log      *log.Helper
	db       *gen.Client
	etcd     *commonClient.EtcdClient
	redis    *commonClient.RedisClient
	rabbitmq *commonClient.RabbitMQClient
}

func NewBaseRepo(conf *conf.Bootstrap, log *log.Helper, db *gen.Client, etcd *commonClient.EtcdClient, redis *commonClient.RedisClient, rabbitmq *commonClient.RabbitMQClient) *BaseRepo {
	return &BaseRepo{
		conf:     conf,
		log:      log,
		etcd:     etcd,
		db:       db,
		redis:    redis,
		rabbitmq: rabbitmq,
	}
}

func NewEtcdClient(log *log.Helper, conf *conf.Bootstrap) (*commonClient.EtcdClient, func(), error) {
	c := &commonModel.EtcdConf{}
	err := copier.Copy(c, conf.Registry.Etcd)
	if err != nil {
		return nil, nil, err
	}
	return commonClient.NewEtcdClient(log, c)
}

func NewRedisClient(log *log.Helper, conf *conf.Bootstrap) (*commonClient.RedisClient, func(), error) {
	c := &commonModel.RedisConf{}
	err := copier.Copy(c, conf.Data.Redis)
	if err != nil {
		return nil, nil, err
	}
	return commonClient.NewRedisClient(log, c)
}

func NewRabbitMQClient(log *log.Helper, conf *conf.Bootstrap) (*commonClient.RabbitMQClient, func(), error) {
	c := &commonModel.RabbitmqConf{}
	err := copier.Copy(c, conf.Data.Rabbitmq)
	if err != nil {
		return nil, nil, err
	}
	return commonClient.NewRabbitMQClient(log, c)
}
