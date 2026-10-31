package enum

import (
	"common/pkg/apperror"
	commonenum "common/pkg/enum"
	commonenums "common/proto/gen/common/enums"
	cerrors "common/proto/gen/common/errors"
	v1 "common/proto/gen/economy/v1/enum"
)

// EconomyRecordType 表示积分流水类型
type EconomyRecordType string

const (
	EconomyRecordTypeSignInReward              EconomyRecordType = "sign_in_reward"               // EconomyRecordTypeSignInReward 表示签到奖励
	EconomyRecordTypeArticleThankReward        EconomyRecordType = "article_thank_reward"         // EconomyRecordTypeArticleThankReward 表示文章感谢奖励
	EconomyRecordTypeCommentThankReward        EconomyRecordType = "comment_thank_reward"         // EconomyRecordTypeCommentThankReward 表示评论感谢奖励
	EconomyRecordTypeArticleRewardOut          EconomyRecordType = "article_reward_out"           // EconomyRecordTypeArticleRewardOut 表示文章打赏支出
	EconomyRecordTypeArticleRewardIn           EconomyRecordType = "article_reward_in"            // EconomyRecordTypeArticleRewardIn 表示文章打赏收入
	EconomyRecordTypeAdminAdd                  EconomyRecordType = "admin_add"                    // EconomyRecordTypeAdminAdd 表示后台增加
	EconomyRecordTypeAdminDeduct               EconomyRecordType = "admin_deduct"                 // EconomyRecordTypeAdminDeduct 表示后台扣减
	EconomyRecordTypeGlobalChatTopicExpense    EconomyRecordType = "global_chat_topic_expense"    // EconomyRecordTypeGlobalChatTopicExpense 表示设置全局聊天室话题支出
	EconomyRecordTypeGlobalChatBarrageExpense  EconomyRecordType = "global_chat_barrage_expense"  // EconomyRecordTypeGlobalChatBarrageExpense 表示发送全局聊天室弹幕支出
	EconomyRecordTypeArticleThankExpense       EconomyRecordType = "article_thank_expense"        // EconomyRecordTypeArticleThankExpense 表示文章感谢支出
	EconomyRecordTypeCommentThankExpense       EconomyRecordType = "comment_thank_expense"        // EconomyRecordTypeCommentThankExpense 表示评论感谢支出
	EconomyRecordTypeArticleRewardSetupExpense EconomyRecordType = "article_reward_setup_expense" // EconomyRecordTypeArticleRewardSetupExpense 表示启用文章打赏区支出
	EconomyRecordTypeCheckinStreakReward       EconomyRecordType = "checkin_streak_reward"        // EconomyRecordTypeCheckinStreakReward 表示连续签到奖励
	EconomyRecordTypeLivenessReward            EconomyRecordType = "liveness_reward"              // EconomyRecordTypeLivenessReward 表示活跃度奖励
	EconomyRecordTypeAccountTransferOut        EconomyRecordType = "account_transfer_out"         // EconomyRecordTypeAccountTransferOut 表示用户转账支出
	EconomyRecordTypeAccountTransferIn         EconomyRecordType = "account_transfer_in"          // EconomyRecordTypeAccountTransferIn 表示用户转账收入
)

