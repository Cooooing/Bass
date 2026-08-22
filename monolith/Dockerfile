ARG GO_VERSION=1.26

# 第一阶段构建单体应用。
FROM golang:${GO_VERSION} AS builder

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

COPY . /build

RUN make init
RUN make gen-all IGNORE_ERROR=0
RUN make -C monolith gen IGNORE_ERROR=0
RUN make -C monolith build IGNORE_ERROR=0
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
