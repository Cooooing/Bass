package main

import (
	v1enum "common/proto/gen/game_town/v1/enum"
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"math/rand/v2"
	"os"
	"strings"
	"sync"
	"time"

	commonClient "common/pkg/client"
	"common/pkg/client/rpc"
	"common/pkg/constant"
	commonpb "common/proto/gen/common"
	commonv1 "common/proto/gen/common/v1"
	v1 "common/proto/gen/game_town/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/durationpb"
)

type config struct {
	addr             string
	consulAddr       string
	consulDatacenter string
	consulToken      string
	players          int
	rounds           int
	bigEventEvery    int
	ollamaURL        string
	model            string
	worldDescription string
	timeout          time.Duration
	settleTimeout    time.Duration
	pollInterval     time.Duration
	strict           bool
}

type playerState struct {
	id         int64
	name       string
	memberID   int64
	locationID int64
	ready      bool
}

type runState struct {
	mu          sync.Mutex
	configID    int64
	worldID     int64
	worldCode   string
	players     []playerState
	playerSeq   map[int64]uint64
	seenEvent   map[int64]bool
	playerStats map[int64]*playerRunStats
	stats       runStats
}

type runStats struct {
	submittedActions int
	bigEvents        int
	pageEvents       int
	streamEvents     int
	streamReconnects int
	streamErrors     int
	npcReplies       int
	worldEvolved     int
	npcPlanned       int
	worldTicks       int
	resolved         int
	rejected         int
	clarification    int
	jobFailed        int
	characterReady   int
}

type playerRunStats struct {
	submitted      int
	pageEvents     int
	streamEvents   int
	visibleEvents  int
	npcReplies     int
	actionResults  int
	characterReady bool
}

func main() {
	conf := parseFlags()
	ctx, cancel := context.WithTimeout(context.Background(), conf.timeout)
	defer cancel()

	client, cleanup, target, err := newGameTownClient(conf)
	if err != nil {
		fatal("connect game_town failed: %v", err)
	}
	defer cleanup()

	fmt.Printf("Game Town load client connected: %s\n", target)
	if err := checkGameTownHealth(ctx, client); err != nil {
		fatal("game_town health check failed: %v", err)
	}
	state, err := run(ctx, client, conf)
	if err != nil {
		fatal("load test failed: %v", err)
	}
	printReport(state)
}

func parseFlags() config {
	var conf config
	flag.StringVar(&conf.addr, "addr", "", "direct game_town gRPC address, for example 192.168.100.1:9105")
	flag.StringVar(&conf.consulAddr, "consul-addr", envDefault("CONSUL_ADDRESS", envDefault("CONSUL_HOST", "127.0.0.1")+":"+envDefault("CONSUL_PORT", "8500")), "Consul address")
	flag.StringVar(&conf.consulDatacenter, "consul-datacenter", envDefault("CONSUL_DATACENTER", "dc1"), "Consul datacenter")
	flag.StringVar(&conf.consulToken, "consul-token", envDefault("CONSUL_ACL_APP_TOKEN", ""), "Consul token")
	flag.IntVar(&conf.players, "players", 2, "number of players; minimum is 2")
	flag.IntVar(&conf.rounds, "rounds", 20, "number of submitted action rounds; use 1000 for acceptance")
	flag.IntVar(&conf.bigEventEvery, "big-event-every", 100, "inject one major event every N rounds; 0 disables major events")
	flag.StringVar(&conf.ollamaURL, "ollama-url", "http://192.168.100.10:31434", "Ollama base URL")
	flag.StringVar(&conf.model, "model", "qwen3:1.7b", "chat model")
	flag.StringVar(&conf.worldDescription, "world-description", defaultWorldDescription(), "world description")
	flag.DurationVar(&conf.timeout, "timeout", 2*time.Hour, "whole load-test timeout")
	flag.DurationVar(&conf.settleTimeout, "settle-timeout", 0, "maximum time to wait for asynchronous world feedback; 0 uses an automatic value")
	flag.DurationVar(&conf.pollInterval, "poll-interval", 2*time.Second, "async polling interval")
	flag.BoolVar(&conf.strict, "strict", false, "fail when acceptance counters are missing")
	flag.Parse()

	if conf.players < 2 {
		conf.players = 2
	}
	if conf.rounds < 1 {
		conf.rounds = 1
	}
	return conf
}

