package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"game_town/internal/biz/model"
	"game_town/internal/biz/repo"
	"game_town/internal/enum"

	"github.com/samber/lo"
)

type AgentClient struct {
	httpClient *http.Client
}

func NewAgentClient() repo.AgentClient {
	return &AgentClient{
		httpClient: &http.Client{},
	}
}

func (c *AgentClient) GenerateWorld(ctx context.Context, req *repo.GenerateWorldReq) (*model.WorldDraft, error) {
	prompt := fmt.Sprintf(`请根据玩家描述创建一个可自行演化的文字游戏世界。

玩家描述：%s

输出要求：
- 只输出合法 JSON，不要输出 Markdown。
- locations 必须正好 %d 个，npcs 必须正好 %d 个，factions 2 到 3 个。
- 所有 code 使用唯一小写 ASCII。
- 内容使用简体中文，但必须短：name 12 字内，summary 120 字内，current_arc 40 字内。
- description、personality、goal、background、system_prompt 每项 60 字内。
- rules 只写最小结构，不要展开长篇规则说明。

JSON 结构：
{
  "name": "",
  "summary": "",
  "current_arc": "",
  "current_era": "",
  "rules": {
    "calendar": {},
    "species": [],
    "relationship_dimensions": [],
    "actions": [],
    "hazards": []
  },
  "locations": [
    {"code": "", "name": "", "description": "", "environment_tags": []}
  ],
  "factions": [
    {"code": "", "name": "", "description": "", "public_goal": ""}
  ],
  "npcs": [
    {
      "code": "",
      "name": "",
      "role": "",
      "species": "",
      "personality": "",
      "goal": "",
      "background": "",
      "location_code": "",
      "faction_code": "",
      "system_prompt": "",
      "attributes": {}
    }
  ]
}`,
		req.World.Description,
		req.LocationCount,
		req.NpcCount,
	)

	output := new(model.WorldDraft)
	if err := c.completeJSON(ctx, req.Config, "你负责生成有因果规则、时间和社会结构的活世界，只输出合法 JSON。所有字符串必须使用简体中文。", prompt, c.worldDraftSchema(req.LocationCount, req.NpcCount), output, 1024); err != nil {
		return nil, err
	}
	return output, nil
}

func (c *AgentClient) GenerateCharacter(ctx context.Context, req *repo.GenerateCharacterReq) (*model.CharacterDraft, error) {
	preference := strings.TrimSpace(req.Preference)
	if preference == "" {
		preference = "由世界根据当前局势安排一个合理的初始身份"
	}
	prompt := fmt.Sprintf(`请根据世界设定为玩家账号生成一个属于该世界的初始角色。

世界：%s
公开背景：%s
当前阶段：%s
初始地点：%s - %s
玩家账号名：%s
玩家角色倾向：%s
最近公开事件：%s

只输出合法 JSON，不要输出 Markdown。
JSON 结构：{"name":"","background":"","goal":"","traits":[]}

规则：
- name 必须是世界内角色姓名，不能使用玩家账号名。
- 角色必须属于这个世界，而不是通用现代玩家。
- 玩家倾向只是参考，最终身份由世界因果、阵营、地点和当前危机裁决。
- background 120 字内，要给出可扮演的出身、处境、与当前世界冲突的连接。
- goal 60 字内，是角色进入世界后的初始目标，不要替玩家决定长期命运。
- traits 返回 2 到 6 个简短标签。`,
		req.World.Name,
		req.State.PublicChronicle,
		req.State.CurrentArc,
		req.Location.Name,
		req.Location.Description,
		req.Player.DisplayName,
		preference,
		c.eventContext(req.RecentEvents),
	)

	output := new(model.CharacterDraft)
	if err := c.completeJSON(ctx, req.Config, "你负责根据活世界状态裁决玩家在该世界中的初始角色，只输出合法 JSON。所有字符串必须使用简体中文。", prompt, c.characterDraftSchema(), output, 384); err != nil {
		return nil, err
	}
	return output, nil
}

