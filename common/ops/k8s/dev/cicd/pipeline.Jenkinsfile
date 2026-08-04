pipeline {
  agent {
    kubernetes {
      cloud 'kubernetes'
      yaml """
apiVersion: v1
kind: Pod
spec:
  serviceAccountName: jenkins
  restartPolicy: Never
  imagePullSecrets:
    - name: aliyun-registry-secret
  containers:
    - name: jnlp
      image: registry.cn-hangzhou.aliyuncs.com/docker-cooooing/bass-ci-agent:20260804
      imagePullPolicy: IfNotPresent
      workingDir: /home/jenkins/agent
      resources:
        requests:
          cpu: 100m
          memory: 256Mi
        limits:
          cpu: "1"
          memory: 1024Mi
"""
    }
  }

  options {
    skipDefaultCheckout(true)
  }

  parameters {
    choice(name: 'SERVICE', choices: ['bbs', 'user', 'content', 'notify', 'im', 'platform', 'scheduler'], description: 'Service to build and deploy')
    string(name: 'BRANCH', defaultValue: 'dev', description: 'Git branch')
    booleanParam(name: 'DEPLOY_TO_TEST', defaultValue: true, description: 'Deploy image to remote test cluster after push')
  }

  environment {
    REGISTRY = 'registry.cn-hangzhou.aliyuncs.com'
    REGISTRY_NAMESPACE = 'docker-cooooing'
    TEST_NAMESPACE = 'bass-test'
    GIT_CACHE_URL = 'git://git-cache.cicd.svc.cluster.local/Bass.git'
    BUILDKIT_HOST = 'tcp://buildkitd.cicd.svc.cluster.local:1234'
    BASS_BUILDER_IMAGE = 'registry.cn-hangzhou.aliyuncs.com/docker-cooooing/base-golang:1.26-bass-20260804'
    GOPROXY = 'https://goproxy.cn,direct'
  }

  stages {
    stage('Refresh Git Cache') {
      steps {
        lock(resource: 'git-cache-Bass') {
          sh '''
            set -euo pipefail
            if ! kubectl exec -n cicd statefulset/git-cache -c git-cache -- /bin/sh /scripts/update.sh; then
              echo "[WARN] git-cache refresh failed; checking cached branch ${BRANCH}."
            fi
            git ls-remote --exit-code "${GIT_CACHE_URL}" "refs/heads/${BRANCH}" >/dev/null
          '''
        }
      }
    }

    stage('Checkout') {
      steps {
        sh '''
          set -euo pipefail
          rm -rf src
          git clone --depth 1 --branch "${BRANCH}" "${GIT_CACHE_URL}" src
          cd src
          git rev-parse --short HEAD > "${WORKSPACE}/commit.txt"
          git log -1 --oneline
        '''
      }
    }

    stage('Build and Push Image') {
      steps {
        withCredentials([usernamePassword(credentialsId: 'aliyun-acr', usernameVariable: 'ACR_USER', passwordVariable: 'ACR_PASS')]) {
          sh '''
            set -euo pipefail

            COMMIT="$(cat "${WORKSPACE}/commit.txt")"
            IMAGE="${REGISTRY}/${REGISTRY_NAMESPACE}/${SERVICE}:${BRANCH}-${BUILD_NUMBER}-${COMMIT}"
            mkdir -p "${WORKSPACE}/.docker"
            AUTH="$(printf "%s:%s" "${ACR_USER}" "${ACR_PASS}" | base64 | tr -d '\n')"
            cat > "${WORKSPACE}/.docker/config.json" <<EOF
{"auths":{"${REGISTRY}":{"auth":"${AUTH}"}}}
EOF

            export DOCKER_CONFIG="${WORKSPACE}/.docker"

            mkdir -p "${WORKSPACE}/.ci-dockerfile"

            awk '
              /^FROM golang:\$\{GO_VERSION\} AS builder$/ {
                print "ARG BASS_BUILDER_IMAGE=golang:1.26"
                print "FROM ${BASS_BUILDER_IMAGE} AS builder"
                next
              }
              /^ARG APP_NAME$/ {
                print
                print "ARG GOPROXY=https://goproxy.cn,direct"
                print "ENV GOPROXY=${GOPROXY}"
                next
              }
              { print }
            ' "${WORKSPACE}/src/common/build/docker/service.Dockerfile" \
              > "${WORKSPACE}/.ci-dockerfile/service.Dockerfile"

            echo "Using generated CI Dockerfile:"
            grep -nE 'BASS_BUILDER_IMAGE|FROM |ARG APP_NAME|GOPROXY|ENV GOPROXY' "${WORKSPACE}/.ci-dockerfile/service.Dockerfile" || true

            buildctl --addr "${BUILDKIT_HOST}" build \
              --frontend dockerfile.v0 \
              --local context="${WORKSPACE}/src" \
              --local dockerfile="${WORKSPACE}/.ci-dockerfile" \
              --opt filename=service.Dockerfile \
              --opt build-arg:APP_NAME="${SERVICE}" \
              --opt build-arg:BASS_BUILDER_IMAGE="${BASS_BUILDER_IMAGE}" \
              --opt build-arg:GOPROXY="${GOPROXY}" \
              --output "type=image,name=${IMAGE},push=true"

            echo "${IMAGE}" > "${WORKSPACE}/image.txt"
            echo "Built image: ${IMAGE}"
          '''
        }
      }
    }

    stage('Deploy to Test') {
      when {
        expression { return params.DEPLOY_TO_TEST }
      }
      steps {
        lock(resource: "bass-${params.SERVICE}") {
          withCredentials([file(credentialsId: 'vm-microk8s-kubeconfig', variable: 'KUBECONFIG_FILE')]) {
            sh '''
              set -euo pipefail
              IMAGE="$(cat "${WORKSPACE}/image.txt")"

              kubectl --kubeconfig "${KUBECONFIG_FILE}" -n "${TEST_NAMESPACE}" \
                set image deployment/bass-${SERVICE} ${SERVICE}="${IMAGE}"

              kubectl --kubeconfig "${KUBECONFIG_FILE}" -n "${TEST_NAMESPACE}" \
                rollout status deployment/bass-${SERVICE} --timeout=180s
            '''
          }
        }
      }
    }
  }
}