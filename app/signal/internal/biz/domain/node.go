package domain

import (
	"bytes"
	cv1 "common/api/gen/common/v1"
	connectorv1 "common/api/gen/connector/v1"
	"common/pkg/client"
	"common/pkg/constant"
	commonModel "common/pkg/model"
	"common/pkg/util"
	"common/pkg/util/server"
	"common/pkg/util/str"
	"common/pkg/util/task"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	domainbase "signal/internal/biz/base"
	"signal/internal/biz/cache"
	"signal/internal/biz/model"
	"signal/internal/biz/repo"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/sony/sonyflake/v2"
)

type NodeDomain struct {
	*domainbase.BaseDomain
	nodeRepo     repo.NodeRepo
	nodeCache    cache.NodeCache
	sessionCache cache.SessionCache
	asynqCache   *task.AsynqCache
	producer     *client.Producer
	httpClient   *http.Client
	sf           *sonyflake.Sonyflake
}

func NewNodeDomain(baseDomain *domainbase.BaseDomain, nodeRepo repo.NodeRepo, nodeCache cache.NodeCache, sessionCache cache.SessionCache, asynqCache *task.AsynqCache, producer *client.Producer, httpClient *http.Client) (*NodeDomain, error) {
	sf, err := str.NewSonyflake()
	if err != nil {
		return nil, err
	}
	return &NodeDomain{
		BaseDomain:   baseDomain,
		nodeRepo:     nodeRepo,
		nodeCache:    nodeCache,
		sessionCache: sessionCache,
		asynqCache:   asynqCache,
		producer:     producer,
		httpClient:   httpClient,
		sf:           sf,
	}, nil
}

// GenerateSecret 生成一个 32 位随机字符串
func (d *NodeDomain) GenerateSecret() string {
	return str.RandStr(d.sf, 32, true, true, true, false)
}

func (d *NodeDomain) GetByKey(ctx context.Context, key string) (*model.Node, error) {
	n, err := d.nodeCache.GetNode(ctx, key)
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
	url := fmt.Sprintf("%s://%s/ping", util.If(d.Conf.Server.Mode == constant.Dev, "http", "https"), node.CallbackURL)

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
	url := fmt.Sprintf("%s://%s/pow", util.If(d.Conf.Server.Mode == constant.Dev, "http", "https"), node.CallbackURL)

	// 生成工作量证明参数
	challenge := str.RandStr(d.sf, 32, true, true, true, false)
	var difficulty int32 = 5
	param, err := json.Marshal(&connectorv1.PowRequest{
		Challenge:  challenge,
		Difficulty: difficulty,
	})
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(param))
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

	data := &server.Result[*connectorv1.PowReply]{}
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

func (d *NodeDomain) Session(node *model.Node) ([]string, error) {
	sessionIds := make([]string, 0)
	url := fmt.Sprintf("%s://%s/session", util.If(d.Conf.Server.Mode == constant.Dev, "http", "https"), node.CallbackURL)

	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return sessionIds, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := d.httpClient.Do(req)
	if err != nil {
		return sessionIds, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			d.Log.Errorf("failed to close body: %v", err)
		}
	}(resp.Body)

	data := &server.Result[*connectorv1.SessionReply]{}
	err = json.NewDecoder(resp.Body).Decode(data)
	if err != nil {
		return sessionIds, err
	}

	if resp.StatusCode != http.StatusOK {
		return sessionIds, fmt.Errorf("failed to session node[%s]: %s", node.Key, node.CallbackURL)
	}

	return data.Data.SessionIds, nil
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

	version := time.Now().UnixMilli()
	pingTaskName := fmt.Sprintf("%s-%s", constant.TaskSignalNodePing.String(), n.Key)
	powTaskName := fmt.Sprintf("%s-%s", constant.TaskSignalNodePow.String(), n.Key)
	sessionTaskName := fmt.Sprintf("%s-%s", constant.TaskSignalNodeSession.String(), n.Key)
	err = d.asynqCache.SetAsynqTaskVersion(ctx, pingTaskName, version, 20*time.Second)
	if err != nil {
		return err
	}
	err = d.asynqCache.SetAsynqTaskVersion(ctx, powTaskName, version, 60*time.Second)
	if err != nil {
		return err
	}
	err = d.asynqCache.SetAsynqTaskVersion(ctx, sessionTaskName, version, 120*time.Second)
	if err != nil {
		return err
	}
	err = d.producer.EnqueueContextTasks(ctx, []*commonModel.Task{
		{
			TaskName: pingTaskName,
			Version:  version,
			Interval: 10 * time.Second,
			Data:     marshal,
		}, {
			TaskName: powTaskName,
			Version:  version,
			Interval: 30 * time.Second,
			Data:     marshal,
		}, {
			TaskName: sessionTaskName,
			Version:  version,
			Interval: 60 * time.Second,
			Data:     marshal,
		},
	})
	if err != nil {
		return err
	}
	d.Log.Infof("node[%s] register", n.Key)
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
	// 尽量保证幂等
	exists, _ := d.nodeCache.ExistsNodeRank(ctx, key)
	if !exists {
		return nil
	}

	d.nodeCache.DelNodeRank(ctx, key)
	d.nodeCache.DelNode(ctx, key)
	d.Log.Infof("node[%s] unregister", key)
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
		if n, exist := nodeMap[key]; exist {
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
	ticket := uuid.New().String()
	err := d.sessionCache.SetTicket(ctx, ticket, user.ID)
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

	userId, err := d.sessionCache.GetTicket(ctx, ticket)
	if err != nil {
		return "", err
	}
	sessionId := uuid.New().String()
	err = d.sessionCache.SetSession(ctx, sessionId, userId, n.Key)
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

	return nil
}
