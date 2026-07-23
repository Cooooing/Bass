package service

import (
	"context"

	"common/pkg/apperror"
	"common/proto/gen/common"
	cerrors "common/proto/gen/common/errors"
	v1 "common/proto/gen/game_town/v1"
	"game_town/internal/biz/base"
	"game_town/internal/biz/model"
	"game_town/internal/biz/usecase"
	gameenum "game_town/internal/enum"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"github.com/samber/lo"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type WorldService struct {
	v1.UnimplementedGameTownWorldServiceServer
	usecase *usecase.WorldUsecase
}

func NewWorldService(usecase *usecase.WorldUsecase) *WorldService {
	return &WorldService{usecase: usecase}
}

func (s *WorldService) RegisterGrpc(server *grpc.Server) {
	v1.RegisterGameTownWorldServiceServer(server, s)
}

func (s *WorldService) RegisterHttp(*http.Server) {}

func (s *WorldService) Create(ctx context.Context, req *v1.CreateGameTownWorld_Request) (*v1.CreateGameTownWorld_Resp, error) {
	resp, err := s.usecase.Create(ctx, &usecase.CreateWorldReq{
		CreatorPlayerID: req.GetCreatorPlayerId(),
		Description:     req.GetDescription(),
		NpcCount:        req.GetNpcCount(),
		LocationCount:   req.GetLocationCount(),
		Seed:            req.Seed,
		AgentConfigID:   req.GetAgentConfigId(),
	})
	if err != nil {
		return nil, err
	}
	row := resp.World
	return &v1.CreateGameTownWorld_Resp{
		Row: &v1.CreateGameTownWorld_Resp_Row{
			Id:                row.ID,
			Code:              row.Code,
			Name:              row.Name,
			Description:       row.Description,
			Status:            gameenum.WorldStatusMap.MustToProto(row.Status),
			CreatorPlayerId:   row.CreatorPlayerID,
			DefaultLocationId: row.DefaultLocationID,
			Seed:              row.Seed,
			AgentConfigId:     row.AgentConfigID,
			GenerationSummary: row.GenerationSummary,
			CreatedAt:         timestamppb.New(*row.CreatedAt),
			UpdatedAt:         timestamppb.New(*row.UpdatedAt),
		},
		EventId: resp.Event.ID,
	}, nil
}

func (s *WorldService) Get(ctx context.Context, req *v1.GetGameTownWorld_Request) (*v1.GetGameTownWorld_Resp, error) {
	if req.GetId() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	row, err := s.usecase.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &v1.GetGameTownWorld_Resp{
		Row: &v1.GetGameTownWorld_Resp_Row{
			Id:                row.ID,
			Code:              row.Code,
			Name:              row.Name,
			Description:       row.Description,
			Status:            gameenum.WorldStatusMap.MustToProto(row.Status),
			CreatorPlayerId:   row.CreatorPlayerID,
			DefaultLocationId: row.DefaultLocationID,
			Seed:              row.Seed,
			AgentConfigId:     row.AgentConfigID,
			GenerationSummary: row.GenerationSummary,
			CreatedAt:         timestamppb.New(*row.CreatedAt),
			UpdatedAt:         timestamppb.New(*row.UpdatedAt),
		},
	}, nil
}

func (s *WorldService) Page(ctx context.Context, req *v1.PageGameTownWorlds_Request) (*v1.PageGameTownWorlds_Resp, error) {
	page := base.PageRequest{}
	if req.GetPage() != nil {
		page.Page = int64(req.GetPage().GetPage())
		page.Size = int64(req.GetPage().GetSize())
	}
	var status *gameenum.WorldStatus
	if req.Status != nil {
		value, ok := gameenum.WorldStatusMap.ToEnum(*req.Status)
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		status = new(value)
	}
	resp, err := s.usecase.Page(ctx, &usecase.PageWorldsReq{
		Page:            page,
		CreatorPlayerID: req.CreatorPlayerId,
		Status:          status,
	})
	if err != nil {
		return nil, err
	}
	reply := &v1.PageGameTownWorlds_Resp{
		Page: &common.PageResp{
			Page:  uint32(resp.Page.Page),
			Size:  uint32(resp.Page.Size),
			Total: uint32(resp.Page.Total),
		},
		Rows: lo.Map(resp.Rows, func(row *model.World, _ int) *v1.PageGameTownWorlds_Resp_Row {
			return &v1.PageGameTownWorlds_Resp_Row{
				Id:                row.ID,
				Code:              row.Code,
				Name:              row.Name,
				Description:       row.Description,
				Status:            gameenum.WorldStatusMap.MustToProto(row.Status),
				CreatorPlayerId:   row.CreatorPlayerID,
				DefaultLocationId: row.DefaultLocationID,
				Seed:              row.Seed,
				AgentConfigId:     row.AgentConfigID,
				GenerationSummary: row.GenerationSummary,
				CreatedAt:         timestamppb.New(*row.CreatedAt),
				UpdatedAt:         timestamppb.New(*row.UpdatedAt),
			}
		}),
	}
	return reply, nil
}
