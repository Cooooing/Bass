package usecase

import (
	"common/pkg/apperror"
	"common/pkg/auth"
	"common/pkg/constant"
	commonModel "common/pkg/model"
	"common/pkg/util/str"
	commonenums "common/proto/gen/common/enums"
	cerrors "common/proto/gen/common/errors"
	"context"
	base "user/internal/biz/base"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/config"
	"user/internal/enum"

	"github.com/sony/sonyflake/v2"
	"log/slog"
)

type AuthUsecase struct {
	conf         *config.Bootstrap
	logger       *slog.Logger
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
	Conf         *config.Bootstrap
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
		logger:       deps.Logger,
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

type CheckPasswordLoginReq struct {
	Account  string
	Password string
}

type CheckPasswordLoginResponse struct {
	Passed bool
}

func (s *AuthUsecase) CheckPasswordLogin(ctx context.Context, req *CheckPasswordLoginReq) (*CheckPasswordLoginResponse, error) {
	if req.Account == "" || req.Password == "" {
		return &CheckPasswordLoginResponse{}, nil
	}
	existResp, err := s.accountRepo.ExistsByAccount(ctx, &repo.AccountExistsByAccountReq{Account: req.Account})
	if err != nil || !existResp.Exists {
		return &CheckPasswordLoginResponse{}, err
	}
	userResp, err := s.accountRepo.Get(ctx, &repo.AccountGetReq{Account: &req.Account})
	if err != nil {
		return nil, err
	}
	return &CheckPasswordLoginResponse{Passed: str.VerifyPassword(userResp.Account.Password, req.Password)}, nil
}

type StartEmailRegistrationReq struct {
	Account *model.Account
}

type StartEmailRegistrationResponse struct {
	Code  string
	Token string
}

func (s *AuthUsecase) StartEmailRegistration(ctx context.Context, req *StartEmailRegistrationReq) (*StartEmailRegistrationResponse, error) {
	u := req.Account
	if u == nil || u.Email == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	token, err := s.tokenUsecase.VerityCodeAccountTokenGen.Generate(model.TokenVerityCodeAccount{Account: *u.Email}, s.conf.Business.Jwt.EmailExpire.AsDuration())
	if err != nil {
		return nil, err
	}
	code := str.RandStr(s.sf, 6, true, true, true, false)
	passwordHash, err := str.HashPassword(u.Password)
	if err != nil {
		return nil, err
	}

	err = s.tokenCache.SaveVerityCode(ctx, constant.VerifyCodeTypeRegisterEmail, *u.Email, code, &registerAccountCache{
		Name:         u.Name,
		Nickname:     u.Nickname,
		PasswordHash: passwordHash,
		Email:        u.Email,
	}, s.conf.Business.Jwt.EmailExpire.AsDuration())
	if err != nil {
		return nil, err
	}
	_, err = s.outboxRepo.Save(ctx, &repo.OutboxEventSave{
		Event: &commonenums.Event{
			Type:    commonenums.EventType_EVENT_TYPE_USER_EMAIL_VERIFICATION_CODE,
			Subject: commonenums.EventSubject_EVENT_SUBJECT_USER_EMAIL_VERIFICATION_CODE,
			Payload: &commonenums.Event_UserEmailVerificationCode{
				UserEmailVerificationCode: &commonenums.UserEmailVerificationCodePayload{
					Email:          *u.Email,
					Code:           code,
					ExpiresSeconds: int64(s.conf.Business.Jwt.EmailExpire.AsDuration().Seconds()),
				},
			},
		},
	})
	if err != nil {
		if delErr := s.tokenCache.DelVerityCode(ctx, constant.VerifyCodeTypeRegisterEmail, *u.Email); delErr != nil {
			s.logger.WarnContext(ctx, "delete email registration verification code failed", constant.LogFieldErr, delErr)
		}
		return nil, err
	}
	return &StartEmailRegistrationResponse{Code: code, Token: token}, nil
}

type CheckEmailRegistrationCodeReq struct {
	CodeToken string
	Code      string
}

type CheckEmailRegistrationCodeResponse struct {
	Passed bool
}

func (s *AuthUsecase) CheckEmailRegistrationCode(ctx context.Context, req *CheckEmailRegistrationCodeReq) (*CheckEmailRegistrationCodeResponse, error) {
	token, err := s.tokenUsecase.VerityCodeAccountTokenGen.Parse(req.CodeToken)
	if err != nil {
		return &CheckEmailRegistrationCodeResponse{}, nil
	}
	saveUser := &registerAccountCache{}
	verityCode, err := s.tokenCache.GetVerityCode(ctx, constant.VerifyCodeTypeRegisterEmail, token.Account, saveUser)
	if err != nil {
		return nil, err
	}
	return &CheckEmailRegistrationCodeResponse{Passed: verityCode == req.Code && saveUser.PasswordHash != ""}, nil
}

type VerifyEmailRegistrationReq struct {
	CodeToken string
	Code      string
}

func (s *AuthUsecase) VerifyEmailRegistration(ctx context.Context, req *VerifyEmailRegistrationReq) error {
	token, err := s.tokenUsecase.VerityCodeAccountTokenGen.Parse(req.CodeToken)
	if err != nil {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_VERIFICATION_CODE_INVALID_OR_EXPIRED)
	}
	saveUser := &registerAccountCache{}
	verityCode, err := s.tokenCache.GetVerityCode(ctx, constant.VerifyCodeTypeRegisterEmail, token.Account, saveUser)
	if err != nil {
		return err
	}
	if verityCode != req.Code || saveUser.PasswordHash == "" {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_VERIFICATION_CODE_INVALID_OR_EXPIRED)
	}

	err = s.tx(ctx, func(ctx context.Context) error {
		user := &model.Account{
			Name:     saveUser.Name,
			Nickname: saveUser.Nickname,
			Password: saveUser.PasswordHash,
			Email:    saveUser.Email,
		}
		createdResp, err := s.accountRepo.Create(ctx, &repo.AccountCreateReq{Account: user})
		if err != nil {
			return err
		}
		err = s.tokenCache.DelVerityCode(ctx, constant.VerifyCodeTypeRegisterEmail, token.Account)
		if err != nil {
			return err
		}
		return s.saveRegisterOutbox(ctx, createdResp.Account.ID)
	})
	if err != nil {
		return err
	}
	return nil
}

