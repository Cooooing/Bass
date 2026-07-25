package model

import (
	commonenum "common/pkg/enum"
	"time"
)

type RbacRole struct {
	ID          int64
	Realm       commonenum.LoginRealm
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
	Realm       commonenum.LoginRealm
	Code        string
	Name        string
	Description string
	Enabled     bool
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}