func checkGameTownHealth(ctx context.Context, client *rpc.GameTownClient) error {
	if client == nil || client.System == nil {
		return errors.New("system health client is nil")
	}
	healthCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	resp, err := client.System.Health(healthCtx, &commonv1.HealthSystem_Req{})
	if err != nil {
		return err
	}
	if strings.TrimSpace(resp.GetMessage()) == "" {
		return errors.New("empty health response")
	}
	fmt.Printf("game_town health: %s\n", resp.GetMessage())
	return nil
}
func run(ctx context.Context, client *rpc.GameTownClient, conf config) (*runState, error) {
	state := &runState{
		playerSeq:   make(map[int64]uint64),
		seenEvent:   make(map[int64]bool),
		playerStats: make(map[int64]*playerRunStats),
	}
	configID, err := ensureAgentConfig(ctx, client, conf)
	if err != nil {
		return nil, err
	}
	state.configID = configID

	players, err := registerPlayers(ctx, client, conf.players)
	if err != nil {
		return nil, err
	}
	state.players = players
	state.initPlayerStats(players)

	world, err := client.World.Create(ctx, &v1.CreateGameTownWorld_Request{
		CreatorPlayerId: players[0].id,
		Description:     conf.worldDescription,
		NpcCount:        4,
		LocationCount:   4,
		AgentConfigId:   configID,
	})
	if err != nil {
		return nil, fmt.Errorf("create world failed: %w", err)
	}
	state.worldID = world.GetRow().GetId()
	state.worldCode = world.GetRow().GetCode()
	fmt.Printf("world queued: id=%d code=%s event=%d\n", state.worldID, state.worldCode, world.GetEventId())

	if err := waitWorldReady(ctx, client, state, conf.pollInterval); err != nil {
		return nil, err
	}
	if err := joinPlayers(ctx, client, state); err != nil {
		return nil, err
	}

	watchCtx, stopWatch := context.WithCancel(ctx)
	var watchWG sync.WaitGroup
	startWatchers(watchCtx, &watchWG, client, state, conf.pollInterval)
	defer func() {
		stopWatch()
		watchWG.Wait()
	}()

	if err := waitCharactersReady(ctx, client, state, conf.pollInterval); err != nil {
		return nil, err
	}
	if err := playRounds(ctx, client, state, conf); err != nil {
		return nil, err
	}
	if err := collectEvents(ctx, client, state); err != nil {
		return nil, err
	}
	if err := validateRun(state, conf); err != nil {
		return nil, err
	}
	return state, nil
}

func ensureAgentConfig(ctx context.Context, client *rpc.GameTownClient, conf config) (int64, error) {
	name := "load-local-ollama-" + strings.NewReplacer(":", "-", "/", "-", ".", "-").Replace(conf.model)
	created, err := client.AgentConfig.Create(ctx, &v1.CreateGameTownAgentConfig_Request{
		Name:           name,
		Provider:       v1enum.GameTownAgentProvider_GAME_TOWN_AGENT_PROVIDER_OLLAMA,
		BaseUrl:        conf.ollamaURL,
		Model:          conf.model,
		TimeoutSeconds: 90,
	})
	if err == nil {
		fmt.Printf("config created: %d\n", created.GetRow().GetId())
		return created.GetRow().GetId(), nil
	}

	listed, listErr := client.AgentConfig.List(ctx, &v1.ListGameTownAgentConfigs_Request{Page: &commonpb.PageReq{Page: 1, Size: 100}})
	if listErr != nil {
		return 0, fmt.Errorf("create config failed and list config failed: create=%w list=%w", err, listErr)
	}
	for _, row := range listed.GetRows() {
		if row.GetName() == name && row.GetBaseUrl() == conf.ollamaURL && row.GetModel() == conf.model {
			fmt.Printf("config reused: %d\n", row.GetId())
			return row.GetId(), nil
		}
	}
	return 0, fmt.Errorf("create config failed and reusable config not found: %w", err)
}

func registerPlayers(ctx context.Context, client *rpc.GameTownClient, count int) ([]playerState, error) {
	players := make([]playerState, 0, count)
	runID := time.Now().Format("20060102150405")
	for i := range count {
		name := fmt.Sprintf("load_player_%s_%02d", runID, i+1)
		displayName := fmt.Sprintf("Load Player %d", i+1)
		reply, err := client.Player.Register(ctx, &v1.RegisterGameTownPlayer_Request{Name: name, DisplayName: displayName})
		if err != nil {
			return nil, fmt.Errorf("register player %s failed: %w", name, err)
		}
		players = append(players, playerState{id: reply.GetRow().GetId(), name: name})
		fmt.Printf("player registered: %s id=%d\n", name, reply.GetRow().GetId())
	}
	return players, nil
}

func waitWorldReady(ctx context.Context, client *rpc.GameTownClient, state *runState, interval time.Duration) error {
	for {
		if err := collectEvents(ctx, client, state); err != nil {
			return err
		}
		reply, err := retryRPC(ctx, "get world", func() (*v1.GetGameTownWorld_Resp, error) {
			return client.World.Get(ctx, &v1.GetGameTownWorld_Request{Id: state.worldID})
		})
		if err != nil {
			return fmt.Errorf("get world status failed: %w", err)
		}
		switch reply.GetRow().GetStatus() {
		case v1enum.GameTownWorldStatus_GAME_TOWN_WORLD_STATUS_ACTIVE:
			fmt.Printf("world ready: %s\n", reply.GetRow().GetName())
			return nil
		case v1enum.GameTownWorldStatus_GAME_TOWN_WORLD_STATUS_FAILED:
			return errors.New("world generation failed")
		}
		if err := sleepContext(ctx, interval); err != nil {
			return err
		}
	}
}

