package usecase

import (
	"common/pkg/apperror"
	"common/pkg/util/str"
	commonenums "common/proto/gen/common/enums"
	cerrors "common/proto/gen/common/errors"
	"context"
	"strings"
	"time"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/config"
	"user/internal/enum"

	"github.com/sony/sonyflake/v2"
)

type EmailOtpUsecase struct {
	conf          *config.Bootstrap
	accountRepo   repo.AccountRepo
	authCacheRepo repo.AuthCacheRepo
	outboxRepo    repo.OutboxEventRepo
	sf            *sonyflake.Sonyflake
}

func NewEmailOtpUsecase(
	conf *config.Bootstrap,
	accountRepo repo.AccountRepo,
	authCacheRepo repo.AuthCacheRepo,
	outboxRepo repo.OutboxEventRepo,
) (*EmailOtpUsecase, error) {
	sf, err := str.NewSonyflake()
	if err != nil {
		return nil, err
	}
	return &EmailOtpUsecase{
		conf:          conf,
		accountRepo:   accountRepo,
		authCacheRepo: authCacheRepo,
		outboxRepo:    outboxRepo,
		sf:            sf,
	}, nil
}

type SendEmailOtpReq struct {
	UserID *int64
	Email  string
}

type SendEmailOtpResp struct {
	Code string
}

func (u *EmailOtpUsecase) SendEmailOtp(ctx context.Context, req *SendEmailOtpReq) (*SendEmailOtpResp, error) {
	email := strings.ToLower(strings.TrimSpace(req.Email))
	if req.UserID == nil {
		user, err := u.accountRepo.Get(ctx, &repo.AccountGetReq{
			Account: new(email),
		})
		if err != nil {
			code, ok := apperror.BusinessCode(err)
			if !(ok && code == cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_ACCOUNT_NOT_FOUND) {
				return nil, err
			}
		}
		if user == nil || user.Status == nil || *user.Status != enum.AccountStatusNormal {
			return &SendEmailOtpResp{}, nil
		}
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
		Type:    enum.VerificationTypeEmail,
		Account: email,
		UserID:  req.UserID,
	}
	if err := u.authCacheRepo.SaveCode(ctx, &model.VerificationCode{
		Type:        enum.VerificationTypeEmail,
		Account:     email,
		UserID:      req.UserID,
		Code:        code,
		MaxAttempts: maxAttempts,
		CreatedAt:   new(now),
		ExpiresAt:   new(now.Add(codeTTL)),
	}, codeTTL); err != nil {
		return nil, err
	}
	if err := u.outboxRepo.Save(ctx, &repo.OutboxEventSave{
		Event: &commonenums.Event{
			Type:    commonenums.EventType_EVENT_TYPE_USER_EMAIL_VERIFICATION_CODE,
			Subject: commonenums.EventSubject_EVENT_SUBJECT_USER_EMAIL_VERIFICATION_CODE,
			Payload: &commonenums.Event_UserEmailVerificationCode{
				UserEmailVerificationCode: &commonenums.UserEmailVerificationCodePayload{
					Email:          email,
					Code:           code,
					ExpiresSeconds: int64(codeTTL.Seconds()),
				},
			},
		},
	}); err != nil {
		_ = u.authCacheRepo.DeleteCode(ctx, key)
		return nil, err
	}
	return &SendEmailOtpResp{
		Code: code,
	}, nil
}

type VerifyEmailOtpReq struct {
	UserID *int64
	Email  string
	Code   string
}

func (u *EmailOtpUsecase) VerifyEmailOtp(ctx context.Context, req *VerifyEmailOtpReq) error {
	email := strings.ToLower(strings.TrimSpace(req.Email))
	key := &repo.VerificationCodeKeyReq{
		Type:    enum.VerificationTypeEmail,
		Account: email,
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
