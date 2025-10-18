package model

import (
	"common/pkg/util"
	"user/internal/data/ent/gen"
)

type User gen.User

func (u *User) PasswordEncrypt() error {
	password, err := util.HashPassword(u.Password)
	u.Password = password
	return err
}

func (u *User) PasswordVerify(password string) bool {
	return util.VerifyPassword(u.Password, password)
}
