package usecase

import (
	"common/pkg/apperror"
	"common/pkg/auth"
	"common/pkg/constant"
	commonModel "common/pkg/model"
	"common/pkg/util"
	"common/pkg/util/str"
	commonenums "common/proto/gen/common/enums"
	cerrors "common/proto/gen/common/errors"
	"context"
	base "user/internal/biz/base"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/conf"
	"user/internal/enum"

	"github.com/sony/sonyflake/v2"
	"log/slog"
)

type AuthUsecase struct {
	conf         *conf.Bootstrap
	log          *util.LogHelper
	tx           base.Tx
	accountRepo  repo.AccountRepo
	prefsRepo    repo.PreferencesRepo
	loginLogRepo repo.LoginLogRepo
	outboxRepo   repo.OutboxEventRepo
	tokenCache   *auth.TokenCache
	tokenUsecase *TokenUsecase

	sf *sonyflake.Sonyflake
}

type AuthUsecaseDeps struct {
	Conf         *conf.Bootstrap
	Logger       *slog.Logger
	Tx           base.Tx
	AccountRepo  repo.AccountRepo
	PrefsRepo    repo.PreferencesRepo
	LoginLogRepo repo.LoginLogRepo
	OutboxRepo   repo.OutboxEventRepo
	TokenCache   *auth.TokenCache
	TokenUsecase *TokenUsecase
}

type registerAccountCache struct {
	Name         string  `json:"name"`
	Nickname     *string `json:"nickname,omitempty"`
	PasswordHash string  `json:"password_hash"`
	Email        *string `json:"email,omitempty"`
	Phone        *string `json:"phone,omitempty"`
}

func NewAuthUsecase(deps AuthUsecaseDeps) (*AuthUsecase, error) {
	sf, err := str.NewSonyflake()
	if err != nil {
		return nil, err
	}
	return &AuthUsecase{
		conf:         deps.Conf,
		log:          util.NewLogHelper(deps.Logger),
		tx:           deps.Tx,
		accountRepo:  deps.AccountRepo,
		prefsRepo:    deps.PrefsRepo,
		loginLogRepo: deps.LoginLogRepo,
		outboxRepo:   deps.OutboxRepo,
		tokenCache:   deps.TokenCache,
		tokenUsecase: deps.TokenUsecase,
		sf:           sf,
	}, nil
}

func (s *AuthUsecase) CheckPasswordLogin(ctx context.Context, account string, password string) (bool, error) {
	if account == "" || password == "" {
		return false, nil
	}
	exist, err := s.accountRepo.ExistsByAccount(ctx, account)
	if err != nil || !exist {
		return false, err
	}
	user, err := s.accountRepo.Get(ctx, &repo.AccountGetReq{Account: &account})
	if err != nil {
		return false, err
	}
	return str.VerifyPassword(user.Password, password), nil
}

