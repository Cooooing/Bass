package usecase

import (
	"bbs/internal/biz/repo"
	bbsuserv1 "common/api/gen/bbs/v1/user"
	"context"
)

type UserUsecase struct {
	userRepo repo.UserRepo
}

func NewUserUsecase(userRepo repo.UserRepo) *UserUsecase {
	return &UserUsecase{userRepo: userRepo}
}

func (u *UserUsecase) StartEmailRegistration(ctx context.Context, req *bbsuserv1.StartEmailRegistration_Request) (*bbsuserv1.StartEmailRegistration_Reply, error) {
	return u.userRepo.StartEmailRegistration(ctx, req)
}

func (u *UserUsecase) VerifyEmailRegistration(ctx context.Context, req *bbsuserv1.VerifyEmailRegistration_Request) (*bbsuserv1.VerifyEmailRegistration_Reply, error) {
	return u.userRepo.VerifyEmailRegistration(ctx, req)
}

func (u *UserUsecase) StartPhoneRegistration(ctx context.Context, req *bbsuserv1.StartPhoneRegistration_Request) (*bbsuserv1.StartPhoneRegistration_Reply, error) {
	return u.userRepo.StartPhoneRegistration(ctx, req)
}

func (u *UserUsecase) VerifyPhoneRegistration(ctx context.Context, req *bbsuserv1.VerifyPhoneRegistration_Request) (*bbsuserv1.VerifyPhoneRegistration_Reply, error) {
	return u.userRepo.VerifyPhoneRegistration(ctx, req)
}

func (u *UserUsecase) LoginByPassword(ctx context.Context, req *bbsuserv1.LoginByPassword_Request) (*bbsuserv1.LoginByPassword_Reply, error) {
	return u.userRepo.LoginByPassword(ctx, req)
}

func (u *UserUsecase) Logout(ctx context.Context, req *bbsuserv1.Logout_Request) (*bbsuserv1.Logout_Reply, error) {
	return u.userRepo.Logout(ctx, req)
}

func (u *UserUsecase) GetCurrentAccount(ctx context.Context, req *bbsuserv1.GetCurrentAccount_Request) (*bbsuserv1.GetCurrentAccount_Reply, error) {
	return u.userRepo.GetCurrentAccount(ctx, req)
}

func (u *UserUsecase) GetProfileAccount(ctx context.Context, req *bbsuserv1.GetProfileAccount_Request) (*bbsuserv1.GetProfileAccount_Reply, error) {
	return u.userRepo.GetProfileAccount(ctx, req)
}

func (u *UserUsecase) UpdateProfileAccount(ctx context.Context, req *bbsuserv1.UpdateProfileAccount_Request) (*bbsuserv1.UpdateProfileAccount_Reply, error) {
	return u.userRepo.UpdateProfileAccount(ctx, req)
}

func (u *UserUsecase) GetCurrentPreferences(ctx context.Context, req *bbsuserv1.GetCurrentPreferences_Request) (*bbsuserv1.GetCurrentPreferences_Reply, error) {
	return u.userRepo.GetCurrentPreferences(ctx, req)
}

func (u *UserUsecase) UpdateCurrentPreferences(ctx context.Context, req *bbsuserv1.UpdateCurrentPreferences_Request) (*bbsuserv1.UpdateCurrentPreferences_Reply, error) {
	return u.userRepo.UpdateCurrentPreferences(ctx, req)
}

func (u *UserUsecase) GetCurrentPrivacySetting(ctx context.Context, req *bbsuserv1.GetCurrentPrivacySetting_Request) (*bbsuserv1.GetCurrentPrivacySetting_Reply, error) {
	return u.userRepo.GetCurrentPrivacySetting(ctx, req)
}

func (u *UserUsecase) UpdateCurrentPrivacySetting(ctx context.Context, req *bbsuserv1.UpdateCurrentPrivacySetting_Request) (*bbsuserv1.UpdateCurrentPrivacySetting_Reply, error) {
	return u.userRepo.UpdateCurrentPrivacySetting(ctx, req)
}

