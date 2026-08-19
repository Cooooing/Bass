package service

import (
	"common/pkg/apperror"
	"common/proto/gen/common"
	cerrors "common/proto/gen/common/errors"
	v1 "common/proto/gen/economy/v1"
	"context"
	"economy/internal/biz/base"
	"economy/internal/biz/usecase"
	"economy/internal/enum"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PointsService struct {
	v1.UnimplementedEconomyServiceServer

	pointsUsecase *usecase.PointsUsecase
}

func NewPointsService(pointsUsecase *usecase.PointsUsecase) *PointsService {
	return &PointsService{pointsUsecase: pointsUsecase}
}

func (s *PointsService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterEconomyServiceServer(gs, s)
}

func (s *PointsService) RegisterHttp(hs *http.Server) {
}

func (s *PointsService) GetAccount(ctx context.Context, req *v1.GetEconomyAccount_Req) (*v1.GetEconomyAccount_Resp, error) {
	account, err := s.pointsUsecase.GetAccount(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}
	if account == nil {
		return &v1.GetEconomyAccount_Resp{}, nil
	}
	resp := &v1.EconomyAccount{
		Id:            account.ID,
		UserId:        account.UserID,
		Balance:       account.Balance,
		FrozenBalance: account.FrozenBalance,
		TotalIncome:   account.TotalIncome,
		TotalExpense:  account.TotalExpense,
	}
	if account.CreatedAt != nil {
		resp.CreatedAt = timestamppb.New(*account.CreatedAt)
	}
	if account.UpdatedAt != nil {
		resp.UpdatedAt = timestamppb.New(*account.UpdatedAt)
	}
	return &v1.GetEconomyAccount_Resp{Account: resp}, nil
}

func (s *PointsService) MapAccounts(ctx context.Context, req *v1.MapEconomyAccounts_Req) (*v1.MapEconomyAccounts_Resp, error) {
	accounts, err := s.pointsUsecase.MapAccounts(ctx, req.GetUserIds())
	if err != nil {
		return nil, err
	}
	reply := &v1.MapEconomyAccounts_Resp{Accounts: make(map[int64]*v1.EconomyAccount, len(accounts))}
	for userID, account := range accounts {
		if account == nil {
			reply.Accounts[userID] = nil
			continue
		}
		item := &v1.EconomyAccount{
			Id:            account.ID,
			UserId:        account.UserID,
			Balance:       account.Balance,
			FrozenBalance: account.FrozenBalance,
			TotalIncome:   account.TotalIncome,
			TotalExpense:  account.TotalExpense,
		}
		if account.CreatedAt != nil {
			item.CreatedAt = timestamppb.New(*account.CreatedAt)
		}
		if account.UpdatedAt != nil {
			item.UpdatedAt = timestamppb.New(*account.UpdatedAt)
		}
		reply.Accounts[userID] = item
	}
	return reply, nil
}

func (s *PointsService) AddPoints(ctx context.Context, req *v1.AddEconomyPoints_Req) (*v1.AddEconomyPoints_Resp, error) {
	recordType, ok := enum.EconomyRecordTypeMap.ToEnum(req.GetRecordType())
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_RECORD_TYPE_INVALID)
	}
	resp, err := s.pointsUsecase.AddPoints(ctx, &usecase.EconomyAddPointsReq{
		UserID:     req.GetUserId(),
		Amount:     req.GetAmount(),
		RecordType: recordType,
		Remark:     req.Remark,
	})
	if err != nil {
		return nil, err
	}
	var account *v1.EconomyAccount
	if resp.Account != nil {
		account = &v1.EconomyAccount{
			Id:            resp.Account.ID,
			UserId:        resp.Account.UserID,
			Balance:       resp.Account.Balance,
			FrozenBalance: resp.Account.FrozenBalance,
			TotalIncome:   resp.Account.TotalIncome,
			TotalExpense:  resp.Account.TotalExpense,
		}
		if resp.Account.CreatedAt != nil {
			account.CreatedAt = timestamppb.New(*resp.Account.CreatedAt)
		}
		if resp.Account.UpdatedAt != nil {
			account.UpdatedAt = timestamppb.New(*resp.Account.UpdatedAt)
		}
	}
	var record *v1.EconomyRecord
	if resp.Record != nil {
		record = &v1.EconomyRecord{
			Id:            resp.Record.ID,
			TransactionNo: resp.Record.TransactionNo,
			UserId:        resp.Record.UserID,
			RecordType:    enum.EconomyRecordTypeMap.MustToProto(resp.Record.RecordType),
			Direction:     enum.EconomyRecordDirectionMap.MustToProto(resp.Record.Direction),
			Amount:        resp.Record.Amount,
			BalanceBefore: resp.Record.BalanceBefore,
			BalanceAfter:  resp.Record.BalanceAfter,
			Remark:        resp.Record.Remark,
		}
		if resp.Record.CreatedAt != nil {
			record.CreatedAt = timestamppb.New(*resp.Record.CreatedAt)
		}
	}
	return &v1.AddEconomyPoints_Resp{Account: account, Record: record}, nil
}

