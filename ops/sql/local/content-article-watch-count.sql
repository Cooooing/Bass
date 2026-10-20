-- 执行本迁移后再部署包含文章关注数的 Content 与 BFF 服务。
ALTER TABLE content_articles
  ADD COLUMN IF NOT EXISTS watch_count integer NOT NULL DEFAULT 0;

-- 用既有关注事实初始化聚合字段；可重复执行，后续由关注/取消关注事务原子维护。
UPDATE content_articles AS article
SET watch_count = COALESCE((
  SELECT count(*)::integer
  FROM content_article_action_records AS record
  WHERE record.article_id = article.id
    AND record.type = 'watch'
), 0);
