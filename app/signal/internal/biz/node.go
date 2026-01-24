package biz

import (
	"bytes"
	cv1 "common/api/common/v1"
	connectorv1 "common/api/connector/v1"
	"common/pkg"
	"common/pkg/constant"
	commonBase "common/pkg/cutil/base"
	"common/pkg/cutil/base/str"
	"common/pkg/util"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"signal/internal/biz/base"
	"signal/internal/biz/model"
	"signal/internal/biz/repo"
	"signal/internal/biz/task"
	"strings"
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

func (d *NodeDomain) Pow(node *model.Node) (int64, error) {
	url := fmt.Sprintf("%s://%s/pow", commonBase.If(d.Conf.Server.Mode == constant.Dev, "http", "https"), node.CallbackURL)

	// 生成工作量证明参数
	challenge := str.RandStr(d.sf, 32, true, true, true, false)
	var difficulty int32 = 5
	b, err := json.Marshal(&connectorv1.PowRequest{
		Challenge:  challenge,
		Difficulty: difficulty,
	})
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(b))
	if err != nil {
		return 0, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

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

	data := &pkg.Result[*connectorv1.PowResponse]{}
	err = json.NewDecoder(resp.Body).Decode(data)
	if err != nil {
		return 0, err
	}

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to Pow node[%s]: %s", node.Key, node.CallbackURL)
	}

	// 验证返回结果
	msg := challenge + ":" + data.Data.Nonce
	sum := sha256.Sum256([]byte(msg))
	h := hex.EncodeToString(sum[:])
	if ok := strings.HasPrefix(h, strings.Repeat("0", int(difficulty))); !ok || h != data.Data.HashHex {
		return 0, fmt.Errorf("PoW verify failed,node[%s]: %s", node.Key, node.CallbackURL)
	}

	// 耗时包含网络耗时
	return end.Sub(start).Milliseconds(), nil
}

func (d *NodeDomain) Register(ctx context.Context) error {
	n, ok := util.GetContextValue[*model.Node](ctx, constant.CtxNodeInfo)
	if !ok {
		return cv1.ErrorUnauthorized("node is not allow")
	}

	// 幂等处理
	isOnline, err := d.nodeRepo.IsOnline(ctx, n.Key)
	if err != nil {
		return err
	}
	if isOnline {
		return nil
	}

	err = d.nodeRepo.Register(ctx, n)
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

func (d *NodeDomain) Negotiate(ctx context.Context) ([]*model.Node, error) {
	nodes := make([]*model.Node, 0)
	keys, err := d.nodeRepo.GetOnlineNodeKeys(ctx)
	if err != nil {
		return nil, err
	}
	nodeMap, err := d.nodeRepo.GetMap(ctx, d.Db, &repo.NodeGetReq{Keys: keys})
	if err != nil {
		return nil, err
	}
	for _, key := range keys {
		if n, exist := nodeMap.Get(key); exist {
			nodes = append(nodes, n)
		}
	}
	return nodes, nil
}
