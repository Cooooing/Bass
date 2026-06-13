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
	"user/internal/conf"

	"github.com/pquerna/otp/totp"
)

type TotpUsecase struct {
	conf            *conf.Bootstrap
	totpSecretCache repo.TotpSecretCache
	tx              base.Tx
	totpRepo        repo.TotpRepo
	outboxRepo      repo.OutboxEventRepo
}

func NewTotpUsecase(
	conf *conf.Bootstrap,
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

func (u *TotpUsecase) GetByUserID(ctx context.Context, userID int64) (*model.Totp, error) {
	return u.totpRepo.Get(ctx, &repo.TotpGetReq{UserID: &userID})
}

func (u *TotpUsecase) ValidateByUserID(ctx context.Context, userID int64, code string) (bool, error) {
	row, err := u.totpRepo.Get(ctx, &repo.TotpGetReq{UserID: &userID})
	if err != nil {
		return false, err
	}
	if row == nil || !row.Enable || row.Secret == "" {
		return false, nil
	}
	return totp.Validate(code, row.Secret), nil
}

func (u *TotpUsecase) BeginEnable(ctx context.Context, userID int64, accountName string) (string, []byte, error) {
	row, err := u.totpRepo.Get(ctx, &repo.TotpGetReq{UserID: &userID})
	if err != nil {
		return "", nil, err
	}
	if row != nil && row.Enable {
		return "", nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOTP_ALREADY_ENABLED)
	}

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      u.conf.Server.App,
		AccountName: accountName,
	})
	if err != nil {
		return "", nil, err
	}
	if err = u.totpSecretCache.Save(ctx, userID, key.Secret(), 5*time.Minute); err != nil {
		return "", nil, err
	}

	image, err := key.Image(256, 256)
	if err != nil {
		return "", nil, err
	}
	buf := &bytes.Buffer{}
	if err = png.Encode(buf, image); err != nil {
		return "", nil, err
	}
	return key.URL(), buf.Bytes(), nil
}

func (u *TotpUsecase) CheckEnableCode(ctx context.Context, userID int64, code string) (bool, error) {
	secret, err := u.totpSecretCache.Get(ctx, userID)
	if err != nil {
		return false, nil
	}
	return totp.Validate(code, secret), nil
}

func (u *TotpUsecase) Disable(ctx context.Context, userID int64, code string) error {
	row, err := u.totpRepo.Get(ctx, &repo.TotpGetReq{UserID: &userID})
	if err != nil {
		return err
	}
	if row == nil || !row.Enable {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOTP_ALREADY_DISABLED)
	}
	if !totp.Validate(code, row.Secret) {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOTP_CODE_INVALID)
	}

	return u.tx(ctx, func(ctx context.Context) error {
		if _, err := u.totpRepo.DisableByUserID(ctx, userID); err != nil {
			return err
		}
		return u.outboxRepo.Save(ctx, &repo.OutboxEventSave{
			Event: &commonenums.Event{
				Type:    commonenums.EventType_EVENT_TYPE_USER_TOTP_DISABLE,
				Subject: commonenums.EventSubject_EVENT_SUBJECT_USER_TOTP_DISABLE,
				Payload: &commonenums.Event_UserTotpDisable{
					UserTotpDisable: &commonenums.UserTotpDisablePayload{
						UserId: userID,
					},
				},
			},
		})
	})
}

func (u *TotpUsecase) ConfirmEnable(ctx context.Context, userID int64, code string) error {
	secret, err := u.totpSecretCache.Get(ctx, userID)
	if err != nil {
		return err
	}
	if !totp.Validate(code, secret) {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOTP_CODE_INVALID)
	}

	return u.tx(ctx, func(ctx context.Context) error {
		if _, err = u.totpRepo.UpsertEnabledByUserID(ctx, userID, secret); err != nil {
			return err
		}
		return u.outboxRepo.Save(ctx, &repo.OutboxEventSave{
			Event: &commonenums.Event{
				Type:    commonenums.EventType_EVENT_TYPE_USER_TOTP_ENABLE,
				Subject: commonenums.EventSubject_EVENT_SUBJECT_USER_TOTP_ENABLE,
				Payload: &commonenums.Event_UserTotpEnable{
					UserTotpEnable: &commonenums.UserTotpEnablePayload{
						UserId: userID,
					},
				},
			},
		})
	})
}
