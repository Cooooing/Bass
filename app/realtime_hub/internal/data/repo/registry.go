package repo

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"common/pkg/client"
	"realtime_hub/internal/biz/model"
	bizrepo "realtime_hub/internal/biz/repo"
	"realtime_hub/internal/enum"

	"github.com/redis/go-redis/v9"
)

type Registry struct {
	redis *client.RedisClient
}

const connectionLease = 30 * time.Second

func NewRegistry(redis *client.RedisClient) bizrepo.Registry {
	return &Registry{redis: redis}
}

func (r *Registry) RegisterNode(ctx context.Context, nodeID string, group enum.NodeGroup, realms []string) error {
	pipe := r.redis.Client.TxPipeline()
	for _, realm := range realms {
		pipe.HSet(ctx, r.nodeKey(realm, nodeID), "group", string(group), "registered_at_ms", time.Now().UnixMilli())
		pipe.PExpire(ctx, r.nodeKey(realm, nodeID), 30*time.Second)
	}
	_, err := pipe.Exec(ctx)
	return err
}

func (r *Registry) RenewNode(ctx context.Context, nodeID string, group enum.NodeGroup, realms []string) (bool, error) {
	for _, realm := range realms {
		valid, err := r.redis.Client.Eval(ctx, `
if redis.call('EXISTS', KEYS[1]) == 0 then
  return 0
end
if redis.call('HGET', KEYS[1], 'has_connections') == '1' and redis.call('EXISTS', KEYS[2]) == 0 then
  return 0
end
redis.call('HSET', KEYS[1], 'group', ARGV[1])
redis.call('PEXPIRE', KEYS[1], 30000)
if redis.call('EXISTS', KEYS[2]) == 1 then
  redis.call('PEXPIRE', KEYS[2], 30000)
end
return 1`, []string{r.nodeKey(realm, nodeID), r.nodeConnectionsKey(realm, nodeID)}, string(group)).Int()
		if err != nil {
			return false, err
		}
		if valid == 0 {
			return false, nil
		}
	}
	return true, nil
}

func (r *Registry) RemoveNode(ctx context.Context, nodeID string, realms []string) error {
	for _, realm := range realms {
		if _, err := r.redis.Client.Eval(ctx, `
local connections = redis.call('HGETALL', KEYS[2])
for index = 1, #connections, 2 do
  local connectionID = connections[index]
  local userID = connections[index + 1]
  local connectionsKey = ARGV[1] .. '{' .. userID .. '}:connections'
  local nodesKey = ARGV[1] .. '{' .. userID .. '}:connection_nodes'
  if redis.call('HGET', nodesKey, connectionID) == ARGV[2] then
    redis.call('ZREM', connectionsKey, connectionID)
    redis.call('HDEL', nodesKey, connectionID)
    if redis.call('ZCARD', connectionsKey) == 0 then
      redis.call('ZREM', KEYS[3], userID)
    end
  end
end

redis.call('DEL', KEYS[1], KEYS[2])
		return #connections / 2`, []string{r.nodeKey(realm, nodeID), r.nodeConnectionsKey(realm, nodeID), r.onlineUsersKey(realm)}, r.nodePrefix(realm), nodeID).Result(); err != nil {
			return fmt.Errorf("remove node: %w", err)
		}
	}
	return nil
}

func (r *Registry) BindConnection(ctx context.Context, nodeID, realm string, userID int64, connectionID string) ([]model.EvictedConnection, error) {
	result, err := r.redis.Client.Eval(ctx, `
if redis.call('EXISTS', KEYS[1]) == 0 then
  return false
end
local existing = redis.call('HGETALL', KEYS[3])
for index = 1, #existing, 2 do
  local existingConnectionID = existing[index]
  local existingNodeID = existing[index + 1]
  if redis.call('EXISTS', ARGV[6] .. '{' .. existingNodeID .. '}:meta') == 0 then
    redis.call('ZREM', KEYS[2], existingConnectionID)
    redis.call('HDEL', KEYS[3], existingConnectionID)
    redis.call('HDEL', ARGV[5] .. '{' .. existingNodeID .. '}:connections', existingConnectionID)
    if redis.call('ZCARD', KEYS[2]) == 0 then
      redis.call('ZREM', KEYS[5], ARGV[4])
    end
  end
end
redis.call('ZADD', KEYS[2], 'NX', ARGV[1], ARGV[2])
redis.call('HSET', KEYS[3], ARGV[2], ARGV[3])
redis.call('HSET', KEYS[4], ARGV[2], ARGV[4])
redis.call('ZADD', KEYS[5], ARGV[1], ARGV[4])
redis.call('HSET', KEYS[1], 'has_connections', '1')
redis.call('PEXPIRE', KEYS[4], 30000)
local evicted = {}
while redis.call('ZCARD', KEYS[2]) > 5 do
  local oldest = redis.call('ZRANGE', KEYS[2], 0, 0)[1]
  local oldNodeID = redis.call('HGET', KEYS[3], oldest)
  redis.call('ZREM', KEYS[2], oldest)
  redis.call('HDEL', KEYS[3], oldest)
  if redis.call('ZCARD', KEYS[2]) == 0 then
    redis.call('ZREM', KEYS[5], ARGV[4])
  end
  if oldNodeID then
    redis.call('HDEL', ARGV[5] .. '{' .. oldNodeID .. '}:connections', oldest)
    table.insert(evicted, oldNodeID)
    table.insert(evicted, oldest)
  end
end
return evicted`, []string{r.nodeKey(realm, nodeID), r.userConnectionsKey(realm, userID), r.userConnectionNodesKey(realm, userID), r.nodeConnectionsKey(realm, nodeID), r.onlineUsersKey(realm)}, strconv.FormatInt(time.Now().UnixMilli(), 10), connectionID, nodeID, strconv.FormatInt(userID, 10), r.nodePrefix(realm), r.nodePrefix(realm)).Result()
	if err != nil {
		return nil, err
	}
	values, ok := result.([]any)
	if !ok {
		return nil, fmt.Errorf("node is not registered")
	}
	evicted := make([]model.EvictedConnection, 0, len(values)/2)
	for index := 0; index+1 < len(values); index += 2 {
		nodeID, nodeOK := values[index].(string)
		connectionID, connectionOK := values[index+1].(string)
		if nodeOK && connectionOK {
			evicted = append(evicted, model.EvictedConnection{NodeID: nodeID, ConnectionID: connectionID})
		}
	}
	return evicted, nil
}

