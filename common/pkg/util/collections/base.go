package collections

// Lockable 提供线程安全的锁操作接口
type Lockable interface {
	Lock()
	Unlock()
	RLock()
	RUnlock()
}
