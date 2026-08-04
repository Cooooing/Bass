package usecase

import (
	"common/pkg/apperror"
	"common/pkg/constant"
	commonenum "common/pkg/enum"
	commonModel "common/pkg/model"
	"common/pkg/util/str"
	commonenums "common/proto/gen/common/enums"
	cerrors "common/proto/gen/common/errors"
	"context"
	"strings"
	"time"
	"user/internal/biz/base"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/config"
	"user/internal/enum"

	"log/slog"

	"github.com/google/uuid"
	"github.com/pquerna/otp/totp"
	"github.com/samber/lo"
)

type AuthUsecase struct {
	conf              *config.Bootstrap
	logger            *slog.Logger
	tx                base.Tx
	accountRepo       repo.AccountRepo
	prefsRepo         repo.PreferencesRepo
	loginLogRepo      repo.LoginLogRepo
	outboxRepo        repo.OutboxEventRepo
	outboxUsecase     *OutboxUsecase
	authCacheRepo     repo.AuthCacheRepo
	emailOtpUsecase   *EmailOtpUsecase
	totpRepo          repo.TotpRepo
	banRecordRepo     repo.BanRecordRepo
	ipClient          repo.IPClient
	delayedTaskClient repo.DelayedTaskClient
	tokenUsecase      *TokenUsecase
	rbacUsecase       *RbacUsecase
}

func NewAuthUsecase(
	conf *config.Bootstrap,
	logger *slog.Logger,
	tx base.Tx,
	accountRepo repo.AccountRepo,
	prefsRepo repo.PreferencesRepo,
	loginLogRepo repo.LoginLogRepo,
	outboxRepo repo.OutboxEventRepo,
	outboxUsecase *OutboxUsecase,
	authCacheRepo repo.AuthCacheRepo,
	emailOtpUsecase *EmailOtpUsecase,
	totpRepo repo.TotpRepo,
	banRecordRepo repo.BanRecordRepo,
	ipClient repo.IPClient,
	delayedTaskClient repo.DelayedTaskClient,
	tokenUsecase *TokenUsecase,
	rbacUsecase *RbacUsecase,
) *AuthUsecase {
	return &AuthUsecase{
		conf:              conf,
		logger:            logger,
		tx:                tx,
		accountRepo:       accountRepo,
		prefsRepo:         prefsRepo,
		loginLogRepo:      loginLogRepo,
		outboxRepo:        outboxRepo,
		outboxUsecase:     outboxUsecase,
		authCacheRepo:     authCacheRepo,
		emailOtpUsecase:   emailOtpUsecase,
		totpRepo:          totpRepo,
		banRecordRepo:     banRecordRepo,
		ipClient:          ipClient,
		delayedTaskClient: delayedTaskClient,
		tokenUsecase:      tokenUsecase,
		rbacUsecase:       rbacUsecase,
	}
}

type RegisterReq struct {
	Type     enum.RegisterType
	Name     string
	Password string
	Nickname *string
	Email    string
	Phone    string
	Code     string
}

