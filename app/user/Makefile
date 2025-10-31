# ============ Global Variables ============
GOHOSTOS := $(shell go env GOHOSTOS)
VERSION := latest
SERVER := user

ROOT_DIR := $(realpath $(dir $(lastword $(MAKEFILE_LIST)))/../..)
COMMON_DIR := $(ROOT_DIR)/common
USER_DIR := $(ROOT_DIR)/app/$(SERVER)
API_DIR := $(COMMON_DIR)/api
THIRD_PARTY_DIR := $(COMMON_DIR)/third_party
INTERNAL_DIR := $(USER_DIR)/internal

# ============ Find Proto Files (Cross-Platform Safe) ============
INTERNAL_PROTO_FILES = $(shell go run $(COMMON_DIR)/build_tools/findproto.go $(INTERNAL_DIR))
API_PROTO_FILES = $(shell go run $(COMMON_DIR)/build_tools/findproto.go $(API_DIR))

# ============ Tool Installation ============
.PHONY: init
init:
	@echo "Installing required tools..."
	@cd $(USER_DIR) && \
	go mod tidy
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/envoyproxy/protoc-gen-validate@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/google/wire/cmd/wire@latest

# ============ Compile config Proto ============
.PHONY: config
config:
	@echo "Generating internal proto files..."
	@cd $(USER_DIR) && \
	protoc -I $(INTERNAL_DIR) \
	       --go_out=paths=source_relative:$(INTERNAL_DIR) \
	       $(INTERNAL_PROTO_FILES)

# ============ Generate Wire Code ============
.PHONY: wire
wire:
	@echo "Generating go wire..."
	@cd $(USER_DIR)/cmd && \
	go run github.com/google/wire/cmd/wire

# ============ Generate Ent Code ============
.PHONY: ent
ent:
	@echo "Generating go ent..."
	@cd $(USER_DIR) && \
	go run entgo.io/ent/cmd/ent generate --target=$(USER_DIR)/internal/data/ent/gen ./internal/data/ent/schema

# ============ Compile API Proto ============
.PHONY: api
api:
	@echo "Generating API proto files..."
	@cd $(USER_DIR) && \
	protoc -I $(API_DIR) -I $(THIRD_PARTY_DIR) \
	       --go_out=paths=source_relative:$(API_DIR) \
	       --go-grpc_out=paths=source_relative:$(API_DIR) \
	       --go-http_out=paths=source_relative:$(API_DIR) \
	       --go-errors_out=paths=source_relative:$(API_DIR) \
	       --openapi_out=fq_schema_naming=true,default_response=false:$(API_DIR) \
	       --validate_out=lang=go,paths=source_relative:$(API_DIR) \
	       $(API_PROTO_FILES)

# ============ Generate All Code ============
.PHONY: gen
gen: config ent wire api

# ============ Build Service ============
.PHONY: build
build:
	@echo "Building user service..."
	@cd $(USER_DIR) && \
	go mod tidy && \
	go build -o $(ROOT_DIR)/$(SERVER) ./cmd/...

# ============ Run All Steps ============
.PHONY: all
all: init gen build
	@echo "Build completed successfully."

