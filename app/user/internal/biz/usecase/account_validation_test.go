package usecase

import (
	"testing"
	"user/internal/biz/repo"
)

func TestAccountValidationUsecaseValidateRegister(t *testing.T) {
	validation := NewAccountValidationUsecase()
	nickname := "用户1"

	if err := validation.ValidateRegister("user-1", &nickname, "abc123"); err != nil {
		t.Fatalf("valid register rejected: %v", err)
	}
	if err := validation.ValidateRegister("u", &nickname, "abc123"); err == nil {
		t.Fatal("short name accepted")
	}
	if err := validation.ValidateRegister("user-1", &nickname, "abc"); err == nil {
		t.Fatal("weak password accepted")
	}
	numericNickname := "1234"
	if err := validation.ValidateRegister("user-1", &numericNickname, "abc123"); err == nil {
		t.Fatal("numeric nickname accepted")
	}
}

func TestAccountValidationUsecaseValidateProfileUpdate(t *testing.T) {
	validation := NewAccountValidationUsecase()
	empty := ""
	invalidNickname := "1"

	if err := validation.ValidateProfileUpdate(&repo.AccountProfilePatch{
		Nickname: repo.NewStringPatch(&empty),
	}); err != nil {
		t.Fatalf("empty nickname clear rejected: %v", err)
	}
	if err := validation.ValidateProfileUpdate(&repo.AccountProfilePatch{
		Nickname: repo.NewStringPatch(&invalidNickname),
	}); err == nil {
		t.Fatal("invalid nickname accepted")
	}
	if err := validation.ValidateProfileUpdate(&repo.AccountProfilePatch{}); err != nil {
		t.Fatalf("empty patch rejected: %v", err)
	}
}
