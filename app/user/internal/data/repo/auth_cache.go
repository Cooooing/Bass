package repo

import (
	"common/pkg/client"
	commonenum "common/pkg/enum"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/enum"

	"github.com/redis/go-redis/v9"
)

var _ repo.AuthCacheRepo = (*AuthCacheRepo)(nil)

type AuthCacheRepo struct {
	redisClient            *client.RedisClient
	authOtpGuestKey        string
	authOtpUserKey         string
	authRefreshSessionKey  string
	authRefreshSessionHead string
	authUserSessionsKey    string
	authUserSessionsScan   string
	authRbacPermissionsKey string
}

func NewAuthCacheRepo(
	redisClient *client.RedisClient,
) repo.AuthCacheRepo {
	return &AuthCacheRepo{
		redisClient:     redisClient,
		authOtpGuestKey: "user:auth:otp:{type:%s}:guest:{account:%s}",
		authOtpUserKey:  "user:auth:otp:{type:%s}:user:{user_id:%d}:account:{account:%s}",

		authRefreshSessionKey:  "user:auth:refresh:{realm:%s}:{session_id:%s}",
		authRefreshSessionHead: "user:auth:refresh:{realm:%s}:",
		authUserSessionsKey:    "user:auth:sessions:{realm:%s}:{user_id:%d}",
		authUserSessionsScan:   "user:auth:sessions:{realm:*}:{user_id:%d}",
		authRbacPermissionsKey: "user:auth:rbac_permissions:{realm:%s}:{user_id:%d}",
	}
}

func (r *AuthCacheRepo) authCodeRedisKey(req *repo.VerificationCodeKeyReq) string {
	if req.UserID != nil {
		return fmt.Sprintf(r.authOtpUserKey, req.Type.String(), *req.UserID, req.Account)
	}
	return fmt.Sprintf(r.authOtpGuestKey, req.Type.String(), req.Account)
}

func (r *AuthCacheRepo) authRefreshSessionRedisKey(realm commonenum.LoginRealm, sessionID string) string {
	return fmt.Sprintf(r.authRefreshSessionKey, realm.String(), sessionID)
}

func (r *AuthCacheRepo) authRefreshSessionRedisHead(realm commonenum.LoginRealm) string {
	return fmt.Sprintf(r.authRefreshSessionHead, realm.String())
}

func (r *AuthCacheRepo) authUserSessionsRedisKey(realm commonenum.LoginRealm, userID int64) string {
	return fmt.Sprintf(r.authUserSessionsKey, realm.String(), userID)
}

func (r *AuthCacheRepo) authUserSessionsRedisScan(userID int64) string {
	return fmt.Sprintf(r.authUserSessionsScan, userID)
}

func (r *AuthCacheRepo) authRbacPermissionsRedisKey(realm string, userID int64) string {
	return fmt.Sprintf(r.authRbacPermissionsKey, realm, userID)
}

func (r *AuthCacheRepo) authUserRbacPermissionsPattern(userID int64) string {
	return fmt.Sprintf("user:auth:rbac_permissions:{realm:*}:{user_id:%d}", userID)
}

func (r *AuthCacheRepo) authRealmRbacPermissionsPattern(realm string) string {
	return fmt.Sprintf("user:auth:rbac_permissions:{realm:%s}:{user_id:*}", realm)
}

func (r *AuthCacheRepo) SaveCode(ctx context.Context, code *model.VerificationCode, ttl time.Duration) error {
	if code == nil || code.CreatedAt == nil || code.ExpiresAt == nil {
		return errors.New("verification code time is required")
	}
	key := r.authCodeRedisKey(&repo.VerificationCodeKeyReq{
		Type:    code.Type,
		Account: code.Account,
		UserID:  code.UserID,
	})
	userID := ""
	if code.UserID != nil {
		userID = strconv.FormatInt(*code.UserID, 10)
	}
	pipe := r.redisClient.Client.TxPipeline()
	pipe.HSet(
		ctx,
		key,
		"type",
		code.Type.String(),
		"account",
		code.Account,
		"user_id",
		userID,
		"code",
		code.Code,
		"attempts",
		strconv.FormatInt(int64(code.Attempts), 10),
		"max_attempts",
		strconv.FormatInt(int64(code.MaxAttempts), 10),
		"created_at",
		code.CreatedAt.Format(time.RFC3339Nano),
		"expires_at",
		code.ExpiresAt.Format(time.RFC3339Nano),
	)
	pipe.Expire(ctx, key, ttl)
	_, err := pipe.Exec(ctx)
	return err
}

