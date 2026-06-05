-- 预设站内信通知规则与模板。
-- 外部通道模板包含供应商配置或密钥，不在初始化 SQL 中预置。

WITH rule AS (
	INSERT INTO notify_notification_rule (event_type, channel, language, enabled, created_at, updated_at)
	VALUES ('user_register', 'station', 'zh_CN', true, NOW(), NOW())
	ON CONFLICT (event_type, channel, language) DO UPDATE
	SET enabled = EXCLUDED.enabled, updated_at = NOW()
	RETURNING id
)
INSERT INTO notify_notification_station_template (rule_id, title_template, content_template, created_at, updated_at)
SELECT id, '欢迎加入 Bass', '你好，账号 {{.User.ID}} 已创建成功。', NOW(), NOW()
FROM rule
ON CONFLICT (rule_id) DO UPDATE
SET title_template = EXCLUDED.title_template,
	content_template = EXCLUDED.content_template,
	updated_at = NOW();

WITH rule AS (
	INSERT INTO notify_notification_rule (event_type, channel, language, enabled, created_at, updated_at)
	VALUES ('user_follow', 'station', 'zh_CN', true, NOW(), NOW())
	ON CONFLICT (event_type, channel, language) DO UPDATE
	SET enabled = EXCLUDED.enabled, updated_at = NOW()
	RETURNING id
)
INSERT INTO notify_notification_station_template (rule_id, title_template, content_template, created_at, updated_at)
SELECT id, '新增关注', '{{.Follower.Nickname}} 关注了你。', NOW(), NOW()
FROM rule
ON CONFLICT (rule_id) DO UPDATE
SET title_template = EXCLUDED.title_template,
	content_template = EXCLUDED.content_template,
	updated_at = NOW();

WITH rule AS (
	INSERT INTO notify_notification_rule (event_type, channel, language, enabled, created_at, updated_at)
	VALUES ('content_article_publish', 'station', 'zh_CN', true, NOW(), NOW())
	ON CONFLICT (event_type, channel, language) DO UPDATE
	SET enabled = EXCLUDED.enabled, updated_at = NOW()
	RETURNING id
)
INSERT INTO notify_notification_station_template (rule_id, title_template, content_template, created_at, updated_at)
SELECT id, '关注作者发布了新文章', '{{.Article.Author.Nickname}} 发布了文章《{{.Article.Title}}》。', NOW(), NOW()
FROM rule
ON CONFLICT (rule_id) DO UPDATE
SET title_template = EXCLUDED.title_template,
	content_template = EXCLUDED.content_template,
	updated_at = NOW();

WITH rule AS (
	INSERT INTO notify_notification_rule (event_type, channel, language, enabled, created_at, updated_at)
	VALUES ('content_article_like', 'station', 'zh_CN', true, NOW(), NOW())
	ON CONFLICT (event_type, channel, language) DO UPDATE
	SET enabled = EXCLUDED.enabled, updated_at = NOW()
	RETURNING id
)
INSERT INTO notify_notification_station_template (rule_id, title_template, content_template, created_at, updated_at)
SELECT id, '文章收到点赞', '{{.Actor.Nickname}} 点赞了你的文章《{{.Article.Title}}》。', NOW(), NOW()
FROM rule
ON CONFLICT (rule_id) DO UPDATE
SET title_template = EXCLUDED.title_template,
	content_template = EXCLUDED.content_template,
	updated_at = NOW();

WITH rule AS (
	INSERT INTO notify_notification_rule (event_type, channel, language, enabled, created_at, updated_at)
	VALUES ('content_article_thank', 'station', 'zh_CN', true, NOW(), NOW())
	ON CONFLICT (event_type, channel, language) DO UPDATE
	SET enabled = EXCLUDED.enabled, updated_at = NOW()
	RETURNING id
)
INSERT INTO notify_notification_station_template (rule_id, title_template, content_template, created_at, updated_at)
SELECT id, '文章收到感谢', '{{.Actor.Nickname}} 感谢了你的文章《{{.Article.Title}}》。', NOW(), NOW()
FROM rule
ON CONFLICT (rule_id) DO UPDATE
SET title_template = EXCLUDED.title_template,
	content_template = EXCLUDED.content_template,
	updated_at = NOW();

WITH rule AS (
	INSERT INTO notify_notification_rule (event_type, channel, language, enabled, created_at, updated_at)
	VALUES ('content_article_collect', 'station', 'zh_CN', true, NOW(), NOW())
	ON CONFLICT (event_type, channel, language) DO UPDATE
	SET enabled = EXCLUDED.enabled, updated_at = NOW()
	RETURNING id
)
INSERT INTO notify_notification_station_template (rule_id, title_template, content_template, created_at, updated_at)
SELECT id, '文章被收藏', '{{.Actor.Nickname}} 收藏了你的文章《{{.Article.Title}}》。', NOW(), NOW()
FROM rule
ON CONFLICT (rule_id) DO UPDATE
SET title_template = EXCLUDED.title_template,
	content_template = EXCLUDED.content_template,
	updated_at = NOW();

WITH rule AS (
	INSERT INTO notify_notification_rule (event_type, channel, language, enabled, created_at, updated_at)
	VALUES ('content_article_watch', 'station', 'zh_CN', true, NOW(), NOW())
	ON CONFLICT (event_type, channel, language) DO UPDATE
	SET enabled = EXCLUDED.enabled, updated_at = NOW()
	RETURNING id
)
INSERT INTO notify_notification_station_template (rule_id, title_template, content_template, created_at, updated_at)
SELECT id, '文章被关注', '{{.Actor.Nickname}} 关注了你的文章《{{.Article.Title}}》。', NOW(), NOW()
FROM rule
ON CONFLICT (rule_id) DO UPDATE
SET title_template = EXCLUDED.title_template,
	content_template = EXCLUDED.content_template,
	updated_at = NOW();

WITH rule AS (
	INSERT INTO notify_notification_rule (event_type, channel, language, enabled, created_at, updated_at)
	VALUES ('content_comment_publish', 'station', 'zh_CN', true, NOW(), NOW())
	ON CONFLICT (event_type, channel, language) DO UPDATE
	SET enabled = EXCLUDED.enabled, updated_at = NOW()
	RETURNING id
)
INSERT INTO notify_notification_station_template (rule_id, title_template, content_template, created_at, updated_at)
SELECT id, '收到新评论', '{{.Comment.User.Nickname}} 评论了文章《{{.Comment.Article.Title}}》。', NOW(), NOW()
FROM rule
ON CONFLICT (rule_id) DO UPDATE
SET title_template = EXCLUDED.title_template,
	content_template = EXCLUDED.content_template,
	updated_at = NOW();

WITH rule AS (
	INSERT INTO notify_notification_rule (event_type, channel, language, enabled, created_at, updated_at)
	VALUES ('content_comment_like', 'station', 'zh_CN', true, NOW(), NOW())
	ON CONFLICT (event_type, channel, language) DO UPDATE
	SET enabled = EXCLUDED.enabled, updated_at = NOW()
	RETURNING id
)
INSERT INTO notify_notification_station_template (rule_id, title_template, content_template, created_at, updated_at)
SELECT id, '评论收到点赞', '{{.Actor.Nickname}} 点赞了你的评论。', NOW(), NOW()
FROM rule
ON CONFLICT (rule_id) DO UPDATE
SET title_template = EXCLUDED.title_template,
	content_template = EXCLUDED.content_template,
	updated_at = NOW();
