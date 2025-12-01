package tests

import (
	"bytes"
	"testing"
	"user/internal/service"
)

func TestVerify(t *testing.T) {
	verifyService := service.NewVerifyService()

	// -----------------------------
	// VerifyNickname
	// -----------------------------
	t.Run("VerifyNickname", func(t *testing.T) {
		tests := []struct {
			input string
			want  bool
		}{
			{"a", true},
			{"test", true},
			{"test123", true},
			{"中文", true},                      // 昵称允许非 ASCII
			{"😀", true},                       // emoji 也按长度统计
			{"", false},                       // 长度不足
			{string(make([]rune, 33)), false}, // 长度超过
		}

		for _, tc := range tests {
			got := verifyService.VerifyNickname(tc.input)
			if got != tc.want {
				t.Errorf("VerifyNickname(%q) = %v; want %v", tc.input, got, tc.want)
			}
		}
	})

	// -----------------------------
	// VerifyName
	// -----------------------------
	t.Run("VerifyName", func(t *testing.T) {
		tests := []struct {
			input string
			want  bool
		}{
			{"a", true},
			{"abc", true},
			{"abc-def", true},
			{"abc-def-123", true},
			{"ABC-123-xyz", true},

			// 非法情况
			{"-abc", false},
			{"abc-", false},
			{"abc--def", false},
			{"abc_def", false}, // 不允许下划线
			{"abc def", false}, // 不允许空格
			{"测试", false},      // 只允许英文与数字
			{"", false},
			{string(make([]rune, 33)), false}, // 超长
		}

		for _, tc := range tests {
			got := verifyService.VerifyName(tc.input)
			if got != tc.want {
				t.Errorf("VerifyName(%q) = %v; want %v", tc.input, got, tc.want)
			}
		}
	})

	// -----------------------------
	// VerifyPassword
	// -----------------------------
	t.Run("VerifyPassword", func(t *testing.T) {
		tests := []struct {
			input string
			want  bool
		}{
			// 合法密码
			{"a1b2", true},
			{"A1@#", true},
			{"Pass1234", true},
			{"1a" + "!@#$%^&*", true}, // 合法字符延伸
			{"a1" + string(bytes.Repeat([]byte("x"), 60)), true}, // 长度接近 64

			// 不合法：缺字母
			{"123456", false},
			{"1111!", false},

			// 不合法：缺数字
			{"abcdef", false},
			{"ABC!", false},

			// 不合法：长度不足
			{"a1!", false},
			{"a1", false},

			// 不合法：长度超过
			{string(make([]byte, 65)), false},

			// 不合法：包含不允许的字符（如中文）
			{"a1你", false},
			{"a1😀", false},
		}

		for _, tc := range tests {
			got := verifyService.VerifyPassword(tc.input)
			if got != tc.want {
				t.Errorf("VerifyPassword(%q) = %v; want %v", tc.input, got, tc.want)
			}
		}
	})
}