func (s *AuthUsecase) Register(ctx context.Context, req *RegisterReq) error {
	switch req.Type {
	case enum.RegisterTypeEmail:
		email := strings.ToLower(strings.TrimSpace(req.Email))
		if exists, err := s.accountRepo.ExistsByAccount(ctx, req.Name); err != nil {
			return err
		} else if exists {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_ACCOUNT_NAME_TAKEN)
		}
		if exists, err := s.accountRepo.ExistsByAccount(ctx, email); err != nil {
			return err
		} else if exists {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_ACCOUNT_ALREADY_EXISTS)
		}
		if err := s.emailOtpUsecase.VerifyEmailOtp(ctx, &VerifyEmailOtpReq{
			Email: email,
			Code:  req.Code,
		}); err != nil {
			return err
		}
		passwordHash, err := str.HashPassword(req.Password)
		if err != nil {
			return err
		}
		var outboxEvent *model.OutboxEvent
		err = s.tx(ctx, func(ctx context.Context) error {
			created, err := s.accountRepo.Create(ctx, &model.Account{
				Name:     req.Name,
				Nickname: req.Nickname,
				Password: passwordHash,
				Email:    new(email),
			})
			if err != nil {
				return err
			}
			outboxEvent, err = s.outboxRepo.Save(ctx, &repo.OutboxEventSave{
				Event: &commonenums.Event{
					Type:    commonenums.EventType_EVENT_TYPE_USER_REGISTER,
					Subject: commonenums.EventSubject_EVENT_SUBJECT_USER_REGISTER,
					Payload: &commonenums.Event_UserRegister{
						UserRegister: &commonenums.UserRegisterPayload{
							UserId: created.ID,
						},
					},
				},
			})
			return err
		})
		if err != nil {
			return err
		}
		if outboxEvent != nil {
			if _, err := s.outboxUsecase.Publish(ctx, &PublishOutboxEventReq{
				ID: outboxEvent.ID,
			}); err != nil {
				s.logger.WarnContext(ctx, "publish outbox event best effort failed", constant.LogFieldEventID, outboxEvent.EventID, constant.LogFieldErr, err)
			}
		}
		return nil
	case enum.RegisterTypePhone:
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_NOT_IMPLEMENTED)
	default:
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
}

type LoginReq struct {
	Type            enum.LoginType
	Realm           commonenum.LoginRealm
	Client          *model.LoginContext
	PasswordAccount string
	Password        string
	Code            string
	Email           string
	Phone           string
}

type LoginResp struct {
	TokenPair model.TokenPair
	Account   *model.Account
}

