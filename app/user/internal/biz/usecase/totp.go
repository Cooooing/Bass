package usecase

import (
	"bytes"
	"common/pkg/apperror"
	commonenums "common/proto/gen/common/enums"
	cerrors "common/proto/gen/common/errors"
	"context"
	"image/png"
	"time"
	base "user/internal/biz/base"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/config"

	"github.com/pquerna/otp/totp"
)

type TotpUsecase struct {
	conf            *config.Bootstrap
	totpSecretCache repo.TotpSecretCache
	tx              base.Tx
	totpRepo        repo.TotpRepo
	outboxRepo      repo.OutboxEventRepo
}

func NewTotpUsecase(
	conf *config.Bootstrap,
	totpSecretCache repo.TotpSecretCache,
	tx base.Tx,
	totpRepo repo.TotpRepo,
	outboxRepo repo.OutboxEventRepo,
) (*TotpUsecase, error) {
	return &TotpUsecase{
		conf:            conf,
		totpSecretCache: totpSecretCache,
		tx:              tx,
		totpRepo:        totpRepo,
		outboxRepo:      outboxRepo,
	}, nil
}

type GetTotpByUserIDReq struct {
	UserID int64
}

type GetTotpByUserIDResponse struct {
	Totp *model.Totp
}

func (u *TotpUsecase) GetByUserID(ctx context.Context, req *GetTotpByUserIDReq) (*GetTotpByUserIDResponse, error) {
	totpResp, err := u.totpRepo.Get(ctx, &repo.TotpGetReq{UserID: &req.UserID})
	if err != nil {
		return nil, err
	}
	return &GetTotpByUserIDResponse{Totp: totpResp.Totp}, nil
}

type ValidateTotpByUserIDReq struct {
	UserID int64
	Code   string
}

type ValidateTotpByUserIDResponse struct {
	Verified bool
}

func (u *TotpUsecase) ValidateByUserID(ctx context.Context, req *ValidateTotpByUserIDReq) (*ValidateTotpByUserIDResponse, error) {
	totpResp, err := u.totpRepo.Get(ctx, &repo.TotpGetReq{UserID: &req.UserID})
	if err != nil {
		return nil, err
	}
	row := totpResp.Totp
	if row == nil || !row.Enable || row.Secret == "" {
		return &ValidateTotpByUserIDResponse{}, nil
	}
	return &ValidateTotpByUserIDResponse{Verified: totp.Validate(req.Code, row.Secret)}, nil
}

type BeginEnableTotpReq struct {
	UserID      int64
	AccountName string
}

type BeginEnableTotpResponse struct {
	URL    string
	QRCode []byte
}

func (u *TotpUsecase) BeginEnable(ctx context.Context, req *BeginEnableTotpReq) (*BeginEnableTotpResponse, error) {
	totpResp, err := u.totpRepo.Get(ctx, &repo.TotpGetReq{UserID: &req.UserID})
	if err != nil {
		return nil, err
	}
	row := totpResp.Totp
	if row != nil && row.Enable {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOTP_ALREADY_ENABLED)
	}

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      u.conf.Business.App,
		AccountName: req.AccountName,
	})
	if err != nil {
		return nil, err
	}
	if _, err = u.totpSecretCache.Save(ctx, &repo.TotpSecretCacheSaveReq{UserID: req.UserID, Secret: key.Secret(), TTL: 5 * time.Minute}); err != nil {
		return nil, err
	}

	image, err := key.Image(256, 256)
	if err != nil {
		return nil, err
	}
	buf := &bytes.Buffer{}
	if err = png.Encode(buf, image); err != nil {
		return nil, err
	}
	return &BeginEnableTotpResponse{URL: key.URL(), QRCode: buf.Bytes()}, nil
}

type CheckEnableTotpCodeReq struct {
	UserID int64
	Code   string
}

type CheckEnableTotpCodeResponse struct {
	Verified bool
}

func (u *TotpUsecase) CheckEnableCode(ctx context.Context, req *CheckEnableTotpCodeReq) (*CheckEnableTotpCodeResponse, error) {
	secretResp, err := u.totpSecretCache.Get(ctx, &repo.TotpSecretCacheGetReq{UserID: req.UserID})
	if err != nil {
		return &CheckEnableTotpCodeResponse{}, nil
	}
	return &CheckEnableTotpCodeResponse{Verified: totp.Validate(req.Code, secretResp.Secret)}, nil
}

type DisableTotpReq struct {
	UserID int64
	Code   string
}

func (u *TotpUsecase) Disable(ctx context.Context, req *DisableTotpReq) error {
	totpResp, err := u.totpRepo.Get(ctx, &repo.TotpGetReq{UserID: &req.UserID})
	if err != nil {
		return err
	}
	row := totpResp.Totp
	if row == nil || !row.Enable {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOTP_ALREADY_DISABLED)
	}
	if !totp.Validate(req.Code, row.Secret) {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOTP_CODE_INVALID)
	}

	return u.tx(ctx, func(ctx context.Context) error {
		if _, err := u.totpRepo.DisableByUserID(ctx, &repo.TotpDisableByUserIDReq{UserID: req.UserID}); err != nil {
			return err
		}
		_, err := u.outboxRepo.Save(ctx, &repo.OutboxEventSave{
			Event: &commonenums.Event{
				Type:    commonenums.EventType_EVENT_TYPE_USER_TOTP_DISABLE,
				Subject: commonenums.EventSubject_EVENT_SUBJECT_USER_TOTP_DISABLE,
				Payload: &commonenums.Event_UserTotpDisable{
					UserTotpDisable: &commonenums.UserTotpDisablePayload{
						UserId: req.UserID,
					},
				},
			},
		})
		return err
	})
}

type ConfirmEnableTotpReq struct {
	UserID int64
	Code   string
}

func (u *TotpUsecase) ConfirmEnable(ctx context.Context, req *ConfirmEnableTotpReq) error {
	secretResp, err := u.totpSecretCache.Get(ctx, &repo.TotpSecretCacheGetReq{UserID: req.UserID})
	if err != nil {
		return err
	}
	if !totp.Validate(req.Code, secretResp.Secret) {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOTP_CODE_INVALID)
	}

	return u.tx(ctx, func(ctx context.Context) error {
		if _, err = u.totpRepo.UpsertEnabledByUserID(ctx, &repo.TotpUpsertEnabledByUserIDReq{UserID: req.UserID, Secret: secretResp.Secret}); err != nil {
			return err
		}
		_, err := u.outboxRepo.Save(ctx, &repo.OutboxEventSave{
			Event: &commonenums.Event{
				Type:    commonenums.EventType_EVENT_TYPE_USER_TOTP_ENABLE,
				Subject: commonenums.EventSubject_EVENT_SUBJECT_USER_TOTP_ENABLE,
				Payload: &commonenums.Event_UserTotpEnable{
					UserTotpEnable: &commonenums.UserTotpEnablePayload{
						UserId: req.UserID,
					},
				},
			},
		})
		return err
	})
}
