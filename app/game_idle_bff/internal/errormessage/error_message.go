package errormessage

import (
	"common/pkg/apperror"
	serverutil "common/pkg/server"
	commonenums "common/proto/gen/common/enums"
	cerrors "common/proto/gen/common/errors"
	"encoding/json"
	"net/http"
)

var messages = serverutil.ErrorMessages{
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_UNKNOWN: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "服务暂时不可用",
			commonenums.Language_LANGUAGE_ZH_TW: "服務暫時不可用",
			commonenums.Language_LANGUAGE_EN:    "Service is temporarily unavailable",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "请求参数无效",
			commonenums.Language_LANGUAGE_ZH_TW: "請求參數無效",
			commonenums.Language_LANGUAGE_EN:    "Invalid request parameters",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_UNAUTHORIZED: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "请先登录",
			commonenums.Language_LANGUAGE_ZH_TW: "請先登入",
			commonenums.Language_LANGUAGE_EN:    "Sign in required",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "没有操作权限",
			commonenums.Language_LANGUAGE_ZH_TW: "沒有操作權限",
			commonenums.Language_LANGUAGE_EN:    "Permission denied",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_NOT_FOUND: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "资源不存在",
			commonenums.Language_LANGUAGE_ZH_TW: "資源不存在",
			commonenums.Language_LANGUAGE_EN:    "Resource not found",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_CONFLICT: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "当前状态不允许该操作",
			commonenums.Language_LANGUAGE_ZH_TW: "目前狀態不允許該操作",
			commonenums.Language_LANGUAGE_EN:    "The current state does not allow this operation",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_TOO_MANY_ReqS: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "操作过于频繁，请 %d 秒后再试",
			commonenums.Language_LANGUAGE_ZH_TW: "操作過於頻繁，請 %d 秒後再試",
			commonenums.Language_LANGUAGE_EN:    "Too many requests, please try again in %d seconds",
		},
		Data: new(cerrors.RetryAfterErrorData),
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INTERNAL: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "服务暂时不可用",
			commonenums.Language_LANGUAGE_ZH_TW: "服務暫時不可用",
			commonenums.Language_LANGUAGE_EN:    "Service is temporarily unavailable",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_NOT_IMPLEMENTED: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "功能暂未开放",
			commonenums.Language_LANGUAGE_ZH_TW: "功能暫未開放",
			commonenums.Language_LANGUAGE_EN:    "Feature is not available yet",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_UPSTREAM_UNAVAILABLE: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "服务暂时不可用",
			commonenums.Language_LANGUAGE_ZH_TW: "服務暫時不可用",
			commonenums.Language_LANGUAGE_EN:    "Service is temporarily unavailable",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_OPERATION_FAILED: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "操作失败",
			commonenums.Language_LANGUAGE_ZH_TW: "操作失敗",
			commonenums.Language_LANGUAGE_EN:    "Operation failed",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_INVALID_CREDENTIALS: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "账号不存在或密码错误",
			commonenums.Language_LANGUAGE_ZH_TW: "帳號不存在或密碼錯誤",
			commonenums.Language_LANGUAGE_EN:    "Account does not exist or password is incorrect",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "请先登录",
			commonenums.Language_LANGUAGE_ZH_TW: "請先登入",
			commonenums.Language_LANGUAGE_EN:    "Sign in required",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_INVALID: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "登录已失效，请重新登录",
			commonenums.Language_LANGUAGE_ZH_TW: "登入已失效，請重新登入",
			commonenums.Language_LANGUAGE_EN:    "Session expired, please sign in again",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_ACCOUNT_NOT_FOUND: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "用户不存在",
			commonenums.Language_LANGUAGE_ZH_TW: "使用者不存在",
			commonenums.Language_LANGUAGE_EN:    "User does not exist",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_ACCOUNT_NAME_TAKEN: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "账号名已被占用",
			commonenums.Language_LANGUAGE_ZH_TW: "帳號名稱已被占用",
			commonenums.Language_LANGUAGE_EN:    "Account name is already taken",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_ACCOUNT_ALREADY_EXISTS: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "账号已存在",
			commonenums.Language_LANGUAGE_ZH_TW: "帳號已存在",
			commonenums.Language_LANGUAGE_EN:    "Account already exists",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_VERIFICATION_CODE_INVALID_OR_EXPIRED: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "验证码无效或已过期",
			commonenums.Language_LANGUAGE_ZH_TW: "驗證碼無效或已過期",
			commonenums.Language_LANGUAGE_EN:    "Verification code is invalid or expired",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_VERIFICATION_CODE_SEND_TOO_FREQUENT: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "验证码发送过于频繁，请 %d 秒后再试",
			commonenums.Language_LANGUAGE_ZH_TW: "驗證碼發送過於頻繁，請 %d 秒後再試",
			commonenums.Language_LANGUAGE_EN:    "Verification code was sent too frequently, please try again in %d seconds",
		},
		Data: new(cerrors.RetryAfterErrorData),
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOTP_CODE_INVALID: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "TOTP 验证码无效",
			commonenums.Language_LANGUAGE_ZH_TW: "TOTP 驗證碼無效",
			commonenums.Language_LANGUAGE_EN:    "Invalid TOTP code",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_ACCOUNT_BANNED: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "账号已被封禁",
			commonenums.Language_LANGUAGE_ZH_TW: "帳號已被封禁",
			commonenums.Language_LANGUAGE_EN:    "Account is banned",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_ACCOUNT_CANCELLED: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "账号已注销",
			commonenums.Language_LANGUAGE_ZH_TW: "帳號已註銷",
			commonenums.Language_LANGUAGE_EN:    "Account has been cancelled",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_CHARACTER_NOT_FOUND: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "角色不存在",
			commonenums.Language_LANGUAGE_ZH_TW: "角色不存在",
			commonenums.Language_LANGUAGE_EN:    "Character does not exist",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_CHARACTER_INVALID: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "角色信息无效",
			commonenums.Language_LANGUAGE_ZH_TW: "角色資訊無效",
			commonenums.Language_LANGUAGE_EN:    "Invalid character information",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_CHARACTER_SESSION_INVALID: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "角色会话已失效，请重新进入游戏",
			commonenums.Language_LANGUAGE_ZH_TW: "角色會話已失效，請重新進入遊戲",
			commonenums.Language_LANGUAGE_EN:    "Character session expired, please enter the game again",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_CHARACTER_LIMIT_EXCEEDED: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "角色数量已达上限",
			commonenums.Language_LANGUAGE_ZH_TW: "角色數量已達上限",
			commonenums.Language_LANGUAGE_EN:    "Character limit reached",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_CHARACTER_NAME_TAKEN: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "角色名已被占用",
			commonenums.Language_LANGUAGE_ZH_TW: "角色名稱已被占用",
			commonenums.Language_LANGUAGE_EN:    "Character name is already taken",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_BACKPACK_CHANGE_INVALID: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "背包变更无效",
			commonenums.Language_LANGUAGE_ZH_TW: "背包變更無效",
			commonenums.Language_LANGUAGE_EN:    "Invalid backpack change",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_BACKPACK_INSUFFICIENT: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "背包物品不足",
			commonenums.Language_LANGUAGE_ZH_TW: "背包物品不足",
			commonenums.Language_LANGUAGE_EN:    "Not enough items in backpack",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_ITEM_INVALID: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "物品无效",
			commonenums.Language_LANGUAGE_ZH_TW: "物品無效",
			commonenums.Language_LANGUAGE_EN:    "Invalid item",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_RECIPE_INVALID: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "游戏配置异常，请稍后再试",
			commonenums.Language_LANGUAGE_ZH_TW: "遊戲配置異常，請稍後再試",
			commonenums.Language_LANGUAGE_EN:    "Game configuration error, please try again later",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_RECIPE_OUTPUT_EMPTY: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "游戏配置异常，请稍后再试",
			commonenums.Language_LANGUAGE_ZH_TW: "遊戲配置異常，請稍後再試",
			commonenums.Language_LANGUAGE_EN:    "Game configuration error, please try again later",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_ACTION_INVALID: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "行动不可用",
			commonenums.Language_LANGUAGE_ZH_TW: "行動不可用",
			commonenums.Language_LANGUAGE_EN:    "Action is unavailable",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_ACTION_QUEUE_FULL: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "行动队列已满",
			commonenums.Language_LANGUAGE_ZH_TW: "行動佇列已滿",
			commonenums.Language_LANGUAGE_EN:    "Action queue is full",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_ACTION_QUEUE_STATE_CONFLICT: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "行动队列状态已变化，请刷新后重试",
			commonenums.Language_LANGUAGE_ZH_TW: "行動佇列狀態已變化，請刷新後重試",
			commonenums.Language_LANGUAGE_EN:    "Action queue changed, please refresh and try again",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_ABILITY_LEVEL_INSUFFICIENT: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "能力等级不足",
			commonenums.Language_LANGUAGE_ZH_TW: "能力等級不足",
			commonenums.Language_LANGUAGE_EN:    "Ability level is insufficient",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_CHAT_MESSAGE_INVALID: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "聊天消息无效",
			commonenums.Language_LANGUAGE_ZH_TW: "聊天訊息無效",
			commonenums.Language_LANGUAGE_EN:    "Invalid chat message",
		},
	},
}

func ResolveHTTP(r *http.Request, code cerrors.BusinessErrorCode, data json.RawMessage) string {
	return messages.Resolve(r, code, data)
}

func ResolveError(err error) (cerrors.BusinessErrorCode, string) {
	code, ok := apperror.BusinessCode(err)
	if !ok || code == cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_SUCCESS {
		code = cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_UNKNOWN
	}
	return code, messages.Resolve(nil, code, apperror.Data(err))
}