func (s *AuthUsecase) Login(ctx context.Context, req *LoginReq) (*LoginResp, error) {
	client := req.Client
	if client == nil {
		client = &model.LoginContext{}
	}
	if client.ClientType == "" {
		client.ClientType = enum.ClientTypeUnknown
	}
	if client.DeviceType == "" {
		client.DeviceType = enum.DeviceTypeUnknown
	}
	if client.IP != "" && s.ipClient != nil {
		if info, err := s.ipClient.Resolve(ctx, client.IP); err != nil {
			s.logger.WarnContext(ctx, "resolve login ip failed", constant.LogFieldErr, err)
		} else if info != nil {
			client.Country = info.Country
			client.CountryCode = info.CountryCode
			client.Province = info.Province
			client.City = info.City
			client.ISP = info.ISP
		}
	}
	audit := &model.LoginLog{
		LoginType:      req.Type,
		Realm:          req.Realm,
		Status:         enum.LoginStatusFailed,
		ClientType:     new(client.ClientType),
		DeviceType:     new(client.DeviceType),
		OSName:         client.OSName,
		OSVersion:      client.OSVersion,
		BrowserName:    client.BrowserName,
		BrowserVersion: client.BrowserVersion,
		AppName:        client.AppName,
		AppVersion:     client.AppVersion,
	}
	if client.IP != "" {
		audit.IP = new(client.IP)
	}
	if client.Country != "" {
		audit.Country = new(client.Country)
	}
	if client.CountryCode != "" {
		audit.CountryCode = new(client.CountryCode)
	}
	if client.Province != "" {
		audit.Province = new(client.Province)
	}
	if client.City != "" {
		audit.City = new(client.City)
	}
	if client.ISP != "" {
		audit.ISP = new(client.ISP)
	}
	if client.UserAgent != "" {
		audit.UserAgent = new(client.UserAgent)
	}
	defer func() {
		if _, err := s.loginLogRepo.Create(ctx, audit); err != nil {
			s.logger.WarnContext(ctx, "record login log failed", constant.LogFieldErr, err)
		}
	}()

	var account *model.Account
	var err error
	switch req.Type {
	case enum.LoginTypePassword:
		accountInput := strings.TrimSpace(req.PasswordAccount)
		audit.AccountInput = accountInput
		if accountInput == "" || strings.TrimSpace(req.Password) == "" {
			audit.FailureReason = new(enum.LoginFailureReasonInvalidCredentials)
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_INVALID_CREDENTIALS)
		}
		user, err := s.accountRepo.Get(ctx, &repo.AccountGetReq{
			Account: new(accountInput),
		})
		if err != nil || user == nil || user.Status == nil || *user.Status != enum.AccountStatusNormal || !str.VerifyPassword(user.Password, req.Password) {
			if err != nil {
				code, ok := apperror.BusinessCode(err)
				if !(ok && code == cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_ACCOUNT_NOT_FOUND) {
					s.logger.WarnContext(ctx, "password login account lookup failed", constant.LogFieldErr, err)
				}
			}
			audit.FailureReason = new(enum.LoginFailureReasonInvalidCredentials)
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_INVALID_CREDENTIALS)
		}
		audit.UserID = new(user.ID)
		totpRow, err := s.totpRepo.Get(ctx, &repo.TotpGetReq{
			UserID: new(user.ID),
		})
		if err != nil {
			return nil, err
		}
		if totpRow != nil && totpRow.Enable {
			if strings.TrimSpace(req.Code) == "" || !totp.Validate(req.Code, totpRow.Secret) {
				audit.FailureReason = new(enum.LoginFailureReasonTotpInvalid)
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOTP_CODE_INVALID)
			}
		}
		account = user
	case enum.LoginTypeEmail:
		email := strings.ToLower(strings.TrimSpace(req.Email))
		audit.AccountInput = email
		if err := s.emailOtpUsecase.VerifyEmailOtp(ctx, &VerifyEmailOtpReq{
			Email: email,
			Code:  req.Code,
		}); err != nil {
			audit.FailureReason = new(enum.LoginFailureReasonCodeInvalidOrExpired)
			return nil, err
		}
		user, err := s.accountRepo.Get(ctx, &repo.AccountGetReq{
			Account: new(email),
		})
		if err != nil || user == nil || user.Status == nil || *user.Status != enum.AccountStatusNormal {
			if err != nil {
				code, ok := apperror.BusinessCode(err)
				if !(ok && code == cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_ACCOUNT_NOT_FOUND) {
					return nil, err
				}
			}
			audit.FailureReason = new(enum.LoginFailureReasonInvalidCredentials)
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_INVALID_CREDENTIALS)
		}
		audit.UserID = new(user.ID)
		account = user
	case enum.LoginTypePhone:
		audit.AccountInput = req.Phone
		audit.FailureReason = new(enum.LoginFailureReasonNotImplemented)
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_NOT_IMPLEMENTED)
	default:
		audit.FailureReason = new(enum.LoginFailureReasonInvalidCredentials)
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}

	allowed, err := s.rbacUsecase.CheckRealmAccess(ctx, account.ID, req.Realm)
	if err != nil {
		audit.FailureReason = new(enum.LoginFailureReasonInternal)
		return nil, err
	}
	if !allowed {
		audit.FailureReason = new(enum.LoginFailureReasonInvalidCredentials)
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_RBAC_PERMISSION_DENIED)
	}

	now := time.Now()
	sid := uuid.NewString()
	jti := uuid.NewString()
	sessionExpiresAt := now.Add(s.tokenUsecase.SessionTTL())
	prefs, err := s.prefsRepo.Get(ctx, &repo.PreferencesGetReq{
		UserID: new(account.ID),
	})
	if err != nil {
		audit.FailureReason = new(enum.LoginFailureReasonInternal)
		return nil, err
	}
	language := commonenums.Language_LANGUAGE_UNSPECIFIED
	timezone := ""
	if prefs != nil {
		if prefs.Language != nil {
			language = enum.LanguageMap.MustToProto(*prefs.Language)
		}
		if prefs.Timezone != nil {
			timezone = *prefs.Timezone
		}
	}
	nickname := ""
	if account.Nickname != nil {
		nickname = *account.Nickname
	}
	accessToken, accessExpiresAt, err := s.tokenUsecase.GenerateAccess(&GenerateAccessTokenReq{
		UserID:    account.ID,
		SessionID: sid,
		Realm:     req.Realm,
		Name:      account.Name,
		Nickname:  nickname,
		Language:  language,
		Timezone:  timezone,
	})
	if err != nil {
		audit.FailureReason = new(enum.LoginFailureReasonInternal)
		return nil, err
	}
	refreshToken, refreshExpiresAt, err := s.tokenUsecase.GenerateRefresh(&GenerateRefreshTokenReq{
		UserID:    account.ID,
		SessionID: sid,
		Realm:     req.Realm,
		JTI:       jti,
	})
	if err != nil {
		audit.FailureReason = new(enum.LoginFailureReasonInternal)
		return nil, err
	}
	session := &model.RefreshSession{
		SessionID:        sid,
		UserID:           account.ID,
		Realm:            req.Realm,
		CurrentJTI:       jti,
		CreatedAt:        new(now),
		LastSeenAt:       new(now),
		SessionExpiresAt: new(sessionExpiresAt),
		Client:           *client,
	}
	sessionRedisTTL := sessionExpiresAt.Sub(now)
	if sessionRedisTTL <= 0 {
		sessionRedisTTL = time.Second
	}
	if refreshTokenTTL := s.tokenUsecase.RefreshTokenTTL(); refreshTokenTTL < sessionRedisTTL {
		sessionRedisTTL = refreshTokenTTL
	}
	maxSessions := int(s.conf.GetBusiness().GetAuth().GetSession().GetMaxSessions())
	if maxSessions <= 0 {
		maxSessions = 5
	}
	if err := s.authCacheRepo.SaveSession(ctx, session, sessionRedisTTL, maxSessions); err != nil {
		audit.FailureReason = new(enum.LoginFailureReasonInternal)
		return nil, err
	}
	refreshTokenExpiresAt := refreshExpiresAt
	if refreshTokenExpiresAt == nil || sessionExpiresAt.Before(*refreshTokenExpiresAt) {
		refreshTokenExpiresAt = new(sessionExpiresAt)
	}
	tokenPair := &model.TokenPair{
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  accessExpiresAt,
		RefreshTokenExpiresAt: refreshTokenExpiresAt,
		SessionExpiresAt:      new(sessionExpiresAt),
		SessionID:             sid,
	}
	audit.Status = enum.LoginStatusSuccess
	audit.SessionID = tokenPair.SessionID
	outboxEvent, err := s.outboxRepo.Save(ctx, &repo.OutboxEventSave{
		Event: &commonenums.Event{
			Type:    commonenums.EventType_EVENT_TYPE_USER_LOGIN,
			Subject: commonenums.EventSubject_EVENT_SUBJECT_USER_LOGIN,
			Payload: &commonenums.Event_UserLogin{
				UserLogin: &commonenums.UserLoginPayload{
					UserId:     account.ID,
					Name:       account.Name,
					Account:    audit.AccountInput,
					Ip:         client.IP,
					UserAgent:  client.UserAgent,
					Realm:      req.Realm.String(),
					LoginType:  req.Type.String(),
					SessionId:  tokenPair.SessionID,
					AppName:    client.AppName,
					AppVersion: client.AppVersion,
					ClientType: client.ClientType.String(),
					DeviceType: client.DeviceType.String(),
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	if outboxEvent != nil {
		if _, err := s.outboxUsecase.Publish(ctx, &PublishOutboxEventReq{
			ID: outboxEvent.ID,
		}); err != nil {
			s.logger.WarnContext(ctx, "publish outbox event best effort failed", constant.LogFieldEventID, outboxEvent.EventID, constant.LogFieldErr, err)
		}
	}
	return &LoginResp{
		TokenPair: *tokenPair,
		Account:   account,
	}, nil
}
func (s *AuthUsecase) RefreshToken(ctx context.Context, refreshToken string, realm commonenum.LoginRealm) (*model.TokenPair, error) {
	claims, err := s.tokenUsecase.Parse(refreshToken)
	if err != nil || claims.Type != enum.TokenTypeRefresh || claims.Realm != realm || claims.SessionID == "" || claims.JTI == "" {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_INVALID)
	}
	session, err := s.authCacheRepo.GetSession(ctx, realm, claims.SessionID)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	if session == nil || session.UserID != claims.UserID || session.Realm != realm || session.SessionExpiresAt == nil || !session.SessionExpiresAt.After(now) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_INVALID)
	}
	account, err := s.accountRepo.Get(ctx, &repo.AccountGetReq{
		UserID: new(claims.UserID),
	})
	if err != nil {
		code, ok := apperror.BusinessCode(err)
		if ok && code == cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_ACCOUNT_NOT_FOUND {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_INVALID)
		}
		return nil, err
	}
	if account == nil || account.Status == nil || *account.Status != enum.AccountStatusNormal {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_INVALID)
	}
	if session.CurrentJTI != claims.JTI {
		_ = s.authCacheRepo.DeleteSession(ctx, realm, claims.UserID, claims.SessionID)
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_INVALID)
	}
	newJTI := uuid.NewString()
	sessionRedisTTL := session.SessionExpiresAt.Sub(now)
	if sessionRedisTTL <= 0 {
		sessionRedisTTL = time.Second
	}
	if refreshTokenTTL := s.tokenUsecase.RefreshTokenTTL(); refreshTokenTTL < sessionRedisTTL {
		sessionRedisTTL = refreshTokenTTL
	}
	rotated, err := s.authCacheRepo.RotateSessionJTI(ctx, realm, claims.SessionID, claims.JTI, newJTI, new(now), sessionRedisTTL)
	if err != nil {
		return nil, err
	}
	if !rotated {
		_ = s.authCacheRepo.DeleteSession(ctx, realm, claims.UserID, claims.SessionID)
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_INVALID)
	}
	session.CurrentJTI = newJTI
	session.LastSeenAt = new(now)
	if err := s.authCacheRepo.TouchSession(ctx, session, sessionRedisTTL); err != nil {
		return nil, err
	}
	nickname := ""
	if account.Nickname != nil {
		nickname = *account.Nickname
	}
	language := commonenums.Language_LANGUAGE_UNSPECIFIED
	timezone := ""
	prefs, err := s.prefsRepo.Get(ctx, &repo.PreferencesGetReq{
		UserID: new(claims.UserID),
	})
	if err != nil {
		return nil, err
	}
	if prefs != nil {
		if prefs.Language != nil {
			language = enum.LanguageMap.MustToProto(*prefs.Language)
		}
		if prefs.Timezone != nil {
			timezone = *prefs.Timezone
		}
	}
	accessToken, accessExpiresAt, err := s.tokenUsecase.GenerateAccess(&GenerateAccessTokenReq{
		UserID:    claims.UserID,
		SessionID: claims.SessionID,
		Realm:     realm,
		Name:      account.Name,
		Nickname:  nickname,
		Language:  language,
		Timezone:  timezone,
	})
	if err != nil {
		return nil, err
	}
	newRefreshToken, refreshExpiresAt, err := s.tokenUsecase.GenerateRefresh(&GenerateRefreshTokenReq{
		UserID:    claims.UserID,
		SessionID: claims.SessionID,
		Realm:     realm,
		JTI:       newJTI,
	})
	if err != nil {
		return nil, err
	}
	sessionExpiresAt := session.SessionExpiresAt
	refreshTokenExpiresAt := refreshExpiresAt
	if refreshTokenExpiresAt == nil || sessionExpiresAt.Before(*refreshTokenExpiresAt) {
		refreshTokenExpiresAt = sessionExpiresAt
	}
	return &model.TokenPair{
		AccessToken:           accessToken,
		RefreshToken:          newRefreshToken,
		AccessTokenExpiresAt:  accessExpiresAt,
		RefreshTokenExpiresAt: refreshTokenExpiresAt,
		SessionExpiresAt:      sessionExpiresAt,
		SessionID:             claims.SessionID,
	}, nil
}

