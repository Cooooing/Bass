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

# run: error-handling wrapper (defined at parse time via ifeq, not $(if))
# Usage: $(call run,shell-command,error-label)
# Suppresses command echo (@), prints label to stderr on failure.
ifeq ($(IGNORE_ERROR),1)
run = @$(1) || echo "[ERROR] $(2) failed" >&2
else
run = @$(1) || { echo "[ERROR] $(2) failed" >&2; exit 1; }
endif

# --- Per-module targets ---

.PHONY: tidy
tidy:
	@echo "[tidy] go mod tidy..."
	$(call run,cd $(APP_DIR) && go mod tidy,[tidy] go mod tidy)

.PHONY: config-clean
config-clean:
	@echo "[config-clean] removing proto go files..."
	@cd $(APP_DIR) 2>/dev/null && find . -type f \( -name "*.pb.go" -o -name "*.pb.validate.go" -o -name "*.pb.gw.go" \) -delete 2>/dev/null; true

.PHONY: config
config: config-clean
	@echo "[config] protoc..."
	$(call run,cd $(APP_DIR) && protoc -I $(INTERNAL_DIR) -I $(PROTO_THIRD_PARTY_DIR) -I $(ROOT_DIR) --go_out=paths=source_relative:$(INTERNAL_DIR) $(APP_PROTO_FILES),[config] protoc)

.PHONY: wire-clean
wire-clean:
	@echo "[wire-clean] removing wire_gen.go..."
	@cd $(APP_DIR)/cmd 2>/dev/null && find . -type f -name "wire_gen.go" -delete 2>/dev/null; true

.PHONY: wire
wire: wire-clean
	@echo "[wire] wire..."
	$(call run,cd $(APP_DIR)/cmd && wire,[wire] wire)

.PHONY: build
build:
	@echo "[build] go build..."
	$(call run,cd $(APP_DIR) && go mod tidy && go mod download && go build -trimpath -ldflags "-s -w" -o $(APP_DIR)/server ./cmd/...,[build] go build)

.PHONY: build-clean
build-clean:
	@echo "[build-clean] removing server binary..."
	@rm -f $(APP_DIR)/server 2>/dev/null; true

.PHONY: clean
clean: config-clean
clean: wire-clean
clean: build-clean

.PHONY: gen
gen: config
gen: wire

.PHONY: all
all: init api tidy gen build

endif
