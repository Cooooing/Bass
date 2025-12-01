package pool

import (
	"context"
	"sync"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type Task func(ctx context.Context)

type Pool struct {
	minWorkers int
	maxWorkers int
	keepAlive  time.Duration

	jobChan chan Task
	ctx     context.Context
	cancel  context.CancelFunc

	mu      sync.Mutex
	running int
	wg      sync.WaitGroup
	log     *log.Helper
}

func NewPool(log *log.Helper, minWorkers, maxWorkers int, keepAlive time.Duration) *Pool {
	ctx, cancel := context.WithCancel(context.Background())

	p := &Pool{
		minWorkers: minWorkers,
		maxWorkers: maxWorkers,
		keepAlive:  keepAlive,
		jobChan:    make(chan Task),
		ctx:        ctx,
		cancel:     cancel,
		log:        log,
	}

	// 初始化最小 worker
	for i := 0; i < p.minWorkers; i++ {
		p.startWorker()
	}

	return p
}

func (p *Pool) Submit(task Task) {
	select {
	case <-p.ctx.Done():
		return
	default:
	}

	p.mu.Lock()
	if p.running < p.maxWorkers {
		p.startWorker()
	}
	p.mu.Unlock()

	p.jobChan <- task
}

func (p *Pool) startWorker() {
	p.running++
	p.wg.Add(1)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				p.log.Errorf("worker panic: %v", r)
			}
		}()
		defer func() {
			p.mu.Lock()
			p.running--
			p.mu.Unlock()
			p.wg.Done()
		}()

		timer := time.NewTimer(p.keepAlive)
		defer timer.Stop()

		for {
			select {
			case <-p.ctx.Done():
				return

			case task := <-p.jobChan:
				// 执行任务
				task(p.ctx)

				// 重置 keepalive
				timer.Stop()
				timer.Reset(p.keepAlive)

			case <-timer.C:
				// idle 超过 keepAlive 退出，但要保证最小 worker 数
				p.mu.Lock()
				canExit := p.running > p.minWorkers
				p.mu.Unlock()

				if canExit {
					return
				}

				// 不退出则继续等待
				timer.Reset(p.keepAlive)
			}
		}
	}()
}

func (p *Pool) Stop() {
	p.cancel()
	p.wg.Wait()
}
