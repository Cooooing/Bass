-- DTM v1.19.0 PostgreSQL 初始化脚本
-- 使用 dtm schema 隔离 DTM 内部表

CREATE SCHEMA IF NOT EXISTS dtm;
SET search_path TO dtm;

CREATE SEQUENCE IF NOT EXISTS trans_global_seq;
CREATE TABLE IF NOT EXISTS trans_global (
  id bigint NOT NULL DEFAULT NEXTVAL('trans_global_seq'),
  gid varchar(128) NOT NULL,
  trans_type varchar(45) NOT NULL,
  status varchar(45) NOT NULL,
  query_prepared varchar(1024) NOT NULL,
  protocol varchar(45) NOT NULL,
  create_time timestamp(0) with time zone DEFAULT NULL,
  update_time timestamp(0) with time zone DEFAULT NULL,
  finish_time timestamp(0) with time zone DEFAULT NULL,
  rollback_time timestamp(0) with time zone DEFAULT NULL,
  options varchar(1024) DEFAULT '',
  custom_data varchar(1024) DEFAULT '',
  next_cron_interval int DEFAULT NULL,
  next_cron_time timestamp(0) with time zone DEFAULT NULL,
  owner varchar(128) NOT NULL DEFAULT '',
  ext_data text,
  result varchar(1024) DEFAULT '',
  rollback_reason varchar(1024) DEFAULT '',
  PRIMARY KEY (id),
  CONSTRAINT gid UNIQUE (gid)
);
CREATE INDEX IF NOT EXISTS owner ON trans_global(owner);
CREATE INDEX IF NOT EXISTS status_next_cron_time ON trans_global(status, next_cron_time);

CREATE SEQUENCE IF NOT EXISTS trans_branch_op_seq;
CREATE TABLE IF NOT EXISTS trans_branch_op (
  id bigint NOT NULL DEFAULT NEXTVAL('trans_branch_op_seq'),
  gid varchar(128) NOT NULL,
  url varchar(1024) NOT NULL,
  data text,
  bin_data bytea,
  branch_id varchar(128) NOT NULL,
  op varchar(45) NOT NULL,
  status varchar(45) NOT NULL,
  finish_time timestamp(0) with time zone DEFAULT NULL,
  rollback_time timestamp(0) with time zone DEFAULT NULL,
  create_time timestamp(0) with time zone DEFAULT NULL,
  update_time timestamp(0) with time zone DEFAULT NULL,
  PRIMARY KEY (id),
  CONSTRAINT gid_branch_uniq UNIQUE (gid, branch_id, op)
);

CREATE SEQUENCE IF NOT EXISTS kv_seq;
CREATE TABLE IF NOT EXISTS kv (
  id bigint NOT NULL DEFAULT NEXTVAL('kv_seq'),
  cat varchar(45) NOT NULL,
  k varchar(128) NOT NULL,
  v text,
  version bigint DEFAULT 1,
  create_time timestamp(0) with time zone DEFAULT NULL,
  update_time timestamp(0) with time zone DEFAULT NULL,
  PRIMARY KEY (id),
  CONSTRAINT uniq_k UNIQUE(cat, k)
);
