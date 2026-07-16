-- 预设验证码通知规则与模板。
-- 当前只初始化验证码模板，其他业务通知模板暂不预置。

WITH rule AS (
    INSERT INTO notify_notification_rule (event_type, channel, language, enabled, created_at, updated_at)
        VALUES ('user_email_verification_code', 'email', 'zh_CN', true, NOW(), NOW())
        ON CONFLICT (event_type, channel, language) DO UPDATE
            SET enabled = EXCLUDED.enabled, updated_at = NOW()
        RETURNING id)
INSERT
INTO notify_notification_email_template (rule_id, subject_template, body_template, content_type, created_at, updated_at)
SELECT id,
       'Bass 验证码',
       $html$<!doctype html>
<html lang="zh-CN">
<head>
	<meta charset="utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<title>Bass 验证码</title>
</head>
<body style="margin:0;padding:0;background:#f6f7f9;font-family:-apple-system,BlinkMacSystemFont,'Segoe UI',sans-serif;color:#1f2937;">
	<div style="max-width:520px;margin:0 auto;padding:32px 20px;">
		<div style="background:#ffffff;border:1px solid #e5e7eb;border-radius:8px;padding:28px;">
			<div style="font-size:18px;font-weight:600;margin-bottom:18px;">Bass 验证码</div>
			<div style="font-size:14px;line-height:1.7;color:#4b5563;margin-bottom:20px;">本次操作需要验证你的身份，请在页面中输入以下验证码。</div>
			<div style="font-size:32px;line-height:1;font-weight:700;letter-spacing:6px;color:#111827;background:#f3f4f6;border-radius:6px;padding:18px 20px;text-align:center;">{{.Code}}</div>
			<div style="font-size:13px;line-height:1.6;color:#6b7280;margin-top:20px;">验证码有效期 {{.ExpiresMinutes}} 分钟，请勿转发给其他人。</div>
		</div>
		<div style="font-size:12px;line-height:1.6;color:#9ca3af;margin-top:16px;text-align:center;">此邮件由系统自动发送，请勿直接回复。</div>
	</div>
</body>
</html>$html$,
       'text/html',
       NOW(),
       NOW()
FROM rule
ON CONFLICT (rule_id) DO UPDATE
    SET subject_template = EXCLUDED.subject_template,
        body_template    = EXCLUDED.body_template,
        content_type     = EXCLUDED.content_type,
        updated_at       = NOW();

WITH rule AS (
    INSERT INTO notify_notification_rule (event_type, channel, language, enabled, created_at, updated_at)
        VALUES ('user_phone_verification_code', 'tencent_sms', 'zh_CN', true, NOW(), NOW())
        ON CONFLICT (event_type, channel, language) DO UPDATE
            SET enabled = EXCLUDED.enabled, updated_at = NOW()
        RETURNING id)
INSERT
INTO notify_notification_tencent_sms_template (rule_id,
                                               sms_sdk_app_id,
                                               sign_name,
                                               provider_template_id,
                                               param_templates,
                                               created_at,
                                               updated_at)
SELECT id,
       '',
       '',
       '',
       '[
         "{{.Code}}"
       ]'::jsonb,
       NOW(),
       NOW()
FROM rule
ON CONFLICT (rule_id) DO UPDATE
    SET sms_sdk_app_id = EXCLUDED.sms_sdk_app_id,
        sign_name = EXCLUDED.sign_name,
        provider_template_id = EXCLUDED.provider_template_id,
        param_templates = EXCLUDED.param_templates,
        updated_at = NOW();
