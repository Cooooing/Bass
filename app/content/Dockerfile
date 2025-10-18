# ===================== 第一阶段：构建 Go 应用 =====================
FROM golang:1.24 AS builder

# 设置工作目录
WORKDIR /build

COPY . .

# 下载 common 模块依赖
WORKDIR /build/common
RUN go mod download

# 下载 content 模块依赖
WORKDIR /build/app/content
RUN go mod download

# 安装依赖
RUN apt-get update && apt-get install -y \
    protobuf-compiler \
    git \
    curl \
    unzip \
    && rm -rf /var/lib/apt/lists/*

# 执行构建
RUN make all

# ===================== 第二阶段：制作轻量运行环境 =====================
FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates netbase && \
    rm -rf /var/lib/apt/lists/*

# 设置工作目录
WORKDIR /app

# 拷贝编译好的可执行文件（根据 Makefile，输出在项目根目录）
COPY --from=builder /build/content /app/server
# 拷贝配置文件
COPY --from=builder /build/app/content/configs/ /app/configs/

# 设置时区（可选，根据需要调整）
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# 暴露端口
EXPOSE 8002 9002

# 启动命令
CMD ["/app/server"]