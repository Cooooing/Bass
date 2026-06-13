ARG GO_VERSION=1.26

# ===================== 第一阶段：构建 Go 应用 =====================
FROM golang:${GO_VERSION} AS builder

ARG APP_NAME

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    PATH="/go/bin/linux_amd64:/go/bin:${PATH}"

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    protobuf-compiler \
    git \
    curl \
    unzip \
    upx-ucl \
    && update-ca-certificates \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /build

COPY common/build/make/ /build/common/build/make/
COPY common/go.mod common/go.sum /build/common/
COPY app/${APP_NAME}/go.mod app/${APP_NAME}/go.sum app/${APP_NAME}/Makefile /build/app/${APP_NAME}/
RUN cd /build/common && go mod download && cd /build/app/${APP_NAME} && go mod download && make init

COPY common/ /build/common/
COPY app/${APP_NAME}/ /build/app/${APP_NAME}/
RUN cd /build/app/${APP_NAME} && make all
RUN upx -9 --lzma /build/app/${APP_NAME}/server -o /build/server

# ===================== 第二阶段：制作轻量运行环境 =====================
FROM scratch

ARG APP_NAME

WORKDIR /app

COPY --from=builder /build/server /app/server
COPY --from=builder /build/app/${APP_NAME}/configs/ /app/configs/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENV SSL_CERT_FILE=/etc/ssl/certs/ca-certificates.crt \
    TZ=Asia/Shanghai

EXPOSE 8000 9000

ENTRYPOINT ["/app/server"]
