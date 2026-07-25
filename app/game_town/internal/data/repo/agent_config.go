package repo

import (
	"context"
	"strings"

	"common/pkg/apperror"
	utilent "common/pkg/util/ent"
	cerrors "common/proto/gen/common/errors"
	"game_town/internal/biz/model"
	bizrepo "game_town/internal/biz/repo"
	"game_town/internal/data/gen"
	"game_town/internal/data/gen/agentconfig"
	"game_town/internal/enum"

	"github.com/samber/lo"
)

var _ bizrepo.AgentConfigRepo = (*AgentConfigRepo)(nil)

type AgentConfigRepo struct {
	db *gen.Client
}

func NewAgentConfigRepo(
	db *gen.Client,
) bizrepo.AgentConfigRepo {
	return &AgentConfigRepo{
		db: db,
	}
}

func (r *AgentConfigRepo) getClient(ctx context.Context) *gen.Client {
	if tx, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return tx
	}
	return r.db
}

func (r *AgentConfigRepo) Save(ctx context.Context, row *model.AgentConfig) (*model.AgentConfig, error) {
	saved, err := r.getClient(ctx).AgentConfig.Create().
		SetName(row.Name).
		SetProvider(agentconfig.Provider(row.Provider)).
		SetBaseURL(row.BaseURL).
		SetModel(row.Model).
		SetSecretEnv(row.SecretEnv).
		SetTimeoutSeconds(row.TimeoutSeconds).
		Save(ctx)
	if err != nil {
		if isAgentConfigNameConflict(err) {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		return nil, err
	}
	return agentConfigModel(saved), nil
}

func isAgentConfigNameConflict(err error) bool {
	message := err.Error()
	return strings.Contains(message, "game_town_agent_configs_name_active_unique") || strings.Contains(message, "duplicate key")
}

func agentConfigQuery(q *gen.AgentConfigQuery, req *bizrepo.AgentConfigQuery) *gen.AgentConfigQuery {
	q = q.Where(agentconfig.DeletedAtIsNil())
	if req == nil {
		return q
	}
	if req.ID != nil {
		q = q.Where(agentconfig.ID(*req.ID))
	}
	if len(req.IDs) > 0 {
		q = q.Where(agentconfig.IDIn(req.IDs...))
	}
	if req.Provider != nil {
		q = q.Where(agentconfig.ProviderEQ(agentconfig.Provider(*req.Provider)))
	}
	return q
}

func (r *AgentConfigRepo) Get(ctx context.Context, req *bizrepo.AgentConfigQuery) (*model.AgentConfig, error) {
	row, err := agentConfigQuery(r.getClient(ctx).AgentConfig.Query(), req).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_TOWN_AGENT_CONFIG_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return agentConfigModel(row), nil
}

func (r *AgentConfigRepo) List(ctx context.Context, req *bizrepo.AgentConfigQuery) ([]*model.AgentConfig, error) {
	rows, err := agentConfigQuery(r.getClient(ctx).AgentConfig.Query(), req).Order(agentconfig.ByID()).All(ctx)
	if err != nil {
		return nil, err
	}
	return lo.Map(rows, func(row *gen.AgentConfig, _ int) *model.AgentConfig {
		return agentConfigModel(row)
	}), nil
}

func (r *AgentConfigRepo) Map(ctx context.Context, req *bizrepo.AgentConfigQuery) (map[int64]*model.AgentConfig, error) {
	rows, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	out := make(map[int64]*model.AgentConfig, len(rows))
	for _, row := range rows {
		out[row.ID] = row
	}
	return out, nil
}

func (r *AgentConfigRepo) Count(ctx context.Context, req *bizrepo.AgentConfigQuery) (int, error) {
	return agentConfigQuery(r.getClient(ctx).AgentConfig.Query(), req).Count(ctx)
}

func (r *AgentConfigRepo) Page(ctx context.Context, req *bizrepo.AgentConfigPageReq) (*bizrepo.AgentConfigPageResp, error) {
	p := page(req.Page)
	q := agentConfigQuery(r.getClient(ctx).AgentConfig.Query(), &req.Query)
	total, err := q.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := q.Order(agentconfig.ByID()).Offset(pageOffset(p)).Limit(pageLimit(p)).All(ctx)
	if err != nil {
		return nil, err
	}
	out := lo.Map(rows, func(row *gen.AgentConfig, _ int) *model.AgentConfig {
		return agentConfigModel(row)
	})
	return &bizrepo.AgentConfigPageResp{
		Rows: out,
		Page: basePage(total, p),
	}, nil
}

func agentConfigModel(row *gen.AgentConfig) *model.AgentConfig {
	return &model.AgentConfig{
		ID:             row.ID,
		Name:           row.Name,
		Provider:       enum.AgentProvider(row.Provider),
		BaseURL:        row.BaseURL,
		Model:          row.Model,
		SecretEnv:      row.SecretEnv,
		TimeoutSeconds: row.TimeoutSeconds,
		CreatedAt:      row.CreatedAt,
		UpdatedAt:      row.UpdatedAt,
	}
}
