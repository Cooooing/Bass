package str

import (
	"crypto/rand"
	"math/big"
	"strings"
)

// RandStr returns a cryptographically secure random string using the selected ASCII character groups.
func RandStr(length int, useLower, useUpper, useDigit, useUnderscore bool) (string, error) {
	var charset string
	if useDigit {
		charset += "0123456789"
	}
	if useLower {
		charset += "abcdefghijklmnopqrstuvwxyz"
	}
	if useUpper {
		charset += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if useUnderscore {
		charset += "_"
	}
	if charset == "" || length <= 0 {
		return "", nil
	}

	bound := big.NewInt(int64(len(charset)))
	var sb strings.Builder
	sb.Grow(length)
	for range length {
		index, err := rand.Int(rand.Reader, bound)
		if err != nil {
			return "", err
		}
		sb.WriteByte(charset[index.Int64()])
	}
	return sb.String(), nil
}
