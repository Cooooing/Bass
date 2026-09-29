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

# --- One-time target ---

.PHONY: init
init:
	@echo "[init] installing pinned development tools..."
	@set -e; \
	for tool in \
		google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.11 \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1 \
		github.com/go-kratos/kratos/cmd/kratos/v3@v3.0.0-20260626125723-668db92c2c00 \
		github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v3@v3.0.0-20260626125723-668db92c2c00 \
		github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v3@v3.0.0-20260626125723-668db92c2c00 \
		github.com/envoyproxy/protoc-gen-validate@v1.3.3 \
		github.com/google/gnostic/cmd/protoc-gen-openapi@v0.7.1 \
		github.com/google/wire/cmd/wire@v0.7.0 \
		entgo.io/ent/cmd/ent@v0.14.5 \
		github.com/bufbuild/buf/cmd/buf@v1.50.0; do \
		for attempt in 1 2 3; do \
			if go install "$$tool"; then break; fi; \
			if [ "$$attempt" -eq 3 ]; then exit 1; fi; \
			sleep "$$attempt"; \
		done; \
	done

.PHONY: api-clean
api-clean:
	@echo "[api-clean] cleaning generated Go files..."
	@cd $(PROTO_GEN_DIR) 2>/dev/null && find . -name "*.go" -type f -delete 2>/dev/null; true
	@cd $(PROTO_GEN_DIR) 2>/dev/null && find . -type d -empty -delete 2>/dev/null; true

.PHONY: api
api: api-clean
	@echo "[api] buf generate..."
	@mkdir -p $(PROTO_GEN_DIR)
	@cd $(ROOT_DIR) && $(BUF) generate $(BUF_CONFIG_DIR) --template $(BUF_GEN_API)

endif
