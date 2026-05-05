# global variables
GOHOSTOS := $(shell go env GOHOSTOS)
VERSION := latest

SERVERS := gateway user content notify im signal connector
ROOT_DIR := $(realpath $(dir $(lastword $(MAKEFILE_LIST))))

.DEFAULT_GOAL := help # default target

IGNORE_ERROR ?= 1

# dispatch
.PHONY: $(SERVERS)
$(SERVERS):
	@echo "===> [$@] $(SUBTARGET)"
	@$(MAKE) -C $(ROOT_DIR)/app/$@ $(SUBTARGET) IGNORE_ERROR=$(IGNORE_ERROR)

# target
.PHONY: wire wire-clean ent ent-clean config config-clean tidy gen all clean
wire wire-clean ent ent-clean config config-clean tidy gen all clean:
	@$(MAKE) SUBTARGET=$@ $(SERVERS)

include $(ROOT_DIR)/common/make/common.mk

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
	@echo ""
	@echo "Parallel example:"
	@echo "  make -j8 tidy"