func (c *AgentClient) Talk(ctx context.Context, req *repo.TalkReq) (*model.NpcReply, error) {
	prompt := fmt.Sprintf(`你正在扮演一个独立 NPC。你不是全知者，只能依据提供给你的观察、记忆和公开世界信息回答。

公开世界背景：%s
当前位置：%s - %s
NPC：%s，%s，%s，性格：%s
当前目标：%s
私有摘要：%s
近期可感知事件：%s
可召回私有记忆：%s
玩家角色：%s
玩家输入：%s

只输出合法 JSON：{"reply":"","context_summary":"","suggested_actions":[{"label":"","content":"","targets":[{"type":"npc","id":0}]}],"claims":[]}

规则：
- suggested_actions 返回 0 到 4 项，只是玩家可选快捷回答。
- reply 120 字内，context_summary 80 字内，每条建议 30 字内。
- 玩家始终可以自由输入，建议回答不能限制玩家。
- 不要假设你知道未出现在上下文中的对话、秘密或事实。
- 如果信息不足，可以按 NPC 性格表达不知道、怀疑或要求玩家说明。`,
		req.State.PublicChronicle,
		req.Location.Name,
		req.Location.Description,
		req.Npc.Name,
		req.Npc.Role,
		req.Npc.Species,
		req.Npc.Personality,
		req.Npc.Goal,
		req.Npc.ContextSummary,
		c.eventContext(req.RecentEvents),
		c.memoryContext(req.Memories),
		c.characterName(req.Member, req.Player),
		req.Content,
	)

	systemPrompt := req.Npc.SystemPrompt + "\n你不是全知者，只能依据提供给你的观察和记忆回答。所有字符串必须使用简体中文。"
	output := new(model.NpcReply)
	if err := c.completeJSON(ctx, req.Config, systemPrompt, prompt, c.npcReplySchema(), output, 384); err != nil {
		return nil, err
	}
	return output, nil
}

func (c *AgentClient) Act(ctx context.Context, req *repo.ActReq) (*model.ActionResolution, error) {
	target := "无明确 NPC 目标"
	systemPrompt := "你负责解释玩家的自由行动，并提出可由世界规则校验的结构化结果，只输出合法 JSON。所有字符串必须使用简体中文。"
	if req.Npc != nil {
		target = fmt.Sprintf("%s（%s），当前目标：%s，私有摘要：%s", req.Npc.Name, req.Npc.Role, req.Npc.Goal, req.Npc.ContextSummary)
		systemPrompt = req.Npc.SystemPrompt + "\n你只知道提供的观察和私有记忆。请先以该 NPC 的立场回应玩家行动，再提出结构化结果。"
	}

	prompt := fmt.Sprintf(`公开世界背景：%s
当前阶段：%s
当前位置：%s - %s
玩家角色：%s
目标 NPC：%s
近期可用信息：%s
目标 NPC 可召回记忆：%s
玩家行动：%s

只输出合法 JSON：
{
  "status": "resolved|rejected|clarification",
  "summary": "",
  "clarification": "",
  "world_summary": "",
  "current_arc": "",
  "actions": [
    {
      "type": "move|move_player|move_npc|change_npc_state|change_location|change_faction|change_relationship|share_claim",
      "target": {"type": "npc|location|faction|player", "id": 0},
      "parameters": {},
      "duration_minutes": 0
    }
  ],
  "claims": [],
  "suggested_actions": [
    {"label": "", "content": "", "targets": []}
  ]
}

规则：
- 你只能提出动作，不能伪造 ID、权限、能力或已经发生的结果。
- 如果信息不足，status 使用 clarification，并填写 clarification。
- summary 80 字内，clarification 60 字内，world_summary 120 字内，current_arc 40 字内。
- actions 最多 3 项，claims 最多 3 项，suggested_actions 最多 3 项。
- 玩家可以做任何想做的事，但世界、NPC、地点、阵营会根据规则给出成功、失败或代价。
- 如果改变 NPC、地点、阵营、关系，请使用 actions 描述，不要只写在 summary 里。`,
		req.State.PublicChronicle,
		req.State.CurrentArc,
		req.Location.Name,
		req.Location.Description,
		c.characterName(req.Member, req.Player),
		target,
		c.eventContext(req.RecentEvents),
		c.memoryContext(req.Memories),
		req.Content,
	)

	output := new(model.ActionResolution)
	if err := c.completeJSON(ctx, req.Config, systemPrompt, prompt, c.actionResolutionSchema(), output, 384); err != nil {
		return nil, err
	}
	return output, nil
}

