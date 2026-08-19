package service

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	v1 "common/proto/gen/economy/v1"
	"context"
	"economy/internal/biz/usecase"
	"economy/internal/enum"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PointsTransferTccService struct {
	v1.UnimplementedPointsTransferTccServiceServer

	pointsUsecase *usecase.PointsUsecase
}

func NewPointsTransferTccService(pointsUsecase *usecase.PointsUsecase) *PointsTransferTccService {
	return &PointsTransferTccService{pointsUsecase: pointsUsecase}
}

func (s *PointsTransferTccService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterPointsTransferTccServiceServer(gs, s)
}

func (s *PointsTransferTccService) RegisterHttp(hs *http.Server) {
}

func (s *PointsTransferTccService) Try(ctx context.Context, req *v1.PointsTransferTccReq) (*v1.PointsTransferTccResp, error) {
	outRecordType, ok := enum.EconomyRecordTypeMap.ToEnum(req.GetOutRecordType())
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_RECORD_TYPE_INVALID)
	}
	inRecordType, ok := enum.EconomyRecordTypeMap.ToEnum(req.GetInRecordType())
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_RECORD_TYPE_INVALID)
	}
	resp, err := s.pointsUsecase.TransferTry(ctx, &usecase.PointsTransferTccReq{
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
	return &v1.PointsTransferTccResp{FromAccount: fromAccount, ToAccount: toAccount, OutRecord: outRecord, InRecord: inRecord}, nil
}

func (s *PointsTransferTccService) Comfirm(ctx context.Context, req *v1.PointsTransferTccReq) (*v1.PointsTransferTccResp, error) {
	outRecordType, ok := enum.EconomyRecordTypeMap.ToEnum(req.GetOutRecordType())
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_RECORD_TYPE_INVALID)
	}
	inRecordType, ok := enum.EconomyRecordTypeMap.ToEnum(req.GetInRecordType())
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_RECORD_TYPE_INVALID)
	}
	resp, err := s.pointsUsecase.TransferComfirm(ctx, &usecase.PointsTransferTccReq{
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
	return &v1.PointsTransferTccResp{FromAccount: fromAccount, ToAccount: toAccount, OutRecord: outRecord, InRecord: inRecord}, nil
}

func (s *PointsTransferTccService) Cancel(ctx context.Context, req *v1.PointsTransferTccReq) (*v1.PointsTransferTccResp, error) {
	outRecordType, ok := enum.EconomyRecordTypeMap.ToEnum(req.GetOutRecordType())
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_RECORD_TYPE_INVALID)
	}
	inRecordType, ok := enum.EconomyRecordTypeMap.ToEnum(req.GetInRecordType())
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_RECORD_TYPE_INVALID)
	}
	resp, err := s.pointsUsecase.TransferCancel(ctx, &usecase.PointsTransferTccReq{
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
	return &v1.PointsTransferTccResp{FromAccount: fromAccount, ToAccount: toAccount, OutRecord: outRecord, InRecord: inRecord}, nil
}
