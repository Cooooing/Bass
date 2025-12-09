package timewheel

import (
	"container/list"
	"errors"
	"runtime/debug"
	"time"

	"common/pkg/cutil/base/logger"
	"github.com/orcaman/concurrent-map/v2"
	"github.com/panjf2000/ants/v2"
)

// Timewheel 核心结构体
type Timewheel struct {
	interval          time.Duration // 时间轮的精度
	slots             []*list.List  // 时间轮每个位置存储的Task列表
	ticker            *time.Ticker  // 时间轮的计时器
	currentPos        int           // 时间轮当前的位置
	slotNums          int           // 时间轮的齿轮数 interval*slotNums就是时间轮转一圈走过的时间
	addTaskChannel    chan *Task
	removeTaskChannel chan *Task
	stopChannel       chan bool
	taskRecords       cmap.ConcurrentMap[string, *list.Element] // Map结构来存储Task对象，key是Task.key，value是Task在双向链表中的存储对象list.Element
	isRunning         bool
	pool              *ants.Pool // 协程池
	poolSize          int
	poolMaxSize       int
}

// Job 需要执行的Job的函数结构体
type Job func(task *Task)

// Task 时间轮上需要执行的任务
type Task struct {
	Key         string        // 用来标识task对象，是唯一的
	Interval    time.Duration // 任务周期
	Times       int           // 任务需要执行的次数，如果需要一直执行，设置成-1
	Job         Job           // 任务需要执行的Job
	createdTime time.Time     // 任务的创建时间
	pos         int           // 任务在轮的位置
	circle      int           // 任务需要在轮走多少圈才能执行
}

// ErrDuplicateTaskKey is an definedError for duplicate task key
var ErrDuplicateTaskKey = errors.New("duplicate task key")

// ErrTaskKeyNotFount is an definedError when task key is not found
var ErrTaskKeyNotFount = errors.New("task key doesn't existed in task list, please check your input")

// NewTimewheel 初始化一个Timewheel对象
func NewTimewheel(interval time.Duration, slotNum int, poolSize int, maxPoolSize int) *Timewheel {
	if interval <= 0 || slotNum <= 0 {
		return nil
	}
	pool, err := ants.NewPool(
		poolSize,
		ants.WithPreAlloc(true),
		ants.WithNonblocking(true),
	)
	if err != nil {
		logger.Error("init [timewheel] ants pool failed: %v", err)
	}
	tw := &Timewheel{
		interval:          interval,
		slots:             make([]*list.List, slotNum),
		currentPos:        0,
		slotNums:          slotNum,
		addTaskChannel:    make(chan *Task, 16),
		removeTaskChannel: make(chan *Task, 16),
		stopChannel:       make(chan bool),
		taskRecords:       cmap.New[*list.Element](),
		isRunning:         false,
		pool:              pool,
		poolSize:          poolSize,
		poolMaxSize:       maxPoolSize,
	}
	tw.initSlots()
	return tw
}

// Start 启动时间轮
func (tw *Timewheel) Start() {
	tw.ticker = time.NewTicker(tw.interval)
	tw.TunePool()
	go func(tw *Timewheel) {
		for {
			select {
			case <-tw.ticker.C:
				tw.checkAndRunTask()
			case task := <-tw.addTaskChannel:
				tw.addTask(task, false)
			case task := <-tw.removeTaskChannel:
				tw.removeTask(task)
			case <-tw.stopChannel:
				tw.ticker.Stop()
				return
			}
		}
	}(tw)
	tw.isRunning = true
}

// Stop 关闭时间轮
func (tw *Timewheel) Stop() {
	tw.stopChannel <- true
	tw.isRunning = false
	tw.pool.Release()
}

// IsRunning 检查全局时间轮是否在正常运行
func (tw *Timewheel) IsRunning() bool {
	return tw.isRunning
}

// Exist 检查任务是否存在
func (tw *Timewheel) Exist(key string) bool {
	_, ok := tw.taskRecords.Get(key)
	return ok
}

// GetTaskTimes 获取任务剩余执行次数
func (tw *Timewheel) GetTaskTimes(key string) int {
	t, ok := tw.taskRecords.Get(key)
	if !ok {
		return 0
	}
	return t.Value.(*Task).Times
}

// AddTask 向时间轮添加固定周期任务
func (tw *Timewheel) AddTask(task *Task) error {
	if task.Interval <= 0 || task.Key == "" {
		return errors.New("invalid task params")
	}

	// 检查Task.Key是否已经存在
	_, ok := tw.taskRecords.Get(task.Key)
	if ok {
		return ErrDuplicateTaskKey
	}
	task.createdTime = time.Now()
	tw.addTaskChannel <- task
	return nil
}

// RemoveTask 从时间轮删除任务
func (tw *Timewheel) RemoveTask(key string) error {
	if key == "" {
		return nil
	}

	// 检查该Task是否存在
	val, ok := tw.taskRecords.Get(key)
	if !ok {
		return ErrTaskKeyNotFount
	}

	task := val.Value.(*Task)
	tw.removeTaskChannel <- task
	return nil
}

