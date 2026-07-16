package repo

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	"context"
	"game_town/internal/biz/model"
	bizrepo "game_town/internal/biz/repo"
	"game_town/internal/data/gen"
	"game_town/internal/data/gen/agentconfig"
)

type AgentConfigRepo struct{ *baseRepo }

func NewAgentConfigRepo(db *gen.Client) bizrepo.AgentConfigRepo {
	return &AgentConfigRepo{baseRepo: &baseRepo{db: db}}
}

func (r *AgentConfigRepo) CreateAgentConfig(ctx context.Context, req *bizrepo.CreateAgentConfigReq) (*bizrepo.CreateAgentConfigResponse, error) {
	row := req.Row
	if row.IsDefault {
		_, err := r.db.AgentConfig.Update().Where(agentconfig.PlayerID(row.PlayerID), agentconfig.DeletedAtIsNil()).SetIsDefault(false).Save(ctx)
		if err != nil {
			return nil, err
		}
	}
	created, err := r.db.AgentConfig.Create().SetPlayerID(row.PlayerID).SetName(row.Name).SetProvider(row.Provider).SetModel(row.Model).SetBaseURL(row.BaseURL).SetAPIKey(row.APIKey).SetTimeoutSeconds(row.TimeoutSeconds).SetIsDefault(row.IsDefault).SetStatus(row.Status).Save(ctx)
	if gen.IsConstraintError(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if err != nil {
		return nil, err
	}
	return &bizrepo.CreateAgentConfigResponse{Row: r.agentConfig(created)}, nil
}

func (r *AgentConfigRepo) GetAgentConfig(ctx context.Context, req *bizrepo.GetAgentConfigReq) (*bizrepo.GetAgentConfigResponse, error) {
	row, err := r.db.AgentConfig.Query().Where(agentconfig.ID(req.ID), agentconfig.PlayerID(req.PlayerID), agentconfig.DeletedAtIsNil()).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_TOWN_AGENT_CONFIG_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return &bizrepo.GetAgentConfigResponse{Row: r.agentConfig(row)}, nil
}

func (r *AgentConfigRepo) GetDefaultAgentConfig(ctx context.Context, req *bizrepo.GetDefaultAgentConfigReq) (*bizrepo.GetDefaultAgentConfigResponse, error) {
	row, err := r.db.AgentConfig.Query().Where(agentconfig.PlayerID(req.PlayerID), agentconfig.IsDefault(true), agentconfig.Status("active"), agentconfig.DeletedAtIsNil()).Order(agentconfig.ByID()).First(ctx)
	if gen.IsNotFound(err) {
		row, err = r.db.AgentConfig.Query().Where(agentconfig.PlayerID(req.PlayerID), agentconfig.Status("active"), agentconfig.DeletedAtIsNil()).Order(agentconfig.ByID()).First(ctx)
	}
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_TOWN_AGENT_CONFIG_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return &bizrepo.GetDefaultAgentConfigResponse{Row: r.agentConfig(row)}, nil
}

func (r *AgentConfigRepo) ListAgentConfigs(ctx context.Context, req *bizrepo.ListAgentConfigsReq) (*bizrepo.ListAgentConfigsResponse, error) {
	query := r.db.AgentConfig.Query().Where(agentconfig.PlayerID(req.PlayerID), agentconfig.DeletedAtIsNil())
	if req.Status != nil {
		query = query.Where(agentconfig.Status(*req.Status))
	}
	rows, err := query.Order(agentconfig.ByID()).All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.AgentConfig, 0, len(rows))
	for _, row := range rows {
		result = append(result, r.agentConfig(row))
	}
	return &bizrepo.ListAgentConfigsResponse{Rows: result}, nil
}
