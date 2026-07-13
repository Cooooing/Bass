-- 初始化任务
INSERT INTO public.scheduler_tasks (created_at, updated_at, name, title, description, enabled, cron_spec, payload, timeout_seconds, allow_overlap, alert_enabled, version)
VALUES (now(), now(), 'noop', '空任务', '用于验证 scheduler 调度链路的空任务，不执行任何业务操作。', true, '0/5 * * * * ? *', '{"name":"test"}', 10, true, true, 0);
