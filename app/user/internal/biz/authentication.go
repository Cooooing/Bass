package biz

import (
	cv1 "common/api/common/v1"
	commonModel "common/pkg/model"
	"common/pkg/util"
	"context"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/data/ent"
	"user/internal/data/ent/gen"

	"github.com/jinzhu/copier"
	"github.com/sony/sonyflake/v2"
)

type AuthenticationDomain struct {
	*BaseDomain
	userRepo     repo.UserRepo
	tokenRepo    *util.TokenRepo
	tokenService *TokenService

	sf *sonyflake.Sonyflake
}

func NewAuthenticationDomain(base *BaseDomain, userRepo repo.UserRepo, tokenRepo *util.TokenRepo, tokenService *TokenService) (*AuthenticationDomain, error) {
	sf, err := util.NewSonyflake()
	if err != nil {
		return nil, err
	}
	return &AuthenticationDomain{
		BaseDomain:   base,
		userRepo:     userRepo,
		tokenRepo:    tokenRepo,
		tokenService: tokenService,
		sf:           sf,
	}, nil
}

func (s *AuthenticationDomain) RegisterEmail(ctx context.Context, u *model.User) (code string, token string, err error) {
	// 验证数据
	exist, err := s.userRepo.ConstantAccount(ctx, s.db, u.Email)
	if exist {
		err = cv1.ErrorBadRequest("email already exists")
	}
	if err != nil {
		return
	}
	exist, err = s.userRepo.ConstantAccount(ctx, s.db, u.Nickname)
	if exist {
		err = cv1.ErrorBadRequest("nickname already exists")
	}
	if err != nil {
		return
	}

	// 该邮箱是否在缓存
	existEmailCode, err := s.tokenRepo.ExistEmailVerificationCode(ctx, u.Email)
	if err != nil {
		return
	}
	if existEmailCode {
		err = cv1.ErrorBadRequest("email verification code has been sent")
		return
	}

	// 生成 code
	code = util.RandStr(s.sf, 6, true, true, true, false)
	token, err = s.tokenService.EmailTokenGen.Generate(model.TokenEmail{
		Email: u.Email,
	})
	if err != nil {
		return
	}
	// Todo 发送邮件

	// 保存 code 到缓存
	saveUser := &commonModel.User{}
	err = copier.Copy(saveUser, u)
	if err != nil {
		return
	}
	err = s.tokenRepo.SaveEmailVerificationCode(ctx, u.Email, code, saveUser, s.conf.Jwt.EmailExpire.AsDuration())
	if err != nil {
		return
	}

	return code, token, nil
}

func (s *AuthenticationDomain) RegisterEmailVerify(ctx context.Context, codeToken string, code string) (err error) {
	// 通过 token 获取 code
	tokenEmail, err := s.tokenService.EmailTokenGen.Parse(codeToken)
	if err != nil {
		return
	}
	emailCode, saveUser, err := s.tokenRepo.GetEmailVerificationCode(ctx, tokenEmail.Email)
	if err != nil {
		return
	}
	// 验证 code
	if emailCode != code {
		err = cv1.ErrorBadRequest("email code invalid")
		return
	}

	err = ent.WithTx(ctx, s.db, func(tx *gen.Client) error {
		// 保存用户信息
		user := &model.User{
			Name:     saveUser.Name,
			Nickname: saveUser.Nickname,
			Password: saveUser.Password,
			Email:    saveUser.Email,
		}
		err = user.PasswordEncrypt()
		if err != nil {
			return err
		}
		_, err = s.userRepo.Save(ctx, tx, user)
		if err != nil {
			return err
		}

		// 删除 code 缓存
		err = s.tokenRepo.DelEmailVerificationCode(ctx, tokenEmail.Email)
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
	user, err = s.userRepo.GetByAccount(ctx, s.db, account)
	if err != nil {
		return
	}
	// 验证密码
	if !user.PasswordVerify(password) {
		return token, nil, cv1.ErrorBadRequest("password invalid")
	}
	// 生成 token
	token, err = s.tokenService.TokenGen.Generate(model.Token{
		User:     user,
		IsOnline: true,
	})
	if err != nil {
		return
	}
	// 保存 token 到缓存
	saveUser := &commonModel.User{}
	err = copier.Copy(saveUser, user)
	if err != nil {
		return
	}
	err = s.tokenRepo.SaveToken(ctx, token, saveUser, s.conf.Jwt.Expires.AsDuration())
	if err != nil {
		return
	}

	return token, user, nil
}

func (s *AuthenticationDomain) Logout(ctx context.Context, token string) (err error) {
	return s.tokenRepo.DelToken(ctx, token)
}
