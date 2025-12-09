-- 创建用户和数据库
CREATE USER bass WITH PASSWORD '123456';
CREATE DATABASE bass OWNER bass;

-- 初始化默认通知模板
insert into notify_notification_template (notification_type, channel, content, processors, enable, created_at, updated_at)
values (100, 2, '{{username}} 注册成功', '[]', true, now(), now())
     , (200, 2, '{{username}} 发布文章 {{title}}', '[]', true, now(), now())
;



-- 安装 zhparser 中文分词扩展
-- git clone https://github.com/amutu/zhparser.git
-- cd zhparser
-- make && sudo make install
-- 数据库中创建拓展
CREATE EXTENSION IF NOT EXISTS zhparser;

ALTER TABLE articles
    ADD COLUMN IF NOT EXISTS tsv tsvector;

UPDATE articles
SET tsv =
        setweight(to_tsvector('chinese', title), 'A') ||
        setweight(to_tsvector('chinese', content), 'B')
WHERE tsv IS NULL; -- 或更新全部

CREATE INDEX idx_article_tsv_gin ON articles USING GIN (tsv);