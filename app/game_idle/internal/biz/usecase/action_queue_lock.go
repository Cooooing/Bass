package usecase

import (
	"sync"
)

// ActionQueueLocker 串行化同一角色的队列编辑与时间轮调度，避免队首任务重复或漏调度。
type ActionQueueLocker struct {
	locks sync.Map
}

func NewActionQueueLocker() *ActionQueueLocker {
	return &ActionQueueLocker{}
}

func (l *ActionQueueLocker) withCharacterLock(characterID int64, fn func() error) error {
	lock, _ := l.locks.LoadOrStore(characterID, &sync.Mutex{})
	mutex := lock.(*sync.Mutex)
	mutex.Lock()
	defer mutex.Unlock()
	return fn()
}
