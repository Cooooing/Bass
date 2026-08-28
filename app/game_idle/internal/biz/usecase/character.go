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

var characterNamePattern = regexp.MustCompile("^[A-Za-z0-9_]{4,32}$")

type CharacterUsecase struct {
	characterRepo              repo.CharacterRepo
	characterSessionRepo       repo.CharacterSessionRepo
	gameIdleEventRepo          repo.GameIdleEventRepo
	maxCharacterCountPerUser   int32
	defaultActionQueueCapacity int32
	defaultMaxOfflineDuration  time.Duration
	webSocketOnlineTTL         time.Duration
}

func NewCharacterUsecase(
	conf *config.Bootstrap,
	characterRepo repo.CharacterRepo,
	characterSessionRepo repo.CharacterSessionRepo,
	gameIdleEventRepo repo.GameIdleEventRepo,
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
	webSocketOnlineTTL := 5 * time.Minute
	if conf.GetGameIdle().GetWebsocket().GetOnlineTtl() != nil && conf.GetGameIdle().GetWebsocket().GetOnlineTtl().AsDuration() > 0 {
		webSocketOnlineTTL = conf.GetGameIdle().GetWebsocket().GetOnlineTtl().AsDuration()
	}
	return &CharacterUsecase{
		characterRepo:              characterRepo,
		characterSessionRepo:       characterSessionRepo,
		gameIdleEventRepo:          gameIdleEventRepo,
		maxCharacterCountPerUser:   maxCharacterCountPerUser,
		defaultActionQueueCapacity: defaultActionQueueCapacity,
		defaultMaxOfflineDuration:  defaultMaxOfflineDuration,
		webSocketOnlineTTL:         webSocketOnlineTTL,
	}
}

type CreateCharacterReq struct {
	UserID int64
	Name   string
}

func (u *CharacterUsecase) Create(ctx context.Context, req *CreateCharacterReq) (*model.Character, error) {
	if req.UserID <= 0 || !characterNamePattern.MatchString(req.Name) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_CHARACTER_INVALID)
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
	return u.characterRepo.Save(ctx, &model.Character{
		UserID:              req.UserID,
		Slot:                slot,
		Name:                req.Name,
		NameKey:             strings.ToLower(req.Name),
		ActionQueueCapacity: u.defaultActionQueueCapacity,
		MaxOfflineDuration:  u.defaultMaxOfflineDuration,
		Status:              enum.CharacterStatusActive,
	})
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
	UserID      int64
	CharacterID int64
}

func (u *CharacterUsecase) Online(ctx context.Context, req *OnlineCharacterReq) (*model.CharacterSession, error) {
	character, err := u.characterRepo.Get(ctx, req.CharacterID)
	if err != nil {
		return nil, err
	}
	if character.UserID != req.UserID {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_CHARACTER_INVALID)
	}
	sessionID := uuid.NewString()
	oldSessionID, err := u.characterSessionRepo.Online(
		ctx,
		req.CharacterID,
		sessionID,
		int64(u.webSocketOnlineTTL/time.Second),
	)
	if err != nil {
		return nil, err
	}
	if oldSessionID != "" && oldSessionID != sessionID {
		err = u.gameIdleEventRepo.Publish(ctx, &model.GameIdleEvent{
			CloseSession: &model.CharacterCloseSessionEvent{
				SessionID:       oldSessionID,
				Reason:          enum.CharacterCloseSessionReasonOccupied,
				Message:         "Disconnected. The game was opened from another device or window.",
				ShouldReconnect: false,
			},
		})
		if err != nil {
			return nil, err
		}
	}
	return &model.CharacterSession{
		CharacterID: req.CharacterID,
		SessionID:   sessionID,
		ExpiresIn:   u.webSocketOnlineTTL,
	}, nil
}

type PingCharacterReq struct {
	CharacterID int64
	SessionID   string
}

func (u *CharacterUsecase) Ping(ctx context.Context, req *PingCharacterReq) (*model.CharacterSession, error) {
	ok, err := u.characterSessionRepo.Ping(
		ctx,
		req.CharacterID,
		req.SessionID,
		int64(u.webSocketOnlineTTL/time.Second),
	)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_CHARACTER_SESSION_INVALID)
	}
	return &model.CharacterSession{
		CharacterID: req.CharacterID,
		SessionID:   req.SessionID,
		ExpiresIn:   u.webSocketOnlineTTL,
	}, nil
}

type OfflineCharacterReq struct {
	CharacterID int64
	SessionID   string
	Timeout     bool
}

func (u *CharacterUsecase) Offline(ctx context.Context, req *OfflineCharacterReq) error {
	offline, err := u.characterSessionRepo.Offline(ctx, req.CharacterID, req.SessionID)
	if err != nil {
		return err
	}
	if offline {
		if err = u.characterRepo.UpdateLastOfflineAt(ctx, req.CharacterID, time.Now()); err != nil {
			return err
		}
	}
	if req.Timeout {
		return u.gameIdleEventRepo.Publish(ctx, &model.GameIdleEvent{
			CloseSession: &model.CharacterCloseSessionEvent{
				SessionID:       req.SessionID,
				Reason:          enum.CharacterCloseSessionReasonTimeout,
				ShouldReconnect: false,
			},
		})
	}
	return nil
}
