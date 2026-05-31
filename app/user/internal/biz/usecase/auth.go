package usecase

import (
	commonenums "common/api/gen/common/enums"
	cerrors "common/api/gen/common/errors"
	"common/pkg/constant"
	commonModel "common/pkg/model"
	"common/pkg/util"
	"common/pkg/util/jwt"
	serverutil "common/pkg/util/server"
	"common/pkg/util/str"
	"context"
	base "user/internal/biz/base"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/conf"
	"user/internal/enum"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/sony/sonyflake/v2"
)

type AuthUsecase struct {
	conf            *conf.Bootstrap
	log             *log.Helper
	tx              base.Tx
	accountRepo     repo.AccountRepo
	prefsRepo       repo.PreferencesRepo
	loginLogRepo    repo.LoginLogRepo
	outboxRepo      repo.OutboxEventRepo
	rateLimitClient repo.NotificationRateLimitClient
	tokenCache      *jwt.TokenCache
	tokenUsecase    *TokenUsecase

	sf *sonyflake.Sonyflake
}

type AuthUsecaseDeps struct {
	Conf            *conf.Bootstrap
	Logger          log.Logger
	Tx              base.Tx
	AccountRepo     repo.AccountRepo
	PrefsRepo       repo.PreferencesRepo
	LoginLogRepo    repo.LoginLogRepo
	OutboxRepo      repo.OutboxEventRepo
	RateLimitClient repo.NotificationRateLimitClient
	TokenCache      *jwt.TokenCache
	TokenUsecase    *TokenUsecase
}

type registerAccountCache struct {
	Name         string  `json:"name"`
	Nickname     *string `json:"nickname,omitempty"`
	PasswordHash string  `json:"password_hash"`
	Email        *string `json:"email,omitempty"`
	Phone        *string `json:"phone,omitempty"`
	Blocked      bool    `json:"blocked,omitempty"`
}

const verificationCodeInvalidOrExpired = "verification code is invalid or expired"

func NewAuthUsecase(deps AuthUsecaseDeps) (*AuthUsecase, error) {
	sf, err := str.NewSonyflake()
	if err != nil {
		return nil, err
	}
	return &AuthUsecase{
		conf:            deps.Conf,
		log:             log.NewHelper(deps.Logger),
		tx:              deps.Tx,
		accountRepo:     deps.AccountRepo,
		prefsRepo:       deps.PrefsRepo,
		loginLogRepo:    deps.LoginLogRepo,
		outboxRepo:      deps.OutboxRepo,
		rateLimitClient: deps.RateLimitClient,
		tokenCache:      deps.TokenCache,
		tokenUsecase:    deps.TokenUsecase,
		sf:              sf,
	}, nil
}

func (s *AuthUsecase) StartEmailRegistration(ctx context.Context, u *model.Account) (code string, token string, err error) {
	if u == nil || u.Email == nil {
		return "", "", cerrors.ErrorBadRequest("email can not be empty")
	}
	exist, err := s.accountRepo.ExistsByAccount(ctx, u.Name)
	if exist {
		err = cerrors.ErrorBadRequest("name is already taken")
	}
	if err != nil {
		return
	}

	token, err = s.tokenUsecase.VerityCodeAccountTokenGen.Generate(model.TokenVerityCodeAccount{Account: *u.Email}, s.conf.Server.Jwt.EmailExpire.AsDuration())
	if err != nil {
		return
	}
	exist, err = s.accountRepo.ExistsByAccount(ctx, *u.Email)
	if err != nil {
		return
	}
	rateLimitState, err := s.rateLimitClient.CheckEmail(ctx, *u.Email)
	if err != nil {
		return
	}
	if rateLimitState != nil && rateLimitState.Limited {
		return "", "", cerrors.ErrorTooManyRequests("verification code send too frequent, retry after %d seconds", int64(rateLimitState.RetryAfter.Seconds()))
	}
	code = str.RandStr(s.sf, 6, true, true, true, false)
	if exist {
		err = s.tokenCache.SaveVerityCode(ctx, constant.VerifyCodeTypeRegisterEmail, *u.Email, "", &registerAccountCache{
			Blocked: true,
			Email:   u.Email,
		}, s.conf.Server.Jwt.EmailExpire.AsDuration())
		return code, token, err
	}
	passwordHash, err := str.HashPassword(u.Password)
	if err != nil {
		return
	}

	err = s.tokenCache.SaveVerityCode(ctx, constant.VerifyCodeTypeRegisterEmail, *u.Email, code, &registerAccountCache{
		Name:         u.Name,
		Nickname:     u.Nickname,
		PasswordHash: passwordHash,
		Email:        u.Email,
	}, s.conf.Server.Jwt.EmailExpire.AsDuration())
	if err != nil {
		return
	}
	err = s.outboxRepo.Save(ctx, &repo.OutboxEventSave{
		Event: &commonenums.Event{
			Type:    commonenums.EventType_EVENT_TYPE_USER_EMAIL_VERIFICATION_CODE,
			Subject: commonenums.EventSubject_EVENT_SUBJECT_USER_EMAIL_VERIFICATION_CODE,
			Payload: &commonenums.Event_UserEmailVerificationCode{
				UserEmailVerificationCode: &commonenums.UserEmailVerificationCodePayload{
					Email:          *u.Email,
					Code:           code,
					ExpiresSeconds: int64(s.conf.Server.Jwt.EmailExpire.AsDuration().Seconds()),
				},
			},
		},
	})
	if err != nil {
		if delErr := s.tokenCache.DelVerityCode(ctx, constant.VerifyCodeTypeRegisterEmail, *u.Email); delErr != nil {
			s.log.Warnf("delete email registration verification code failed: %v", delErr)
		}
		return
	}
	return code, token, nil
}

