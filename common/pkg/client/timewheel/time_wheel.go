package timewheel

import (
	"common/proto/gen/common"
	"context"
	"fmt"
	"sync"
	"time"
)

// TimeWheel 是单层时间轮，适合高频短周期任务调度。
type TimeWheel struct {
	mutex       sync.Mutex
	tick        time.Duration
	slots       []map[string]*Task
	tasks       map[string]*Task
	currentTick int64
	currentTime time.Time
	stop        context.CancelFunc
	running     bool
}

// NewTimeWheel 创建时间轮。
func NewTimeWheel(config *common.TimeWheel) (*TimeWheel, error) {
	if config == nil {
		return nil, fmt.Errorf("time wheel config is required")
	}
	if config.GetInterval() == nil {
		return nil, fmt.Errorf("time wheel interval is required")
	}
	interval := config.GetInterval().AsDuration()
	if interval <= 0 {
		return nil, fmt.Errorf("time wheel interval must be positive")
	}
	if config.GetWheelSlots() == 0 {
		return nil, fmt.Errorf("time wheel slot count must be positive")
	}

	slots := make([]map[string]*Task, config.GetWheelSlots())
	for index := range slots {
		slots[index] = make(map[string]*Task)
	}

	return &TimeWheel{
		tick:        interval,
		slots:       slots,
		tasks:       make(map[string]*Task),
		currentTime: time.Now(),
	}, nil
}

// Start 启动时间轮自动推进。
func (w *TimeWheel) Start(ctx context.Context) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	w.mutex.Lock()
	if w.running {
		w.mutex.Unlock()
		return nil
	}

	runCtx, cancel := context.WithCancel(ctx)
	ticker := time.NewTicker(w.tick)
	w.stop = cancel
	w.running = true
	w.mutex.Unlock()

	go func() {
		defer ticker.Stop()

		for {
			select {
			case <-runCtx.Done():
				return
			case now := <-ticker.C:
				for _, task := range w.Advance(now) {
					if task.Job != nil {
						go func(item *Task) {
							_ = item.Job(runCtx, item)
						}(task)
					}
				}
			}
		}
	}()

	return nil
}

// Stop 停止时间轮自动推进。
func (w *TimeWheel) Stop(ctx context.Context) error {
	w.mutex.Lock()
	if !w.running {
		w.mutex.Unlock()
		return nil
	}

	stop := w.stop
	w.stop = nil
	w.running = false
	w.mutex.Unlock()

	if stop != nil {
		stop()
	}
	return nil
}

// Add 新增或替换任务，同一个任务 ID 只会保留一个调度位置。
func (w *TimeWheel) Add(task *Task) error {
	if task == nil {
		return fmt.Errorf("time wheel task nil")
	}
	if task.ID == "" {
		return fmt.Errorf("time wheel task id empty")
	}

	w.mutex.Lock()
	defer w.mutex.Unlock()

	if oldTask, ok := w.tasks[task.ID]; ok {
		delete(w.slots[oldTask.slot], task.ID)
		delete(w.tasks, task.ID)
	}

	delayTicks := int64(0)
	if task.DueAt.After(w.currentTime) {
		delay := task.DueAt.Sub(w.currentTime)
		delayTicks = int64(delay / w.tick)
		if delay%w.tick != 0 {
			delayTicks++
		}
	}
	slotCount := int64(len(w.slots))
	task.slot = (w.currentTick + delayTicks) % slotCount
	task.rounds = delayTicks / slotCount
	w.slots[task.slot][task.ID] = task
	w.tasks[task.ID] = task
	return nil
}

// Remove 移除任务，任务不存在时保持幂等。
func (w *TimeWheel) Remove(taskID string) bool {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	task, ok := w.tasks[taskID]
	if !ok {
		return false
	}
	delete(w.slots[task.slot], taskID)
	delete(w.tasks, taskID)
	return true
}

// Advance 推进时间轮并返回所有到期任务。
func (w *TimeWheel) Advance(now time.Time) []*Task {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	if now.Before(w.currentTime) {
		return nil
	}

	dueTasks := make([]*Task, 0)
	for !w.currentTime.After(now) {
		slot := w.currentTick % int64(len(w.slots))
		for taskID, task := range w.slots[slot] {
			if task.rounds > 0 {
				task.rounds--
				continue
			}
			delete(w.slots[slot], taskID)
			delete(w.tasks, taskID)
			dueTasks = append(dueTasks, task)
		}
		w.currentTick++
		w.currentTime = w.currentTime.Add(w.tick)
	}
	return dueTasks
}

// Len 返回当前调度中的任务数量。
func (w *TimeWheel) Len() int {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	return len(w.tasks)
}

// DelayTicks 将任务到期时间转换成相对当前时间轮的 tick 偏移。
func (w *TimeWheel) DelayTicks(dueAt time.Time) int64 {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	if !dueAt.After(w.currentTime) {
		return 0
	}
	delay := dueAt.Sub(w.currentTime)
	ticks := int64(delay / w.tick)
	if delay%w.tick != 0 {
		ticks++
	}
	return ticks
}
