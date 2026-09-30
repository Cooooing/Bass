# Bass dev environment

`dev` 环境用于开发联调，namespace 为 `bass-dev`。该目录是可版本化的默认样板；部署前仍应检查 Secret、域名、镜像仓库和资源配额是否适合当前开发集群。

## 首次配置 Secret

真实 Secret 不提交仓库。首次部署前创建本机文件，并填入强密码：

```bash
cp ./dev/manifests/secrets.example.yaml ./dev/manifests/secrets.local.yaml
cp ./dev/monitoring/secrets.example.yaml ./dev/monitoring/secrets.local.yaml
```

## 目录

```text
dev/
├── kustomization.yaml
├── manifests/       # Namespace、Secret、Traefik HTTP/TCP 路由
├── infra/           # Consul/PostgreSQL/Redis/NATS/MinIO Helm values
├── ingress/         # Traefik Helm values
└── monitoring/      # Prometheus、Grafana、Tempo、OTel Collector 等监控栈，单独部署
```

## 部署中间件和入口配置

```bash
kubectl apply -f ./dev/manifests/namespace.yaml -f ./dev/manifests/secrets.local.yaml
helm upgrade --install traefik traefik/traefik -n traefik --create-namespace --timeout 5m -f ./dev/ingress/traefik-values.yaml --wait
helm upgrade --install bass-consul hashicorp/consul -n bass-dev --create-namespace --timeout 5m -f ./dev/infra/consul.yaml --wait
helm upgrade --install bass-pg bitnami/postgresql -n bass-dev --create-namespace --timeout 5m -f ./dev/infra/postgres.yaml --wait
helm upgrade --install bass-redis bitnami/redis -n bass-dev --create-namespace --timeout 5m -f ./dev/infra/redis.yaml --wait
helm upgrade --install bass-nats nats/nats -n bass-dev --create-namespace --timeout 5m -f ./dev/infra/nats.yaml --wait
helm upgrade --install bass-minio minio/minio -n bass-dev --create-namespace --timeout 5m -f ./dev/infra/minio.yaml --wait
kubectl apply -k ./dev
```

## 部署监控

监控栈不挂在 `dev/kustomization.yaml` 中，避免每次 dev 应用都顺手部署。需要时单独执行：

```bash
kubectl apply -k ./dev/monitoring
```

## 验证

```bash
kubectl get pods,svc,pvc -n bass-dev
kubectl get ingressroute,ingressroutetcp -n bass-dev
helm list -n bass-dev
kubectl get pods,svc,pvc -n monitoring
```

## 注意

- `manifests/secrets.local.yaml` 与 `monitoring/secrets.local.yaml` 是本机文件，均由对应的 `*.example.yaml` 创建，绝不能提交。
- MinIO 通过 `bass-nats-secret` 的 `config.env` 向 NATS JetStream 发布资源上传事件；其中 subject 必须与 Platform 的 `ASSET_EVENT_SUBJECT` 相同。`minio-asset-event-rule` Job 会为 `assets/sha256/` 创建幂等的 `put` 事件规则。修改规则后删除该 Job 再重新 `kubectl apply -k ./dev`。
- dev 环境不部署业务服务，业务服务默认在本地运行，通过 dev 中间件联调。