func (s *AuthUsecase) VerifyEmailRegistration(ctx context.Context, codeToken string, code string) (err error) {
	token, err := s.tokenUsecase.VerityCodeAccountTokenGen.Parse(codeToken)
	if err != nil {
		return cerrors.ErrorBadRequest(verificationCodeInvalidOrExpired)
	}
	saveUser := &registerAccountCache{}
	verityCode, err := s.tokenCache.GetVerityCode(ctx, constant.VerifyCodeTypeRegisterEmail, token.Account, saveUser)
	if err != nil {
		return cerrors.ErrorBadRequest(verificationCodeInvalidOrExpired)
	}
	if saveUser.Blocked || verityCode != code || saveUser.PasswordHash == "" {
		err = cerrors.ErrorBadRequest(verificationCodeInvalidOrExpired)
		return
	}

	err = s.tx(ctx, func(ctx context.Context) error {
		user := &model.Account{
			Name:     saveUser.Name,
			Nickname: saveUser.Nickname,
			Password: saveUser.PasswordHash,
			Email:    saveUser.Email,
		}
		created, err := s.accountRepo.Create(ctx, user)
		if err != nil {
			return err
		}
		err = s.tokenCache.DelVerityCode(ctx, constant.VerifyCodeTypeRegisterEmail, token.Account)
		if err != nil {
			return err
		}
		return s.saveRegisterOutbox(ctx, created.ID)
	})
	if err != nil {
		return
	}
	return
}

func (s *AuthUsecase) StartPhoneRegistration(ctx context.Context, u *model.Account) (code string, token string, err error) {
	if u == nil || u.Phone == nil {
		return "", "", cerrors.ErrorBadRequest("phone can not be empty")
	}
	exist, err := s.accountRepo.ExistsByAccount(ctx, u.Name)
	if exist {
		err = cerrors.ErrorBadRequest("name is already taken")
	}
	if err != nil {
		return
	}

	token, err = s.tokenUsecase.VerityCodeAccountTokenGen.Generate(model.TokenVerityCodeAccount{Account: *u.Phone}, s.conf.Server.Jwt.PhoneExpire.AsDuration())
	if err != nil {
		return
	}
	exist, err = s.accountRepo.ExistsByAccount(ctx, *u.Phone)
	if err != nil {
		return
	}
	rateLimitState, err := s.rateLimitClient.CheckPhone(ctx, *u.Phone)
	if err != nil {
		return
	}
	if rateLimitState != nil && rateLimitState.Limited {
		return "", "", cerrors.ErrorTooManyRequests("verification code send too frequent, retry after %d seconds", int64(rateLimitState.RetryAfter.Seconds()))
	}
	if exist {
		code = str.RandStr(s.sf, 6, true, true, true, false)
		err = s.tokenCache.SaveVerityCode(ctx, constant.VerifyCodeTypeRegisterPhone, *u.Phone, "", &registerAccountCache{
			Blocked: true,
			Phone:   u.Phone,
		}, s.conf.Server.Jwt.PhoneExpire.AsDuration())
		return code, token, err
	}

	code = str.RandStr(s.sf, 6, true, true, true, false)
	passwordHash, err := str.HashPassword(u.Password)
	if err != nil {
		return
	}

	err = s.tokenCache.SaveVerityCode(ctx, constant.VerifyCodeTypeRegisterPhone, *u.Phone, code, &registerAccountCache{
		Name:         u.Name,
		Nickname:     u.Nickname,
		PasswordHash: passwordHash,
		Phone:        u.Phone,
	}, s.conf.Server.Jwt.PhoneExpire.AsDuration())
	if err != nil {
		return
	}
	err = s.outboxRepo.Save(ctx, &repo.OutboxEventSave{
		Event: &commonenums.Event{
			Type:    commonenums.EventType_EVENT_TYPE_USER_PHONE_VERIFICATION_CODE,
			Subject: commonenums.EventSubject_EVENT_SUBJECT_USER_PHONE_VERIFICATION_CODE,
			Payload: &commonenums.Event_UserPhoneVerificationCode{
				UserPhoneVerificationCode: &commonenums.UserPhoneVerificationCodePayload{
					Phone:          *u.Phone,
					Code:           code,
					ExpiresSeconds: int64(s.conf.Server.Jwt.PhoneExpire.AsDuration().Seconds()),
				},
			},
		},
	})
	if err != nil {
		if delErr := s.tokenCache.DelVerityCode(ctx, constant.VerifyCodeTypeRegisterPhone, *u.Phone); delErr != nil {
			s.log.Warnf("delete phone registration verification code failed: %v", delErr)
		}
		return
	}
	return code, token, nil
}