func (r *AuthCacheRepo) GetCode(ctx context.Context, req *repo.VerificationCodeKeyReq) (*model.VerificationCode, error) {
	values, err := r.redisClient.Client.HGetAll(ctx, r.authCodeRedisKey(req)).Result()
	if err != nil {
		return nil, err
	}
	if len(values) == 0 {
		return nil, nil
	}
	attempts, err := strconv.ParseInt(values["attempts"], 10, 32)
	if err != nil {
		return nil, err
	}
	maxAttempts, err := strconv.ParseInt(values["max_attempts"], 10, 32)
	if err != nil {
		return nil, err
	}
	createdAt, err := time.Parse(time.RFC3339Nano, values["created_at"])
	if err != nil {
		return nil, err
	}
	expiresAt, err := time.Parse(time.RFC3339Nano, values["expires_at"])
	if err != nil {
		return nil, err
	}
	out := &model.VerificationCode{
		Type:        enum.VerificationType(values["type"]),
		Account:     values["account"],
		Code:        values["code"],
		Attempts:    int32(attempts),
		MaxAttempts: int32(maxAttempts),
		CreatedAt:   new(createdAt),
		ExpiresAt:   new(expiresAt),
	}
	if values["user_id"] != "" {
		userID, err := strconv.ParseInt(values["user_id"], 10, 64)
		if err != nil {
			return nil, err
		}
		out.UserID = new(userID)
	}
	return out, nil
}

func (r *AuthCacheRepo) IncrCodeAttempts(ctx context.Context, req *repo.VerificationCodeKeyReq) (int64, error) {
	return r.redisClient.Client.HIncrBy(ctx, r.authCodeRedisKey(req), "attempts", 1).Result()
}

func (r *AuthCacheRepo) DeleteCode(ctx context.Context, req *repo.VerificationCodeKeyReq) error {
	return r.redisClient.Client.Del(ctx, r.authCodeRedisKey(req)).Err()
}

func (r *AuthCacheRepo) SaveSession(ctx context.Context, session *model.RefreshSession, ttl time.Duration, maxSessions int) error {
	if session == nil || session.CreatedAt == nil || session.LastSeenAt == nil || session.SessionExpiresAt == nil {
		return errors.New("refresh session time is required")
	}
	key := r.authRefreshSessionRedisKey(session.Realm, session.SessionID)
	expiresAt := time.Now().Add(ttl)
	script := redis.NewScript(`
redis.call('HSET', KEYS[1], unpack(ARGV, 7))
redis.call('PEXPIRE', KEYS[1], ARGV[1])
redis.call('ZREMRANGEBYSCORE', KEYS[2], '-inf', ARGV[6])
redis.call('ZADD', KEYS[2], ARGV[2], ARGV[3])
local max_sessions = tonumber(ARGV[4])
if max_sessions and max_sessions > 0 then
  local count = redis.call('ZCARD', KEYS[2])
  if count > max_sessions then
    local overflow = count - max_sessions
    local sids = redis.call('ZRANGE', KEYS[2], 0, overflow - 1)
    for _, sid in ipairs(sids) do
      redis.call('DEL', ARGV[5] .. '{session_id:' .. sid .. '}')
      redis.call('ZREM', KEYS[2], sid)
    end
  end
end
return 1
`)
	return script.Run(
		ctx,
		r.redisClient.Client,
		[]string{key, r.authUserSessionsRedisKey(session.Realm, session.UserID)},
		strconv.FormatInt(ttl.Milliseconds(), 10),
		strconv.FormatInt(expiresAt.Unix(), 10),
		session.SessionID,
		strconv.FormatInt(int64(maxSessions), 10),
		r.authRefreshSessionRedisHead(session.Realm),
		strconv.FormatInt(time.Now().Unix(), 10),
		"user_id",
		strconv.FormatInt(session.UserID, 10),
		"realm",
		session.Realm.String(),
		"current_jti",
		session.CurrentJTI,
		"created_at",
		session.CreatedAt.Format(time.RFC3339Nano),
		"last_seen_at",
		session.LastSeenAt.Format(time.RFC3339Nano),
		"session_expires_at",
		session.SessionExpiresAt.Format(time.RFC3339Nano),
		"client_type",
		session.Client.ClientType.String(),
		"device_type",
		session.Client.DeviceType.String(),
		"os_name",
		session.Client.OSName,
		"os_version",
		session.Client.OSVersion,
		"browser_name",
		session.Client.BrowserName,
		"browser_version",
		session.Client.BrowserVersion,
		"app_name",
		session.Client.AppName,
		"app_version",
		session.Client.AppVersion,
		"user_agent",
		session.Client.UserAgent,
	).Err()
}

