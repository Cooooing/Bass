package enum

import (
	"common/pkg/apperror"
	commonenum "common/pkg/enum"
	cerrors "common/proto/gen/common/errors"
	v1 "common/proto/gen/economy/v1/enum"
)

// EconomyRecordType 表示积分流水类型
type EconomyRecordType string

const (
	EconomyRecordTypeSignInReward             EconomyRecordType = "sign_in_reward"              // EconomyRecordTypeSignInReward 表示签到奖励
	EconomyRecordTypeArticleThankReward       EconomyRecordType = "article_thank_reward"        // EconomyRecordTypeArticleThankReward 表示文章感谢奖励
	EconomyRecordTypeCommentThankReward       EconomyRecordType = "comment_thank_reward"        // EconomyRecordTypeCommentThankReward 表示评论感谢奖励
	EconomyRecordTypeArticleRewardOut         EconomyRecordType = "article_reward_out"          // EconomyRecordTypeArticleRewardOut 表示文章打赏支出
	EconomyRecordTypeArticleRewardIn          EconomyRecordType = "article_reward_in"           // EconomyRecordTypeArticleRewardIn 表示文章打赏收入
	EconomyRecordTypeAdminAdd                 EconomyRecordType = "admin_add"                   // EconomyRecordTypeAdminAdd 表示后台增加
	EconomyRecordTypeAdminDeduct              EconomyRecordType = "admin_deduct"                // EconomyRecordTypeAdminDeduct 表示后台扣减
	EconomyRecordTypeGlobalChatTopicExpense   EconomyRecordType = "global_chat_topic_expense"   // EconomyRecordTypeGlobalChatTopicExpense 表示设置全局聊天室话题支出
	EconomyRecordTypeGlobalChatBarrageExpense EconomyRecordType = "global_chat_barrage_expense" // EconomyRecordTypeGlobalChatBarrageExpense 表示发送全局聊天室弹幕支出
)

// EconomyRecordTypeMap 维护流水类型映射
var EconomyRecordTypeMap = commonenum.NewMapping[EconomyRecordType, v1.EconomyRecordType](map[EconomyRecordType]commonenum.Entry[EconomyRecordType, v1.EconomyRecordType]{
	EconomyRecordTypeSignInReward:             {Proto: v1.EconomyRecordType_ECONOMY_RECORD_TYPE_SIGN_IN_REWARD},
	EconomyRecordTypeArticleThankReward:       {Proto: v1.EconomyRecordType_ECONOMY_RECORD_TYPE_ARTICLE_THANK_REWARD},
	EconomyRecordTypeCommentThankReward:       {Proto: v1.EconomyRecordType_ECONOMY_RECORD_TYPE_COMMENT_THANK_REWARD},
	EconomyRecordTypeArticleRewardOut:         {Proto: v1.EconomyRecordType_ECONOMY_RECORD_TYPE_ARTICLE_REWARD_OUT},
	EconomyRecordTypeArticleRewardIn:          {Proto: v1.EconomyRecordType_ECONOMY_RECORD_TYPE_ARTICLE_REWARD_IN},
	EconomyRecordTypeAdminAdd:                 {Proto: v1.EconomyRecordType_ECONOMY_RECORD_TYPE_ADMIN_ADD},
	EconomyRecordTypeAdminDeduct:              {Proto: v1.EconomyRecordType_ECONOMY_RECORD_TYPE_ADMIN_DEDUCT},
	EconomyRecordTypeGlobalChatTopicExpense:   {Proto: v1.EconomyRecordType_ECONOMY_RECORD_TYPE_GLOBAL_CHAT_TOPIC_EXPENSE},
	EconomyRecordTypeGlobalChatBarrageExpense: {Proto: v1.EconomyRecordType_ECONOMY_RECORD_TYPE_GLOBAL_CHAT_BARRAGE_EXPENSE},
})

func (e EconomyRecordType) String() string {
	return string(e)
}

func (e EconomyRecordType) Direction() (EconomyRecordDirection, error) {
	switch e {
	case EconomyRecordTypeSignInReward, EconomyRecordTypeArticleThankReward, EconomyRecordTypeCommentThankReward, EconomyRecordTypeArticleRewardIn, EconomyRecordTypeAdminAdd:
		return EconomyRecordDirectionIncome, nil
	case EconomyRecordTypeArticleRewardOut, EconomyRecordTypeAdminDeduct, EconomyRecordTypeGlobalChatTopicExpense, EconomyRecordTypeGlobalChatBarrageExpense:
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
