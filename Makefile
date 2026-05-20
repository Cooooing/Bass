.DEFAULT_GOAL := help

ROOT_DIR := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))
APP_DIR := $(ROOT_DIR)/app

# Auto-discover modules (all directories with a Makefile under app/)
MODULES ?= $(sort $(patsubst $(APP_DIR)/%/Makefile,%,$(wildcard $(APP_DIR)/*/Makefile)))
BFF_SERVERS ?= bbs

IGNORE_ERROR ?= 1

include $(ROOT_DIR)/common/make/common.mk

# --- Root-only targets (unique names, no module collision) ---

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

.PHONY: doc-all
doc-all:
	@for module in $(BFF_SERVERS); do \
		echo "---- [$$module] ----"; \
		$(MAKE) -C $(APP_DIR)/$$module doc IGNORE_ERROR=$(IGNORE_ERROR) || exit 1; \
	done

# --- help ---
.PHONY: help
help:
	@echo "Available targets:"
	@echo ""
	@echo "Root-level (run once for all modules):"
	@echo "  make init         - install tools"
	@echo "  make api          - generate shared API proto codes"
	@echo "  make api-clean    - clean shared API proto codes"
	@echo "  make all          - full pipeline: init + api + gen + build"
	@echo ""
	@echo "Batch (run across all modules):"
	@echo "  make gen-all      - generate all codes (api + per-module)"
	@echo "  make clean-all    - clean all generated files"
	@echo "  make tidy-all     - run go mod tidy for all modules"
	@echo "  make build-all    - build all services"
	@echo "  make doc-all      - generate BFF OpenAPI docs"
	@echo ""
	@echo "Single module (make -C app/<module> <target>):"
	@echo "  make -C app/<module> gen       - generate codes for one module"
	@echo "  make -C app/<module> clean     - clean one module"
	@echo "  make -C app/<module> build     - build one module"
	@echo "  make -C app/<module> tidy      - go mod tidy for one module"
	@echo "  make -C app/<module> ent       - generate ent codes (if ent.mk included)"
	@echo "  make -C app/<module> doc       - generate openapi docs (if doc.mk included)"
	@echo ""
	@echo "  Examples:"
	@echo "    make -C app/bbs gen"
	@echo "    make -C app/user build"
	@echo "    make -C app/bbs ent"
	@echo "    make -C app/bbs doc"