func (r *AuthCacheRepo) GetSession(ctx context.Context, realm commonenum.LoginRealm, sessionID string) (*model.RefreshSession, error) {
	values, err := r.redisClient.Client.HGetAll(ctx, r.authRefreshSessionRedisKey(realm, sessionID)).Result()
	if err != nil {
		return nil, err
	}
	if len(values) == 0 {
		return nil, nil
	}
	userID, err := strconv.ParseInt(values["user_id"], 10, 64)
	if err != nil {
		return nil, err
	}
	createdAt, err := time.Parse(time.RFC3339Nano, values["created_at"])
	if err != nil {
		return nil, err
	}
	lastSeenAt, err := time.Parse(time.RFC3339Nano, values["last_seen_at"])
	if err != nil {
		return nil, err
	}
	sessionExpiresAt, err := time.Parse(time.RFC3339Nano, values["session_expires_at"])
	if err != nil {
		return nil, err
	}
	return &model.RefreshSession{
		SessionID:        sessionID,
		UserID:           userID,
		Realm:            commonenum.LoginRealm(values["realm"]),
		CurrentJTI:       values["current_jti"],
		CreatedAt:        new(createdAt),
		LastSeenAt:       new(lastSeenAt),
		SessionExpiresAt: new(sessionExpiresAt),
		Client: model.LoginContext{
			ClientType:     enum.ClientType(values["client_type"]),
			DeviceType:     enum.DeviceType(values["device_type"]),
			OSName:         values["os_name"],
			OSVersion:      values["os_version"],
			BrowserName:    values["browser_name"],
			BrowserVersion: values["browser_version"],
			AppName:        values["app_name"],
			AppVersion:     values["app_version"],
			UserAgent:      values["user_agent"],
		},
	}, nil
}

func (r *AuthCacheRepo) TouchSession(ctx context.Context, session *model.RefreshSession, ttl time.Duration) error {
	if session == nil || session.LastSeenAt == nil {
		return errors.New("refresh session last_seen_at is required")
	}
	key := r.authRefreshSessionRedisKey(session.Realm, session.SessionID)
	expiresAt := time.Now().Add(ttl)
	pipe := r.redisClient.Client.TxPipeline()
	pipe.HSet(ctx, key, "last_seen_at", session.LastSeenAt.Format(time.RFC3339Nano))
	pipe.Expire(ctx, key, ttl)
	pipe.ZAdd(ctx, r.authUserSessionsRedisKey(session.Realm, session.UserID), redis.Z{
		Score:  float64(expiresAt.Unix()),
		Member: session.SessionID,
	})
	_, err := pipe.Exec(ctx)
	return err
}

func (r *AuthCacheRepo) RotateSessionJTI(ctx context.Context, realm commonenum.LoginRealm, sessionID string, oldJTI string, newJTI string, lastSeenAt *time.Time, ttl time.Duration) (bool, error) {
	if lastSeenAt == nil {
		return false, errors.New("refresh session last_seen_at is required")
	}
	key := r.authRefreshSessionRedisKey(realm, sessionID)
	script := redis.NewScript(`
local current = redis.call('HGET', KEYS[1], 'current_jti')
if not current then
  return 0
end
if current ~= ARGV[1] then
  return -1
end
redis.call('HSET', KEYS[1], 'current_jti', ARGV[2], 'last_seen_at', ARGV[3])
redis.call('PEXPIRE', KEYS[1], ARGV[4])
return 1
`)
	result, err := script.Run(ctx, r.redisClient.Client, []string{key}, oldJTI, newJTI, lastSeenAt.Format(time.RFC3339Nano), strconv.FormatInt(ttl.Milliseconds(), 10)).Int()
	if err != nil {
		return false, err
	}
	if result == -1 {
		return false, nil
	}
	return result == 1, nil
}

