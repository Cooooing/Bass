package main

import (
	v1enum "common/proto/gen/game_town/v1/enum"
	"context"
	"strings"
	"testing"

	v1 "common/proto/gen/game_town/v1"

	"github.com/samber/lo"
)

func TestResolveClientMode(
	t *testing.T,
) {
	tests := []struct {
		name     string
		mode     string
		terminal bool
		want     string
		wantErr  bool
	}{
		{name: "auto terminal", mode: clientModeAuto, terminal: true, want: clientModeTUI},
		{name: "auto console", mode: clientModeAuto, terminal: false, want: clientModeConsole},
		{name: "forced tui", mode: clientModeTUI, want: clientModeTUI},
		{name: "forced console", mode: clientModeConsole, want: clientModeConsole},
		{name: "invalid", mode: "invalid", wantErr: true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := resolveClientMode(test.mode, test.terminal)
			if (err != nil) != test.wantErr {
				t.Fatalf("resolveClientMode() error = %v", err)
			}
			if got != test.want {
				t.Fatalf("resolveClientMode() = %q, want %q", got, test.want)
			}
		})
	}
}

func TestExecuteHelp(
	t *testing.T,
) {
	result := executeCommand(context.Background(), nil, 0, 0, 0, nil, "/help")
	if result.err != nil {
		t.Fatalf("executeCommand() error = %v", result.err)
	}
	if !lo.Contains(result.lines, "/register <name>") {
		t.Fatalf("help lines = %v", result.lines)
	}
}

func TestRunConsoleAcceptsLineInput(
	t *testing.T,
) {
	input := strings.NewReader("/help\n/quit\n")
	var output strings.Builder
	if err := runConsole(context.Background(), nil, "test", input, &output); err != nil {
		t.Fatalf("runConsole() error = %v", err)
	}
	if !strings.Contains(output.String(), "/register <name>") {
		t.Fatalf("console output = %q", output.String())
	}
}

func TestExecuteIncompleteCommandShowsUsage(
	t *testing.T,
) {
	result := executeCommand(context.Background(), nil, 0, 0, 0, nil, "/world create")
	if result.err != nil {
		t.Fatalf("executeCommand() error = %v", result.err)
	}
	if len(result.lines) == 0 || !strings.Contains(result.lines[0], "用法") {
		t.Fatalf("executeCommand() lines = %v", result.lines)
	}
}

func TestExecuteUnknownCommand(
	t *testing.T,
) {
	result := executeCommand(context.Background(), nil, 0, 0, 0, nil, "/missing")
	if result.err != nil {
		t.Fatalf("executeCommand() error = %v", result.err)
	}
	if len(result.lines) != 1 || !strings.Contains(result.lines[0], "未知命令") {
		t.Fatalf("executeCommand() lines = %v", result.lines)
	}
}

func TestExecuteBackClearsDialog(
	t *testing.T,
) {
	result := executeCommand(context.Background(), nil, 1, 1, 7, nil, "/back")
	if result.err != nil {
		t.Fatalf("executeCommand() error = %v", result.err)
	}
	if !result.clearDialog {
		t.Fatalf("clearDialog = false")
	}
}

func TestJoinWorldRejectsPlaceholderCode(
	t *testing.T,
) {
	result := joinWorld(context.Background(), nil, 1, []string{"<world_code>", "角色倾向"})
	if result.err == nil {
		t.Fatalf("expected placeholder error")
	}
	if !strings.Contains(result.err.Error(), "替换") {
		t.Fatalf("unexpected error: %v", result.err)
	}
}

func TestSubmitSuggestedChoiceRejectsOutOfRange(
	t *testing.T,
) {
	_, ok := submitSuggestedChoice(context.Background(), nil, 1, 1, 0, nil, "1")
	if ok {
		t.Fatalf("unexpected suggested choice match")
	}
}

func TestExecuteNumericInputWithoutSuggestionsShowsHint(
	t *testing.T,
) {
	result := executeCommand(context.Background(), nil, 1, 1, 0, nil, "3")
	if result.err != nil {
		t.Fatalf("executeCommand() error = %v", result.err)
	}
	if len(result.lines) == 0 || !strings.Contains(result.lines[0], "当前没有可选回答") {
		t.Fatalf("executeCommand() lines = %v", result.lines)
	}
}

func TestEventSuggestedChoicesFiltersInvalidTargetsAndFallsBackToNpc(
	t *testing.T,
) {
	npcID := int64(9)
	choices := eventSuggestedChoices(&v1.WatchGameTownEvents_Resp{
		NpcId: &npcID,
		SuggestedActions: []*v1.WatchGameTownEvents_Resp_SuggestedAction{
			{
				Label:   "继续询问",
				Content: "继续询问星图来源",
				Targets: []*v1.WatchGameTownEvents_Resp_SuggestedAction_EntityRef{
					{Type: v1enum.GameTownEntityType_GAME_TOWN_ENTITY_TYPE_UNSPECIFIED, Id: 0},
				},
			},
		},
	})
	if len(choices) != 1 {
		t.Fatalf("choices len = %d, want 1", len(choices))
	}
	if len(choices[0].targets) != 1 {
		t.Fatalf("targets len = %d, want 1", len(choices[0].targets))
	}
	if choices[0].targets[0].GetType() != v1enum.GameTownEntityType_GAME_TOWN_ENTITY_TYPE_NPC {
		t.Fatalf("target type = %v, want NPC", choices[0].targets[0].GetType())
	}
	if choices[0].targets[0].GetId() != 9 {
		t.Fatalf("target id = %d, want 9", choices[0].targets[0].GetId())
	}
}
