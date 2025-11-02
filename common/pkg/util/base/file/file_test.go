package file

import (
	"strings"
	"testing"
)

func TestCcalcFileHash(t *testing.T) {
	r := strings.NewReader("hello world")
	hash, err := Hash(r, MD5)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(hash)
}
