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
	"strings"
	base "user/internal/biz/base"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/conf"
	"user/internal/enum"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/sony/sonyflake/v2"
)

type AuthUsecase struct {
	conf         *conf.Bootstrap
	log          *log.Helper
	tx           base.Tx
	accountRepo  repo.AccountRepo
	loginLogRepo repo.LoginLogRepo
	outboxRepo   repo.OutboxEventRepo
	tokenCache   *jwt.TokenCache
	tokenUsecase *TokenUsecase

	sf *sonyflake.Sonyflake
}

type registerAccountCache struct {
	Name     string  `json:"name"`
	Nickname *string `json:"nickname,omitempty"`
	Password string  `json:"password"`
	Email    *string `json:"email,omitempty"`
	Phone    *string `json:"phone,omitempty"`
}

func NewAuthUsecase(
	conf *conf.Bootstrap,
	logger log.Logger,
	tx base.Tx,
	accountRepo repo.AccountRepo,
	loginLogRepo repo.LoginLogRepo,
	outboxRepo repo.OutboxEventRepo,
	tokenCache *jwt.TokenCache,
	tokenUsecase *TokenUsecase,
) (*AuthUsecase, error) {
	sf, err := str.NewSonyflake()
	if err != nil {
		return nil, err
	}
	return &AuthUsecase{
		conf:         conf,
		log:          log.NewHelper(logger),
		tx:           tx,
		accountRepo:  accountRepo,
		loginLogRepo: loginLogRepo,
		outboxRepo:   outboxRepo,
		tokenCache:   tokenCache,
		tokenUsecase: tokenUsecase,
		sf:           sf,
	}, nil
}

func (s *AuthUsecase) StartEmailRegistration(ctx context.Context, u *model.Account) (code string, token string, err error) {
	if u.Email == nil {
		return "", "", cerrors.ErrorBadRequest("email can not be empty")
	}
	exist, err := s.accountRepo.ExistsByAccount(ctx, *u.Email)
	if exist {
		err = cerrors.ErrorBadRequest("email already exists")
	}
	if err != nil {
		return
	}
	exist, err = s.accountRepo.ExistsByAccount(ctx, u.Name)
	if exist {
		err = cerrors.ErrorBadRequest("name already exists")
	}
	if err != nil {
		return
	}

	existEmailCode, err := s.tokenCache.ExistVerityCode(ctx, constant.VerifyCodeTypeRegisterEmail, *u.Email)
	if err != nil {
		return
	}
	if existEmailCode {
		err = cerrors.ErrorBadRequest("email verification code has been sent")
		return
	}

	code = str.RandStr(s.sf, 6, true, true, true, false)
	token, err = s.tokenUsecase.VerityCodeAccountTokenGen.Generate(model.TokenVerityCodeAccount{Account: *u.Email}, s.conf.Server.Jwt.EmailExpire.AsDuration())
	if err != nil {
		return
	}

	err = s.tokenCache.SaveVerityCode(ctx, constant.VerifyCodeTypeRegisterEmail, *u.Email, code, &registerAccountCache{
		Name:     u.Name,
		Nickname: u.Nickname,
		Password: u.Password,
		Email:    u.Email,
	}, s.conf.Server.Jwt.EmailExpire.AsDuration())
	if err != nil {
		return
	}
	return code, token, nil
}

