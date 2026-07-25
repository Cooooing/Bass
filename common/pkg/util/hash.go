package util

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
)

// Sha256Bytes 返回字节内容的 SHA-256 十六进制摘要。
func Sha256Bytes(
	content []byte,
) string {
	sum := sha256.Sum256(content)
	return hex.EncodeToString(sum[:])
}

// Sha256File 返回文件内容的 SHA-256 十六进制摘要。
func Sha256File(
	path string,
) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return Sha256Bytes(content), nil
}
