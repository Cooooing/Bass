ifndef COMMON_MK_INCLUDED
COMMON_MK_INCLUDED := 1

# --- Variables ---
COMMON_MAKE_DIR := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))
ROOT_DIR := $(abspath $(COMMON_MAKE_DIR)/../../..)
COMMON_DIR := $(ROOT_DIR)/common

# Proto paths for shared API contracts.
PROTO_DIR := $(COMMON_DIR)/proto/app
COMMON_PROTO_DIR := $(COMMON_DIR)/proto/app/common
PROTO_GEN_DIR := $(COMMON_DIR)/proto/gen
BUF_DIR := $(COMMON_DIR)/proto/buf
BUF_CONFIG_DIR := $(PROTO_DIR)
BUF_CONFIG := $(BUF_CONFIG_DIR)/buf.yaml
BUF_GEN_API := $(BUF_DIR)/gen.api.yaml
BUF_GEN_CONFIG := $(BUF_DIR)/gen.config.yaml
BUF_GEN_OPENAPI := $(BUF_DIR)/gen.openapi.yaml
BUF ?= buf

# --- One-time targets ---

.PHONY: init
init:
	@echo "[init] installing development tools..."
	@cd $(COMMON_DIR) && \
	go mod tidy && go mod download
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v3@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v3@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v3@latest
	go install github.com/envoyproxy/protoc-gen-validate@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/google/wire/cmd/wire@latest
	go install entgo.io/ent/cmd/ent@latest
	go install github.com/bufbuild/buf/cmd/buf@latest

.PHONY: api-clean
api-clean:
	@echo "[api-clean] cleaning generated Go files..."
	@cd $(PROTO_GEN_DIR) 2>/dev/null && find . -name "*.go" -type f -delete 2>/dev/null; true
	@cd $(PROTO_GEN_DIR) 2>/dev/null && find . -type d -empty -delete 2>/dev/null; true

.PHONY: api-lint
api-lint:
	@echo "[api-lint] buf lint..."
	@cd $(ROOT_DIR) && $(BUF) lint $(BUF_CONFIG_DIR)

.PHONY: api-dep
api-dep:
	@echo "[api-dep] buf dep update..."
	@cd $(ROOT_DIR) && $(BUF) dep update $(BUF_CONFIG_DIR)

.PHONY: api-format
api-format:
	@echo "[api-format] format shared API proto..."
	@cd $(ROOT_DIR) && $(BUF) format -w $(BUF_CONFIG_DIR)
	@echo "[api-format] gofmt common Go files..."
	@find $(COMMON_DIR) -type f -name "*.go" -not -path "$(PROTO_GEN_DIR)/*" -exec gofmt -w {} +

.PHONY: api
api: api-clean
	@echo "[api] buf generate..."
	@mkdir -p $(PROTO_GEN_DIR)
	@cd $(ROOT_DIR) && $(BUF) generate $(BUF_CONFIG_DIR) --template $(BUF_GEN_API)

endif