func (s *PointsService) DeductPoints(ctx context.Context, req *v1.DeductEconomyPoints_Req) (*v1.DeductEconomyPoints_Resp, error) {
	recordType, ok := enum.EconomyRecordTypeMap.ToEnum(req.GetRecordType())
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_RECORD_TYPE_INVALID)
	}
	resp, err := s.pointsUsecase.DeductPoints(ctx, &usecase.EconomyDeductPointsReq{
		UserID:     req.GetUserId(),
		Amount:     req.GetAmount(),
		RecordType: recordType,
		Remark:     req.Remark,
	})
	if err != nil {
		return nil, err
	}
	var account *v1.EconomyAccount
	if resp.Account != nil {
		account = &v1.EconomyAccount{
			Id:            resp.Account.ID,
			UserId:        resp.Account.UserID,
			Balance:       resp.Account.Balance,
			FrozenBalance: resp.Account.FrozenBalance,
			TotalIncome:   resp.Account.TotalIncome,
			TotalExpense:  resp.Account.TotalExpense,
		}
		if resp.Account.CreatedAt != nil {
			account.CreatedAt = timestamppb.New(*resp.Account.CreatedAt)
		}
		if resp.Account.UpdatedAt != nil {
			account.UpdatedAt = timestamppb.New(*resp.Account.UpdatedAt)
		}
	}
	var record *v1.EconomyRecord
	if resp.Record != nil {
		record = &v1.EconomyRecord{
			Id:            resp.Record.ID,
			TransactionNo: resp.Record.TransactionNo,
			UserId:        resp.Record.UserID,
			RecordType:    enum.EconomyRecordTypeMap.MustToProto(resp.Record.RecordType),
			Direction:     enum.EconomyRecordDirectionMap.MustToProto(resp.Record.Direction),
			Amount:        resp.Record.Amount,
			BalanceBefore: resp.Record.BalanceBefore,
			BalanceAfter:  resp.Record.BalanceAfter,
			Remark:        resp.Record.Remark,
		}
		if resp.Record.CreatedAt != nil {
			record.CreatedAt = timestamppb.New(*resp.Record.CreatedAt)
		}
	}
	return &v1.DeductEconomyPoints_Resp{Account: account, Record: record}, nil
}