func (u *UserUsecase) GetCurrentLocation(ctx context.Context, req *bbsuserv1.GetCurrentLocation_Request) (*bbsuserv1.GetCurrentLocation_Reply, error) {
	return u.userRepo.GetCurrentLocation(ctx, req)
}

func (u *UserUsecase) UpsertCurrentLocation(ctx context.Context, req *bbsuserv1.UpsertCurrentLocation_Request) (*bbsuserv1.UpsertCurrentLocation_Reply, error) {
	return u.userRepo.UpsertCurrentLocation(ctx, req)
}

func (u *UserUsecase) Follow(ctx context.Context, req *bbsuserv1.FollowRelation_Request) (*bbsuserv1.FollowRelation_Reply, error) {
	return u.userRepo.Follow(ctx, req)
}

func (u *UserUsecase) Unfollow(ctx context.Context, req *bbsuserv1.UnfollowRelation_Request) (*bbsuserv1.UnfollowRelation_Reply, error) {
	return u.userRepo.Unfollow(ctx, req)
}

func (u *UserUsecase) Block(ctx context.Context, req *bbsuserv1.BlockRelation_Request) (*bbsuserv1.BlockRelation_Reply, error) {
	return u.userRepo.Block(ctx, req)
}

func (u *UserUsecase) Unblock(ctx context.Context, req *bbsuserv1.UnblockRelation_Request) (*bbsuserv1.UnblockRelation_Reply, error) {
	return u.userRepo.Unblock(ctx, req)
}

func (u *UserUsecase) ListFollowing(ctx context.Context, req *bbsuserv1.ListFollowingRelations_Request) (*bbsuserv1.ListFollowingRelations_Reply, error) {
	return u.userRepo.ListFollowing(ctx, req)
}

func (u *UserUsecase) ListFollowers(ctx context.Context, req *bbsuserv1.ListFollowersRelations_Request) (*bbsuserv1.ListFollowersRelations_Reply, error) {
	return u.userRepo.ListFollowers(ctx, req)
}

func (u *UserUsecase) ListBlocked(ctx context.Context, req *bbsuserv1.ListBlockedRelations_Request) (*bbsuserv1.ListBlockedRelations_Reply, error) {
	return u.userRepo.ListBlocked(ctx, req)
}

func (u *UserUsecase) GetStatus(ctx context.Context, req *bbsuserv1.GetStatusRelation_Request) (*bbsuserv1.GetStatusRelation_Reply, error) {
	return u.userRepo.GetStatus(ctx, req)
}

func (u *UserUsecase) ValidateTfa(ctx context.Context, req *bbsuserv1.ValidateTfa_Request) (*bbsuserv1.ValidateTfa_Reply, error) {
	return u.userRepo.ValidateTfa(ctx, req)
}

func (u *UserUsecase) BeginEnableTfa(ctx context.Context, req *bbsuserv1.BeginEnableTfa_Request) (*bbsuserv1.BeginEnableTfa_Reply, error) {
	return u.userRepo.BeginEnableTfa(ctx, req)
}

func (u *UserUsecase) ConfirmEnableTfa(ctx context.Context, req *bbsuserv1.ConfirmEnableTfa_Request) (*bbsuserv1.ConfirmEnableTfa_Reply, error) {
	return u.userRepo.ConfirmEnableTfa(ctx, req)
}

func (u *UserUsecase) DisableTfa(ctx context.Context, req *bbsuserv1.DisableTfa_Request) (*bbsuserv1.DisableTfa_Reply, error) {
	return u.userRepo.DisableTfa(ctx, req)
}

func (u *UserUsecase) GetCurrentTfa(ctx context.Context, req *bbsuserv1.GetCurrentTfa_Request) (*bbsuserv1.GetCurrentTfa_Reply, error) {
	return u.userRepo.GetCurrentTfa(ctx, req)
}
