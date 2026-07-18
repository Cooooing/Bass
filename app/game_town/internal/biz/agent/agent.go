package agent

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Runner interface {
	GenerateWorld(ctx context.Context, input *GenerateWorldInput) (*GenerateWorldOutput, error)
	Talk(ctx context.Context, input *TalkInput) (*TalkOutput, error)
	Direct(ctx context.Context, input *DirectInput) (*DirectOutput, error)
}

type RunConfig struct {
	ID             int64
	Provider       string
	Model          string
	BaseURL        string
	APIKey         string
	TimeoutSeconds int32
}

type GenerateWorldInput struct {
	Config        *RunConfig
	Description   string
	NpcCount      uint32
	LocationCount uint32
	Scale         string
	Seed          int64
	StyleTags     []string
}

type GeneratedLocation struct {
	Code        string
	Name        string
	Description string
	Tags        map[string]any
}

type GeneratedNpc struct {
	Code         string
	Name         string
	Role         string
	Personality  string
	Goal         string
	Background   string
	LocationCode string
	SystemPrompt string
	Profile      map[string]any
}

type GeneratedMetric struct {
	Key          string
	Name         string
	Description  string
	MinValue     int32
	MaxValue     int32
	InitialValue int32
}

type GenerateWorldOutput struct {
	WorldName       string
	WorldSummary    string
	Locations       []GeneratedLocation
	Npcs            []GeneratedNpc
	Metrics         []GeneratedMetric
	InitialMetrics  map[string]any
	CurrentArc      string
	OpeningEvents   []string
	RawProviderName string
}

type TalkInput struct {
	Config         *RunConfig
	PlayerName     string
	WorldName      string
	LocationName   string
	NpcName        string
	NpcRole        string
	NpcPersonality string
	Relationship   map[string]any
	Memories       []string
	WorldMetrics   map[string]any
	Content        string
}

type TalkOutput struct {
	Reply             string
	MemoryCandidates  []string
	RelationshipDelta map[string]int32
	WorldMetricDelta  map[string]int32
	Events            []string
}

type DirectInput struct {
	Config    *RunConfig
	WorldName string
	Arc       string
	Metrics   map[string]any
	Events    []string
}

type DirectOutput struct {
	Summary          string
	CurrentArc       string
	WorldMetricDelta map[string]int32
	Events           []string
}

type BladesRunner struct{}

func NewBladesRunner() Runner {
	return &BladesRunner{}
}

func (r *BladesRunner) GenerateWorld(ctx context.Context, input *GenerateWorldInput) (*GenerateWorldOutput, error) {
	_ = ctx
	seed := input.Seed
	if seed == 0 {
		seed = time.Now().UnixNano()
	}
	rnd := rand.New(rand.NewSource(seed))
	locationCount := int(input.LocationCount)
	if locationCount <= 0 {
		locationCount = 4
	}
	npcCount := int(input.NpcCount)
	if npcCount <= 0 {
		npcCount = 4
	}
	base := strings.TrimSpace(input.Description)
	if base == "" {
		base = "一个由玩家共同塑造的文字世界"
	}
	worldName := strings.TrimSpace(base)
	if len([]rune(worldName)) > 18 {
		worldName = string([]rune(worldName)[:18])
	}
	locations := make([]GeneratedLocation, 0, locationCount)
	for i := 0; i < locationCount; i++ {
		code := fmt.Sprintf("loc_%02d", i+1)
		locations = append(locations, GeneratedLocation{Code: code, Name: fmt.Sprintf("区域%d", i+1), Description: fmt.Sprintf("%s中的第%d个关键地点。", worldName, i+1), Tags: map[string]any{"scale": input.Scale}})
	}
	npcs := make([]GeneratedNpc, 0, npcCount)
	for i := 0; i < npcCount; i++ {
		loc := locations[rnd.Intn(len(locations))]
		code := fmt.Sprintf("npc_%02d", i+1)
		name := fmt.Sprintf("角色%d", i+1)
		npcs = append(npcs, GeneratedNpc{Code: code, Name: name, Role: "世界居民", Personality: "会根据玩家行为调整态度", Goal: "推动世界事件演化", Background: fmt.Sprintf("%s中的重要 NPC。", worldName), LocationCode: loc.Code, SystemPrompt: fmt.Sprintf("你是%s，属于%s。", name, worldName), Profile: map[string]any{"generated": true}})
	}
	metrics := []GeneratedMetric{
		{Key: "stability", Name: "稳定度", Description: "世界秩序和结构稳定程度", MinValue: 0, MaxValue: 100, InitialValue: 60},
		{Key: "activity", Name: "活跃度", Description: "玩家和 NPC 的行动活跃程度", MinValue: 0, MaxValue: 100, InitialValue: 40},
		{Key: "tension", Name: "紧张度", Description: "世界内部冲突压力", MinValue: 0, MaxValue: 100, InitialValue: 30},
	}
	initial := map[string]any{"stability": float64(60), "activity": float64(40), "tension": float64(30)}
	provider := "blades"
	if input.Config != nil && strings.TrimSpace(input.Config.Provider) != "" {
		provider = strings.TrimSpace(input.Config.Provider)
	}
	return &GenerateWorldOutput{WorldName: worldName, WorldSummary: base, Locations: locations, Npcs: npcs, Metrics: metrics, InitialMetrics: initial, CurrentArc: "世界刚刚生成，玩家的行动会决定后续走向", OpeningEvents: []string{"世界被创建"}, RawProviderName: provider}, nil
}

func (r *BladesRunner) Talk(ctx context.Context, input *TalkInput) (*TalkOutput, error) {
	_ = ctx
	reply := fmt.Sprintf("%s回应：我理解你提到的“%s”。这件事会影响%s当前的发展。", input.NpcName, input.Content, input.WorldName)
	return &TalkOutput{Reply: reply, MemoryCandidates: []string{fmt.Sprintf("玩家提到：%s", input.Content)}, RelationshipDelta: map[string]int32{"trust": 1}, WorldMetricDelta: map[string]int32{"activity": 1}, Events: []string{fmt.Sprintf("%s与玩家完成了一次对话", input.NpcName)}}, nil
}

func (r *BladesRunner) Direct(ctx context.Context, input *DirectInput) (*DirectOutput, error) {
	_ = ctx
	return &DirectOutput{Summary: "世界根据近期事件完成一次演化", CurrentArc: input.Arc, WorldMetricDelta: map[string]int32{"activity": 1, "tension": -1}, Events: []string{"世界状态发生了细微变化"}}, nil
}
