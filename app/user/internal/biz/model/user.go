package model

import (
	v1 "common/api/gen/user/v1"
	"common/pkg/util/str"
	"user/internal/data/ent/gen"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type User struct {
	*gen.User
}

func (u *User) PasswordEncrypt() error {
	password, err := str.HashPassword(u.Password)
	u.Password = password
	return err
}

func (u *User) PasswordVerify(password string) bool {
	return str.VerifyPassword(u.Password, password)
}

func (u *User) ConvertToRpc() *v1.User {
	reply := &v1.User{
		Id:                   u.ID,
		Name:                 u.Name,
		Nickname:             u.Nickname,
		Email:                u.Email,
		Phone:                u.Phone,
		Url:                  u.URL,
		AvatarUrl:            u.AvatarURL,
		Introduction:         u.Introduction,
		Mbti:                 u.Mbti,
		Status:               (*v1.UserStatus)(u.Status),
		GroupName:            u.GroupName,
		FollowCount:          u.FollowCount,
		FollowerCount:        u.FollowerCount,
		BlockCount:           u.BlockCount,
		BlockedCount:         u.BlockedCount,
		LastLoginIp:          u.LastLoginIP,
		OnlineMinutes:        u.OnlineMinutes,
		CurrentCheckinStreak: u.CurrentCheckinStreak,
		LongestCheckinStreak: u.LongestCheckinStreak,
		Language:             u.Language,
		Timezone:             u.Timezone,
		Theme:                u.Theme,
		MobileTheme:          u.MobileTheme,
		EnableWebNotify:      u.EnableWebNotify,
		EnableEmailSubscribe: u.EnableEmailSubscribe,
		PublicPoints:         u.PublicPoints,
		PublicFollowers:      u.PublicFollowers,
		PublicArticles:       u.PublicArticles,
		PublicComments:       u.PublicComments,
		PublicOnlineStatus:   u.PublicOnlineStatus,
		Country:              u.Country,
		Province:             u.Province,
		City:                 u.City,
		PublicLocation:       u.PublicLocation,
		TwofaEnable:          u.TwofaEnable,
		CreatedAt:            timestamppb.New(*u.CreatedAt),
		UpdatedAt:            timestamppb.New(*u.UpdatedAt),
	}
	if u.TwofaEnable && u.TwofaEnableTime != nil {
		reply.TwofaEnableTime = timestamppb.New(*u.TwofaEnableTime)
	}
	if u.LastLoginTime != nil {
		reply.LastLoginTime = timestamppb.New(*u.LastLoginTime)
	}
	if u.LastCheckinTime != nil {
		reply.LastCheckinTime = timestamppb.New(*u.LastCheckinTime)
	}
	return reply
}
