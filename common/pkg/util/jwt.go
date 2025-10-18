package util

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// TokenGenerator 负责生成和解析 JWT Token
type TokenGenerator[T any] struct {
	secret string
	expire time.Duration
}

// Claims 自定义载荷结构
type Claims[T any] struct {
	Data T `json:"data"`
	jwt.RegisteredClaims
}

// NewTokenGenerator 创建一个新的 Token 生成器
func NewTokenGenerator[T any](secret string, expire time.Duration) *TokenGenerator[T] {
	return &TokenGenerator[T]{
		secret: secret,
		expire: expire,
	}
}

// Generate 生成 JWT Token
func (g *TokenGenerator[T]) Generate(data T) (string, error) {
	claims := &Claims[T]{
		Data: data,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(g.expire)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()), // Token 立即可用
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(g.secret))
}

// Parse 解析 Token 并返回业务数据 T
// 如果 Token 无效、过期或解析失败，返回错误
func (g *TokenGenerator[T]) Parse(tokenStr string) (T, error) {
	var zero T // 泛型的零值，用于返回错误时

	parsedToken, err := jwt.ParseWithClaims(tokenStr, &Claims[T]{}, func(token *jwt.Token) (interface{}, error) {
		// 确保签名方法正确
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(g.secret), nil
	})
	if err != nil {
		return zero, err
	}

	// 断言 Claims
	claims, ok := parsedToken.Claims.(*Claims[T])
	if !ok {
		return zero, errors.New("invalid token claims")
	}
	if !parsedToken.Valid {
		return zero, errors.New("invalid token")
	}
	return claims.Data, nil
}
