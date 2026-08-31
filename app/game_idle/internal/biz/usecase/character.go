package usecase

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	"context"
	"game_idle/internal/biz/model"
	"game_idle/internal/biz/repo"
	"game_idle/internal/config"
	"game_idle/internal/enum"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
)

type CharacterUsecase struct {
	characterRepo              repo.CharacterRepo
	characterSessionRepo       repo.CharacterSessionRepo
	stateUsecase               *CharacterStateUsecase
	eventUsecase               *GameIdleEventUsecase
	namePattern                *regexp.Regexp
	maxCharacterCountPerUser   int32
	defaultActionQueueCapacity int32
	defaultMaxOfflineDuration  time.Duration
	onlineSessionTTL           time.Duration
}

func NewCharacterUsecase(
	conf *config.Bootstrap,
	characterRepo repo.CharacterRepo,
	characterSessionRepo repo.CharacterSessionRepo,
	stateUsecase *CharacterStateUsecase,
	eventUsecase *GameIdleEventUsecase,
) *CharacterUsecase {
	maxCharacterCountPerUser := int32(3)
	if conf.GetGameIdle().GetCharacter().GetMaxCountPerUser() > 0 {
		maxCharacterCountPerUser = int32(conf.GetGameIdle().GetCharacter().GetMaxCountPerUser())
	}
	defaultActionQueueCapacity := int32(3)
	if conf.GetGameIdle().GetCharacter().GetDefaultActionQueueCapacity() > 0 {
		defaultActionQueueCapacity = int32(conf.GetGameIdle().GetCharacter().GetDefaultActionQueueCapacity())
	}
	defaultMaxOfflineDuration := 8 * time.Hour
	if conf.GetGameIdle().GetCharacter().GetDefaultMaxOfflineDuration() != nil {
		defaultMaxOfflineDuration = conf.GetGameIdle().GetCharacter().GetDefaultMaxOfflineDuration().AsDuration()
	}
	onlineSessionTTL := 5 * time.Minute
	if conf.GetGameIdle().GetOnlineSession().GetTtl() != nil && conf.GetGameIdle().GetOnlineSession().GetTtl().AsDuration() > 0 {
		onlineSessionTTL = conf.GetGameIdle().GetOnlineSession().GetTtl().AsDuration()
	}
	return &CharacterUsecase{
		characterRepo:              characterRepo,
		characterSessionRepo:       characterSessionRepo,
		stateUsecase:               stateUsecase,
		eventUsecase:               eventUsecase,
		namePattern:                regexp.MustCompile("^[A-Za-z0-9_]{4,32}$"),
		maxCharacterCountPerUser:   maxCharacterCountPerUser,
		defaultActionQueueCapacity: defaultActionQueueCapacity,
		defaultMaxOfflineDuration:  defaultMaxOfflineDuration,
		onlineSessionTTL:           onlineSessionTTL,
	}
}

type CreateCharacterReq struct {
	UserID int64
	Name   string
}

func (u *CharacterUsecase) Create(ctx context.Context, req *CreateCharacterReq) (*model.Character, error) {
	if req.UserID <= 0 || !u.namePattern.MatchString(req.Name) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_CHARACTER_INVALID)
	}
	nameKey := strings.ToLower(req.Name)
	sameNameRows, err := u.characterRepo.List(ctx, &repo.ListCharacterReq{
		NameKey: &nameKey,
	})
	if err != nil {
		return nil, err
	}
	if len(sameNameRows) > 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_CHARACTER_NAME_TAKEN)
	}
	rows, err := u.characterRepo.List(ctx, &repo.ListCharacterReq{
		UserID: &req.UserID,
	})
	if err != nil {
		return nil, err
	}
	if len(rows) >= int(u.maxCharacterCountPerUser) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_CHARACTER_LIMIT_EXCEEDED)
	}
	usedSlots := make(map[int32]struct{}, len(rows))
	for _, row := range rows {
		if row.Slot > 0 {
			usedSlots[row.Slot] = struct{}{}
		}
	}
	slot := int32(0)
	for candidate := int32(1); candidate <= u.maxCharacterCountPerUser; candidate++ {
		if _, ok := usedSlots[candidate]; !ok {
			slot = candidate
			break
		}
	}
	if slot == 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_CHARACTER_LIMIT_EXCEEDED)
	}
	character, err := u.characterRepo.Save(ctx, &model.Character{
		UserID:              req.UserID,
		Slot:                slot,
		Name:                req.Name,
		NameKey:             nameKey,
		ActionQueueCapacity: u.defaultActionQueueCapacity,
		MaxOfflineDuration:  u.defaultMaxOfflineDuration,
		Status:              enum.CharacterStatusActive,
	})
	if err != nil && strings.Contains(err.Error(), "game_idle_characters_name_key_active_unique") {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_CHARACTER_NAME_TAKEN)
	}
	if err != nil && strings.Contains(err.Error(), "game_idle_characters_user_slot_active_unique") {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_CHARACTER_LIMIT_EXCEEDED)
	}
	return character, err
}

