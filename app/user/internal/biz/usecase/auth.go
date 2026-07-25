package usecase

import (
	commonenum "common/pkg/enum"
	"common/pkg/apperror"
	"common/pkg/constant"
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
	"github.com/sony/sonyflake/v2"
)

const (
	verificationTypeEmail = "email"
	verificationTypePhone = "phone"
)

type AuthUsecase struct {
	conf              *config.Bootstrap
	logger            *slog.Logger
	tx                base.Tx
	accountRepo       repo.AccountRepo
	prefsRepo         repo.PreferencesRepo
	loginLogRepo      repo.LoginLogRepo
	outboxRepo        repo.OutboxEventRepo
	authCacheRepo     repo.AuthCacheRepo
	totpRepo          repo.TotpRepo
	banRecordRepo     repo.BanRecordRepo
	ipClient          repo.IPClient
	delayedTaskClient repo.DelayedTaskClient
	tokenUsecase      *TokenUsecase

	sf *sonyflake.Sonyflake
}

type AuthUsecaseDeps struct {
	Conf              *config.Bootstrap
	Logger            *slog.Logger
	Tx                base.Tx
	AccountRepo       repo.AccountRepo
	PrefsRepo         repo.PreferencesRepo
	LoginLogRepo      repo.LoginLogRepo
	OutboxRepo        repo.OutboxEventRepo
	AuthCacheRepo     repo.AuthCacheRepo
	TotpRepo          repo.TotpRepo
	BanRecordRepo     repo.BanRecordRepo
	IPClient          repo.IPClient
	DelayedTaskClient repo.DelayedTaskClient
	TokenUsecase      *TokenUsecase
}

func NewAuthUsecase(deps AuthUsecaseDeps) (*AuthUsecase, error) {
	sf, err := str.NewSonyflake()
	if err != nil {
		return nil, err
	}
	return &AuthUsecase{
		conf:              deps.Conf,
		logger:            deps.Logger,
		tx:                deps.Tx,
		accountRepo:       deps.AccountRepo,
		prefsRepo:         deps.PrefsRepo,
		loginLogRepo:      deps.LoginLogRepo,
		outboxRepo:        deps.OutboxRepo,
		authCacheRepo:     deps.AuthCacheRepo,
		totpRepo:          deps.TotpRepo,
		banRecordRepo:     deps.BanRecordRepo,
		ipClient:          deps.IPClient,
		delayedTaskClient: deps.DelayedTaskClient,
		tokenUsecase:      deps.TokenUsecase,
		sf:                sf,
	}, nil
}

type CheckPasswordLoginReq struct {
	Account  string
	Password string
}

func (s *AuthUsecase) CheckPasswordLogin(ctx context.Context, req *CheckPasswordLoginReq) (bool, error) {
	if req.Account == "" || req.Password == "" {
		return false, nil
	}
	user, err := s.accountRepo.Get(ctx, &repo.AccountGetReq{Account: &req.Account})
	if isAccountNotFound(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return user != nil && str.VerifyPassword(user.Password, req.Password), nil
}

type StartEmailRegistrationResp struct {
	Code string
}

func (s *AuthUsecase) StartEmailRegistration(ctx context.Context, account *model.Account) (*StartEmailRegistrationResp, error) {
	if account == nil || account.Email == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	email := normalizeAccountKey(*account.Email)
	if exists, err := s.accountRepo.ExistsByAccount(ctx, account.Name); err != nil {
		return nil, err
	} else if exists {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_ACCOUNT_NAME_TAKEN)
	}
	if exists, err := s.accountRepo.ExistsByAccount(ctx, email); err != nil {
		return nil, err
	} else if exists {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_ACCOUNT_ALREADY_EXISTS)
	}
	passwordHash, err := str.HashPassword(account.Password)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	codeTTL := s.verificationCodeTTL()
	draftTTL := s.registerDraftTTL()
	code := str.RandStr(s.sf, 6, true, true, true, false)
	if err := s.authCacheRepo.SaveRegisterDraft(ctx, verificationTypeEmail, email, &model.RegisterDraft{
		Name:          account.Name,
		Nickname:      account.Nickname,
		PasswordHash:  passwordHash,
		Email:         &email,
		CreatedAtUnix: now.Unix(),
		ExpiresAtUnix: now.Add(draftTTL).Unix(),
	}, draftTTL); err != nil {
		return nil, err
	}
	if err := s.authCacheRepo.SaveCode(ctx, &model.VerificationCode{
		Type:          verificationTypeEmail,
		Account:       email,
		Code:          code,
		MaxAttempts:   s.verificationMaxAttempts(),
		CreatedAtUnix: now.Unix(),
		ExpiresAtUnix: now.Add(codeTTL).Unix(),
	}, codeTTL); err != nil {
		return nil, err
	}
	if err := s.saveVerificationCodeOutbox(ctx, verificationTypeEmail, email, code, codeTTL); err != nil {
		_ = s.authCacheRepo.DeleteCode(ctx, verificationTypeEmail, email)
		_ = s.authCacheRepo.DeleteRegisterDraft(ctx, verificationTypeEmail, email)
		return nil, err
	}
	return &StartEmailRegistrationResp{Code: code}, nil
}