func joinPlayers(ctx context.Context, client *rpc.GameTownClient, state *runState) error {
	for i := range state.players {
		preference := characterPreference(i)
		reply, err := retryRPC(ctx, "join world", func() (*v1.JoinGameTownWorld_Resp, error) {
			return client.WorldMember.Join(ctx, &v1.JoinGameTownWorld_Request{PlayerId: state.players[i].id, WorldCode: state.worldCode, CharacterPreference: &preference})
		})
		if err != nil {
			return fmt.Errorf("player %d join world failed: %w", state.players[i].id, err)
		}
		state.players[i].memberID = reply.GetMemberId()
		state.players[i].locationID = reply.GetLocationId()
		fmt.Printf("player joined: id=%d member=%d event=%d\n", state.players[i].id, reply.GetMemberId(), reply.GetEventId())
	}
	return collectEvents(ctx, client, state)
}

func startWatchers(ctx context.Context, wg *sync.WaitGroup, client *rpc.GameTownClient, state *runState, retryDelay time.Duration) {
	for _, player := range state.players {
		playerID := player.id
		wg.Add(1)
		go func() {
			defer wg.Done()
			watchPlayer(ctx, client, state, playerID, retryDelay)
		}()
	}
}

func watchPlayer(ctx context.Context, client *rpc.GameTownClient, state *runState, playerID int64, retryDelay time.Duration) {
	for ctx.Err() == nil {
		afterSequence := state.lastSequence(playerID)
		stream, err := client.Event.Watch(ctx, &v1.WatchGameTownEvents_Request{WorldId: state.worldID, PlayerId: playerID, AfterSequence: afterSequence})
		if err != nil {
			state.recordStreamError()
			if !retryableWatchError(ctx, err) {
				return
			}
			_ = sleepContext(ctx, retryDelay)
			continue
		}
		for ctx.Err() == nil {
			event, err := stream.Recv()
			if err != nil {
				state.recordStreamReconnect(err)
				if !retryableWatchError(ctx, err) {
					return
				}
				_ = sleepContext(ctx, retryDelay)
				break
			}
			state.recordStreamEvent(playerID, event.GetId(), event.GetSequence(), event.GetType())
		}
	}
}

func retryableWatchError(ctx context.Context, err error) bool {
	if err == nil || ctx.Err() != nil {
		return false
	}
	if errors.Is(err, io.EOF) {
		return true
	}
	return retryableRPCError(ctx, err)
}

func retryableRPCError(ctx context.Context, err error) bool {
	if err == nil || ctx.Err() != nil {
		return false
	}
	switch status.Code(err) {
	case codes.Canceled:
		return ctx.Err() == nil
	case codes.Unavailable, codes.DeadlineExceeded, codes.Unknown:
		return true
	}
	message := strings.ToLower(err.Error())
	return strings.Contains(message, "57p03") ||
		strings.Contains(message, "recovery mode") ||
		strings.Contains(message, "not yet accepting connections") ||
		strings.Contains(message, "rejecting connections")
}

func waitCharactersReady(ctx context.Context, client *rpc.GameTownClient, state *runState, interval time.Duration) error {
	for {
		if err := collectEvents(ctx, client, state); err != nil {
			return err
		}
		ready := 0
		for i := range state.players {
			member, err := retryRPC(ctx, "get world member", func() (*v1.GetGameTownWorldMember_Resp, error) {
				return client.WorldMember.Get(ctx, &v1.GetGameTownWorldMember_Request{WorldId: state.worldID, PlayerId: state.players[i].id})
			})
			if err != nil {
				return fmt.Errorf("get player character failed: %w", err)
			}
			state.players[i].locationID = member.GetCurrentLocationId()
			state.players[i].ready = member.GetCharacterReady()
			if state.players[i].ready {
				ready++
			}
		}
		if ready == len(state.players) {
			fmt.Println("all player characters ready")
			return nil
		}
		if err := sleepContext(ctx, interval); err != nil {
			return err
		}
	}
}

