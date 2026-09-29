ifndef COMMON_MK_INCLUDED
COMMON_MK_INCLUDED := 1

# --- Variables ---
COMMON_MAKE_DIR := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))
ROOT_DIR := $(abspath $(COMMON_MAKE_DIR)/../..)
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

include $(COMMON_MAKE_DIR)/toolchain.mk

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
