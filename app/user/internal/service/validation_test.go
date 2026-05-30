package service

import "testing"

func TestAuthServiceValidateRegister(t *testing.T) {
	service := NewAuthService(nil, nil)
	nickname := "Tester"
	if err := service.validateRegister("user-1", &nickname, "abc123"); err != nil {
		t.Fatalf("valid register should pass: %v", err)
	}
	if err := service.validateRegister("u", &nickname, "abc123"); err == nil {
		t.Fatalf("short name should fail")
	}
	if err := service.validateRegister("user-1", &nickname, "abc"); err == nil {
		t.Fatalf("weak password should fail")
	}
	if err := service.validateRegister("user-1", new("1234"), "abc123"); err == nil {
		t.Fatalf("numeric nickname should fail")
	}
}

func TestAccountServiceValidateProfileUpdate(t *testing.T) {
	service := NewAccountService(nil)
	if err := service.validateProfileUpdate(nil, new(""), nil, nil); err != nil {
		t.Fatalf("empty nickname should clear profile: %v", err)
	}
	if err := service.validateProfileUpdate(nil, new("1"), nil, nil); err == nil {
		t.Fatalf("short nickname should fail")
	}
	if err := service.validateProfileUpdate(nil, nil, nil, nil); err != nil {
		t.Fatalf("empty profile update should pass: %v", err)
	}
}
