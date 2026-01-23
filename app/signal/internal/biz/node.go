package biz

import (
	cv1 "common/api/common/v1"
	"common/pkg/constant"
	commonBase "common/pkg/cutil/base"
	"common/pkg/cutil/base/str"
	"common/pkg/util"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"signal/internal/biz/base"
	"signal/internal/biz/model"
	"signal/internal/biz/repo"
	"signal/internal/biz/task"
	"time"

	"github.com/sony/sonyflake/v2"
)

type NodeDomain struct {
	*base.BaseDomain
	nodeRepo   repo.NodeRepo
	producer   *Producer
	httpClient *http.Client
	sf         *sonyflake.Sonyflake
}

func NewNodeDomain(baseDomain *base.BaseDomain, nodeRepo repo.NodeRepo, producer *Producer, httpClient *http.Client) (*NodeDomain, error) {
	sf, err := str.NewSonyflake()
	if err != nil {
		return nil, err
	}
	return &NodeDomain{
		BaseDomain: baseDomain,
		nodeRepo:   nodeRepo,
		producer:   producer,
		httpClient: httpClient,
		sf:         sf,
	}, nil
}

// GenerateSecret 生成一个 32 位随机字符串
func (d *NodeDomain) GenerateSecret() string {
	return str.RandStr(d.sf, 32, true, true, true, false)
}

func (d *NodeDomain) Ping(node *model.Node) (int64, error) {
	url := fmt.Sprintf("%s://%s/ping", commonBase.If(d.Conf.Server.Mode == constant.Dev, "http", "https"), node.CallbackURL)

	// Todo 工作量证明参数

	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return 0, err
	}

	start := time.Now()
	resp, err := d.httpClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			d.Log.Errorf("failed to close body: %v", err)
		}
	}(resp.Body)
	end := time.Now()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to ping node[%s]: %s", node.Key, node.CallbackURL)
	}
	return end.Sub(start).Milliseconds(), nil
}

func (d *NodeDomain) Pow(node *model.Node) (int64, error) {
	url := fmt.Sprintf("%s://%s/pow", commonBase.If(d.Conf.Server.Mode == constant.Dev, "http", "https"), node.CallbackURL)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return 0, err
	}

	start := time.Now()
	resp, err := d.httpClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			d.Log.Errorf("failed to close body: %v", err)
		}
	}(resp.Body)
	end := time.Now()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to ping node[%s]: %s", node.Key, node.CallbackURL)
	}
	return end.Sub(start).Milliseconds(), nil
}

func (d *NodeDomain) Register(ctx context.Context) error {
	n, ok := util.GetContextValue[*model.Node](ctx, constant.CtxNodeInfo)
	if !ok {
		return cv1.ErrorUnauthorized("node is not allow")
	}
	err := d.nodeRepo.Register(ctx, n)
	if err != nil {
		return err
	}
	marshal, err := json.Marshal(n)
	if err != nil {
		return err
	}
	err = d.producer.EnqueueTasks([]*model.Task{
		{
			TaskName: task.TaskNodePing.String(),
			Interval: 10 * time.Second,
			MaxRetry: 3,
			Data:     marshal,
		}, {
			TaskName: task.TaskNodePow.String(),
			Interval: 30 * time.Second,
			MaxRetry: 3,
			Data:     marshal,
		},
	})
	if err != nil {
		return err
	}

	return nil
}

func (d *NodeDomain) Unregister(ctx context.Context, key string) error {
	if key == "" {
		n, ok := util.GetContextValue[*model.Node](ctx, constant.CtxNodeInfo)
		if !ok {
			return cv1.ErrorUnauthorized("node is not allow")
		}
		key = n.Key
	}
	err := d.nodeRepo.Unregister(ctx, key)
	if err != nil {
		return err
	}
	return nil
}
