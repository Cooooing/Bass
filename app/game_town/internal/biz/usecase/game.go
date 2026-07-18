package usecase

import (
	"common/pkg/apperror"
	"common/proto/gen/common"
	cerrors "common/proto/gen/common/errors"
	"context"
	"fmt"
	"game_town/internal/biz/agent"
	"game_town/internal/biz/model"
	"game_town/internal/biz/repo"
	"game_town/internal/config"
	"strconv"
	"strings"
	"time"
)

type GameUsecase struct {
	conf                      *config.Bootstrap
	playerRepo                repo.PlayerRepo
	agentConfigRepo           repo.AgentConfigRepo
	sessionRepo               repo.SessionRepo
	worldRepo                 repo.WorldRepo
	worldMemberRepo           repo.WorldMemberRepo
	locationRepo              repo.LocationRepo
	npcRepo                   repo.NpcRepo
	commandRepo               repo.CommandRepo
	eventRepo                 repo.EventRepo
	worldStateSnapshotRepo    repo.WorldStateSnapshotRepo
	worldMetricDefinitionRepo repo.WorldMetricDefinitionRepo
	memoryRepo                repo.MemoryRepo
	relationshipRepo          repo.RelationshipRepo
	agentRunRepo              repo.AgentRunRepo
	agent                     agent.Runner
}

type CommandResult struct {
	Command         *model.Command
	FeedbackLines   []string
	CurrentWorld    *model.World
	CurrentLocation *model.Location
	VisibleNpcs     []*model.Npc
	WorldState      *model.WorldStateSnapshot
	Events          []*model.Event
}

func NewGameUsecase(conf *config.Bootstrap, playerRepo repo.PlayerRepo, agentConfigRepo repo.AgentConfigRepo, sessionRepo repo.SessionRepo, worldRepo repo.WorldRepo, worldMemberRepo repo.WorldMemberRepo, locationRepo repo.LocationRepo, npcRepo repo.NpcRepo, commandRepo repo.CommandRepo, eventRepo repo.EventRepo, worldStateSnapshotRepo repo.WorldStateSnapshotRepo, worldMetricDefinitionRepo repo.WorldMetricDefinitionRepo, memoryRepo repo.MemoryRepo, relationshipRepo repo.RelationshipRepo, agentRunRepo repo.AgentRunRepo, runner agent.Runner) *GameUsecase {
	return &GameUsecase{conf: conf, playerRepo: playerRepo, agentConfigRepo: agentConfigRepo, sessionRepo: sessionRepo, worldRepo: worldRepo, worldMemberRepo: worldMemberRepo, locationRepo: locationRepo, npcRepo: npcRepo, commandRepo: commandRepo, eventRepo: eventRepo, worldStateSnapshotRepo: worldStateSnapshotRepo, worldMetricDefinitionRepo: worldMetricDefinitionRepo, memoryRepo: memoryRepo, relationshipRepo: relationshipRepo, agentRunRepo: agentRunRepo, agent: runner}
}

type RegisterPlayerReq struct {
	Name        string
	DisplayName string
}

func (u *GameUsecase) RegisterPlayer(ctx context.Context, req *RegisterPlayerReq) (*model.Player, error) {
	if req == nil {
		req = &RegisterPlayerReq{}
	}
	return u.registerPlayer(ctx, &registerPlayerReq{Name: req.Name, DisplayName: req.DisplayName})
}

type registerPlayerReq struct {
	Name        string
	DisplayName string
}