func (s *AuthUsecase) VerifyPhoneRegistration(ctx context.Context, codeToken string, code string) (err error) {
	token, err := s.tokenUsecase.VerityCodeAccountTokenGen.Parse(codeToken)
	if err != nil {
		return cerrors.ErrorBadRequest(verificationCodeInvalidOrExpired)
	}
	saveUser := &registerAccountCache{}
	verityCode, err := s.tokenCache.GetVerityCode(ctx, constant.VerifyCodeTypeRegisterPhone, token.Account, saveUser)
	if err != nil {
		return cerrors.ErrorBadRequest(verificationCodeInvalidOrExpired)
	}
	if saveUser.Blocked || verityCode != code || saveUser.PasswordHash == "" {
		err = cerrors.ErrorBadRequest(verificationCodeInvalidOrExpired)
		return
	}

	err = s.tx(ctx, func(ctx context.Context) error {
		user := &model.Account{
			Name:     saveUser.Name,
			Nickname: saveUser.Nickname,
			Password: saveUser.PasswordHash,
			Phone:    saveUser.Phone,
		}
		created, err := s.accountRepo.Create(ctx, user)
		if err != nil {
			return err
		}
		err = s.tokenCache.DelVerityCode(ctx, constant.VerifyCodeTypeRegisterPhone, token.Account)
		if err != nil {
			return err
		}
		return s.saveRegisterOutbox(ctx, created.ID)
	})
	if err != nil {
		return
	}
	return
}

