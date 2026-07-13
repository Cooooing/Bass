package repo

import (
	commonclient "common/pkg/client"
	"context"
	"fmt"
	"math"
	bizrepo "scheduler/internal/biz/repo"
	schedulerenum "scheduler/internal/enum"
	"strconv"
	"time"

	"github.com/google/uuid"
)

var _ bizrepo.TaskLockRepo = (*TaskLockRepo)(nil)

// TaskLockRepo 使用单任务 Hash 保存周期准入、互斥槽和执行心跳。
type TaskLockRepo struct {
	redisClient           *commonclient.RedisClient
	taskLockKeyFormat     string
	exclusiveField        string
	tryAcquireScript      string
	registerRunningScript string
	refreshRunningScript  string
	releaseRunningScript  string
}

func NewTaskLockRepo(redisClient *commonclient.RedisClient) bizrepo.TaskLockRepo {
	return &TaskLockRepo{
		redisClient:       redisClient,
		taskLockKeyFormat: "Scheduler:TaskLock:{%d}",
		exclusiveField:    "exclusive",
		tryAcquireScript: `
local ok = redis.call("HSETNX", KEYS[1], ARGV[1], ARGV[2])
if ok == 0 then
	return "skip"
end
redis.call("HEXPIRE", KEYS[1], ARGV[3], "FIELDS", 1, ARGV[1])
if ARGV[4] == "1" then
	return "run"
end
ok = redis.call("HSETNX", KEYS[1], ARGV[5], ARGV[6])
if ok == 1 then
	redis.call("HEXPIRE", KEYS[1], ARGV[7], "FIELDS", 1, ARGV[5])
	return "run"
end
return "overlap"
`,
		registerRunningScript: `
if ARGV[5] == "1" and redis.call("HGET", KEYS[1], ARGV[1]) ~= ARGV[3] then
	return 0
end
redis.call("HSET", KEYS[1], ARGV[2], ARGV[3])
redis.call("HEXPIRE", KEYS[1], ARGV[4], "FIELDS", 1, ARGV[2])
if ARGV[5] == "1" then
	redis.call("HEXPIRE", KEYS[1], ARGV[4], "FIELDS", 1, ARGV[1])
end
return 1
`,
		refreshRunningScript: `
if redis.call("HGET", KEYS[1], ARGV[2]) ~= ARGV[3] then
	return 0
end
if ARGV[5] == "1" and redis.call("HGET", KEYS[1], ARGV[1]) ~= ARGV[3] then
	return 0
end
redis.call("HEXPIRE", KEYS[1], ARGV[4], "FIELDS", 1, ARGV[2])
if ARGV[5] == "1" then
	redis.call("HEXPIRE", KEYS[1], ARGV[4], "FIELDS", 1, ARGV[1])
end
return 1
`,
		releaseRunningScript: `
if redis.call("HGET", KEYS[1], ARGV[2]) == ARGV[3] then
	redis.call("HDEL", KEYS[1], ARGV[2])
end
if ARGV[4] == "1" and redis.call("HGET", KEYS[1], ARGV[1]) == ARGV[3] then
	redis.call("HDEL", KEYS[1], ARGV[1])
end
return 0
`,
	}
}

