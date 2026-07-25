package str

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func GenerateHMAC(
	message, secret string,
) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}

func VerifyHMAC(
	message, receivedHMAC, secret string,
) bool {
	expected := GenerateHMAC(message, secret)

	expectedBytes, err1 := hex.DecodeString(expected)
	receivedBytes, err2 := hex.DecodeString(receivedHMAC)
	if err1 != nil || err2 != nil {
		return false
	}

	return hmac.Equal(expectedBytes, receivedBytes)
}
