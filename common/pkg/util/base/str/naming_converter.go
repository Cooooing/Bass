package str

import (
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// ToUpperCamelCase 将字符串转换为大驼峰命名
func ToUpperCamelCase(s string) string {
	words := splitWords(s)
	for i, word := range words {
		if word == "" {
			continue
		}
		// 每个单词首字母大写，其余小写
		words[i] = cases.Title(language.English).String(strings.ToLower(word))
	}
	return strings.Join(words, "")
}

// ToLowerCamelCase 将字符串转换为小驼峰命名
func ToLowerCamelCase(s string) string {
	words := splitWords(s)
	for i, word := range words {
		if word == "" {
			continue
		}
		if i == 0 {
			// 首单词全小写
			words[i] = strings.ToLower(word)
		} else {
			// 其他单词首字母大写
			words[i] = cases.Title(language.English).String(strings.ToLower(word))
		}
	}
	return strings.Join(words, "")
}

// ToSnakeCase 将字符串转换为蛇形命名
func ToSnakeCase(s string) string {
	words := splitWords(s)
	for i, word := range words {
		if word == "" {
			continue
		}
		// 所有单词转为小写
		words[i] = strings.ToLower(word)
	}
	return strings.Join(words, "_")
}

// ToKebabCase 将字符串转换为短横线命名
func ToKebabCase(s string) string {
	words := splitWords(s)
	for i, word := range words {
		if word == "" {
			continue
		}
		// 所有单词转为小写
		words[i] = strings.ToLower(word)
	}
	return strings.Join(words, "-")
}

// splitWords 将输入字符串按下划线、短横线、空格或驼峰规则分割为单词
func splitWords(s string) []string {
	var words []string
	var word strings.Builder
	isUpper := false
	lastWasUpper := false

	for _, r := range s {
		isUpper = unicode.IsUpper(r)
		// 遇到下划线、短横线或空格，直接分割
		if r == '_' || r == '-' || unicode.IsSpace(r) {
			if word.Len() > 0 {
				words = append(words, word.String())
				word.Reset()
			}
			continue
		}
		// 驼峰规则：当前是大写且前一个不是大写时，分割
		if isUpper && !lastWasUpper && word.Len() > 0 {
			words = append(words, word.String())
			word.Reset()
		}
		word.WriteRune(r)
		lastWasUpper = isUpper
	}
	// 添加最后一个单词
	if word.Len() > 0 {
		words = append(words, word.String())
	}
	return words
}
