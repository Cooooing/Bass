package repo

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"common/pkg/client"
	"push_hub/internal/biz/model"

	"github.com/redis/go-redis/v9"
)

const (
	// Redis Key 前缀
	keyNodeHash    = "push_hub:nodes"        // Hash: node_id -> JSON{node_info}
	keyUserNodes   = "push_hub:user_nodes"    // Hash: user_id -> JSON[node_id 列表]
	keyOfflineList = "push_hub:offline:"      // List prefix: push_hub:offline:{user_id}
	keyOnlineSet   = "push_hub:online_nodes"  // Set: 所有在线节点 ID

	// 节点过期时间（无心跳超过 90s 判定离线）
	nodeExpire = 90 * time.Second
	// 离线事件过期时间
	offlineEventExpire = 24 * time.Hour
)

// NodeRegistryRepo Redis 实现的节点注册表。
type NodeRegistryRepo struct {
	rdb *client.RedisClient
}

func NewNodeRegistryRepo(rdb *client.RedisClient) *NodeRegistryRepo {
	return &NodeRegistryRepo{rdb: rdb}
}

// nodeRecord Redis 中存储的节点 JSON 结构。
type nodeRecord struct {
	NodeID          string `json:"node_id"`
	Address         string `json:"address"`
	ConnectionCount int64  `json:"connection_count"`
	Status          int32  `json:"status"`
	RegisteredAt    string `json:"registered_at"`
	LastHeartbeatAt string `json:"last_heartbeat_at"`
}

func (r *NodeRegistryRepo) RegisterNode(ctx context.Context, nodeID, address string) error {
	now := time.Now().Format(time.RFC3339)
	record := nodeRecord{
		NodeID:          nodeID,
		Address:         address,
		Status:          1, // NODE_STATUS_ONLINE
		RegisteredAt:    now,
		LastHeartbeatAt: now,
	}

	data, err := json.Marshal(record)
	if err != nil {
		return fmt.Errorf("序列化节点: %w", err)
	}

	pipe := r.rdb.Client.Pipeline()
	pipe.HSet(ctx, keyNodeHash, nodeID, data)
	pipe.SAdd(ctx, keyOnlineSet, nodeID)
	_, err = pipe.Exec(ctx)
	return err
}

func (r *NodeRegistryRepo) UpdateHeartbeat(ctx context.Context, nodeID string, connectionCount int64) error {
	// 获取当前节点信息
	data, err := r.rdb.Client.HGet(ctx, keyNodeHash, nodeID).Bytes()
	if err != nil {
		return fmt.Errorf("获取节点信息: %w", err)
	}

	var record nodeRecord
	if err := json.Unmarshal(data, &record); err != nil {
		return fmt.Errorf("反序列化节点: %w", err)
	}

	record.LastHeartbeatAt = time.Now().Format(time.RFC3339)
	record.ConnectionCount = connectionCount
	record.Status = 1 // NODE_STATUS_ONLINE

	updatedData, err := json.Marshal(record)
	if err != nil {
		return fmt.Errorf("序列化节点: %w", err)
	}

	pipe := r.rdb.Client.Pipeline()
	pipe.HSet(ctx, keyNodeHash, nodeID, updatedData)
	pipe.SAdd(ctx, keyOnlineSet, nodeID)
	_, err = pipe.Exec(ctx)
	return err
}

func (r *NodeRegistryRepo) GetNode(ctx context.Context, nodeID string) (*model.NodeInfo, error) {
	data, err := r.rdb.Client.HGet(ctx, keyNodeHash, nodeID).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, fmt.Errorf("获取节点: %w", err)
	}

	var record nodeRecord
	if err := json.Unmarshal(data, &record); err != nil {
		return nil, fmt.Errorf("反序列化节点: %w", err)
	}

	return r.toModel(&record), nil
}

func (r *NodeRegistryRepo) ListNodes(ctx context.Context) ([]*model.NodeInfo, error) {
	results, err := r.rdb.Client.HGetAll(ctx, keyNodeHash).Result()
	if err != nil {
		return nil, fmt.Errorf("列出节点: %w", err)
	}

	nodes := make([]*model.NodeInfo, 0, len(results))
	now := time.Now()
	for _, data := range results {
		var record nodeRecord
		if err := json.Unmarshal([]byte(data), &record); err != nil {
			continue
		}
		info := r.toModel(&record)
		// 判断节点是否在线（心跳超时则标记离线）
		if now.Sub(info.LastHeartbeatAt) > nodeExpire {
			info.Status = 2 // NODE_STATUS_OFFLINE
		}
		nodes = append(nodes, info)
	}
	return nodes, nil
}

func (r *NodeRegistryRepo) RemoveNode(ctx context.Context, nodeID string) error {
	pipe := r.rdb.Client.Pipeline()
	pipe.HDel(ctx, keyNodeHash, nodeID)
	pipe.SRem(ctx, keyOnlineSet, nodeID)
	_, err := pipe.Exec(ctx)
	return err
}

