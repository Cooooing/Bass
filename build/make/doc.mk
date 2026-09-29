ifndef DOC_MK_INCLUDED
DOC_MK_INCLUDED := 1

# Requires app.mk for PROTO_DIR and PROTO_GEN_DIR.

BFF_SERVER := $(notdir $(CURDIR))
BFF_PROTO_DIR := $(PROTO_DIR)/$(BFF_SERVER)
BFF_OPENAPI_DIR := $(COMMON_DIR)/proto/gen-openapi/$(BFF_SERVER)
BFF_OPENAPI_FILE := $(BFF_OPENAPI_DIR)/openapi.yaml
BFF_GEN_TYPESCRIPT_AXIOS_DIR := $(COMMON_DIR)/proto/gen-sdk/typescript-axios/$(BFF_SERVER)
BFF_GEN_TYPESCRIPT_FETCH_DIR := $(COMMON_DIR)/proto/gen-sdk/typescript-fetch/$(BFF_SERVER)
BFF_GEN_GO_DIR := $(COMMON_DIR)/proto/gen-sdk/go/$(BFF_SERVER)
BFF_GEN_JAVA_DIR := $(COMMON_DIR)/proto/gen-sdk/java/$(BFF_SERVER)
BFF_GEN_RUST_DIR := $(COMMON_DIR)/proto/gen-sdk/rust/$(BFF_SERVER)
OPENAPI_GENERATOR_CLI_VERSION ?= 2.41.0
OPENAPI_GENERATOR ?= cd $(ROOT_DIR)/common/proto/sdk && npx --yes @openapitools/openapi-generator-cli@$(OPENAPI_GENERATOR_CLI_VERSION)
SDK_SHORT_NAME_OPTS := --remove-operation-id-prefix --additional-properties=apiNameSuffix=
SDK_LANGUAGES ?= typescript-axios typescript-fetch go java rust

# CI uses the same SDK language list as the local generation target.
.PHONY: ci-sdk-languages
ci-sdk-languages:
	@echo $(SDK_LANGUAGES)

# Append to composite target sequence.
MODULE_GEN_TARGETS += doc
gen-clean: doc-clean

# Clean OpenAPI artifacts.
.PHONY: doc-clean
doc-clean:
	@echo "[doc-clean] cleaning OpenAPI documents..."
	@rm -rf $(BFF_OPENAPI_DIR) 2>/dev/null; true

# Generate OpenAPI document for the current BFF service.
.PHONY: doc
doc: doc-clean
	@echo "[doc] buf generate openapi..."
	@mkdir -p $(BFF_OPENAPI_DIR)
	@cd $(ROOT_DIR) && $(BUF) generate $(BUF_CONFIG_DIR) --path common/proto/app/$(BFF_SERVER) --template $(BUF_GEN_OPENAPI) --output $(BFF_OPENAPI_DIR)

# Clean generated SDK artifacts.
.PHONY: sdk-clean
sdk-clean:
	@echo "[sdk-clean] cleaning generated SDKs..."
	@rm -rf $(BFF_GEN_TYPESCRIPT_AXIOS_DIR) $(BFF_GEN_TYPESCRIPT_FETCH_DIR) $(BFF_GEN_GO_DIR) $(BFF_GEN_JAVA_DIR) $(BFF_GEN_RUST_DIR) 2>/dev/null; true

.PHONY: sdk
sdk: sdk-clean doc
	@command -v npx >/dev/null 2>&1 || { echo "[ERROR] [sdk] OpenAPI Generator requires npx" >&2; exit 1; }
	@command -v java >/dev/null 2>&1 || { echo "[ERROR] [sdk] OpenAPI Generator CLI requires Java" >&2; exit 1; }
	@for language in $(SDK_LANGUAGES); do \
		case "$$language" in \
			typescript-axios) output="$(BFF_GEN_TYPESCRIPT_AXIOS_DIR)" ;; \
			typescript-fetch) output="$(BFF_GEN_TYPESCRIPT_FETCH_DIR)" ;; \
			go) output="$(BFF_GEN_GO_DIR)" ;; \
			java) output="$(BFF_GEN_JAVA_DIR)" ;; \
			rust) output="$(BFF_GEN_RUST_DIR)" ;; \
			*) echo "[ERROR] unsupported SDK language: $$language" >&2; exit 1 ;; \
		esac; \
		echo "[sdk] openapi-generator $$language..."; \
		mkdir -p "$$output"; \
		$(OPENAPI_GENERATOR) generate \
			-i $(BFF_OPENAPI_FILE) \
			-g "$$language" \
			-o "$$output" \
			-c $(ROOT_DIR)/common/proto/sdk/openapi-generator/$$language.json \
			$(SDK_SHORT_NAME_OPTS) || exit 1; \
	done

endif
