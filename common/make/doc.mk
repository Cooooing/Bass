ifndef DOC_MK_INCLUDED
DOC_MK_INCLUDED := 1

# Requires app.mk to be included first (provides PROTO_DIR, PROTO_THIRD_PARTY_DIR, PROTO_GEN_DIR)

BFF_SERVER := $(notdir $(CURDIR))
BFF_PROTO_DIR := $(PROTO_DIR)/$(BFF_SERVER)
BFF_PROTO_FILES := $(shell find $(BFF_PROTO_DIR) -type f -name "*.proto" | sort)
BFF_OPENAPI_DIR := $(PROTO_GEN_DIR)/openapi/$(BFF_SERVER)

# Append to composite targets
gen: doc
clean: doc-clean

# clean openapi
.PHONY: doc-clean
doc-clean:
	@echo "[doc-clean] removing openapi..."
	@rm -rf $(PROTO_GEN_DIR)/openapi/$(BFF_SERVER) 2>/dev/null; true

# Generate OpenAPI spec for this BFF service
.PHONY: doc
doc: doc-clean
	@echo "[doc] protoc openapi..."
	@mkdir -p $(BFF_OPENAPI_DIR)
	@protoc -I $(PROTO_DIR) -I $(PROTO_THIRD_PARTY_DIR) \
	       --openapi_out=fq_schema_naming=true,naming=proto,default_response=false:$(BFF_OPENAPI_DIR) \
	       $(BFF_PROTO_FILES) || \
	{ echo "[ERROR] [doc] protoc openapi failed" >&2; exit 1; }

endif
