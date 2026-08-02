# Bass dev environment

`dev` 环境用于开发联调，namespace 为 `bass-dev`。该目录是可版本化的默认样板；部署前仍应检查 Secret、域名、镜像仓库和资源配额是否适合当前开发集群。

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
kubectl apply -f ./dev/manifests/namespace.yaml -f ./dev/manifests/secrets.yaml
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

- `manifests/secrets.yaml` 是 dev 样板，公开或共享环境部署前必须替换。
- dev 环境不部署业务服务，业务服务默认在本地运行，通过 dev 中间件联调。
- `monitoring/secrets.yaml` 包含 Grafana 管理员凭据，部署前必须替换默认值。