func (r *NodeRegistryRepo) MapUserToNode(ctx context.Context, userID int64, nodeID string) error {
	key := strconv.FormatInt(userID, 10)
	data, err := r.rdb.Client.HGet(ctx, keyUserNodes, key).Bytes()
	if err != nil && err != redis.Nil {
		return fmt.Errorf("获取用户节点映射: %w", err)
	}

	var nodeIDs []string
	if err == nil {
		_ = json.Unmarshal(data, &nodeIDs)
	}

	// 去重
	for _, id := range nodeIDs {
		if id == nodeID {
			return nil
		}
	}
	nodeIDs = append(nodeIDs, nodeID)

	updatedData, err := json.Marshal(nodeIDs)
	if err != nil {
		return fmt.Errorf("序列化用户节点映射: %w", err)
	}

	return r.rdb.Client.HSet(ctx, keyUserNodes, key, updatedData).Err()
}

func (r *NodeRegistryRepo) UnmapUserFromNode(ctx context.Context, userID int64, nodeID string) error {
	key := strconv.FormatInt(userID, 10)
	data, err := r.rdb.Client.HGet(ctx, keyUserNodes, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil
		}
		return fmt.Errorf("获取用户节点映射: %w", err)
	}

	var nodeIDs []string
	_ = json.Unmarshal(data, &nodeIDs)

	filtered := make([]string, 0, len(nodeIDs))
	for _, id := range nodeIDs {
		if id != nodeID {
			filtered = append(filtered, id)
		}
	}

	updatedData, err := json.Marshal(filtered)
	if err != nil {
		return fmt.Errorf("序列化用户节点映射: %w", err)
	}

	return r.rdb.Client.HSet(ctx, keyUserNodes, key, updatedData).Err()
}

func (r *NodeRegistryRepo) GetUserNodes(ctx context.Context, userID int64) ([]string, error) {
	key := strconv.FormatInt(userID, 10)
	data, err := r.rdb.Client.HGet(ctx, keyUserNodes, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, fmt.Errorf("获取用户节点映射: %w", err)
	}

	var nodeIDs []string
	if err := json.Unmarshal(data, &nodeIDs); err != nil {
		return nil, fmt.Errorf("反序列化用户节点映射: %w", err)
	}

	// 过滤掉不在线的节点
	onlineIDs := make([]string, 0, len(nodeIDs))
	for _, id := range nodeIDs {
		exists, err := r.rdb.Client.SIsMember(ctx, keyOnlineSet, id).Result()
		if err == nil && exists {
			onlineIDs = append(onlineIDs, id)
		}
	}

	return onlineIDs, nil
}

func (r *NodeRegistryRepo) GetAllOnlineNodes(ctx context.Context) ([]*model.NodeInfo, error) {
	nodeIDs, err := r.rdb.Client.SMembers(ctx, keyOnlineSet).Result()
	if err != nil {
		return nil, fmt.Errorf("获取在线节点集合: %w", err)
	}

	nodes := make([]*model.NodeInfo, 0, len(nodeIDs))
	now := time.Now()
	for _, nodeID := range nodeIDs {
		data, err := r.rdb.Client.HGet(ctx, keyNodeHash, nodeID).Bytes()
		if err != nil {
			continue
		}
		var record nodeRecord
		if err := json.Unmarshal(data, &record); err != nil {
			continue
		}
		info := r.toModel(&record)
		// 心跳超时的节点不返回
		if now.Sub(info.LastHeartbeatAt) > nodeExpire {
			// 异步清理过期节点
			_ = r.rdb.Client.SRem(ctx, keyOnlineSet, nodeID).Err()
			continue
		}
		nodes = append(nodes, info)
	}
	return nodes, nil
}

func (r *NodeRegistryRepo) SaveOfflineEvent(ctx context.Context, userID int64, event *model.PushEvent) error {
	key := keyOfflineList + strconv.FormatInt(userID, 10)
	data, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("序列化离线事件: %w", err)
	}

	pipe := r.rdb.Client.Pipeline()
	pipe.RPush(ctx, key, data)
	pipe.Expire(ctx, key, offlineEventExpire)
	_, err = pipe.Exec(ctx)
	return err
}

func (r *NodeRegistryRepo) GetOfflineEvents(ctx context.Context, userID int64) ([]*model.PushEvent, error) {
	key := keyOfflineList + strconv.FormatInt(userID, 10)
	results, err := r.rdb.Client.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		return nil, fmt.Errorf("获取离线事件: %w", err)
	}

	events := make([]*model.PushEvent, 0, len(results))
	for _, data := range results {
		var event model.PushEvent
		if err := json.Unmarshal([]byte(data), &event); err != nil {
			continue
		}
		events = append(events, &event)
	}
	return events, nil
}

func (r *NodeRegistryRepo) ClearOfflineEvents(ctx context.Context, userID int64) error {
	key := keyOfflineList + strconv.FormatInt(userID, 10)
	return r.rdb.Client.Del(ctx, key).Err()
}

func (r *NodeRegistryRepo) toModel(record *nodeRecord) *model.NodeInfo {
	info := &model.NodeInfo{
		NodeID:          record.NodeID,
		Address:         record.Address,
		ConnectionCount: record.ConnectionCount,
		Status:          record.Status,
	}
	info.RegisteredAt, _ = time.Parse(time.RFC3339, record.RegisteredAt)
	info.LastHeartbeatAt, _ = time.Parse(time.RFC3339, record.LastHeartbeatAt)
	return info
}