func (s *AuthUsecase) LoginByPassword(ctx context.Context, account string, password string) (token string, user *model.Account, err error) {
	var loginUserID *int64
	loginStatus := enum.LoginStatusFailed
	defer func() {
		loginLog := &model.LoginLog{
			UserID:      loginUserID,
			LoginMethod: enum.LoginMethodPassword,
			Status:      loginStatus,
		}
		if ipInfo, ok := util.GetContextValue[*commonModel.IpInfo](ctx, constant.CtxIpInfo); ok && ipInfo != nil {
			if ipInfo.Ip != "" {
				loginLog.IP = new(ipInfo.Ip)
			}
			if ipInfo.Country != "" {
				loginLog.Country = new(ipInfo.Country)
			}
			if ipInfo.CountryCode != "" {
				loginLog.CountryCode = new(ipInfo.CountryCode)
			}
			if ipInfo.Province != "" {
				loginLog.Province = new(ipInfo.Province)
			}
			if ipInfo.City != "" {
				loginLog.City = new(ipInfo.City)
			}
			if ipInfo.ISP != "" {
				loginLog.ISP = new(ipInfo.ISP)
			}
		}
		if loginLog.IP == nil {
			ip := serverutil.ClientIP(ctx)
			if ip != "" {
				loginLog.IP = new(ip)
			}
		}
		if userAgent := serverutil.GetHeader(ctx, constant.HeaderUserAgent); userAgent != "" {
			loginLog.UserAgent = new(userAgent)
		}
		if deviceID := serverutil.GetHeader(ctx, constant.HeaderDeviceID); deviceID != "" {
			loginLog.DeviceID = new(deviceID)
		}
		if _, logErr := s.loginLogRepo.Create(ctx, loginLog); logErr != nil {
			s.log.Warnf("record login log failed: %v", logErr)
		}
	}()

	user, err = s.accountRepo.GetByAccount(ctx, account)
	if err != nil {
		return
	}
	loginUserID = new(user.ID)
	if !str.VerifyPassword(user.Password, password) {
		return token, nil, cerrors.ErrorBadRequest("password invalid")
	}

	token, err = s.tokenUsecase.TokenGen.Generate(model.Token{Id: user.ID}, s.conf.Server.Jwt.Expires.AsDuration())
	if err != nil {
		return
	}
	saveUser := &commonModel.User{
		ID:   user.ID,
		Name: user.Name,
	}
	if user.Nickname != nil {
		saveUser.Nickname = *user.Nickname
	}
	prefs, err := s.prefsRepo.FindByUserID(ctx, user.ID)
	if err != nil {
		return
	}
	if prefs != nil {
		if prefs.Language != nil {
			saveUser.Language = string(*prefs.Language)
		}
		if prefs.Timezone != nil {
			saveUser.Timezone = *prefs.Timezone
		}
	}
	err = s.tokenCache.SaveToken(ctx, token, saveUser, s.conf.Server.Jwt.Expires.AsDuration())
	if err != nil {
		return
	}
	loginStatus = enum.LoginStatusSuccess
	loginPayload := &commonenums.UserLoginPayload{
		UserId:    user.ID,
		Name:      user.Name,
		Account:   account,
		UserAgent: serverutil.GetHeader(ctx, constant.HeaderUserAgent),
		DeviceId:  serverutil.GetHeader(ctx, constant.HeaderDeviceID),
		Platform:  serverutil.GetHeader(ctx, constant.HeaderPlatform),
		RequestId: serverutil.GetHeader(ctx, constant.HeaderRequestID),
	}
	if loginPayload.RequestId == "" {
		loginPayload.RequestId = serverutil.GetHeader(ctx, constant.HeaderTraceID)
	}
	loginPayload.Ip = serverutil.ClientIP(ctx)
	if outboxErr := s.outboxRepo.Save(ctx, &repo.OutboxEventSave{
		Event: &commonenums.Event{
			Type:    commonenums.EventType_EVENT_TYPE_USER_LOGIN,
			Subject: commonenums.EventSubject_EVENT_SUBJECT_USER_LOGIN,
			Payload: &commonenums.Event_UserLogin{
				UserLogin: loginPayload,
			},
		},
	}); outboxErr != nil {
		s.log.Errorf("create login outbox failed: %v", outboxErr)
		return "", nil, cerrors.ErrorInternalServerError("create login outbox failed").WithCause(outboxErr)
	}

	return token, user, nil
}

func (s *AuthUsecase) Logout(ctx context.Context, token string) (err error) {
	tokenUser, getErr := s.tokenCache.GetToken(ctx, token)
	err = s.tokenCache.DelToken(ctx, token)
	if err != nil {
		return err
	}
	if getErr != nil {
		s.log.Warnf("get logout token user failed: %v", getErr)
		return nil
	}
	if tokenUser != nil {
		if outboxErr := s.outboxRepo.Save(ctx, &repo.OutboxEventSave{
			Event: &commonenums.Event{
				Type:    commonenums.EventType_EVENT_TYPE_USER_LOGOUT,
				Subject: commonenums.EventSubject_EVENT_SUBJECT_USER_LOGOUT,
				Payload: &commonenums.Event_UserLogout{
					UserLogout: &commonenums.UserLogoutPayload{
						UserId: tokenUser.ID,
						Name:   tokenUser.Name,
					},
				},
			},
		}); outboxErr != nil {
			s.log.Errorf("create logout outbox failed: %v", outboxErr)
			return cerrors.ErrorInternalServerError("create logout outbox failed").WithCause(outboxErr)
		}
	}
	return nil
}

func (s *AuthUsecase) ParseToken(ctx context.Context, token string) (*commonModel.User, error) {
	if token == "" {
		return nil, cerrors.ErrorUnauthorized("token is invalid")
	}
	tokenData, err := s.tokenUsecase.TokenGen.Parse(token)
	if err != nil {
		return nil, cerrors.ErrorUnauthorized("token is invalid").WithCause(err)
	}
	user, err := s.tokenCache.GetToken(ctx, token)
	if err != nil {
		return nil, err
	}
	if user == nil || user.ID != tokenData.Id {
		return nil, cerrors.ErrorUnauthorized("token is invalid")
	}
	return user, nil
}

func (s *AuthUsecase) saveRegisterOutbox(ctx context.Context, userID int64) error {
	return s.outboxRepo.Save(ctx, &repo.OutboxEventSave{
		Event: &commonenums.Event{
			Type:    commonenums.EventType_EVENT_TYPE_USER_REGISTER,
			Subject: commonenums.EventSubject_EVENT_SUBJECT_USER_REGISTER,
			Payload: &commonenums.Event_UserRegister{
				UserRegister: &commonenums.UserRegisterPayload{
					UserId: userID,
				},
			},
		},
	})
}