type StartPhoneRegistrationReq struct {
	Account *model.Account
}

type StartPhoneRegistrationResponse struct {
	Code  string
	Token string
}

func (s *AuthUsecase) StartPhoneRegistration(ctx context.Context, req *StartPhoneRegistrationReq) (*StartPhoneRegistrationResponse, error) {
	u := req.Account
	if u == nil || u.Phone == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	token, err := s.tokenUsecase.VerityCodeAccountTokenGen.Generate(model.TokenVerityCodeAccount{Account: *u.Phone}, s.conf.Business.Jwt.PhoneExpire.AsDuration())
	if err != nil {
		return nil, err
	}
	code := str.RandStr(s.sf, 6, true, true, true, false)
	passwordHash, err := str.HashPassword(u.Password)
	if err != nil {
		return nil, err
	}

	err = s.tokenCache.SaveVerityCode(ctx, constant.VerifyCodeTypeRegisterPhone, *u.Phone, code, &registerAccountCache{
		Name:         u.Name,
		Nickname:     u.Nickname,
		PasswordHash: passwordHash,
		Phone:        u.Phone,
	}, s.conf.Business.Jwt.PhoneExpire.AsDuration())
	if err != nil {
		return nil, err
	}
	_, err = s.outboxRepo.Save(ctx, &repo.OutboxEventSave{
		Event: &commonenums.Event{
			Type:    commonenums.EventType_EVENT_TYPE_USER_PHONE_VERIFICATION_CODE,
			Subject: commonenums.EventSubject_EVENT_SUBJECT_USER_PHONE_VERIFICATION_CODE,
			Payload: &commonenums.Event_UserPhoneVerificationCode{
				UserPhoneVerificationCode: &commonenums.UserPhoneVerificationCodePayload{
					Phone:          *u.Phone,
					Code:           code,
					ExpiresSeconds: int64(s.conf.Business.Jwt.PhoneExpire.AsDuration().Seconds()),
				},
			},
		},
	})
	if err != nil {
		if delErr := s.tokenCache.DelVerityCode(ctx, constant.VerifyCodeTypeRegisterPhone, *u.Phone); delErr != nil {
			s.logger.WarnContext(ctx, "delete phone registration verification code failed", constant.LogFieldErr, delErr)
		}
		return nil, err
	}
	return &StartPhoneRegistrationResponse{Code: code, Token: token}, nil
}

type CheckPhoneRegistrationCodeReq struct {
	CodeToken string
	Code      string
}

