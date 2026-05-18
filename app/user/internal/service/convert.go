package service

import (
	v1 "common/api/gen/user/v1"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/enum"

	"context"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func assembleUserProto(ctx context.Context, u *model.User, prefRepo repo.UserPreferencesRepo, privRepo repo.UserPrivacyRepo, locRepo repo.UserLocationRepo, tfaRepo repo.UserTfaRepo, checkinRepo repo.UserCheckinRepo) *v1.User {
	prefs, _ := prefRepo.GetByUserID(ctx, u.ID)
	privacy, _ := privRepo.GetByUserID(ctx, u.ID)
	loc, _ := locRepo.GetByUserID(ctx, u.ID)
	tfa, _ := tfaRepo.GetByUserID(ctx, u.ID)
	stat, _ := checkinRepo.GetStatByUserID(ctx, u.ID)
	return userToProto(u, prefs, privacy, loc, tfa, stat)
}

func userToProto(
	u *model.User,
	prefs *model.UserPreferences,
	privacy *model.UserPrivacy,
	loc *model.UserLocation,
	tfa *model.UserTFA,
	stat *model.UserCheckinStat,
) *v1.User {
	reply := &v1.User{
		Id:            u.ID,
		Name:          u.Name,
		Nickname:      u.Nickname,
		Email:         u.Email,
		Phone:         u.Phone,
		Url:           u.URL,
		AvatarUrl:     u.AvatarURL,
		Introduction:  u.Introduction,
		Mbti:          u.Mbti,
		GroupName:     u.GroupName,
		FollowCount:   u.FollowCount,
		FollowerCount: u.FollowerCount,
		BlockCount:    u.BlockCount,
		BlockedCount:  u.BlockedCount,
		LastLoginIp:   u.LastLoginIP,
	}
	if u.Status != nil {
		v := enum.UserStatusMap.MustToProto(*u.Status)
		reply.Status = &v
	}
	if u.LastLoginTime != nil {
		reply.LastLoginTime = timestamppb.New(*u.LastLoginTime)
	}
	if u.CreatedAt != nil {
		reply.CreatedAt = timestamppb.New(*u.CreatedAt)
	}
	if u.UpdatedAt != nil {
		reply.UpdatedAt = timestamppb.New(*u.UpdatedAt)
	}
	if prefs != nil {
		reply.Language = prefs.Language
		reply.Timezone = prefs.Timezone
		reply.Theme = prefs.Theme
		reply.MobileTheme = prefs.MobileTheme
		reply.EnableWebNotify = prefs.EnableWebNotify
		reply.EnableEmailSubscribe = prefs.EnableEmailSubscribe
	}
	if privacy != nil {
		reply.PublicPoints = privacy.PublicPoints
		reply.PublicFollowers = privacy.PublicFollowers
		reply.PublicArticles = privacy.PublicArticles
		reply.PublicComments = privacy.PublicComments
		reply.PublicOnlineStatus = privacy.PublicOnlineStatus
		reply.PublicLocation = privacy.PublicLocation
	}
	if loc != nil {
		reply.Country = loc.Country
		reply.Province = loc.Province
		reply.City = loc.City
	}
	if tfa != nil {
		reply.TwofaEnable = tfa.Enable
		if tfa.Enable && tfa.EnableTime != nil {
			reply.TwofaEnableTime = timestamppb.New(*tfa.EnableTime)
		}
	}
	if stat != nil {
		reply.OnlineMinutes = stat.TotalOnlineMinutes
		reply.CurrentCheckinStreak = stat.CurrentStreak
		reply.LongestCheckinStreak = stat.LongestStreak
	}
	return reply
}

func userRelationsToProto(relations []*model.UserRelation) []*v1.UserRelation {
	result := make([]*v1.UserRelation, 0, len(relations))
	for _, r := range relations {
		result = append(result, relationToProto(r))
	}
	return result
}

func relationToProto(r *model.UserRelation) *v1.UserRelation {
	reply := &v1.UserRelation{
		Id:       r.ID,
		Type:     enum.UserRelationTypeMap.MustToProto(r.Type),
		ActorId:  r.ActorID,
		TargetId: r.TargetID,
	}
	if r.CreatedAt != nil {
		reply.CreatedAt = timestamppb.New(*r.CreatedAt)
	}
	if r.UpdatedAt != nil {
		reply.UpdatedAt = timestamppb.New(*r.UpdatedAt)
	}
	return reply
}
