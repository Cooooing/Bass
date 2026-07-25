package repo

import (
	commonenum "common/pkg/enum"
	"common/pkg/client"
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
	redisClient *client.RedisClient
}

func NewAuthCacheRepo(redisClient *client.RedisClient) repo.AuthCacheRepo {
	return &AuthCacheRepo{redisClient: redisClient}
}

const (
	authCodeKey            = "Auth:Code:{%s}:{%s}"
	authRegisterDraftKey   = "Auth:RegisterDraft:{%s}:{%s}"
	authRefreshSessionKey  = "Auth:Refresh:{%s}"
	authUserSessionsKey    = "Auth:UserSessions:{%d}"
	authRbacPermissionsKey = "Auth:Rbac:{%s}:{%d}"
)

func authCodeRedisKey(codeType string, account string) string {
	return fmt.Sprintf(authCodeKey, codeType, account)
}

func authRegisterDraftRedisKey(draftType string, account string) string {
	return fmt.Sprintf(authRegisterDraftKey, draftType, account)
}

func authRefreshSessionRedisKey(sessionID string) string {
	return fmt.Sprintf(authRefreshSessionKey, sessionID)
}

func authUserSessionsRedisKey(userID int64) string {
	return fmt.Sprintf(authUserSessionsKey, userID)
}

func authRbacPermissionsRedisKey(realm string, userID int64) string {
	return fmt.Sprintf(authRbacPermissionsKey, realm, userID)
}

func authUserRbacPermissionsPattern(userID int64) string {
	return fmt.Sprintf("Auth:Rbac:*:{%d}", userID)
}

func authRealmRbacPermissionsPattern(realm string) string {
	return fmt.Sprintf("Auth:Rbac:%s:*", realm)
}

func (r *AuthCacheRepo) SaveCode(ctx context.Context, code *model.VerificationCode, ttl time.Duration) error {
	key := authCodeRedisKey(code.Type, code.Account)
	values := map[string]any{
		"type":            code.Type,
		"account":         code.Account,
		"code":            code.Code,
		"attempts":        code.Attempts,
		"max_attempts":    code.MaxAttempts,
		"created_at_unix": code.CreatedAtUnix,
		"expires_at_unix": code.ExpiresAtUnix,
	}
	pipe := r.redisClient.Client.TxPipeline()
	pipe.HSet(ctx, key, values)
	pipe.Expire(ctx, key, ttl)
	_, err := pipe.Exec(ctx)
	return err
}

func (r *AuthCacheRepo) GetCode(ctx context.Context, codeType string, account string) (*model.VerificationCode, error) {
	values, err := r.redisClient.Client.HGetAll(ctx, authCodeRedisKey(codeType, account)).Result()
	if err != nil {
		return nil, err
	}
	if len(values) == 0 {
		return nil, nil
	}
	return &model.VerificationCode{
		Type:          values["type"],
		Account:       values["account"],
		Code:          values["code"],
		Attempts:      parseInt32(values["attempts"]),
		MaxAttempts:   parseInt32(values["max_attempts"]),
		CreatedAtUnix: parseInt64(values["created_at_unix"]),
		ExpiresAtUnix: parseInt64(values["expires_at_unix"]),
	}, nil
}

func (r *AuthCacheRepo) IncrCodeAttempts(ctx context.Context, codeType string, account string) (int64, error) {
	return r.redisClient.Client.HIncrBy(ctx, authCodeRedisKey(codeType, account), "attempts", 1).Result()
}

func (r *AuthCacheRepo) DeleteCode(ctx context.Context, codeType string, account string) error {
	return r.redisClient.Client.Del(ctx, authCodeRedisKey(codeType, account)).Err()
}

func (r *AuthCacheRepo) SaveRegisterDraft(ctx context.Context, draftType string, account string, draft *model.RegisterDraft, ttl time.Duration) error {
	data, err := json.Marshal(draft)
	if err != nil {
		return err
	}
	return r.redisClient.Client.Set(ctx, authRegisterDraftRedisKey(draftType, account), data, ttl).Err()
}

func (r *AuthCacheRepo) GetRegisterDraft(ctx context.Context, draftType string, account string) (*model.RegisterDraft, error) {
	data, err := r.redisClient.Client.Get(ctx, authRegisterDraftRedisKey(draftType, account)).Bytes()
	if errors.Is(err, redis.Nil) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var draft model.RegisterDraft
	if err := json.Unmarshal(data, &draft); err != nil {
		return nil, err
	}
	return &draft, nil
}

func (r *AuthCacheRepo) DeleteRegisterDraft(ctx context.Context, draftType string, account string) error {
	return r.redisClient.Client.Del(ctx, authRegisterDraftRedisKey(draftType, account)).Err()
}

func (r *AuthCacheRepo) SaveSession(ctx context.Context, session *model.RefreshSession, ttl time.Duration) error {
	key := authRefreshSessionRedisKey(session.SessionID)
	values := sessionHash(session)
	expiresAt := time.Now().Add(ttl).Unix()
	pipe := r.redisClient.Client.TxPipeline()
	pipe.HSet(ctx, key, values)
	pipe.Expire(ctx, key, ttl)
	pipe.ZAdd(ctx, authUserSessionsRedisKey(session.UserID), redis.Z{Score: float64(expiresAt), Member: session.SessionID})
	_, err := pipe.Exec(ctx)
	return err
}

