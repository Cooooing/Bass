ifndef APP_MK_INCLUDED
APP_MK_INCLUDED := 1

# Include common.mk for one-time targets and shared variables.
MAKE_DIR := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))
include $(MAKE_DIR)/common.mk

# --- Module variables ---
SERVER := $(notdir $(CURDIR))
APP_DIR := $(ROOT_DIR)/app/$(SERVER)
INTERNAL_DIR := $(APP_DIR)/internal
APP_PROTO_FILES := $(shell find $(INTERNAL_DIR) -type f -name "*.proto" | sort)
BUF_SERVICE_CONFIG := '{"version":"v2","modules":[{"path":"common/proto/app"},{"path":"app/$(SERVER)"}],"deps":["buf.build/googleapis/googleapis"],"lint":{"use":["MINIMAL"],"except":["PACKAGE_DIRECTORY_MATCH","PACKAGE_SAME_DIRECTORY"]},"breaking":{"use":["FILE"]}}'

IGNORE_ERROR ?= 0
MODULE_GEN_TARGETS := config wire

# run is an error handling wrapper selected at parse time by ifeq.
# Usage: $(call run,command,error-label)
# Commands are hidden by default. Failures print labels to stderr.
ifeq ($(IGNORE_ERROR),1)
run = @$(1) || echo "[ERROR] $(2) failed" >&2
else
run = @$(1) || { echo "[ERROR] $(2) failed" >&2; exit 1; }
endif

# --- Module targets ---

.PHONY: tidy
tidy:
	@echo "[tidy] go mod tidy..."
	$(call run,cd $(APP_DIR) && go mod tidy,[tidy] go mod tidy)

.PHONY: format
format:
	@echo "[format] gofmt module Go files..."
	$(call run,find $(APP_DIR) -type f -name "*.go" -exec gofmt -w {} +,[format] gofmt)
	@echo "[format] format module proto..."
	$(call run,cd $(ROOT_DIR) && $(BUF) format -w app/$(SERVER),[format] buf format)

.PHONY: config-clean
config-clean:
	@echo "[config-clean] cleaning proto generated Go files..."
	@cd $(APP_DIR) 2>/dev/null && find . -type f \( -name "*.pb.go" -o -name "*.pb.validate.go" -o -name "*.pb.gw.go" \) -delete 2>/dev/null; true

.PHONY: config
config: config-clean
	@echo "[config] buf generate..."
	$(call run,cd $(ROOT_DIR) && $(BUF) generate --config $(BUF_SERVICE_CONFIG) --template $(BUF_GEN_CONFIG) --path app/$(SERVER)/internal/config/config.proto --output app/$(SERVER),[config] buf generate)

.PHONY: wire-clean
wire-clean:
	@echo "[wire-clean] cleaning wire_gen.go..."
	@find $(APP_DIR) -type f -name "wire_gen.go" -delete; true

.PHONY: wire
wire: wire-clean
	@echo "[wire] wire..."
	@find $(APP_DIR) -type f -name "wire.go" -exec dirname {} \; | sort -u | while read dir; do \
		(cd "$$dir" && wire) || exit 1; \
	done

.PHONY: build
build:
	@echo "[build] go build..."
	$(call run,cd $(APP_DIR) && go mod tidy && go mod download && go build -trimpath -ldflags "-s -w" -o $(APP_DIR)/server ./cmd/...,[build] go build)

.PHONY: build-clean
build-clean:
	@echo "[build-clean] cleaning service binary..."
	@rm -f $(APP_DIR)/server 2>/dev/null; true

.PHONY: clean
clean: config-clean
clean: wire-clean
clean: build-clean

.PHONY: gen
gen:
	@for target in $(MODULE_GEN_TARGETS); do \
		$(MAKE) $$target IGNORE_ERROR=$(IGNORE_ERROR) || exit 1; \
	done

.PHONY: all
all: init api tidy gen build

endif
