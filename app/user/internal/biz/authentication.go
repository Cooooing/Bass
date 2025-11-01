package biz

import (
	commonModel "common/pkg/model"
	"common/pkg/util"
	"context"
	"errors"
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
		err = errors.New("email already exists")
	}
	if err != nil {
		return
	}
	exist, err = s.userRepo.ConstantAccount(ctx, s.db, u.Nickname)
	if exist {
		err = errors.New("nickname already exists")
	}
	if err != nil {
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
	err = s.tokenRepo.SaveEmailToken(ctx, token, code, saveUser, s.conf.Jwt.EmailExpire.AsDuration())
	if err != nil {
		return
	}

	return code, token, nil
}

func (s *AuthenticationDomain) RegisterEmailVerify(ctx context.Context, codeToken string, code string) (err error) {
	// 通过 token 获取 code
	emailCode, saveUser, err := s.tokenRepo.GetEmailToken(ctx, codeToken)
	if err != nil {
		return
	}
	// 验证 code
	if emailCode != code {
		err = errors.New("email code invalid")
		return
	}

	err = ent.WithTx(ctx, s.db, func(tx *gen.Client) error {
		// 保存用户信息
		user := &model.User{}
		err = copier.Copy(user, saveUser)
		if err != nil {
			return err
		}
		err = user.PasswordEncrypt()
		if err != nil {
			return err
		}
		_, err = s.userRepo.Save(ctx, s.db, user)
		if err != nil {
			return err
		}

		// 删除 code 缓存
		err = s.tokenRepo.DelEmailToken(ctx, codeToken)
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

func (s *AuthenticationDomain) LoginAccount(ctx context.Context, account string, password string) (token string, err error) {
	// 获取用户信息
	user, err := s.userRepo.GetUserByAccount(ctx, s.db, account)
	if err != nil {
		return
	}
	// 验证密码
	if !user.PasswordVerify(password) {
		return token, errors.New("password invalid")
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

	return token, nil
}
