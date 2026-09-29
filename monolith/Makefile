ROOT_DIR := $(abspath $(CURDIR)/..)

include $(ROOT_DIR)/build/make/common.mk

MONOLITH_DIR := $(abspath $(CURDIR))
MONOLITH_CONFIG_PROTO := monolith/internal/config/config.proto
BUF_MONOLITH_CONFIG := '{"version":"v2","modules":[{"path":"common/proto/app"},{"path":"monolith"}],"deps":["buf.build/googleapis/googleapis"],"lint":{"use":["MINIMAL"],"except":["PACKAGE_DIRECTORY_MATCH","PACKAGE_SAME_DIRECTORY"]},"breaking":{"use":["FILE"]}}'

IGNORE_ERROR ?= 0

ifeq ($(IGNORE_ERROR),1)
run = @$(1) || echo "[ERROR] $(2) failed" >&2
else
run = @$(1) || { echo "[ERROR] $(2) failed" >&2; exit 1; }
endif

.PHONY: cfg-clean
cfg-clean:
	@echo "[cfg-clean] cleaning monolith config generated Go files..."
	@find internal/config -type f \( -name "*.pb.go" -o -name "*.pb.validate.go" -o -name "*.pb.gw.go" \) -delete; true

.PHONY: cfg
cfg: cfg-clean
	@echo "[cfg] buf generate monolith config..."
	$(call run,cd $(ROOT_DIR) && $(BUF) generate --config $(BUF_MONOLITH_CONFIG) --template $(BUF_GEN_CONFIG) --path $(MONOLITH_CONFIG_PROTO) --output monolith,[cfg] buf generate)

.PHONY: wire-clean
wire-clean:
	@echo "[wire-clean] cleaning monolith wire_gen.go..."
	@find . -type f -name "wire_gen.go" -delete; true

.PHONY: wire
wire: wire-clean
	@echo "[wire] wire monolith..."
	@find . -type f -name "wire.go" -exec dirname {} \; | sort -u | while read dir; do \
		(cd "$$dir" && wire) || exit 1; \
	done

.PHONY: build
build:
	@echo "[build] go build monolith..."
	$(call run,go build -trimpath -ldflags "-s -w" -o $(MONOLITH_DIR)/server ./cmd/...,[build] go build)

.PHONY: build-clean
build-clean:
	@echo "[build-clean] cleaning monolith binary..."
	@rm -f $(MONOLITH_DIR)/server; true

.PHONY: gen-clean
gen-clean: cfg-clean wire-clean

.PHONY: gen
gen: cfg wire
