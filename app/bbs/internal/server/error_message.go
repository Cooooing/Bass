package server

import (
	serverutil "common/pkg/server"
	commonenums "common/proto/gen/common/enums"
	cerrors "common/proto/gen/common/errors"
)

var bbsErrorMessages = serverutil.ErrorMessages{
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
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_ACCOUNT_NAME_INVALID: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "账号名格式不正确",
			commonenums.Language_LANGUAGE_ZH_TW: "帳號名格式不正確",
			commonenums.Language_LANGUAGE_EN:    "Invalid account name format",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_PASSWORD_INVALID: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "密码格式不正确",
			commonenums.Language_LANGUAGE_ZH_TW: "密碼格式不正確",
			commonenums.Language_LANGUAGE_EN:    "Invalid password format",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_EMAIL_INVALID: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "邮箱格式不正确",
			commonenums.Language_LANGUAGE_ZH_TW: "信箱格式不正確",
			commonenums.Language_LANGUAGE_EN:    "Invalid email format",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_PHONE_INVALID: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "手机号格式不正确",
			commonenums.Language_LANGUAGE_ZH_TW: "手機號格式不正確",
			commonenums.Language_LANGUAGE_EN:    "Invalid phone number format",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_NICKNAME_INVALID: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "昵称格式不正确",
			commonenums.Language_LANGUAGE_ZH_TW: "暱稱格式不正確",
			commonenums.Language_LANGUAGE_EN:    "Invalid nickname format",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_VERIFICATION_CODE_INVALID: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "验证码格式不正确",
			commonenums.Language_LANGUAGE_ZH_TW: "驗證碼格式不正確",
			commonenums.Language_LANGUAGE_EN:    "Invalid verification code format",
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
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOTP_ALREADY_ENABLED: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "TOTP 已启用",
			commonenums.Language_LANGUAGE_ZH_TW: "TOTP 已啟用",
			commonenums.Language_LANGUAGE_EN:    "TOTP is already enabled",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOTP_ALREADY_DISABLED: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "TOTP 已关闭",
			commonenums.Language_LANGUAGE_ZH_TW: "TOTP 已關閉",
			commonenums.Language_LANGUAGE_EN:    "TOTP is already disabled",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOTP_CODE_INVALID: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "TOTP 验证码无效",
			commonenums.Language_LANGUAGE_ZH_TW: "TOTP 驗證碼無效",
			commonenums.Language_LANGUAGE_EN:    "Invalid TOTP code",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_RELATION_INVALID: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "关系操作无效",
			commonenums.Language_LANGUAGE_ZH_TW: "關係操作無效",
			commonenums.Language_LANGUAGE_EN:    "Invalid relation operation",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_RELATION_ALREADY_EXISTS: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "关系已存在",
			commonenums.Language_LANGUAGE_ZH_TW: "關係已存在",
			commonenums.Language_LANGUAGE_EN:    "Relation already exists",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_RELATION_NOT_FOUND: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "关系不存在",
			commonenums.Language_LANGUAGE_ZH_TW: "關係不存在",
			commonenums.Language_LANGUAGE_EN:    "Relation does not exist",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_SELF_OPERATION_NOT_ALLOWED: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "不能对自己执行该操作",
			commonenums.Language_LANGUAGE_ZH_TW: "不能對自己執行該操作",
			commonenums.Language_LANGUAGE_EN:    "This operation cannot target yourself",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_PREFERENCE_INVALID: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "偏好设置无效",
			commonenums.Language_LANGUAGE_ZH_TW: "偏好設定無效",
			commonenums.Language_LANGUAGE_EN:    "Invalid preference settings",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_PROFILE_INVALID: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "资料信息无效",
			commonenums.Language_LANGUAGE_ZH_TW: "資料資訊無效",
			commonenums.Language_LANGUAGE_EN:    "Invalid profile information",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_NOT_FOUND: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "文章不存在",
			commonenums.Language_LANGUAGE_ZH_TW: "文章不存在",
			commonenums.Language_LANGUAGE_EN:    "Article does not exist",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_TAG_NOT_FOUND: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "标签不存在",
			commonenums.Language_LANGUAGE_ZH_TW: "標籤不存在",
			commonenums.Language_LANGUAGE_EN:    "Tag does not exist",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_DOMAIN_NOT_FOUND: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "领域不存在",
			commonenums.Language_LANGUAGE_ZH_TW: "領域不存在",
			commonenums.Language_LANGUAGE_EN:    "Domain does not exist",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_COMMENT_NOT_FOUND: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "评论不存在",
			commonenums.Language_LANGUAGE_ZH_TW: "評論不存在",
			commonenums.Language_LANGUAGE_EN:    "Comment does not exist",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_ACTION_RECORD_NOT_FOUND: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "文章操作记录不存在",
			commonenums.Language_LANGUAGE_ZH_TW: "文章操作記錄不存在",
			commonenums.Language_LANGUAGE_EN:    "Article action record does not exist",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_STATUS: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "文章状态无效",
			commonenums.Language_LANGUAGE_ZH_TW: "文章狀態無效",
			commonenums.Language_LANGUAGE_EN:    "Invalid article status",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_TYPE: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "文章类型无效",
			commonenums.Language_LANGUAGE_ZH_TW: "文章類型無效",
			commonenums.Language_LANGUAGE_EN:    "Invalid article type",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_ACTION: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "文章操作无效",
			commonenums.Language_LANGUAGE_ZH_TW: "文章操作無效",
			commonenums.Language_LANGUAGE_EN:    "Invalid article action",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_COMMENT_STATUS: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "评论状态无效",
			commonenums.Language_LANGUAGE_ZH_TW: "評論狀態無效",
			commonenums.Language_LANGUAGE_EN:    "Invalid comment status",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_COMMENT_ACTION: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "评论操作无效",
			commonenums.Language_LANGUAGE_ZH_TW: "評論操作無效",
			commonenums.Language_LANGUAGE_EN:    "Invalid comment action",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_NOT_COMMENTABLE: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "文章不允许评论",
			commonenums.Language_LANGUAGE_ZH_TW: "文章不允許評論",
			commonenums.Language_LANGUAGE_EN:    "This article does not allow comments",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "文章当前状态不允许该操作",
			commonenums.Language_LANGUAGE_ZH_TW: "文章目前狀態不允許該操作",
			commonenums.Language_LANGUAGE_EN:    "The current article status does not allow this operation",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_REWARD_NOT_IMPLEMENTED: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "文章打赏暂未开放",
			commonenums.Language_LANGUAGE_ZH_TW: "文章打賞暫未開放",
			commonenums.Language_LANGUAGE_EN:    "Article reward is not available yet",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_TAG_INVALID: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "标签信息无效",
			commonenums.Language_LANGUAGE_ZH_TW: "標籤資訊無效",
			commonenums.Language_LANGUAGE_EN:    "Invalid tag information",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_DOMAIN_INVALID: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "领域信息无效",
			commonenums.Language_LANGUAGE_ZH_TW: "領域資訊無效",
			commonenums.Language_LANGUAGE_EN:    "Invalid domain information",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_COMMENT_INVALID: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "评论信息无效",
			commonenums.Language_LANGUAGE_ZH_TW: "評論資訊無效",
			commonenums.Language_LANGUAGE_EN:    "Invalid comment information",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_NOTIFY_RATE_LIMIT_Req_INVALID: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "限流请求无效",
			commonenums.Language_LANGUAGE_ZH_TW: "限流請求無效",
			commonenums.Language_LANGUAGE_EN:    "Invalid rate limit request",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_NOTIFY_CHANNEL_INVALID: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "通知渠道无效",
			commonenums.Language_LANGUAGE_ZH_TW: "通知渠道無效",
			commonenums.Language_LANGUAGE_EN:    "Invalid notification channel",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_NOTIFY_RECIPIENT_INVALID: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "通知接收人无效",
			commonenums.Language_LANGUAGE_ZH_TW: "通知接收人無效",
			commonenums.Language_LANGUAGE_EN:    "Invalid notification recipient",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_IM_CHAT_SESSION_NOT_FOUND: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "会话不存在",
			commonenums.Language_LANGUAGE_ZH_TW: "會話不存在",
			commonenums.Language_LANGUAGE_EN:    "Chat session does not exist",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_IM_CHAT_GROUP_NOT_FOUND: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "群组不存在",
			commonenums.Language_LANGUAGE_ZH_TW: "群組不存在",
			commonenums.Language_LANGUAGE_EN:    "Chat group does not exist",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_IM_CHAT_GROUP_STATUS_INVALID: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "群组状态无效",
			commonenums.Language_LANGUAGE_ZH_TW: "群組狀態無效",
			commonenums.Language_LANGUAGE_EN:    "Invalid chat group status",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_IM_CHAT_SESSION_INVALID: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "会话信息无效",
			commonenums.Language_LANGUAGE_ZH_TW: "會話資訊無效",
			commonenums.Language_LANGUAGE_EN:    "Invalid chat session information",
		},
	},
	cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_BBS_CALLBACK_SIGNATURE_INVALID: {
		Text: map[commonenums.Language]string{
			commonenums.Language_LANGUAGE_ZH_CN: "回调签名无效",
			commonenums.Language_LANGUAGE_ZH_TW: "回調簽名無效",
			commonenums.Language_LANGUAGE_EN:    "Invalid callback signature",
		},
	},
}
