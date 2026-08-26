package usecase

import (
	"context"
	"game_idle/internal/biz/model"
	"game_idle/internal/biz/repo"
	"game_idle/internal/config"
	"game_idle/internal/enum"
	"regexp"
	"strings"
	"time"
)

var characterNicknamePattern = regexp.MustCompile("^[A-Za-z0-9_]{4,32}$")

type CharacterUsecase struct {
	characterRepo              repo.CharacterRepo
	maxCharacterCountPerUser   int32
	defaultActionQueueCapacity int32
	defaultMaxOfflineDuration  time.Duration
}

func NewCharacterUsecase(conf *config.Bootstrap, characterRepo repo.CharacterRepo) *CharacterUsecase {
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
	return &CharacterUsecase{
		characterRepo:              characterRepo,
		maxCharacterCountPerUser:   maxCharacterCountPerUser,
		defaultActionQueueCapacity: defaultActionQueueCapacity,
		defaultMaxOfflineDuration:  defaultMaxOfflineDuration,
	}
}

type CreateCharacterReq struct {
	UserID   int64
	Nickname string
}

func (u *CharacterUsecase) Create(ctx context.Context, req *CreateCharacterReq) (*model.Character, error) {
	if req.UserID <= 0 || !characterNicknamePattern.MatchString(req.Nickname) {
		return nil, model.ErrCharacterInvalid
	}
	rows, err := u.characterRepo.ListByUserID(ctx, req.UserID)
	if err != nil {
		return nil, err
	}
	if len(rows) >= int(u.maxCharacterCountPerUser) {
		return nil, model.ErrCharacterLimitExceeded
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
		return nil, model.ErrCharacterLimitExceeded
	}
	return u.characterRepo.Save(ctx, &model.Character{
		UserID:              req.UserID,
		Slot:                slot,
		Nickname:            req.Nickname,
		NicknameKey:         strings.ToLower(req.Nickname),
		ActionQueueCapacity: u.defaultActionQueueCapacity,
		MaxOfflineDuration:  u.defaultMaxOfflineDuration,
		Status:              enum.CharacterStatusActive,
	})
}

func (u *CharacterUsecase) Get(ctx context.Context, userID int64) ([]*model.Character, error) {
	if userID <= 0 {
		return nil, model.ErrCharacterInvalid
	}
	return u.characterRepo.ListByUserID(ctx, userID)
}