func (r *Registry) UnbindConnection(ctx context.Context, nodeID, realm string, userID int64, connectionID string) error {
	_, err := r.redis.Client.Eval(ctx, `
if redis.call('HGET', KEYS[3], ARGV[1]) == ARGV[2] then
  redis.call('ZREM', KEYS[1], ARGV[1])
  redis.call('HDEL', KEYS[2], ARGV[1])
end
if redis.call('HGET', KEYS[4], ARGV[1]) == ARGV[3] then
  redis.call('HDEL', KEYS[4], ARGV[1])
end
if redis.call('ZCARD', KEYS[1]) == 0 then
  redis.call('ZREM', KEYS[6], ARGV[4])
end
if redis.call('HLEN', KEYS[4]) == 0 then
  redis.call('HDEL', KEYS[5], 'has_connections')
end
return 1`, []string{r.userConnectionsKey(realm, userID), r.userConnectionNodesKey(realm, userID), r.userConnectionNodesKey(realm, userID), r.nodeConnectionsKey(realm, nodeID), r.nodeKey(realm, nodeID), r.onlineUsersKey(realm)}, connectionID, nodeID, strconv.FormatInt(userID, 10)).Result()
	return err
}

func (r *Registry) RenewConnection(ctx context.Context, nodeID, realm string, userID int64, connectionID string) error {
	_, err := r.redis.Client.Eval(ctx, `
if redis.call('EXISTS', KEYS[1]) == 0 then
  return 0
end
if not redis.call('ZSCORE', KEYS[2], ARGV[1]) then
  return 0
end
if redis.call('HGET', KEYS[3], ARGV[1]) ~= ARGV[2] then
  return 0
end
if redis.call('HGET', KEYS[4], ARGV[1]) ~= ARGV[3] then
  return 0
end
redis.call('ZADD', KEYS[2], 'XX', ARGV[4], ARGV[1])
redis.call('ZADD', KEYS[5], ARGV[4], ARGV[3])
return 1`, []string{r.nodeKey(realm, nodeID), r.userConnectionsKey(realm, userID), r.userConnectionNodesKey(realm, userID), r.nodeConnectionsKey(realm, nodeID), r.onlineUsersKey(realm)}, connectionID, nodeID, strconv.FormatInt(userID, 10), strconv.FormatInt(time.Now().UnixMilli(), 10)).Int()
	return err
}

func (r *Registry) UserNodes(ctx context.Context, realm string, userID int64) ([]string, error) {
	values, err := r.redis.Client.Eval(ctx, `
local nodes = {}
local connections = redis.call('HGETALL', KEYS[2])
for index = 1, #connections, 2 do
  local connectionID = connections[index]
  local nodeID = connections[index + 1]
  if redis.call('EXISTS', ARGV[1] .. '{' .. nodeID .. '}:meta') == 1 then
    table.insert(nodes, nodeID)
  else
    redis.call('ZREM', KEYS[1], connectionID)
    redis.call('HDEL', KEYS[2], connectionID)
    redis.call('HDEL', ARGV[1] .. '{' .. nodeID .. '}:connections', connectionID)
  end
end
return nodes`, []string{r.userConnectionsKey(realm, userID), r.userConnectionNodesKey(realm, userID)}, r.nodePrefix(realm)).Result()
	if err != nil {
		return nil, err
	}
	connectionNodes, ok := values.([]any)
	if !ok {
		return nil, fmt.Errorf("user node routes are invalid")
	}
	nodes := make(map[string]struct{}, len(connectionNodes))
	for _, value := range connectionNodes {
		nodeID, ok := value.(string)
		if ok {
			nodes[nodeID] = struct{}{}
		}
	}
	result := make([]string, 0, len(nodes))
	for nodeID := range nodes {
		result = append(result, nodeID)
	}
	return result, nil
}

