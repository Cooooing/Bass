package timewheel

import (
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
}

// NewTimeWheel 创建时间轮。
func NewTimeWheel(tick time.Duration, slotCount int, startAt time.Time) (*TimeWheel, error) {
	if tick <= 0 {
		return nil, fmt.Errorf("time wheel tick must be positive")
	}
	if slotCount <= 0 {
		return nil, fmt.Errorf("time wheel slot count must be positive")
	}

	slots := make([]map[string]*Task, slotCount)
	for index := range slots {
		slots[index] = make(map[string]*Task)
	}

	return &TimeWheel{
		tick:        tick,
		slots:       slots,
		tasks:       make(map[string]*Task),
		currentTime: startAt,
	}, nil
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

	delayTicks := w.delayTicks(task.DueAt)
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

	return w.delayTicks(dueAt)
}

func (w *TimeWheel) delayTicks(dueAt time.Time) int64 {
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
