package util

import (
	"hash/crc32"
	"os"
	"strings"

	"github.com/bytedance/gopkg/lang/fastrand"
	"github.com/sony/sonyflake/v2"
)

func NewSonyflake() (*sonyflake.Sonyflake, error) {
	return sonyflake.New(sonyflake.Settings{
		MachineID: func() (int, error) {
			h, err := os.Hostname()
			if err != nil {
				return 0, err
			}
			sum := crc32.ChecksumIEEE([]byte(h))
			u := sum % 65536
			return int(u), nil
		},
	})
}

func RandStr(sf *sonyflake.Sonyflake, length int, useLower, useUpper, useDigit, useUnderscore bool) string {
	// 构建字符集
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
		return ""
	}

	base := int64(len(charset))

	var sb strings.Builder
	var n int64
	for sb.Len() < length {
		if n <= 0 {
			n, _ = sf.NextID()
		}
		sb.WriteByte(charset[n%base])
		n /= base
	}

	return sb.String()
}

func RandomInRange(min, max int) int {
	if min > max {
		min, max = max, min // 处理min>max的情况
	} else if min == max {
		return min
	}
	return fastrand.Intn(max-min) + min
}