func (c *AgentClient) PlanNpc(ctx context.Context, req *repo.PlanNpcReq) (*model.NpcPlan, error) {
	prompt := fmt.Sprintf(`公开世界背景：%s
世界时间：%s
当前位置：%s - %s
NPC：%s，%s，%s，当前目标：%s
私有摘要：%s
最近感知事件：%s
私有记忆：%s

只输出合法 JSON：{"summary":"","goal":"","next_decision_minutes":1440,"actions":[{"type":"","target":{"type":"location|npc|faction|player","id":0},"parameters":{},"duration_minutes":0}]}

计划必须符合角色认知、能力、资源和当前位置，不得引用未提供的秘密。summary 80 字内，goal 60 字内，actions 最多 3 项。`,
		req.State.PublicChronicle,
		req.State.WorldTime.Format(time.RFC3339),
		req.Location.Name,
		req.Location.Description,
		req.Npc.Name,
		req.Npc.Role,
		req.Npc.Species,
		req.Npc.Goal,
		req.Npc.ContextSummary,
		c.eventContext(req.RecentEvents),
		c.memoryContext(req.Memories),
	)

	systemPrompt := req.Npc.SystemPrompt + "\n你需要为自己制定下一步计划，只输出合法 JSON。所有字符串必须使用简体中文。"
	output := new(model.NpcPlan)
	if err := c.completeJSON(ctx, req.Config, systemPrompt, prompt, c.npcPlanSchema(), output, 384); err != nil {
		return nil, err
	}
	return output, nil
}

func (c *AgentClient) Tick(ctx context.Context, req *repo.TickReq) (*model.ActionResolution, error) {
	prompt := fmt.Sprintf(`世界：%s
内部世界摘要：%s
当前阶段：%s
世界时间：%s
近期已提交事件：%s
NPC 当前状态：%s
地点当前状态：%s
阵营当前状态：%s

推动世界自然演进一步。只输出合法 JSON：
{
  "status": "resolved",
  "summary": "",
  "world_summary": "",
  "current_arc": "",
  "actions": [
    {
      "type": "move_npc|change_npc_state|change_location|change_faction|change_relationship|share_claim",
      "target": {"type": "npc|location|faction", "id": 0},
      "parameters": {},
      "duration_minutes": 0
    }
  ],
  "claims": [],
  "suggested_actions": []
}

所有变化必须有近期事件、世界规则或角色目标支撑，避免重复已有内容。世界可以在玩家不在线时继续变化。
summary 80 字内，world_summary 120 字内，current_arc 40 字内，actions 最多 3 项，claims 最多 3 项。`,
		req.World.Name,
		req.State.Summary,
		req.State.CurrentArc,
		req.State.WorldTime.Format(time.RFC3339),
		c.eventContext(req.RecentEvents),
		c.npcContext(req.Npcs),
		c.locationContext(req.Locations),
		c.factionContext(req.Factions),
	)

	output := new(model.ActionResolution)
	if err := c.completeJSON(ctx, req.Config, "你是世界裁决者，负责提出结构化演化动作，只输出合法 JSON。所有字符串必须使用简体中文。", prompt, c.actionResolutionSchema(), output, 384); err != nil {
		return nil, err
	}
	return output, nil
}

func (c *AgentClient) eventContext(events []*model.Event) string {
	parts := lo.FilterMap(events, func(event *model.Event, _ int) (string, bool) {
		if event == nil {
			return "", false
		}
		return event.Summary + ": " + event.Content, true
	})
	return strings.Join(parts, " | ")
}

func (c *AgentClient) memoryContext(memories []*model.NpcMemory) string {
	parts := lo.FilterMap(memories, func(memory *model.NpcMemory, _ int) (string, bool) {
		if memory == nil {
			return "", false
		}
		return memory.Content, true
	})
	return strings.Join(parts, " | ")
}

