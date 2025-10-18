package client

import (
	"content/internal/conf"
	"content/internal/data/ent"
	"context"
	"fmt"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/lib/pq"
)

type DatabaseClient struct {
	log           *log.Helper
	clients       map[string]*ent.Client
	defaultClient *ent.Client
}

func NewDataBaseClient(log *log.Helper, conf *conf.Bootstrap) (*DatabaseClient, func(), error) {
	var dbClient = &DatabaseClient{
		log:     log,
		clients: make(map[string]*ent.Client),
	}

	for s, connections := range conf.Data.Database {
		err := dbClient.newEntClient(s, connections.Driver, connections.Source, true)
		if err != nil {
			panic(fmt.Errorf("failed to new ent client: %w", err))
		}
	}
	dbClient.defaultClient = dbClient.GetClient("default")

	cleanup := dbClient.CleanUp
	return dbClient, cleanup, nil
}

// GetClient 根据名称获取客户端
func (c *DatabaseClient) GetClient(name string) *ent.Client {
	return c.clients[name]
}

// Default 返回默认数据库客户端
func (c *DatabaseClient) Default() *ent.Client {
	return c.defaultClient
}

func (c *DatabaseClient) newEntClient(name string, driver string, source string, merge bool) error {
	logFunc := func(args ...any) {
		text := fmt.Sprint(args[0])
		text = strings.ReplaceAll(strings.ReplaceAll(text, "\t", " "), "\n", " ") // 将换行替换为空格
		c.log.Debugf("%s", text)
	}

	client, err := ent.Open(driver, source, ent.Log(logFunc))
	if err != nil {
		return fmt.Errorf("failed to open db: %w", err)
	}
	c.log.Infof("ent: created database client %s", name)
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
}
