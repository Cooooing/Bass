package tests

import (
	"testing"
	"time"
	"user/internal/biz"
	"user/internal/biz/model"
	"user/internal/conf"

	"github.com/golang/protobuf/ptypes/duration"
)

var c = &conf.Bootstrap{
	Jwt: &conf.Jwt{
		Secret:      "123456",
		EmailExpire: &duration.Duration{Seconds: 1},
		Expires:     &duration.Duration{Seconds: 1},
	},
}

func TestJwt(t *testing.T) {
	service := biz.NewTokenService(c)
	token, err := service.EmailTokenGen.Generate(model.TokenEmail{
		UserId: 1,
		Email:  "2222",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(token)
	time.Sleep(2 * time.Second)
	claims, err := service.EmailTokenGen.Parse(token)
	if err != nil {
		t.Error(err)
	}
	t.Log(claims)
}
