.DEFAULT_GOAL := help

ROOT_DIR := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))
APP_DIR := $(ROOT_DIR)/app

# Auto-discover modules that have a Makefile
#SERVERS := $(sort $(patsubst $(APP_DIR)/%/Makefile,%,$(wildcard $(APP_DIR)/*/Makefile)))
SERVERS := bbs user content notify

IGNORE_ERROR ?= 1

include $(ROOT_DIR)/common/make/common.mk

# Dispatch to single module: make user gen
.PHONY: $(SERVERS)
$(SERVERS):
	@echo "===> [$@] $(SUBTARGET)"
	@$(MAKE) -C $(APP_DIR)/$@ $(SUBTARGET) IGNORE_ERROR=$(IGNORE_ERROR)

# gen: api once, then gen each module (ROOT_LEVEL=1 skips api per-module)
.PHONY: gen
gen: api
	@for module in $(SERVERS); do \
		echo "===> [$$module] gen"; \
		$(MAKE) -C $(APP_DIR)/$$module gen IGNORE_ERROR=$(IGNORE_ERROR) ROOT_LEVEL=1 || exit 1; \
	done

# clean: api-clean once, then clean each module
.PHONY: clean
clean: api-clean
	@for module in $(SERVERS); do \
		echo "===> [$$module] clean"; \
		$(MAKE) -C $(APP_DIR)/$$module clean IGNORE_ERROR=$(IGNORE_ERROR) ROOT_LEVEL=1 || exit 1; \
	done

# all: init+api once, then all each module
.PHONY: all
all: init api
	@for module in $(SERVERS); do \
		echo "===> [$$module] all"; \
		$(MAKE) -C $(APP_DIR)/$$module all IGNORE_ERROR=$(IGNORE_ERROR) ROOT_LEVEL=1 || exit 1; \
	done

# Per-module targets: dispatch to each module directly
.PHONY: tidy build config config-clean wire wire-clean ent ent-clean
tidy build config config-clean wire wire-clean ent ent-clean:
	@for module in $(SERVERS); do \
		echo "===> [$$module] $@"; \
		$(MAKE) -C $(APP_DIR)/$$module $@ IGNORE_ERROR=$(IGNORE_ERROR) || exit 1; \
	done

# help
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  make init     - install tools"
	@echo "  make tidy     - run go mod tidy"
	@echo "  make config   - generate proto config codes"
	@echo "  make wire     - generate wire codes"
	@echo "  make ent      - generate ent codes"
	@echo "  make api      - generate proto API codes"
	@echo "  make gen      - generate all codes"
	@echo "  make build    - build all services"
	@echo "  make clean    - clean all generated files"
	@echo "  make all      - init + gen + build all"
	@echo ""
	@echo "Single module:"
	@echo "  make user gen       - gen for user module only"
	@echo "  make user SUBTARGET=build - build user module"