func (r *AuthCacheRepo) GetSession(ctx context.Context, sessionID string) (*model.RefreshSession, error) {
	values, err := r.redisClient.Client.HGetAll(ctx, authRefreshSessionRedisKey(sessionID)).Result()
	if err != nil {
		return nil, err
	}
	if len(values) == 0 {
		return nil, nil
	}
	return sessionFromHash(sessionID, values), nil
}

func (r *AuthCacheRepo) TouchSession(ctx context.Context, session *model.RefreshSession, ttl time.Duration) error {
	key := authRefreshSessionRedisKey(session.SessionID)
	expiresAt := time.Now().Add(ttl).Unix()
	pipe := r.redisClient.Client.TxPipeline()
	pipe.HSet(ctx, key, "last_seen_at_unix", session.LastSeenAtUnix)
	pipe.Expire(ctx, key, ttl)
	pipe.ZAdd(ctx, authUserSessionsRedisKey(session.UserID), redis.Z{Score: float64(expiresAt), Member: session.SessionID})
	_, err := pipe.Exec(ctx)
	return err
}

func (r *AuthCacheRepo) RotateSessionJTI(ctx context.Context, sessionID string, oldJTI string, newJTI string, lastSeenAtUnix int64, ttl time.Duration) (bool, error) {
	key := authRefreshSessionRedisKey(sessionID)
	script := redis.NewScript(`
local current = redis.call('HGET', KEYS[1], 'current_jti')
if not current then
  return 0
end
if current ~= ARGV[1] then
  return -1
end
redis.call('HSET', KEYS[1], 'current_jti', ARGV[2], 'last_seen_at_unix', ARGV[3])
redis.call('PEXPIRE', KEYS[1], ARGV[4])
return 1
`)
	result, err := script.Run(ctx, r.redisClient.Client, []string{key}, oldJTI, newJTI, strconv.FormatInt(lastSeenAtUnix, 10), strconv.FormatInt(ttl.Milliseconds(), 10)).Int()
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
	pipe.Del(ctx, authRefreshSessionRedisKey(sessionID))
	if userID > 0 {
		pipe.ZRem(ctx, authUserSessionsRedisKey(userID), sessionID)
	}
	_, err := pipe.Exec(ctx)
	return err
}

func (r *AuthCacheRepo) DeleteUserSessions(ctx context.Context, userID int64) error {
	indexKey := authUserSessionsRedisKey(userID)
	sids, err := r.redisClient.Client.ZRange(ctx, indexKey, 0, -1).Result()
	if err != nil {
		return err
	}
	pipe := r.redisClient.Client.TxPipeline()
	for _, sid := range sids {
		pipe.Del(ctx, authRefreshSessionRedisKey(sid))
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
	return r.redisClient.Client.Set(ctx, authRbacPermissionsRedisKey(realm, userID), data, ttl).Err()
}

func (r *AuthCacheRepo) GetRbacPermissions(ctx context.Context, realm string, userID int64) ([]string, bool, error) {
	data, err := r.redisClient.Client.Get(ctx, authRbacPermissionsRedisKey(realm, userID)).Bytes()
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
	return r.redisClient.Client.Del(ctx, authRbacPermissionsRedisKey(realm, userID)).Err()
}

func (r *AuthCacheRepo) DeleteUserRbacPermissions(ctx context.Context, userID int64) error {
	return r.deleteByPattern(ctx, authUserRbacPermissionsPattern(userID))
}

func (r *AuthCacheRepo) DeleteRealmRbacPermissions(ctx context.Context, realm string) error {
	return r.deleteByPattern(ctx, authRealmRbacPermissionsPattern(realm))
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

func sessionHash(session *model.RefreshSession) map[string]any {
	return map[string]any{
		"user_id":                 session.UserID,
		"realm":                   string(session.Realm),
		"current_jti":             session.CurrentJTI,
		"created_at_unix":         session.CreatedAtUnix,
		"last_seen_at_unix":       session.LastSeenAtUnix,
		"session_expires_at_unix": session.SessionExpiresAtUnix,
		"client_type":             string(session.Client.ClientType),
		"device_type":             string(session.Client.DeviceType),
		"os_name":                 session.Client.OSName,
		"os_version":              session.Client.OSVersion,
		"browser_name":            session.Client.BrowserName,
		"browser_version":         session.Client.BrowserVersion,
		"app_name":                session.Client.AppName,
		"app_version":             session.Client.AppVersion,
		"user_agent":              session.Client.UserAgent,
	}
}

func sessionFromHash(sessionID string, values map[string]string) *model.RefreshSession {
	return &model.RefreshSession{
		SessionID:            sessionID,
		UserID:               parseInt64(values["user_id"]),
		Realm:                commonenum.LoginRealm(values["realm"]),
		CurrentJTI:           values["current_jti"],
		CreatedAtUnix:        parseInt64(values["created_at_unix"]),
		LastSeenAtUnix:       parseInt64(values["last_seen_at_unix"]),
		SessionExpiresAtUnix: parseInt64(values["session_expires_at_unix"]),
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
	}
}

func parseInt64(value string) int64 {
	parsed, _ := strconv.ParseInt(value, 10, 64)
	return parsed
}

func parseInt32(value string) int32 {
	return int32(parseInt64(value))
}
