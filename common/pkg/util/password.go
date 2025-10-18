package util

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"

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

// MD5Hash 返回输入字符串的 32 位小写 MD5 加密结果
func MD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// Sha256Hash 返回输入字符串的 SHA256 加密结果（16进制）
func Sha256Hash(text string) string {
	hash := sha256.New()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}