func playRounds(ctx context.Context, client *rpc.GameTownClient, state *runState, conf config) error {
	for round := 1; round <= conf.rounds; round++ {
		playerIndex := (round - 1) % len(state.players)
		player := state.players[playerIndex]
		content, targets, bigEvent, err := nextAction(ctx, client, state, player, round, conf)
		if err != nil {
			return err
		}
		_, err = retryRPC(ctx, "submit action", func() (*v1.SubmitGameTownAction_Resp, error) {
			return client.WorldMember.SubmitAction(ctx, &v1.SubmitGameTownAction_Request{
				WorldId:         state.worldID,
				PlayerId:        player.id,
				Content:         content,
				Targets:         targets,
				ClientRequestId: stringPtr(fmt.Sprintf("load-%d-%d", time.Now().UnixNano(), round)),
			})
		})
		if err != nil {
			return fmt.Errorf("submit action at round %d failed: %w", round, err)
		}
		state.recordSubmitted(player.id, bigEvent)
		if round%10 == 0 || bigEvent {
			if err := collectEvents(ctx, client, state); err != nil {
				return err
			}
			stats := state.snapshotStats()
			fmt.Printf("round=%d submitted=%d world_evolved=%d npc_replied=%d failed=%d stream=%d\n", round, stats.submittedActions, stats.worldEvolved, stats.npcReplies, stats.jobFailed, stats.streamEvents)
		}
		if err := sleepContext(ctx, 150*time.Millisecond); err != nil {
			return err
		}
	}
	return waitAsyncSettled(ctx, client, state, conf)
}

func nextAction(ctx context.Context, client *rpc.GameTownClient, state *runState, player playerState, round int, conf config) (string, []*v1.SubmitGameTownAction_Request_EntityRef, bool, error) {
	if round <= len(state.players) {
		content, targets, ok, err := npcTalkAction(ctx, client, state, player)
		if err != nil {
			return "", nil, false, err
		}
		if ok {
			return content, targets, false, nil
		}
	}
	if conf.bigEventEvery > 0 && round%conf.bigEventEvery == 0 {
		return bigEvent(round), nil, true, nil
	}
	if shouldTalkToNpc(round) {
		content, targets, ok, err := npcTalkAction(ctx, client, state, player)
		if err != nil {
			return "", nil, false, err
		}
		if ok {
			return content, targets, false, nil
		}
	}
	if round%25 == 0 {
		return smallAction(round), nil, false, nil
	}
	content, targets, ok, err := moveAction(ctx, client, state, player)
	if err != nil {
		return "", nil, false, err
	}
	if ok {
		return content, targets, false, nil
	}
	return smallAction(round), nil, false, nil
}

func moveAction(ctx context.Context, client *rpc.GameTownClient, state *runState, player playerState) (string, []*v1.SubmitGameTownAction_Request_EntityRef, bool, error) {
	locations, err := retryRPC(ctx, "list locations", func() (*v1.ListGameTownLocations_Resp, error) {
		return client.Location.List(ctx, &v1.ListGameTownLocations_Request{WorldId: state.worldID, PlayerId: player.id})
	})
	if err != nil {
		return "", nil, false, fmt.Errorf("list locations failed: %w", err)
	}
	for _, location := range locations.GetRows() {
		if location.GetAccessible() && !location.GetCurrent() {
			content := "I travel to " + location.GetName() + " and watch how the route has changed."
			targets := []*v1.SubmitGameTownAction_Request_EntityRef{
				{
					Type: v1enum.GameTownEntityType_GAME_TOWN_ENTITY_TYPE_LOCATION,
					Id:   location.GetId(),
				},
			}
			return content, targets, true, nil
		}
	}
	return "", nil, false, nil
}

func shouldTalkToNpc(round int) bool {
	return round%40 == 0
}

func npcTalkAction(ctx context.Context, client *rpc.GameTownClient, state *runState, player playerState) (string, []*v1.SubmitGameTownAction_Request_EntityRef, bool, error) {
	npcs, err := retryRPC(ctx, "list npcs", func() (*v1.ListGameTownNpcs_Resp, error) {
		return client.Npc.List(ctx, &v1.ListGameTownNpcs_Request{WorldId: state.worldID, PlayerId: player.id, LocationId: &player.locationID})
	})
	if err != nil {
		return "", nil, false, fmt.Errorf("list NPCs failed: %w", err)
	}
	if len(npcs.GetRows()) == 0 {
		npcs, err = retryRPC(ctx, "list npcs", func() (*v1.ListGameTownNpcs_Resp, error) {
			return client.Npc.List(ctx, &v1.ListGameTownNpcs_Request{WorldId: state.worldID, PlayerId: player.id})
		})
		if err != nil {
			return "", nil, false, fmt.Errorf("list NPCs failed: %w", err)
		}
	}
	if len(npcs.GetRows()) == 0 {
		return "", nil, false, nil
	}
	npc := npcs.GetRows()[rand.IntN(len(npcs.GetRows()))]
	content := fmt.Sprintf("I ask %s: what recent sign could change the balance of power?", npc.GetName())
	targets := []*v1.SubmitGameTownAction_Request_EntityRef{{Type: v1enum.GameTownEntityType_GAME_TOWN_ENTITY_TYPE_NPC, Id: npc.GetId()}}
	return content, targets, true, nil
}

