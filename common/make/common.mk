# global variables
ROOT_DIR := $(abspath $(dir $(lastword $(MAKEFILE_LIST)))/../../)
COMMON_DIR := $(ROOT_DIR)/common
PROTO_DIR := $(COMMON_DIR)/api/app
PROTO_GEN_DIR := $(COMMON_DIR)/api/gen
PROTO_THIRD_PARTY_DIR := $(COMMON_DIR)/api/third_party

# find proto files
COMMON_PROTO_FILES = $(shell go run $(COMMON_DIR)/build_tools/findproto.go $(PROTO_DIR))

# tools install
.PHONY: init
init:
	@echo "installing required tools..."
	@cd $(COMMON_DIR) && \
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

# clean API proto files
.PHONY: api-clean
api-clean:
	@echo "clean API proto files..."
	@cd $(PROTO_GEN_DIR) && find . -type f ! -name ".gitkeep" -delete 2>/dev/null || true && find . -type d -empty -delete 2>/dev/null || true

# generate API proto files
.PHONY: api
api: api-clean
	@echo "generating API proto files..."
	@protoc -I $(PROTO_DIR) -I $(PROTO_THIRD_PARTY_DIR) \
	       --go_out=paths=source_relative:$(PROTO_GEN_DIR) \
	       --go-grpc_out=paths=source_relative:$(PROTO_GEN_DIR) \
	       --go-http_out=paths=source_relative:$(PROTO_GEN_DIR) \
	       --go-errors_out=paths=source_relative:$(PROTO_GEN_DIR) \
	       --openapi_out=fq_schema_naming=true,naming=proto,default_response=false:$(PROTO_GEN_DIR) \
	       --validate_out=lang=go,paths=source_relative:$(PROTO_GEN_DIR) \
	       $(COMMON_PROTO_FILES)