func (s *AuthUsecase) Logout(ctx context.Context, accessToken string, realm commonenum.LoginRealm) error {
	claims, err := s.tokenUsecase.Parse(accessToken)
	if err != nil || claims.Type != enum.TokenTypeAccess || claims.Realm != realm || claims.SessionID == "" {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_INVALID)
	}
	if err := s.authCacheRepo.DeleteSession(ctx, realm, claims.UserID, claims.SessionID); err != nil {
		return err
	}
	outboxEvent, err := s.outboxRepo.Save(ctx, &repo.OutboxEventSave{
		Event: &commonenums.Event{
			Type:    commonenums.EventType_EVENT_TYPE_USER_LOGOUT,
			Subject: commonenums.EventSubject_EVENT_SUBJECT_USER_LOGOUT,
			Payload: &commonenums.Event_UserLogout{
				UserLogout: &commonenums.UserLogoutPayload{
					UserId:    claims.UserID,
					Name:      claims.Name,
					SessionId: claims.SessionID,
					Realm:     realm.String(),
				},
			},
		},
	})
	if err != nil {
		return err
	}
	if outboxEvent != nil {
		if _, err := s.outboxUsecase.Publish(ctx, &PublishOutboxEventReq{
			ID: outboxEvent.ID,
		}); err != nil {
			s.logger.WarnContext(ctx, "publish outbox event best effort failed", constant.LogFieldEventID, outboxEvent.EventID, constant.LogFieldErr, err)
		}
	}
	return nil
}

