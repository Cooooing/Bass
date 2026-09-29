ARG TOOLCHAIN_IMAGE=ghcr.io/cooooing/bass-build-tools:latest

# 第一阶段构建单体应用。
FROM ${TOOLCHAIN_IMAGE} AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    PATH="/go/bin/linux_amd64:/go/bin:${PATH}"

WORKDIR /build

COPY . /build

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    make api IGNORE_ERROR=0 \
    && find app -mindepth 2 -maxdepth 2 -name Makefile -printf '%h\n' | sort | xargs --no-run-if-empty --max-procs=4 --max-args=1 sh -c 'make -C "$1" gen IGNORE_ERROR=0' _ \
    && make -C monolith gen build IGNORE_ERROR=0
RUN upx -9 --lzma /build/monolith/server -o /build/server

# 第二阶段提供最小运行环境。
FROM scratch

WORKDIR /app

COPY --from=builder /build/server /app/server
COPY --from=builder /build/monolith/configs/ /app/configs/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENV SSL_CERT_FILE=/etc/ssl/certs/ca-certificates.crt \
    TZ=Asia/Shanghai

EXPOSE 8000 9000

ENTRYPOINT ["/app/server"]