type CheckPhoneRegistrationCodeResponse struct {
	Passed bool
}

func (s *AuthUsecase) CheckPhoneRegistrationCode(ctx context.Context, req *CheckPhoneRegistrationCodeReq) (*CheckPhoneRegistrationCodeResponse, error) {
	token, err := s.tokenUsecase.VerityCodeAccountTokenGen.Parse(req.CodeToken)
	if err != nil {
		return &CheckPhoneRegistrationCodeResponse{}, nil
	}
	saveUser := &registerAccountCache{}
	verityCode, err := s.tokenCache.GetVerityCode(ctx, constant.VerifyCodeTypeRegisterPhone, token.Account, saveUser)
	if err != nil {
		return nil, err
	}
	return &CheckPhoneRegistrationCodeResponse{Passed: verityCode == req.Code && saveUser.PasswordHash != ""}, nil
}

type VerifyPhoneRegistrationReq struct {
	CodeToken string
	Code      string
}

func (s *AuthUsecase) VerifyPhoneRegistration(ctx context.Context, req *VerifyPhoneRegistrationReq) error {
	token, err := s.tokenUsecase.VerityCodeAccountTokenGen.Parse(req.CodeToken)
	if err != nil {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_VERIFICATION_CODE_INVALID_OR_EXPIRED)
	}
	saveUser := &registerAccountCache{}
	verityCode, err := s.tokenCache.GetVerityCode(ctx, constant.VerifyCodeTypeRegisterPhone, token.Account, saveUser)
	if err != nil {
		return err
	}
	if verityCode != req.Code || saveUser.PasswordHash == "" {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_VERIFICATION_CODE_INVALID_OR_EXPIRED)
	}

	err = s.tx(ctx, func(ctx context.Context) error {
		user := &model.Account{
			Name:     saveUser.Name,
			Nickname: saveUser.Nickname,
			Password: saveUser.PasswordHash,
			Phone:    saveUser.Phone,
		}
		createdResp, err := s.accountRepo.Create(ctx, &repo.AccountCreateReq{Account: user})
		if err != nil {
			return err
		}
		err = s.tokenCache.DelVerityCode(ctx, constant.VerifyCodeTypeRegisterPhone, token.Account)
		if err != nil {
			return err
		}
		return s.saveRegisterOutbox(ctx, createdResp.Account.ID)
	})
	if err != nil {
		return err
	}
	return nil
}

type LoginByPasswordReq struct {
	Account      string
	Password     string
	LoginContext *model.LoginContext
}

type LoginByPasswordResponse struct {
	Token   string
	Account *model.Account
}

