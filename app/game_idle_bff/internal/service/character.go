package service

import (
	"common/pkg/apperror"
	"common/pkg/constant"
	commonmodel "common/pkg/model"
	"common/pkg/util"
	cerrors "common/proto/gen/common/errors"
	v1 "common/proto/gen/game_idle_bff/v1"
	"context"
	"game_idle_bff/internal/biz/usecase"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CharacterService struct {
	v1.UnimplementedCharacterServiceServer
	characterUsecase *usecase.CharacterUsecase
}

func NewCharacterService(
	characterUsecase *usecase.CharacterUsecase,
) *CharacterService {
	return &CharacterService{
		characterUsecase: characterUsecase,
	}
}

func (s *CharacterService) RegisterGrpc(*grpc.Server) {
}

func (s *CharacterService) RegisterHttp(hs *http.Server) {
	v1.RegisterCharacterServiceHTTPServer(hs, s)
}

func (s *CharacterService) Create(ctx context.Context, req *v1.CreateCharacter_Req) (*v1.CreateCharacter_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	row, err := s.characterUsecase.Create(ctx, &usecase.CreateCharacterReq{
		UserID: user.ID,
		Name:   req.GetName(),
	})
	if err != nil {
		return nil, err
	}
	character := &v1.Character{
		Id:                  row.ID,
		Name:                row.Name,
		Status:              row.Status,
		Slot:                row.Slot,
		ActionQueueCapacity: row.ActionQueueCapacity,
		MaxOfflineSeconds:   row.MaxOfflineSeconds,
	}
	if row.CreatedAt != nil {
		character.CreatedAt = timestamppb.New(*row.CreatedAt)
	}
	if row.UpdatedAt != nil {
		character.UpdatedAt = timestamppb.New(*row.UpdatedAt)
	}
	return &v1.CreateCharacter_Resp{Row: character}, nil
}

func (s *CharacterService) List(ctx context.Context, req *v1.ListCharacter_Req) (*v1.ListCharacter_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	rows, err := s.characterUsecase.List(ctx, &usecase.ListCharacterReq{
		UserID: user.ID,
	})
	if err != nil {
		return nil, err
	}
	characters := make([]*v1.Character, 0, len(rows))
	for _, row := range rows {
		character := &v1.Character{
			Id:                  row.ID,
			Name:                row.Name,
			Status:              row.Status,
			Slot:                row.Slot,
			ActionQueueCapacity: row.ActionQueueCapacity,
			MaxOfflineSeconds:   row.MaxOfflineSeconds,
		}
		if row.CreatedAt != nil {
			character.CreatedAt = timestamppb.New(*row.CreatedAt)
		}
		if row.UpdatedAt != nil {
			character.UpdatedAt = timestamppb.New(*row.UpdatedAt)
		}
		characters = append(characters, character)
	}
	return &v1.ListCharacter_Resp{
		Rows: characters,
	}, nil
}