func (r *AuthCacheRepo) DeleteSession(ctx context.Context, realm commonenum.LoginRealm, userID int64, sessionID string) error {
	pipe := r.redisClient.Client.TxPipeline()
	pipe.Del(ctx, r.authRefreshSessionRedisKey(realm, sessionID))
	if userID > 0 {
		pipe.ZRem(ctx, r.authUserSessionsRedisKey(realm, userID), sessionID)
	}
	_, err := pipe.Exec(ctx)
	return err
}

func (r *AuthCacheRepo) DeleteUserSessions(ctx context.Context, userID int64) error {
	var cursor uint64
	for {
		indexKeys, next, err := r.redisClient.Client.Scan(ctx, cursor, r.authUserSessionsRedisScan(userID), 100).Result()
		if err != nil {
			return err
		}
		for _, indexKey := range indexKeys {
			sids, err := r.redisClient.Client.ZRange(ctx, indexKey, 0, -1).Result()
			if err != nil {
				return err
			}
			realm := ""
			if strings.HasPrefix(indexKey, "user:auth:sessions:{realm:") {
				value := strings.TrimPrefix(indexKey, "user:auth:sessions:{realm:")
				end := strings.Index(value, "}")
				if end > 0 {
					realm = value[:end]
				}
			}
			pipe := r.redisClient.Client.TxPipeline()
			if realm != "" {
				for _, sid := range sids {
					pipe.Del(ctx, r.authRefreshSessionRedisKey(commonenum.LoginRealm(realm), sid))
				}
			}
			pipe.Del(ctx, indexKey)
			if _, err = pipe.Exec(ctx); err != nil {
				return err
			}
		}
		if next == 0 {
			return nil
		}
		cursor = next
	}
}
func (r *AuthCacheRepo) SaveRbacPermissions(ctx context.Context, realm string, userID int64, permissions []string, ttl time.Duration) error {
	data, err := json.Marshal(permissions)
	if err != nil {
		return err
	}
	return r.redisClient.Client.Set(ctx, r.authRbacPermissionsRedisKey(realm, userID), data, ttl).Err()
}

func (r *AuthCacheRepo) GetRbacPermissions(ctx context.Context, realm string, userID int64) ([]string, bool, error) {
	data, err := r.redisClient.Client.Get(ctx, r.authRbacPermissionsRedisKey(realm, userID)).Bytes()
	if errors.Is(err, redis.Nil) {
		return nil, false, nil
	}
	if err != nil {
		return nil, false, err
	}
	var permissions []string
	if err := json.Unmarshal(data, &permissions); err != nil {
		return nil, false, err
	}
	return permissions, true, nil
}

func (r *AuthCacheRepo) DeleteRbacPermissions(ctx context.Context, realm string, userID int64) error {
	return r.redisClient.Client.Del(ctx, r.authRbacPermissionsRedisKey(realm, userID)).Err()
}

func (r *AuthCacheRepo) DeleteUserRbacPermissions(ctx context.Context, userID int64) error {
	return r.deleteByPattern(ctx, r.authUserRbacPermissionsPattern(userID))
}

func (r *AuthCacheRepo) DeleteRealmRbacPermissions(ctx context.Context, realm string) error {
	return r.deleteByPattern(ctx, r.authRealmRbacPermissionsPattern(realm))
}

func (r *AuthCacheRepo) deleteByPattern(ctx context.Context, pattern string) error {
	var cursor uint64
	for {
		keys, next, err := r.redisClient.Client.Scan(ctx, cursor, pattern, 100).Result()
		if err != nil {
			return err
		}
		if len(keys) > 0 {
			if err := r.redisClient.Client.Del(ctx, keys...).Err(); err != nil {
				return err
			}
		}
		if next == 0 {
			return nil
		}
		cursor = next
	}
}