func (c *AgentClient) characterName(member *model.WorldMember, player *model.Player) string {
	if member != nil && strings.TrimSpace(member.CharacterName) != "" {
		return member.CharacterName
	}
	if player != nil {
		return player.DisplayName
	}
	return "未知玩家"
}

func (c *AgentClient) npcContext(rows []*model.Npc) string {
	parts := lo.Map(rows, func(row *model.Npc, _ int) string {
		return fmt.Sprintf("%d:%s[%s]@%d 目标=%s", row.ID, row.Name, row.LifeStatus, row.CurrentLocationID, row.Goal)
	})
	return strings.Join(parts, " | ")
}

func (c *AgentClient) locationContext(rows []*model.Location) string {
	parts := lo.Map(rows, func(row *model.Location, _ int) string {
		return fmt.Sprintf("%d:%s[%s] %s", row.ID, row.Name, row.Status, row.Description)
	})
	return strings.Join(parts, " | ")
}

func (c *AgentClient) factionContext(rows []*model.Faction) string {
	parts := lo.Map(rows, func(row *model.Faction, _ int) string {
		return fmt.Sprintf("%d:%s[%s] 目标=%s", row.ID, row.Name, row.Status, row.PublicGoal)
	})
	return strings.Join(parts, " | ")
}

func (c *AgentClient) completeJSON(ctx context.Context, config *model.AgentConfig, systemPrompt string, userPrompt string, schema any, output any, maxTokens int) error {
	if config == nil {
		return fmt.Errorf("agent config is nil")
	}

	timeout := config.Timeout
	if timeout <= 0 {
		timeout = 60 * time.Second
	}

	callCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	content, err := c.complete(callCtx, config, systemPrompt, userPrompt, schema, maxTokens)
	if err != nil {
		return err
	}
	if err = json.Unmarshal([]byte(content), output); err != nil {
		return fmt.Errorf("decode agent json: %w; content=%s", err, c.truncateBody(content, 512))
	}
	return nil
}

func (c *AgentClient) complete(ctx context.Context, config *model.AgentConfig, systemPrompt string, userPrompt string, schema any, maxTokens int) (string, error) {
	switch config.Provider {
	case enum.AgentProviderOllama:
		return c.callOllama(ctx, config, systemPrompt, userPrompt, schema, maxTokens)
	case enum.AgentProviderOpenAICompatible:
		return c.callOpenAICompatible(ctx, config, systemPrompt, userPrompt, schema, maxTokens)
	default:
		return "", fmt.Errorf("unsupported agent provider: %s", config.Provider)
	}
}

func (c *AgentClient) callOllama(ctx context.Context, config *model.AgentConfig, systemPrompt string, userPrompt string, schema any, maxTokens int) (string, error) {
	if maxTokens <= 0 {
		maxTokens = 512
	}
	body := map[string]any{
		"model":  config.Model,
		"stream": false,
		"format": c.ollamaFormat(schema),
		"think":  false,
		"options": map[string]any{
			"temperature": 0.2,
			"num_predict": maxTokens,
		},
		"messages": []map[string]string{
			{"role": "system", "content": systemPrompt},
			{"role": "user", "content": userPrompt},
		},
	}

	var response struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
		DoneReason string `json:"done_reason"`
	}
	url := strings.TrimRight(config.BaseURL, "/") + "/api/chat"
	if err := c.post(ctx, url, config, body, &response); err != nil {
		return "", err
	}
	if response.DoneReason == "length" {
		return "", fmt.Errorf("agent response was truncated by token limit")
	}
	return response.Message.Content, nil
}

func (c *AgentClient) callOpenAICompatible(ctx context.Context, config *model.AgentConfig, systemPrompt string, userPrompt string, _ any, maxTokens int) (string, error) {
	if maxTokens <= 0 {
		maxTokens = 512
	}
	body := map[string]any{
		"model":      config.Model,
		"max_tokens": maxTokens,
		"response_format": map[string]string{
			"type": "json_object",
		},
		"messages": []map[string]string{
			{"role": "system", "content": systemPrompt},
			{"role": "user", "content": userPrompt},
		},
	}

	var response struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}
	url := strings.TrimRight(config.BaseURL, "/") + "/v1/chat/completions"
	if err := c.post(ctx, url, config, body, &response); err != nil {
		return "", err
	}
	if len(response.Choices) == 0 {
		return "", fmt.Errorf("empty model choices")
	}
	return response.Choices[0].Message.Content, nil
}