func (r *TaskLockRepo) TryAcquireSchedule(ctx context.Context, req *bizrepo.TaskScheduleAcquireReq) (*bizrepo.TaskScheduleAcquireResult, error) {
	if req == nil {
		return nil, fmt.Errorf("scheduler task schedule acquire request is nil")
	}
	scheduleField := "schedule:" + req.ScheduledAt.Format(time.RFC3339Nano)
	runningToken := uuid.NewString()
	allowOverlap := "0"
	runningTTLSeconds := int64(1)
	if req.AllowOverlap {
		allowOverlap = "1"
	} else {
		runningTTLSeconds = int64(math.Ceil(req.RunningLockTTL.Seconds()))
	}
	result, err := r.redisClient.Client.Eval(
		ctx,
		r.tryAcquireScript,
		[]string{fmt.Sprintf(r.taskLockKeyFormat, req.TaskID)},
		scheduleField,
		uuid.NewString(),
		int64(math.Ceil(req.SchedulePeriodTTL.Seconds())),
		allowOverlap,
		r.exclusiveField,
		runningToken,
		runningTTLSeconds,
	).Text()
	if err != nil {
		return nil, err
	}
	decision := schedulerenum.TaskScheduleDecision(result)
	if decision != schedulerenum.TaskScheduleDecisionRun {
		runningToken = ""
	}
	return &bizrepo.TaskScheduleAcquireResult{
		Decision:     decision,
		RunningToken: runningToken,
	}, nil
}

func (r *TaskLockRepo) RegisterRunning(ctx context.Context, taskID int64, executionRecordID int64, runningToken string, exclusive bool, ttl time.Duration) (bool, error) {
	if runningToken == "" || executionRecordID == 0 {
		return false, nil
	}
	exclusiveValue := "0"
	if exclusive {
		exclusiveValue = "1"
	}
	result, err := r.redisClient.Client.Eval(
		ctx,
		r.registerRunningScript,
		[]string{fmt.Sprintf(r.taskLockKeyFormat, taskID)},
		r.exclusiveField,
		"running:"+strconv.FormatInt(executionRecordID, 10),
		runningToken,
		int64(math.Ceil(ttl.Seconds())),
		exclusiveValue,
	).Int64()
	if err != nil {
		return false, err
	}
	return result == 1, nil
}

func (r *TaskLockRepo) RefreshRunning(ctx context.Context, taskID int64, executionRecordID int64, runningToken string, exclusive bool, ttl time.Duration) (bool, error) {
	if runningToken == "" || executionRecordID == 0 {
		return false, nil
	}
	exclusiveValue := "0"
	if exclusive {
		exclusiveValue = "1"
	}
	result, err := r.redisClient.Client.Eval(
		ctx,
		r.refreshRunningScript,
		[]string{fmt.Sprintf(r.taskLockKeyFormat, taskID)},
		r.exclusiveField,
		"running:"+strconv.FormatInt(executionRecordID, 10),
		runningToken,
		int64(math.Ceil(ttl.Seconds())),
		exclusiveValue,
	).Int64()
	if err != nil {
		return false, err
	}
	return result == 1, nil
}

func (r *TaskLockRepo) ReleaseRunning(ctx context.Context, taskID int64, executionRecordID int64, runningToken string, exclusive bool) error {
	if runningToken == "" {
		return nil
	}
	exclusiveValue := "0"
	if exclusive {
		exclusiveValue = "1"
	}
	return r.redisClient.Client.Eval(
		ctx,
		r.releaseRunningScript,
		[]string{fmt.Sprintf(r.taskLockKeyFormat, taskID)},
		r.exclusiveField,
		"running:"+strconv.FormatInt(executionRecordID, 10),
		runningToken,
		exclusiveValue,
	).Err()
}

func (r *TaskLockRepo) MapRunning(ctx context.Context, taskID int64, executionRecordIDs []int64) (map[int64]bool, error) {
	result := make(map[int64]bool, len(executionRecordIDs))
	if len(executionRecordIDs) == 0 {
		return result, nil
	}
	fields := make([]string, 0, len(executionRecordIDs))
	for _, id := range executionRecordIDs {
		result[id] = false
		fields = append(fields, "running:"+strconv.FormatInt(id, 10))
	}
	values, err := r.redisClient.Client.HMGet(ctx, fmt.Sprintf(r.taskLockKeyFormat, taskID), fields...).Result()
	if err != nil {
		return nil, err
	}
	for i, value := range values {
		if value != nil {
			result[executionRecordIDs[i]] = true
		}
	}
	return result, nil
}
