package service

import (
	"common/pkg/apperror"
	"common/proto/gen/common"
	cerrors "common/proto/gen/common/errors"
	v1 "common/proto/gen/economy/v1"
	"context"
	"economy/internal/biz/base"
	"economy/internal/biz/model"
	"economy/internal/biz/usecase"
	"economy/internal/enum"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type EconomyService struct {
	v1.UnimplementedEconomyServiceServer

	economyUsecase *usecase.EconomyUsecase
}

func NewEconomyService(economyUsecase *usecase.EconomyUsecase) *EconomyService {
	return &EconomyService{economyUsecase: economyUsecase}
}

func (s *EconomyService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterEconomyServiceServer(gs, s)
}

func (s *EconomyService) RegisterHttp(hs *http.Server) {
}

func (s *EconomyService) GetAccount(ctx context.Context, req *v1.GetEconomyAccount_Req) (*v1.GetEconomyAccount_Resp, error) {
	account, err := s.economyUsecase.GetAccount(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}
	return &v1.GetEconomyAccount_Resp{Account: s.account(account)}, nil
}

func (s *EconomyService) MapAccounts(ctx context.Context, req *v1.MapEconomyAccounts_Req) (*v1.MapEconomyAccounts_Resp, error) {
	accounts, err := s.economyUsecase.MapAccounts(ctx, req.GetUserIds())
	if err != nil {
		return nil, err
	}
	reply := &v1.MapEconomyAccounts_Resp{Accounts: make(map[int64]*v1.EconomyAccount, len(accounts))}
	for userID, account := range accounts {
		reply.Accounts[userID] = s.account(account)
	}
	return reply, nil
}

func (s *EconomyService) AddPoints(ctx context.Context, req *v1.AddEconomyPoints_Req) (*v1.AddEconomyPoints_Resp, error) {
	recordType, ok := enum.EconomyRecordTypeMap.ToEnum(req.GetRecordType())
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_RECORD_TYPE_INVALID)
	}
	resp, err := s.economyUsecase.AddPoints(ctx, &usecase.EconomyAddPointsReq{UserID: req.GetUserId(), Amount: req.GetAmount(), RecordType: recordType, SourceID: req.SourceId, IdempotencyKey: req.GetIdempotencyKey(), Remark: req.Remark})
	if err != nil {
		return nil, err
	}
	return &v1.AddEconomyPoints_Resp{Account: s.account(resp.Account), Record: s.record(resp.Record)}, nil
}

func (s *EconomyService) DeductPoints(ctx context.Context, req *v1.DeductEconomyPoints_Req) (*v1.DeductEconomyPoints_Resp, error) {
	recordType, ok := enum.EconomyRecordTypeMap.ToEnum(req.GetRecordType())
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_RECORD_TYPE_INVALID)
	}
	resp, err := s.economyUsecase.DeductPoints(ctx, &usecase.EconomyDeductPointsReq{UserID: req.GetUserId(), Amount: req.GetAmount(), RecordType: recordType, SourceID: req.SourceId, IdempotencyKey: req.GetIdempotencyKey(), Remark: req.Remark})
	if err != nil {
		return nil, err
	}
	return &v1.DeductEconomyPoints_Resp{Account: s.account(resp.Account), Record: s.record(resp.Record)}, nil
}

func (s *EconomyService) TransferPoints(ctx context.Context, req *v1.TransferEconomyPoints_Req) (*v1.TransferEconomyPoints_Resp, error) {
	outRecordType, ok := enum.EconomyRecordTypeMap.ToEnum(req.GetOutRecordType())
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_RECORD_TYPE_INVALID)
	}
	inRecordType, ok := enum.EconomyRecordTypeMap.ToEnum(req.GetInRecordType())
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_RECORD_TYPE_INVALID)
	}
	resp, err := s.economyUsecase.TransferPoints(ctx, &usecase.EconomyTransferPointsReq{FromUserID: req.GetFromUserId(), ToUserID: req.GetToUserId(), Amount: req.GetAmount(), OutRecordType: outRecordType, InRecordType: inRecordType, SourceID: req.SourceId, IdempotencyKey: req.GetIdempotencyKey(), Remark: req.Remark})
	if err != nil {
		return nil, err
	}
	return &v1.TransferEconomyPoints_Resp{FromAccount: s.account(resp.FromAccount), ToAccount: s.account(resp.ToAccount), OutRecord: s.record(resp.OutRecord), InRecord: s.record(resp.InRecord)}, nil
}

func (s *EconomyService) ListRecords(ctx context.Context, req *v1.ListEconomyRecords_Req) (*v1.ListEconomyRecords_Resp, error) {
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
	resp, err := s.economyUsecase.ListRecords(ctx, &usecase.EconomyListRecordsReq{UserID: req.GetUserId(), Page: page, Direction: direction, RecordType: recordType})
	if err != nil {
		return nil, err
	}
	rows := make([]*v1.EconomyRecord, 0, len(resp.Rows))
	for _, row := range resp.Rows {
		rows = append(rows, s.record(row))
	}
	return &v1.ListEconomyRecords_Resp{Rows: rows, Page: &common.PageResp{Page: uint32(resp.Page.Page), Size: uint32(resp.Page.Size), Total: uint32(resp.Page.Total)}}, nil
}

func (s *EconomyService) account(account *model.Account) *v1.EconomyAccount {
	if account == nil {
		return nil
	}
	out := &v1.EconomyAccount{Id: account.ID, UserId: account.UserID, Balance: account.Balance, TotalIncome: account.TotalIncome, TotalExpense: account.TotalExpense}
	if account.CreatedAt != nil {
		out.CreatedAt = timestamppb.New(*account.CreatedAt)
	}
	if account.UpdatedAt != nil {
		out.UpdatedAt = timestamppb.New(*account.UpdatedAt)
	}
	return out
}

func (s *EconomyService) record(record *model.Record) *v1.EconomyRecord {
	if record == nil {
		return nil
	}
	out := &v1.EconomyRecord{Id: record.ID, TransactionNo: record.TransactionNo, UserId: record.UserID, RecordType: enum.EconomyRecordTypeMap.MustToProto(record.RecordType), Direction: enum.EconomyRecordDirectionMap.MustToProto(record.Direction), Amount: record.Amount, BalanceBefore: record.BalanceBefore, BalanceAfter: record.BalanceAfter, SourceId: record.SourceID, IdempotencyKey: record.IdempotencyKey, Remark: record.Remark}
	if record.CreatedAt != nil {
		out.CreatedAt = timestamppb.New(*record.CreatedAt)
	}
	return out
}
