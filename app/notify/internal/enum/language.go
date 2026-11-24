package enum

import (
	"common/pkg/enum"
	"common/proto/gen/common/enums/v1"
)

type Language string

const (
	LanguageZhCN Language = "zh_CN" // 简体中文
	LanguageZhTW Language = "zh_TW" // 繁体中文
	LanguageEn   Language = "en"    // 英文
)

var LanguageMap = enum.NewMapping[Language, enums.Language](map[Language]enum.Entry[Language, enums.Language]{
	LanguageZhCN: {Proto: enums.Language_LANGUAGE_ZH_CN},
	LanguageZhTW: {Proto: enums.Language_LANGUAGE_ZH_TW},
	LanguageEn:   {Proto: enums.Language_LANGUAGE_EN},
})

func (e Language) String() string {
	return string(e)
}
