# Bass dev CI/CD

`common/ops/k8s/dev/cicd` contains the local MicroK8s CI/CD stack for Bass. The current setup runs Jenkins in the `cicd` namespace, uses a local `git-cache` mirror for source checkout, builds service images through a persistent BuildKit daemon, pushes images to Aliyun ACR, and can deploy the selected service to the remote test cluster.

## Files

```text
cicd/
|-- README.md              # This operational guide
|-- jenkins.yaml           # Namespace, Jenkins controller, RBAC, PVC, services, NodePort
|-- git-cache.yaml         # Local bare Git mirror exposed as git://git-cache:9418
|-- buildkit.yaml          # Persistent BuildKit daemon and layer cache
|-- pipeline.Jenkinsfile   # Jenkins Pipeline script
`-- secrets.example.yaml   # Credential placeholder example, do not apply directly
```

The service image build still uses `common/build/docker/service.Dockerfile`. The Pipeline does not modify that repository file. Instead, it generates a temporary Dockerfile in the Jenkins workspace so CI can use the configured builder image and `GOPROXY` while keeping the committed Dockerfile unchanged.

## Current Runtime Shape

The local `cicd` namespace currently runs:

```text
jenkins      Deployment, image jenkins/jenkins:2.568.1-lts-jdk21, 10Gi jenkins-home PVC
buildkitd    StatefulSet, image moby/buildkit:v0.24.0, 20Gi buildkit-data PVC
git-cache    StatefulSet, image alpine:3.20, 5Gi git-cache-data PVC
```

Services:

```text
jenkins             ClusterIP, ports 8080 and 50000
jenkins-nodeport    NodePort 31980 -> Jenkins HTTP
buildkitd           ClusterIP, port 1234
git-cache           ClusterIP, port 9418
```

Required Kubernetes secret in namespace `cicd`:

```text
aliyun-registry-secret    kubernetes.io/dockerconfigjson for registry.cn-hangzhou.aliyuncs.com
```

## Deploy Local CI/CD Infra

Switch to the local cluster first:

```bash
kubectl config use-context microk8s
```

Create or update the namespace and registry pull secret:

```bash
kubectl create namespace cicd --dry-run=client -o yaml | kubectl apply -f -

kubectl -n cicd create secret docker-registry aliyun-registry-secret \
  --docker-server=registry.cn-hangzhou.aliyuncs.com \
  --docker-username=<aliyun-acr-username> \
  --docker-password=<aliyun-acr-password> \
  --dry-run=client -o yaml | kubectl apply -f -
```

Apply the stack:

```bash
kubectl apply -f common/ops/k8s/dev/cicd/jenkins.yaml
kubectl apply -f common/ops/k8s/dev/cicd/git-cache.yaml
kubectl apply -f common/ops/k8s/dev/cicd/buildkit.yaml

kubectl rollout status deployment/jenkins -n cicd --timeout=5m
kubectl rollout status statefulset/git-cache -n cicd --timeout=5m
kubectl rollout status statefulset/buildkitd -n cicd --timeout=5m
```

Validate:

```bash
kubectl get pods,pvc,svc -n cicd
kubectl exec -n cicd statefulset/git-cache -c git-cache -- git -C /git-cache/Bass.git log -1 --oneline
```

## Jenkins Access

Prefer port-forward on Windows + local MicroK8s:

```bash
kubectl port-forward -n cicd svc/jenkins 31980:8080
```

Open:

```text
http://127.0.0.1:31980
```

`jenkins.yaml` also exposes NodePort `31980`, but direct Windows access depends on the local Kubernetes networking setup.

Initial admin password:

```bash
kubectl exec -n cicd deploy/jenkins -- cat /var/jenkins_home/secrets/initialAdminPassword
```

## Jenkins Plugins And Credentials

Required plugins:

```text
Kubernetes
Pipeline
Git
Credentials Binding
Lockable Resources
```

Required Jenkins credentials:

```text
aliyun-acr                  Username with password, used by buildctl image push
vm-microk8s-kubeconfig      Secret file, remote test cluster kubeconfig
```

The dynamic Jenkins agent pod also uses the Kubernetes `aliyun-registry-secret` image pull secret.

## Pipeline

Create a parameterized Pipeline job and paste `pipeline.Jenkinsfile` as the script.

Parameters:

```text
SERVICE=bbs|user|content|notify|im|platform|scheduler
BRANCH=dev
DEPLOY_TO_TEST=true
```

Pipeline flow:

```text
1. lock git-cache-Bass and refresh /git-cache/Bass.git
2. clone the requested branch from git://git-cache.cicd.svc.cluster.local/Bass.git
3. generate a temporary CI Dockerfile in the Jenkins workspace
4. build with buildctl against tcp://buildkitd.cicd.svc.cluster.local:1234
5. push registry.cn-hangzhou.aliyuncs.com/docker-cooooing/<service>:<branch>-<build>-<commit>
6. optionally lock bass-<service> and roll out the remote bass-test deployment
```

The Pipeline passes these build settings into the temporary Dockerfile:

```text
BASS_BUILDER_IMAGE=registry.cn-hangzhou.aliyuncs.com/docker-cooooing/base-golang:1.26-bass-20260804
GOPROXY=https://goproxy.cn,direct
```

Different services can build concurrently. Deployment of the same service is serialized by `lock("bass-${SERVICE}")`.

## Cache Model

```text
Git cache      git-cache StatefulSet + 5Gi PVC, updated from GitHub and reused by Jenkins checkouts
Image cache    buildkitd StatefulSet + 20Gi PVC, reused across service image builds
Workspace      dynamic Jenkins agent pods use ephemeral workspace volumes
Final images   Aliyun ACR
```

Because local MicroK8s uses `microk8s-hostpath` with `ReadWriteOnce`, cache PVCs are mounted only by their owning StatefulSet pods. Dynamic build pods do not share source PVCs.