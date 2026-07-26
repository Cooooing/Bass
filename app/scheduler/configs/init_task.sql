-- 初始化任务
INSERT INTO public.scheduler_tasks (created_at, updated_at, name, title, description, enabled, cron_spec, payload, timeout_seconds, allow_overlap, alert_enabled, version)
VALUES (now(), now(), 'noop', '空任务', '用于验证 scheduler 调度链路的空任务，不执行任何业务操作。', true, '0/5 * * * * ? *', '{"name":"test"}', 10, true, true, 0);

INSERT INTO public.scheduler_tasks (created_at, updated_at, name, title, description, enabled, cron_spec, payload, timeout_seconds, allow_overlap, alert_enabled, version)
VALUES (now(), now(), 'user.outbox_publish_batch', 'User outbox publish batch', 'Call user.OutboxService.PublishBatch to publish pending outbox events.', true, '0/10 * * * * ? *', '{"limit":1000,"publish_timeout_seconds":300,"max_retry":10}', 30, false, false, 0);
