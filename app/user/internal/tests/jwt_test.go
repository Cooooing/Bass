package tests

import (
	"testing"
	"time"
	"user/internal/biz/doamin"
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
	service := doamin.NewTokenService(c)
	token, err := service.VerityCodeAccountTokenGen.Generate(model.TokenVerityCodeAccount{
		Account: "2222",
	}, 5*time.Minute)
	if err != nil {
		t.Error(err)
	}
	t.Log(token)
	time.Sleep(2 * time.Second)
	claims, err := service.VerityCodeAccountTokenGen.Parse(token)
	if err != nil {
		t.Error(err)
	}
	t.Log(claims)
}
