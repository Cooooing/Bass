package doamin

import (
	"bytes"
	"context"
	"image/png"
	"user/internal/conf"

	"github.com/MuhammadSaim/goavatar"
	"github.com/go-kratos/kratos/v2/log"
)

type UserDomain struct {
	conf *conf.Bootstrap
	log  *log.Helper
}

func NewUserDomain(
	conf *conf.Bootstrap,
	logger log.Logger,
) (*UserDomain, error) {
	return &UserDomain{
		conf: conf,
		log:  log.NewHelper(logger),
	}, nil
}

func (s *UserDomain) Avatar(ctx context.Context, name string) ([]byte, error) {
	buf := &bytes.Buffer{}
	avatar := goavatar.Make(name, goavatar.WithSize(512))
	err := png.Encode(buf, avatar)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
