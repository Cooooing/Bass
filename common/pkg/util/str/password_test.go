package str

import (
	"fmt"
	"testing"
)

func TestPassword(t *testing.T) {
	password, err := HashPassword("123456")
	if err != nil {
		return
	}
	fmt.Printf("password: %s\n", password)
}
