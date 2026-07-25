package repo

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"common/pkg/client"
	"push_hub/internal/biz/model"
	bizrepo "push_hub/internal/biz/repo"

	"github.com/redis/go-redis/v9"
)

// NodeRegistryRepo Redis 实现的节点注册表。
type NodeRegistryRepo struct {
	rdb                *client.RedisClient
	keyNodeHash        string
	keyUserNodes       string
	keyOfflineList     string
	keyOnlineSet       string
	nodeExpire         time.Duration
	offlineEventExpire time.Duration
}

func NewNodeRegistryRepo(
	rdb *client.RedisClient,
) *NodeRegistryRepo {
	return &NodeRegistryRepo{
		rdb:                rdb,
		keyNodeHash:        "push_hub:nodes",
		keyUserNodes:       "push_hub:user_nodes",
		keyOfflineList:     "push_hub:offline:",
		keyOnlineSet:       "push_hub:online_nodes",
		nodeExpire:         90 * time.Second,
		offlineEventExpire: 24 * time.Hour,
	}
}

func (r *NodeRegistryRepo) RegisterNode(ctx context.Context, req *bizrepo.RegisterNodeReq) error {
	return r.registerNode(ctx, req.NodeID, req.Address)
}

func (r *NodeRegistryRepo) UpdateHeartbeat(ctx context.Context, req *bizrepo.UpdateHeartbeatReq) error {
	return r.updateHeartbeat(ctx, req.NodeID, req.ConnectionCount)
}

func (r *NodeRegistryRepo) GetNode(ctx context.Context, nodeID string) (*model.NodeInfo, error) {
	return r.getNode(ctx, nodeID)
}

func (r *NodeRegistryRepo) ListNodes(ctx context.Context) ([]*model.NodeInfo, error) {
	return r.listNodes(ctx)
}

func (r *NodeRegistryRepo) RemoveNode(ctx context.Context, nodeID string) error {
	return r.removeNode(ctx, nodeID)
}

func (r *NodeRegistryRepo) MapUserToNode(ctx context.Context, req *bizrepo.MapUserToNodeReq) error {
	return r.mapUserToNode(ctx, req.UserID, req.NodeID)
}

func (r *NodeRegistryRepo) UnmapUserFromNode(ctx context.Context, req *bizrepo.UnmapUserFromNodeReq) error {
	return r.unmapUserFromNode(ctx, req.UserID, req.NodeID)
}

func (r *NodeRegistryRepo) GetUserNodes(ctx context.Context, userID int64) ([]string, error) {
	return r.getUserNodes(ctx, userID)
}

func (r *NodeRegistryRepo) GetAllOnlineNodes(ctx context.Context) ([]*model.NodeInfo, error) {
	return r.getAllOnlineNodes(ctx)
}

func (r *NodeRegistryRepo) SaveOfflineEvent(ctx context.Context, req *bizrepo.SaveOfflineEventReq) error {
	return r.saveOfflineEvent(ctx, req.UserID, req.Event)
}

func (r *NodeRegistryRepo) GetOfflineEvents(ctx context.Context, userID int64) ([]*model.PushEvent, error) {
	return r.getOfflineEvents(ctx, userID)
}

