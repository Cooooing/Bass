package doamin

import (
	"bytes"
	"context"
	"image/png"
	domainbase "user/internal/biz/base"

	"github.com/MuhammadSaim/goavatar"
)

type UserDomain struct {
	*domainbase.BaseDomain
}

func NewUserDomain(
	base *domainbase.BaseDomain) (*UserDomain, error) {
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
