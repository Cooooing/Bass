package usecase

import (
	"common/pkg/constant"
	"context"
	"game_idle/internal/biz/repo"
	"log/slog"
	"sync"
	"time"
)

const (
	metadataCacheRefreshInterval = time.Hour
	metadataCacheRefreshJitter   = 5 * time.Minute
)

type MetadataCacheUsecase struct {
	logger     *slog.Logger
	actionRepo repo.ActionRepo
	regionRepo repo.RegionRepo
	itemRepo   repo.ItemRepo
	recipeRepo repo.RecipeRepo
	mutex      sync.Mutex
	stop       context.CancelFunc
	running    bool
}

func NewMetadataCacheUsecase(
	logger *slog.Logger,
	actionRepo repo.ActionRepo,
	regionRepo repo.RegionRepo,
	itemRepo repo.ItemRepo,
	recipeRepo repo.RecipeRepo,
) *MetadataCacheUsecase {
	return &MetadataCacheUsecase{
		logger:     logger,
		actionRepo: actionRepo,
		regionRepo: regionRepo,
		itemRepo:   itemRepo,
		recipeRepo: recipeRepo,
	}
}

// Start 启动游戏元数据本地缓存定时刷新。
func (u *MetadataCacheUsecase) Start(ctx context.Context) error {
	u.mutex.Lock()
	if u.running {
		u.mutex.Unlock()
		return nil
	}
	runCtx, cancel := context.WithCancel(ctx)
	u.stop = cancel
	u.running = true
	u.mutex.Unlock()

	go u.refreshLoop(runCtx)
	return nil
}

// Stop 停止游戏元数据本地缓存定时刷新。
func (u *MetadataCacheUsecase) Stop(ctx context.Context) error {
	_ = ctx
	u.mutex.Lock()
	if !u.running {
		u.mutex.Unlock()
		return nil
	}
	stop := u.stop
	u.stop = nil
	u.running = false
	u.mutex.Unlock()
	stop()
	return nil
}

// RefreshRedis 从数据库重建 Redis 元数据缓存。
func (u *MetadataCacheUsecase) RefreshRedis(ctx context.Context) error {
	if _, err := u.regionRepo.Refresh(ctx); err != nil {
		return err
	}
	if _, err := u.itemRepo.Refresh(ctx); err != nil {
		return err
	}
	if _, err := u.recipeRepo.Refresh(ctx); err != nil {
		return err
	}
	if _, err := u.actionRepo.Refresh(ctx); err != nil {
		return err
	}
	return nil
}

func (u *MetadataCacheUsecase) refreshLoop(ctx context.Context) {
	timer := time.NewTimer(u.refreshDelay())
	defer timer.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			u.refreshLocal(ctx)
			timer.Reset(metadataCacheRefreshInterval + u.refreshDelay())
		}
	}
}

func (u *MetadataCacheUsecase) refreshLocal(ctx context.Context) {
	if err := u.regionRepo.RefreshLocal(ctx); err != nil {
		u.logger.WarnContext(ctx, "game idle refresh region metadata cache failed", constant.LogFieldErr, err)
	}
	if err := u.itemRepo.RefreshLocal(ctx); err != nil {
		u.logger.WarnContext(ctx, "game idle refresh item metadata cache failed", constant.LogFieldErr, err)
	}
	if err := u.recipeRepo.RefreshLocal(ctx); err != nil {
		u.logger.WarnContext(ctx, "game idle refresh recipe metadata cache failed", constant.LogFieldErr, err)
	}
	if err := u.actionRepo.RefreshLocal(ctx); err != nil {
		u.logger.WarnContext(ctx, "game idle refresh action metadata cache failed", constant.LogFieldErr, err)
	}
}

func (u *MetadataCacheUsecase) refreshDelay() time.Duration {
	return time.Duration(time.Now().UnixNano() % int64(metadataCacheRefreshJitter))
}
