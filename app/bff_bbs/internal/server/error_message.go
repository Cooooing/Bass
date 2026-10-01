package server

import (
	serverutil "common/pkg/server"
	commonenums "common/proto/gen/common/enums"
	cerrors "common/proto/gen/common/errors"
)

func NewBBSErrorMessages() serverutil.ErrorMessages {
	return serverutil.ErrorMessages{
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
		cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_SCHEDULED_AT_OUT_OF_RANGE: {
			Text: map[commonenums.Language]string{
				commonenums.Language_LANGUAGE_ZH_CN: "定时发布时间需在 5 分钟后且不超过 3 个月",
				commonenums.Language_LANGUAGE_ZH_TW: "定時發佈時間需在 5 分鐘後且不超過 3 個月",
				commonenums.Language_LANGUAGE_EN:    "Scheduled publishing must be at least 5 minutes away and within 3 months",
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
		cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_REGISTER_TYPE_INVALID: {
			Text: map[commonenums.Language]string{
				commonenums.Language_LANGUAGE_ZH_CN: "注册方式无效",
				commonenums.Language_LANGUAGE_ZH_TW: "註冊方式無效",
				commonenums.Language_LANGUAGE_EN:    "Invalid registration method",
			},
		},
		cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_REGISTER_CREDENTIAL_REQUIRED: {
			Text: map[commonenums.Language]string{
				commonenums.Language_LANGUAGE_ZH_CN: "请填写完整的注册信息",
				commonenums.Language_LANGUAGE_ZH_TW: "請填寫完整的註冊資訊",
				commonenums.Language_LANGUAGE_EN:    "Complete registration details are required",
			},
		},
		cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_LOGIN_TYPE_INVALID: {
			Text: map[commonenums.Language]string{
				commonenums.Language_LANGUAGE_ZH_CN: "登录方式无效",
				commonenums.Language_LANGUAGE_ZH_TW: "登入方式無效",
				commonenums.Language_LANGUAGE_EN:    "Invalid sign-in method",
			},
		},
		cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_LOGIN_CREDENTIAL_REQUIRED: {
			Text: map[commonenums.Language]string{
				commonenums.Language_LANGUAGE_ZH_CN: "请填写完整的登录信息",
				commonenums.Language_LANGUAGE_ZH_TW: "請填寫完整的登入資訊",
				commonenums.Language_LANGUAGE_EN:    "Complete sign-in details are required",
			},
		},
		cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_REFRESH_TOKEN_REQUIRED: {
			Text: map[commonenums.Language]string{
				commonenums.Language_LANGUAGE_ZH_CN: "刷新令牌不能为空",
				commonenums.Language_LANGUAGE_ZH_TW: "刷新權杖不能為空",
				commonenums.Language_LANGUAGE_EN:    "Refresh token is required",
			},
		},
		cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_PROFILE_ARTICLE_LIST_PRIVATE: {
			Text: map[commonenums.Language]string{
				commonenums.Language_LANGUAGE_ZH_CN: "该用户未公开帖子列表",
				commonenums.Language_LANGUAGE_ZH_TW: "該使用者未公開貼文列表",
				commonenums.Language_LANGUAGE_EN:    "This user has not made posts public",
			},
		},
		cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_PROFILE_COMMENT_LIST_PRIVATE: {
			Text: map[commonenums.Language]string{
				commonenums.Language_LANGUAGE_ZH_CN: "该用户未公开回复列表",
				commonenums.Language_LANGUAGE_ZH_TW: "該使用者未公開回覆列表",
				commonenums.Language_LANGUAGE_EN:    "This user has not made comments public",
			},
		},
		cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_PROFILE_FOLLOWING_LIST_PRIVATE: {
			Text: map[commonenums.Language]string{
				commonenums.Language_LANGUAGE_ZH_CN: "该用户未公开关注列表",
				commonenums.Language_LANGUAGE_ZH_TW: "該使用者未公開關注列表",
				commonenums.Language_LANGUAGE_EN:    "This user has not made following public",
			},
		},
		cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_PROFILE_FOLLOWER_LIST_PRIVATE: {
			Text: map[commonenums.Language]string{
				commonenums.Language_LANGUAGE_ZH_CN: "该用户未公开粉丝列表",
				commonenums.Language_LANGUAGE_ZH_TW: "該使用者未公開粉絲列表",
				commonenums.Language_LANGUAGE_EN:    "This user has not made followers public",
			},
		},
		cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_PROFILE_MOONBREEZE_LIST_PRIVATE: {
			Text: map[commonenums.Language]string{
				commonenums.Language_LANGUAGE_ZH_CN: "该用户未公开清风明月列表",
				commonenums.Language_LANGUAGE_ZH_TW: "該使用者未公開清風明月列表",
				commonenums.Language_LANGUAGE_EN:    "This user has not made their moonbreezes list public",
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
		cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_PROFILE_IMAGE_INVALID: {
			Text: map[commonenums.Language]string{
				commonenums.Language_LANGUAGE_ZH_CN: "图片格式、大小或上传信息无效",
				commonenums.Language_LANGUAGE_ZH_TW: "圖片格式、大小或上傳資訊無效",
				commonenums.Language_LANGUAGE_EN:    "Invalid image format, size, or upload information",
			},
		},
		cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_PROFILE_IMAGE_TOO_LARGE: {
			Text: map[commonenums.Language]string{
				commonenums.Language_LANGUAGE_ZH_CN: "图片文件不能超过 2 MiB",
				commonenums.Language_LANGUAGE_ZH_TW: "圖片檔案不能超過 2 MiB",
				commonenums.Language_LANGUAGE_EN:    "Image files cannot exceed 2 MiB",
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
		cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_INVALID: {
			Text: map[commonenums.Language]string{
				commonenums.Language_LANGUAGE_ZH_CN: "文章信息无效",
				commonenums.Language_LANGUAGE_ZH_TW: "文章資訊無效",
				commonenums.Language_LANGUAGE_EN:    "Invalid article information",
			},
		},
		cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_MOONBREEZE_INVALID: {
			Text: map[commonenums.Language]string{
				commonenums.Language_LANGUAGE_ZH_CN: "动态内容必须是最多 512 个字符的单行文本",
				commonenums.Language_LANGUAGE_ZH_TW: "動態內容必須是最多 512 個字元的單行文字",
				commonenums.Language_LANGUAGE_EN:    "A moonbreeze must be one line with at most 512 characters",
			},
		},
		cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_MOONBREEZE_RATE_LIMITED: {
			Text: map[commonenums.Language]string{
				commonenums.Language_LANGUAGE_ZH_CN: "发布过于频繁，请 %d 秒后再试",
				commonenums.Language_LANGUAGE_ZH_TW: "發布過於頻繁，請 %d 秒後再試",
				commonenums.Language_LANGUAGE_EN:    "You are posting too frequently. Please try again in %d seconds",
			},
			Data: new(cerrors.RetryAfterErrorData),
		},
		cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_PUBLISH_AT_REQUIRED: {
			Text: map[commonenums.Language]string{
				commonenums.Language_LANGUAGE_ZH_CN: "请选择定时发布时间",
				commonenums.Language_LANGUAGE_ZH_TW: "請選擇定時發布時間",
				commonenums.Language_LANGUAGE_EN:    "Scheduled publish time is required",
			},
		},
		cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_AMOUNT_INVALID: {
			Text: map[commonenums.Language]string{
				commonenums.Language_LANGUAGE_ZH_CN: "积分数量无效",
				commonenums.Language_LANGUAGE_ZH_TW: "積分數量無效",
				commonenums.Language_LANGUAGE_EN:    "Invalid point amount",
			},
		},
		cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_INSUFFICIENT_BALANCE: {
			Text: map[commonenums.Language]string{
				commonenums.Language_LANGUAGE_ZH_CN: "积分余额不足",
				commonenums.Language_LANGUAGE_ZH_TW: "積分餘額不足",
				commonenums.Language_LANGUAGE_EN:    "Insufficient point balance",
			},
		},
		cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_IDEMPOTENCY_CONFLICT: {
			Text: map[commonenums.Language]string{
				commonenums.Language_LANGUAGE_ZH_CN: "请求正在处理中，请勿重复提交",
				commonenums.Language_LANGUAGE_ZH_TW: "請求正在處理中，請勿重複提交",
				commonenums.Language_LANGUAGE_EN:    "Request is already being processed",
			},
		},
		cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_TRANSFER_SELF_NOT_ALLOWED: {
			Text: map[commonenums.Language]string{
				commonenums.Language_LANGUAGE_ZH_CN: "不能给自己转积分",
				commonenums.Language_LANGUAGE_ZH_TW: "不能給自己轉積分",
				commonenums.Language_LANGUAGE_EN:    "You cannot transfer points to yourself",
			},
		},
		cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_RECORD_NOT_FOUND: {
			Text: map[commonenums.Language]string{
				commonenums.Language_LANGUAGE_ZH_CN: "积分流水不存在",
				commonenums.Language_LANGUAGE_ZH_TW: "積分流水不存在",
				commonenums.Language_LANGUAGE_EN:    "Point record does not exist",
			},
		},
		cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_RECORD_TYPE_INVALID: {
			Text: map[commonenums.Language]string{
				commonenums.Language_LANGUAGE_ZH_CN: "积分流水类型无效",
				commonenums.Language_LANGUAGE_ZH_TW: "積分流水類型無效",
				commonenums.Language_LANGUAGE_EN:    "Invalid point record type",
			},
		},
		cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_RECORD_QUERY_INVALID: {
			Text: map[commonenums.Language]string{
				commonenums.Language_LANGUAGE_ZH_CN: "积分流水查询条件无效",
				commonenums.Language_LANGUAGE_ZH_TW: "積分流水查詢條件無效",
				commonenums.Language_LANGUAGE_EN:    "Invalid point record query",
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
}