func (r *Registry) QueryOnline(ctx context.Context, realm string, userIDs []int64) (map[int64]bool, error) {
	if len(userIDs) == 0 {
		return map[int64]bool{}, nil
	}
	statusesByUserID := make(map[int64]bool, len(userIDs))
	for _, userID := range userIDs {
		connections, err := r.activeConnectionCount(ctx, realm, userID)
		if err != nil {
			return nil, err
		}
		statusesByUserID[userID] = connections > 0
	}
	return statusesByUserID, nil
}

func (r *Registry) ListOnlineUsers(ctx context.Context, realm string) ([]int64, error) {
	values, err := r.redis.Client.ZRevRange(ctx, r.onlineUsersKey(realm), 0, -1).Result()
	if err != nil {
		return nil, err
	}
	result := make([]int64, 0, len(values))
	for _, value := range values {
		userID, parseErr := strconv.ParseInt(value, 10, 64)
		if parseErr != nil || userID <= 0 {
			continue
		}
		connections, countErr := r.activeConnectionCount(ctx, realm, userID)
		if countErr != nil {
			return nil, countErr
		}
		if connections > 0 {
			result = append(result, userID)
		}
	}
	return result, nil
}

// activeConnectionCount 清理失效连接，并返回账号的有效连接数。
func (r *Registry) activeConnectionCount(ctx context.Context, realm string, userID int64) (int64, error) {
	connections, err := r.redis.Client.Eval(ctx, `
local connectionIDs = redis.call('ZRANGE', KEYS[1], 0, -1)
for _, connectionID in ipairs(connectionIDs) do
  local nodeID = redis.call('HGET', KEYS[2], connectionID)
  local connectedAt = redis.call('ZSCORE', KEYS[1], connectionID)
  if not nodeID or not connectedAt or tonumber(ARGV[3]) - tonumber(connectedAt) > tonumber(ARGV[4]) or redis.call('EXISTS', ARGV[1] .. '{' .. nodeID .. '}:meta') == 0 then
    redis.call('ZREM', KEYS[1], connectionID)
    redis.call('HDEL', KEYS[2], connectionID)
    if nodeID then
      redis.call('HDEL', ARGV[1] .. '{' .. nodeID .. '}:connections', connectionID)
    end
  end
end
local count = redis.call('ZCARD', KEYS[1])
if count == 0 then
  redis.call('ZREM', KEYS[3], ARGV[2])
end
return count`, []string{r.userConnectionsKey(realm, userID), r.userConnectionNodesKey(realm, userID), r.onlineUsersKey(realm)}, r.nodePrefix(realm), strconv.FormatInt(userID, 10), strconv.FormatInt(time.Now().UnixMilli(), 10), strconv.FormatInt(connectionLease.Milliseconds(), 10)).Int64()
	if err != nil {
		return 0, err
	}
	return connections, nil
}

func (r *Registry) IssueTicket(ctx context.Context, realm, ticket string, userID int64) error {
	return r.redis.Client.Set(ctx, r.ticketKey(realm, ticket), strconv.FormatInt(userID, 10), time.Minute).Err()
}

func (r *Registry) ConsumeTicket(ctx context.Context, realm, ticket string) (int64, error) {
	userID, err := r.redis.Client.GetDel(ctx, r.ticketKey(realm, ticket)).Result()
	if errors.Is(err, redis.Nil) {
		return 0, fmt.Errorf("connection ticket is invalid")
	}
	if err != nil {
		return 0, err
	}
	parsed, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		return 0, err
	}
	return parsed, nil
}

func (*Registry) nodePrefix(realm string) string {
	return "realtime:{" + realm + "}:"
}

func (r *Registry) nodeKey(realm, nodeID string) string {
	return r.nodePrefix(realm) + "{" + nodeID + "}:meta"
}

func (r *Registry) nodeConnectionsKey(realm, nodeID string) string {
	return r.nodePrefix(realm) + "{" + nodeID + "}:connections"
}

func (r *Registry) userConnectionsKey(realm string, userID int64) string {
	return r.nodePrefix(realm) + "{" + strconv.FormatInt(userID, 10) + "}:connections"
}

func (r *Registry) userConnectionNodesKey(realm string, userID int64) string {
	return r.nodePrefix(realm) + "{" + strconv.FormatInt(userID, 10) + "}:connection_nodes"
}

func (r *Registry) onlineUsersKey(realm string) string {
	return r.nodePrefix(realm) + "online_users"
}

func (*Registry) ticketKey(realm, ticket string) string {
	return "realtime:{" + realm + "}:{" + ticket + "}"
}
