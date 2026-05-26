-- 预设通知模板（upsert：已存在则跳过）
-- 唯一约束: (event_type, channel, language)
-- event_type/channel/language 使用内部枚举值（非 proto 名称）

-- ====================================================================
-- 站内信 · 中文
-- ====================================================================

INSERT INTO notify_notification_template (event_type, channel, language, title, content, enable, created_at, updated_at)
VALUES ('user_follow', 'station', 'zh_CN', '新增关注', '{{.SenderName}} 关注了你', true, NOW(), NOW())
ON CONFLICT (event_type, channel, language) DO NOTHING;

INSERT INTO notify_notification_template (event_type, channel, language, title, content, enable, created_at, updated_at)
VALUES ('user_register', 'email', 'zh_CN', '注册验证码', '你的注册验证码是 {{.Code}}，有效期 {{.ExpiresSeconds}} 秒。', true, NOW(), NOW())
ON CONFLICT (event_type, channel, language) DO NOTHING;

INSERT INTO notify_notification_template (event_type, channel, language, title, content, enable, created_at, updated_at)
VALUES ('user_register', 'sms', 'zh_CN', '注册验证码', '你的注册验证码是 {{.Code}}，有效期 {{.ExpiresSeconds}} 秒。', true, NOW(), NOW())
ON CONFLICT (event_type, channel, language) DO NOTHING;

INSERT INTO notify_notification_template (event_type, channel, language, title, content, enable, created_at, updated_at)
VALUES ('content_article_publish', 'station', 'zh_CN', '新文章发布', '{{.SenderName}} 发布了文章「{{.Title}}」', true, NOW(), NOW())
ON CONFLICT (event_type, channel, language) DO NOTHING;

INSERT INTO notify_notification_template (event_type, channel, language, title, content, enable, created_at, updated_at)
VALUES ('content_article_like', 'station', 'zh_CN', '文章被点赞', '{{.SenderName}} 点赞了你的文章「{{.Title}}」', true, NOW(), NOW())
ON CONFLICT (event_type, channel, language) DO NOTHING;

INSERT INTO notify_notification_template (event_type, channel, language, title, content, enable, created_at, updated_at)
VALUES ('content_article_thank', 'station', 'zh_CN', '文章被感谢', '{{.SenderName}} 感谢了你的文章「{{.Title}}」', true, NOW(), NOW())
ON CONFLICT (event_type, channel, language) DO NOTHING;

INSERT INTO notify_notification_template (event_type, channel, language, title, content, enable, created_at, updated_at)
VALUES ('content_article_collect', 'station', 'zh_CN', '文章被收藏', '{{.SenderName}} 收藏了你的文章「{{.Title}}」', true, NOW(), NOW())
ON CONFLICT (event_type, channel, language) DO NOTHING;

INSERT INTO notify_notification_template (event_type, channel, language, title, content, enable, created_at, updated_at)
VALUES ('content_article_watch', 'station', 'zh_CN', '文章被关注', '{{.SenderName}} 关注了你的文章「{{.Title}}」', true, NOW(), NOW())
ON CONFLICT (event_type, channel, language) DO NOTHING;

INSERT INTO notify_notification_template (event_type, channel, language, title, content, enable, created_at, updated_at)
VALUES ('content_comment_publish', 'station', 'zh_CN', '收到新评论', '{{.SenderName}} 评论了你的文章「{{.Title}}」', true, NOW(), NOW())
ON CONFLICT (event_type, channel, language) DO NOTHING;

INSERT INTO notify_notification_template (event_type, channel, language, title, content, enable, created_at, updated_at)
VALUES ('content_comment_like', 'station', 'zh_CN', '评论被点赞', '{{.SenderName}} 点赞了你的评论', true, NOW(), NOW())
ON CONFLICT (event_type, channel, language) DO NOTHING;