type GetCharacterReq struct {
	UserID      int64
	CharacterID int64
}

func (u *CharacterUsecase) Get(ctx context.Context, req *GetCharacterReq) ([]*model.Character, error) {
	if req.UserID <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_CHARACTER_INVALID)
	}
	return u.characterRepo.List(ctx, &repo.ListCharacterReq{
		UserID:      &req.UserID,
		CharacterID: &req.CharacterID,
	})
}

type OnlineCharacterReq struct {
	CharacterID int64
}

func (u *CharacterUsecase) Online(ctx context.Context, req *OnlineCharacterReq) (*model.CharacterSession, error) {
	var session *model.CharacterSession
	err := u.stateUsecase.Execute(ctx, req.CharacterID, func(runCtx context.Context) error {
		character, err := u.characterRepo.Get(runCtx, req.CharacterID)
		if err != nil {
			return err
		}
		if character == nil {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_CHARACTER_NOT_FOUND)
		}
		sessionID := uuid.NewString()
		oldSessionID, err := u.characterSessionRepo.Online(
			runCtx,
			req.CharacterID,
			sessionID,
			int64(u.onlineSessionTTL/time.Second),
		)
		if err != nil {
			return err
		}
		if oldSessionID != "" && oldSessionID != sessionID {
			err = u.eventUsecase.PublishCloseSession(runCtx, &model.CharacterCloseSessionEvent{
				SessionID:       oldSessionID,
				Reason:          enum.CharacterCloseSessionReasonOccupied,
				Message:         "Disconnected. The game was opened from another device or window.",
				ShouldReconnect: false,
			})
			if err != nil {
				return err
			}
		}
		session = &model.CharacterSession{
			CharacterID: req.CharacterID,
			SessionID:   sessionID,
			ExpiresIn:   u.onlineSessionTTL,
		}
		return nil
	})
	return session, err
}

type PingCharacterReq struct {
	CharacterID int64
	SessionID   string
}

func (u *CharacterUsecase) Ping(ctx context.Context, req *PingCharacterReq) (*model.CharacterSession, error) {
	var session *model.CharacterSession
	err := u.stateUsecase.Execute(ctx, req.CharacterID, func(runCtx context.Context) error {
		ok, err := u.characterSessionRepo.Ping(
			runCtx,
			req.CharacterID,
			req.SessionID,
			int64(u.onlineSessionTTL/time.Second),
		)
		if err != nil {
			return err
		}
		if !ok {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_CHARACTER_SESSION_INVALID)
		}
		session = &model.CharacterSession{
			CharacterID: req.CharacterID,
			SessionID:   req.SessionID,
			ExpiresIn:   u.onlineSessionTTL,
		}
		return nil
	})
	return session, err
}

type OfflineCharacterReq struct {
	CharacterID int64
	SessionID   string
	Timeout     bool
}

func (u *CharacterUsecase) Offline(ctx context.Context, req *OfflineCharacterReq) error {
	return u.stateUsecase.Execute(ctx, req.CharacterID, func(runCtx context.Context) error {
		offline, err := u.characterSessionRepo.Offline(runCtx, req.CharacterID, req.SessionID)
		if err != nil {
			return err
		}
		if !offline {
			online, err := u.characterSessionRepo.IsOnline(runCtx, req.CharacterID)
			if err != nil {
				return err
			}
			offline = !online
		}
		if offline {
			updated, err := u.characterRepo.UpdateLastOfflineAt(runCtx, req.CharacterID, time.Now())
			if err != nil {
				return err
			}
			if !updated {
				return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_CHARACTER_NOT_FOUND)
			}
		}
		if req.Timeout {
			return u.eventUsecase.PublishCloseSession(runCtx, &model.CharacterCloseSessionEvent{
				SessionID:       req.SessionID,
				Reason:          enum.CharacterCloseSessionReasonTimeout,
				ShouldReconnect: false,
			})
		}
		return nil
	})
}