type CheckEmailRegistrationCodeReq struct {
	Email string
	Code  string
}

func (s *AuthUsecase) CheckEmailRegistrationCode(ctx context.Context, req *CheckEmailRegistrationCodeReq) (bool, error) {
	return s.checkVerificationCode(ctx, verificationTypeEmail, normalizeAccountKey(req.Email), req.Code, false)
}

type VerifyEmailRegistrationReq struct {
	Email string
	Code  string
}

func (s *AuthUsecase) VerifyEmailRegistration(ctx context.Context, req *VerifyEmailRegistrationReq) error {
	email := normalizeAccountKey(req.Email)
	ok, err := s.checkVerificationCode(ctx, verificationTypeEmail, email, req.Code, true)
	if err != nil {
		return err
	}
	if !ok {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_VERIFICATION_CODE_INVALID_OR_EXPIRED)
	}
	draft, err := s.authCacheRepo.GetRegisterDraft(ctx, verificationTypeEmail, email)
	if err != nil {
		return err
	}
	if draft == nil || draft.PasswordHash == "" || draft.Email == nil {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_VERIFICATION_CODE_INVALID_OR_EXPIRED)
	}
	return s.tx(ctx, func(ctx context.Context) error {
		created, err := s.accountRepo.Create(ctx, &model.Account{
			Name:     draft.Name,
			Nickname: draft.Nickname,
			Password: draft.PasswordHash,
			Email:    draft.Email,
		})
		if err != nil {
			return err
		}
		if err := s.authCacheRepo.DeleteCode(ctx, verificationTypeEmail, email); err != nil {
			return err
		}
		if err := s.authCacheRepo.DeleteRegisterDraft(ctx, verificationTypeEmail, email); err != nil {
			return err
		}
		return s.saveRegisterOutbox(ctx, created.ID)
	})
}

type StartPhoneRegistrationResp struct {
	Code string
}

func (s *AuthUsecase) StartPhoneRegistration(ctx context.Context, account *model.Account) (*StartPhoneRegistrationResp, error) {
	return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_NOT_IMPLEMENTED)
}

type CheckPhoneRegistrationCodeReq struct {
	Phone string
	Code  string
}

func (s *AuthUsecase) CheckPhoneRegistrationCode(ctx context.Context, req *CheckPhoneRegistrationCodeReq) (bool, error) {
	return false, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_NOT_IMPLEMENTED)
}

type VerifyPhoneRegistrationReq struct {
	Phone string
	Code  string
}

func (s *AuthUsecase) VerifyPhoneRegistration(ctx context.Context, req *VerifyPhoneRegistrationReq) error {
	return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_NOT_IMPLEMENTED)
}

type StartEmailLoginResp struct {
	Code string
}

