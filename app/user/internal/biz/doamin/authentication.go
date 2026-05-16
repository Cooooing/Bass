package doamin

import (
	cerrors "common/api/gen/common/errors"
	"common/pkg/constant"
	commonModel "common/pkg/model"
	"common/pkg/util/jwt"
	"common/pkg/util/str"
	"context"
	domainbase "user/internal/biz/base"
	"user/internal/biz/model"
	"user/internal/biz/repo"

	"github.com/jinzhu/copier"
	"github.com/sony/sonyflake/v2"
)

type AuthenticationDomain struct {
	*domainbase.BaseDomain
	userRepo     repo.UserRepo
	tokenCache   *jwt.TokenCache
	tokenService *TokenService

	sf *sonyflake.Sonyflake
}

func NewAuthenticationDomain(
	base *domainbase.BaseDomain,
	userRepo repo.UserRepo,
	tokenCache *jwt.TokenCache,
	tokenService *TokenService,
) (*AuthenticationDomain, error) {
	sf, err := str.NewSonyflake()
	if err != nil {
		return nil, err
	}
	return &AuthenticationDomain{
		BaseDomain:   base,
		userRepo:     userRepo,
		tokenCache:   tokenCache,
		tokenService: tokenService,
		sf:           sf,
	}, nil
}

func (s *AuthenticationDomain) RegisterEmail(ctx context.Context, u *model.User) (code string, token string, err error) {
	// 验证数据
	if u.Email == nil {
		return "", "", cerrors.ErrorBadRequest("email can not be empty")
	}
	exist, err := s.userRepo.ConstantAccount(ctx, *u.Email)
	if exist {
		err = cerrors.ErrorBadRequest("email already exists")
	}
	if err != nil {
		return
	}
	exist, err = s.userRepo.ConstantAccount(ctx, u.Name)
	if exist {
		err = cerrors.ErrorBadRequest("name already exists")
	}
	if err != nil {
		return
	}

	// 该邮箱是否在缓存
	existEmailCode, err := s.tokenCache.ExistVerityCode(ctx, constant.VerifyCodeTypeRegisterEmail, *u.Email)
	if err != nil {
		return
	}
	if existEmailCode {
		err = cerrors.ErrorBadRequest("email verification code has been sent")
		return
	}

	// 生成 code
	code = str.RandStr(s.sf, 6, true, true, true, false)
	token, err = s.tokenService.VerityCodeAccountTokenGen.Generate(model.TokenVerityCodeAccount{Account: *u.Email}, s.Conf.Server.Jwt.EmailExpire.AsDuration())
	if err != nil {
		return
	}
	// 发送邮件验证码通知
	err = s.EventPool.Submit(func() {
	})
	if err != nil {
		return
	}

	// 保存 code 到缓存
	saveUser := &commonModel.User{}
	err = copier.Copy(saveUser, u)
	if err != nil {
		return
	}
	err = s.tokenCache.SaveVerityCode(ctx, constant.VerifyCodeTypeRegisterEmail, *u.Email, code, saveUser, s.Conf.Server.Jwt.EmailExpire.AsDuration())
	if err != nil {
		return
	}

	return code, token, nil
}

