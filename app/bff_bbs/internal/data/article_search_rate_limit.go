package data

import (
	"common/pkg/client"
	"context"
	"fmt"

	"bff_bbs/internal/biz/repo"

	"github.com/redis/go-redis/v9"
)

var _ repo.ArticleSearchRateLimitCache = (*ArticleSearchRateLimitCache)(nil)

type ArticleSearchRateLimitCache struct {
	redisClient *client.RedisClient
	allowScript *redis.Script
}

func NewArticleSearchRateLimitCache(
	redisClient *client.RedisClient,
) repo.ArticleSearchRateLimitCache {
	return &ArticleSearchRateLimitCache{
		redisClient: redisClient,
		allowScript: redis.NewScript(`
local count = redis.call("INCR", KEYS[1])
if count == 1 then
  redis.call("PEXPIRE", KEYS[1], ARGV[1])
end
return count
		`),
	}
}

func (c *ArticleSearchRateLimitCache) Allow(ctx context.Context, req *repo.ArticleSearchRateLimitReq) (bool, error) {
	if c == nil || c.redisClient == nil || c.redisClient.Client == nil || req == nil || req.Key == "" || req.Window <= 0 || req.Limit <= 0 {
		return false, fmt.Errorf("invalid article search rate limit request")
	}
	key := fmt.Sprintf("bbs:article_search:rate_limit:{principal:%s}", req.Key)
	count, err := c.allowScript.Run(
		ctx,
		c.redisClient.Client,
		[]string{key},
		req.Window.Milliseconds(),
	).Int64()
	if err != nil {
		return false, err
	}
	return count <= req.Limit, nil
}
