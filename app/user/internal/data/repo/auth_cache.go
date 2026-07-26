package repo

import (
	"common/pkg/client"
	commonenum "common/pkg/enum"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/enum"

	"github.com/redis/go-redis/v9"
)

var _ repo.AuthCacheRepo = (*AuthCacheRepo)(nil)

type AuthCacheRepo struct {
	redisClient            *client.RedisClient
	authCodeKey            string
	authRegisterDraftKey   string
	authRefreshSessionKey  string
	authUserSessionsKey    string
	authRbacPermissionsKey string
}

func NewAuthCacheRepo(
	redisClient *client.RedisClient,
) repo.AuthCacheRepo {
	return &AuthCacheRepo{
		redisClient:            redisClient,
		authCodeKey:            "Auth:Code:{%s}:{%s}",
		authRegisterDraftKey:   "Auth:RegisterDraft:{%s}:{%s}",
		authRefreshSessionKey:  "Auth:Refresh:{%s}",
		authUserSessionsKey:    "Auth:UserSessions:{%d}",
		authRbacPermissionsKey: "Auth:Rbac:{%s}:{%d}",
	}
}

func (r *AuthCacheRepo) authCodeRedisKey(codeType enum.VerificationType, account string) string {
	return fmt.Sprintf(r.authCodeKey, string(codeType), account)
}

func (r *AuthCacheRepo) authRegisterDraftRedisKey(draftType enum.VerificationType, account string) string {
	return fmt.Sprintf(r.authRegisterDraftKey, string(draftType), account)
}

func (r *AuthCacheRepo) authRefreshSessionRedisKey(sessionID string) string {
	return fmt.Sprintf(r.authRefreshSessionKey, sessionID)
}

func (r *AuthCacheRepo) authUserSessionsRedisKey(userID int64) string {
	return fmt.Sprintf(r.authUserSessionsKey, userID)
}

func (r *AuthCacheRepo) authRbacPermissionsRedisKey(realm string, userID int64) string {
	return fmt.Sprintf(r.authRbacPermissionsKey, realm, userID)
}

func (r *AuthCacheRepo) authUserRbacPermissionsPattern(userID int64) string {
	return fmt.Sprintf("Auth:Rbac:*:{%d}", userID)
}

func (r *AuthCacheRepo) authRealmRbacPermissionsPattern(realm string) string {
	return fmt.Sprintf("Auth:Rbac:%s:*", realm)
}

