ifndef ENT_MK_INCLUDED
ENT_MK_INCLUDED := 1

# Requires app.mk to be included first (provides APP_DIR, IGNORE_ERROR)

# Append to composite targets (prerequisites are accumulated, no recipe conflict)
gen: ent
clean: ent-clean

# clean ent products
.PHONY: ent-clean
ent-clean:
	@echo "clean go ent products..."
	@rm -rf $(APP_DIR)/internal/data/ent/gen 2>/dev/null || true

# generate ent
.PHONY: ent
ent: ent-clean
	@echo "generating go ent..."
	@cd $(APP_DIR) && \
	ent generate --target=$(APP_DIR)/internal/data/ent/gen ./internal/data/ent/schema || \
	{ echo "[ERROR] generate ent failed"; [ "$(IGNORE_ERROR)" = "1" ]; }

endif
