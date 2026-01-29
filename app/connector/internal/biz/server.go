package biz

import (
	"common/pkg/client"
	"connector/internal/biz/base"
	"connector/internal/biz/cache"
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

type ServerDomain struct {
	*base.BaseDomain
	sessionCache cache.SessionCache

	ctx        context.Context
	httpClient *http.Client
}

func NewServerDomain(baseDomain *base.BaseDomain, sessionCache cache.SessionCache) (*ServerDomain, func()) {
	s := &ServerDomain{
		BaseDomain:   baseDomain,
		sessionCache: sessionCache,
		httpClient:   client.NewHttpClient(),
		ctx:          context.Background(),
	}
	return s, s.cleanup
}

func (d *ServerDomain) Register() error {
	url := fmt.Sprintf("http://127.0.0.1:8000/api/signal/v1/node/register")

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
			d.Log.Errorf("failed to close body: %v", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to register node: %s", url)
	}

	return nil
}

func (d *ServerDomain) Unregister() error {
	url := fmt.Sprintf("http://127.0.0.1:8000/api/signal/v1/node/unregister")

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
			d.Log.Errorf("failed to close body: %v", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to unregister node: %s", url)
	}
	return nil
}

func (d *ServerDomain) Run() {
	// 单机模式 Todo 集群模式需分布式任务
	f := func() {
		err := d.Register()
		if err != nil {
			d.Log.Errorf("failed to register node: %v", err)
		}
	}
	interval := 60 * time.Second
	if d.Conf.Server.Cluster {

	} else {
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
		d.Log.Errorf("failed to unregister node: %v", err)
	}
}
