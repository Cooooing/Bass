.DEFAULT_GOAL := help

ROOT_DIR := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))
APP_DIR := $(ROOT_DIR)/app

# Auto-discover modules with Makefile under app.
MODULES ?= $(sort $(patsubst $(APP_DIR)/%/Makefile,%,$(wildcard $(APP_DIR)/*/Makefile)))
BFF_SERVERS ?= bff_bbs bff_bbs_admin bff_game_idle

IGNORE_ERROR ?= 0

include $(ROOT_DIR)/common/build/make/common.mk

# --- Root-only targets. ---

.PHONY: gen
gen: api
	@for module in $(MODULES); do \
		echo "---- [$$module] ----"; \
		$(MAKE) -C $(APP_DIR)/$$module gen IGNORE_ERROR=$(IGNORE_ERROR) || exit 1; \
	done

.PHONY: gen-clean
gen-clean: api-clean
	@for module in $(MODULES); do \
		echo "---- [$$module] ----"; \
		$(MAKE) -C $(APP_DIR)/$$module gen-clean IGNORE_ERROR=$(IGNORE_ERROR) || exit 1; \
	done

.PHONY: build
build:
	@for module in $(MODULES); do \
		echo "---- [$$module] ----"; \
		$(MAKE) -C $(APP_DIR)/$$module build IGNORE_ERROR=$(IGNORE_ERROR) || exit 1; \
	done

.PHONY: build-clean
build-clean:
	@for module in $(MODULES); do \
		echo "---- [$$module] ----"; \
		$(MAKE) -C $(APP_DIR)/$$module build-clean IGNORE_ERROR=$(IGNORE_ERROR) || exit 1; \
	done

.PHONY: doc
doc:
	@for module in $(BFF_SERVERS); do \
		echo "---- [$$module] ----"; \
		$(MAKE) -C $(APP_DIR)/$$module doc IGNORE_ERROR=$(IGNORE_ERROR) || exit 1; \
	done

.PHONY: doc-clean
doc-clean:
	@for module in $(BFF_SERVERS); do \
		echo "---- [$$module] ----"; \
		$(MAKE) -C $(APP_DIR)/$$module doc-clean IGNORE_ERROR=$(IGNORE_ERROR) || exit 1; \
	done

.PHONY: sdk
sdk:
	@for module in $(BFF_SERVERS); do \
		echo "---- [$$module] ----"; \
		$(MAKE) -C $(APP_DIR)/$$module sdk IGNORE_ERROR=$(IGNORE_ERROR) || exit 1; \
	done

.PHONY: sdk-clean
sdk-clean:
	@for module in $(BFF_SERVERS); do \
		echo "---- [$$module] ----"; \
		$(MAKE) -C $(APP_DIR)/$$module sdk-clean IGNORE_ERROR=$(IGNORE_ERROR) || exit 1; \
	done

# --- Help ---
.PHONY: help
help:
	@echo "Available targets (all direct Make targets):"
	@echo ""
	@echo "Root:"
	@echo "  make help         - show this command list"
	@echo "  make init         - install pinned development tools"
	@echo "  make api          - generate shared API code"
	@echo "  make api-clean    - clean shared API code"
	@echo "  make gen          - generate all module code"
	@echo "  make gen-clean    - clean all generated module code"
	@echo "  make build        - build all app modules"
	@echo "  make build-clean  - clean all app binaries"
	@echo "  make doc          - generate OpenAPI documents for all BFFs"
	@echo "  make doc-clean    - clean all BFF OpenAPI documents"
	@echo "  make sdk          - generate SDKs for all BFFs"
	@echo "  make sdk-clean    - clean all BFF SDKs"
	@echo ""
	@echo "App module (make -C app/<module> <target>):"
	@echo "  init                            - install pinned development tools (same global setup)"
	@echo "  api / api-clean                 - generate / clean shared API code (same root artifact)"
	@echo "  gen / gen-clean                 - generate / clean this module's code"
	@echo "  build / build-clean             - build / clean this module's binary"
	@echo "  cfg / cfg-clean                 - generate / clean configuration protobuf code"
	@echo "  wire / wire-clean               - generate / clean Wire injection code"
	@echo ""
	@echo "Ent module only (content, economy, game_idle, game_town, im, notify, platform, scheduler, user):"
	@echo "  ent / ent-clean                 - generate / clean Ent data-access code"
	@echo ""
	@echo "BFF module only (make -C app/<bff> <target>):"
	@echo "  doc / doc-clean                 - generate / clean OpenAPI document"
	@echo "  sdk / sdk-clean                 - generate / clean all SDK languages"
	@echo ""
	@echo "Monolith (make -C monolith <target>):"
	@echo "  init                            - install pinned development tools (same global setup)"
	@echo "  api / api-clean                 - generate / clean shared API code (same root artifact)"
	@echo "  gen / gen-clean                 - generate / clean monolith code"
	@echo "  build / build-clean             - build / clean monolith binary"
	@echo "  cfg / cfg-clean                 - generate / clean configuration protobuf code"
	@echo "  wire / wire-clean               - generate / clean Wire injection code"
	@echo ""
	@echo "Examples:"
	@echo "    make -C app/bff_bbs gen"
	@echo "    make -C app/user build"
	@echo "    make -C app/bff_bbs doc"
