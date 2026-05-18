-- 预设通知模板（upsert：已存在则跳过）
-- 唯一约束: (event_type, channel, language)

-- ========== 站内信 · 中文 ==========

INSERT INTO notify_notification_template (event_type, channel, language, title, content, enable, created_at, updated_at)
VALUES ('EVENT_TYPE_USER_FOLLOW_CREATED', 'NOTIFICATION_CHANNEL_STATION', 'zh_CN', '新增关注', '{{.SenderName}} 关注了你', true, NOW(), NOW())
ON CONFLICT (event_type, channel, language) DO NOTHING;

INSERT INTO notify_notification_template (event_type, channel, language, title, content, enable, created_at, updated_at)
VALUES ('EVENT_TYPE_ARTICLE_PUBLISHED', 'NOTIFICATION_CHANNEL_STATION', 'zh_CN', '新文章发布', '{{.SenderName}} 发布了文章「{{.Title}}」', true, NOW(), NOW())
ON CONFLICT (event_type, channel, language) DO NOTHING;

INSERT INTO notify_notification_template (event_type, channel, language, title, content, enable, created_at, updated_at)
VALUES ('EVENT_TYPE_ARTICLE_LIKED', 'NOTIFICATION_CHANNEL_STATION', 'zh_CN', '文章被点赞', '{{.SenderName}} 点赞了你的文章「{{.Title}}」', true, NOW(), NOW())
ON CONFLICT (event_type, channel, language) DO NOTHING;

INSERT INTO notify_notification_template (event_type, channel, language, title, content, enable, created_at, updated_at)
VALUES ('EVENT_TYPE_ARTICLE_THANKED', 'NOTIFICATION_CHANNEL_STATION', 'zh_CN', '文章被感谢', '{{.SenderName}} 感谢了你的文章「{{.Title}}」', true, NOW(), NOW())
ON CONFLICT (event_type, channel, language) DO NOTHING;

INSERT INTO notify_notification_template (event_type, channel, language, title, content, enable, created_at, updated_at)
VALUES ('EVENT_TYPE_ARTICLE_COLLECTED', 'NOTIFICATION_CHANNEL_STATION', 'zh_CN', '文章被收藏', '{{.SenderName}} 收藏了你的文章「{{.Title}}」', true, NOW(), NOW())
ON CONFLICT (event_type, channel, language) DO NOTHING;

INSERT INTO notify_notification_template (event_type, channel, language, title, content, enable, created_at, updated_at)
VALUES ('EVENT_TYPE_ARTICLE_WATCHED', 'NOTIFICATION_CHANNEL_STATION', 'zh_CN', '文章被关注', '{{.SenderName}} 关注了你的文章「{{.Title}}」', true, NOW(), NOW())
ON CONFLICT (event_type, channel, language) DO NOTHING;

INSERT INTO notify_notification_template (event_type, channel, language, title, content, enable, created_at, updated_at)
VALUES ('EVENT_TYPE_COMMENT_PUBLISHED', 'NOTIFICATION_CHANNEL_STATION', 'zh_CN', '收到新评论', '{{.SenderName}} 评论了你的文章', true, NOW(), NOW())
ON CONFLICT (event_type, channel, language) DO NOTHING;

INSERT INTO notify_notification_template (event_type, channel, language, title, content, enable, created_at, updated_at)
VALUES ('EVENT_TYPE_COMMENT_LIKED', 'NOTIFICATION_CHANNEL_STATION', 'zh_CN', '评论被点赞', '{{.SenderName}} 点赞了你的评论', true, NOW(), NOW())
ON CONFLICT (event_type, channel, language) DO NOTHING;
