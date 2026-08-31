package repo

import (
	"context"
	"fmt"
	"game_idle/internal/biz/model"
	bizrepo "game_idle/internal/biz/repo"
	"game_idle/internal/enum"
	"runtime"
	"sync"
	"time"
)

var _ bizrepo.CharacterStateRepo = (*CharacterStateCache)(nil)

// CharacterStateCache 保存单体模式下的角色运行时热状态。
type CharacterStateCache struct {
	mutex sync.RWMutex

	// key: characterID，value: 角色完整运行时状态；全局锁只保护该索引表，不参与角色状态计算。
	characters map[int64]*cachedCharacterState

	// executorMutex 只保护 mailbox 和 worker 运行信息；单个角色状态由各自的 mailbox 串行化。
	executorMutex     sync.Mutex
	mailboxes         map[int64]*characterStateMailbox
	readyMailboxes    chan *characterStateMailbox
	workerCount       int
	maxWorkerCount    int
	workerIdleTimeout time.Duration
}

// cachedCharacterState 聚合单个角色的全部本地热状态，角色锁用于保护这些状态的一致读写。
type cachedCharacterState struct {
	mutex sync.RWMutex

	// value: 角色行动队列缓存；nil 表示未回表加载，Items 为空表示队列确实为空。
	queue *cachedActionQueue
	// value: 角色背包缓存；nil 表示未回表加载，items 为空表示背包确实为空。
	backpack *cachedBackpack
	// value: 角色能力缓存；nil 表示未回表加载，items 为空表示能力确实为空。
	abilities *cachedAbilities
	// value: 角色在线会话缓存；nil 表示当前进程内未在线。
	session *characterSession
}

// cachedActionQueue 保存角色行动队列快照，队首会被时间轮调度。
type cachedActionQueue struct {
	queue *model.CharacterActionQueue
}

// cachedBackpack 保存角色背包快照和变更计数，计数达到阈值时会触发刷库。
type cachedBackpack struct {
	items          map[string]*model.CharacterItem
	operationCount int64
}

// cachedAbilities 保存角色能力快照和变更计数，升级或计数达到阈值时会触发刷库。
type cachedAbilities struct {
	items          map[enum.Ability]*model.CharacterAbility
	operationCount int64
}

// characterSession 保存当前进程内的在线会话，断线或踢下线时会被清理。
type characterSession struct {
	sessionID string
	expiresAt time.Time
}

// characterStateMailbox 是单角色的待执行操作队列，保证同一角色命令按提交顺序执行。
type characterStateMailbox struct {
	mutex       sync.Mutex
	characterID int64
	jobs        []*characterStateJob
	scheduled   bool
}

// characterStateJob 是一次需要进入单角色串行上下文的状态操作。
type characterStateJob struct {
	ctx       context.Context
	operation bizrepo.CharacterStateOperation
	done      chan error
}

// characterStateTx 持有角色锁和状态草稿，Commit 或 Rollback 后才会释放锁。
type characterStateTx struct {
	character *cachedCharacterState
	draft     *bizrepo.CharacterStateDraft
	closed    bool
}

func NewCharacterStateCache() *CharacterStateCache {
	maxWorkerCount := runtime.GOMAXPROCS(0) * 32
	if maxWorkerCount < 4 {
		maxWorkerCount = 4
	}
	return &CharacterStateCache{
		characters:        make(map[int64]*cachedCharacterState),
		mailboxes:         make(map[int64]*characterStateMailbox),
		readyMailboxes:    make(chan *characterStateMailbox, 4096),
		maxWorkerCount:    maxWorkerCount,
		workerIdleTimeout: 30 * time.Second,
	}
}

