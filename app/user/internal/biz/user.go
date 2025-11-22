package biz

import (
	"bytes"
	"context"
	"image/png"

	"github.com/MuhammadSaim/goavatar"
)

type UserDomain struct {
	*BaseDomain
}

func NewUserDomain(base *BaseDomain) (*UserDomain, error) {
	return &UserDomain{
		BaseDomain: base,
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
