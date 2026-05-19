-- 预设通知模板（upsert：已存在则跳过）
-- 唯一约束: (event_type, channel, language)
-- event_type/channel/language 使用内部枚举值（非 proto 名称）

-- ====================================================================
-- 站内信 · 中文
-- ====================================================================

INSERT INTO notify_notification_template (event_type, channel, language, title, content, enable, created_at, updated_at)
VALUES ('user_follow_created', 'station', 'zh_CN', '新增关注', '{{.ActorName}} 关注了你', true, NOW(), NOW())
ON CONFLICT (event_type, channel, language) DO NOTHING;

INSERT INTO notify_notification_template (event_type, channel, language, title, content, enable, created_at, updated_at)
VALUES ('user_follow_deleted', 'station', 'zh_CN', '取消关注', '{{.ActorName}} 取消了关注你', true, NOW(), NOW())
ON CONFLICT (event_type, channel, language) DO NOTHING;

INSERT INTO notify_notification_template (event_type, channel, language, title, content, enable, created_at, updated_at)
VALUES ('article_published', 'station', 'zh_CN', '新文章发布', '{{.ActorName}} 发布了文章「{{.Title}}」', true, NOW(), NOW())
ON CONFLICT (event_type, channel, language) DO NOTHING;

INSERT INTO notify_notification_template (event_type, channel, language, title, content, enable, created_at, updated_at)
VALUES ('article_liked', 'station', 'zh_CN', '文章被点赞', '{{.ActorName}} 点赞了你的文章「{{.Title}}」', true, NOW(), NOW())
ON CONFLICT (event_type, channel, language) DO NOTHING;

INSERT INTO notify_notification_template (event_type, channel, language, title, content, enable, created_at, updated_at)
VALUES ('article_thanked', 'station', 'zh_CN', '文章被感谢', '{{.ActorName}} 感谢了你的文章「{{.Title}}」', true, NOW(), NOW())
ON CONFLICT (event_type, channel, language) DO NOTHING;

INSERT INTO notify_notification_template (event_type, channel, language, title, content, enable, created_at, updated_at)
VALUES ('article_collected', 'station', 'zh_CN', '文章被收藏', '{{.ActorName}} 收藏了你的文章「{{.Title}}」', true, NOW(), NOW())
ON CONFLICT (event_type, channel, language) DO NOTHING;

INSERT INTO notify_notification_template (event_type, channel, language, title, content, enable, created_at, updated_at)
VALUES ('article_watched', 'station', 'zh_CN', '文章被关注', '{{.ActorName}} 关注了你的文章「{{.Title}}」', true, NOW(), NOW())
ON CONFLICT (event_type, channel, language) DO NOTHING;

INSERT INTO notify_notification_template (event_type, channel, language, title, content, enable, created_at, updated_at)
VALUES ('comment_published', 'station', 'zh_CN', '收到新评论', '{{.ActorName}} 评论了你的文章「{{.Title}}」', true, NOW(), NOW())
ON CONFLICT (event_type, channel, language) DO NOTHING;

INSERT INTO notify_notification_template (event_type, channel, language, title, content, enable, created_at, updated_at)
VALUES ('comment_liked', 'station', 'zh_CN', '评论被点赞', '{{.ActorName}} 点赞了你的评论', true, NOW(), NOW())
ON CONFLICT (event_type, channel, language) DO NOTHING;
