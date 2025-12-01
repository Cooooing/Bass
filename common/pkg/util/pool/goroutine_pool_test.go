package pool

import (
	"common/pkg"
	"context"
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

func TestGoroutinePool(t *testing.T) {
	minWorkers := 2
	maxWorkers := 5
	keepAlive := 2 * time.Second

	logger := pkg.NewLogger("test", "1.0.0", "debug", "debug", "")
	p := NewPool(log.NewHelper(logger), minWorkers, maxWorkers, keepAlive)

	var counter int32

	// 模拟任务函数
	task := func(ctx context.Context) {
		atomic.AddInt32(&counter, 1)
		fmt.Println("执行任务, 当前 counter:", atomic.LoadInt32(&counter))
		time.Sleep(500 * time.Millisecond) // 模拟工作负载
	}

	// 提交10个任务
	for i := 0; i < 10; i++ {
		p.Submit(task)
	}

	// 等待任务完成
	time.Sleep(3 * time.Second)

	// 检查 counter 是否正确
	if atomic.LoadInt32(&counter) != 10 {
		t.Errorf("预期 counter=10, 实际=%d", counter)
	}

	fmt.Println("当前 worker 数:", p.RunningWorkers())

	// 测试 ctx 取消
	ctxTask := func(ctx context.Context) {
		select {
		case <-ctx.Done():
			fmt.Println("任务被取消")
		case <-time.After(3 * time.Second):
			fmt.Println("任务完成")
		}
	}

	p.Submit(ctxTask)
	time.Sleep(500 * time.Millisecond)
	p.Stop() // 取消所有任务
	fmt.Println("池已停止")
}
