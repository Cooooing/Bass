# K8s Monitoring Stack

监控栈包含 Prometheus、Grafana、Tempo、OpenTelemetry Collector、kube-state-metrics 和 node-exporter，部署到 `monitoring` 命名空间。

## 部署

在 `common/ops/k8s` 目录执行：

```bash
kubectl apply -k ./monitoring
```

## 状态检查

```bash
kubectl get pods,svc,pvc -n monitoring
```

## 访问

```bash
kubectl get svc prometheus -n monitoring -o jsonpath='{.spec.clusterIP}:{.spec.ports[0].port}'
kubectl get svc grafana -n monitoring -o jsonpath='{.spec.clusterIP}:{.spec.ports[0].port}'
kubectl get svc tempo -n monitoring -o jsonpath='{.spec.clusterIP}:{.spec.ports[0].port}'
kubectl get svc otel-collector -n monitoring -o jsonpath='{.spec.clusterIP}:{.spec.ports[0].port}'
kubectl get svc -n monitoring -o wide
```

Grafana 默认账号来自 `monitoring/secrets.yaml` 中的 `monitoring-grafana-secret`。

## 组件

| 组件                      | 镜像                                             | 用途           |
|-------------------------|------------------------------------------------|--------------|
| Prometheus              | `prom/prometheus:v3.4.1`                       | 指标采集与存储      |
| Grafana                 | `grafana/grafana:11.6.0`                       | 指标和链路可视化     |
| Tempo                   | `grafana/tempo:2.8.1`                          | 链路数据存储       |
| OpenTelemetry Collector | `otel/opentelemetry-collector-contrib:0.129.0` | OTLP 链路接收和转发 |
| kube-state-metrics      | `registry.k8s.io/kube-state-metrics:v2.15.0`   | K8s 对象指标     |
| node-exporter           | `prom/node-exporter:v1.9.1`                    | 主机指标         |

## 业务服务接入

业务服务开启链路上报时，将 `trace.endpoint` 指向：

```yaml
otel-collector.monitoring.svc:4317
```

业务 Pod 或 Service 添加以下注解后可被 Prometheus 自动发现：

```yaml
prometheus.io/scrape: "true"
prometheus.io/path: /metrics
prometheus.io/port: "9100"
```

`prometheus.io/port` 使用各服务 `server.http.port` 的实际端口。

## 自定义

- PVC 默认使用 `microk8s-hostpath`，其他集群需要替换为实际 StorageClass。
- Prometheus 默认保留 15 天数据，配置在 `prometheus-deployment.yaml`。
- Tempo 默认保留 7 天链路块，配置在 `tempo.yaml`。