func (s *AuthUsecase) StartEmailLogin(ctx context.Context, email string) (*StartEmailLoginResp, error) {
	email = normalizeAccountKey(email)
	user, err := s.accountRepo.Get(ctx, &repo.AccountGetReq{Account: &email})
	if err != nil || !isNormalAccount(user) {
		if err != nil && !isAccountNotFound(err) {
			s.logger.WarnContext(ctx, "email login account lookup failed", constant.LogFieldErr, err)
		}
		return &StartEmailLoginResp{}, nil
	}
	now := time.Now()
	codeTTL := s.verificationCodeTTL()
	code := str.RandStr(s.sf, 6, true, true, true, false)
	if err := s.authCacheRepo.SaveCode(ctx, &model.VerificationCode{
		Type:          verificationTypeEmail,
		Account:       email,
		Code:          code,
		MaxAttempts:   s.verificationMaxAttempts(),
		CreatedAtUnix: now.Unix(),
		ExpiresAtUnix: now.Add(codeTTL).Unix(),
	}, codeTTL); err != nil {
		return nil, err
	}
	if err := s.saveVerificationCodeOutbox(ctx, verificationTypeEmail, email, code, codeTTL); err != nil {
		_ = s.authCacheRepo.DeleteCode(ctx, verificationTypeEmail, email)
		return nil, err
	}
	return &StartEmailLoginResp{Code: code}, nil
}

type StartPhoneLoginResp struct {
	Code string
}

func (s *AuthUsecase) StartPhoneLogin(ctx context.Context, phone string) (*StartPhoneLoginResp, error) {
	return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_NOT_IMPLEMENTED)
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
	if req == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	client := normalizeLoginContext(req.Client)
	s.enrichLoginContext(ctx, client)
	audit := &model.LoginLog{LoginType: req.Type, Realm: req.Realm, Status: enum.LoginStatusFailed}
	fillLoginLogClient(audit, client)
	defer func() {
		if _, err := s.loginLogRepo.Create(ctx, audit); err != nil {
			s.logger.WarnContext(ctx, "record login log failed", constant.LogFieldErr, err)
		}
	}()

	var account *model.Account
	var err error
	switch req.Type {
	case enum.LoginTypePassword:
		audit.AccountInput = req.PasswordAccount
		account, err = s.loginByPassword(ctx, req, audit)
	case enum.LoginTypeEmail:
		email := normalizeAccountKey(req.Email)
		audit.AccountInput = email
		account, err = s.loginByEmail(ctx, email, req.Code, audit)
	case enum.LoginTypePhone:
		audit.AccountInput = req.Phone
		audit.FailureReason = ptr(enum.LoginFailureReasonNotImplemented)
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_NOT_IMPLEMENTED)
	default:
		audit.FailureReason = ptr(enum.LoginFailureReasonInvalidCredentials)
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if err != nil {
		return nil, err
	}
	audit.UserID = &account.ID

	tokenPair, err := s.createSession(ctx, account, req.Realm, client)
	if err != nil {
		audit.FailureReason = ptr(enum.LoginFailureReasonInternal)
		return nil, err
	}
	audit.Status = enum.LoginStatusSuccess
	audit.SessionID = tokenPair.SessionID
	if err := s.saveLoginOutbox(ctx, account, req.Type, req.Realm, tokenPair.SessionID, client, audit.AccountInput); err != nil {
		return nil, err
	}
	return &LoginResp{TokenPair: *tokenPair, Account: account}, nil
}

func (s *AuthUsecase) loginByPassword(ctx context.Context, req *LoginReq, audit *model.LoginLog) (*model.Account, error) {
	accountInput := strings.TrimSpace(req.PasswordAccount)
	if accountInput == "" || strings.TrimSpace(req.Password) == "" {
		audit.FailureReason = ptr(enum.LoginFailureReasonInvalidCredentials)
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_INVALID_CREDENTIALS)
	}
	user, err := s.accountRepo.Get(ctx, &repo.AccountGetReq{Account: &accountInput})
	if err != nil || user == nil || !str.VerifyPassword(user.Password, req.Password) || !isNormalAccount(user) {
		if err != nil && !isAccountNotFound(err) {
			s.logger.WarnContext(ctx, "password login account lookup failed", constant.LogFieldErr, err)
		}
		audit.FailureReason = ptr(enum.LoginFailureReasonInvalidCredentials)
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_INVALID_CREDENTIALS)
	}
	audit.UserID = &user.ID
	row, err := s.totpRepo.Get(ctx, &repo.TotpGetReq{UserID: &user.ID})
	if err != nil {
		return nil, err
	}
	if row != nil && row.Enable {
		if strings.TrimSpace(req.Code) == "" || !totp.Validate(req.Code, row.Secret) {
			audit.FailureReason = ptr(enum.LoginFailureReasonTotpInvalid)
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOTP_CODE_INVALID)
		}
	}
	return user, nil
}