func (s *AuthUsecase) LoginByPassword(ctx context.Context, req *LoginByPasswordReq) (*LoginByPasswordResponse, error) {
	var loginUserID *int64
	loginStatus := enum.LoginStatusFailed
	defer func() {
		loginLog := &model.LoginLog{
			UserID:      loginUserID,
			LoginMethod: enum.LoginMethodPassword,
			Status:      loginStatus,
		}
		if req.LoginContext != nil {
			if req.LoginContext.IP != "" {
				loginLog.IP = new(req.LoginContext.IP)
			}
			if req.LoginContext.Country != "" {
				loginLog.Country = new(req.LoginContext.Country)
			}
			if req.LoginContext.CountryCode != "" {
				loginLog.CountryCode = new(req.LoginContext.CountryCode)
			}
			if req.LoginContext.Province != "" {
				loginLog.Province = new(req.LoginContext.Province)
			}
			if req.LoginContext.City != "" {
				loginLog.City = new(req.LoginContext.City)
			}
			if req.LoginContext.ISP != "" {
				loginLog.ISP = new(req.LoginContext.ISP)
			}
			if req.LoginContext.UserAgent != "" {
				loginLog.UserAgent = new(req.LoginContext.UserAgent)
			}
			if req.LoginContext.DeviceID != "" {
				loginLog.DeviceID = new(req.LoginContext.DeviceID)
			}
		}
		if _, logErr := s.loginLogRepo.Create(ctx, &repo.LoginLogCreateReq{Log: loginLog}); logErr != nil {
			s.logger.WarnContext(ctx, "record login log failed", constant.LogFieldErr, logErr)
		}
	}()

	userResp, err := s.accountRepo.Get(ctx, &repo.AccountGetReq{Account: &req.Account})
	if err != nil {
		return nil, err
	}
	user := userResp.Account
	loginUserID = new(user.ID)
	if !str.VerifyPassword(userResp.Account.Password, req.Password) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_INVALID_CREDENTIALS)
	}

	token, err := s.tokenUsecase.TokenGen.Generate(model.Token{Id: user.ID}, s.conf.Business.Jwt.Expires.AsDuration())
	if err != nil {
		return nil, err
	}
	saveUser := &commonModel.User{
		ID:   user.ID,
		Name: user.Name,
	}
	if user.Nickname != nil {
		saveUser.Nickname = *user.Nickname
	}
	prefsResp, err := s.prefsRepo.Get(ctx, &repo.PreferencesGetReq{UserID: &user.ID})
	if err != nil {
		return nil, err
	}
	if prefsResp.Preferences != nil {
		prefs := prefsResp.Preferences
		if prefs.Language != nil {
			saveUser.Language = enum.LanguageMap.MustToProto(*prefs.Language)
		}
		if prefs.Timezone != nil {
			saveUser.Timezone = *prefs.Timezone
		}
	}
	err = s.tokenCache.SaveToken(ctx, token, saveUser, s.conf.Business.Jwt.Expires.AsDuration())
	if err != nil {
		return nil, err
	}
	loginStatus = enum.LoginStatusSuccess
	loginPayload := &commonenums.UserLoginPayload{
		UserId:  user.ID,
		Name:    user.Name,
		Account: req.Account,
	}
	if req.LoginContext != nil {
		loginPayload.UserAgent = req.LoginContext.UserAgent
		loginPayload.DeviceId = req.LoginContext.DeviceID
		loginPayload.Platform = req.LoginContext.Platform
		loginPayload.RequestId = req.LoginContext.RequestID
		loginPayload.Ip = req.LoginContext.IP
	}
	if _, outboxErr := s.outboxRepo.Save(ctx, &repo.OutboxEventSave{
		Event: &commonenums.Event{
			Type:    commonenums.EventType_EVENT_TYPE_USER_LOGIN,
			Subject: commonenums.EventSubject_EVENT_SUBJECT_USER_LOGIN,
			Payload: &commonenums.Event_UserLogin{
				UserLogin: loginPayload,
			},
		},
	}); outboxErr != nil {
		s.logger.ErrorContext(ctx, "create login outbox failed", constant.LogFieldEventType, commonenums.EventType_EVENT_TYPE_USER_LOGIN.String(), constant.LogFieldErr, outboxErr)
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INTERNAL).WithCause(outboxErr)
	}

	return &LoginByPasswordResponse{Token: token, Account: user}, nil
}

type LogoutReq struct {
	Token string
}

func (s *AuthUsecase) Logout(ctx context.Context, req *LogoutReq) error {
	tokenUser, getErr := s.tokenCache.GetToken(ctx, req.Token)
	err := s.tokenCache.DelToken(ctx, req.Token)
	if err != nil {
		return err
	}
	if getErr != nil {
		s.logger.WarnContext(ctx, "get logout token user failed", constant.LogFieldErr, getErr)
		return nil
	}
	if tokenUser != nil {
		if _, outboxErr := s.outboxRepo.Save(ctx, &repo.OutboxEventSave{
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
			s.logger.ErrorContext(ctx, "create logout outbox failed", constant.LogFieldEventType, commonenums.EventType_EVENT_TYPE_USER_LOGOUT.String(), constant.LogFieldErr, outboxErr)
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INTERNAL).WithCause(outboxErr)
		}
	}
	return nil
}

type ParseTokenReq struct {
	Token string
}

type ParseTokenResponse struct {
	User *commonModel.User
}

func (s *AuthUsecase) ParseToken(ctx context.Context, req *ParseTokenReq) (*ParseTokenResponse, error) {
	if req.Token == "" {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_INVALID)
	}
	tokenData, err := s.tokenUsecase.TokenGen.Parse(req.Token)
	if err != nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_INVALID).WithCause(err)
	}
	user, err := s.tokenCache.GetToken(ctx, req.Token)
	if err != nil {
		return nil, err
	}
	if user == nil || user.ID != tokenData.Id {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_INVALID)
	}
	return &ParseTokenResponse{User: user}, nil
}

func (s *AuthUsecase) saveRegisterOutbox(ctx context.Context, userID int64) error {
	_, err := s.outboxRepo.Save(ctx, &repo.OutboxEventSave{
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
	return err
}