func (s *AuthenticationDomain) RegisterEmailVerify(ctx context.Context, codeToken string, code string) (err error) {
	// 通过 token 获取 code
	token, err := s.tokenService.VerityCodeAccountTokenGen.Parse(codeToken)
	if err != nil {
		return
	}
	verityCode, saveUser, err := s.tokenCache.GetVerityCode(ctx, constant.VerifyCodeTypeRegisterEmail, token.Account)
	if err != nil {
		return
	}
	// 验证 code
	if verityCode != code {
		err = cerrors.ErrorBadRequest("email code invalid")
		return
	}

	err = s.TxRunner(ctx, func(ctx context.Context) error {
		// 保存用户信息
		user := &model.User{
			Name:     saveUser.Name,
			Nickname: &saveUser.Nickname,
			Password: saveUser.Password,
			Email:    &saveUser.Email,
		}
		user.Password, err = str.HashPassword(user.Password)
		if err != nil {
			return err
		}
		_, err = s.userRepo.Save(ctx, user)
		if err != nil {
			return err
		}

		// 删除 code 缓存
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

func (s *AuthenticationDomain) RegisterPhone(ctx context.Context, u *model.User) (code string, token string, err error) {
	// 验证数据
	if u.Phone == nil {
		return "", "", cerrors.ErrorBadRequest("phone can not be empty")
	}
	exist, err := s.userRepo.ConstantAccount(ctx, *u.Phone)
	if exist {
		err = cerrors.ErrorBadRequest("phone already exists")
	}
	if err != nil {
		return
	}
	exist, err = s.userRepo.ConstantAccount(ctx, u.Name)
	if exist {
		err = cerrors.ErrorBadRequest("name already exists")
	}
	if err != nil {
		return
	}

	// 该邮箱是否在缓存
	existPhoneCode, err := s.tokenCache.ExistVerityCode(ctx, constant.VerifyCodeTypeRegisterPhone, *u.Phone)
	if err != nil {
		return
	}
	if existPhoneCode {
		err = cerrors.ErrorBadRequest("phone verification code has been sent")
		return
	}

	// 生成 code
	code = str.RandStr(s.sf, 6, true, true, true, false)
	token, err = s.tokenService.VerityCodeAccountTokenGen.Generate(model.TokenVerityCodeAccount{Account: *u.Phone}, s.Conf.Server.Jwt.PhoneExpire.AsDuration())
	if err != nil {
		return
	}
	// 发送邮件验证码通知
	err = s.EventPool.Submit(func() {
	})
	if err != nil {
		return
	}

	// 保存 code 到缓存
	saveUser := &commonModel.User{}
	err = copier.Copy(saveUser, u)
	if err != nil {
		return
	}
	err = s.tokenCache.SaveVerityCode(ctx, constant.VerifyCodeTypeRegisterPhone, *u.Phone, code, saveUser, s.Conf.Server.Jwt.PhoneExpire.AsDuration())
	if err != nil {
		return
	}

	return code, token, nil
}

func (s *AuthenticationDomain) RegisterPhoneVerify(ctx context.Context, codeToken string, code string) (err error) {
	// 通过 token 获取 code
	token, err := s.tokenService.VerityCodeAccountTokenGen.Parse(codeToken)
	if err != nil {
		return
	}
	verityCode, saveUser, err := s.tokenCache.GetVerityCode(ctx, constant.VerifyCodeTypeRegisterPhone, token.Account)
	if err != nil {
		return
	}
	// 验证 code
	if verityCode != code {
		err = cerrors.ErrorBadRequest("phone code invalid")
		return
	}

	err = s.TxRunner(ctx, func(ctx context.Context) error {
		// 保存用户信息
		user := &model.User{
			Name:     saveUser.Name,
			Nickname: &saveUser.Nickname,
			Password: saveUser.Password,
			Phone:    &saveUser.Phone,
		}
		user.Password, err = str.HashPassword(user.Password)
		if err != nil {
			return err
		}
		_, err = s.userRepo.Save(ctx, user)
		if err != nil {
			return err
		}

		// 删除 code 缓存
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

func (s *AuthenticationDomain) LoginAccount(ctx context.Context, account string, password string) (token string, user *model.User, err error) {
	// 获取用户信息
	user, err = s.userRepo.GetByAccount(ctx, account)
	if err != nil {
		return
	}
	// 验证密码
	if !str.VerifyPassword(user.Password, password) {
		return token, nil, cerrors.ErrorBadRequest("password invalid")
	}
	// 生成 token
	token, err = s.tokenService.TokenGen.Generate(model.Token{Id: user.ID}, s.Conf.Server.Jwt.Expires.AsDuration())
	if err != nil {
		return
	}
	// 保存 token 到缓存
	saveUser := &commonModel.User{}
	err = copier.Copy(saveUser, user)
	if err != nil {
		return
	}
	err = s.tokenCache.SaveToken(ctx, token, saveUser, s.Conf.Server.Jwt.Expires.AsDuration())
	if err != nil {
		return
	}

	return token, user, nil
}

func (s *AuthenticationDomain) Logout(ctx context.Context, token string) (err error) {
	return s.tokenCache.DelToken(ctx, token)
}