func (s *AuthUsecase) loginByEmail(ctx context.Context, email string, code string, audit *model.LoginLog) (*model.Account, error) {
	ok, err := s.checkVerificationCode(ctx, verificationTypeEmail, email, code, true)
	if err != nil {
		return nil, err
	}
	if !ok {
		audit.FailureReason = ptr(enum.LoginFailureReasonCodeInvalidOrExpired)
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_VERIFICATION_CODE_INVALID_OR_EXPIRED)
	}
	user, err := s.accountRepo.Get(ctx, &repo.AccountGetReq{Account: &email})
	if err != nil || !isNormalAccount(user) {
		if err != nil && !isAccountNotFound(err) {
			return nil, err
		}
		audit.FailureReason = ptr(enum.LoginFailureReasonInvalidCredentials)
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_INVALID_CREDENTIALS)
	}
	audit.UserID = &user.ID
	_ = s.authCacheRepo.DeleteCode(ctx, verificationTypeEmail, email)
	return user, nil
}

func (s *AuthUsecase) RefreshToken(ctx context.Context, refreshToken string, realm commonenum.LoginRealm) (*model.TokenPair, error) {
	claims, err := s.tokenUsecase.Parse(refreshToken)
	if err != nil || claims.Type != tokenTypeRefresh || claims.Realm != realm || claims.SessionID == "" || claims.JTI == "" {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_INVALID)
	}
	session, err := s.authCacheRepo.GetSession(ctx, claims.SessionID)
	if err != nil {
		return nil, err
	}
	if session == nil || session.UserID != claims.UserID || session.Realm != realm || time.Now().Unix() >= session.SessionExpiresAtUnix {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_INVALID)
	}
	identity, err := s.tokenIdentity(ctx, claims.UserID)
	if err != nil {
		return nil, err
	}
	if session.CurrentJTI != claims.JTI {
		_ = s.authCacheRepo.DeleteSession(ctx, claims.UserID, claims.SessionID)
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_INVALID)
	}
	newJTI := uuid.NewString()
	now := time.Now()
	ttl := s.sessionRedisTTL(session, now)
	rotated, err := s.authCacheRepo.RotateSessionJTI(ctx, claims.SessionID, claims.JTI, newJTI, now.Unix(), ttl)
	if err != nil {
		return nil, err
	}
	if !rotated {
		_ = s.authCacheRepo.DeleteSession(ctx, claims.UserID, claims.SessionID)
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_INVALID)
	}
	session.CurrentJTI = newJTI
	session.LastSeenAtUnix = now.Unix()
	if err := s.authCacheRepo.TouchSession(ctx, session, ttl); err != nil {
		return nil, err
	}
	accessToken, accessExpiresAt, err := s.tokenUsecase.GenerateAccess(&GenerateAccessTokenReq{
		UserID:    claims.UserID,
		SessionID: claims.SessionID,
		Realm:     realm,
		Name:      identity.Name,
		Nickname:  identity.Nickname,
		Language:  identity.Language,
		Timezone:  identity.Timezone,
	})
	if err != nil {
		return nil, err
	}
	newRefreshToken, refreshExpiresAt, err := s.tokenUsecase.GenerateRefresh(&GenerateRefreshTokenReq{UserID: claims.UserID, SessionID: claims.SessionID, Realm: realm, JTI: newJTI})
	if err != nil {
		return nil, err
	}
	return &model.TokenPair{
		AccessToken:           accessToken,
		RefreshToken:          newRefreshToken,
		AccessTokenExpiresAt:  accessExpiresAt,
		RefreshTokenExpiresAt: minTime(refreshExpiresAt, time.Unix(session.SessionExpiresAtUnix, 0)),
		SessionExpiresAt:      time.Unix(session.SessionExpiresAtUnix, 0),
		SessionID:             claims.SessionID,
	}, nil
}

