# ============ Global Variables ============
GOHOSTOS := $(shell go env GOHOSTOS)
VERSION := latest
SERVERS := gateway user content notify im signal connector infra

ROOT_DIR := $(realpath $(dir $(lastword $(MAKEFILE_LIST))))

.PHONY: $(SERVERS)
$(SERVERS):
	@echo "===> [$@] $(SUBTARGET)"
	@$(MAKE) -C $(ROOT_DIR)/app/$@ $(SUBTARGET)

.PHONY: gen
gen:
	@$(MAKE) SUBTARGET=gen $(SERVERS)

.PHONY: wire
wire:
	@$(MAKE) SUBTARGET=wire $(SERVERS)

.PHONY: ent
ent:
	@$(MAKE) SUBTARGET=ent $(SERVERS)

.PHONY: config
config:
	@$(MAKE) SUBTARGET=config $(SERVERS)

.PHONY: api
api:
	@$(MAKE) SUBTARGET=api $(SERVERS)
