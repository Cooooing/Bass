package util

import (
	"github.com/bytedance/gopkg/lang/fastrand"
	"github.com/sony/sonyflake/v2"
	"strings"
)

var SF *sonyflake.Sonyflake

func init() {
	var err error
	SF, err = sonyflake.New(sonyflake.Settings{})
	if err != nil {
		panic(err)
	}
}

func RandStr(length int, useLower, useUpper, useDigit, useUnderscore bool) string {
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
			n, _ = SF.NextID()
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