func (s *AuthUsecase) Logout(ctx context.Context, accessToken string, realm commonenum.LoginRealm) error {
	claims, err := s.tokenUsecase.Parse(accessToken)
	if err != nil || claims.Type != tokenTypeAccess || claims.Realm != realm || claims.SessionID == "" {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_INVALID)
	}
	if err := s.authCacheRepo.DeleteSession(ctx, claims.UserID, claims.SessionID); err != nil {
		return err
	}
	return s.saveLogoutOutbox(ctx, claims.UserID, claims.Name, claims.SessionID, realm)
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
	if err != nil || claims.Type != tokenTypeAccess || claims.Realm != realm || claims.SessionID == "" {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_INVALID)
	}
	session, err := s.authCacheRepo.GetSession(ctx, claims.SessionID)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	if session == nil || session.UserID != claims.UserID || session.Realm != realm || now.Unix() >= session.SessionExpiresAtUnix {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_INVALID)
	}
	session.LastSeenAtUnix = now.Unix()
	if err := s.authCacheRepo.TouchSession(ctx, session, s.sessionRedisTTL(session, now)); err != nil {
		return nil, err
	}
	return &ParseTokenResp{
		User:      &commonModel.User{ID: claims.UserID, Name: claims.Name, Nickname: claims.Nickname, Language: claims.Language, Timezone: claims.Timezone},
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
	user, err := s.accountRepo.Get(ctx, &repo.AccountGetReq{UserID: &req.UserID})
	if err != nil {
		return err
	}
	if !isNormalAccount(user) || !str.VerifyPassword(user.Password, req.Password) {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_INVALID_CREDENTIALS)
	}
	if ok, err := s.validateTotpIfEnabled(ctx, req.UserID, req.Code); err != nil {
		return err
	} else if !ok {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOTP_CODE_INVALID)
	}
	return s.tx(ctx, func(ctx context.Context) error {
		if _, err := s.accountRepo.UpdateStatus(ctx, req.UserID, enum.AccountStatusCancelled); err != nil {
			return err
		}
		if err := s.authCacheRepo.DeleteUserSessions(ctx, req.UserID); err != nil {
			return err
		}
		if err := s.authCacheRepo.DeleteUserRbacPermissions(ctx, req.UserID); err != nil {
			return err
		}
		return s.outboxRepo.Save(ctx, &repo.OutboxEventSave{Event: &commonenums.Event{
			Type:    commonenums.EventType_EVENT_TYPE_USER_ACCOUNT_CANCELLED,
			Subject: commonenums.EventSubject_EVENT_SUBJECT_USER_ACCOUNT_CANCELLED,
			Payload: &commonenums.Event_UserAccountCancelled{UserAccountCancelled: &commonenums.UserAccountCancelledPayload{UserId: req.UserID}},
		}})
	})
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
	if req == nil || req.UserID == 0 || req.OperatorID == 0 || strings.TrimSpace(req.Reason) == "" {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	var record *model.BanRecord
	err := s.tx(ctx, func(ctx context.Context) error {
		var err error
		record, err = s.banRecordRepo.Create(ctx, &model.BanRecord{
			UserID:        req.UserID,
			OperatorID:    req.OperatorID,
			OperatorRealm: req.OperatorRealm,
			Reason:        strings.TrimSpace(req.Reason),
			Remark:        req.Remark,
			StartedAt:     time.Now(),
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
		payload := &commonenums.UserAccountBannedPayload{UserId: req.UserID, OperatorId: req.OperatorID, OperatorRealm: string(req.OperatorRealm), BanRecordId: record.ID, Reason: strings.TrimSpace(req.Reason)}
		if req.BannedUntil != nil {
			payload.BannedUntilUnixSeconds = req.BannedUntil.Unix()
		}
		return s.outboxRepo.Save(ctx, &repo.OutboxEventSave{Event: &commonenums.Event{
			Type:    commonenums.EventType_EVENT_TYPE_USER_ACCOUNT_BANNED,
			Subject: commonenums.EventSubject_EVENT_SUBJECT_USER_ACCOUNT_BANNED,
			Payload: &commonenums.Event_UserAccountBanned{UserAccountBanned: payload},
		}})
	})
	if err != nil {
		return nil, err
	}
	if req.BannedUntil != nil {
		if err := s.delayedTaskClient.RegisterUnbanExpired(ctx, req.UserID, record.ID, *req.BannedUntil); err != nil {
			return nil, err
		}
	}
	return &BanAccountResp{BanRecordID: record.ID}, nil
}

func (s *AuthUsecase) UnbanExpired(ctx context.Context, userID int64, banRecordID int64) (bool, error) {
	user, err := s.accountRepo.Get(ctx, &repo.AccountGetReq{UserID: &userID})
	if err != nil {
		return false, err
	}
	if user == nil || user.Status == nil || *user.Status != enum.AccountStatusBanned {
		return false, nil
	}
	record, err := s.banRecordRepo.Get(ctx, banRecordID)
	if err != nil || record == nil || record.UserID != userID || record.BannedUntil == nil || record.BannedUntil.After(time.Now()) {
		return false, err
	}
	latest, err := s.banRecordRepo.LatestByUserID(ctx, userID)
	if err != nil {
		return false, err
	}
	if latest == nil || latest.ID != banRecordID {
		return false, nil
	}
	err = s.tx(ctx, func(ctx context.Context) error {
		if _, err := s.accountRepo.UpdateStatus(ctx, userID, enum.AccountStatusNormal); err != nil {
			return err
		}
		return s.outboxRepo.Save(ctx, &repo.OutboxEventSave{Event: &commonenums.Event{
			Type:    commonenums.EventType_EVENT_TYPE_USER_ACCOUNT_UNBANNED,
			Subject: commonenums.EventSubject_EVENT_SUBJECT_USER_ACCOUNT_UNBANNED,
			Payload: &commonenums.Event_UserAccountUnbanned{UserAccountUnbanned: &commonenums.UserAccountUnbannedPayload{UserId: userID, BanRecordId: banRecordID}},
		}})
	})
	return err == nil, err
}

func (s *AuthUsecase) createSession(ctx context.Context, account *model.Account, realm commonenum.LoginRealm, client *model.LoginContext) (*model.TokenPair, error) {
	now := time.Now()
	sid := uuid.NewString()
	jti := uuid.NewString()
	sessionExpiresAt := now.Add(s.tokenUsecase.SessionTTL())
	prefs, err := s.prefsRepo.Get(ctx, &repo.PreferencesGetReq{UserID: &account.ID})
	if err != nil {
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
		UserID: account.ID, SessionID: sid, Realm: realm, Name: account.Name, Nickname: nickname, Language: language, Timezone: timezone,
	})
	if err != nil {
		return nil, err
	}
	refreshToken, refreshExpiresAt, err := s.tokenUsecase.GenerateRefresh(&GenerateRefreshTokenReq{UserID: account.ID, SessionID: sid, Realm: realm, JTI: jti})
	if err != nil {
		return nil, err
	}
	session := &model.RefreshSession{
		SessionID:            sid,
		UserID:               account.ID,
		Realm:                realm,
		CurrentJTI:           jti,
		CreatedAtUnix:        now.Unix(),
		LastSeenAtUnix:       now.Unix(),
		SessionExpiresAtUnix: sessionExpiresAt.Unix(),
		Client:               *client,
	}
	if err := s.authCacheRepo.SaveSession(ctx, session, s.sessionRedisTTL(session, now)); err != nil {
		return nil, err
	}
	return &model.TokenPair{
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  accessExpiresAt,
		RefreshTokenExpiresAt: minTime(refreshExpiresAt, sessionExpiresAt),
		SessionExpiresAt:      sessionExpiresAt,
		SessionID:             sid,
	}, nil
}

type tokenIdentity struct {
	Name     string
	Nickname string
	Language commonenums.Language
	Timezone string
}

func (s *AuthUsecase) tokenIdentity(ctx context.Context, userID int64) (*tokenIdentity, error) {
	account, err := s.accountRepo.Get(ctx, &repo.AccountGetReq{UserID: &userID})
	if isAccountNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_INVALID)
	}
	if err != nil {
		return nil, err
	}
	if !isNormalAccount(account) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_INVALID)
	}
	identity := &tokenIdentity{Name: account.Name, Language: commonenums.Language_LANGUAGE_UNSPECIFIED}
	if account.Nickname != nil {
		identity.Nickname = *account.Nickname
	}
	prefs, err := s.prefsRepo.Get(ctx, &repo.PreferencesGetReq{UserID: &userID})
	if err != nil {
		return nil, err
	}
	if prefs != nil {
		if prefs.Language != nil {
			identity.Language = enum.LanguageMap.MustToProto(*prefs.Language)
		}
		if prefs.Timezone != nil {
			identity.Timezone = *prefs.Timezone
		}
	}
	return identity, nil
}

