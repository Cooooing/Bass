# Minimal observability stack

The development monitoring stack uses Tempo, Prometheus, Loki, Grafana Alloy, and Grafana. It intentionally excludes an OTLP collector, Alertmanager, node-exporter, and kube-state-metrics.

Tempo receives OTLP/gRPC directly at `tempo.monitoring.svc:4317`. Prometheus discovers only opted-in Pods. Alloy collects only opted-in Pod logs. Grafana exposes the three data sources and the runtime dashboard.

## Workload opt-in

Apply the following Pod-template annotations to every application workload that should be observed:

```yaml
metadata:
  annotations:
    prometheus.io/scrape: "true"
    prometheus.io/path: /metrics
    prometheus.io/port: "<http-port>"
    observability.grafana.com/logs: "true"
    observability.grafana.com/service_name: "<runtime-service-name>"
```

The workload must also have the `app.kubernetes.io/name` label. Set `service_name` to the service's configured runtime name so logs can be joined to Tempo traces; it is intentionally not an indexed business identifier.

## Runtime configuration

When publishing application configuration to Consul for Kubernetes, set:

```dotenv
ENABLE_OTEL=true
OTEL_ENDPOINT=tempo.monitoring.svc:4317
OTEL_SAMPLER=0.1
ENABLE_METRICS=true
```

The `OTEL_ENDPOINT` value is environment-specific. Docker Compose uses `tempo:4317`; Kubernetes uses the monitoring Service DNS name.

## Retention and scope

Tempo, Prometheus, and Loki retain data for 90 days. Their development PVC requests are 30Gi, 20Gi, and 50Gi respectively; monitor actual ingestion and expand them for workloads that exceed this baseline. Grafana is the only externally exposed monitoring UI. Log streams retain only bounded labels: environment, namespace, pod, service name, and level. Request IDs and trace IDs remain log fields rather than labels.
