package main

import (
	"context"
	"io"
	"strings"
	"testing"

	v1 "common/proto/gen/game_town/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestRecordEventsKeepsPlayerSequencesAndDeduplicatesGlobalStats(t *testing.T) {
	state := &runState{
		playerSeq: make(map[int64]uint64),
		seenEvent: make(map[int64]bool),
	}

	state.recordPageEvent(1, 10, 1, v1.GameTownEventType_GAME_TOWN_EVENT_TYPE_NPC_REPLIED)
	state.recordPageEvent(2, 10, 1, v1.GameTownEventType_GAME_TOWN_EVENT_TYPE_NPC_REPLIED)
	state.recordStreamEvent(1, 11, 2, v1.GameTownEventType_GAME_TOWN_EVENT_TYPE_WORLD_EVOLVED)
	state.recordStreamEvent(1, 11, 2, v1.GameTownEventType_GAME_TOWN_EVENT_TYPE_WORLD_EVOLVED)

	stats := state.snapshotStats()
	if stats.pageEvents != 2 {
		t.Fatalf("pageEvents = %d, want 2", stats.pageEvents)
	}
	if stats.streamEvents != 2 {
		t.Fatalf("streamEvents = %d, want 2", stats.streamEvents)
	}
	if stats.npcReplies != 1 {
		t.Fatalf("npcReplies = %d, want 1", stats.npcReplies)
	}
	if stats.worldEvolved != 1 {
		t.Fatalf("worldEvolved = %d, want 1", stats.worldEvolved)
	}
	if got := state.lastSequence(1); got != 2 {
		t.Fatalf("player 1 sequence = %d, want 2", got)
	}
	if got := state.lastSequence(2); got != 1 {
		t.Fatalf("player 2 sequence = %d, want 1", got)
	}
}

func TestRetryableWatchError(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{name: "eof", err: io.EOF, want: true},
		{name: "unavailable", err: status.Error(codes.Unavailable, "server closed"), want: true},
		{name: "deadline", err: status.Error(codes.DeadlineExceeded, "timeout"), want: true},
		{name: "invalid", err: status.Error(codes.InvalidArgument, "bad request"), want: false},
		{name: "nil", err: nil, want: false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := retryableWatchError(ctx, test.err); got != test.want {
				t.Fatalf("retryableWatchError() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestRetryableWatchErrorStopsWhenContextCanceled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if retryableWatchError(ctx, io.EOF) {
		t.Fatal("retryableWatchError() = true after context cancellation")
	}
}

func TestValidateRunStrictRequiresWatchEvents(t *testing.T) {
	state := &runState{
		playerSeq: make(map[int64]uint64),
		seenEvent: make(map[int64]bool),
		stats: runStats{
			submittedActions: 10,
			characterReady:   2,
			bigEvents:        1,
		},
	}
	err := validateRun(state, config{players: 2, rounds: 10, bigEventEvery: 5, strict: true})
	if err == nil {
		t.Fatal("validateRun() error = nil, want missing Watch event error")
	}
	if !strings.Contains(err.Error(), "Watch") {
		t.Fatalf("validateRun() error = %v, want Watch error", err)
	}
}

func TestValidateRunStrictRequiresFeedbackDensity(t *testing.T) {
	state := &runState{
		playerSeq: make(map[int64]uint64),
		seenEvent: make(map[int64]bool),
		stats: runStats{
			submittedActions: 20,
			characterReady:   2,
			bigEvents:        4,
			streamEvents:     10,
			npcReplies:       1,
			resolved:         1,
		},
		playerStats: map[int64]*playerRunStats{
			1: {submitted: 10, pageEvents: 5, streamEvents: 5, visibleEvents: 5, npcReplies: 1, characterReady: true},
			2: {submitted: 10, pageEvents: 5, streamEvents: 5, visibleEvents: 5, actionResults: 1, characterReady: true},
		},
	}

	err := validateRun(state, config{players: 2, rounds: 20, bigEventEvery: 5, strict: true})
	if err == nil {
		t.Fatal("validateRun() error = nil, want feedback density error")
	}
	if !strings.Contains(err.Error(), "feedback") {
		t.Fatalf("validateRun() error = %v, want feedback error", err)
	}
}

func TestValidateRunStrictRequiresWorldEvolutionForLongRun(t *testing.T) {
	state := &runState{
		playerSeq: make(map[int64]uint64),
		seenEvent: make(map[int64]bool),
		stats: runStats{
			submittedActions: 1000,
			bigEvents:        10,
			streamEvents:     100,
			npcReplies:       400,
			resolved:         100,
			characterReady:   4,
		},
		playerStats: map[int64]*playerRunStats{
			1: {submitted: 250, pageEvents: 100, streamEvents: 100, visibleEvents: 250, npcReplies: 90, actionResults: 10, characterReady: true},
			2: {submitted: 250, pageEvents: 100, streamEvents: 100, visibleEvents: 250, npcReplies: 90, actionResults: 10, characterReady: true},
			3: {submitted: 250, pageEvents: 100, streamEvents: 100, visibleEvents: 250, npcReplies: 90, actionResults: 10, characterReady: true},
			4: {submitted: 250, pageEvents: 100, streamEvents: 100, visibleEvents: 250, npcReplies: 90, actionResults: 10, characterReady: true},
		},
	}

	err := validateRun(state, config{players: 4, rounds: 1000, bigEventEvery: 100, strict: true})
	if err == nil {
		t.Fatal("validateRun() error = nil, want world evolution error")
	}
	if !strings.Contains(err.Error(), "world evolution") {
		t.Fatalf("validateRun() error = %v, want world evolution error", err)
	}
}
