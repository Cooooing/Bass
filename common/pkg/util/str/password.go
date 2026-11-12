package str

import (
	"golang.org/x/crypto/bcrypt"
)

// DefaultCost 默认加密复杂度（cost） 越大越安全，但加密越慢；10 是默认推荐值
const DefaultCost = bcrypt.DefaultCost

// HashPassword 将明文密码加密为 bcrypt 哈希
func HashPassword(plain string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plain), DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// VerifyPassword 验证明文密码是否与加密密码匹配
func VerifyPassword(hashed string, plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	return err == nil
}
