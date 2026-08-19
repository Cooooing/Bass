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

type PointsTccService struct {
	v1.UnimplementedPointsTccServiceServer

	pointsUsecase *usecase.PointsUsecase
}

func NewPointsTccService(pointsUsecase *usecase.PointsUsecase) *PointsTccService {
	return &PointsTccService{pointsUsecase: pointsUsecase}
}

func (s *PointsTccService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterPointsTccServiceServer(gs, s)
}

func (s *PointsTccService) RegisterHttp(hs *http.Server) {
}

func (s *PointsTccService) Try(ctx context.Context, req *v1.PointsTccReq) (*v1.PointsTccResp, error) {
	recordType, ok := enum.EconomyRecordTypeMap.ToEnum(req.GetRecordType())
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_RECORD_TYPE_INVALID)
	}
	resp, err := s.pointsUsecase.Try(ctx, &usecase.PointsTccReq{
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
	return &v1.PointsTccResp{Account: account, Record: record}, nil
}

func (s *PointsTccService) Comfirm(ctx context.Context, req *v1.PointsTccReq) (*v1.PointsTccResp, error) {
	recordType, ok := enum.EconomyRecordTypeMap.ToEnum(req.GetRecordType())
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_RECORD_TYPE_INVALID)
	}
	resp, err := s.pointsUsecase.Comfirm(ctx, &usecase.PointsTccReq{
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
	return &v1.PointsTccResp{Account: account, Record: record}, nil
}

func (s *PointsTccService) Cancel(ctx context.Context, req *v1.PointsTccReq) (*v1.PointsTccResp, error) {
	recordType, ok := enum.EconomyRecordTypeMap.ToEnum(req.GetRecordType())
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_RECORD_TYPE_INVALID)
	}
	resp, err := s.pointsUsecase.Cancel(ctx, &usecase.PointsTccReq{
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
	return &v1.PointsTccResp{Account: account, Record: record}, nil
}
