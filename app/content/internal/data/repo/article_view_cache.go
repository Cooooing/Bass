package repo

import (
	"common/pkg/client"
	"content/internal/biz/repo"
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/samber/lo"
)

var _ repo.ArticleViewCacheRepo = (*ArticleViewCacheRepo)(nil)

type ArticleViewCacheRepo struct {
	redisClient     *client.RedisClient
	countKey        string
	recordKeyPrefix string
	recordTTL       time.Duration
}

func NewArticleViewCacheRepo(
	redisClient *client.RedisClient,
) repo.ArticleViewCacheRepo {
	return &ArticleViewCacheRepo{
		redisClient:     redisClient,
		countKey:        "Content:ArticleViewCount",
		recordKeyPrefix: "Content:ArticleViewRecord",
		recordTTL:       24 * time.Hour,
	}
}

func (r *ArticleViewCacheRepo) Record(ctx context.Context, req *repo.ArticleViewCacheRecordReq) (bool, error) {
	viewerKey := ""
	if req.ViewerUserID != nil && *req.ViewerUserID > 0 {
		viewerKey = fmt.Sprintf("user:%d", *req.ViewerUserID)
	} else if req.BrowserFingerprint != nil && *req.BrowserFingerprint != "" {
		viewerKey = fmt.Sprintf("fp:%s", *req.BrowserFingerprint)
	} else {
		ip := ""
		if req.IP != nil {
			ip = *req.IP
		}
		userAgent := ""
		if req.UserAgent != nil {
			userAgent = *req.UserAgent
		}
		viewerKey = fmt.Sprintf("ipua:%s:%s", ip, userAgent)
	}
	if viewerKey == "ipua::" {
		return false, nil
	}
	recordKey := fmt.Sprintf("%s:%d", r.recordKeyPrefix, req.ArticleID)
	ok, err := r.redisClient.Client.HSetNX(ctx, recordKey, viewerKey, time.Now().Format(time.RFC3339Nano)).Result()
	if err != nil || !ok {
		return ok, err
	}
	if err = r.redisClient.Client.Expire(ctx, recordKey, r.recordTTL).Err(); err != nil {
		return false, err
	}
	if err = r.redisClient.Client.HIncrBy(ctx, r.countKey, strconv.FormatInt(req.ArticleID, 10), 1).Err(); err != nil {
		return false, err
	}
	return true, nil
}

func (r *ArticleViewCacheRepo) PopCounts(ctx context.Context, limit int32) (map[int64]int32, error) {
	values, err := r.redisClient.Client.HGetAll(ctx, r.countKey).Result()
	if err != nil {
		return nil, err
	}
	if limit <= 0 {
		limit = 1000
	}
	rows := make(map[int64]int32, len(values))
	fields := make([]string, 0, len(values))
	for articleIDValue, countValue := range values {
		if int32(len(rows)) >= limit {
			break
		}
		articleID, err := strconv.ParseInt(articleIDValue, 10, 64)
		if err != nil {
			continue
		}
		count, err := strconv.ParseInt(countValue, 10, 32)
		if err != nil || count <= 0 {
			fields = append(fields, articleIDValue)
			continue
		}
		rows[articleID] = int32(count)
		fields = append(fields, articleIDValue)
	}
	if len(fields) > 0 {
		if err = r.redisClient.Client.HDel(ctx, r.countKey, lo.Uniq(fields)...).Err(); err != nil {
			return nil, err
		}
	}
	return rows, nil
}
