package cache

import (
	cv1 "common/api/common/v1"
	"common/pkg/constant"
	"common/pkg/cutil/collections/set"
	"context"
	"errors"
	"signal/internal/biz/cache"
	"signal/internal/data/base"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type SessionCache struct {
	*base.BaseData

	sessionExpire time.Duration
}

func NewSessionCache(baseData *base.BaseData) cache.SessionCache {
	return &SessionCache{
		BaseData:      baseData,
		sessionExpire: 3 * time.Minute,
	}
}

func (c *SessionCache) SetTicket(ctx context.Context, ticket string, userId int64) error {
	return c.Redis.Client.SetEx(ctx, constant.GetKeySignalTicket(ticket), userId, time.Minute).Err()
}

func (c *SessionCache) GetTicket(ctx context.Context, ticket string) (int64, error) {
	result, err := c.Redis.Client.Get(ctx, constant.GetKeySignalTicket(ticket)).Result()
	if errors.Is(err, redis.Nil) {
		return 0, cv1.ErrorUnauthorized("ticket is invalid")
	}
	if err != nil {
		return 0, err
	}
	err = c.Redis.Client.Del(ctx, constant.GetKeySignalTicket(ticket)).Err()
	if err != nil {
		return 0, err
	}
	userId, err := strconv.ParseInt(result, 10, 64)
	if err != nil {
		return 0, err
	}
	return userId, nil
}

func (c *SessionCache) SetSession(ctx context.Context, sessionId string, userId int64, nodeKey string) error {
	key := constant.GetKeySignalSession(userId)
	_, err := c.Redis.Client.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.HSet(ctx, key, map[string]interface{}{
			sessionId: nodeKey,
		})
		pipe.HExpire(ctx, key, c.sessionExpire, sessionId)
		pipe.Set(ctx, sessionId, userId, c.sessionExpire)
		pipe.SAdd(ctx, constant.GetKeySignalNodeKeySessions(nodeKey), sessionId)
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *SessionCache) RenewalSessions(ctx context.Context, sessionIds []string) error {
	if len(sessionIds) == 0 {
		return nil
	}
	var err error

	// Pipeline 批量 Get session -> user
	getCmds := make([]*redis.StringCmd, 0, len(sessionIds))
	_, err = c.Redis.Client.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		for _, sessionId := range sessionIds {
			getCmds = append(getCmds, pipe.Get(ctx, constant.GetKeySignalSessionUser(sessionId)))
		}
		return nil
	})
	if errors.Is(err, redis.Nil) {
		return nil
	}
	if err != nil {
		return err
	}

	// 本地去重 userId
	userIdsSet := set.New[int64](0)
	for _, cmd := range getCmds {
		if errors.Is(cmd.Err(), redis.Nil) {
			continue
		}
		uid, err := cmd.Int64()
		if err != nil {
			continue
		}
		userIdsSet.Add(uid)
	}
	if userIdsSet.Len() == 0 {
		return nil
	}

	// Pipeline 批量 Expire user routes
	_, err = c.Redis.Client.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		userIdsSet.ForEach(func(id int64) bool {
			pipe.Expire(ctx, constant.GetKeySignalSession(id), c.sessionExpire)
			return true
		})
		for _, id := range sessionIds {
			pipe.Expire(ctx, id, c.sessionExpire)
		}
		return nil
	})
	return err
}

func (c *SessionCache) GetNodeKeyByUserIds(ctx context.Context, userIds []int64) ([]string, error) {
	var (
		nodeKeys = make([]string, 0)
		err      error
	)
	if len(userIds) == 0 {
		return nodeKeys, nil
	}

	// Pipeline 批量 HGETALL
	cmds := make([]*redis.MapStringStringCmd, 0, len(userIds))
	_, err = c.Redis.Client.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		for _, uid := range userIds {
			cmds = append(cmds, pipe.HGetAll(ctx, constant.GetKeySignalSession(uid)))
		}
		return nil
	})
	if errors.Is(err, redis.Nil) {
		return nodeKeys, nil
	}
	if err != nil {
		return nil, err
	}

	// 去重 nodeKey
	nodeKeysSet := set.New[string](0)
	for _, cmd := range cmds {
		if err := cmd.Err(); err != nil && !errors.Is(err, redis.Nil) {
			continue
		}
		for _, nodeKey := range cmd.Val() {
			nodeKeysSet.Add(nodeKey)
		}
	}
	return nodeKeysSet.ToSlice(), nil
}
