package model

import "time"

type Account struct {
	ID           int64
	UserID       int64
	Balance      int64
	TotalIncome  int64
	TotalExpense int64
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
	DeletedAt    *time.Time
}

func (a *Account) CanAdd(amount int64) bool {
	return a != nil && amount > 0
}

func (a *Account) CanDeduct(amount int64) bool {
	return a != nil && amount > 0 && a.Balance >= amount
}
