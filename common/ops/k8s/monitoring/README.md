# K8s Monitoring Stack

监控栈包含 Prometheus、Grafana、kube-state-metrics 和 node-exporter，部署到 `monitoring` 命名空间。

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
kubectl get svc -n monitoring -o wide
```

Grafana 默认账号来自 `monitoring/secrets.yaml` 中的 `monitoring-grafana-secret`。

## 组件

| 组件 | 镜像 | 用途 |
|------|------|------|
| Prometheus | `prom/prometheus:v3.4.1` | 指标采集与存储 |
| Grafana | `grafana/grafana:11.6.0` | 监控可视化 |
| kube-state-metrics | `registry.k8s.io/kube-state-metrics:v2.15.0` | K8s 对象指标 |
| node-exporter | `prom/node-exporter:v1.9.1` | 主机指标 |

## 自定义

- PVC 使用本地默认 `microk8s-hostpath`；其他集群需要替换为实际 StorageClass。
- 业务 Pod 或 Service 添加 `prometheus.io/scrape: "true"` 后可被 Prometheus 自动发现。
- Prometheus 默认保留 15 天数据，配置在 `prometheus-deployment.yaml`。
