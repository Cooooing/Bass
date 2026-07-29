package service

import (
	"context"
	"time"

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

type AgentConfigService struct {
	v1.UnimplementedGameTownAgentConfigServiceServer
	usecase *usecase.AgentConfigUsecase
}

func NewAgentConfigService(
	usecase *usecase.AgentConfigUsecase,
) *AgentConfigService {
	return &AgentConfigService{
		usecase: usecase,
	}
}

func (s *AgentConfigService) RegisterGrpc(server *grpc.Server) {
	v1.RegisterGameTownAgentConfigServiceServer(server, s)
}

func (s *AgentConfigService) RegisterHttp(*http.Server) {
}

func (s *AgentConfigService) Create(ctx context.Context, req *v1.CreateGameTownAgentConfig_Request) (*v1.CreateGameTownAgentConfig_Resp, error) {
	provider, ok := gameenum.AgentProviderMap.ToEnum(req.GetProvider())
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	row, err := s.usecase.Create(ctx, &usecase.CreateAgentConfigReq{
		Name:      req.GetName(),
		Provider:  provider,
		BaseURL:   req.GetBaseUrl(),
		Model:     req.GetModel(),
		SecretEnv: req.GetSecretEnv(),
		Timeout:   time.Duration(req.GetTimeoutSeconds()) * time.Second,
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateGameTownAgentConfig_Resp{
		Row: &v1.CreateGameTownAgentConfig_Resp_Row{
			Id:             row.ID,
			Name:           row.Name,
			Provider:       gameenum.AgentProviderMap.MustToProto(row.Provider),
			BaseUrl:        row.BaseURL,
			Model:          row.Model,
			SecretEnv:      row.SecretEnv,
			TimeoutSeconds: int32(row.Timeout / time.Second),
			CreatedAt:      timestamppb.New(*row.CreatedAt),
			UpdatedAt:      timestamppb.New(*row.UpdatedAt),
		},
	}, nil
}

func (s *AgentConfigService) Get(ctx context.Context, req *v1.GetGameTownAgentConfig_Request) (*v1.GetGameTownAgentConfig_Resp, error) {
	if req.GetId() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	row, err := s.usecase.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &v1.GetGameTownAgentConfig_Resp{
		Row: &v1.GetGameTownAgentConfig_Resp_Row{
			Id:             row.ID,
			Name:           row.Name,
			Provider:       gameenum.AgentProviderMap.MustToProto(row.Provider),
			BaseUrl:        row.BaseURL,
			Model:          row.Model,
			SecretEnv:      row.SecretEnv,
			TimeoutSeconds: int32(row.Timeout / time.Second),
			CreatedAt:      timestamppb.New(*row.CreatedAt),
			UpdatedAt:      timestamppb.New(*row.UpdatedAt),
		},
	}, nil
}

func (s *AgentConfigService) List(ctx context.Context, req *v1.ListGameTownAgentConfigs_Request) (*v1.ListGameTownAgentConfigs_Resp, error) {
	page := base.PageRequest{}
	if req.GetPage() != nil {
		page.Page = int64(req.GetPage().GetPage())
		page.Size = int64(req.GetPage().GetSize())
	}
	resp, err := s.usecase.Page(ctx, page)
	if err != nil {
		return nil, err
	}
	reply := &v1.ListGameTownAgentConfigs_Resp{
		Page: &common.PageResp{
			Page:  uint32(resp.Page.Page),
			Size:  uint32(resp.Page.Size),
			Total: uint32(resp.Page.Total),
		},
		Rows: lo.Map(resp.Rows, func(row *model.AgentConfig, _ int) *v1.ListGameTownAgentConfigs_Resp_Row {
			return &v1.ListGameTownAgentConfigs_Resp_Row{
				Id:             row.ID,
				Name:           row.Name,
				Provider:       gameenum.AgentProviderMap.MustToProto(row.Provider),
				BaseUrl:        row.BaseURL,
				Model:          row.Model,
				SecretEnv:      row.SecretEnv,
				TimeoutSeconds: int32(row.Timeout / time.Second),
				CreatedAt:      timestamppb.New(*row.CreatedAt),
				UpdatedAt:      timestamppb.New(*row.UpdatedAt),
			}
		}),
	}
	return reply, nil
}