// EconomyRecordTypeMap 维护流水类型映射
var EconomyRecordTypeMap = commonenum.NewMapping[EconomyRecordType, v1.EconomyRecordType](map[EconomyRecordType]commonenum.Entry[EconomyRecordType, v1.EconomyRecordType]{
	EconomyRecordTypeSignInReward:              {Proto: v1.EconomyRecordType_ECONOMY_RECORD_TYPE_SIGN_IN_REWARD},
	EconomyRecordTypeArticleThankReward:        {Proto: v1.EconomyRecordType_ECONOMY_RECORD_TYPE_ARTICLE_THANK_REWARD},
	EconomyRecordTypeCommentThankReward:        {Proto: v1.EconomyRecordType_ECONOMY_RECORD_TYPE_COMMENT_THANK_REWARD},
	EconomyRecordTypeArticleRewardOut:          {Proto: v1.EconomyRecordType_ECONOMY_RECORD_TYPE_ARTICLE_REWARD_OUT},
	EconomyRecordTypeArticleRewardIn:           {Proto: v1.EconomyRecordType_ECONOMY_RECORD_TYPE_ARTICLE_REWARD_IN},
	EconomyRecordTypeAdminAdd:                  {Proto: v1.EconomyRecordType_ECONOMY_RECORD_TYPE_ADMIN_ADD},
	EconomyRecordTypeAdminDeduct:               {Proto: v1.EconomyRecordType_ECONOMY_RECORD_TYPE_ADMIN_DEDUCT},
	EconomyRecordTypeGlobalChatTopicExpense:    {Proto: v1.EconomyRecordType_ECONOMY_RECORD_TYPE_GLOBAL_CHAT_TOPIC_EXPENSE},
	EconomyRecordTypeGlobalChatBarrageExpense:  {Proto: v1.EconomyRecordType_ECONOMY_RECORD_TYPE_GLOBAL_CHAT_BARRAGE_EXPENSE},
	EconomyRecordTypeArticleThankExpense:       {Proto: v1.EconomyRecordType_ECONOMY_RECORD_TYPE_ARTICLE_THANK_EXPENSE},
	EconomyRecordTypeCommentThankExpense:       {Proto: v1.EconomyRecordType_ECONOMY_RECORD_TYPE_COMMENT_THANK_EXPENSE},
	EconomyRecordTypeArticleRewardSetupExpense: {Proto: v1.EconomyRecordType_ECONOMY_RECORD_TYPE_ARTICLE_REWARD_SETUP_EXPENSE},
	EconomyRecordTypeCheckinStreakReward:       {Proto: v1.EconomyRecordType_ECONOMY_RECORD_TYPE_CHECKIN_STREAK_REWARD},
	EconomyRecordTypeLivenessReward:            {Proto: v1.EconomyRecordType_ECONOMY_RECORD_TYPE_LIVENESS_REWARD},
	EconomyRecordTypeAccountTransferOut:        {Proto: v1.EconomyRecordType_ECONOMY_RECORD_TYPE_ACCOUNT_TRANSFER_OUT},
	EconomyRecordTypeAccountTransferIn:         {Proto: v1.EconomyRecordType_ECONOMY_RECORD_TYPE_ACCOUNT_TRANSFER_IN},
})

func (e EconomyRecordType) String() string {
	return string(e)
}

func (e EconomyRecordType) Direction() (EconomyRecordDirection, error) {
	switch e {
	case EconomyRecordTypeSignInReward, EconomyRecordTypeArticleThankReward, EconomyRecordTypeCommentThankReward, EconomyRecordTypeArticleRewardIn, EconomyRecordTypeAdminAdd, EconomyRecordTypeCheckinStreakReward, EconomyRecordTypeLivenessReward, EconomyRecordTypeAccountTransferIn:
		return EconomyRecordDirectionIncome, nil
	case EconomyRecordTypeArticleRewardOut, EconomyRecordTypeAdminDeduct, EconomyRecordTypeGlobalChatTopicExpense, EconomyRecordTypeGlobalChatBarrageExpense, EconomyRecordTypeArticleThankExpense, EconomyRecordTypeCommentThankExpense, EconomyRecordTypeArticleRewardSetupExpense, EconomyRecordTypeAccountTransferOut:
		return EconomyRecordDirectionExpense, nil
	default:
		return "", apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_RECORD_TYPE_INVALID)
	}
}

func (e EconomyRecordType) IsIncome() bool {
	direction, err := e.Direction()
	return err == nil && direction == EconomyRecordDirectionIncome
}

func (e EconomyRecordType) IsExpense() bool {
	direction, err := e.Direction()
	return err == nil && direction == EconomyRecordDirectionExpense
}

