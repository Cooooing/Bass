package repo

import (
	commonclient "common/pkg/client"
	"context"
	"fmt"
	bizrepo "game_idle/internal/biz/repo"

	"github.com/redis/go-redis/v9"
)

var _ bizrepo.CharacterSessionRepo = (*CharacterSessionRepo)(nil)

type CharacterSessionRepo struct {
	redisClient          *commonclient.RedisClient
	onlineRedisKeyFormat string
	onlineScript         *redis.Script
	pingScript           *redis.Script
	offlineScript        *redis.Script
}

func NewCharacterSessionRepo(redisClient *commonclient.RedisClient) bizrepo.CharacterSessionRepo {
	return &CharacterSessionRepo{
		redisClient:          redisClient,
		onlineRedisKeyFormat: "game_idle:character:online:%d",
		onlineScript: redis.NewScript(`
local old_session_id = redis.call('GET', KEYS[1])
redis.call('SET', KEYS[1], ARGV[1], 'EX', ARGV[2])
return old_session_id or ''
`),
		pingScript: redis.NewScript(`
if redis.call('GET', KEYS[1]) ~= ARGV[1] then
  return 0
end
redis.call('EXPIRE', KEYS[1], ARGV[2])
return 1
`),
		offlineScript: redis.NewScript(`
if redis.call('GET', KEYS[1]) ~= ARGV[1] then
  return 0
end
redis.call('DEL', KEYS[1])
return 1
`),
	}
}

func (r *CharacterSessionRepo) Online(ctx context.Context, characterID int64, sessionID string, ttlSeconds int64) (string, error) {
	return r.onlineScript.Run(
		ctx,
		r.redisClient.Client,
		[]string{r.onlineRedisKey(characterID)},
		sessionID,
		ttlSeconds,
	).Text()
}

func (r *CharacterSessionRepo) Ping(ctx context.Context, characterID int64, sessionID string, ttlSeconds int64) (bool, error) {
	result, err := r.pingScript.Run(
		ctx,
		r.redisClient.Client,
		[]string{r.onlineRedisKey(characterID)},
		sessionID,
		ttlSeconds,
	).Int()
	return result == 1, err
}

func (r *CharacterSessionRepo) Offline(ctx context.Context, characterID int64, sessionID string) (bool, error) {
	result, err := r.offlineScript.Run(
		ctx,
		r.redisClient.Client,
		[]string{r.onlineRedisKey(characterID)},
		sessionID,
	).Int()
	return result == 1, err
}

func (r *CharacterSessionRepo) onlineRedisKey(characterID int64) string {
	return fmt.Sprintf(r.onlineRedisKeyFormat, characterID)
}