func waitAsyncSettled(ctx context.Context, client *rpc.GameTownClient, state *runState, conf config) error {
	if !conf.strict {
		for i := 0; i < 3; i++ {
			if err := collectEvents(ctx, client, state); err != nil {
				return err
			}
			if err := sleepContext(ctx, conf.pollInterval); err != nil {
				return err
			}
		}
		return nil
	}

	deadline := time.Now().Add(strictSettleTimeout(conf))
	nextReport := time.Now().Add(30 * time.Second)
	for {
		if err := collectEvents(ctx, client, state); err != nil {
			return err
		}
		now := time.Now()
		if now.After(nextReport) {
			stats := state.snapshotStats()
			fmt.Printf(
				"settle submitted=%d feedback=%d/%d world_evolved=%d/%d world_evolution=%d/%d npc_replied=%d failed=%d stream=%d\n",
				stats.submittedActions,
				feedbackCount(stats),
				minTotalFeedback(conf),
				stats.worldEvolved,
				minWorldEvolved(conf),
				worldEvolutionCount(stats),
				minWorldEvolution(conf),
				stats.npcReplies,
				stats.jobFailed,
				stats.streamEvents,
			)
			nextReport = now.Add(30 * time.Second)
		}
		if strictAsyncSatisfied(state.snapshotStats(), state.snapshotPlayerStats(), conf) {
			return nil
		}
		if now.After(deadline) || ctx.Err() != nil {
			stats := state.snapshotStats()
			return fmt.Errorf(
				"strict async settle timeout: feedback=%d/%d world_evolution=%d/%d world_evolved=%d/%d npc_replied=%d stream_events=%d",
				feedbackCount(stats),
				minTotalFeedback(conf),
				worldEvolutionCount(stats),
				minWorldEvolution(conf),
				stats.worldEvolved,
				minWorldEvolved(conf),
				stats.npcReplies,
				stats.streamEvents,
			)
		}
		if err := sleepContext(ctx, conf.pollInterval); err != nil {
			return err
		}
	}
}

func strictAsyncSatisfied(stats runStats, playerStats map[int64]playerRunStats, conf config) bool {
	if stats.characterReady < conf.players || stats.streamEvents == 0 {
		return false
	}
	if stats.jobFailed > 0 {
		return false
	}
	if conf.bigEventEvery > 0 && stats.bigEvents == 0 {
		return false
	}
	if stats.npcReplies == 0 {
		return false
	}
	if feedbackCount(stats) < minTotalFeedback(conf) {
		return false
	}
	if minEvolution := minWorldEvolution(conf); minEvolution > 0 && worldEvolutionCount(stats) < minEvolution {
		return false
	}
	if minEvolved := minWorldEvolved(conf); minEvolved > 0 && stats.worldEvolved < minEvolved {
		return false
	}
	if stats.resolved+stats.rejected+stats.clarification == 0 {
		return false
	}
	return validatePlayerStats(playerStats, conf) == nil
}

func strictSettleTimeout(conf config) time.Duration {
	if conf.settleTimeout > 0 {
		return conf.settleTimeout
	}
	if conf.rounds >= 1000 {
		return 2 * time.Hour
	}
	if conf.rounds >= 300 {
		return 45 * time.Minute
	}
	if conf.rounds >= 100 {
		return 20 * time.Minute
	}
	return 5 * time.Minute
}

func feedbackCount(stats runStats) int {
	return stats.npcReplies + stats.resolved + stats.rejected + stats.clarification
}

func worldEvolutionCount(stats runStats) int {
	return stats.worldEvolved + stats.npcPlanned + stats.worldTicks
}

func minTotalFeedback(conf config) int {
	value := conf.rounds / 4
	if conf.rounds >= 1000 {
		value = conf.rounds / 3
	}
	if value < conf.players {
		return conf.players
	}
	return value
}

func minPlayerFeedback(player playerRunStats, conf config) int {
	value := player.submitted / 4
	if conf.rounds >= 1000 {
		value = player.submitted / 3
	}
	if value < 1 {
		return 1
	}
	return value
}

func minWorldEvolution(conf config) int {
	if conf.rounds >= 1000 {
		value := conf.rounds / 100
		if value < 10 {
			return 10
		}
		return value
	}
	if conf.rounds >= 300 {
		return 3
	}
	if conf.rounds >= 100 {
		return 1
	}
	return 0
}

func minWorldEvolved(conf config) int {
	if conf.rounds >= 1000 {
		return 5
	}
	if conf.rounds >= 100 {
		return 1
	}
	return 0
}

func collectEvents(ctx context.Context, client *rpc.GameTownClient, state *runState) error {
	for _, player := range state.players {
		reply, err := pageEventsWithRetry(ctx, client, state, player.id)
		if err != nil {
			return fmt.Errorf("page events failed: %w", err)
		}
		for _, event := range reply.GetRows() {
			state.recordPageEvent(player.id, event.GetId(), event.GetSequence(), event.GetType())
		}
	}
	return nil
}

