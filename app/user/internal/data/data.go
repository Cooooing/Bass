package data

import (
	"common/api/gen/common"
	commonClient "common/pkg/client"
	utilent "common/pkg/util/ent"
	"context"
	bizbase "user/internal/biz/base"
	"user/internal/conf"
	"user/internal/data/client"
	"user/internal/data/gen"
	"user/internal/data/repo"

	"github.com/google/wire"
)

// DataProviderSet is data providers.
var DataProviderSet = wire.NewSet(
	client.NewDataBaseClient,
	ProvideRedis,
	ProvideConsul,
	ProvideNats,
	commonClient.NewConsulClient,
	commonClient.NewRedisClient,
	commonClient.NewNatsClient,

	ProvideTxRunner,

	repo.NewUserRepo,
	repo.NewUserRelationRepo,
	repo.NewUserPreferencesRepo,
	repo.NewUserPrivacyRepo,
	repo.NewUserLocationRepo,
	repo.NewUserTfaRepo,
	repo.NewUserCheckinRepo,
)

func ProvideRedis(c *conf.Bootstrap) *common.Redis {
	return c.Data.Redis
}

func ProvideConsul(c *conf.Bootstrap) *common.Consul {
	return c.Data.Consul
}

func ProvideNats(c *conf.Bootstrap) *common.Nats {
	return c.Data.Nats
}

func ProvideTxRunner(db *gen.Client) bizbase.TxRunner {
	starter := func(ctx context.Context) (utilent.Tx, error) {
		tx, err := db.Tx(ctx)
		if err != nil {
			return nil, err
		}
		return &genTxWrapper{tx: tx}, nil
	}
	return func(ctx context.Context, fn func(ctx context.Context) error) error {
		return utilent.WithTx(ctx, starter, fn)
	}
}

// genTxWrapper adapts gen.Tx to the utilent.Tx interface.
type genTxWrapper struct{ tx *gen.Tx }

func (w *genTxWrapper) Commit() error       { return w.tx.Commit() }
func (w *genTxWrapper) Rollback() error     { return w.tx.Rollback() }
func (w *genTxWrapper) Client() interface{} { return w.tx.Client() }
