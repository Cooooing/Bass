package enum

import (
	"common/api/gen/common/enums"
	"common/pkg/enum"
)

type Language string

const (
	LanguageZhCN Language = "zh_CN"
	LanguageZhTW Language = "zh_TW"
	LanguageEn   Language = "en"
)

var LanguageMap = enum.NewMapping[Language, enums.Language](map[Language]enum.Entry[Language, enums.Language]{
	LanguageZhCN: {Proto: enums.Language_LANGUAGE_ZH_CN},
	LanguageZhTW: {Proto: enums.Language_LANGUAGE_ZH_TW},
	LanguageEn:   {Proto: enums.Language_LANGUAGE_EN},
})
