package doamin

import (
	"common/api/gen/common"
	notifyv1 "common/api/gen/notify/v1"
	"common/pkg/constant"
	commonModel "common/pkg/model"
	"common/pkg/util/jwt"
	"common/pkg/util/str"
	"context"
	domainbase "user/internal/biz/base"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/data/ent"
	"user/internal/data/ent/gen"

	"github.com/google/uuid"
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

func NewAuthenticationDomain(base *domainbase.BaseDomain, userRepo repo.UserRepo, tokenCache *jwt.TokenCache, tokenService *TokenService) (*AuthenticationDomain, error) {
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
		return "", "", common.ErrorBadRequest("email can not be empty")
	}
	exist, err := s.userRepo.ConstantAccount(ctx, s.Db, *u.Email)
	if exist {
		err = common.ErrorBadRequest("email already exists")
	}
	if err != nil {
		return
	}
	exist, err = s.userRepo.ConstantAccount(ctx, s.Db, u.Name)
	if exist {
		err = common.ErrorBadRequest("name already exists")
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
		err = common.ErrorBadRequest("email verification code has been sent")
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
		err := s.Rabbitmq.Publish(constant.ExchangeUser.String(), constant.RoutingKeyUserRegisterVerifyCode.String(), &commonModel.Notification{
			UUID:       uuid.New().String(),
			Type:       new(notifyv1.NotificationType_NOTIFICATION_TYPE_USER_REGISTER_VERIFY_CODE),
			SenderId:   u.ID,
			SenderName: u.Name,
			Channels:   []*notifyv1.NotificationChannel{new(notifyv1.NotificationChannel_NOTIFICATION_CHANNEL_EMAIL)},
			Meta: commonModel.Meta{
				RegisterVerifyCode: &commonModel.RegisterVerifyCode{
					Email:  *u.Email,
					Code:   code,
					Expire: s.Conf.Server.Jwt.EmailExpire.AsDuration(),
				},
			},
		})
		if err != nil {
			s.Log.Errorf("publish user register verfity code event error: %v", err)
		}
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
		err = common.ErrorBadRequest("email code invalid")
		return
	}

	err = ent.WithTx(ctx, s.Db, func(tx *gen.Client) error {
		// 保存用户信息
		user := &model.User{User: &gen.User{
			Name:     saveUser.Name,
			Nickname: new(saveUser.Nickname),
			Password: saveUser.Password,
			Email:    new(saveUser.Email),
		}}
		err = user.PasswordEncrypt()
		if err != nil {
			return err
		}
		_, err = s.userRepo.Save(ctx, tx, user)
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
		return "", "", common.ErrorBadRequest("phone can not be empty")
	}
	exist, err := s.userRepo.ConstantAccount(ctx, s.Db, *u.Phone)
	if exist {
		err = common.ErrorBadRequest("phone already exists")
	}
	if err != nil {
		return
	}
	exist, err = s.userRepo.ConstantAccount(ctx, s.Db, u.Name)
	if exist {
		err = common.ErrorBadRequest("name already exists")
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
		err = common.ErrorBadRequest("phone verification code has been sent")
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
		err := s.Rabbitmq.Publish(constant.ExchangeUser.String(), constant.RoutingKeyUserRegisterVerifyCode.String(), &commonModel.Notification{
			UUID:       uuid.New().String(),
			Type:       new(notifyv1.NotificationType_NOTIFICATION_TYPE_USER_REGISTER_VERIFY_CODE),
			SenderId:   u.ID,
			SenderName: u.Name,
			Channels:   []*notifyv1.NotificationChannel{new(notifyv1.NotificationChannel_NOTIFICATION_CHANNEL_SMS)},
			Meta: commonModel.Meta{
				RegisterVerifyCode: &commonModel.RegisterVerifyCode{
					Phone:  *u.Phone,
					Code:   code,
					Expire: s.Conf.Server.Jwt.PhoneExpire.AsDuration(),
				},
			},
		})
		if err != nil {
			s.Log.Errorf("publish user register verfity code event error: %v", err)
		}
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
		err = common.ErrorBadRequest("phone code invalid")
		return
	}

	err = ent.WithTx(ctx, s.Db, func(tx *gen.Client) error {
		// 保存用户信息
		user := &model.User{User: &gen.User{
			Name:     saveUser.Name,
			Nickname: new(saveUser.Nickname),
			Password: saveUser.Password,
			Phone:    new(saveUser.Phone),
		}}
		err = user.PasswordEncrypt()
		if err != nil {
			return err
		}
		_, err = s.userRepo.Save(ctx, tx, user)
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
	user, err = s.userRepo.GetByAccount(ctx, s.Db, account)
	if err != nil {
		return
	}
	// 验证密码
	if !user.PasswordVerify(password) {
		return token, nil, common.ErrorBadRequest("password invalid")
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
