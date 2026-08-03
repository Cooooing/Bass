package usecase

import (
	"common/pkg/apperror"
	"common/pkg/constant"
	"common/pkg/util/str"
	commonenums "common/proto/gen/common/enums"
	cerrors "common/proto/gen/common/errors"
	"context"
	"log/slog"
	"strings"
	"time"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/config"
	"user/internal/enum"

	"github.com/sony/sonyflake/v2"
)

type SmsOtpUsecase struct {
	logger                      *slog.Logger
	conf                        *config.Bootstrap
	authCacheRepo               repo.AuthCacheRepo
	outboxRepo                  repo.OutboxEventRepo
	outboxUsecase               *OutboxUsecase
	notificationRateLimitClient repo.NotificationRateLimitClient
	sf                          *sonyflake.Sonyflake
}

func NewSmsOtpUsecase(
	logger *slog.Logger,
	conf *config.Bootstrap,
	authCacheRepo repo.AuthCacheRepo,
	outboxRepo repo.OutboxEventRepo,
	outboxUsecase *OutboxUsecase,
	notificationRateLimitClient repo.NotificationRateLimitClient,
) (*SmsOtpUsecase, error) {
	sf, err := str.NewSonyflake()
	if err != nil {
		return nil, err
	}
	return &SmsOtpUsecase{
		logger:                      logger,
		conf:                        conf,
		authCacheRepo:               authCacheRepo,
		outboxRepo:                  outboxRepo,
		outboxUsecase:               outboxUsecase,
		notificationRateLimitClient: notificationRateLimitClient,
		sf:                          sf,
	}, nil
}

type SendPhoneOtpReq struct {
	UserID *int64
	Phone  string
}

type SendPhoneOtpResp struct {
	Code string
}

func (u *SmsOtpUsecase) SendPhoneOtp(ctx context.Context, req *SendPhoneOtpReq) (*SendPhoneOtpResp, error) {
	phone := strings.TrimSpace(req.Phone)

	rateLimitState, err := u.notificationRateLimitClient.CheckPhone(ctx, phone)
	if err != nil {
		return nil, err
	}
	if rateLimitState != nil && rateLimitState.Limited {
		return nil, apperror.New(
			cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_VERIFICATION_CODE_SEND_TOO_FREQUENT,
			apperror.WithData(&cerrors.RetryAfterErrorData{
				RetryAfterSeconds: rateLimitState.RetryAfterSeconds,
			}),
		)
	}

	verificationCodeConf := u.conf.GetBusiness().GetAuth().GetVerificationCode()
	codeTTL := 5 * time.Minute
	if verificationCodeConf.GetCodeTtl() != nil && verificationCodeConf.GetCodeTtl().AsDuration() > 0 {
		codeTTL = verificationCodeConf.GetCodeTtl().AsDuration()
	}
	maxAttempts := verificationCodeConf.GetMaxAttempts()
	if maxAttempts <= 0 {
		maxAttempts = 5
	}
	now := time.Now()
	code := str.RandStr(u.sf, 6, true, true, true, false)
	key := &repo.VerificationCodeKeyReq{
		Type:    enum.VerificationTypePhone,
		Account: phone,
		UserID:  req.UserID,
	}
	if err := u.authCacheRepo.SaveCode(ctx, &model.VerificationCode{
		Type:        enum.VerificationTypePhone,
		Account:     phone,
		UserID:      req.UserID,
		Code:        code,
		MaxAttempts: maxAttempts,
		CreatedAt:   new(now),
		ExpiresAt:   new(now.Add(codeTTL)),
	}, codeTTL); err != nil {
		return nil, err
	}
	outboxEvent, err := u.outboxRepo.Save(ctx, &repo.OutboxEventSave{
		Event: &commonenums.Event{
			Type:    commonenums.EventType_EVENT_TYPE_USER_PHONE_VERIFICATION_CODE,
			Subject: commonenums.EventSubject_EVENT_SUBJECT_USER_PHONE_VERIFICATION_CODE,
			Payload: &commonenums.Event_UserPhoneVerificationCode{
				UserPhoneVerificationCode: &commonenums.UserPhoneVerificationCodePayload{
					Phone:          phone,
					Code:           code,
					ExpiresSeconds: int64(codeTTL.Seconds()),
				},
			},
		},
	})
	if err != nil {
		_ = u.authCacheRepo.DeleteCode(ctx, key)
		return nil, err
	}
	if outboxEvent != nil {
		if _, err := u.outboxUsecase.Publish(ctx, &PublishOutboxEventReq{
			ID: outboxEvent.ID,
		}); err != nil {
			u.logger.WarnContext(ctx, "publish outbox event best effort failed", constant.LogFieldEventID, outboxEvent.EventID, constant.LogFieldErr, err)
		}
	}
	return &SendPhoneOtpResp{
		Code: code,
	}, nil
}

type VerifyPhoneOtpReq struct {
	UserID *int64
	Phone  string
	Code   string
}

func (u *SmsOtpUsecase) VerifyPhoneOtp(ctx context.Context, req *VerifyPhoneOtpReq) error {
	phone := strings.TrimSpace(req.Phone)
	key := &repo.VerificationCodeKeyReq{
		Type:    enum.VerificationTypePhone,
		Account: phone,
		UserID:  req.UserID,
	}
	row, err := u.authCacheRepo.GetCode(ctx, key)
	if err != nil {
		return err
	}
	if row == nil || row.ExpiresAt == nil || !row.ExpiresAt.After(time.Now()) || row.Attempts >= row.MaxAttempts || row.Code != strings.TrimSpace(req.Code) {
		if row != nil && row.Attempts < row.MaxAttempts {
			_, _ = u.authCacheRepo.IncrCodeAttempts(ctx, key)
		}
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_VERIFICATION_CODE_INVALID_OR_EXPIRED)
	}
	return u.authCacheRepo.DeleteCode(ctx, key)
}
