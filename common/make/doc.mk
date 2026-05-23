ifndef DOC_MK_INCLUDED
DOC_MK_INCLUDED := 1

# Requires app.mk for PROTO_DIR and PROTO_GEN_DIR.

BFF_SERVER := $(notdir $(CURDIR))
BFF_PROTO_DIR := $(PROTO_DIR)/$(BFF_SERVER)
BFF_OPENAPI_DIR := $(PROTO_GEN_DIR)/openapi/$(BFF_SERVER)
BFF_OPENAPI_FILE := $(BFF_OPENAPI_DIR)/openapi.yaml
BFF_GEN_TS_DIR := $(ROOT_DIR)/common/api/gen-ts/$(BFF_SERVER)
BFF_GEN_GO_DIR := $(ROOT_DIR)/common/api/gen-go/$(BFF_SERVER)
OPENAPI_GENERATOR ?= cd $(ROOT_DIR)/common/api/sdk && npx --yes @openapitools/openapi-generator-cli

# Append to composite target sequence.
MODULE_GEN_TARGETS += doc
clean: doc-clean
clean: sdk-clean

# Clean OpenAPI artifacts.
.PHONY: doc-clean
doc-clean:
	@echo "[doc-clean] cleaning OpenAPI documents..."
	@rm -rf $(PROTO_GEN_DIR)/openapi/$(BFF_SERVER) 2>/dev/null; true

# Generate OpenAPI document for the current BFF service.
.PHONY: doc
doc: doc-clean
	@echo "[doc] buf generate openapi..."
	@mkdir -p $(BFF_OPENAPI_DIR)
	@cd $(ROOT_DIR) && $(BUF) generate $(BUF_CONFIG_DIR) --path common/api/app/$(BFF_SERVER) --template $(BUF_GEN_OPENAPI) --output $(BFF_OPENAPI_DIR)

# Check SDK generator prerequisites.
.PHONY: sdk-prereq
sdk-prereq:
	@command -v npx >/dev/null 2>&1 || { echo "[ERROR] [sdk] OpenAPI Generator requires npx" >&2; exit 1; }
	@command -v java >/dev/null 2>&1 || { echo "[ERROR] [sdk] OpenAPI Generator CLI requires Java" >&2; exit 1; }

# Validate OpenAPI document with OpenAPI Generator.
.PHONY: sdk-validate
sdk-validate: doc sdk-prereq
	@echo "[sdk-validate] validating OpenAPI document..."
	@$(OPENAPI_GENERATOR) validate -i $(BFF_OPENAPI_FILE)

# Clean generated SDK artifacts.
.PHONY: sdk-clean
sdk-clean:
	@echo "[sdk-clean] cleaning generated SDKs..."
	@rm -rf $(BFF_GEN_TS_DIR) $(BFF_GEN_GO_DIR) 2>/dev/null; true

# Generate TypeScript fetch SDK.
.PHONY: sdk-ts
sdk-ts: doc sdk-prereq
	@echo "[sdk-ts] openapi-generator typescript-fetch..."
	@mkdir -p $(BFF_GEN_TS_DIR)
	@$(OPENAPI_GENERATOR) generate \
		-i $(BFF_OPENAPI_FILE) \
		-g typescript-fetch \
		-o $(BFF_GEN_TS_DIR) \
		-c $(ROOT_DIR)/common/api/sdk/openapi-generator/typescript-fetch.json

# Generate Go SDK.
.PHONY: sdk-go
sdk-go: doc sdk-prereq
	@echo "[sdk-go] openapi-generator go..."
	@mkdir -p $(BFF_GEN_GO_DIR)
	@$(OPENAPI_GENERATOR) generate \
		-i $(BFF_OPENAPI_FILE) \
		-g go \
		-o $(BFF_GEN_GO_DIR) \
		-c $(ROOT_DIR)/common/api/sdk/openapi-generator/go.json

.PHONY: sdk
sdk: sdk-ts sdk-go

endif
