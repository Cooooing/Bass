package repo

import (
	"context"
	bizrepo "game_idle/internal/biz/repo"
	"time"
)

var _ bizrepo.CharacterSessionRepo = (*CharacterSessionRepo)(nil)

type CharacterSessionRepo struct {
	cache *CharacterStateCache
}

func NewCharacterSessionRepo(cache *CharacterStateCache) bizrepo.CharacterSessionRepo {
	return &CharacterSessionRepo{cache: cache}
}

func (r *CharacterSessionRepo) Online(ctx context.Context, characterID int64, sessionID string, ttlSeconds int64) (string, error) {
	character := r.cache.character(characterID)
	character.mutex.Lock()
	defer character.mutex.Unlock()
	oldSessionID := ""
	if session := character.session; session != nil && time.Now().Before(session.expiresAt) {
		oldSessionID = session.sessionID
	}
	character.session = &characterSession{
		sessionID: sessionID,
		expiresAt: time.Now().Add(time.Duration(ttlSeconds) * time.Second),
	}
	return oldSessionID, nil
}

func (r *CharacterSessionRepo) Ping(ctx context.Context, characterID int64, sessionID string, ttlSeconds int64) (bool, error) {
	character := r.cache.character(characterID)
	character.mutex.Lock()
	defer character.mutex.Unlock()
	session := character.session
	if session == nil || session.sessionID != sessionID || !time.Now().Before(session.expiresAt) {
		return false, nil
	}
	session.expiresAt = time.Now().Add(time.Duration(ttlSeconds) * time.Second)
	return true, nil
}

func (r *CharacterSessionRepo) Offline(ctx context.Context, characterID int64, sessionID string) (bool, error) {
	character := r.cache.character(characterID)
	character.mutex.Lock()
	defer character.mutex.Unlock()
	session := character.session
	if session == nil || session.sessionID != sessionID {
		return false, nil
	}
	character.session = nil
	return true, nil
}

func (r *CharacterSessionRepo) IsOnline(ctx context.Context, characterID int64) (bool, error) {
	character := r.cache.characterIfExists(characterID)
	if character == nil {
		return false, nil
	}
	character.mutex.RLock()
	session := character.session
	character.mutex.RUnlock()
	return session != nil && time.Now().Before(session.expiresAt), nil
}
