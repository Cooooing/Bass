package service

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	v1 "common/proto/gen/game_idle/v1"
	"context"
	"game_idle/internal/biz/usecase"
	"game_idle/internal/enum"
	"time"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CharacterService struct {
	v1.UnimplementedGameIdleCharacterServiceServer
	characterUsecase *usecase.CharacterUsecase
}

func NewCharacterService(characterUsecase *usecase.CharacterUsecase) *CharacterService {
	return &CharacterService{
		characterUsecase: characterUsecase,
	}
}

func (s *CharacterService) RegisterGrpc(server *grpc.Server) {
	v1.RegisterGameIdleCharacterServiceServer(server, s)
}

func (s *CharacterService) RegisterHttp(*http.Server) {
}

func (s *CharacterService) Create(ctx context.Context, req *v1.CreateGameIdleCharacter_Request) (*v1.CreateGameIdleCharacter_Resp, error) {
	if req.GetUserId() <= 0 || req.GetName() == "" {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	row, err := s.characterUsecase.Create(ctx, &usecase.CreateCharacterReq{
		UserID: req.GetUserId(),
		Name:   req.GetName(),
	})
	if err != nil {
		return nil, err
	}
	character := &v1.Character{
		Id:                  row.ID,
		UserId:              row.UserID,
		Name:                row.Name,
		Status:              enum.CharacterStatusMap.MustToProto(row.Status),
		Slot:                row.Slot,
		ActionQueueCapacity: row.ActionQueueCapacity,
		MaxOfflineSeconds:   int64(row.MaxOfflineDuration / time.Second),
	}
	if row.CreatedAt != nil {
		character.CreatedAt = timestamppb.New(*row.CreatedAt)
	}
	if row.UpdatedAt != nil {
		character.UpdatedAt = timestamppb.New(*row.UpdatedAt)
	}
	return &v1.CreateGameIdleCharacter_Resp{Row: character}, nil
}

func (s *CharacterService) Get(ctx context.Context, req *v1.GetGameIdleCharacter_Request) (*v1.GetGameIdleCharacter_Resp, error) {
	if req.GetUserId() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	rows, err := s.characterUsecase.Get(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}
	characters := make([]*v1.Character, 0, len(rows))
	for _, row := range rows {
		character := &v1.Character{
			Id:                  row.ID,
			UserId:              row.UserID,
			Name:                row.Name,
			Status:              enum.CharacterStatusMap.MustToProto(row.Status),
			Slot:                row.Slot,
			ActionQueueCapacity: row.ActionQueueCapacity,
			MaxOfflineSeconds:   int64(row.MaxOfflineDuration / time.Second),
		}
		if row.CreatedAt != nil {
			character.CreatedAt = timestamppb.New(*row.CreatedAt)
		}
		if row.UpdatedAt != nil {
			character.UpdatedAt = timestamppb.New(*row.UpdatedAt)
		}
		characters = append(characters, character)
	}
	return &v1.GetGameIdleCharacter_Resp{Rows: characters}, nil
}
