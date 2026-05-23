package usecase

import (
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
	"github.com/jinzhu/copier"
	"github.com/sony/sonyflake/v2"
)

type AuthUsecase struct {
	conf         *conf.Bootstrap
	log          *log.Helper
	tx           base.Tx
	eventPool    *util.EventPool
	accountRepo  repo.AccountRepo
	loginLogRepo repo.LoginLogRepo
	tokenCache   *jwt.TokenCache
	tokenService *TokenUsecase

	sf *sonyflake.Sonyflake
}

func NewAuthUsecase(
	conf *conf.Bootstrap,
	logger log.Logger,
	tx base.Tx,
	eventPool *util.EventPool,
	accountRepo repo.AccountRepo,
	loginLogRepo repo.LoginLogRepo,
	tokenCache *jwt.TokenCache,
	tokenService *TokenUsecase,
) (*AuthUsecase, error) {
	sf, err := str.NewSonyflake()
	if err != nil {
		return nil, err
	}
	return &AuthUsecase{
		conf:         conf,
		log:          log.NewHelper(logger),
		tx:           tx,
		eventPool:    eventPool,
		accountRepo:  accountRepo,
		loginLogRepo: loginLogRepo,
		tokenCache:   tokenCache,
		tokenService: tokenService,
		sf:           sf,
	}, nil
}

func (s *AuthUsecase) RegisterEmail(ctx context.Context, u *model.Account) (code string, token string, err error) {
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
	token, err = s.tokenService.VerityCodeAccountTokenGen.Generate(model.TokenVerityCodeAccount{Account: *u.Email}, s.conf.Server.Jwt.EmailExpire.AsDuration())
	if err != nil {
		return
	}
	err = s.eventPool.Submit(func() {})
	if err != nil {
		return
	}

	saveUser := &commonModel.User{}
	err = copier.Copy(saveUser, u)
	if err != nil {
		return
	}
	err = s.tokenCache.SaveVerityCode(ctx, constant.VerifyCodeTypeRegisterEmail, *u.Email, code, saveUser, s.conf.Server.Jwt.EmailExpire.AsDuration())
	if err != nil {
		return
	}

	return code, token, nil
}

func (s *AuthUsecase) RegisterEmailVerify(ctx context.Context, codeToken string, code string) (err error) {
	token, err := s.tokenService.VerityCodeAccountTokenGen.Parse(codeToken)
	if err != nil {
		return
	}
	verityCode, saveUser, err := s.tokenCache.GetVerityCode(ctx, constant.VerifyCodeTypeRegisterEmail, token.Account)
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
			Nickname: &saveUser.Nickname,
			Password: saveUser.Password,
			Email:    &saveUser.Email,
		}
		user.Password, err = str.HashPassword(user.Password)
		if err != nil {
			return err
		}
		_, err = s.accountRepo.Create(ctx, user)
		if err != nil {
			return err
		}
		err = s.tokenCache.DelVerityCode(ctx, constant.VerifyCodeTypeRegisterEmail, token.Account)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return
	}
	return
}

func (s *AuthUsecase) RegisterPhone(ctx context.Context, u *model.Account) (code string, token string, err error) {
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
	token, err = s.tokenService.VerityCodeAccountTokenGen.Generate(model.TokenVerityCodeAccount{Account: *u.Phone}, s.conf.Server.Jwt.PhoneExpire.AsDuration())
	if err != nil {
		return
	}
	err = s.eventPool.Submit(func() {})
	if err != nil {
		return
	}

	saveUser := &commonModel.User{}
	err = copier.Copy(saveUser, u)
	if err != nil {
		return
	}
	err = s.tokenCache.SaveVerityCode(ctx, constant.VerifyCodeTypeRegisterPhone, *u.Phone, code, saveUser, s.conf.Server.Jwt.PhoneExpire.AsDuration())
	if err != nil {
		return
	}

	return code, token, nil
}

