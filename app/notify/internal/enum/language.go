package enum

import (
	"common/pkg/enum"
	"common/proto/gen/common/enums"
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
