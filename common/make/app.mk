ifndef APP_MK_INCLUDED
APP_MK_INCLUDED := 1

# Include common.mk (run-once targets + shared variables)
MAKE_DIR := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))
include $(MAKE_DIR)/common.mk

# --- Per-module variables ---
SERVER := $(notdir $(CURDIR))
APP_DIR := $(ROOT_DIR)/app/$(SERVER)
INTERNAL_DIR := $(APP_DIR)/internal
APP_PROTO_FILES := $(shell find $(INTERNAL_DIR) -type f -name "*.proto" | sort)

IGNORE_ERROR ?= 0

# --- Per-module targets ---

.PHONY: tidy
tidy:
	@echo "go mod tidy..."
	@cd $(APP_DIR) && \
	go mod tidy || \
	{ echo "[ERROR] go mod tidy failed"; [ "$(IGNORE_ERROR)" = "1" ]; }

.PHONY: config-clean
config-clean:
	@echo "clean internal proto files products..."
	@cd $(APP_DIR) && find . -type f \( -name "*.pb.go" -o -name "*.pb.validate.go" -o -name "*.pb.gw.go" \) -delete || true

.PHONY: config
config: config-clean
	@echo "generating internal proto files..."
	@cd $(APP_DIR) && \
	protoc -I $(INTERNAL_DIR) -I $(PROTO_THIRD_PARTY_DIR) -I $(ROOT_DIR) \
	       --go_out=paths=source_relative:$(INTERNAL_DIR) \
	       $(APP_PROTO_FILES) || \
	{ echo "[ERROR] generate config proto failed"; [ "$(IGNORE_ERROR)" = "1" ]; }

.PHONY: wire-clean
wire-clean:
	@echo "clean go wire products..."
	@cd $(APP_DIR)/cmd && find . -type f -name "wire_gen.go" -delete || true

.PHONY: wire
wire: wire-clean
	@echo "generating go wire..."
	@cd $(APP_DIR)/cmd && \
	wire || { echo "[ERROR] generate wire failed"; [ "$(IGNORE_ERROR)" = "1" ]; }

.PHONY: build
build:
	@echo "building ${SERVER} service..."
	@cd $(APP_DIR) && \
	go mod tidy && go mod download && \
	go build -trimpath -ldflags "-s -w" -o $(APP_DIR)/server ./cmd/...

# clean: skip api-clean at root level (root runs it once)
.PHONY: clean
ifdef ROOT_LEVEL
clean: config-clean wire-clean
else
clean: api-clean config-clean wire-clean
endif
	@echo "all generated files have been cleaned."

# gen: skip api at root level (root runs it once)
.PHONY: gen
ifdef ROOT_LEVEL
gen: config wire
else
gen: config api wire
endif

# all: skip init at root level (root runs it once)
.PHONY: all
ifndef ROOT_LEVEL
all: init
endif
all: tidy gen build
	@echo "build completed successfully."

endif
