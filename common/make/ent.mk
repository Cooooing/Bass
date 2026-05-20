ifndef ENT_MK_INCLUDED
ENT_MK_INCLUDED := 1

# Requires app.mk to be included first (provides APP_DIR, IGNORE_ERROR, run)

# Append to composite targets (prerequisites are accumulated, no recipe conflict)
gen: ent
clean: ent-clean

# clean ent products
.PHONY: ent-clean
ent-clean:
	@echo "[ent-clean] removing ent gen..."
	@rm -rf $(APP_DIR)/internal/data/gen 2>/dev/null; true

# generate ent
.PHONY: ent
ent: ent-clean
	@echo "[ent] ent generate..."
	$(call run,cd $(APP_DIR) && ent generate --target=$(APP_DIR)/internal/data/gen ./internal/data/schema,[ent] ent generate)

endif