func (s *AuthUsecase) VerifyEmailRegistration(ctx context.Context, codeToken string, code string) (err error) {
	token, err := s.tokenUsecase.VerityCodeAccountTokenGen.Parse(codeToken)
	if err != nil {
		return
	}
	saveUser := &registerAccountCache{}
	verityCode, err := s.tokenCache.GetVerityCode(ctx, constant.VerifyCodeTypeRegisterEmail, token.Account, saveUser)
	if err != nil {
		return
	}
	if verityCode != code {
		err = cerrors.ErrorBadRequest("email code invalid")
		return
	}

	err = s.tx(ctx, func(ctx context.Context) error {
		user := &model.Account{
			Name:     saveUser.Name,
			Nickname: saveUser.Nickname,
			Password: saveUser.Password,
			Email:    saveUser.Email,
		}
		user.Password, err = str.HashPassword(user.Password)
		if err != nil {
			return err
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
	if u.Phone == nil {
		return "", "", cerrors.ErrorBadRequest("phone can not be empty")
	}
	exist, err := s.accountRepo.ExistsByAccount(ctx, *u.Phone)
	if exist {
		err = cerrors.ErrorBadRequest("phone already exists")
	}
	if err != nil {
		return
	}
	exist, err = s.accountRepo.ExistsByAccount(ctx, u.Name)
	if exist {
		err = cerrors.ErrorBadRequest("name already exists")
	}
	if err != nil {
		return
	}

	existPhoneCode, err := s.tokenCache.ExistVerityCode(ctx, constant.VerifyCodeTypeRegisterPhone, *u.Phone)
	if err != nil {
		return
	}
	if existPhoneCode {
		err = cerrors.ErrorBadRequest("phone verification code has been sent")
		return
	}

	code = str.RandStr(s.sf, 6, true, true, true, false)
	token, err = s.tokenUsecase.VerityCodeAccountTokenGen.Generate(model.TokenVerityCodeAccount{Account: *u.Phone}, s.conf.Server.Jwt.PhoneExpire.AsDuration())
	if err != nil {
		return
	}

	err = s.tokenCache.SaveVerityCode(ctx, constant.VerifyCodeTypeRegisterPhone, *u.Phone, code, &registerAccountCache{
		Name:     u.Name,
		Nickname: u.Nickname,
		Password: u.Password,
		Phone:    u.Phone,
	}, s.conf.Server.Jwt.PhoneExpire.AsDuration())
	if err != nil {
		return
	}
	return code, token, nil
}

func (s *AuthUsecase) VerifyPhoneRegistration(ctx context.Context, codeToken string, code string) (err error) {
	token, err := s.tokenUsecase.VerityCodeAccountTokenGen.Parse(codeToken)
	if err != nil {
		return
	}
	saveUser := &registerAccountCache{}
	verityCode, err := s.tokenCache.GetVerityCode(ctx, constant.VerifyCodeTypeRegisterPhone, token.Account, saveUser)
	if err != nil {
		return
	}
	if verityCode != code {
		err = cerrors.ErrorBadRequest("phone code invalid")
		return
	}

	err = s.tx(ctx, func(ctx context.Context) error {
		user := &model.Account{
			Name:     saveUser.Name,
			Nickname: saveUser.Nickname,
			Password: saveUser.Password,
			Phone:    saveUser.Phone,
		}
		user.Password, err = str.HashPassword(user.Password)
		if err != nil {
			return err
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
	user, err = s.accountRepo.GetByAccount(ctx, account)
	if err != nil {
		s.recordLoginLog(ctx, nil, enum.LoginStatusFailed)
		return
	}
	if !str.VerifyPassword(user.Password, password) {
		s.recordLoginLog(ctx, &user.ID, enum.LoginStatusFailed)
		return token, nil, cerrors.ErrorBadRequest("password invalid")
	}

	token, err = s.tokenUsecase.TokenGen.Generate(model.Token{Id: user.ID}, s.conf.Server.Jwt.Expires.AsDuration())
	if err != nil {
		s.recordLoginLog(ctx, &user.ID, enum.LoginStatusFailed)
		return
	}
	saveUser := &commonModel.User{
		ID:   user.ID,
		Name: user.Name,
	}
	if user.Nickname != nil {
		saveUser.Nickname = *user.Nickname
	}
	err = s.tokenCache.SaveToken(ctx, token, saveUser, s.conf.Server.Jwt.Expires.AsDuration())
	if err != nil {
		s.recordLoginLog(ctx, &user.ID, enum.LoginStatusFailed)
		return
	}
	s.recordLoginLog(ctx, &user.ID, enum.LoginStatusSuccess)
	loginPayload := &commonenums.UserLoginPayload{
		UserId:    user.ID,
		Name:      user.Name,
		Account:   account,
		UserAgent: serverutil.GetHeader(ctx, "User-Agent"),
		DeviceId:  serverutil.GetHeader(ctx, "X-Device-ID"),
		Platform:  serverutil.GetHeader(ctx, "X-Platform"),
		RequestId: firstHeaderValue(
			serverutil.GetHeader(ctx, "X-Request-ID"),
			serverutil.GetHeader(ctx, "X-Trace-ID"),
		),
	}
	if ipInfo, ok := util.GetContextValue[*commonModel.IpInfo](ctx, constant.CtxIpInfo); ok && ipInfo != nil {
		loginPayload.Ip = ipInfo.Ip
	}
	if loginPayload.Ip == "" {
		loginPayload.Ip = firstHeaderValue(
			serverutil.GetHeader(ctx, "X-Forwarded-For"),
			serverutil.GetHeader(ctx, "X-Real-IP"),
			serverutil.GetHeader(ctx, "X-Client-IP"),
		)
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

func (s *AuthUsecase) recordLoginLog(ctx context.Context, userID *int64, status enum.LoginStatus) {
	loginLog := &model.LoginLog{
		UserID:      userID,
		LoginMethod: enum.LoginMethodPassword,
		Status:      status,
	}

	if ipInfo, ok := util.GetContextValue[*commonModel.IpInfo](ctx, constant.CtxIpInfo); ok && ipInfo != nil {
		loginLog.IP = optionalString(ipInfo.Ip)
		loginLog.Country = optionalString(ipInfo.Country)
		loginLog.CountryCode = optionalString(ipInfo.CountryCode)
		loginLog.Province = optionalString(ipInfo.Province)
		loginLog.City = optionalString(ipInfo.City)
		loginLog.ISP = optionalString(ipInfo.ISP)
	}
	if loginLog.IP == nil {
		loginLog.IP = optionalString(firstHeaderValue(
			serverutil.GetHeader(ctx, "X-Forwarded-For"),
			serverutil.GetHeader(ctx, "X-Real-IP"),
			serverutil.GetHeader(ctx, "X-Client-IP"),
		))
	}
	loginLog.UserAgent = optionalString(serverutil.GetHeader(ctx, "User-Agent"))
	loginLog.DeviceID = optionalString(serverutil.GetHeader(ctx, "X-Device-ID"))

	if _, err := s.loginLogRepo.Create(ctx, loginLog); err != nil {
		s.log.Warnf("record login log failed: %v", err)
	}
}

func firstHeaderValue(values ...string) string {
	for _, value := range values {
		for _, item := range strings.Split(value, ",") {
			item = strings.TrimSpace(item)
			if item != "" {
				return item
			}
		}
	}
	return ""
}

func optionalString(value string) *string {
	if value == "" {
		return nil
	}
	return new(value)
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
