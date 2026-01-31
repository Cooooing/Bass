package cache

import (
	"common/pkg/constant"
	commonBase "common/pkg/cutil/base"
	"context"
	"encoding/json"
	"errors"
	"math"
	"signal/internal/biz/cache"
	"signal/internal/biz/model"
	"signal/internal/data/base"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type NodeCache struct {
	*base.BaseData
}

func NewNodeCache(baseData *base.BaseData) cache.NodeCache {
	return &NodeCache{
		BaseData: baseData,
	}
}

func (r *NodeCache) SetNode(ctx context.Context, n *model.Node) error {
	marshal, err := json.Marshal(n)
	if err != nil {
		return err
	}
	_, err = r.Redis.Client.HSet(ctx, constant.GetKeySignalNode(n.Key), map[string]interface{}{
		constant.SignalNodeData:               string(marshal),
		constant.SignalNodeCurrentConnections: 0,
		constant.SignalNodePingMs:             math.MaxInt64,
		constant.SignalNodePowCostMs:          math.MaxInt64,
		constant.SignalNodeLastPingTime:       time.Now().UnixMilli(),
	}).Result()
	if err != nil {
		return err
	}
	return nil
}

func (r *NodeCache) DelNode(ctx context.Context, key string) {
	r.Redis.Client.Del(ctx, constant.GetKeySignalNode(key))
}

func (r *NodeCache) UpdateNodeConnections(ctx context.Context, key string, connections int64) error {
	redisKey := constant.GetKeySignalNode(key)
	_, err := r.Redis.Client.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.HSet(ctx, redisKey, constant.SignalNodeCurrentConnections, connections)
		pipe.HSet(ctx, redisKey, constant.SignalNodeLastPingTime, time.Now().UnixMilli())
		return nil
	})
	return err
}

func (r *NodeCache) UpdateNodeConnectionsDelta(ctx context.Context, key string, delta int64) error {
	redisKey := constant.GetKeySignalNode(key)
	_, err := r.Redis.Client.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.HIncrBy(ctx, redisKey, constant.SignalNodeCurrentConnections, delta)
		pipe.HSet(ctx, redisKey, constant.SignalNodeLastPingTime, time.Now().UnixMilli())
		return nil
	})
	return err
}

func (r *NodeCache) UpdateNodePing(ctx context.Context, key string, pingMs int64) error {
	return r.Redis.Client.HSet(ctx,
		constant.GetKeySignalNode(key),
		constant.SignalNodePingMs, pingMs,
		constant.SignalNodeLastPingTime, time.Now().UnixMilli(),
	).Err()
}

func (r *NodeCache) UpdateNodePowCost(ctx context.Context, key string, powCostMs int64) error {
	return r.Redis.Client.HSet(ctx,
		constant.GetKeySignalNode(key),
		constant.SignalNodePowCostMs, powCostMs,
		constant.SignalNodeLastPingTime, time.Now().UnixMilli(),
	).Err()
}

func (r *NodeCache) GetNode(ctx context.Context, key string) (*model.Node, error) {
	result, err := r.Redis.Client.HGetAll(ctx, constant.GetKeySignalNode(key)).Result()
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, redis.Nil
	}
	// 解析字段
	var node model.Node
	err = json.Unmarshal([]byte(result[constant.SignalNodeData]), &node)
	if err != nil {
		return nil, err
	}
	node.Connections, _ = strconv.ParseInt(result[constant.SignalNodeCurrentConnections], 10, 64)
	node.PingMs, _ = strconv.ParseInt(result[constant.SignalNodePingMs], 10, 64)
	node.PowCostMs, _ = strconv.ParseInt(result[constant.SignalNodePowCostMs], 10, 64)
	t, _ := strconv.ParseInt(result[constant.SignalNodeLastPingTime], 10, 64)
	node.LastPingTime = commonBase.Ptr(time.UnixMilli(t))
	return &node, nil
}

func (r *NodeCache) SetNodeRank(ctx context.Context, key string, score float64) error {
	err := r.Redis.Client.ZAdd(ctx, constant.SignalNodeRank, redis.Z{
		Score:  score,
		Member: key,
	}).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *NodeCache) DelNodeRank(ctx context.Context, key string) {
	r.Redis.Client.ZRem(ctx, constant.SignalNodeRank, key)
}

func (r *NodeCache) ExistsNodeRank(ctx context.Context, key string) (bool, error) {
	_, err := r.Redis.Client.ZScore(ctx, constant.SignalNodeRank, key).Result()
	if errors.Is(err, redis.Nil) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

func (r *NodeCache) GetOnlineNodeKeys(ctx context.Context) ([]string, error) {
	return r.Redis.Client.ZRevRange(ctx, constant.SignalNodeRank, 0, -1).Result()
}

func (r *NodeCache) UpdateScore(ctx context.Context, key string) error {
	node, err := r.GetNode(ctx, key)
	if err != nil {
		return err
	}
	score := node.CalculateScore()
	return r.SetNodeRank(ctx, key, score)
}
