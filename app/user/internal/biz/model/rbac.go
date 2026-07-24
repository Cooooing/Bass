package model

import (
	"time"
	"user/internal/enum"
)

type RbacRole struct {
	ID          int64
	Realm       enum.LoginRealm
	Code        string
	Name        string
	Description string
	BuiltIn     bool
	Enabled     bool
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}
type RbacPermission struct {
	ID          int64
	Realm       enum.LoginRealm
	Code        string
	Name        string
	Description string
	Enabled     bool
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}
