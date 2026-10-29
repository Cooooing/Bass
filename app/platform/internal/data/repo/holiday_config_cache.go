package repo

import (
	"common/pkg/client"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"platform/internal/biz/model"
	bizrepo "platform/internal/biz/repo"

	"github.com/redis/go-redis/v9"
)

var _ bizrepo.HolidayConfigCacheRepo = (*HolidayConfigCacheRepo)(nil)

type HolidayConfigCacheRepo struct {
	redisClient *client.RedisClient
	ttl         time.Duration
	// setIfFresherScript 只在入参快照比缓存更新时写入，避免多实例互相覆盖新数据。
	setIfFresherScript *redis.Script
}

func NewHolidayConfigCacheRepo(redisClient *client.RedisClient) bizrepo.HolidayConfigCacheRepo {
	return &HolidayConfigCacheRepo{
		redisClient: redisClient,
		ttl:         12 * time.Hour,
		setIfFresherScript: redis.NewScript(`
local current = redis.call("GET", KEYS[1])
if current then
	local decoded = cjson.decode(current)
	if (decoded["fetched_at"] or 0) > tonumber(ARGV[1]) then
		return 0
	end
end
redis.call("SET", KEYS[1], ARGV[2], "PX", ARGV[3])
return 1
`),
	}
}

func (r *HolidayConfigCacheRepo) Get(ctx context.Context, year int) (*model.HolidayConfigCache, error) {
	if r == nil || r.redisClient == nil || r.redisClient.Client == nil {
		return nil, fmt.Errorf("holiday config Redis client is required")
	}
	payload, err := r.redisClient.Client.Get(ctx, r.key(year)).Bytes()
	if errors.Is(err, redis.Nil) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var cached holidayConfigSnapshot
	if err = json.Unmarshal(payload, &cached); err != nil {
		return nil, fmt.Errorf("decode holiday config cache: %w", err)
	}
	entries := make(map[string]*model.HolidayEntry, len(cached.Entries))
	for key, entry := range cached.Entries {
		if entry == nil {
			continue
		}
		entries[key] = &model.HolidayEntry{
			Date:    time.UnixMilli(entry.DateMillis).In(time.Local),
			Name:    entry.Name,
			Holiday: entry.Holiday,
		}
	}
	return &model.HolidayConfigCache{
		Entries:   entries,
		FetchedAt: time.UnixMilli(cached.FetchedAtMillis),
	}, nil
}

func (r *HolidayConfigCacheRepo) SetIfFresher(ctx context.Context, year int, cache *model.HolidayConfigCache) error {
	if r == nil || r.redisClient == nil || r.redisClient.Client == nil {
		return fmt.Errorf("holiday config Redis client is required")
	}
	if cache == nil {
		return fmt.Errorf("holiday config cache is required")
	}
	snapshot := holidayConfigSnapshot{
		FetchedAtMillis: cache.FetchedAt.UnixMilli(),
		Entries:         make(map[string]*holidayConfigSnapshotEntry, len(cache.Entries)),
	}
	for key, entry := range cache.Entries {
		if entry == nil {
			continue
		}
		snapshot.Entries[key] = &holidayConfigSnapshotEntry{
			DateMillis: entry.Date.UnixMilli(),
			Name:       entry.Name,
			Holiday:    entry.Holiday,
		}
	}
	payload, err := json.Marshal(snapshot)
	if err != nil {
		return fmt.Errorf("encode holiday config cache: %w", err)
	}
	return r.setIfFresherScript.Run(
		ctx,
		r.redisClient.Client,
		[]string{r.key(year)},
		snapshot.FetchedAtMillis,
		payload,
		r.ttl.Milliseconds(),
	).Err()
}

func (*HolidayConfigCacheRepo) key(year int) string {
	return fmt.Sprintf("platform:holiday_config:%d", year)
}

type holidayConfigSnapshot struct {
	FetchedAtMillis int64                                  `json:"fetched_at"`
	Entries         map[string]*holidayConfigSnapshotEntry `json:"entries"`
}

type holidayConfigSnapshotEntry struct {
	DateMillis int64  `json:"date"`
	Name       string `json:"name"`
	Holiday    bool   `json:"holiday"`
}
