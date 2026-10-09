ifndef TOOLCHAIN_MK_INCLUDED
TOOLCHAIN_MK_INCLUDED := 1

# 代码生成器和构建工具的固定版本。此处的更改会重新构建 CI 工具链镜像。
.PHONY: init
init:
	@echo "[init] installing pinned development tools..."
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.11
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1
	@go install github.com/go-kratos/kratos/cmd/kratos/v3@v3.0.0-20260626125723-668db92c2c00
	@go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v3@v3.0.0-20260626125723-668db92c2c00
	@go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v3@v3.0.0-20260626125723-668db92c2c00
	@go install github.com/envoyproxy/protoc-gen-validate@v1.3.3
	@go install github.com/google/gnostic/cmd/protoc-gen-openapi@v0.7.1
	@go install github.com/google/wire/cmd/wire@v0.7.0
	@go install entgo.io/ent/cmd/ent@v0.14.5
	@go install github.com/bufbuild/buf/cmd/buf@v1.50.0

endif
