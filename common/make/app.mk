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
APP_PROTO_FILES = $(shell go run $(COMMON_DIR)/build_tools/findproto.go $(INTERNAL_DIR))

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

# generate config proto
.PHONY: config
config:
	@echo "Generating internal proto files..."
	@cd $(APP_DIR) && \
	protoc -I $(INTERNAL_DIR) -I $(PROTO_THIRD_PARTY_DIR) \
	       --go_out=paths=source_relative:$(INTERNAL_DIR) \
	       $(APP_PROTO_FILES) || \
	{ echo "[ERROR] generate config proto failed"; [ "$(IGNORE_ERROR)" = "1" ]; }

# generate wire
.PHONY: wire
wire:
	@echo "Generating go wire..."
	@cd $(APP_DIR)/cmd && \
	wire || { echo "[ERROR] generate wire failed"; [ "$(IGNORE_ERROR)" = "1" ]; }

# generate ent
.PHONY: ent
ent:
	@echo "Generating go ent..."
	@cd $(APP_DIR) && \
	ent generate --target=$(APP_DIR)/internal/data/ent/gen ./internal/data/ent/schema || \
	{ echo "[ERROR] generate ent failed"; [ "$(IGNORE_ERROR)" = "1" ]; }

# build service
.PHONY: build
build:
	@echo "Building ${SERVER} service..."
	@cd $(APP_DIR) && \
	go mod tidy && go mod download && \
	go build -trimpath -ldflags "-s -w" -o $(APP_DIR)/server ./cmd/...

# generate all code
.PHONY: gen
gen: config ent wire api

# run all of targets
.PHONY: all
all: init tidy gen build
	@echo "Build completed successfully."
