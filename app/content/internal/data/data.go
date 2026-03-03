package data

import (
	commonClient "common/pkg/client"
	commonModel "common/pkg/model"
	"common/pkg/util/jwt"
	"content/internal/conf"
	"content/internal/data/base"
	"content/internal/data/client"
	"content/internal/data/repo"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/jinzhu/copier"
)

// DataProviderSet is data providers.
var DataProviderSet = wire.NewSet(
	base.NewBaseData,

	client.NewDataBaseClient,
	NewConsulClient,
	NewRedisClient,
	NewRabbitMQClient,

	repo.NewArticleRepo,
	repo.NewCommentRepo,
	repo.NewCommentActionRecordRepo,
	repo.NewArticlePostscriptRepo,
	repo.NewArticleActionRecordRepo,
	repo.NewDomainRepo,
	repo.NewTagRepo,

	jwt.NewTokenCache,
)

func NewConsulClient(log *log.Helper, conf *conf.Bootstrap) (*commonClient.ConsulClient, func(), error) {
	c := &commonModel.ConsulConf{}
	err := copier.Copy(c, conf.Data.Consul)
	if err != nil {
		return nil, nil, err
	}
	return commonClient.NewConsulClient(log, c)
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
