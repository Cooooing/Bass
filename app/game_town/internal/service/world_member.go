package service

import (
	"context"
	"time"

	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	v1 "common/proto/gen/game_town/v1"
	"game_town/internal/biz/usecase"
	"game_town/internal/enum"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type WorldMemberService struct {
	v1.UnimplementedGameTownWorldMemberServiceServer
	usecase *usecase.WorldMemberUsecase
}

func NewWorldMemberService(
	usecase *usecase.WorldMemberUsecase,
) *WorldMemberService {
	return &WorldMemberService{
		usecase: usecase,
	}
}

func (s *WorldMemberService) RegisterGrpc(server *grpc.Server) {
	v1.RegisterGameTownWorldMemberServiceServer(server, s)
}

func (s *WorldMemberService) RegisterHttp(*http.Server) {
}

func (s *WorldMemberService) Join(ctx context.Context, req *v1.JoinGameTownWorld_Request) (*v1.JoinGameTownWorld_Resp, error) {
	if req.GetPlayerId() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}

	resp, err := s.usecase.Join(ctx, &usecase.JoinWorldReq{
		PlayerID:            req.GetPlayerId(),
		WorldCode:           req.GetWorldCode(),
		CharacterPreference: req.GetCharacterPreference(),
	})
	if err != nil {
		return nil, err
	}

	return &v1.JoinGameTownWorld_Resp{
		WorldId:    resp.WorldID,
		MemberId:   resp.MemberID,
		LocationId: resp.LocationID,
		EventId:    resp.EventID,
	}, nil
}

func (s *WorldMemberService) Get(ctx context.Context, req *v1.GetGameTownWorldMember_Request) (*v1.GetGameTownWorldMember_Resp, error) {
	if req.GetWorldId() <= 0 || req.GetPlayerId() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}

	resp, err := s.usecase.Get(ctx, &usecase.GetWorldMemberReq{
		WorldID:  req.GetWorldId(),
		PlayerID: req.GetPlayerId(),
	})
	if err != nil {
		return nil, err
	}

	timeScale := resp.State.TimeScale
	if timeScale == 0 {
		timeScale = 24
	}
	worldTime := resp.State.WorldTime.Add(
		time.Duration(float64(time.Since(resp.State.TimeAnchor)) * float64(timeScale)),
	)

	return &v1.GetGameTownWorldMember_Resp{
		Id:                  resp.Member.ID,
		WorldId:             resp.Member.WorldID,
		PlayerId:            resp.Member.PlayerID,
		CurrentLocationId:   resp.Member.CurrentLocationID,
		Role:                enum.WorldMemberRoleMap.MustToProto(resp.Member.Role),
		WorldTime:           timestamppb.New(worldTime),
		JoinedAt:            timestamppb.New(resp.Member.JoinedAt),
		CharacterName:       resp.Member.CharacterName,
		CharacterBackground: resp.Member.CharacterBackground,
		CharacterGoal:       resp.Member.CharacterGoal,
		CharacterTraits:     resp.Member.CharacterTraits,
		CharacterReady:      resp.Member.CharacterReady,
		CharacterPreference: resp.Member.CharacterPreference,
	}, nil
}

func (s *WorldMemberService) SubmitAction(ctx context.Context, req *v1.SubmitGameTownAction_Request) (*v1.SubmitGameTownAction_Resp, error) {
	if req.GetWorldId() <= 0 || req.GetPlayerId() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}

	targets := make([]usecase.SubmitActionTarget, 0, len(req.GetTargets()))
	for _, target := range req.GetTargets() {
		entityType, ok := enum.EntityTypeMap.ToEnum(target.GetType())
		if !ok || target.GetId() <= 0 {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		targets = append(targets, usecase.SubmitActionTarget{
			Type: entityType,
			ID:   target.GetId(),
		})
	}

	event, err := s.usecase.SubmitAction(ctx, &usecase.SubmitActionReq{
		WorldID:         req.GetWorldId(),
		PlayerID:        req.GetPlayerId(),
		Content:         req.GetContent(),
		Targets:         targets,
		ClientRequestID: req.GetClientRequestId(),
	})
	if err != nil {
		return nil, err
	}

	return &v1.SubmitGameTownAction_Resp{
		EventId: event.ID,
	}, nil
}
