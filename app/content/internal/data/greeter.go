package data

import (
	"content/internal/biz"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type greeterRepo struct {
	log *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(logger log.Logger) biz.GreeterRepo {
	return &greeterRepo{
		log: log.NewHelper(logger),
	}
}

func (r *greeterRepo) Save(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
	return g, nil
}

func (r *greeterRepo) Update(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
	return g, nil
}

func (r *greeterRepo) FindByID(context.Context, int64) (*biz.Greeter, error) {
	return nil, nil
}

func (r *greeterRepo) ListByHello(context.Context, string) ([]*biz.Greeter, error) {
	return nil, nil
}

func (r *greeterRepo) ListAll(context.Context) ([]*biz.Greeter, error) {
	return nil, nil
}
