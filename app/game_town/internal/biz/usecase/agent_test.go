package usecase

import (
	"reflect"
	"strings"
	"testing"
	"time"

	"game_town/internal/biz/model"
	"game_town/internal/config"
	"game_town/internal/enum"

	"github.com/samber/lo"
	"google.golang.org/protobuf/types/known/durationpb"
)

func TestFairAgentJobsPrioritizesAndRoundsWorlds(
	t *testing.T,
) {
	jobs := []*model.AgentJob{
		{ID: 1, WorldID: 10, Priority: enum.AgentJobPriorityHigh},
		{ID: 2, WorldID: 10, Priority: enum.AgentJobPriorityHigh},
		{ID: 3, WorldID: 20, Priority: enum.AgentJobPriorityHigh},
		{ID: 4, WorldID: 10, Priority: enum.AgentJobPriorityLow},
		{ID: 5, WorldID: 20, Priority: enum.AgentJobPriorityLow},
	}
	got := jobIDs(fairAgentJobs(jobs, new(0)))
	want := []int64{1, 3, 2, 4, 5}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected fair order: got %v want %v", got, want)
	}
}

func TestFairAgentJobsRotatesStartingWorld(
	t *testing.T,
) {
	jobs := []*model.AgentJob{
		{ID: 1, WorldID: 10, Priority: enum.AgentJobPriorityHigh},
		{ID: 2, WorldID: 20, Priority: enum.AgentJobPriorityHigh},
	}
	cursor := new(0)
	_ = fairAgentJobs(jobs, cursor)
	got := jobIDs(fairAgentJobs(jobs, cursor))
	want := []int64{2, 1}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected rotated order: got %v want %v", got, want)
	}
}

func TestPromoteOverdueTickJobsPreventsLowPriorityStarvation(
	t *testing.T,
) {
	now := time.Now()
	runner := &WorldAgentRunner{
		conf: &config.Bootstrap{
			GameTown: &config.GameTown{
				Agent: &config.Agent{
					TickInterval: durationpb.New(time.Minute),
				},
			},
		},
	}
	jobs := []*model.AgentJob{
		{ID: 1, WorldID: 10, Type: enum.AgentJobTypePlayerActionInterpret, Priority: enum.AgentJobPriorityHigh, AvailableAt: now},
		{ID: 2, WorldID: 10, Type: enum.AgentJobTypeWorldTick, Priority: enum.AgentJobPriorityLow, AvailableAt: now.Add(-time.Minute)},
		{ID: 3, WorldID: 10, Type: enum.AgentJobTypeNpcPlan, Priority: enum.AgentJobPriorityLow, AvailableAt: now.Add(-time.Minute)},
	}

	got := jobIDs(runner.promoteOverdueTickJobs(jobs, now))
	want := []int64{2, 1, 3}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected promoted order: got %v want %v", got, want)
	}
}

func TestSchedulerOrderPromotesOverdueTickBeforeHighPriorityBacklog(
	t *testing.T,
) {
	now := time.Now()
	runner := &WorldAgentRunner{
		conf: &config.Bootstrap{
			GameTown: &config.GameTown{
				Agent: &config.Agent{
					TickInterval: durationpb.New(time.Minute),
				},
			},
		},
	}
	jobs := []*model.AgentJob{
		{ID: 1, WorldID: 10, Type: enum.AgentJobTypeWorldTick, Priority: enum.AgentJobPriorityLow, AvailableAt: now.Add(-time.Minute)},
		{ID: 2, WorldID: 20, Type: enum.AgentJobTypeWorldGenerate, Priority: enum.AgentJobPriorityHigh, AvailableAt: now},
		{ID: 3, WorldID: 10, Type: enum.AgentJobTypeNpcPlan, Priority: enum.AgentJobPriorityLow, AvailableAt: now.Add(-time.Minute)},
	}

	got := jobIDs(runner.orderDispatchJobs(jobs, now))
	want := []int64{1, 2, 3}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected scheduler order: got %v want %v", got, want)
	}
}

func TestTickScanInterval(
	t *testing.T,
) {
	tests := []struct {
		name         string
		tickInterval time.Duration
		want         time.Duration
	}{
		{name: "minimum", tickInterval: 2 * time.Second, want: time.Second},
		{name: "quarter", tickInterval: 20 * time.Second, want: 5 * time.Second},
		{name: "maximum", tickInterval: time.Minute, want: 15 * time.Second},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			runner := &WorldAgentRunner{
				conf: &config.Bootstrap{
					GameTown: &config.GameTown{
						Agent: &config.Agent{
							TickInterval: durationpb.New(test.tickInterval),
						},
					},
				},
			}
			if got := runner.tickScanInterval(); got != test.want {
				t.Fatalf("tickScanInterval() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestNormalizeModelTextUsesRuneLimit(
	t *testing.T,
) {
	value := "  " + strings.Repeat("雾", maxCurrentArcRunes+20) + "  "
	got := normalizeModelText(value, maxCurrentArcRunes)
	if len([]rune(got)) != maxCurrentArcRunes {
		t.Fatalf("rune length = %d", len([]rune(got)))
	}
	if strings.HasPrefix(got, " ") || strings.HasSuffix(got, " ") {
		t.Fatalf("value was not trimmed: %q", got)
	}
}

func jobIDs(
	jobs []*model.AgentJob,
) []int64 {
	return lo.Map(jobs, func(job *model.AgentJob, _ int) int64 {
		return job.ID
	})
}