// Label 返回积分流水类型在指定语言下的展示文案。
func (e EconomyRecordType) Label(language commonenums.Language) string {
	labels := map[commonenums.Language]string{}
	switch e {
	case EconomyRecordTypeSignInReward:
		labels = map[commonenums.Language]string{commonenums.Language_LANGUAGE_ZH_CN: "签到奖励", commonenums.Language_LANGUAGE_ZH_TW: "簽到獎勵", commonenums.Language_LANGUAGE_EN: "Check-in reward"}
	case EconomyRecordTypeCheckinStreakReward:
		labels = map[commonenums.Language]string{commonenums.Language_LANGUAGE_ZH_CN: "连续签到奖励", commonenums.Language_LANGUAGE_ZH_TW: "連續簽到獎勵", commonenums.Language_LANGUAGE_EN: "Check-in streak reward"}
	case EconomyRecordTypeLivenessReward:
		labels = map[commonenums.Language]string{commonenums.Language_LANGUAGE_ZH_CN: "活跃度奖励", commonenums.Language_LANGUAGE_ZH_TW: "活躍度獎勵", commonenums.Language_LANGUAGE_EN: "Liveness reward"}
	case EconomyRecordTypeArticleThankReward:
		labels = map[commonenums.Language]string{commonenums.Language_LANGUAGE_ZH_CN: "感谢帖子收入", commonenums.Language_LANGUAGE_ZH_TW: "感謝文章收入", commonenums.Language_LANGUAGE_EN: "Article thanks received"}
	case EconomyRecordTypeArticleThankExpense:
		labels = map[commonenums.Language]string{commonenums.Language_LANGUAGE_ZH_CN: "感谢帖子", commonenums.Language_LANGUAGE_ZH_TW: "感謝文章", commonenums.Language_LANGUAGE_EN: "Article thanks"}
	case EconomyRecordTypeCommentThankReward:
		labels = map[commonenums.Language]string{commonenums.Language_LANGUAGE_ZH_CN: "感谢回帖收入", commonenums.Language_LANGUAGE_ZH_TW: "感謝回覆收入", commonenums.Language_LANGUAGE_EN: "Comment thanks received"}
	case EconomyRecordTypeCommentThankExpense:
		labels = map[commonenums.Language]string{commonenums.Language_LANGUAGE_ZH_CN: "感谢回帖", commonenums.Language_LANGUAGE_ZH_TW: "感謝回覆", commonenums.Language_LANGUAGE_EN: "Comment thanks"}
	case EconomyRecordTypeArticleRewardIn:
		labels = map[commonenums.Language]string{commonenums.Language_LANGUAGE_ZH_CN: "帖子打赏收入", commonenums.Language_LANGUAGE_ZH_TW: "文章打賞收入", commonenums.Language_LANGUAGE_EN: "Article reward received"}
	case EconomyRecordTypeArticleRewardOut:
		labels = map[commonenums.Language]string{commonenums.Language_LANGUAGE_ZH_CN: "帖子打赏", commonenums.Language_LANGUAGE_ZH_TW: "文章打賞", commonenums.Language_LANGUAGE_EN: "Article reward"}
	case EconomyRecordTypeArticleRewardSetupExpense:
		labels = map[commonenums.Language]string{commonenums.Language_LANGUAGE_ZH_CN: "启用打赏区", commonenums.Language_LANGUAGE_ZH_TW: "啟用打賞區", commonenums.Language_LANGUAGE_EN: "Enable reward section"}
	case EconomyRecordTypeGlobalChatTopicExpense:
		labels = map[commonenums.Language]string{commonenums.Language_LANGUAGE_ZH_CN: "设置聊天室话题", commonenums.Language_LANGUAGE_ZH_TW: "設定聊天室話題", commonenums.Language_LANGUAGE_EN: "Set chat topic"}
	case EconomyRecordTypeGlobalChatBarrageExpense:
		labels = map[commonenums.Language]string{commonenums.Language_LANGUAGE_ZH_CN: "发送弹幕", commonenums.Language_LANGUAGE_ZH_TW: "發送彈幕", commonenums.Language_LANGUAGE_EN: "Send barrage"}
	case EconomyRecordTypeAdminAdd:
		labels = map[commonenums.Language]string{commonenums.Language_LANGUAGE_ZH_CN: "管理员增加", commonenums.Language_LANGUAGE_ZH_TW: "管理員增加", commonenums.Language_LANGUAGE_EN: "Administrator credit"}
	case EconomyRecordTypeAdminDeduct:
		labels = map[commonenums.Language]string{commonenums.Language_LANGUAGE_ZH_CN: "管理员扣减", commonenums.Language_LANGUAGE_ZH_TW: "管理員扣減", commonenums.Language_LANGUAGE_EN: "Administrator deduction"}
	case EconomyRecordTypeAccountTransferOut:
		labels = map[commonenums.Language]string{commonenums.Language_LANGUAGE_ZH_CN: "转账", commonenums.Language_LANGUAGE_ZH_TW: "轉帳", commonenums.Language_LANGUAGE_EN: "Points transfer"}
	case EconomyRecordTypeAccountTransferIn:
		labels = map[commonenums.Language]string{commonenums.Language_LANGUAGE_ZH_CN: "收到转账", commonenums.Language_LANGUAGE_ZH_TW: "收到轉帳", commonenums.Language_LANGUAGE_EN: "Points transfer received"}
	}
	if labels[language] != "" {
		return labels[language]
	}
	return labels[commonenums.Language_LANGUAGE_ZH_CN]
}
