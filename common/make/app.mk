# global variables
SERVER := $(notdir $(CURDIR))

ROOT_DIR := $(abspath $(dir $(lastword $(MAKEFILE_LIST)))/../../)
COMMON_DIR := $(ROOT_DIR)/common
APP_DIR := $(ROOT_DIR)/app/$(SERVER)
INTERNAL_DIR := $(APP_DIR)/internal
PROTO_DIR := $(COMMON_DIR)/api/app
PROTO_GEN_DIR := $(COMMON_DIR)/api/gen
PROTO_THIRD_PARTY_DIR := $(COMMON_DIR)/api/third_party

# Find proto files
APP_PROTO_FILES := $(shell find $(INTERNAL_DIR) -type f -name "*.proto" | sort)

# Include common makefile
include $(ROOT_DIR)/common/make/common.mk

IGNORE_ERROR ?= 0 # 0: exit when error, 1: ignore error

# go mod tidy
.PHONY: tidy
tidy:
	@echo "go mod tidy..."
	@cd $(APP_DIR) && \
	go mod tidy || \
	{ echo "[ERROR] go mod tidy failed"; [ "$(IGNORE_ERROR)" = "1" ]; }

# clean config proto
.PHONY: config-clean
config-clean:
	@echo "clean internal proto files products..."
	@cd $(APP_DIR) && find . -type f \( -name "*.pb.go" -o -name "*.pb.validate.go" -o -name "*.pb.gw.go" \) -delete 2>/dev/null || true

# generate config proto
.PHONY: config
config: config-clean
	@echo "generating internal proto files..."
	@cd $(APP_DIR) && \
	protoc -I $(INTERNAL_DIR) -I $(PROTO_THIRD_PARTY_DIR) \
	       --go_out=paths=source_relative:$(INTERNAL_DIR) \
	       $(APP_PROTO_FILES) || \
	{ echo "[ERROR] generate config proto failed"; [ "$(IGNORE_ERROR)" = "1" ]; }

# clean wire products
.PHONY: wire-clean
wire-clean:
	@echo "clean go wire products..."
	@cd $(APP_DIR)/cmd && find . -type f -name "wire_gen.go" -delete 2>/dev/null || true

# generate wire
.PHONY: wire
wire: wire-clean
	@echo "generating go wire..."
	@cd $(APP_DIR)/cmd && \
	wire || { echo "[ERROR] generate wire failed"; [ "$(IGNORE_ERROR)" = "1" ]; }

# clean ent products
.PHONY: ent-clean
ent-clean:
	@echo "clean go ent products..."
	@cd $(APP_DIR) && \
	rm -rf $(APP_DIR)/internal/data/ent/gen 2>/dev/null || true

# generate ent
.PHONY: ent
ent: ent-clean
	@echo "generating go ent..."
	@cd $(APP_DIR) && \
	ent generate --target=$(APP_DIR)/internal/data/ent/gen ./internal/data/ent/schema || \
	{ echo "[ERROR] generate ent failed"; [ "$(IGNORE_ERROR)" = "1" ]; }

# build service
.PHONY: build
build:
	@echo "building ${SERVER} service..."
	@cd $(APP_DIR) && \
	go mod tidy && go mod download && \
	go build -trimpath -ldflags "-s -w" -o $(APP_DIR)/server ./cmd/...

# clean all generated files
.PHONY: clean
clean: api-clean config-clean wire-clean ent-clean
	@echo "all generated files have been cleaned."

# generate all code
.PHONY: gen
gen: config ent wire api

# run all of targets
.PHONY: all
all: init tidy gen build
	@echo "build completed successfully."
