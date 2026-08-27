package client

import (
	"common/pkg/constant"
	"common/proto/gen/common"
	"fmt"
	"log/slog"
	"runtime/debug"
	"sync"
	"time"

	"google.golang.org/protobuf/types/known/durationpb"
)

// WorkerPool 是可配置的通用协程池，core worker 常驻，临时 worker 按空闲时间回收。
type WorkerPool struct {
	logger      *slog.Logger
	tasks       chan func()
	closed      chan struct{}
	closeOnce   sync.Once
	mutex       sync.Mutex
	coreSize    int
	maxSize     int
	idleTimeout time.Duration
	nonblocking bool
	running     int
}

func NewWorkerPool(
	logger *slog.Logger,
	conf *common.WorkerPool,
) (*WorkerPool, func(), error) {
	if logger == nil {
		logger = slog.Default()
	}
	if conf == nil {
		conf = &common.WorkerPool{}
	}
	if conf.CoreSize == 0 {
		conf.CoreSize = 16
	}
	if conf.MaxSize == 0 {
		conf.MaxSize = conf.CoreSize
	}
	if conf.MaxSize < conf.CoreSize {
		return nil, nil, fmt.Errorf("worker pool max size must be greater than or equal to core size")
	}
	if conf.QueueSize == 0 {
		conf.QueueSize = conf.MaxSize
	}
	if conf.GetIdleTimeout() == nil || conf.GetIdleTimeout().AsDuration() <= 0 {
		conf.IdleTimeout = durationpb.New(time.Minute)
	}
	pool := &WorkerPool{
		logger:      logger,
		tasks:       make(chan func(), conf.QueueSize),
		closed:      make(chan struct{}),
		coreSize:    int(conf.CoreSize),
		maxSize:     int(conf.MaxSize),
		idleTimeout: conf.GetIdleTimeout().AsDuration(),
		nonblocking: conf.Nonblocking,
	}
	pool.running = pool.coreSize
	for index := 0; index < pool.coreSize; index++ {
		go pool.runCoreWorker()
	}
	return pool, func() {
		pool.Release()
	}, nil
}

// Submit 提交任务，队列满时按配置阻塞等待或立即返回错误。
func (p *WorkerPool) Submit(task func()) error {
	if task == nil {
		return fmt.Errorf("worker pool task is required")
	}
	select {
	case <-p.closed:
		return fmt.Errorf("worker pool closed")
	case p.tasks <- task:
		return nil
	default:
	}

	p.mutex.Lock()
	if p.running < p.maxSize {
		p.running++
		p.mutex.Unlock()
		go p.runTemporaryWorker(task)
		return nil
	}
	p.mutex.Unlock()

	if p.nonblocking {
		return fmt.Errorf("worker pool overload")
	}
	select {
	case <-p.closed:
		return fmt.Errorf("worker pool closed")
	case p.tasks <- task:
		return nil
	}
}

// Release 停止协程池，不再接收新的任务。
func (p *WorkerPool) Release() {
	p.closeOnce.Do(func() {
		close(p.closed)
	})
}

// Running 返回当前 worker 数量。
func (p *WorkerPool) Running() int {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.running
}

// Cap 返回最大 worker 数量。
func (p *WorkerPool) Cap() int {
	return p.maxSize
}

// Waiting 返回等待队列中的任务数量。
func (p *WorkerPool) Waiting() int {
	return len(p.tasks)
}

func (p *WorkerPool) runCoreWorker() {
	for {
		select {
		case <-p.closed:
			return
		case task := <-p.tasks:
			p.execute(task)
		}
	}
}

func (p *WorkerPool) runTemporaryWorker(task func()) {
	defer func() {
		p.mutex.Lock()
		p.running--
		p.mutex.Unlock()
	}()
	p.execute(task)
	timer := time.NewTimer(p.idleTimeout)
	defer timer.Stop()
	for {
		select {
		case <-p.closed:
			return
		case task := <-p.tasks:
			p.execute(task)
			if !timer.Stop() {
				select {
				case <-timer.C:
				default:
				}
			}
			timer.Reset(p.idleTimeout)
		case <-timer.C:
			return
		}
	}
}

func (p *WorkerPool) execute(task func()) {
	defer func() {
		if err := recover(); err != nil {
			p.logger.Error("worker panic recovered", constant.LogFieldKind, constant.LogKindSystem, "panic", err, "stack", string(debug.Stack()))
		}
	}()
	task()
}