func retryRPC[T any](ctx context.Context, operation string, call func() (T, error)) (T, error) {
	var zero T
	var lastErr error
	for attempt := 1; attempt <= 10; attempt++ {
		resp, err := call()
		if err == nil {
			return resp, nil
		}
		lastErr = err
		if !retryableRPCError(ctx, err) {
			return zero, fmt.Errorf("%s failed: %w", operation, err)
		}
		if sleepErr := sleepContext(ctx, time.Duration(attempt)*time.Second); sleepErr != nil {
			return zero, errors.Join(lastErr, sleepErr)
		}
	}
	return zero, fmt.Errorf("%s failed after retries: %w", operation, lastErr)
}
func pageEventsWithRetry(ctx context.Context, client *rpc.GameTownClient, state *runState, playerID int64) (*v1.PageGameTownEvents_Resp, error) {
	var lastErr error
	for attempt := 1; attempt <= 10; attempt++ {
		afterSequence := state.lastSequence(playerID)
		reply, err := client.Event.Page(ctx, &v1.PageGameTownEvents_Request{
			WorldId:       state.worldID,
			PlayerId:      playerID,
			AfterSequence: afterSequence,
			Page:          &commonpb.PageReq{Page: 1, Size: 100},
		})
		if err == nil {
			return reply, nil
		}
		lastErr = err
		if !retryableRPCError(ctx, err) {
			return nil, err
		}
		if sleepErr := sleepContext(ctx, time.Duration(attempt)*time.Second); sleepErr != nil {
			return nil, errors.Join(lastErr, sleepErr)
		}
	}
	return nil, lastErr
}

func (s *runState) lastSequence(playerID int64) uint64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.playerSeq[playerID]
}

func (s *runState) initPlayerStats(players []playerState) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, player := range players {
		if s.playerStats[player.id] == nil {
			s.playerStats[player.id] = &playerRunStats{}
		}
	}
}

func (s *runState) recordSubmitted(playerID int64, bigEvent bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.stats.submittedActions++
	s.ensurePlayerStatsLocked(playerID).submitted++
	if bigEvent {
		s.stats.bigEvents++
	}
}

func (s *runState) recordPageEvent(playerID int64, eventID int64, sequence uint64, eventType v1enum.GameTownEventType) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.stats.pageEvents++
	s.ensurePlayerStatsLocked(playerID).pageEvents++
	s.recordEventLocked(playerID, eventID, sequence, eventType)
}

func (s *runState) recordStreamEvent(playerID int64, eventID int64, sequence uint64, eventType v1enum.GameTownEventType) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.stats.streamEvents++
	s.ensurePlayerStatsLocked(playerID).streamEvents++
	s.recordEventLocked(playerID, eventID, sequence, eventType)
}

func (s *runState) recordStreamError() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.stats.streamErrors++
}

