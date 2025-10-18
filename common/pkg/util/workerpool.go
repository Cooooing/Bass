package util

import (
	"context"
	"sync"

	"github.com/go-kratos/kratos/v2/log"
)

// Task 定义任务接口
type Task func()

type WorkerPool struct {
	log       *log.Helper
	taskCh    chan Task
	wg        sync.WaitGroup
	ctx       context.Context
	cancel    context.CancelFunc
	numWorker int
}

// NewWorkerPool 创建 WorkerPool
func NewWorkerPool(log *log.Helper, numWorker, queueSize int) *WorkerPool {
	ctx, cancel := context.WithCancel(context.Background())
	pool := &WorkerPool{
		taskCh:    make(chan Task, queueSize),
		ctx:       ctx,
		cancel:    cancel,
		numWorker: numWorker,
	}
	pool.start()
	return pool
}

// start 启动 worker
func (p *WorkerPool) start() {
	for i := 0; i < p.numWorker; i++ {
		p.wg.Add(1)
		go func(id int) {
			defer p.wg.Done()
			for {
				select {
				case <-p.ctx.Done():
					return
				case task, ok := <-p.taskCh:
					if !ok {
						return
					}
					p.safeExecute(task)
				}
			}
		}(i)
	}
}

// safeExecute 对单个任务做异常兜底
func (p *WorkerPool) safeExecute(task Task) {
	defer func() {
		if r := recover(); r != nil {
			p.log.Errorf("[WorkerPool] task panic recovered: %v", r)
		}
	}()
	task()
}

// Submit 提交任务
func (p *WorkerPool) Submit(task Task) bool {
	select {
	case p.taskCh <- task:
		return true
	case <-p.ctx.Done():
		return false
	}
}

// Stop 优雅停止工作池
func (p *WorkerPool) Stop() {
	p.cancel()
	close(p.taskCh)
	p.wg.Wait()
}
