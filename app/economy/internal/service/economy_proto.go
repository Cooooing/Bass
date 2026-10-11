package service

import (
	v1 "common/proto/gen/economy/v1"

	"economy/internal/biz/model"
	"economy/internal/enum"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type economyProtoMapper struct{}

func (economyProtoMapper) account(row *model.Account) *v1.EconomyAccount {
	if row == nil {
		return nil
	}
	account := &v1.EconomyAccount{Id: row.ID, UserId: row.UserID, Balance: row.Balance, FrozenBalance: row.FrozenBalance, TotalIncome: row.TotalIncome, TotalExpense: row.TotalExpense}
	if row.CreatedAt != nil {
		account.CreatedAt = timestamppb.New(*row.CreatedAt)
	}
	if row.UpdatedAt != nil {
		account.UpdatedAt = timestamppb.New(*row.UpdatedAt)
	}
	return account
}

func (economyProtoMapper) record(row *model.Record) *v1.EconomyRecord {
	if row == nil {
		return nil
	}
	record := &v1.EconomyRecord{Id: row.ID, TransactionNo: row.TransactionNo, UserId: row.UserID, RecordType: enum.EconomyRecordTypeMap.MustToProto(row.RecordType), Direction: enum.EconomyRecordDirectionMap.MustToProto(row.Direction), Amount: row.Amount, BalanceBefore: row.BalanceBefore, BalanceAfter: row.BalanceAfter, Remark: row.Remark}
	if row.CreatedAt != nil {
		record.CreatedAt = timestamppb.New(*row.CreatedAt)
	}
	return record
}
