ARG GO_VERSION=1.26

FROM golang:${GO_VERSION}

ENV PATH="/go/bin/linux_amd64:/go/bin:${PATH}"

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    git \
    protobuf-compiler \
    upx-ucl \
    && update-ca-certificates \
    && rm -rf /var/lib/apt/lists/*

COPY common/build/make/ /build/common/build/make/
RUN --mount=type=cache,target=/go/pkg/mod \
    make -f /build/common/build/make/common.mk init