type ParseTokenResp struct {
	User      *commonModel.User
	SessionID string
	Realm     commonenum.LoginRealm
}

func (s *AuthUsecase) ParseToken(ctx context.Context, accessToken string, realm commonenum.LoginRealm) (*ParseTokenResp, error) {
	if accessToken == "" {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_INVALID)
	}
	claims, err := s.tokenUsecase.Parse(accessToken)
	if err != nil || claims.Type != enum.TokenTypeAccess || claims.Realm != realm || claims.SessionID == "" {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_INVALID)
	}
	session, err := s.authCacheRepo.GetSession(ctx, realm, claims.SessionID)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	if session == nil || session.UserID != claims.UserID || session.Realm != realm || session.SessionExpiresAt == nil || !session.SessionExpiresAt.After(now) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_INVALID)
	}
	session.LastSeenAt = new(now)
	sessionRedisTTL := session.SessionExpiresAt.Sub(now)
	if sessionRedisTTL <= 0 {
		sessionRedisTTL = time.Second
	}
	if refreshTokenTTL := s.tokenUsecase.RefreshTokenTTL(); refreshTokenTTL < sessionRedisTTL {
		sessionRedisTTL = refreshTokenTTL
	}
	if err := s.authCacheRepo.TouchSession(ctx, session, sessionRedisTTL); err != nil {
		return nil, err
	}
	return &ParseTokenResp{
		User: &commonModel.User{
			ID:       claims.UserID,
			Name:     claims.Name,
			Nickname: claims.Nickname,
			Language: claims.Language,
			Timezone: claims.Timezone,
		},
		SessionID: claims.SessionID,
		Realm:     realm,
	}, nil
}