func (u *GameUsecase) registerPlayer(ctx context.Context, req *registerPlayerReq) (*model.Player, error) {
	if req == nil {
		req = &registerPlayerReq{}
	}
	name := strings.ToLower(strings.TrimSpace(req.Name))
	displayName := strings.TrimSpace(req.DisplayName)
	if name == "" || len(name) > 64 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if displayName == "" {
		displayName = name
	}
	resp, err := u.playerRepo.CreatePlayer(ctx, &repo.CreatePlayerReq{Name: name, DisplayName: displayName})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type CreateAgentConfigReq struct {
	PlayerID       int64
	Name           string
	Provider       string
	ModelName      string
	BaseURL        string
	APIKey         string
	TimeoutSeconds int32
	IsDefault      bool
}

func (u *GameUsecase) CreateAgentConfig(ctx context.Context, req *CreateAgentConfigReq) (*model.AgentConfig, error) {
	if req == nil {
		req = &CreateAgentConfigReq{}
	}
	return u.createAgentConfig(ctx, &createAgentConfigReq{
		PlayerID:       req.PlayerID,
		Name:           req.Name,
		Provider:       req.Provider,
		ModelName:      req.ModelName,
		BaseURL:        req.BaseURL,
		APIKey:         req.APIKey,
		TimeoutSeconds: req.TimeoutSeconds,
		IsDefault:      req.IsDefault,
	})
}

type createAgentConfigReq struct {
	PlayerID       int64
	Name           string
	Provider       string
	ModelName      string
	BaseURL        string
	APIKey         string
	TimeoutSeconds int32
	IsDefault      bool
}

func (u *GameUsecase) createAgentConfig(ctx context.Context, req *createAgentConfigReq) (*model.AgentConfig, error) {
	if req == nil {
		req = &createAgentConfigReq{}
	}
	name := strings.TrimSpace(req.Name)
	provider := strings.TrimSpace(req.Provider)
	modelName := strings.TrimSpace(req.ModelName)
	baseURL := strings.TrimSpace(req.BaseURL)
	apiKey := strings.TrimSpace(req.APIKey)
	if req.PlayerID <= 0 || name == "" || provider == "" || modelName == "" {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	timeoutSeconds := req.TimeoutSeconds
	if timeoutSeconds <= 0 {
		timeoutSeconds = 30
	}
	isDefault := req.IsDefault
	listResp, err := u.agentConfigRepo.ListAgentConfigs(ctx, &repo.ListAgentConfigsReq{PlayerID: req.PlayerID})
	if err != nil {
		return nil, err
	}
	if len(listResp) == 0 {
		isDefault = true
	}
	createResp, err := u.agentConfigRepo.CreateAgentConfig(ctx, &model.AgentConfig{PlayerID: req.PlayerID, Name: name, Provider: provider, Model: modelName, BaseURL: baseURL, APIKey: apiKey, TimeoutSeconds: timeoutSeconds, IsDefault: isDefault, Status: "active"})
	if err != nil {
		return nil, err
	}
	return createResp, nil
}

type GetAgentConfigReq struct {
	ID       int64
	PlayerID int64
}

func (u *GameUsecase) GetAgentConfig(ctx context.Context, req *GetAgentConfigReq) (*model.AgentConfig, error) {
	if req == nil {
		req = &GetAgentConfigReq{}
	}
	return u.getAgentConfig(ctx, &getAgentConfigReq{ID: req.ID, PlayerID: req.PlayerID})
}

type getAgentConfigReq struct {
	ID       int64
	PlayerID int64
}

func (u *GameUsecase) getAgentConfig(ctx context.Context, req *getAgentConfigReq) (*model.AgentConfig, error) {
	if req == nil {
		req = &getAgentConfigReq{}
	}
	if req.ID <= 0 || req.PlayerID <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	resp, err := u.agentConfigRepo.GetAgentConfig(ctx, &repo.GetAgentConfigReq{ID: req.ID, PlayerID: req.PlayerID})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ListAgentConfigsReq struct {
	PlayerID int64
	Status   *string
}

func (u *GameUsecase) ListAgentConfigs(ctx context.Context, req *ListAgentConfigsReq) ([]*model.AgentConfig, error) {
	if req == nil {
		req = &ListAgentConfigsReq{}
	}
	return u.listAgentConfigs(ctx, &listAgentConfigsReq{PlayerID: req.PlayerID, Status: req.Status})
}

type listAgentConfigsReq struct {
	PlayerID int64
	Status   *string
}

func (u *GameUsecase) listAgentConfigs(ctx context.Context, req *listAgentConfigsReq) ([]*model.AgentConfig, error) {
	if req == nil {
		req = &listAgentConfigsReq{}
	}
	if req.PlayerID <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	resp, err := u.agentConfigRepo.ListAgentConfigs(ctx, &repo.ListAgentConfigsReq{PlayerID: req.PlayerID, Status: req.Status})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (u *GameUsecase) GetPlayer(ctx context.Context, id int64) (*model.Player, error) {
	return u.getPlayer(ctx, id)
}

func (u *GameUsecase) getPlayer(ctx context.Context, id int64) (*model.Player, error) {
	resp, err := u.playerRepo.GetPlayer(ctx, id)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type StartSessionReq struct {
	PlayerID   *int64
	ClientType string
}

func (u *GameUsecase) StartSession(ctx context.Context, req *StartSessionReq) (*model.Session, error) {
	if req == nil {
		req = &StartSessionReq{}
	}
	return u.startSession(ctx, &startSessionReq{PlayerID: req.PlayerID, ClientType: req.ClientType})
}

type startSessionReq struct {
	PlayerID   *int64
	ClientType string
}

func (u *GameUsecase) startSession(ctx context.Context, req *startSessionReq) (*model.Session, error) {
	if req == nil {
		req = &startSessionReq{}
	}
	clientType := req.ClientType
	if clientType == "" {
		clientType = "text_client"
	}
	resp, err := u.sessionRepo.StartSession(ctx, &repo.StartSessionReq{PlayerID: req.PlayerID, ClientType: clientType})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (u *GameUsecase) EndSession(ctx context.Context, id int64) (*model.Session, error) {
	return u.endSession(ctx, id)
}

func (u *GameUsecase) endSession(ctx context.Context, id int64) (*model.Session, error) {
	resp, err := u.sessionRepo.EndSession(ctx, id)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type CreateWorldReq struct {
	PlayerID      int64
	Description   string
	NpcCount      uint32
	LocationCount uint32
	Scale         string
	Seed          *int64
	Tags          []string
	AgentConfigID *int64
}

func (u *GameUsecase) CreateWorld(ctx context.Context, req *CreateWorldReq) (*repo.CreateWorldResp, error) {
	if req == nil {
		req = &CreateWorldReq{}
	}
	return u.createWorld(ctx, &createWorldReq{
		PlayerID:      req.PlayerID,
		Description:   req.Description,
		NpcCount:      req.NpcCount,
		LocationCount: req.LocationCount,
		Scale:         req.Scale,
		Seed:          req.Seed,
		Tags:          req.Tags,
		AgentConfigID: req.AgentConfigID,
	})
}

type createWorldReq struct {
	PlayerID      int64
	Description   string
	NpcCount      uint32
	LocationCount uint32
	Scale         string
	Seed          *int64
	Tags          []string
	AgentConfigID *int64
}

func (u *GameUsecase) createWorld(ctx context.Context, req *createWorldReq) (*repo.CreateWorldResp, error) {
	if req == nil {
		req = &createWorldReq{}
	}
	if req.PlayerID <= 0 || strings.TrimSpace(req.Description) == "" {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_TOWN_WORLD_INVALID)
	}
	npcCount := req.NpcCount
	if npcCount == 0 {
		npcCount = 4
	}
	locationCount := req.LocationCount
	if locationCount == 0 {
		locationCount = 4
	}
	if npcCount > u.conf.GetGameTown().GetWorld().GetMaxNpcCount() || locationCount > u.conf.GetGameTown().GetWorld().GetMaxLocationCount() {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_TOWN_WORLD_INVALID)
	}
	scale := req.Scale
	if scale == "" {
		scale = u.conf.GetGameTown().GetWorld().GetDefaultScale()
	}
	actualSeed := time.Now().UnixNano()
	if req.Seed != nil {
		actualSeed = *req.Seed
	}
	agentConfigResp, err := u.selectedAgentConfig(ctx, &selectedAgentConfigReq{PlayerID: req.PlayerID, AgentConfigID: req.AgentConfigID})
	if err != nil {
		return nil, err
	}
	agentConfig := agentConfigResp
	out, err := u.agent.GenerateWorld(ctx, &agent.GenerateWorldInput{Config: agentRunConfig(agentConfig), Description: req.Description, NpcCount: npcCount, LocationCount: locationCount, Scale: scale, Seed: actualSeed, StyleTags: req.Tags})
	if err != nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_TOWN_AGENT_FAILED).WithCause(err)
	}
	result, err := u.worldRepo.CreateWorld(ctx, &repo.CreateWorldReq{CreatorPlayerID: req.PlayerID, Description: req.Description, NpcCount: npcCount, LocationCount: locationCount, Scale: scale, Seed: actualSeed, StyleTags: req.Tags, AgentConfigID: &agentConfig.ID, Generated: out})
	if err != nil {
		return nil, err
	}
	_, _ = u.agentRunRepo.CreateAgentRun(ctx, &model.AgentRun{WorldID: &result.World.ID, RunType: "world_generate", AgentConfigID: &agentConfig.ID, Model: agentConfig.Model, InputJSON: map[string]any{"description": req.Description, "seed": float64(actualSeed)}, OutputJSON: map[string]any{"world_name": out.WorldName}, Status: "succeeded"})
	return result, nil
}

type JoinWorldReq struct {
	PlayerID  int64
	WorldCode string
}

func (u *GameUsecase) JoinWorld(ctx context.Context, req *JoinWorldReq) (*repo.JoinWorldResp, error) {
	if req == nil {
		req = &JoinWorldReq{}
	}
	return u.joinWorld(ctx, &joinWorldReq{PlayerID: req.PlayerID, WorldCode: req.WorldCode})
}

type joinWorldReq struct {
	PlayerID  int64
	WorldCode string
}

func (u *GameUsecase) joinWorld(ctx context.Context, req *joinWorldReq) (*repo.JoinWorldResp, error) {
	if req == nil {
		req = &joinWorldReq{}
	}
	if req.PlayerID <= 0 || strings.TrimSpace(req.WorldCode) == "" {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	resp, err := u.worldMemberRepo.JoinWorld(ctx, &repo.JoinWorldReq{PlayerID: req.PlayerID, WorldCode: strings.TrimSpace(req.WorldCode)})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (u *GameUsecase) GetWorld(ctx context.Context, id int64) (*model.World, error) {
	return u.getWorld(ctx, id)
}

func (u *GameUsecase) getWorld(ctx context.Context, id int64) (*model.World, error) {
	resp, err := u.worldRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type PageWorldsReq struct {
	Page            *common.PageReq
	CreatorPlayerID *int64
	Status          *string
}

type PageWorldsResp struct {
	Rows []*model.World
	Page *common.PageResp
}

func (u *GameUsecase) PageWorlds(ctx context.Context, req *PageWorldsReq) (*PageWorldsResp, error) {
	if req == nil {
		req = &PageWorldsReq{}
	}
	pageResp, err := u.worldRepo.Page(ctx, &repo.WorldPageReq{Page: req.Page, Query: repo.WorldQuery{CreatorPlayerID: req.CreatorPlayerID, Status: req.Status}})
	if err != nil {
		return nil, err
	}
	return &PageWorldsResp{Rows: pageResp.Rows, Page: pageResp.Page}, nil
}

type GetStateResp struct {
	State   *model.WorldStateSnapshot
	Metrics []*model.WorldMetricDefinition
}

func (u *GameUsecase) GetState(ctx context.Context, worldID int64) (*GetStateResp, error) {
	resp, err := u.getState(ctx, worldID)
	if err != nil {
		return nil, err
	}
	return &GetStateResp{State: resp.State, Metrics: resp.Metrics}, nil
}

type getStateResp struct {
	State   *model.WorldStateSnapshot
	Metrics []*model.WorldMetricDefinition
}

func (u *GameUsecase) getState(ctx context.Context, worldID int64) (*getStateResp, error) {
	stateResp, err := u.worldStateSnapshotRepo.GetLatestWorldState(ctx, worldID)
	if err != nil {
		return nil, err
	}
	metricsResp, err := u.worldMetricDefinitionRepo.ListWorldMetrics(ctx, worldID)
	if err != nil {
		return nil, err
	}
	return &getStateResp{State: stateResp, Metrics: metricsResp}, nil
}

func (u *GameUsecase) GetNpc(ctx context.Context, id int64) (*model.Npc, error) {
	return u.getNpc(ctx, id)
}

func (u *GameUsecase) getNpc(ctx context.Context, id int64) (*model.Npc, error) {
	resp, err := u.npcRepo.GetNpc(ctx, id)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ListNpcsReq struct {
	WorldID    int64
	LocationID *int64
}

func (u *GameUsecase) ListNpcs(ctx context.Context, req *ListNpcsReq) ([]*model.Npc, error) {
	if req == nil {
		req = &ListNpcsReq{}
	}
	return u.listNpcs(ctx, &listNpcsReq{WorldID: req.WorldID, LocationID: req.LocationID})
}

type listNpcsReq struct {
	WorldID    int64
	LocationID *int64
}

func (u *GameUsecase) listNpcs(ctx context.Context, req *listNpcsReq) ([]*model.Npc, error) {
	if req == nil {
		req = &listNpcsReq{}
	}
	resp, err := u.npcRepo.ListNpcs(ctx, &repo.ListNpcsReq{WorldID: req.WorldID, LocationID: req.LocationID})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ListMemoriesReq struct {
	WorldID  int64
	PlayerID int64
	NpcID    *int64
	Type     *string
}

func (u *GameUsecase) ListMemories(ctx context.Context, req *ListMemoriesReq) ([]*model.Memory, error) {
	if req == nil {
		req = &ListMemoriesReq{}
	}
	listResp, err := u.memoryRepo.ListMemories(ctx, &repo.MemoryQuery{WorldID: req.WorldID, PlayerID: req.PlayerID, NpcID: req.NpcID, Type: req.Type})
	if err != nil {
		return nil, err
	}
	return listResp, nil
}

type PageEventsReq struct {
	Page          *common.PageReq
	WorldID       int64
	ActorPlayerID *int64
	TargetNpcID   *int64
	Type          *string
}

type PageEventsResp struct {
	Rows []*model.Event
	Page *common.PageResp
}

func (u *GameUsecase) PageEvents(ctx context.Context, req *PageEventsReq) (*PageEventsResp, error) {
	if req == nil {
		req = &PageEventsReq{}
	}
	pageResp, err := u.eventRepo.Page(ctx, &repo.EventPageReq{Page: req.Page, Query: repo.EventQuery{WorldID: req.WorldID, ActorPlayerID: req.ActorPlayerID, TargetNpcID: req.TargetNpcID, Type: req.Type}})
	if err != nil {
		return nil, err
	}
	return &PageEventsResp{Rows: pageResp.Rows, Page: pageResp.Page}, nil
}

type PageCommandsReq struct {
	Page      *common.PageReq
	WorldID   *int64
	SessionID *int64
	PlayerID  *int64
}

type PageCommandsResp struct {
	Rows []*model.Command
	Page *common.PageResp
}

func (u *GameUsecase) PageCommands(ctx context.Context, req *PageCommandsReq) (*PageCommandsResp, error) {
	if req == nil {
		req = &PageCommandsReq{}
	}
	pageResp, err := u.commandRepo.Page(ctx, &repo.CommandPageReq{Page: req.Page, Query: repo.CommandQuery{WorldID: req.WorldID, SessionID: req.SessionID, PlayerID: req.PlayerID}})
	if err != nil {
		return nil, err
	}
	return &PageCommandsResp{Rows: pageResp.Rows, Page: pageResp.Page}, nil
}

type ReplayCommandsReq struct {
	SessionID int64
	PlayerID  int64
}

type ReplayCommandsResp struct {
	Commands []*model.Command
	Events   []*model.Event
}

func (u *GameUsecase) ReplayCommands(ctx context.Context, req *ReplayCommandsReq) (*ReplayCommandsResp, error) {
	if req == nil {
		req = &ReplayCommandsReq{}
	}
	resp, err := u.replayCommands(ctx, &replayCommandsReq{SessionID: req.SessionID, PlayerID: req.PlayerID})
	if err != nil {
		return nil, err
	}
	return &ReplayCommandsResp{Commands: resp.Commands, Events: resp.Events}, nil
}

type replayCommandsReq struct {
	SessionID int64
	PlayerID  int64
}

type replayCommandsResp struct {
	Commands []*model.Command
	Events   []*model.Event
}

func (u *GameUsecase) replayCommands(ctx context.Context, req *replayCommandsReq) (*replayCommandsResp, error) {
	if req == nil {
		req = &replayCommandsReq{}
	}
	commandResp, err := u.commandRepo.List(ctx, &repo.CommandListReq{SessionID: req.SessionID, PlayerID: req.PlayerID})
	if err != nil {
		return nil, err
	}
	var worldID int64
	commands := commandResp
	for _, item := range commands {
		if item.WorldID != nil {
			worldID = *item.WorldID
			break
		}
	}
	if worldID == 0 {
		return &replayCommandsResp{Commands: commands}, nil
	}
	eventResp, err := u.eventRepo.Page(ctx, &repo.EventPageReq{Page: &common.PageReq{Page: 1, Size: 100}, Query: repo.EventQuery{WorldID: worldID}})
	if err != nil {
		return nil, err
	}
	return &replayCommandsResp{Commands: commands, Events: eventResp.Rows}, nil
}

type TickReq struct {
	WorldID          int64
	OperatorPlayerID int64
	Limit            uint32
}

type TickResp struct {
	State  *model.WorldStateSnapshot
	Events []*model.Event
}

func (u *GameUsecase) Tick(ctx context.Context, req *TickReq) (*TickResp, error) {
	if req == nil {
		req = &TickReq{}
	}
	resp, err := u.tick(ctx, &tickReq{WorldID: req.WorldID, OperatorPlayerID: req.OperatorPlayerID, Limit: req.Limit})
	if err != nil {
		return nil, err
	}
	return &TickResp{State: resp.State, Events: resp.Events}, nil
}

type tickReq struct {
	WorldID          int64
	OperatorPlayerID int64
	Limit            uint32
}

type tickResp struct {
	State  *model.WorldStateSnapshot
	Events []*model.Event
}

func (u *GameUsecase) tick(ctx context.Context, req *tickReq) (*tickResp, error) {
	if req == nil {
		req = &tickReq{}
	}
	worldResp, err := u.worldRepo.Get(ctx, req.WorldID)
	if err != nil {
		return nil, err
	}
	stateResp, err := u.worldStateSnapshotRepo.GetLatestWorldState(ctx, req.WorldID)
	if err != nil {
		return nil, err
	}
	recentResp, err := u.eventRepo.ListRecentEvents(ctx, &repo.ListRecentEventsReq{WorldID: req.WorldID, Limit: int(req.Limit)})
	if err != nil {
		return nil, err
	}
	recent := recentResp
	summaries := make([]string, 0, len(recent))
	for _, item := range recent {
		summaries = append(summaries, item.Summary)
	}
	worldRow := worldResp
	agentConfigResp, err := u.worldAgentConfig(ctx, worldRow)
	if err != nil {
		return nil, err
	}
	agentConfig := agentConfigResp
	out, err := u.agent.Direct(ctx, &agent.DirectInput{Config: agentRunConfig(agentConfig), WorldName: worldRow.Name, Arc: stateResp.CurrentArc, Metrics: stateResp.Metrics, Events: summaries})
	if err != nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_TOWN_AGENT_FAILED).WithCause(err)
	}
	metrics := applyMetricDelta(stateResp.Metrics, out.WorldMetricDelta, u.conf.GetGameTown().GetEffect().GetMaxDelta())
	currentArc := out.CurrentArc
	if strings.TrimSpace(currentArc) == "" {
		currentArc = stateResp.CurrentArc
	}
	eventResp, err := u.eventRepo.CreateEvent(ctx, &model.Event{WorldID: req.WorldID, Type: "world_tick", ActorPlayerID: &req.OperatorPlayerID, Summary: out.Summary, Content: strings.Join(out.Events, "\n"), Effects: map[string]any{"metrics": out.WorldMetricDelta}, Metadata: map[string]any{}, OccurredAt: time.Now()})
	if err != nil {
		return nil, err
	}
	newStateResp, err := u.worldStateSnapshotRepo.CreateState(ctx, &model.WorldStateSnapshot{WorldID: req.WorldID, TickCount: stateResp.TickCount + 1, CurrentArc: currentArc, Metrics: metrics, Summary: out.Summary, ReasonEventID: &eventResp.ID})
	if err != nil {
		return nil, err
	}
	return &tickResp{State: newStateResp, Events: []*model.Event{eventResp}}, nil
}

type ExecuteCommandReq struct {
	SessionID int64
	PlayerID  int64
	Raw       string
}

func (u *GameUsecase) ExecuteCommand(ctx context.Context, req *ExecuteCommandReq) (*CommandResult, error) {
	if req == nil {
		req = &ExecuteCommandReq{}
	}
	return u.executeCommand(ctx, &executeCommandReq{SessionID: req.SessionID, PlayerID: req.PlayerID, Raw: req.Raw})
}

type executeCommandReq struct {
	SessionID int64
	PlayerID  int64
	Raw       string
}

func (u *GameUsecase) executeCommand(ctx context.Context, req *executeCommandReq) (*CommandResult, error) {
	if req == nil {
		req = &executeCommandReq{}
	}
	sessionID := req.SessionID
	playerID := req.PlayerID
	raw := strings.TrimSpace(req.Raw)
	if raw == "" {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_TOWN_COMMAND_INVALID)
	}
	cmdType, payload := parseCommand(raw)
	createdResp, err := u.commandRepo.CreateCommand(ctx, &model.Command{SessionID: sessionID, PlayerID: &playerID, RawText: raw, Type: cmdType, ParsedPayload: payload, Status: "received", ResultSummary: ""})
	if err != nil {
		return nil, err
	}
	result := &CommandResult{Command: createdResp}
	var finishWorldID *int64
	defer func() {
		if result.Command != nil && result.Command.Status == "received" {
			updated, _ := u.commandRepo.FinishCommand(ctx, &repo.FinishCommandReq{ID: result.Command.ID, Status: "succeeded", Summary: result.Command.ResultSummary, WorldID: finishWorldID})
			if updated != nil {
				result.Command = updated
			}
		}
	}()
	sessionResp, err := u.sessionRepo.GetSession(ctx, sessionID)
	if err != nil {
		return nil, err
	}
	switch cmdType {
	case "help":
		result.FeedbackLines = []string{"/register <name>", "/agent config create <name> --provider blades --model <model> --base-url <url> --api-key <key> --default", "/agent configs", "/world create <闂傚倸鍊搁崐鎼佸磹閻戣姤鍤勯柛顐ｆ磸閳ь兛鐒︾换婵嬪炊閼稿灚娅旈柣鐔哥矊缁绘﹢鐛幋锕€顫呴柣妯荤墬濡啫鐣烽妸鈺婃晣闁靛繒濮甸弳? --npc 4 --locations 4 --scale small --seed 42 --agent-config <id>", "/world join <world_code>", "/look", "/move <location_code>", "/talk <npc_code> <content>", "/do <content>", "/tick"}
		result.Command.ResultSummary = "help"
	case "agent_config_create":
		createResp, err := u.createAgentConfig(ctx, &createAgentConfigReq{PlayerID: playerID, Name: stringValue(payload, "name"), Provider: stringValue(payload, "provider"), ModelName: stringValue(payload, "model"), BaseURL: stringValue(payload, "base_url"), APIKey: stringValue(payload, "api_key"), TimeoutSeconds: int32(uint32Value(payload, "timeout_seconds")), IsDefault: boolValue(payload, "is_default")})
		if err != nil {
			return nil, err
		}
		row := createResp
		result.FeedbackLines = []string{fmt.Sprintf("Agent config created: %s (%d)", row.Name, row.ID)}
		result.Command.ResultSummary = "create agent config"
	case "agent_config_list":
		listResp, err := u.listAgentConfigs(ctx, &listAgentConfigsReq{PlayerID: playerID})
		if err != nil {
			return nil, err
		}
		rows := listResp
		lines := make([]string, 0, len(rows))
		for _, row := range rows {
			lines = append(lines, fmt.Sprintf("%d %s provider=%s model=%s default=%t", row.ID, row.Name, row.Provider, row.Model, row.IsDefault))
		}
		if len(lines) == 0 {
			lines = []string{"no agent config"}
		}
		result.FeedbackLines = lines
		result.Command.ResultSummary = "list agent config"
	case "world_create":
		desc := stringValue(payload, "description")
		npcCount := uint32Value(payload, "npc_count")
		locationCount := uint32Value(payload, "location_count")
		seedValue, hasSeed := int64Value(payload, "seed")
		var seed *int64
		if hasSeed {
			seed = &seedValue
		}
		agentConfigIDValue, hasAgentConfigID := int64Value(payload, "agent_config_id")
		var agentConfigID *int64
		if hasAgentConfigID {
			agentConfigID = &agentConfigIDValue
		}
		createWorldResp, err := u.createWorld(ctx, &createWorldReq{PlayerID: playerID, Description: desc, NpcCount: npcCount, LocationCount: locationCount, Scale: stringValue(payload, "scale"), Seed: seed, AgentConfigID: agentConfigID})
		if err != nil {
			return nil, err
		}
		createdWorld := createWorldResp
		finishWorldID = &createdWorld.World.ID
		_, _ = u.sessionRepo.UpdateSessionWorld(ctx, &repo.UpdateSessionWorldReq{ID: sessionID, PlayerID: playerID, WorldID: createdWorld.World.ID})
		result.CurrentWorld = createdWorld.World
		result.CurrentLocation = createdWorld.DefaultLocation
		result.VisibleNpcs = createdWorld.Npcs
		result.WorldState = createdWorld.State
		result.Events = createdWorld.Events
		result.FeedbackLines = []string{fmt.Sprintf("world created: %s (%s)", createdWorld.World.Name, createdWorld.World.Code)}
		result.Command.ResultSummary = "create world"
	case "world_join":
		joinResp, err := u.joinWorld(ctx, &joinWorldReq{PlayerID: playerID, WorldCode: stringValue(payload, "world_code")})
		if err != nil {
			return nil, err
		}
		joined := joinResp
		finishWorldID = &joined.World.ID
		_, _ = u.sessionRepo.UpdateSessionWorld(ctx, &repo.UpdateSessionWorldReq{ID: sessionID, PlayerID: playerID, WorldID: joined.World.ID})
		result.CurrentWorld = joined.World
		result.CurrentLocation = joined.Location
		npcResp, _ := u.npcRepo.ListNpcs(ctx, &repo.ListNpcsReq{WorldID: joined.World.ID, LocationID: &joined.Location.ID})
		result.VisibleNpcs = npcResp
		stateResp, _ := u.worldStateSnapshotRepo.GetLatestWorldState(ctx, joined.World.ID)
		result.WorldState = stateResp
		result.FeedbackLines = []string{fmt.Sprintf("joined world: %s", joined.World.Name)}
		result.Command.ResultSummary = "join world"
	default:
		sessionRow := sessionResp
		if sessionRow.CurrentWorldID == nil {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_TOWN_WORLD_NOT_FOUND)
		}
		worldID := *sessionRow.CurrentWorldID
		finishWorldID = &worldID
		executeResp, err := u.executeWorldCommand(ctx, &executeWorldCommandReq{Result: result, WorldID: worldID, PlayerID: playerID, CommandType: cmdType, Payload: payload})
		if err != nil {
			return nil, err
		}
		return executeResp, nil
	}
	return result, nil
}

type executeWorldCommandReq struct {
	Result      *CommandResult
	WorldID     int64
	PlayerID    int64
	CommandType string
	Payload     map[string]any
}

func (u *GameUsecase) executeWorldCommand(ctx context.Context, req *executeWorldCommandReq) (*CommandResult, error) {
	if req == nil {
		req = &executeWorldCommandReq{}
	}
	result := req.Result
	if result == nil {
		result = &CommandResult{}
	}
	worldID := req.WorldID
	playerID := req.PlayerID
	cmdType := req.CommandType
	payload := req.Payload
	worldResp, err := u.worldRepo.Get(ctx, worldID)
	if err != nil {
		return nil, err
	}
	memberResp, err := u.worldMemberRepo.GetMember(ctx, &repo.GetMemberReq{WorldID: worldID, PlayerID: playerID})
	if err != nil {
		return nil, err
	}
	member := memberResp
	locationResp, err := u.locationRepo.GetLocation(ctx, member.CurrentLocationID)
	if err != nil {
		return nil, err
	}
	stateResp, _ := u.worldStateSnapshotRepo.GetLatestWorldState(ctx, worldID)
	locationRow := locationResp
	npcResp, _ := u.npcRepo.ListNpcs(ctx, &repo.ListNpcsReq{WorldID: worldID, LocationID: &locationRow.ID})
	worldRow := worldResp
	result.CurrentWorld = worldRow
	result.CurrentLocation = locationRow
	result.VisibleNpcs = npcResp
	result.WorldState = stateResp
	switch cmdType {
	case "look":
		result.FeedbackLines = []string{fmt.Sprintf("%s: %s", locationRow.Name, locationRow.Description)}
		result.Command.ResultSummary = "look"
	case "move":
		locResp, err := u.locationRepo.GetLocationByCode(ctx, &repo.GetLocationByCodeReq{WorldID: worldID, Code: stringValue(payload, "location_code")})
		if err != nil {
			return nil, err
		}
		loc := locResp
		_, err = u.worldMemberRepo.MoveMember(ctx, &repo.MoveMemberReq{WorldID: worldID, PlayerID: playerID, LocationID: loc.ID})
		if err != nil {
			return nil, err
		}
		eventResp, err := u.eventRepo.CreateEvent(ctx, &model.Event{WorldID: worldID, Type: "player_moved", ActorPlayerID: &playerID, LocationID: &loc.ID, CommandID: &result.Command.ID, Summary: fmt.Sprintf("player moved to %s", loc.Name), Effects: map[string]any{}, Metadata: map[string]any{}, OccurredAt: time.Now()})
		if err != nil {
			return nil, err
		}
		result.CurrentLocation = loc
		moveNpcResp, _ := u.npcRepo.ListNpcs(ctx, &repo.ListNpcsReq{WorldID: worldID, LocationID: &loc.ID})
		result.VisibleNpcs = moveNpcResp
		result.Events = []*model.Event{eventResp}
		result.FeedbackLines = []string{fmt.Sprintf("arrived at %s", loc.Name)}
		result.Command.ResultSummary = "move"
	case "talk":
		npcResp, err := u.npcRepo.GetNpcByCode(ctx, &repo.GetNpcByCodeReq{WorldID: worldID, Code: stringValue(payload, "npc_code")})
		if err != nil {
			return nil, err
		}
		content := stringValue(payload, "content")
		npcRow := npcResp
		relResp, _ := u.relationshipRepo.GetRelationship(ctx, &repo.GetRelationshipReq{WorldID: worldID, PlayerID: playerID, NpcID: npcRow.ID})
		memoryResp, _ := u.memoryRepo.ListMemories(ctx, &repo.MemoryQuery{WorldID: worldID, PlayerID: playerID, NpcID: &npcRow.ID})
		memories := memoryResp
		memoryTexts := make([]string, 0, len(memories))
		for _, item := range memories {
			memoryTexts = append(memoryTexts, item.Content)
		}
		worldRow := worldResp
		agentConfigResp, err := u.worldAgentConfig(ctx, worldRow)
		if err != nil {
			return nil, err
		}
		agentConfig := agentConfigResp
		worldMetrics := map[string]any{}
		if stateResp != nil && stateResp.Metrics != nil {
			worldMetrics = stateResp.Metrics
		}
		out, err := u.agent.Talk(ctx, &agent.TalkInput{Config: agentRunConfig(agentConfig), WorldName: worldRow.Name, LocationName: locationRow.Name, NpcName: npcRow.Name, NpcRole: npcRow.Role, NpcPersonality: npcRow.Personality, Relationship: map[string]any{"affinity": relResp.Affinity, "trust": relResp.Trust, "tension": relResp.Tension}, Memories: memoryTexts, WorldMetrics: worldMetrics, Content: content})
		if err != nil {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_TOWN_AGENT_FAILED).WithCause(err)
		}
		rel := relResp
		rel.Affinity = clamp(rel.Affinity + limitedDelta(out.RelationshipDelta["affinity"], u.conf.GetGameTown().GetEffect().GetMaxDelta()))
		rel.Trust = clamp(rel.Trust + limitedDelta(out.RelationshipDelta["trust"], u.conf.GetGameTown().GetEffect().GetMaxDelta()))
		rel.Tension = clamp(rel.Tension + limitedDelta(out.RelationshipDelta["tension"], u.conf.GetGameTown().GetEffect().GetMaxDelta()))
		_, _ = u.relationshipRepo.UpsertRelationship(ctx, rel)
		eventResp, err := u.eventRepo.CreateEvent(ctx, &model.Event{WorldID: worldID, Type: "player_talked_to_npc", ActorPlayerID: &playerID, TargetNpcID: &npcRow.ID, LocationID: &locationRow.ID, CommandID: &result.Command.ID, Summary: fmt.Sprintf("player talked to %s", npcRow.Name), Content: content, Effects: map[string]any{"relationship": out.RelationshipDelta, "metrics": out.WorldMetricDelta}, Metadata: map[string]any{}, OccurredAt: time.Now()})
		if err != nil {
			return nil, err
		}
		for _, candidate := range out.MemoryCandidates {
			_, _ = u.memoryRepo.CreateMemory(ctx, &model.Memory{WorldID: worldID, PlayerID: playerID, NpcID: npcRow.ID, Type: "long_term", Content: candidate, Importance: 50, SourceEventID: &eventResp.ID})
		}
		_, _ = u.agentRunRepo.CreateAgentRun(ctx, &model.AgentRun{WorldID: &worldID, RunType: "npc_talk", CommandID: &result.Command.ID, NpcID: &npcRow.ID, AgentConfigID: &agentConfig.ID, Model: agentConfig.Model, InputJSON: map[string]any{"content": content}, OutputJSON: map[string]any{"reply": out.Reply}, Status: "succeeded"})
		result.Events = []*model.Event{eventResp}
		result.FeedbackLines = []string{out.Reply}
		result.Command.ResultSummary = "npc talk"
	case "do":
		content := stringValue(payload, "content")
		eventResp, err := u.eventRepo.CreateEvent(ctx, &model.Event{WorldID: worldID, Type: "player_action", ActorPlayerID: &playerID, LocationID: &locationRow.ID, CommandID: &result.Command.ID, Summary: "player action", Content: content, Effects: map[string]any{"activity": 1}, Metadata: map[string]any{}, OccurredAt: time.Now()})
		if err != nil {
			return nil, err
		}
		result.Events = []*model.Event{eventResp}
		result.FeedbackLines = []string{"action recorded"}
		result.Command.ResultSummary = "player action"
	case "tick":
		tickResp, err := u.tick(ctx, &tickReq{WorldID: worldID, OperatorPlayerID: playerID, Limit: 20})
		if err != nil {
			return nil, err
		}
		result.WorldState = tickResp.State
		result.Events = tickResp.Events
		result.FeedbackLines = []string{tickResp.State.Summary}
		result.Command.ResultSummary = "world tick"
	case "events":
		eventsResp, err := u.eventRepo.Page(ctx, &repo.EventPageReq{Page: &common.PageReq{Page: 1, Size: 10}, Query: repo.EventQuery{WorldID: worldID}})
		if err != nil {
			return nil, err
		}
		result.Events = eventsResp.Rows
		result.FeedbackLines = []string{fmt.Sprintf("recent events: %d", len(eventsResp.Rows))}
		result.Command.ResultSummary = "events"
	case "memory":
		npcCode := stringValue(payload, "npc_code")
		var npcID *int64
		if npcCode != "" {
			npcResp, err := u.npcRepo.GetNpcByCode(ctx, &repo.GetNpcByCodeReq{WorldID: worldID, Code: npcCode})
			if err != nil {
				return nil, err
			}
			npcID = &npcResp.ID
		}
		memoryResp, err := u.memoryRepo.ListMemories(ctx, &repo.MemoryQuery{WorldID: worldID, PlayerID: playerID, NpcID: npcID})
		memories := memoryResp
		if err != nil {
			return nil, err
		}
		lines := make([]string, 0, len(memories))
		for _, item := range memories {
			lines = append(lines, item.Content)
		}
		if len(lines) == 0 {
			lines = []string{"no memory"}
		}
		result.FeedbackLines = lines
		result.Command.ResultSummary = "list memory"
	case "npcs":
		result.FeedbackLines = []string{fmt.Sprintf("current location npc count: %d", len(result.VisibleNpcs))}
		result.Command.ResultSummary = "list npc"
	case "status":
		result.FeedbackLines = []string{fmt.Sprintf("world: %s, arc: %s", worldRow.Name, stateResp.CurrentArc)}
		result.Command.ResultSummary = "status"
	default:
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_TOWN_COMMAND_INVALID)
	}
	return result, nil
}

func parseCommand(raw string) (string, map[string]any) {
	payload := map[string]any{}
	if raw == "/help" {
		return "help", payload
	}
	if strings.HasPrefix(raw, "/agent config create") {
		content := strings.TrimSpace(strings.TrimPrefix(raw, "/agent config create"))
		payload["name"] = quotedOrRest(content)
		payload["provider"] = stringFlag(raw, "--provider", "blades")
		payload["model"] = stringFlag(raw, "--model", "")
		payload["base_url"] = stringFlag(raw, "--base-url", "")
		payload["api_key"] = stringFlag(raw, "--api-key", "")
		payload["timeout_seconds"] = float64(intFlag(raw, "--timeout", 30))
		payload["is_default"] = boolFlag(raw, "--default")
		return "agent_config_create", payload
	}
	if raw == "/agent configs" {
		return "agent_config_list", payload
	}
	if strings.HasPrefix(raw, "/world create") {
		payload["description"] = quotedOrRest(strings.TrimSpace(strings.TrimPrefix(raw, "/world create")))
		payload["npc_count"] = float64(intFlag(raw, "--npc", 4))
		payload["location_count"] = float64(intFlag(raw, "--locations", 4))
		payload["scale"] = stringFlag(raw, "--scale", "small")
		if agentConfigID, ok := int64Flag(raw, "--agent-config"); ok {
			payload["agent_config_id"] = float64(agentConfigID)
		}
		if seed, ok := int64Flag(raw, "--seed"); ok {
			payload["seed"] = float64(seed)
		}
		return "world_create", payload
	}
	if strings.HasPrefix(raw, "/world join ") {
		payload["world_code"] = strings.TrimSpace(strings.TrimPrefix(raw, "/world join "))
		return "world_join", payload
	}
	if raw == "/look" {
		return "look", payload
	}
	if strings.HasPrefix(raw, "/move ") {
		payload["location_code"] = strings.TrimSpace(strings.TrimPrefix(raw, "/move "))
		return "move", payload
	}
	if strings.HasPrefix(raw, "/talk ") {
		rest := strings.TrimSpace(strings.TrimPrefix(raw, "/talk "))
		parts := strings.SplitN(rest, " ", 2)
		if len(parts) > 0 {
			payload["npc_code"] = parts[0]
		}
		if len(parts) > 1 {
			payload["content"] = parts[1]
		}
		return "talk", payload
	}
	if strings.HasPrefix(raw, "/do ") {
		payload["content"] = strings.TrimSpace(strings.TrimPrefix(raw, "/do "))
		return "do", payload
	}
	if raw == "/tick" {
		return "tick", payload
	}
	if raw == "/events" {
		return "events", payload
	}
	if raw == "/npcs" {
		return "npcs", payload
	}
	if strings.HasPrefix(raw, "/memory") {
		payload["npc_code"] = strings.TrimSpace(strings.TrimPrefix(raw, "/memory"))
		return "memory", payload
	}
	if raw == "/status" {
		return "status", payload
	}
	return "help", payload
}

func quotedOrRest(value string) string {
	value = strings.TrimSpace(value)
	if strings.HasPrefix(value, "\"") {
		end := strings.Index(value[1:], "\"")
		if end >= 0 {
			return value[1 : end+1]
		}
	}
	return strings.Split(value, " --")[0]
}
func intFlag(raw string, name string, fallback int) int {
	v, ok := int64Flag(raw, name)
	if !ok {
		return fallback
	}
	return int(v)
}
func int64Flag(raw string, name string) (int64, bool) {
	parts := strings.Split(raw, " ")
	for i := 0; i < len(parts)-1; i++ {
		if parts[i] == name {
			v, err := strconv.ParseInt(parts[i+1], 10, 64)
			return v, err == nil
		}
	}
	return 0, false
}
func boolFlag(raw string, name string) bool {
	for _, part := range strings.Split(raw, " ") {
		if part == name {
			return true
		}
	}
	return false
}
func stringFlag(raw string, name string, fallback string) string {
	parts := strings.Split(raw, " ")
	for i := 0; i < len(parts)-1; i++ {
		if parts[i] == name {
			return parts[i+1]
		}
	}
	return fallback
}
func stringValue(m map[string]any, key string) string {
	if v, ok := m[key].(string); ok {
		return strings.TrimSpace(v)
	}
	return ""
}
func uint32Value(m map[string]any, key string) uint32 {
	if v, ok := m[key].(float64); ok {
		return uint32(v)
	}
	return 0
}
func boolValue(m map[string]any, key string) bool {
	if v, ok := m[key].(bool); ok {
		return v
	}
	return false
}
func int64Value(m map[string]any, key string) (int64, bool) {
	if v, ok := m[key].(float64); ok {
		return int64(v), true
	}
	return 0, false
}
func applyMetricDelta(metrics map[string]any, delta map[string]int32, maxDelta int32) map[string]any {
	result := map[string]any{}
	for k, v := range metrics {
		result[k] = v
	}
	for k, v := range delta {
		current := numeric(result[k])
		result[k] = float64(clamp(int32(current) + limitedDelta(v, maxDelta)))
	}
	return result
}
func numeric(v any) int {
	switch val := v.(type) {
	case float64:
		return int(val)
	case int:
		return val
	case int32:
		return int(val)
	case int64:
		return int(val)
	default:
		return 0
	}
}
func limitedDelta(v int32, max int32) int32 {
	if max <= 0 {
		return v
	}
	if v > max {
		return max
	}
	if v < -max {
		return -max
	}
	return v
}
func clamp(v int32) int32 {
	if v < 0 {
		return 0
	}
	if v > 100 {
		return 100
	}
	return v
}

type selectedAgentConfigReq struct {
	PlayerID      int64
	AgentConfigID *int64
}

func (u *GameUsecase) selectedAgentConfig(ctx context.Context, req *selectedAgentConfigReq) (*model.AgentConfig, error) {
	if req == nil {
		req = &selectedAgentConfigReq{}
	}
	if req.AgentConfigID != nil {
		resp, err := u.agentConfigRepo.GetAgentConfig(ctx, &repo.GetAgentConfigReq{ID: *req.AgentConfigID, PlayerID: req.PlayerID})
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
	resp, err := u.agentConfigRepo.GetDefaultAgentConfig(ctx, req.PlayerID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (u *GameUsecase) worldAgentConfig(ctx context.Context, worldRow *model.World) (*model.AgentConfig, error) {
	if worldRow == nil || worldRow.AgentConfigID == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_TOWN_AGENT_CONFIG_NOT_FOUND)
	}
	resp, err := u.agentConfigRepo.GetAgentConfig(ctx, &repo.GetAgentConfigReq{ID: *worldRow.AgentConfigID, PlayerID: worldRow.CreatorPlayerID})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func agentRunConfig(config *model.AgentConfig) *agent.RunConfig {
	if config == nil {
		return nil
	}
	return &agent.RunConfig{ID: config.ID, Provider: config.Provider, Model: config.Model, BaseURL: config.BaseURL, APIKey: config.APIKey, TimeoutSeconds: config.TimeoutSeconds}
}