func (s *AuthUsecase) checkVerificationCode(ctx context.Context, codeType string, account string, code string, consumeAttempt bool) (bool, error) {
	row, err := s.authCacheRepo.GetCode(ctx, codeType, account)
	if err != nil {
		return false, err
	}
	if row == nil || row.ExpiresAtUnix <= time.Now().Unix() || row.Attempts >= row.MaxAttempts || row.Code != strings.TrimSpace(code) {
		if consumeAttempt && row != nil && row.Attempts < row.MaxAttempts {
			_, _ = s.authCacheRepo.IncrCodeAttempts(ctx, codeType, account)
		}
		return false, nil
	}
	return true, nil
}

func (s *AuthUsecase) validateTotpIfEnabled(ctx context.Context, userID int64, code string) (bool, error) {
	row, err := s.totpRepo.Get(ctx, &repo.TotpGetReq{UserID: &userID})
	if err != nil {
		return false, err
	}
	if row == nil || !row.Enable {
		return true, nil
	}
	return strings.TrimSpace(code) != "" && totp.Validate(code, row.Secret), nil
}

func (s *AuthUsecase) saveVerificationCodeOutbox(ctx context.Context, codeType string, account string, code string, ttl time.Duration) error {
	payloadSeconds := int64(ttl.Seconds())
	if codeType == verificationTypeEmail {
		return s.outboxRepo.Save(ctx, &repo.OutboxEventSave{Event: &commonenums.Event{
			Type:    commonenums.EventType_EVENT_TYPE_USER_EMAIL_VERIFICATION_CODE,
			Subject: commonenums.EventSubject_EVENT_SUBJECT_USER_EMAIL_VERIFICATION_CODE,
			Payload: &commonenums.Event_UserEmailVerificationCode{UserEmailVerificationCode: &commonenums.UserEmailVerificationCodePayload{Email: account, Code: code, ExpiresSeconds: payloadSeconds}},
		}})
	}
	return s.outboxRepo.Save(ctx, &repo.OutboxEventSave{Event: &commonenums.Event{
		Type:    commonenums.EventType_EVENT_TYPE_USER_PHONE_VERIFICATION_CODE,
		Subject: commonenums.EventSubject_EVENT_SUBJECT_USER_PHONE_VERIFICATION_CODE,
		Payload: &commonenums.Event_UserPhoneVerificationCode{UserPhoneVerificationCode: &commonenums.UserPhoneVerificationCodePayload{Phone: account, Code: code, ExpiresSeconds: payloadSeconds}},
	}})
}