func (s *AuthUsecase) RegisterPhoneVerify(ctx context.Context, codeToken string, code string) (err error) {
	token, err := s.tokenService.VerityCodeAccountTokenGen.Parse(codeToken)
	if err != nil {
		return
	}
	verityCode, saveUser, err := s.tokenCache.GetVerityCode(ctx, constant.VerifyCodeTypeRegisterPhone, token.Account)
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
			Nickname: &saveUser.Nickname,
			Password: saveUser.Password,
			Phone:    &saveUser.Phone,
		}
		user.Password, err = str.HashPassword(user.Password)
		if err != nil {
			return err
		}
		_, err = s.accountRepo.Create(ctx, user)
		if err != nil {
			return err
		}
		err = s.tokenCache.DelVerityCode(ctx, constant.VerifyCodeTypeRegisterPhone, token.Account)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return
	}
	return
}

func (s *AuthUsecase) LoginAccount(ctx context.Context, account string, password string) (token string, user *model.Account, err error) {
	user, err = s.accountRepo.GetByAccount(ctx, account)
	if err != nil {
		s.recordLoginLog(ctx, nil, account, enum.LoginStatusFailed, "account lookup failed")
		return
	}
	if !str.VerifyPassword(user.Password, password) {
		s.recordLoginLog(ctx, &user.ID, account, enum.LoginStatusFailed, "password invalid")
		return token, nil, cerrors.ErrorBadRequest("password invalid")
	}

	token, err = s.tokenService.TokenGen.Generate(model.Token{Id: user.ID}, s.conf.Server.Jwt.Expires.AsDuration())
	if err != nil {
		s.recordLoginLog(ctx, &user.ID, account, enum.LoginStatusFailed, "token generate failed")
		return
	}
	saveUser := &commonModel.User{}
	err = copier.Copy(saveUser, user)
	if err != nil {
		s.recordLoginLog(ctx, &user.ID, account, enum.LoginStatusFailed, "token user copy failed")
		return
	}
	err = s.tokenCache.SaveToken(ctx, token, saveUser, s.conf.Server.Jwt.Expires.AsDuration())
	if err != nil {
		s.recordLoginLog(ctx, &user.ID, account, enum.LoginStatusFailed, "token cache save failed")
		return
	}
	s.recordLoginLog(ctx, &user.ID, account, enum.LoginStatusSuccess, "")

	return token, user, nil
}

func (s *AuthUsecase) Logout(ctx context.Context, token string) (err error) {
	return s.tokenCache.DelToken(ctx, token)
}

func (s *AuthUsecase) recordLoginLog(ctx context.Context, userID *int64, account string, status enum.LoginStatus, reason string) {
	loginLog := &model.LoginLog{
		UserID:      userID,
		Account:     account,
		LoginMethod: enum.LoginMethodPassword,
		Status:      status,
	}
	if reason != "" {
		loginLog.FailureReason = stringPtr(reason)
	}

	if ipInfo, ok := util.GetContextValue[*commonModel.IpInfo](ctx, constant.CtxIpInfo); ok && ipInfo != nil {
		loginLog.IP = stringPtr(ipInfo.Ip)
		loginLog.Country = stringPtr(ipInfo.Country)
		loginLog.CountryCode = stringPtr(ipInfo.CountryCode)
		loginLog.Province = stringPtr(ipInfo.Province)
		loginLog.City = stringPtr(ipInfo.City)
		loginLog.ISP = stringPtr(ipInfo.ISP)
	}
	if loginLog.IP == nil {
		loginLog.IP = stringPtr(firstHeaderValue(
			serverutil.GetHeader(ctx, "X-Forwarded-For"),
			serverutil.GetHeader(ctx, "X-Real-IP"),
			serverutil.GetHeader(ctx, "X-Client-IP"),
		))
	}
	loginLog.UserAgent = stringPtr(serverutil.GetHeader(ctx, "User-Agent"))
	loginLog.DeviceID = stringPtr(serverutil.GetHeader(ctx, "X-Device-ID"))
	loginLog.DeviceName = stringPtr(serverutil.GetHeader(ctx, "X-Device-Name"))
	loginLog.Platform = stringPtr(serverutil.GetHeader(ctx, "X-Platform"))
	loginLog.OS = stringPtr(serverutil.GetHeader(ctx, "X-OS"))
	loginLog.Browser = stringPtr(serverutil.GetHeader(ctx, "X-Browser"))
	loginLog.RequestID = stringPtr(firstHeaderValue(
		serverutil.GetHeader(ctx, "X-Request-ID"),
		serverutil.GetHeader(ctx, "X-Trace-ID"),
	))

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

func stringPtr(value string) *string {
	if value == "" {
		return nil
	}
	return &value
}
