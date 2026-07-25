package repo

import (
	commonclient "common/pkg/client"
	"context"
	"fmt"
	bizrepo "scheduler/internal/biz/repo"
	schedulerenum "scheduler/internal/enum"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
)

func TestTaskLockRepoScheduleExclusiveAndRunningHeartbeat(t *testing.T) {
	redisServer := miniredis.RunT(t)
	redisClient := redis.NewClient(&redis.Options{
		Addr: redisServer.Addr(),
	})
	t.Cleanup(func() {
		_ = redisClient.Close()
	})
	lockRepo := NewTaskLockRepo(&commonclient.RedisClient{
		Client: redisClient,
	}).(*TaskLockRepo)
	ctx := context.Background()
	taskID := int64(1)
	executionID := int64(100)
	scheduledAt := time.Date(2026, 7, 4, 10, 0, 0, 0, time.UTC)
	req := &bizrepo.TaskScheduleAcquireReq{
		TaskID:            taskID,
		ScheduledAt:       scheduledAt,
		AllowOverlap:      false,
		SchedulePeriodTTL: time.Minute,
		RunningLockTTL:    time.Minute,
	}

	acquired, err := lockRepo.TryAcquireSchedule(ctx, req)
	if err != nil {
		t.Fatalf("TryAcquireSchedule returned error: %v", err)
	}
	if acquired.Decision != schedulerenum.TaskScheduleDecisionRun || acquired.RunningToken == "" {
		t.Fatalf("expected run decision with running token, got %#v", acquired)
	}

	duplicated, err := lockRepo.TryAcquireSchedule(ctx, req)
	if err != nil {
		t.Fatalf("TryAcquireSchedule duplicated returned error: %v", err)
	}
	if duplicated.Decision != schedulerenum.TaskScheduleDecisionSkip || duplicated.RunningToken != "" {
		t.Fatalf("expected duplicated schedule skip, got %#v", duplicated)
	}

	overlapped, err := lockRepo.TryAcquireSchedule(ctx, &bizrepo.TaskScheduleAcquireReq{
		TaskID:            taskID,
		ScheduledAt:       scheduledAt.Add(time.Second),
		AllowOverlap:      false,
		SchedulePeriodTTL: time.Minute,
		RunningLockTTL:    time.Minute,
	})
	if err != nil {
		t.Fatalf("TryAcquireSchedule overlapped returned error: %v", err)
	}
	if overlapped.Decision != schedulerenum.TaskScheduleDecisionOverlap || overlapped.RunningToken != "" {
		t.Fatalf("expected overlap decision, got %#v", overlapped)
	}

	registerResp, err := lockRepo.RegisterRunning(ctx, &bizrepo.TaskRunningLockReq{
		TaskID:            taskID,
		ExecutionRecordID: executionID,
		RunningToken:      acquired.RunningToken,
		Exclusive:         true,
		TTL:               time.Minute,
	})
	if err != nil {
		t.Fatalf("RegisterRunning returned error: %v", err)
	}
	if !registerResp {
		t.Fatal("expected running registration to succeed")
	}
	value, err := redisClient.HGet(ctx, fmt.Sprintf(lockRepo.taskLockKeyFormat, taskID), "running:100").Result()
	if err != nil || value != acquired.RunningToken {
		t.Fatalf("expected execution running field to keep token, value=%q err=%v", value, err)
	}
	value, err = redisClient.HGet(ctx, fmt.Sprintf(lockRepo.taskLockKeyFormat, taskID), lockRepo.exclusiveField).Result()
	if err != nil || value != acquired.RunningToken {
		t.Fatalf("expected exclusive field to keep token, value=%q err=%v", value, err)
	}

	refreshResp, err := lockRepo.RefreshRunning(ctx, &bizrepo.TaskRunningLockReq{
		TaskID:            taskID,
		ExecutionRecordID: executionID,
		RunningToken:      "wrong-token",
		Exclusive:         true,
		TTL:               time.Minute,
	})
	if err != nil {
		t.Fatalf("RefreshRunning wrong token returned error: %v", err)
	}
	if refreshResp {
		t.Fatal("expected wrong running token refresh to fail")
	}
	if err = lockRepo.ReleaseRunning(ctx, &bizrepo.TaskRunningLockReq{
		TaskID:            taskID,
		ExecutionRecordID: executionID,
		RunningToken:      "wrong-token",
		Exclusive:         true,
	}); err != nil {
		t.Fatalf("ReleaseRunning wrong token returned error: %v", err)
	}
	value, err = redisClient.HGet(ctx, fmt.Sprintf(lockRepo.taskLockKeyFormat, taskID), lockRepo.exclusiveField).Result()
	if err != nil || value != acquired.RunningToken {
		t.Fatalf("expected exclusive field to keep original token, value=%q err=%v", value, err)
	}

	refreshResp, err = lockRepo.RefreshRunning(ctx, &bizrepo.TaskRunningLockReq{
		TaskID:            taskID,
		ExecutionRecordID: executionID,
		RunningToken:      acquired.RunningToken,
		Exclusive:         true,
		TTL:               time.Minute,
	})
	if err != nil {
		t.Fatalf("RefreshRunning correct token returned error: %v", err)
	}
	if !refreshResp {
		t.Fatal("expected correct running token refresh to succeed")
	}
	runningResp, err := lockRepo.MapRunning(ctx, &bizrepo.TaskRunningMapReq{
		TaskID:             taskID,
		ExecutionRecordIDs: []int64{executionID, 101},
	})
	if err != nil {
		t.Fatalf("MapRunning returned error: %v", err)
	}
	if !runningResp[executionID] || runningResp[101] {
		t.Fatalf("unexpected running map: %#v", runningResp)
	}
	if err = lockRepo.ReleaseRunning(ctx, &bizrepo.TaskRunningLockReq{
		TaskID:            taskID,
		ExecutionRecordID: executionID,
		RunningToken:      acquired.RunningToken,
		Exclusive:         true,
	}); err != nil {
		t.Fatalf("ReleaseRunning correct token returned error: %v", err)
	}
	if exists, err := redisClient.HExists(ctx, fmt.Sprintf(lockRepo.taskLockKeyFormat, taskID), "running:100").Result(); err != nil || exists {
		t.Fatalf("expected execution running field deleted, exists=%v err=%v", exists, err)
	}
	if exists, err := redisClient.HExists(ctx, fmt.Sprintf(lockRepo.taskLockKeyFormat, taskID), lockRepo.exclusiveField).Result(); err != nil || exists {
		t.Fatalf("expected exclusive field deleted, exists=%v err=%v", exists, err)
	}
}