// Execute 将状态操作投递到角色 mailbox，调用方会阻塞等待操作完成。
func (c *CharacterStateCache) Execute(ctx context.Context, characterID int64, operation bizrepo.CharacterStateOperation) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	job := &characterStateJob{
		ctx:       ctx,
		operation: operation,
		done:      make(chan error, 1),
	}
	c.enqueue(job, characterID)
	select {
	case err := <-job.done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// enqueue 只在 mailbox 未被 worker 接管时投递一次，后续任务由同一个 worker 连续消费。
func (c *CharacterStateCache) enqueue(job *characterStateJob, characterID int64) {
	mailbox := c.mailbox(characterID)
	mailbox.mutex.Lock()
	mailbox.jobs = append(mailbox.jobs, job)
	shouldSchedule := !mailbox.scheduled
	if shouldSchedule {
		mailbox.scheduled = true
	}
	mailbox.mutex.Unlock()

	if shouldSchedule {
		c.readyMailboxes <- mailbox
		c.scaleWorkers()
	}
}

// mailbox 获取或创建角色 mailbox；创建过程使用 executorMutex 保护全局索引。
func (c *CharacterStateCache) mailbox(characterID int64) *characterStateMailbox {
	c.executorMutex.Lock()
	defer c.executorMutex.Unlock()

	mailbox := c.mailboxes[characterID]
	if mailbox == nil {
		mailbox = &characterStateMailbox{characterID: characterID}
		c.mailboxes[characterID] = mailbox
	}
	return mailbox
}

// scaleWorkers 根据积压 mailbox 增加 worker，空闲 worker 会在超时后退出。
func (c *CharacterStateCache) scaleWorkers() {
	c.executorMutex.Lock()
	defer c.executorMutex.Unlock()

	if c.workerCount >= c.maxWorkerCount {
		return
	}
	if c.workerCount > 0 && len(c.readyMailboxes) == 0 {
		return
	}
	c.workerCount++
	go c.runWorker()
}

// runWorker 消费已经就绪的角色 mailbox；同一 mailbox 一次只会被一个 worker 执行。
func (c *CharacterStateCache) runWorker() {
	timer := time.NewTimer(c.workerIdleTimeout)
	defer timer.Stop()

	for {
		select {
		case mailbox := <-c.readyMailboxes:
			if !timer.Stop() {
				select {
				case <-timer.C:
				default:
				}
			}
			for c.runMailbox(mailbox) {
				select {
				case c.readyMailboxes <- mailbox:
					mailbox = nil
				default:
				}
				if mailbox == nil {
					break
				}
			}
			timer.Reset(c.workerIdleTimeout)
		case <-timer.C:
			c.executorMutex.Lock()
			shouldStop := len(c.readyMailboxes) == 0
			if shouldStop {
				c.workerCount--
			}
			c.executorMutex.Unlock()
			if shouldStop {
				return
			}
			timer.Reset(c.workerIdleTimeout)
		}
	}
}

// runMailbox 执行单角色队列里的一个任务，并返回该角色是否还有后续任务。
func (c *CharacterStateCache) runMailbox(mailbox *characterStateMailbox) bool {
	mailbox.mutex.Lock()
	if len(mailbox.jobs) == 0 {
		mailbox.scheduled = false
		mailbox.mutex.Unlock()
		return false
	}
	job := mailbox.jobs[0]
	copy(mailbox.jobs, mailbox.jobs[1:])
	mailbox.jobs[len(mailbox.jobs)-1] = nil
	mailbox.jobs = mailbox.jobs[:len(mailbox.jobs)-1]
	mailbox.mutex.Unlock()

	job.done <- c.runJob(job)

	mailbox.mutex.Lock()
	hasNext := len(mailbox.jobs) > 0
	if !hasNext {
		mailbox.scheduled = false
	}
	mailbox.mutex.Unlock()
	return hasNext
}

// runJob 隔离单个状态操作的 panic，避免 worker 因业务异常退出。
func (c *CharacterStateCache) runJob(job *characterStateJob) (err error) {
	if err = job.ctx.Err(); err != nil {
		return err
	}
	defer func() {
		if recovered := recover(); recovered != nil {
			err = fmt.Errorf("game idle character state operation panic: %v", recovered)
		}
	}()
	return job.operation(job.ctx)
}

// Begin 创建角色状态事务草稿；调用方只修改草稿，提交前不会影响真实热状态。
func (c *CharacterStateCache) Begin(ctx context.Context, characterID int64) (bizrepo.CharacterStateTx, error) {
	character := c.character(characterID)
	character.mutex.Lock()
	if character.queue == nil {
		character.mutex.Unlock()
		return nil, fmt.Errorf("game idle character action queue state missing: %d", characterID)
	}
	if character.backpack == nil {
		character.mutex.Unlock()
		return nil, fmt.Errorf("game idle character backpack state missing: %d", characterID)
	}
	queueItems := make([]*model.CharacterActionQueueItem, len(character.queue.queue.Items))
	for index, item := range character.queue.queue.Items {
		copyItem := *item
		queueItems[index] = &copyItem
	}
	backpackItems := make(map[string]*model.CharacterItem, len(character.backpack.items))
	for itemID, item := range character.backpack.items {
		copyItem := *item
		backpackItems[itemID] = &copyItem
	}
	draft := &bizrepo.CharacterStateDraft{
		CharacterID:            characterID,
		Queue:                  &model.CharacterActionQueue{CharacterID: character.queue.queue.CharacterID, Items: queueItems},
		BackpackItems:          backpackItems,
		BackpackOperationCount: character.backpack.operationCount,
	}
	if character.abilities != nil {
		draft.Abilities = make(map[enum.Ability]*model.CharacterAbility, len(character.abilities.items))
		for abilityID, ability := range character.abilities.items {
			copyAbility := *ability
			draft.Abilities[abilityID] = &copyAbility
		}
		draft.AbilityOperationCount = character.abilities.operationCount
	}
	return &characterStateTx{
		character: character,
		draft:     draft,
	}, nil
}

// Clear 清理指定角色的全部热状态，通常在离线超时或服务停止刷库后调用。
func (c *CharacterStateCache) Clear(ctx context.Context, characterID int64) error {
	c.mutex.Lock()
	delete(c.characters, characterID)
	c.mutex.Unlock()
	return nil
}

func (tx *characterStateTx) Draft() *bizrepo.CharacterStateDraft {
	return tx.draft
}

// Commit 用事务草稿覆盖角色热状态，并释放角色锁。
func (tx *characterStateTx) Commit(ctx context.Context) error {
	if tx.closed {
		return nil
	}
	queueItems := make([]*model.CharacterActionQueueItem, len(tx.draft.Queue.Items))
	for index, item := range tx.draft.Queue.Items {
		copyItem := *item
		queueItems[index] = &copyItem
	}
	tx.character.queue.queue = &model.CharacterActionQueue{CharacterID: tx.draft.Queue.CharacterID, Items: queueItems}
	tx.character.backpack.items = make(map[string]*model.CharacterItem, len(tx.draft.BackpackItems))
	for itemID, item := range tx.draft.BackpackItems {
		copyItem := *item
		tx.character.backpack.items[itemID] = &copyItem
	}
	tx.character.backpack.operationCount = tx.draft.BackpackOperationCount
	if tx.draft.Abilities != nil {
		if tx.character.abilities == nil {
			tx.character.abilities = &cachedAbilities{}
		}
		tx.character.abilities.items = make(map[enum.Ability]*model.CharacterAbility, len(tx.draft.Abilities))
		for abilityID, ability := range tx.draft.Abilities {
			copyAbility := *ability
			tx.character.abilities.items[abilityID] = &copyAbility
		}
		tx.character.abilities.operationCount = tx.draft.AbilityOperationCount
	}
	tx.closed = true
	tx.character.mutex.Unlock()
	return nil
}

// Rollback 放弃事务草稿，并释放角色锁。
func (tx *characterStateTx) Rollback(ctx context.Context) error {
	if tx.closed {
		return nil
	}
	tx.closed = true
	tx.character.mutex.Unlock()
	return nil
}

func (c *CharacterStateCache) character(characterID int64) *cachedCharacterState {
	c.mutex.RLock()
	character := c.characters[characterID]
	c.mutex.RUnlock()
	if character != nil {
		return character
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()
	character = c.characters[characterID]
	if character == nil {
		character = &cachedCharacterState{}
		c.characters[characterID] = character
	}
	return character
}

func (c *CharacterStateCache) characterIfExists(characterID int64) *cachedCharacterState {
	c.mutex.RLock()
	character := c.characters[characterID]
	c.mutex.RUnlock()
	return character
}

func (c *CharacterStateCache) charactersSnapshot() map[int64]*cachedCharacterState {
	c.mutex.RLock()
	characters := make(map[int64]*cachedCharacterState, len(c.characters))
	for characterID, character := range c.characters {
		characters[characterID] = character
	}
	c.mutex.RUnlock()
	return characters
}
