ifndef ENT_MK_INCLUDED
ENT_MK_INCLUDED := 1

# Requires app.mk for APP_DIR, IGNORE_ERROR, run, and MODULE_GEN_TARGETS.

MODULE_GEN_TARGETS := config ent wire

ENT_SCHEMA_DIR ?= $(if $(wildcard $(APP_DIR)/internal/data/schema),./internal/data/schema,./internal/data/ent/schema)
ENT_GEN_DIR ?= $(APP_DIR)/internal/data/gen

ENT_FEATURES := sql/upsert sql/modifier intercept
clean: ent-clean

# Clean Ent generated artifacts.
.PHONY: ent-clean
ent-clean:
	@echo "[ent-clean] cleaning Ent generated code..."
	@rm -rf $(APP_DIR)/internal/data/gen 2>/dev/null; true

# Generate Ent code.
.PHONY: ent
ent: ent-clean
	@echo "[ent] ent generate..."
	$(call run,cd $(APP_DIR) && ent generate --target=$(ENT_GEN_DIR) $(foreach feature,$(ENT_FEATURES),--feature $(feature)) $(ENT_SCHEMA_DIR),[ent] ent generate)

endif