func (c *AgentClient) post(ctx context.Context, url string, config *model.AgentConfig, body any, output any) error {
	data, err := json.Marshal(body)
	if err != nil {
		return err
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")

	if config.SecretEnv != "" {
		secret, ok := os.LookupEnv(config.SecretEnv)
		if !ok || secret == "" {
			return fmt.Errorf("agent secret environment variable is missing")
		}
		request.Header.Set("Authorization", "Bearer "+secret)
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusMultipleChoices {
		data, _ := io.ReadAll(io.LimitReader(response.Body, 4096))
		return fmt.Errorf("agent http status %d: %s", response.StatusCode, strings.TrimSpace(string(data)))
	}
	if err = json.NewDecoder(response.Body).Decode(output); err != nil {
		return fmt.Errorf("decode agent response: %w", err)
	}
	return nil
}

func (c *AgentClient) ollamaFormat(schema any) any {
	if schema == nil {
		return "json"
	}
	return schema
}

func (c *AgentClient) worldDraftSchema(locationCount uint32, npcCount uint32) map[string]any {
	return c.objectSchema(
		[]string{"name", "summary", "current_arc", "current_era", "rules", "locations", "factions", "npcs"},
		map[string]any{
			"name":        c.stringSchema(),
			"summary":     c.stringSchema(),
			"current_arc": c.stringSchema(),
			"current_era": c.stringSchema(),
			"rules":       map[string]any{"type": "object"},
			"locations": c.arraySchema(int(locationCount), int(locationCount), c.objectSchema(
				[]string{"code", "name", "description", "environment_tags"},
				map[string]any{
					"code":             c.stringSchema(),
					"name":             c.stringSchema(),
					"description":      c.stringSchema(),
					"environment_tags": c.stringArraySchema(0, 8),
				},
			)),
			"factions": c.arraySchema(2, 3, c.objectSchema(
				[]string{"code", "name", "description", "public_goal"},
				map[string]any{
					"code":        c.stringSchema(),
					"name":        c.stringSchema(),
					"description": c.stringSchema(),
					"public_goal": c.stringSchema(),
				},
			)),
			"npcs": c.arraySchema(int(npcCount), int(npcCount), c.objectSchema(
				[]string{"code", "name", "role", "species", "personality", "goal", "background", "location_code", "faction_code", "system_prompt", "attributes"},
				map[string]any{
					"code":          c.stringSchema(),
					"name":          c.stringSchema(),
					"role":          c.stringSchema(),
					"species":       c.stringSchema(),
					"personality":   c.stringSchema(),
					"goal":          c.stringSchema(),
					"background":    c.stringSchema(),
					"location_code": c.stringSchema(),
					"faction_code":  c.stringSchema(),
					"system_prompt": c.stringSchema(),
					"attributes":    map[string]any{"type": "object"},
				},
			)),
		},
	)
}

func (c *AgentClient) characterDraftSchema() map[string]any {
	return c.objectSchema(
		[]string{"name", "background", "goal", "traits"},
		map[string]any{
			"name":       c.stringSchema(32),
			"background": c.stringSchema(180),
			"goal":       c.stringSchema(100),
			"traits":     c.stringArraySchema(0, 6),
		},
	)
}

func (c *AgentClient) npcReplySchema() map[string]any {
	return c.objectSchema(
		[]string{"reply", "context_summary", "suggested_actions", "claims"},
		map[string]any{
			"reply":             c.stringSchema(180),
			"context_summary":   c.stringSchema(120),
			"suggested_actions": c.suggestedActionsSchema(),
			"claims":            c.claimsSchema(),
		},
	)
}

func (c *AgentClient) actionResolutionSchema() map[string]any {
	return c.objectSchema(
		[]string{"status", "summary", "clarification", "world_summary", "current_arc", "actions", "claims", "suggested_actions"},
		map[string]any{
			"status":            c.enumStringSchema([]string{"resolved", "rejected", "clarification"}),
			"summary":           c.stringSchema(120),
			"clarification":     c.stringSchema(100),
			"world_summary":     c.stringSchema(180),
			"current_arc":       c.stringSchema(80),
			"actions":           c.actionStepsSchema(),
			"claims":            c.claimsSchema(),
			"suggested_actions": c.suggestedActionsSchema(),
		},
	)
}

func (c *AgentClient) npcPlanSchema() map[string]any {
	return c.objectSchema(
		[]string{"summary", "goal", "next_decision_minutes", "actions"},
		map[string]any{
			"summary":               c.stringSchema(120),
			"goal":                  c.stringSchema(100),
			"next_decision_minutes": map[string]any{"type": "integer"},
			"actions":               c.actionStepsSchema(),
		},
	)
}

func (c *AgentClient) suggestedActionsSchema() map[string]any {
	return c.arraySchema(0, 4, c.objectSchema(
		[]string{"label", "content", "targets"},
		map[string]any{
			"label":   c.stringSchema(40),
			"content": c.stringSchema(100),
			"targets": c.entityRefsSchema(),
		},
	))
}

func (c *AgentClient) claimsSchema() map[string]any {
	return c.arraySchema(0, 3, c.objectSchema(
		[]string{"subject_type", "subject_id", "predicate", "object", "truth"},
		map[string]any{
			"subject_type": c.enumStringSchema([]string{"world", "player", "npc", "location", "faction"}),
			"subject_id":   map[string]any{"type": "integer"},
			"predicate":    c.stringSchema(80),
			"object":       map[string]any{"type": "object"},
			"truth":        c.enumStringSchema([]string{"true", "false", "unknown"}),
		},
	))
}

func (c *AgentClient) actionStepsSchema() map[string]any {
	return c.arraySchema(0, 3, c.objectSchema(
		[]string{"type", "target", "parameters", "duration_minutes"},
		map[string]any{
			"type":             c.enumStringSchema([]string{"move", "move_player", "move_npc", "change_npc_state", "change_location", "change_faction", "change_relationship", "share_claim"}),
			"target":           c.entityRefSchema(),
			"parameters":       map[string]any{"type": "object"},
			"duration_minutes": map[string]any{"type": "integer"},
		},
	))
}

func (c *AgentClient) entityRefsSchema() map[string]any {
	return c.arraySchema(0, 4, c.entityRefSchema())
}

func (c *AgentClient) entityRefSchema() map[string]any {
	return c.objectSchema(
		[]string{"type", "id"},
		map[string]any{
			"type": c.enumStringSchema([]string{"world", "player", "npc", "location", "faction"}),
			"id":   map[string]any{"type": "integer"},
		},
	)
}

func (c *AgentClient) objectSchema(required []string, properties map[string]any) map[string]any {
	return map[string]any{
		"type":                 "object",
		"additionalProperties": false,
		"required":             required,
		"properties":           properties,
	}
}

func (c *AgentClient) arraySchema(minItems int, maxItems int, items any) map[string]any {
	return map[string]any{
		"type":     "array",
		"minItems": minItems,
		"maxItems": maxItems,
		"items":    items,
	}
}

func (c *AgentClient) stringArraySchema(minItems int, maxItems int) map[string]any {
	return c.arraySchema(minItems, maxItems, c.stringSchema())
}

func (c *AgentClient) stringSchema(maxLength ...int) map[string]any {
	schema := map[string]any{"type": "string"}
	if len(maxLength) > 0 && maxLength[0] > 0 {
		schema["maxLength"] = maxLength[0]
	}
	return schema
}

func (c *AgentClient) enumStringSchema(values []string) map[string]any {
	return map[string]any{"type": "string", "enum": values}
}

func (c *AgentClient) truncateBody(content string, limit int) string {
	if len(content) <= limit {
		return content
	}
	return content[:limit]
}