// 初始化时间轮，每个轮上的卡槽用一个双向队列表示，便于插入和删除
func (tw *Timewheel) initSlots() {
	for i := 0; i < tw.slotNums; i++ {
		tw.slots[i] = list.New()
	}
}

// 检查该轮点位上的Task，看哪个需要执行
func (tw *Timewheel) checkAndRunTask() {

	// 获取该轮位置的双向链表
	currentList := tw.slots[tw.currentPos]

	if currentList != nil {
		for item := currentList.Front(); item != nil; {
			task, ok := item.Value.(*Task)
			if !ok {
				item = item.Next()
				continue
			}
			next := item.Next()
			if task.circle > 0 {
				task.circle--
			} else {
				if task.Job != nil {
					// 使用协程池执行任务
					err := tw.pool.Submit(func() {
						defer func() {
							if err := recover(); err != nil {
								stack := debug.Stack()
								logger.Error("task %v panic: %v %s\n", task.Key, err, string(stack))
							}
						}()
						task.Job(task)
					})
					if err != nil {
						logger.Error("task %v submit failed: %v\n", task.Key, err)
					}
				} else {
					logger.Warn("The task %s don't have job to run\n", task.Key)
				}

				tw.taskRecords.Remove(task.Key)
				currentList.Remove(item)

				if task.Times < 0 {
					tw.addTask(task, true) // 无限次，继续添加
				} else if task.Times > 1 {
					task.Times--
					tw.addTask(task, true) // 剩余次数大于1，继续添加
				}

			}
			item = next
		}
	}

	// 轮前进一步
	tw.currentPos = (tw.currentPos + 1) % tw.slotNums
}

// 添加任务的内部函数
func (tw *Timewheel) addTask(task *Task, byInterval bool) {
	var pos, circle int
	// 使用任务周期或创建时间生成
	if byInterval {
		pos, circle = tw.getPosAndCircleByInterval(task.Interval)
	} else {
		pos, circle = tw.getPosAndCircleByCreatedTime(task.createdTime, task.Interval)
	}

	task.circle = circle
	task.pos = pos

	element := tw.slots[pos].PushBack(task)
	tw.taskRecords.Set(task.Key, element)
}

// 删除任务的内部函数
func (tw *Timewheel) removeTask(task *Task) {
	val, ok := tw.taskRecords.Get(task.Key)
	if !ok {
		return
	}
	tw.taskRecords.Remove(task.Key)
	if t, ok := val.Value.(*Task); ok && t.pos < tw.slotNums {
		tw.slots[t.pos].Remove(val)
	}
}

// 该函数通过任务的周期来计算下次执行的位置和圈数
func (tw *Timewheel) getPosAndCircleByInterval(d time.Duration) (int, int) {
	delayMs := int(d.Milliseconds())
	intervalMs := int(tw.interval.Milliseconds())
	ticks := delayMs / intervalMs
	circle := ticks / tw.slotNums
	pos := (tw.currentPos + ticks) % tw.slotNums

	// 特殊case，当计算的位置和当前位置重叠时，因为当前位置已经走过了，所以circle需要减一
	if pos == tw.currentPos && circle != 0 {
		circle--
	}
	return pos, circle
}

// 该函数用任务的创建时间来计算下次执行的位置和圈数
func (tw *Timewheel) getPosAndCircleByCreatedTime(createdTime time.Time, d time.Duration) (int, int) {
	delayMs := int(d.Milliseconds())
	intervalMs := int(tw.interval.Milliseconds())
	ticksPassed := int(time.Since(createdTime).Milliseconds()) / intervalMs
	totalTicks := delayMs / intervalMs
	remainingTicks := totalTicks - ticksPassed

	if remainingTicks <= 0 {
		remainingTicks = 1 // 防止立即过期
	}

	circle := remainingTicks / tw.slotNums
	pos := (tw.currentPos + remainingTicks) % tw.slotNums

	// 特殊case，当计算的位置和当前位置重叠时，因为当前位置已经走过了，所以circle需要减一
	if pos == tw.currentPos && circle != 0 {
		circle--
	}

	return pos, circle
}

// TunePool 打印协程池状态并判断扩缩容
func (tw *Timewheel) TunePool() {

	task := &Task{
		Interval: time.Minute,
		Key:      "logAndTunePool",
		Times:    -1,
		Job: func(task *Task) {
			running := tw.pool.Running()
			poolCap := tw.pool.Cap()
			poolFree := tw.pool.Free()
			if running >= poolCap && poolCap < tw.poolMaxSize {
				newCap := poolCap * 2
				if newCap > tw.poolMaxSize {
					newCap = tw.poolMaxSize
				}
				tw.pool.Tune(newCap)
				logger.Info("[定期扩容] 协程池扩容: %d -> %d", poolCap, newCap)
			} else if poolFree > poolCap/2 && poolCap > tw.poolSize {
				newCap := poolCap / 2
				if newCap < tw.poolSize {
					newCap = tw.poolSize
				}
				tw.pool.Tune(newCap)
				logger.Info("[定期缩容] 协程池缩容: %d -> %d", poolCap, newCap)
			}
		},
	}

	err := tw.AddTask(task)
	if err != nil {
		logger.Error("add logAndTunePool task err: %v", err)
	}
}
