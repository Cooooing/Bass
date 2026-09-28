ARG TOOLCHAIN_IMAGE=ghcr.io/cooooing/bass-build-tools:latest

# ===================== 第一阶段：构建 Go 应用 =====================
FROM ${TOOLCHAIN_IMAGE} AS builder

ARG APP_NAME

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    PATH="/go/bin/linux_amd64:/go/bin:${PATH}"

WORKDIR /build

COPY common/ /build/common/
COPY app/${APP_NAME}/ /build/app/${APP_NAME}/
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    make -C /build/app/${APP_NAME} api gen build
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
