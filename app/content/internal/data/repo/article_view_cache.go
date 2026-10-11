package repo

import (
	"common/pkg/client"
	"context"
	"fmt"
	"strconv"
	"time"

	"content/internal/biz/repo"

	"github.com/samber/lo"
)

var _ repo.ArticleViewCacheRepo = (*ArticleViewCacheRepo)(nil)

type ArticleViewCacheRepo struct {
	redisClient     *client.RedisClient
	countKey        string
	recordKeyFormat string
	recordTTL       time.Duration
}

func NewArticleViewCacheRepo(
	redisClient *client.RedisClient,
) repo.ArticleViewCacheRepo {
	return &ArticleViewCacheRepo{
		redisClient:     redisClient,
		countKey:        "content:article_view:count",
		recordKeyFormat: "content:article_view:record:{article_id:%d}",
		recordTTL:       24 * time.Hour,
	}
}

func (r *ArticleViewCacheRepo) Record(ctx context.Context, req *repo.ArticleViewCacheRecordReq) (bool, error) {
	var viewerKey string
	switch {
	case req.ViewerUserID != nil && *req.ViewerUserID > 0:
		viewerKey = fmt.Sprintf("viewer:{user_id:%d}", *req.ViewerUserID)
	case req.BrowserFingerprint != nil && *req.BrowserFingerprint != "":
		viewerKey = fmt.Sprintf("viewer:{fingerprint:%s}", *req.BrowserFingerprint)
	default:
		ip := ""
		if req.IP != nil {
			ip = *req.IP
		}
		userAgent := ""
		if req.UserAgent != nil {
			userAgent = *req.UserAgent
		}
		viewerKey = fmt.Sprintf("viewer:{ip:%s}:ua:{user_agent:%s}", ip, userAgent)
	}
	if viewerKey == "viewer:{ip:}:ua:{user_agent:}" {
		return false, nil
	}
	recordKey := fmt.Sprintf(r.recordKeyFormat, req.ArticleID)
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
		if len(rows) >= int(limit) {
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