func (r *AuthCacheRepo) SaveCode(ctx context.Context, code *model.VerificationCode, ttl time.Duration) error {
	if code == nil || code.CreatedAt == nil || code.ExpiresAt == nil {
		return errors.New("verification code time is required")
	}
	key := r.authCodeRedisKey(code.Type, code.Account)
	pipe := r.redisClient.Client.TxPipeline()
	pipe.HSet(
		ctx,
		key,
		"type",
		string(code.Type),
		"account",
		code.Account,
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

func (r *AuthCacheRepo) GetCode(ctx context.Context, codeType enum.VerificationType, account string) (*model.VerificationCode, error) {
	values, err := r.redisClient.Client.HGetAll(ctx, r.authCodeRedisKey(codeType, account)).Result()
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
	return &model.VerificationCode{
		Type:        enum.VerificationType(values["type"]),
		Account:     values["account"],
		Code:        values["code"],
		Attempts:    int32(attempts),
		MaxAttempts: int32(maxAttempts),
		CreatedAt:   new(createdAt),
		ExpiresAt:   new(expiresAt),
	}, nil
}

func (r *AuthCacheRepo) IncrCodeAttempts(ctx context.Context, codeType enum.VerificationType, account string) (int64, error) {
	return r.redisClient.Client.HIncrBy(ctx, r.authCodeRedisKey(codeType, account), "attempts", 1).Result()
}

func (r *AuthCacheRepo) DeleteCode(ctx context.Context, codeType enum.VerificationType, account string) error {
	return r.redisClient.Client.Del(ctx, r.authCodeRedisKey(codeType, account)).Err()
}

func (r *AuthCacheRepo) SaveRegisterDraft(ctx context.Context, draftType enum.VerificationType, account string, draft *model.RegisterDraft, ttl time.Duration) error {
	data, err := json.Marshal(draft)
	if err != nil {
		return err
	}
	return r.redisClient.Client.Set(ctx, r.authRegisterDraftRedisKey(draftType, account), data, ttl).Err()
}

func (r *AuthCacheRepo) GetRegisterDraft(ctx context.Context, draftType enum.VerificationType, account string) (*model.RegisterDraft, error) {
	data, err := r.redisClient.Client.Get(ctx, r.authRegisterDraftRedisKey(draftType, account)).Bytes()
	if errors.Is(err, redis.Nil) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	draft := new(model.RegisterDraft)
	if err := json.Unmarshal(data, draft); err != nil {
		return nil, err
	}
	return draft, nil
}

func (r *AuthCacheRepo) DeleteRegisterDraft(ctx context.Context, draftType enum.VerificationType, account string) error {
	return r.redisClient.Client.Del(ctx, r.authRegisterDraftRedisKey(draftType, account)).Err()
}

func (r *AuthCacheRepo) SaveSession(ctx context.Context, session *model.RefreshSession, ttl time.Duration) error {
	if session == nil || session.CreatedAt == nil || session.LastSeenAt == nil || session.SessionExpiresAt == nil {
		return errors.New("refresh session time is required")
	}
	key := r.authRefreshSessionRedisKey(session.SessionID)
	expiresAt := time.Now().Add(ttl)
	pipe := r.redisClient.Client.TxPipeline()
	pipe.HSet(
		ctx,
		key,
		"user_id",
		strconv.FormatInt(session.UserID, 10),
		"realm",
		string(session.Realm),
		"current_jti",
		session.CurrentJTI,
		"created_at",
		session.CreatedAt.Format(time.RFC3339Nano),
		"last_seen_at",
		session.LastSeenAt.Format(time.RFC3339Nano),
		"session_expires_at",
		session.SessionExpiresAt.Format(time.RFC3339Nano),
		"client_type",
		string(session.Client.ClientType),
		"device_type",
		string(session.Client.DeviceType),
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
	)
	pipe.Expire(ctx, key, ttl)
	pipe.ZAdd(ctx, r.authUserSessionsRedisKey(session.UserID), redis.Z{
		Score:  float64(expiresAt.Unix()),
		Member: session.SessionID,
	})
	_, err := pipe.Exec(ctx)
	return err
}

func (r *AuthCacheRepo) GetSession(ctx context.Context, sessionID string) (*model.RefreshSession, error) {
	values, err := r.redisClient.Client.HGetAll(ctx, r.authRefreshSessionRedisKey(sessionID)).Result()
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
	key := r.authRefreshSessionRedisKey(session.SessionID)
	expiresAt := time.Now().Add(ttl)
	pipe := r.redisClient.Client.TxPipeline()
	pipe.HSet(ctx, key, "last_seen_at", session.LastSeenAt.Format(time.RFC3339Nano))
	pipe.Expire(ctx, key, ttl)
	pipe.ZAdd(ctx, r.authUserSessionsRedisKey(session.UserID), redis.Z{
		Score:  float64(expiresAt.Unix()),
		Member: session.SessionID,
	})
	_, err := pipe.Exec(ctx)
	return err
}

func (r *AuthCacheRepo) RotateSessionJTI(ctx context.Context, sessionID string, oldJTI string, newJTI string, lastSeenAt *time.Time, ttl time.Duration) (bool, error) {
	if lastSeenAt == nil {
		return false, errors.New("refresh session last_seen_at is required")
	}
	key := r.authRefreshSessionRedisKey(sessionID)
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

func (r *AuthCacheRepo) DeleteSession(ctx context.Context, userID int64, sessionID string) error {
	pipe := r.redisClient.Client.TxPipeline()
	pipe.Del(ctx, r.authRefreshSessionRedisKey(sessionID))
	if userID > 0 {
		pipe.ZRem(ctx, r.authUserSessionsRedisKey(userID), sessionID)
	}
	_, err := pipe.Exec(ctx)
	return err
}

func (r *AuthCacheRepo) DeleteUserSessions(ctx context.Context, userID int64) error {
	indexKey := r.authUserSessionsRedisKey(userID)
	sids, err := r.redisClient.Client.ZRange(ctx, indexKey, 0, -1).Result()
	if err != nil {
		return err
	}
	pipe := r.redisClient.Client.TxPipeline()
	for _, sid := range sids {
		pipe.Del(ctx, r.authRefreshSessionRedisKey(sid))
	}
	pipe.Del(ctx, indexKey)
	_, err = pipe.Exec(ctx)
	return err
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