func (s *AuthUsecase) saveLoginOutbox(ctx context.Context, account *model.Account, loginType enum.LoginType, realm commonenum.LoginRealm, sessionID string, client *model.LoginContext, accountInput string) error {
	return s.outboxRepo.Save(ctx, &repo.OutboxEventSave{Event: &commonenums.Event{
		Type:    commonenums.EventType_EVENT_TYPE_USER_LOGIN,
		Subject: commonenums.EventSubject_EVENT_SUBJECT_USER_LOGIN,
		Payload: &commonenums.Event_UserLogin{UserLogin: &commonenums.UserLoginPayload{
			UserId: account.ID, Name: account.Name, Account: accountInput, Ip: client.IP, UserAgent: client.UserAgent,
			Realm: string(realm), LoginType: string(loginType), SessionId: sessionID, AppName: client.AppName, AppVersion: client.AppVersion,
			ClientType: string(client.ClientType), DeviceType: string(client.DeviceType),
		}},
	}})
}

func (s *AuthUsecase) saveLogoutOutbox(ctx context.Context, userID int64, name string, sessionID string, realm commonenum.LoginRealm) error {
	return s.outboxRepo.Save(ctx, &repo.OutboxEventSave{Event: &commonenums.Event{
		Type:    commonenums.EventType_EVENT_TYPE_USER_LOGOUT,
		Subject: commonenums.EventSubject_EVENT_SUBJECT_USER_LOGOUT,
		Payload: &commonenums.Event_UserLogout{UserLogout: &commonenums.UserLogoutPayload{UserId: userID, Name: name, SessionId: sessionID, Realm: string(realm)}},
	}})
}