type CancelAccountReq struct {
	UserID   int64
	Password string
	Code     string
}

func (s *AuthUsecase) CancelAccount(ctx context.Context, req *CancelAccountReq) error {
	user, err := s.accountRepo.Get(ctx, &repo.AccountGetReq{
		UserID: new(req.UserID),
	})
	if err != nil {
		return err
	}
	if user == nil || user.Status == nil || *user.Status != enum.AccountStatusNormal || !str.VerifyPassword(user.Password, req.Password) {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_INVALID_CREDENTIALS)
	}
	totpRow, err := s.totpRepo.Get(ctx, &repo.TotpGetReq{
		UserID: new(req.UserID),
	})
	if err != nil {
		return err
	}
	if totpRow != nil && totpRow.Enable && (strings.TrimSpace(req.Code) == "" || !totp.Validate(req.Code, totpRow.Secret)) {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOTP_CODE_INVALID)
	}
	var outboxEvent *model.OutboxEvent
	err = s.tx(ctx, func(ctx context.Context) error {
		if _, err := s.accountRepo.UpdateStatus(ctx, req.UserID, enum.AccountStatusCancelled); err != nil {
			return err
		}
		if err := s.authCacheRepo.DeleteUserSessions(ctx, req.UserID); err != nil {
			return err
		}
		if err := s.authCacheRepo.DeleteUserRbacPermissions(ctx, req.UserID); err != nil {
			return err
		}
		outboxEvent, err = s.outboxRepo.Save(ctx, &repo.OutboxEventSave{
			Event: &commonenums.Event{
				Type:    commonenums.EventType_EVENT_TYPE_USER_ACCOUNT_CANCELLED,
				Subject: commonenums.EventSubject_EVENT_SUBJECT_USER_ACCOUNT_CANCELLED,
				Payload: &commonenums.Event_UserAccountCancelled{
					UserAccountCancelled: &commonenums.UserAccountCancelledPayload{
						UserId: req.UserID,
					},
				},
			},
		})
		return err
	})
	if err != nil {
		return err
	}
	if outboxEvent != nil {
		if _, err := s.outboxUsecase.Publish(ctx, &PublishOutboxEventReq{
			ID: outboxEvent.ID,
		}); err != nil {
			s.logger.WarnContext(ctx, "publish outbox event best effort failed", constant.LogFieldEventID, outboxEvent.EventID, constant.LogFieldErr, err)
		}
	}
	return nil
}

