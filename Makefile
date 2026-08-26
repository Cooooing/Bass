.DEFAULT_GOAL := help

ROOT_DIR := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))
APP_DIR := $(ROOT_DIR)/app

# Auto-discover modules with Makefile under app.
MODULES ?= $(sort $(patsubst $(APP_DIR)/%/Makefile,%,$(wildcard $(APP_DIR)/*/Makefile)))
BFF_SERVERS ?= bbs game_idle_bff

IGNORE_ERROR ?= 1

include $(ROOT_DIR)/common/build/make/common.mk

# --- Root-only targets, avoiding module target collisions. ---

.PHONY: all
all: init api
	@for module in $(MODULES); do \
		echo "---- [$$module] ----"; \
		$(MAKE) -C $(APP_DIR)/$$module all IGNORE_ERROR=$(IGNORE_ERROR) || exit 1; \
	done

.PHONY: gen-all
gen-all: api
	@for module in $(MODULES); do \
		echo "---- [$$module] ----"; \
		$(MAKE) -C $(APP_DIR)/$$module gen IGNORE_ERROR=$(IGNORE_ERROR) || exit 1; \
	done

.PHONY: clean-all
clean-all: api-clean
	@for module in $(MODULES); do \
		echo "---- [$$module] ----"; \
		$(MAKE) -C $(APP_DIR)/$$module clean IGNORE_ERROR=$(IGNORE_ERROR) || exit 1; \
	done

.PHONY: tidy-all build-all
tidy-all build-all:
	@for module in $(MODULES); do \
		echo "---- [$$module] ----"; \
		$(MAKE) -C $(APP_DIR)/$$module $(patsubst %-all,%,$@) IGNORE_ERROR=$(IGNORE_ERROR) || exit 1; \
	done

.PHONY: format fmt
format: api-format
	@for module in $(MODULES); do \
		echo "---- [$$module] ----"; \
		$(MAKE) -C $(APP_DIR)/$$module format IGNORE_ERROR=$(IGNORE_ERROR) || exit 1; \
	done

fmt: format

.PHONY: doc-all
doc-all:
	@for module in $(BFF_SERVERS); do \
		echo "---- [$$module] ----"; \
		$(MAKE) -C $(APP_DIR)/$$module doc IGNORE_ERROR=$(IGNORE_ERROR) || exit 1; \
	done

.PHONY: sdk-validate-all
sdk-validate-all:
	@for module in $(BFF_SERVERS); do \
		echo "---- [$$module] ----"; \
		$(MAKE) -C $(APP_DIR)/$$module sdk-validate IGNORE_ERROR=$(IGNORE_ERROR) || exit 1; \
	done

.PHONY: sdk-all
sdk-all:
	@for module in $(BFF_SERVERS); do \
		echo "---- [$$module] ----"; \
		$(MAKE) -C $(APP_DIR)/$$module sdk IGNORE_ERROR=$(IGNORE_ERROR) || exit 1; \
	done

# --- Help ---
.PHONY: help
help:
	@echo "Available targets:"
	@echo ""
	@echo "Root targets:"
	@echo "  make init         - install development tools"
	@echo "  make api          - generate shared API proto code"
	@echo "  make api-dep      - update Buf dependencies"
	@echo "  make api-lint     - lint shared API proto with Buf"
	@echo "  make api-clean    - clean shared API generated code"
	@echo "  make api-format   - format shared API proto and common Go files"
	@echo "  make format       - format shared API and all app modules"
	@echo "  make all          - run init, api, gen, and build"
	@echo ""
	@echo "Batch targets:"
	@echo "  make gen-all      - generate code for all modules"
	@echo "  make clean-all    - clean generated files for all modules"
	@echo "  make tidy-all     - run go mod tidy for all modules"
	@echo "  make build-all    - build all services"
	@echo "  make doc-all      - generate BFF OpenAPI documents"
	@echo "  make sdk-validate-all - validate BFF OpenAPI documents"
	@echo "  make sdk-all      - generate BFF SDKs from OpenAPI"
	@echo ""
	@echo "Single module targets:"
	@echo "  make -C app/<module> gen       - generate code for one module"
	@echo "  make -C app/<module> clean     - clean one module"
	@echo "  make -C app/<module> build     - build one module"
	@echo "  make -C app/<module> tidy      - run go mod tidy for one module"
	@echo "  make -C app/<module> format    - format one module"
	@echo "  make -C app/<module> ent       - generate Ent code"
	@echo "  make -C app/<module> doc       - generate OpenAPI document"
	@echo "  make -C app/<module> sdk       - generate BFF TypeScript Axios/Fetch, Go, Java and Rust SDKs"
	@echo ""
	@echo "Examples:"
	@echo "    make -C app/bbs gen"
	@echo "    make -C app/user build"
	@echo "    make -C app/bbs ent"
	@echo "    make -C app/bbs doc"