func (s *AuthUsecase) saveRegisterOutbox(ctx context.Context, userID int64) error {
	return s.outboxRepo.Save(ctx, &repo.OutboxEventSave{Event: &commonenums.Event{
		Type:    commonenums.EventType_EVENT_TYPE_USER_REGISTER,
		Subject: commonenums.EventSubject_EVENT_SUBJECT_USER_REGISTER,
		Payload: &commonenums.Event_UserRegister{UserRegister: &commonenums.UserRegisterPayload{UserId: userID}},
	}})
}

func (s *AuthUsecase) sessionRedisTTL(session *model.RefreshSession, now time.Time) time.Duration {
	remaining := time.Until(time.Unix(session.SessionExpiresAtUnix, 0))
	if remaining <= 0 {
		return time.Second
	}
	refreshTTL := s.tokenUsecase.RefreshTokenTTL()
	if refreshTTL < remaining {
		return refreshTTL
	}
	return remaining
}

func (s *AuthUsecase) verificationCodeTTL() time.Duration {
	return durationOrDefault(s.conf.GetBusiness().GetAuth().GetVerificationCode().GetCodeTtl(), 5*time.Minute)
}

func (s *AuthUsecase) registerDraftTTL() time.Duration {
	return durationOrDefault(s.conf.GetBusiness().GetAuth().GetVerificationCode().GetRegisterDraftTtl(), 30*time.Minute)
}

func (s *AuthUsecase) verificationMaxAttempts() int32 {
	maxAttempts := s.conf.GetBusiness().GetAuth().GetVerificationCode().GetMaxAttempts()
	if maxAttempts <= 0 {
		return 5
	}
	return maxAttempts
}

func normalizeLoginContext(client *model.LoginContext) *model.LoginContext {
	if client == nil {
		client = &model.LoginContext{}
	}
	if client.ClientType == "" {
		client.ClientType = enum.ClientTypeUnknown
	}
	if client.DeviceType == "" {
		client.DeviceType = enum.DeviceTypeUnknown
	}
	return client
}

func (s *AuthUsecase) enrichLoginContext(ctx context.Context, client *model.LoginContext) {
	if client == nil || client.IP == "" || s.ipClient == nil {
		return
	}
	info, err := s.ipClient.Resolve(ctx, client.IP)
	if err != nil {
		s.logger.WarnContext(ctx, "resolve login ip failed", constant.LogFieldErr, err)
		return
	}
	if info == nil {
		return
	}
	client.Country = info.Country
	client.CountryCode = info.CountryCode
	client.Province = info.Province
	client.City = info.City
	client.ISP = info.ISP
}

func fillLoginLogClient(log *model.LoginLog, client *model.LoginContext) {
	if client.IP != "" {
		log.IP = &client.IP
	}
	if client.Country != "" {
		log.Country = &client.Country
	}
	if client.CountryCode != "" {
		log.CountryCode = &client.CountryCode
	}
	if client.Province != "" {
		log.Province = &client.Province
	}
	if client.City != "" {
		log.City = &client.City
	}
	if client.ISP != "" {
		log.ISP = &client.ISP
	}
	if client.UserAgent != "" {
		log.UserAgent = &client.UserAgent
	}
	log.ClientType = &client.ClientType
	log.DeviceType = &client.DeviceType
	log.OSName = client.OSName
	log.OSVersion = client.OSVersion
	log.BrowserName = client.BrowserName
	log.BrowserVersion = client.BrowserVersion
	log.AppName = client.AppName
	log.AppVersion = client.AppVersion
}

func normalizeAccountKey(account string) string {
	return strings.ToLower(strings.TrimSpace(account))
}

func isNormalAccount(account *model.Account) bool {
	return account != nil && account.Status != nil && *account.Status == enum.AccountStatusNormal
}

func isAccountNotFound(err error) bool {
	if err == nil {
		return false
	}
	code, ok := apperror.BusinessCode(err)
	return ok && code == cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_ACCOUNT_NOT_FOUND
}

func minTime(a time.Time, b time.Time) time.Time {
	if a.Before(b) {
		return a
	}
	return b
}

func ptr[T any](value T) *T {
	return &value
}