type BanAccountReq struct {
	UserID        int64
	OperatorID    int64
	OperatorRealm commonenum.LoginRealm
	Reason        string
	Remark        string
	BannedUntil   *time.Time
}

type BanAccountResp struct {
	BanRecordID int64
}

func (s *AuthUsecase) BanAccount(ctx context.Context, req *BanAccountReq) (*BanAccountResp, error) {
	var record *model.BanRecord
	var outboxEvent *model.OutboxEvent
	err := s.tx(ctx, func(ctx context.Context) error {
		var err error
		record, err = s.banRecordRepo.Create(ctx, &model.BanRecord{
			UserID:        req.UserID,
			OperatorID:    req.OperatorID,
			OperatorRealm: req.OperatorRealm,
			Reason:        strings.TrimSpace(req.Reason),
			Remark:        req.Remark,
			StartedAt:     new(time.Now()),
			BannedUntil:   req.BannedUntil,
		})
		if err != nil {
			return err
		}
		if _, err := s.accountRepo.UpdateStatus(ctx, req.UserID, enum.AccountStatusBanned); err != nil {
			return err
		}
		if err := s.authCacheRepo.DeleteUserSessions(ctx, req.UserID); err != nil {
			return err
		}
		if err := s.authCacheRepo.DeleteUserRbacPermissions(ctx, req.UserID); err != nil {
			return err
		}
		payload := &commonenums.UserAccountBannedPayload{
			UserId:        req.UserID,
			OperatorId:    req.OperatorID,
			OperatorRealm: req.OperatorRealm.String(),
			BanRecordId:   record.ID,
			Reason:        strings.TrimSpace(req.Reason),
		}
		if req.BannedUntil != nil {
			payload.BannedUntilUnixSeconds = req.BannedUntil.Unix()
		}
		outboxEvent, err = s.outboxRepo.Save(ctx, &repo.OutboxEventSave{
			Event: &commonenums.Event{
				Type:    commonenums.EventType_EVENT_TYPE_USER_ACCOUNT_BANNED,
				Subject: commonenums.EventSubject_EVENT_SUBJECT_USER_ACCOUNT_BANNED,
				Payload: &commonenums.Event_UserAccountBanned{
					UserAccountBanned: payload,
				},
			},
		})
		return err
	})
	if err != nil {
		return nil, err
	}
	if outboxEvent != nil {
		if _, err := s.outboxUsecase.Publish(ctx, &PublishOutboxEventReq{
			ID: outboxEvent.ID,
		}); err != nil {
			s.logger.WarnContext(ctx, "publish outbox event best effort failed", constant.LogFieldEventID, outboxEvent.EventID, constant.LogFieldErr, err)
		}
	}
	if req.BannedUntil != nil {
		if err := s.delayedTaskClient.RegisterUnbanAccounts(ctx, req.UserID, req.BannedUntil); err != nil {
			return nil, err
		}
	}
	return &BanAccountResp{
		BanRecordID: record.ID,
	}, nil
}

func (s *AuthUsecase) UnbanAccounts(ctx context.Context, userIDs []int64) error {
	return s.accountRepo.UnbanBanned(ctx, lo.Uniq(userIDs))
}
