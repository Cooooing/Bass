ifndef DOC_MK_INCLUDED
DOC_MK_INCLUDED := 1

# Requires app.mk to be included first (provides PROTO_DIR, PROTO_THIRD_PARTY_DIR, PROTO_GEN_DIR)

BFF_SERVER := $(notdir $(CURDIR))
BFF_PROTO_DIR := $(PROTO_DIR)/$(BFF_SERVER)
BFF_OPENAPI_DIR := $(PROTO_GEN_DIR)/openapi/$(BFF_SERVER)

gen: doc
clean: doc-clean

.PHONY: doc-clean
doc-clean:
	@echo "clean openapi..."
	@rm -rf $(PROTO_GEN_DIR)/openapi/$(BFF_SERVER) || true

# Generate OpenAPI spec for this BFF service
.PHONY: doc
doc: doc-clean
	@echo "generating openapi..."
	@mkdir -p $(BFF_OPENAPI_DIR)
	protoc -I $(PROTO_DIR) -I $(PROTO_THIRD_PARTY_DIR) \
	       --openapi_out=fq_schema_naming=true,naming=proto,default_response=false:$(BFF_OPENAPI_DIR) \
	       $(COMMON_PROTO_FILES)

endif
