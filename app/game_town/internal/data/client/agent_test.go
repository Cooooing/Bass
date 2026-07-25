package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"game_town/internal/biz/model"
	"game_town/internal/biz/repo"
	"game_town/internal/enum"
)

func TestAgentClientGenerateWorldWithOllama(
	t *testing.T,
) {
	draftJSON, err := json.Marshal(&model.WorldDraft{
		Name:       "Town",
		Summary:    "Ready",
		CurrentArc: "opening",
		Locations: []model.WorldDraftLocation{
			{Code: "square", Name: "Square", Description: "Center"},
		},
		Npcs: []model.WorldDraftNpc{
			{
				Code: "guard", Name: "Guard", Role: "guard", Personality: "calm",
				Goal: "protect", Background: "local", LocationCode: "square",
				SystemPrompt: "You are Guard",
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		if request.URL.Path != "/api/chat" {
			t.Fatalf("unexpected path: %s", request.URL.Path)
		}
		var body map[string]any
		if err := json.NewDecoder(request.Body).Decode(&body); err != nil {
			t.Fatal(err)
		}
		format, ok := body["format"].(map[string]any)
		if !ok || format["type"] != "object" {
			t.Fatalf("expected json schema format, got %#v", body["format"])
		}
		if body["think"] != false {
			t.Fatalf("expected thinking disabled, got %v", body["think"])
		}
		options, ok := body["options"].(map[string]any)
		if !ok {
			t.Fatalf("expected options, got %T", body["options"])
		}
		if options["num_predict"] != float64(1024) {
			t.Fatalf("unexpected num_predict: %v", options["num_predict"])
		}
		if err := json.NewEncoder(w).Encode(map[string]any{
			"message": map[string]string{"content": string(draftJSON)},
		}); err != nil {
			t.Fatal(err)
		}
	}))
	defer server.Close()

	client := &AgentClient{
		httpClient: server.Client(),
	}
	draft, err := client.GenerateWorld(context.Background(), &repo.GenerateWorldReq{
		Config: &model.AgentConfig{
			Provider:       enum.AgentProviderOllama,
			BaseURL:        server.URL,
			Model:          "qwen",
			TimeoutSeconds: 5,
		},
		World: &model.World{
			Description: "test",
		},
		NpcCount:      1,
		LocationCount: 1,
	})
	if err != nil {
		t.Fatal(err)
	}
	if draft.Name != "Town" || len(draft.Npcs) != 1 || len(draft.Locations) != 1 {
		t.Fatalf("unexpected draft: %#v", draft)
	}
}

func TestAgentClientOpenAISecretRequired(
	t *testing.T,
) {
	client := &AgentClient{
		httpClient: http.DefaultClient,
	}
	_, err := client.Talk(context.Background(), &repo.TalkReq{
		Config: &model.AgentConfig{
			Provider:       enum.AgentProviderOpenAICompatible,
			BaseURL:        "http://127.0.0.1",
			Model:          "test",
			SecretEnv:      "GAME_TOWN_MISSING_SECRET",
			TimeoutSeconds: 1,
		},
		World:    &model.World{},
		State:    &model.WorldState{},
		Player:   &model.Player{},
		Location: &model.Location{},
		Npc:      &model.Npc{},
	})
	if err == nil || !strings.Contains(err.Error(), "environment variable") {
		t.Fatalf("expected missing secret error, got %v", err)
	}
}