func (s *AuthUsecase) StartEmailRegistration(ctx context.Context, u *model.Account) (code string, token string, err error) {
	if u == nil || u.Email == nil {
		return "", "", apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	token, err = s.tokenUsecase.VerityCodeAccountTokenGen.Generate(model.TokenVerityCodeAccount{Account: *u.Email}, s.conf.Server.Jwt.EmailExpire.AsDuration())
	if err != nil {
		return
	}
	code = str.RandStr(s.sf, 6, true, true, true, false)
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

func (s *AuthUsecase) CheckEmailRegistrationCode(ctx context.Context, codeToken string, code string) (bool, error) {
	token, err := s.tokenUsecase.VerityCodeAccountTokenGen.Parse(codeToken)
	if err != nil {
		return false, nil
	}
	saveUser := &registerAccountCache{}
	verityCode, err := s.tokenCache.GetVerityCode(ctx, constant.VerifyCodeTypeRegisterEmail, token.Account, saveUser)
	if err != nil {
		return false, err
	}
	return verityCode == code && saveUser.PasswordHash != "", nil
}

func (s *AuthUsecase) VerifyEmailRegistration(ctx context.Context, codeToken string, code string) (err error) {
	token, err := s.tokenUsecase.VerityCodeAccountTokenGen.Parse(codeToken)
	if err != nil {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_VERIFICATION_CODE_INVALID_OR_EXPIRED)
	}
	saveUser := &registerAccountCache{}
	verityCode, err := s.tokenCache.GetVerityCode(ctx, constant.VerifyCodeTypeRegisterEmail, token.Account, saveUser)
	if err != nil {
		return err
	}
	if verityCode != code || saveUser.PasswordHash == "" {
		err = apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_VERIFICATION_CODE_INVALID_OR_EXPIRED)
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
		return "", "", apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	token, err = s.tokenUsecase.VerityCodeAccountTokenGen.Generate(model.TokenVerityCodeAccount{Account: *u.Phone}, s.conf.Server.Jwt.PhoneExpire.AsDuration())
	if err != nil {
		return
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

func (s *AuthUsecase) CheckPhoneRegistrationCode(ctx context.Context, codeToken string, code string) (bool, error) {
	token, err := s.tokenUsecase.VerityCodeAccountTokenGen.Parse(codeToken)
	if err != nil {
		return false, nil
	}
	saveUser := &registerAccountCache{}
	verityCode, err := s.tokenCache.GetVerityCode(ctx, constant.VerifyCodeTypeRegisterPhone, token.Account, saveUser)
	if err != nil {
		return false, err
	}
	return verityCode == code && saveUser.PasswordHash != "", nil
}

func (s *AuthUsecase) VerifyPhoneRegistration(ctx context.Context, codeToken string, code string) (err error) {
	token, err := s.tokenUsecase.VerityCodeAccountTokenGen.Parse(codeToken)
	if err != nil {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_VERIFICATION_CODE_INVALID_OR_EXPIRED)
	}
	saveUser := &registerAccountCache{}
	verityCode, err := s.tokenCache.GetVerityCode(ctx, constant.VerifyCodeTypeRegisterPhone, token.Account, saveUser)
	if err != nil {
		return err
	}
	if verityCode != code || saveUser.PasswordHash == "" {
		err = apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_VERIFICATION_CODE_INVALID_OR_EXPIRED)
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

func (s *AuthUsecase) LoginByPassword(ctx context.Context, account string, password string, loginContext *model.LoginContext) (token string, user *model.Account, err error) {
	var loginUserID *int64
	loginStatus := enum.LoginStatusFailed
	defer func() {
		loginLog := &model.LoginLog{
			UserID:      loginUserID,
			LoginMethod: enum.LoginMethodPassword,
			Status:      loginStatus,
		}
		if loginContext != nil {
			if loginContext.IP != "" {
				loginLog.IP = new(loginContext.IP)
			}
			if loginContext.Country != "" {
				loginLog.Country = new(loginContext.Country)
			}
			if loginContext.CountryCode != "" {
				loginLog.CountryCode = new(loginContext.CountryCode)
			}
			if loginContext.Province != "" {
				loginLog.Province = new(loginContext.Province)
			}
			if loginContext.City != "" {
				loginLog.City = new(loginContext.City)
			}
			if loginContext.ISP != "" {
				loginLog.ISP = new(loginContext.ISP)
			}
			if loginContext.UserAgent != "" {
				loginLog.UserAgent = new(loginContext.UserAgent)
			}
			if loginContext.DeviceID != "" {
				loginLog.DeviceID = new(loginContext.DeviceID)
			}
		}
		if _, logErr := s.loginLogRepo.Create(ctx, loginLog); logErr != nil {
			s.log.Warnf("record login log failed: %v", logErr)
		}
	}()

	user, err = s.accountRepo.Get(ctx, &repo.AccountGetReq{Account: &account})
	if err != nil {
		return
	}
	loginUserID = new(user.ID)
	if !str.VerifyPassword(user.Password, password) {
		return token, nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_INVALID_CREDENTIALS)
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
	prefs, err := s.prefsRepo.Get(ctx, &repo.PreferencesGetReq{UserID: &user.ID})
	if err != nil {
		return
	}
	if prefs != nil {
		if prefs.Language != nil {
			saveUser.Language = enum.LanguageMap.MustToProto(*prefs.Language)
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
		UserId:  user.ID,
		Name:    user.Name,
		Account: account,
	}
	if loginContext != nil {
		loginPayload.UserAgent = loginContext.UserAgent
		loginPayload.DeviceId = loginContext.DeviceID
		loginPayload.Platform = loginContext.Platform
		loginPayload.RequestId = loginContext.RequestID
		loginPayload.Ip = loginContext.IP
	}
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
		return "", nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INTERNAL).WithCause(outboxErr)
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
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INTERNAL).WithCause(outboxErr)
		}
	}
	return nil
}

func (s *AuthUsecase) ParseToken(ctx context.Context, token string) (*commonModel.User, error) {
	if token == "" {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_INVALID)
	}
	tokenData, err := s.tokenUsecase.TokenGen.Parse(token)
	if err != nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_INVALID).WithCause(err)
	}
	user, err := s.tokenCache.GetToken(ctx, token)
	if err != nil {
		return nil, err
	}
	if user == nil || user.ID != tokenData.Id {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_INVALID)
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
