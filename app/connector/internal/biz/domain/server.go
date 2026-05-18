package domain

import (
	"common/pkg/client"
	"common/pkg/constant"
	"common/pkg/model"
	"common/pkg/util"
	"common/pkg/util/task"
	"connector/internal/biz/cache"
	"connector/internal/conf"
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type ServerDomain struct {
	conf         *conf.Bootstrap
	log          *log.Helper
	sessionCache cache.SessionCache
	asynqCache   *task.AsynqCache
	producer     *client.Producer
	ctx          context.Context
	httpClient   *http.Client
}

func NewServerDomain(
	conf *conf.Bootstrap,
	logger log.Logger,
	sessionCache cache.SessionCache,
	asynqCache *task.AsynqCache,
	producer *client.Producer,
) (*ServerDomain, func()) {
	s := &ServerDomain{
		conf:         conf,
		log:          log.NewHelper(logger),
		sessionCache: sessionCache,
		asynqCache:   asynqCache,
		producer:     producer,
		httpClient:   client.NewHttpClient(),
		ctx:          context.Background(),
	}
	return s, s.cleanup
}

func (d *ServerDomain) Register() error {
	url := fmt.Sprintf("%s://%s/api/signal/v1/node/register", util.If(d.conf.Server.Mode == constant.Dev, "http", "https"), d.conf.Server.MasterUrl)

	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Signal-NodeKey", "main")

	resp, err := d.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			d.log.Errorf("failed to close body: %v", err)
		}
	}(resp.Body)

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to register node: %s %d %s", url, resp.StatusCode, string(bytes))
	}

	return nil
}

func (d *ServerDomain) Unregister() error {
	url := fmt.Sprintf("%s://%s/api/signal/v1/node/unregister", util.If(d.conf.Server.Mode == constant.Dev, "http", "https"), d.conf.Server.MasterUrl)

	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Signal-NodeKey", "main")

	resp, err := d.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			d.log.Errorf("failed to close body: %v", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to unregister node: %s", url)
	}
	return nil
}

func (d *ServerDomain) Run() {
	f := func() {
		err := d.Register()
		if err != nil {
			d.log.Errorf("failed to register node: %v", err)
		}
	}
	interval := 60 * time.Second
	if d.conf.Server.Cluster {
		// 集群模式，使用分布式定时任务
		registerTaskName := fmt.Sprintf("%s-%s", constant.TaskConnectorRegister.String(), d.conf.Server.Key)
		version := time.Now().UnixMilli()
		err := d.asynqCache.SetAsynqTaskVersion(d.ctx, registerTaskName, version, interval*2)
		if err != nil {
			d.log.Errorf("failed to set asynq task version: %v", err)
			return
		}
		err = d.producer.EnqueueContextTask(d.ctx, &model.Task{
			TaskName: registerTaskName,
			Version:  version,
			Interval: interval,
		})
		if err != nil {
			d.log.Errorf("failed to register node: %v", err)
			return
		}
	} else {
		// 单机模式，使用本地定时任务
		ticker := time.NewTicker(interval)
		f()
		for {
			select {
			case <-ticker.C:
				f()
			case <-d.ctx.Done():
				return
			}
		}
	}
}

func (d *ServerDomain) cleanup() {
	d.ctx.Done()
	err := d.Unregister()
	if err != nil {
		d.log.Errorf("failed to unregister node: %v", err)
	}
}
