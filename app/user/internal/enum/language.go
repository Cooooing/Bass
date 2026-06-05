package enum

import (
	commonenums "common/api/gen/common/enums"
	commonenum "common/pkg/enum"
)

type Language string

const (
	LanguageZhCN Language = "zh-CN"
	LanguageZhTW Language = "zh-TW"
	LanguageEn   Language = "en"
)

var LanguageMap = commonenum.NewMapping[Language, commonenums.Language](map[Language]commonenum.Entry[Language, commonenums.Language]{
	LanguageZhCN: {Proto: commonenums.Language_LANGUAGE_ZH_CN},
	LanguageZhTW: {Proto: commonenums.Language_LANGUAGE_ZH_TW},
	LanguageEn:   {Proto: commonenums.Language_LANGUAGE_EN},
})
