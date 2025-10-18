package client

import (
	"context"
	"fmt"
	"user/internal/conf"
	"user/internal/data/ent/gen"

	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/lib/pq"
)

type DatabaseClient struct {
	log     *log.Helper
	clients map[string]*gen.Client
}

// NewDefault 创建默认的数据库客户端
func NewDefault(dbClient *DatabaseClient) *gen.Client {
	return dbClient.GetClient("default")
}

func NewDataBaseClient(log *log.Helper, conf *conf.Bootstrap) (*DatabaseClient, func(), error) {
	var dbClient = &DatabaseClient{
		log:     log,
		clients: make(map[string]*gen.Client),
	}

	for s, connections := range conf.Data.Database {
		err := dbClient.newEntClient(s, connections.Driver, connections.Source, connections.Merge)
		if err != nil {
			panic(fmt.Errorf("failed to new ent client: %w", err))
		}
	}

	cleanup := dbClient.CleanUp
	return dbClient, cleanup, nil
}

// GetClient 根据名称获取客户端
func (c *DatabaseClient) GetClient(name string) *gen.Client {
	return c.clients[name]
}

func (c *DatabaseClient) newEntClient(name string, driver string, source string, merge bool) error {
	logFunc := func(args ...any) {
		text := fmt.Sprint(args[0])
		c.log.Debugf("%s", text)
	}

	client, err := gen.Open(driver, source, gen.Log(logFunc))
	if err != nil {
		return fmt.Errorf("failed to open db: %w", err)
	}
	c.log.Infof("database: ent created database client [%s]", name)
	client = client.Debug()
	// 可选：自动迁移
	if merge {
		ctx := context.Background()
		if err := client.Schema.Create(ctx); err != nil {
			return fmt.Errorf("failed creating schema resources: %w", err)
		}
	}
	c.clients[name] = client
	return nil
}

func (c *DatabaseClient) CleanUp() {
	for name, client := range c.clients {
		err := client.Close()
		if err != nil {
			log.Errorf("failed to close %s db: %s", name, err.Error())
		}
	}
	c.log.Infof("database client closed")
}
