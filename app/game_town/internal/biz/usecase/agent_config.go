package usecase

import (
	"context"
	"strings"

	"common/pkg/apperror"
	"game_town/internal/biz/base"
	"game_town/internal/biz/model"
	"game_town/internal/biz/repo"
	"game_town/internal/enum"
)

type AgentConfigUsecase struct {
	agentConfigRepo repo.AgentConfigRepo
}

func NewAgentConfigUsecase(
	agentConfigRepo repo.AgentConfigRepo,
) *AgentConfigUsecase {
	return &AgentConfigUsecase{
		agentConfigRepo: agentConfigRepo,
	}
}

type CreateAgentConfigReq struct {
	Name           string
	Provider       enum.AgentProvider
	BaseURL        string
	Model          string
	SecretEnv      string
	TimeoutSeconds int32
}

func (u *AgentConfigUsecase) Create(
	ctx context.Context,
	req *CreateAgentConfigReq,
) (*model.AgentConfig, error) {
	name := strings.TrimSpace(req.Name)
	baseURL := strings.TrimSpace(req.BaseURL)
	modelName := strings.TrimSpace(req.Model)
	if name == "" || baseURL == "" || modelName == "" {
		return nil, apperror.CommonInvalidArgument()
	}
	if _, ok := enum.AgentProviderMap.ToProto(req.Provider); !ok {
		return nil, apperror.CommonInvalidArgument()
	}
	timeoutSeconds := req.TimeoutSeconds
	if timeoutSeconds <= 0 {
		timeoutSeconds = 60
	}
	return u.agentConfigRepo.Save(ctx, &model.AgentConfig{
		Name:           name,
		Provider:       req.Provider,
		BaseURL:        baseURL,
		Model:          modelName,
		SecretEnv:      strings.TrimSpace(req.SecretEnv),
		TimeoutSeconds: timeoutSeconds,
	})
}

func (u *AgentConfigUsecase) Get(
	ctx context.Context,
	agentConfigID int64,
) (*model.AgentConfig, error) {
	return u.agentConfigRepo.Get(ctx, &repo.AgentConfigQuery{
		ID: new(agentConfigID),
	})
}

type PageAgentConfigsResp struct {
	Rows []*model.AgentConfig
	Page base.PageResp
}

func (u *AgentConfigUsecase) Page(
	ctx context.Context,
	page base.PageRequest,
) (*PageAgentConfigsResp, error) {
	resp, err := u.agentConfigRepo.Page(ctx, &repo.AgentConfigPageReq{
		Page: page,
	})
	if err != nil {
		return nil, err
	}
	return &PageAgentConfigsResp{
		Rows: resp.Rows,
		Page: resp.Page,
	}, nil
}
