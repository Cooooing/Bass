package service

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	v1 "common/proto/gen/game_idle/v1"
	"context"
	"game_idle/internal/biz/usecase"
	"game_idle/internal/enum"
	"sort"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

type CharacterAbilityService struct {
	v1.UnimplementedCharacterAbilityServiceServer
	characterAbilityUsecase *usecase.CharacterAbilityUsecase
}

func NewCharacterAbilityService(characterAbilityUsecase *usecase.CharacterAbilityUsecase) *CharacterAbilityService {
	return &CharacterAbilityService{
		characterAbilityUsecase: characterAbilityUsecase,
	}
}

func (s *CharacterAbilityService) RegisterGrpc(server *grpc.Server) {
	v1.RegisterCharacterAbilityServiceServer(server, s)
}

func (s *CharacterAbilityService) RegisterHttp(*http.Server) {
}

func (s *CharacterAbilityService) Get(
	ctx context.Context,
	req *v1.GetCharacterAbility_Request,
) (*v1.GetCharacterAbility_Resp, error) {
	if req.GetCharacterId() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	rows, err := s.characterAbilityUsecase.Map(ctx, req.GetCharacterId())
	if err != nil {
		return nil, err
	}
	abilityIDs := make([]string, 0, len(rows))
	for abilityID := range rows {
		abilityIDs = append(abilityIDs, abilityID.String())
	}
	sort.Strings(abilityIDs)
	abilities := make([]*v1.CharacterAbility, 0, len(abilityIDs))
	for _, abilityID := range abilityIDs {
		row := rows[enum.Ability(abilityID)]
		abilities = append(abilities, &v1.CharacterAbility{
			AbilityId:    row.AbilityID.String(),
			Level:        row.Level,
			Exp:          row.Exp,
			NextLevelExp: row.NextLevelExp,
		})
	}
	return &v1.GetCharacterAbility_Resp{Rows: abilities}, nil
}
