# ============ Global Variables ============
GOHOSTOS := $(shell go env GOHOSTOS)
VERSION := latest
SERVER := user

ROOT_DIR := $(abspath $(dir $(lastword $(MAKEFILE_LIST)))/../../)
COMMON_DIR := $(ROOT_DIR)/common
APP_DIR := $(ROOT_DIR)/app/$(SERVER)
API_DIR := $(COMMON_DIR)/api/app
API_GEN_DIR := $(COMMON_DIR)/api/gen
THIRD_PARTY_DIR := $(COMMON_DIR)/api/third_party
INTERNAL_DIR := $(APP_DIR)/internal

# ============ Find Proto Files (Cross-Platform Safe) ============
INTERNAL_PROTO_FILES = $(shell go run $(COMMON_DIR)/build_tools/findproto.go $(INTERNAL_DIR))
# Filter out any stray non-proto tokens (e.g., lines like 'root: <path>') emitted by the finder on some platforms
API_PROTO_FILES = $(shell go run $(COMMON_DIR)/build_tools/findproto.go $(API_DIR))

# ============ Tool Installation ============
.PHONY: init
init:
	@echo "Installing required tools..."
	@cd $(APP_DIR) && \
	go mod tidy && go mod download
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest
	go install github.com/envoyproxy/protoc-gen-validate@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/google/wire/cmd/wire@latest
	go install entgo.io/ent/cmd/ent@latest

# ============ Compile config Proto ============
.PHONY: config
config:
	@echo "Generating internal proto files..."
	protoc -I $(INTERNAL_DIR) -I $(THIRD_PARTY_DIR) \
	       --go_out=paths=source_relative:$(INTERNAL_DIR) \
	       $(INTERNAL_PROTO_FILES)

# ============ Generate Wire Code ============
.PHONY: wire
wire:
	@echo "Generating go wire..."
	@cd $(APP_DIR)/cmd && \
	wire

# ============ Generate Ent Code ============
.PHONY: ent
ent:
	@echo "Generating go ent..."
	@cd $(APP_DIR) && \
	ent generate --target=$(APP_DIR)/internal/data/ent/gen ./internal/data/ent/schema

# ============ Compile API Proto ============
.PHONY: api
api:
	@echo "Generating API proto files..."
	protoc -I $(API_DIR) -I $(THIRD_PARTY_DIR) \
	       --go_out=paths=source_relative:$(API_GEN_DIR) \
	       --go-grpc_out=paths=source_relative:$(API_GEN_DIR) \
	       --go-http_out=paths=source_relative:$(API_GEN_DIR) \
	       --go-errors_out=paths=source_relative:$(API_GEN_DIR) \
	       --openapi_out=fq_schema_naming=true,naming=proto,default_response=false:$(API_GEN_DIR) \
	       --validate_out=lang=go,paths=source_relative:$(API_GEN_DIR) \
	       $(API_PROTO_FILES)

# ============ Generate All Code ============
.PHONY: gen
gen: config ent wire api

# ============ Build Service ============
.PHONY: build
build:
	@echo "Building ${SERVER} service..."
	@cd $(APP_DIR) && \
	go mod tidy && go mod download && \
	go build -trimpath -ldflags "-s -w" -o $(ROOT_DIR)/$(SERVER) ./cmd/...

# ============ Run All Steps ============
.PHONY: all
all: init gen build
	@echo "Build completed successfully."