func (s *runState) recordStreamReconnect(err error) {
	if err == nil || errors.Is(err, context.Canceled) {
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.stats.streamReconnects++
}

func (s *runState) recordEventLocked(playerID int64, eventID int64, sequence uint64, eventType v1enum.GameTownEventType) {
	if sequence <= s.playerSeq[playerID] {
		return
	}
	s.playerSeq[playerID] = sequence
	playerStats := s.ensurePlayerStatsLocked(playerID)
	playerStats.visibleEvents++
	countPlayerEvent(playerStats, eventType)
	if s.seenEvent[eventID] {
		return
	}
	s.seenEvent[eventID] = true
	countEvent(&s.stats, eventType)
}

func (s *runState) ensurePlayerStatsLocked(playerID int64) *playerRunStats {
	if s.playerStats == nil {
		s.playerStats = make(map[int64]*playerRunStats)
	}
	stats := s.playerStats[playerID]
	if stats == nil {
		stats = &playerRunStats{}
		s.playerStats[playerID] = stats
	}
	return stats
}

func (s *runState) snapshotStats() runStats {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.stats
}

func (s *runState) snapshotPlayerStats() map[int64]playerRunStats {
	s.mu.Lock()
	defer s.mu.Unlock()
	result := make(map[int64]playerRunStats, len(s.playerStats))
	for playerID, stats := range s.playerStats {
		if stats == nil {
			continue
		}
		result[playerID] = *stats
	}
	return result
}

func countEvent(stats *runStats, eventType v1enum.GameTownEventType) {
	switch eventType {
	case v1enum.GameTownEventType_GAME_TOWN_EVENT_TYPE_NPC_REPLIED:
		stats.npcReplies++
	case v1enum.GameTownEventType_GAME_TOWN_EVENT_TYPE_WORLD_EVOLVED:
		stats.worldEvolved++
	case v1enum.GameTownEventType_GAME_TOWN_EVENT_TYPE_NPC_PLANNED:
		stats.npcPlanned++
	case v1enum.GameTownEventType_GAME_TOWN_EVENT_TYPE_WORLD_TICK_REQUESTED:
		stats.worldTicks++
	case v1enum.GameTownEventType_GAME_TOWN_EVENT_TYPE_ACTION_RESOLVED:
		stats.resolved++
	case v1enum.GameTownEventType_GAME_TOWN_EVENT_TYPE_ACTION_REJECTED:
		stats.rejected++
	case v1enum.GameTownEventType_GAME_TOWN_EVENT_TYPE_ACTION_CLARIFICATION_REQUIRED:
		stats.clarification++
	case v1enum.GameTownEventType_GAME_TOWN_EVENT_TYPE_AGENT_JOB_FAILED:
		stats.jobFailed++
	case v1enum.GameTownEventType_GAME_TOWN_EVENT_TYPE_PLAYER_CHARACTER_READY:
		stats.characterReady++
	}
}

func countPlayerEvent(stats *playerRunStats, eventType v1enum.GameTownEventType) {
	switch eventType {
	case v1enum.GameTownEventType_GAME_TOWN_EVENT_TYPE_NPC_REPLIED:
		stats.npcReplies++
	case v1enum.GameTownEventType_GAME_TOWN_EVENT_TYPE_ACTION_RESOLVED,
		v1enum.GameTownEventType_GAME_TOWN_EVENT_TYPE_ACTION_REJECTED,
		v1enum.GameTownEventType_GAME_TOWN_EVENT_TYPE_ACTION_CLARIFICATION_REQUIRED:
		stats.actionResults++
	case v1enum.GameTownEventType_GAME_TOWN_EVENT_TYPE_PLAYER_CHARACTER_READY:
		stats.characterReady = true
	}
}

func validateRun(state *runState, conf config) error {
	stats := state.snapshotStats()
	if stats.submittedActions != conf.rounds {
		return fmt.Errorf("submitted action count mismatch: got=%d want=%d", stats.submittedActions, conf.rounds)
	}
	if !conf.strict {
		return nil
	}
	if stats.characterReady < conf.players {
		return fmt.Errorf("strict validation failed: character_ready=%d players=%d", stats.characterReady, conf.players)
	}
	if stats.streamEvents == 0 {
		return errors.New("strict validation failed: EventService.Watch produced no events")
	}
	if stats.jobFailed > 0 {
		return fmt.Errorf("strict validation failed: agent_job_failed=%d", stats.jobFailed)
	}
	if conf.bigEventEvery > 0 && stats.bigEvents == 0 {
		return errors.New("strict validation failed: no major event submitted")
	}
	if conf.rounds >= 1000 && stats.bigEvents < 2 {
		return fmt.Errorf("strict validation failed: big_events=%d, expected at least 2 for 1000 rounds", stats.bigEvents)
	}
	if stats.npcReplies == 0 {
		return errors.New("strict validation failed: no npc_replied event observed")
	}
	if stats.resolved+stats.rejected+stats.clarification == 0 {
		return errors.New("strict validation failed: no action result event observed")
	}
	if feedback := feedbackCount(stats); feedback < minTotalFeedback(conf) {
		return fmt.Errorf("strict validation failed: feedback=%d expected>=%d", feedback, minTotalFeedback(conf))
	}
	if evolution := worldEvolutionCount(stats); evolution < minWorldEvolution(conf) {
		return fmt.Errorf("strict validation failed: world evolution events=%d expected>=%d", evolution, minWorldEvolution(conf))
	}
	if stats.worldEvolved < minWorldEvolved(conf) {
		return fmt.Errorf("strict validation failed: world_evolved=%d expected>=%d", stats.worldEvolved, minWorldEvolved(conf))
	}
	return validatePlayerStats(state.snapshotPlayerStats(), conf)
}

func validatePlayerStats(stats map[int64]playerRunStats, conf config) error {
	if len(stats) < conf.players {
		return fmt.Errorf("strict validation failed: player stats count=%d players=%d", len(stats), conf.players)
	}
	minVisibleEvents := 1
	if conf.rounds >= 1000 {
		minVisibleEvents = conf.rounds / conf.players / 4
	}
	for playerID, player := range stats {
		if player.submitted == 0 {
			return fmt.Errorf("strict validation failed: player %d submitted no action", playerID)
		}
		if player.pageEvents == 0 {
			return fmt.Errorf("strict validation failed: player %d has no page events", playerID)
		}
		if player.streamEvents == 0 {
			return fmt.Errorf("strict validation failed: player %d has no stream events", playerID)
		}
		if player.visibleEvents < minVisibleEvents {
			return fmt.Errorf("strict validation failed: player %d visible_events=%d expected>=%d", playerID, player.visibleEvents, minVisibleEvents)
		}
		if !player.characterReady {
			return fmt.Errorf("strict validation failed: player %d character is not ready", playerID)
		}
		minFeedback := minPlayerFeedback(player, conf)
		feedback := player.actionResults + player.npcReplies
		if feedback < minFeedback {
			return fmt.Errorf("strict validation failed: player %d feedback=%d expected>=%d", playerID, feedback, minFeedback)
		}
	}
	return nil
}

func newGameTownClient(conf config) (*rpc.GameTownClient, func(), string, error) {
	if conf.addr != "" {
		conn, err := grpc.NewClient(conf.addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return nil, nil, conf.addr, err
		}
		return rpc.NewGameTownClient(conn), func() { _ = conn.Close() }, conf.addr, nil
	}
	consul, cleanup, err := commonClient.NewConsulClient(slog.Default(), &commonpb.Consul{Address: conf.consulAddr, Datacenter: conf.consulDatacenter, Token: conf.consulToken, DialTimeout: durationpb.New(5 * time.Second)}, nil)
	if err != nil {
		return nil, nil, conf.consulAddr, err
	}
	client, err := rpc.ProvideGameTownClient(consul)
	if err != nil {
		cleanup()
		return nil, nil, conf.consulAddr, err
	}
	return client, cleanup, "consul://" + conf.consulAddr + "/" + constant.GameTownServiceName.String(), nil
}

func printReport(state *runState) {
	stats := state.snapshotStats()
	playerStats := state.snapshotPlayerStats()
	fmt.Println("\nGame Town load test report")
	fmt.Printf("world_id=%d world_code=%s config_id=%d\n", state.worldID, state.worldCode, state.configID)
	fmt.Printf("players=%d submitted_actions=%d big_events=%d page_events=%d stream_events=%d\n", len(state.players), stats.submittedActions, stats.bigEvents, stats.pageEvents, stats.streamEvents)
	fmt.Printf("stream_reconnects=%d stream_errors=%d npc_replied=%d world_evolved=%d npc_planned=%d world_ticks=%d character_ready=%d\n", stats.streamReconnects, stats.streamErrors, stats.npcReplies, stats.worldEvolved, stats.npcPlanned, stats.worldTicks, stats.characterReady)
	fmt.Printf("action_resolved=%d action_rejected=%d clarification=%d agent_job_failed=%d\n", stats.resolved, stats.rejected, stats.clarification, stats.jobFailed)
	for _, player := range state.players {
		stats := playerStats[player.id]
		fmt.Printf(
			"player id=%d submitted=%d page_events=%d stream_events=%d visible_events=%d npc_replies=%d action_results=%d character_ready=%t\n",
			player.id,
			stats.submitted,
			stats.pageEvents,
			stats.streamEvents,
			stats.visibleEvents,
			stats.npcReplies,
			stats.actionResults,
			stats.characterReady,
		)
	}
	fmt.Println("Acceptance hint: run with -rounds 1000 -strict and inspect world_evolved, npc_replied, major events, failures, and stream counters.")
}

func defaultWorldDescription() string {
	return "A floating-islands steam fantasy world. Broken continents drift through storm routes. Cloud whales, a machine church, airship guilds, and exiled royal houses compete for wind cores that can stabilize the islands. Recently the sky rift has widened, some islands have started to fall, and the sleeping cloud whales are sending warnings. The world evolves even without player input; people can grow, get hurt, ally, betray, disappear, or die, and locations and factions can change through disasters, wars, trade, and secret operations."
}

func characterPreference(index int) string {
	values := []string{
		"I want to play an airship mechanic who can hear cloud whale songs and read omens in engine noise.",
		"I want to play a young courier from the exiled royal houses who carries a secret wind-core diagram.",
		"I want to play a runaway apprentice from the machine church who knows a forbidden rite but not its full cost.",
	}
	return values[index%len(values)]
}

func smallAction(round int) string {
	values := []string{
		"I gather rumors nearby and record which faction has moved strangely.",
		"I observe the sky rift and try to predict the next dangerous area.",
		"I help locals repair equipment while asking about wind cores and cloud whales.",
		"I search for evidence that one faction is hiding the truth.",
	}
	return values[round%len(values)]
}

func bigEvent(round int) string {
	values := []string{
		"Major event: an outer floating island starts falling, refugees flood into the main harbor, and the airship guild demands control of the route.",
		"Major event: the machine church announces a new wind-core fragment and locks down the news, triggering conflict with royal exiles and merchants.",
		"Major event: the cloud whale herd changes course at night, the sky rift widens, and an important NPC disappears during the investigation.",
		"Major event: a storm destroys a key supply location, forcing factions to renew alliances or betray each other openly.",
	}
	return values[(round/100)%len(values)]
}

func stringPtr(value string) *string {
	return &value
}

func sleepContext(ctx context.Context, duration time.Duration) error {
	timer := time.NewTimer(duration)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

func envDefault(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func fatal(format string, args ...any) {
	_, _ = fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}
