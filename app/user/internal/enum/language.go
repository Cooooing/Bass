package enum

import (
	commonenum "common/pkg/enum"
	commonenums "common/proto/gen/common/enums/v1"
)

type Language string

const (
	LanguageZhCN Language = "zh-CN" // 简体中文
	LanguageZhTW Language = "zh-TW" // 繁体中文
	LanguageEn   Language = "en"    // 英文
)

var LanguageMap = commonenum.NewMapping[Language, commonenums.Language](map[Language]commonenum.Entry[Language, commonenums.Language]{
	LanguageZhCN: {Proto: commonenums.Language_LANGUAGE_ZH_CN},
	LanguageZhTW: {Proto: commonenums.Language_LANGUAGE_ZH_TW},
	LanguageEn:   {Proto: commonenums.Language_LANGUAGE_EN},
})

func (e Language) String() string {
	return string(e)
}
