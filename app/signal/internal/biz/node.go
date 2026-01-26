package biz

import (
	"bytes"
	cv1 "common/api/common/v1"
	connectorv1 "common/api/connector/v1"
	"common/pkg"
	"common/pkg/constant"
	commonBase "common/pkg/cutil/base"
	"common/pkg/cutil/base/str"
	commonModel "common/pkg/model"
	"common/pkg/util"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"signal/internal/biz/base"
	"signal/internal/biz/cache"
	"signal/internal/biz/model"
	"signal/internal/biz/repo"
	"signal/internal/biz/task"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/sony/sonyflake/v2"
)

type NodeDomain struct {
	*base.BaseDomain
	nodeRepo   repo.NodeRepo
	nodeCache  cache.NodeCache
	producer   *Producer
	httpClient *http.Client
	sf         *sonyflake.Sonyflake
}

func NewNodeDomain(baseDomain *base.BaseDomain, nodeRepo repo.NodeRepo, nodeCache cache.NodeCache, producer *Producer, httpClient *http.Client) (*NodeDomain, error) {
	sf, err := str.NewSonyflake()
	if err != nil {
		return nil, err
	}
	return &NodeDomain{
		BaseDomain: baseDomain,
		nodeRepo:   nodeRepo,
		nodeCache:  nodeCache,
		producer:   producer,
		httpClient: httpClient,
		sf:         sf,
	}, nil
}

// GenerateSecret 生成一个 32 位随机字符串
func (d *NodeDomain) GenerateSecret() string {
	return str.RandStr(d.sf, 32, true, true, true, false)
}

func (d *NodeDomain) GetByKey(ctx context.Context, key string) (*model.Node, error) {
	cacheKey := constant.GetKeySignalNode(key)

	n, err := d.nodeCache.GetNode(ctx, cacheKey)
	if err == nil {
		return n, nil
	}
	if !errors.Is(err, redis.Nil) {
		return nil, err
	}

	n, err = d.nodeRepo.GetOne(ctx, d.Db, &repo.NodeGetReq{Key: &key})
	if err != nil {
		return nil, err
	}

	err = d.nodeCache.SetNode(ctx, n)
	if err != nil {
		return nil, err
	}
	return n, nil
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
	isOnline, err := d.nodeCache.ExistsNodeRank(ctx, n.Key)
	if err != nil {
		return err
	}
	if isOnline {
		return nil
	}

	err = d.nodeCache.SetNode(ctx, n)
	if err != nil {
		return err
	}
	err = d.nodeCache.SetNodeRank(ctx, n.Key, 0)
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
	d.nodeCache.DelNodeRank(ctx, key)
	d.nodeCache.DelNode(ctx, key)
	// Todo 清理缓存
	return nil
}

func (d *NodeDomain) Negotiate(ctx context.Context) ([]*model.Node, error) {
	nodes := make([]*model.Node, 0)
	keys, err := d.nodeCache.GetOnlineNodeKeys(ctx)
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

func (d *NodeDomain) Ticket(ctx context.Context) (string, error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return "", cv1.ErrorUnauthorized("user not login")
	}
	u, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	ticket := u.String()
	err = d.Redis.Client.SetEx(ctx, constant.GetKeySignalTicket(ticket), user.ID, time.Minute).Err()
	if err != nil {
		return "", err
	}
	return ticket, nil
}

func (d *NodeDomain) Online(ctx context.Context, ticket string) (string, error) {
	n, ok := util.GetContextValue[*model.Node](ctx, constant.CtxNodeInfo)
	if !ok {
		return "", cv1.ErrorUnauthorized("node is not allow")
	}
	result, err := d.Redis.Client.Get(ctx, constant.GetKeySignalTicket(ticket)).Result()
	if errors.Is(err, redis.Nil) {
		return "", cv1.ErrorUnauthorized("ticket is invalid")
	}
	if err != nil {
		return "", err
	}
	err = d.Redis.Client.Del(ctx, constant.GetKeySignalTicket(ticket)).Err()
	if err != nil {
		return "", err
	}
	// Todo 上线操作缓存
	sessionId := uuid.New().String()
	_ = n
	err = d.Redis.Client.Set(ctx, constant.GetKeySignalSession(sessionId), result, 0).Err()
	if err != nil {
		return "", err
	}
	return sessionId, nil
}

func (d *NodeDomain) Offline(ctx context.Context, sessionId string) error {
	n, ok := util.GetContextValue[*model.Node](ctx, constant.CtxNodeInfo)
	if !ok {
		return cv1.ErrorUnauthorized("node is not allow")
	}

	// Todo 下线操作缓存
	_ = n
	_, err := d.Redis.Client.Get(ctx, constant.GetKeySignalSession(sessionId)).Result()
	if errors.Is(err, redis.Nil) {
		return nil
	}
	if err != nil {
		return err
	}
	err = d.Redis.Client.Del(ctx, constant.GetKeySignalSession(sessionId)).Err()
	if err != nil {
		return err
	}
	return nil
}
