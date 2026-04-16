# ============ Global Variables ============
GOHOSTOS := $(shell go env GOHOSTOS)
VERSION := latest

SERVERS := gateway infra user content notify im signal connector
ROOT_DIR := $(realpath $(dir $(lastword $(MAKEFILE_LIST))))

# 默认目标
.DEFAULT_GOAL := help

# ============ Core Dispatcher ============

.PHONY: $(SERVERS)
$(SERVERS):
	@echo "===> [$@] $(SUBTARGET)"
	@$(MAKE) -C $(ROOT_DIR)/app/$@ $(SUBTARGET)

# ============ High-level Targets ============

.PHONY: init gen wire ent config api tidy
init gen wire ent config api tidy:
	@$(MAKE) SUBTARGET=$@ $(SERVERS)

# ============ Help ============

.PHONY: help
help:
	@echo "Available targets:"
	@echo "  make init     - run init for all services"
	@echo "  make gen      - run code generation"
	@echo "  make wire     - run wire injection"
	@echo "  make ent      - run ent generation"
	@echo "  make config   - generate configs"
	@echo "  make api      - generate API"
	@echo "  make tidy     - run go mod tidy"
	@echo ""
	@echo "Parallel example:"
	@echo "  make -j8 tidy"