func (r *NodeRegistryRepo) ClearOfflineEvents(ctx context.Context, userID int64) error {
	return r.clearOfflineEvents(ctx, userID)
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

func (r *NodeRegistryRepo) registerNode(ctx context.Context, nodeID, address string) error {
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
	pipe.HSet(ctx, r.keyNodeHash, nodeID, data)
	pipe.SAdd(ctx, r.keyOnlineSet, nodeID)
	_, err = pipe.Exec(ctx)
	return err
}

func (r *NodeRegistryRepo) updateHeartbeat(ctx context.Context, nodeID string, connectionCount int64) error {
	// 获取当前节点信息
	data, err := r.rdb.Client.HGet(ctx, r.keyNodeHash, nodeID).Bytes()
	if err != nil {
		return fmt.Errorf("获取节点信息: %w", err)
	}

	record := new(nodeRecord)
	if err := json.Unmarshal(data, record); err != nil {
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
	pipe.HSet(ctx, r.keyNodeHash, nodeID, updatedData)
	pipe.SAdd(ctx, r.keyOnlineSet, nodeID)
	_, err = pipe.Exec(ctx)
	return err
}

func (r *NodeRegistryRepo) getNode(ctx context.Context, nodeID string) (*model.NodeInfo, error) {
	data, err := r.rdb.Client.HGet(ctx, r.keyNodeHash, nodeID).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, fmt.Errorf("获取节点: %w", err)
	}

	record := new(nodeRecord)
	if err := json.Unmarshal(data, record); err != nil {
		return nil, fmt.Errorf("反序列化节点: %w", err)
	}

	return r.toModel(record), nil
}

func (r *NodeRegistryRepo) listNodes(ctx context.Context) ([]*model.NodeInfo, error) {
	results, err := r.rdb.Client.HGetAll(ctx, r.keyNodeHash).Result()
	if err != nil {
		return nil, fmt.Errorf("列出节点: %w", err)
	}

	nodes := make([]*model.NodeInfo, 0, len(results))
	now := time.Now()
	for _, data := range results {
		record := new(nodeRecord)
		if err := json.Unmarshal([]byte(data), record); err != nil {
			continue
		}
		info := r.toModel(record)
		// 判断节点是否在线（心跳超时则标记离线）
		if now.Sub(info.LastHeartbeatAt) > r.nodeExpire {
			info.Status = 2 // NODE_STATUS_OFFLINE
		}
		nodes = append(nodes, info)
	}
	return nodes, nil
}

func (r *NodeRegistryRepo) removeNode(ctx context.Context, nodeID string) error {
	pipe := r.rdb.Client.Pipeline()
	pipe.HDel(ctx, r.keyNodeHash, nodeID)
	pipe.SRem(ctx, r.keyOnlineSet, nodeID)
	_, err := pipe.Exec(ctx)
	return err
}

func (r *NodeRegistryRepo) mapUserToNode(ctx context.Context, userID int64, nodeID string) error {
	key := strconv.FormatInt(userID, 10)
	data, err := r.rdb.Client.HGet(ctx, r.keyUserNodes, key).Bytes()
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

	return r.rdb.Client.HSet(ctx, r.keyUserNodes, key, updatedData).Err()
}

func (r *NodeRegistryRepo) unmapUserFromNode(ctx context.Context, userID int64, nodeID string) error {
	key := strconv.FormatInt(userID, 10)
	data, err := r.rdb.Client.HGet(ctx, r.keyUserNodes, key).Bytes()
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

	return r.rdb.Client.HSet(ctx, r.keyUserNodes, key, updatedData).Err()
}

func (r *NodeRegistryRepo) getUserNodes(ctx context.Context, userID int64) ([]string, error) {
	key := strconv.FormatInt(userID, 10)
	data, err := r.rdb.Client.HGet(ctx, r.keyUserNodes, key).Bytes()
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
		exists, err := r.rdb.Client.SIsMember(ctx, r.keyOnlineSet, id).Result()
		if err == nil && exists {
			onlineIDs = append(onlineIDs, id)
		}
	}

	return onlineIDs, nil
}

func (r *NodeRegistryRepo) getAllOnlineNodes(ctx context.Context) ([]*model.NodeInfo, error) {
	nodeIDs, err := r.rdb.Client.SMembers(ctx, r.keyOnlineSet).Result()
	if err != nil {
		return nil, fmt.Errorf("获取在线节点集合: %w", err)
	}

	nodes := make([]*model.NodeInfo, 0, len(nodeIDs))
	now := time.Now()
	for _, nodeID := range nodeIDs {
		data, err := r.rdb.Client.HGet(ctx, r.keyNodeHash, nodeID).Bytes()
		if err != nil {
			continue
		}
		record := new(nodeRecord)
		if err := json.Unmarshal(data, record); err != nil {
			continue
		}
		info := r.toModel(record)
		// 心跳超时的节点不返回
		if now.Sub(info.LastHeartbeatAt) > r.nodeExpire {
			// 异步清理过期节点
			_ = r.rdb.Client.SRem(ctx, r.keyOnlineSet, nodeID).Err()
			continue
		}
		nodes = append(nodes, info)
	}
	return nodes, nil
}

func (r *NodeRegistryRepo) saveOfflineEvent(ctx context.Context, userID int64, event *model.PushEvent) error {
	key := r.keyOfflineList + strconv.FormatInt(userID, 10)
	data, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("序列化离线事件: %w", err)
	}

	pipe := r.rdb.Client.Pipeline()
	pipe.RPush(ctx, key, data)
	pipe.Expire(ctx, key, r.offlineEventExpire)
	_, err = pipe.Exec(ctx)
	return err
}

func (r *NodeRegistryRepo) getOfflineEvents(ctx context.Context, userID int64) ([]*model.PushEvent, error) {
	key := r.keyOfflineList + strconv.FormatInt(userID, 10)
	results, err := r.rdb.Client.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		return nil, fmt.Errorf("获取离线事件: %w", err)
	}

	events := make([]*model.PushEvent, 0, len(results))
	for _, data := range results {
		event := new(model.PushEvent)
		if err := json.Unmarshal([]byte(data), event); err != nil {
			continue
		}
		events = append(events, event)
	}
	return events, nil
}

func (r *NodeRegistryRepo) clearOfflineEvents(ctx context.Context, userID int64) error {
	key := r.keyOfflineList + strconv.FormatInt(userID, 10)
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