func (s *PointsService) TransferPoints(ctx context.Context, req *v1.TransferEconomyPoints_Req) (*v1.TransferEconomyPoints_Resp, error) {
	outRecordType, ok := enum.EconomyRecordTypeMap.ToEnum(req.GetOutRecordType())
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_RECORD_TYPE_INVALID)
	}
	inRecordType, ok := enum.EconomyRecordTypeMap.ToEnum(req.GetInRecordType())
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_RECORD_TYPE_INVALID)
	}
	resp, err := s.pointsUsecase.TransferPoints(ctx, &usecase.EconomyTransferPointsReq{
		FromUserID:    req.GetFromUserId(),
		ToUserID:      req.GetToUserId(),
		Amount:        req.GetAmount(),
		OutRecordType: outRecordType,
		InRecordType:  inRecordType,
		Remark:        req.Remark,
	})
	if err != nil {
		return nil, err
	}
	var fromAccount *v1.EconomyAccount
	if resp.FromAccount != nil {
		fromAccount = &v1.EconomyAccount{
			Id:            resp.FromAccount.ID,
			UserId:        resp.FromAccount.UserID,
			Balance:       resp.FromAccount.Balance,
			FrozenBalance: resp.FromAccount.FrozenBalance,
			TotalIncome:   resp.FromAccount.TotalIncome,
			TotalExpense:  resp.FromAccount.TotalExpense,
		}
		if resp.FromAccount.CreatedAt != nil {
			fromAccount.CreatedAt = timestamppb.New(*resp.FromAccount.CreatedAt)
		}
		if resp.FromAccount.UpdatedAt != nil {
			fromAccount.UpdatedAt = timestamppb.New(*resp.FromAccount.UpdatedAt)
		}
	}
	var toAccount *v1.EconomyAccount
	if resp.ToAccount != nil {
		toAccount = &v1.EconomyAccount{
			Id:            resp.ToAccount.ID,
			UserId:        resp.ToAccount.UserID,
			Balance:       resp.ToAccount.Balance,
			FrozenBalance: resp.ToAccount.FrozenBalance,
			TotalIncome:   resp.ToAccount.TotalIncome,
			TotalExpense:  resp.ToAccount.TotalExpense,
		}
		if resp.ToAccount.CreatedAt != nil {
			toAccount.CreatedAt = timestamppb.New(*resp.ToAccount.CreatedAt)
		}
		if resp.ToAccount.UpdatedAt != nil {
			toAccount.UpdatedAt = timestamppb.New(*resp.ToAccount.UpdatedAt)
		}
	}
	var outRecord *v1.EconomyRecord
	if resp.OutRecord != nil {
		outRecord = &v1.EconomyRecord{
			Id:            resp.OutRecord.ID,
			TransactionNo: resp.OutRecord.TransactionNo,
			UserId:        resp.OutRecord.UserID,
			RecordType:    enum.EconomyRecordTypeMap.MustToProto(resp.OutRecord.RecordType),
			Direction:     enum.EconomyRecordDirectionMap.MustToProto(resp.OutRecord.Direction),
			Amount:        resp.OutRecord.Amount,
			BalanceBefore: resp.OutRecord.BalanceBefore,
			BalanceAfter:  resp.OutRecord.BalanceAfter,
			Remark:        resp.OutRecord.Remark,
		}
		if resp.OutRecord.CreatedAt != nil {
			outRecord.CreatedAt = timestamppb.New(*resp.OutRecord.CreatedAt)
		}
	}
	var inRecord *v1.EconomyRecord
	if resp.InRecord != nil {
		inRecord = &v1.EconomyRecord{
			Id:            resp.InRecord.ID,
			TransactionNo: resp.InRecord.TransactionNo,
			UserId:        resp.InRecord.UserID,
			RecordType:    enum.EconomyRecordTypeMap.MustToProto(resp.InRecord.RecordType),
			Direction:     enum.EconomyRecordDirectionMap.MustToProto(resp.InRecord.Direction),
			Amount:        resp.InRecord.Amount,
			BalanceBefore: resp.InRecord.BalanceBefore,
			BalanceAfter:  resp.InRecord.BalanceAfter,
			Remark:        resp.InRecord.Remark,
		}
		if resp.InRecord.CreatedAt != nil {
			inRecord.CreatedAt = timestamppb.New(*resp.InRecord.CreatedAt)
		}
	}
	return &v1.TransferEconomyPoints_Resp{FromAccount: fromAccount, ToAccount: toAccount, OutRecord: outRecord, InRecord: inRecord}, nil
}

func (s *PointsService) ListRecords(ctx context.Context, req *v1.ListEconomyRecords_Req) (*v1.ListEconomyRecords_Resp, error) {
	var direction *enum.EconomyRecordDirection
	if req.Direction != nil {
		value, ok := enum.EconomyRecordDirectionMap.ToEnum(req.GetDirection())
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		direction = new(value)
	}
	var recordType *enum.EconomyRecordType
	if req.RecordType != nil {
		value, ok := enum.EconomyRecordTypeMap.ToEnum(req.GetRecordType())
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_RECORD_TYPE_INVALID)
		}
		recordType = new(value)
	}
	page := &base.PageRequest{}
	if req.GetPage() != nil {
		page.Page = int64(req.GetPage().GetPage())
		page.Size = int64(req.GetPage().GetSize())
	}
	resp, err := s.pointsUsecase.ListRecords(ctx, &usecase.EconomyListRecordsReq{UserID: req.GetUserId(), Page: page, Direction: direction, RecordType: recordType})
	if err != nil {
		return nil, err
	}
	rows := make([]*v1.EconomyRecord, 0, len(resp.Rows))
	for _, row := range resp.Rows {
		if row == nil {
			rows = append(rows, nil)
			continue
		}
		item := &v1.EconomyRecord{
			Id:            row.ID,
			TransactionNo: row.TransactionNo,
			UserId:        row.UserID,
			RecordType:    enum.EconomyRecordTypeMap.MustToProto(row.RecordType),
			Direction:     enum.EconomyRecordDirectionMap.MustToProto(row.Direction),
			Amount:        row.Amount,
			BalanceBefore: row.BalanceBefore,
			BalanceAfter:  row.BalanceAfter,
			Remark:        row.Remark,
		}
		if row.CreatedAt != nil {
			item.CreatedAt = timestamppb.New(*row.CreatedAt)
		}
		rows = append(rows, item)
	}
	return &v1.ListEconomyRecords_Resp{Rows: rows, Page: &common.PageResp{Page: uint32(resp.Page.Page), Size: uint32(resp.Page.Size), Total: uint32(resp.Page.Total)}}, nil
}
