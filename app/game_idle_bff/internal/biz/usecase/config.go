package usecase

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"game_idle_bff/internal/biz/model"
	"game_idle_bff/internal/biz/repo"
	"log/slog"
	"sort"
	"sync"
	"time"
)

const configCacheTTL = 60 * time.Minute

type ConfigUsecase struct {
	logger         *slog.Logger
	regionRepo     repo.RegionRepo
	actionRepo     repo.ActionRepo
	itemRepo       repo.ItemRepo
	cache          *model.GameConfig
	cacheExpiresAt time.Time
	lock           sync.RWMutex
}

func NewConfigUsecase(
	logger *slog.Logger,
	regionRepo repo.RegionRepo,
	actionRepo repo.ActionRepo,
	itemRepo repo.ItemRepo,
) *ConfigUsecase {
	return &ConfigUsecase{
		logger:     logger,
		regionRepo: regionRepo,
		actionRepo: actionRepo,
		itemRepo:   itemRepo,
	}
}

func (u *ConfigUsecase) Get(ctx context.Context) (*model.GameConfig, error) {
	row, err := u.current(ctx)
	if err != nil {
		return nil, err
	}
	out := *row
	out.Regions = append([]*model.RegionConfig(nil), row.Regions...)
	out.Actions = append([]*model.ActionConfig(nil), row.Actions...)
	out.Items = append([]*model.ItemConfig(nil), row.Items...)
	out.ServerTime = time.Now().Unix()
	return &out, nil
}

func (u *ConfigUsecase) Version(ctx context.Context) (*model.GameConfigVersion, error) {
	row, err := u.current(ctx)
	if err != nil {
		return nil, err
	}
	return &model.GameConfigVersion{
		ConfigVersion: row.ConfigVersion,
		ServerTime:    time.Now().Unix(),
	}, nil
}

func (u *ConfigUsecase) current(ctx context.Context) (*model.GameConfig, error) {
	now := time.Now()
	u.lock.RLock()
	if u.cache != nil && now.Before(u.cacheExpiresAt) {
		row := u.cache
		u.lock.RUnlock()
		return row, nil
	}
	u.lock.RUnlock()

	u.lock.Lock()
	defer u.lock.Unlock()
	now = time.Now()
	if u.cache != nil && now.Before(u.cacheExpiresAt) {
		return u.cache, nil
	}
	row, err := u.load(ctx)
	if err != nil {
		return nil, err
	}
	u.cache = row
	u.cacheExpiresAt = now.Add(configCacheTTL)
	return row, nil
}

func (u *ConfigUsecase) load(ctx context.Context) (*model.GameConfig, error) {
	regions, err := u.regionRepo.List(ctx)
	if err != nil {
		return nil, err
	}
	actions, err := u.actionRepo.List(ctx)
	if err != nil {
		return nil, err
	}
	items, err := u.itemRepo.List(ctx)
	if err != nil {
		return nil, err
	}
	u.sort(regions, actions, items)
	row := &model.GameConfig{
		Regions:    regions,
		Actions:    actions,
		Items:      items,
		ServerTime: time.Now().Unix(),
	}
	row.ConfigVersion = u.hash(row)
	return row, nil
}

func (u *ConfigUsecase) GetActionDetail(ctx context.Context, actionID string) (*model.ActionDetailConfig, error) {
	return u.actionRepo.GetDetail(ctx, actionID)
}

func (u *ConfigUsecase) sort(regions []*model.RegionConfig, actions []*model.ActionConfig, items []*model.ItemConfig) {
	sort.SliceStable(regions, func(i, j int) bool {
		if regions[i].Sort == regions[j].Sort {
			return regions[i].RegionID < regions[j].RegionID
		}
		return regions[i].Sort < regions[j].Sort
	})
	sort.SliceStable(actions, func(i, j int) bool {
		if actions[i].Sort == actions[j].Sort {
			return actions[i].ActionID < actions[j].ActionID
		}
		return actions[i].Sort < actions[j].Sort
	})
	sort.SliceStable(items, func(i, j int) bool {
		if items[i].Sort == items[j].Sort {
			return items[i].ItemID < items[j].ItemID
		}
		return items[i].Sort < items[j].Sort
	})
}

func (u *ConfigUsecase) hash(row *model.GameConfig) string {
	payload, err := json.Marshal(struct {
		Regions []*model.RegionConfig `json:"regions"`
		Actions []*model.ActionConfig `json:"actions"`
		Items   []*model.ItemConfig   `json:"items"`
	}{
		Regions: row.Regions,
		Actions: row.Actions,
		Items:   row.Items,
	})
	if err != nil {
		u.logger.Error("game idle bff config hash marshal failed", "err", err)
		return ""
	}
	sum := sha256.Sum256(payload)
	return "sha256:" + hex.EncodeToString(sum[:])
}
