ROOT_DIR := $(abspath $(CURDIR)/..)

include $(ROOT_DIR)/common/build/make/common.mk

MONOLITH_DIR := $(abspath $(CURDIR))
MONOLITH_CONFIG_PROTO := monolith/internal/config/config.proto
BUF_MONOLITH_CONFIG := '{"version":"v2","modules":[{"path":"common/proto/app"},{"path":"monolith"}],"deps":["buf.build/googleapis/googleapis"],"lint":{"use":["MINIMAL"],"except":["PACKAGE_DIRECTORY_MATCH","PACKAGE_SAME_DIRECTORY"]},"breaking":{"use":["FILE"]}}'

IGNORE_ERROR ?= 0

ifeq ($(IGNORE_ERROR),1)
run = @$(1) || echo "[ERROR] $(2) failed" >&2
else
run = @$(1) || { echo "[ERROR] $(2) failed" >&2; exit 1; }
endif

.PHONY: tidy
tidy:
	@echo "[tidy] go mod tidy..."
	$(call run,go mod tidy,[tidy] go mod tidy)

.PHONY: format
format:
	@echo "[format] gofmt monolith Go files..."
	$(call run,find . -type f -name "*.go" -exec gofmt -w {} +,[format] gofmt)
	@echo "[format] format monolith proto..."
	$(call run,cd $(ROOT_DIR) && $(BUF) format -w monolith,[format] buf format)

.PHONY: config-clean
config-clean:
	@echo "[config-clean] cleaning monolith config generated Go files..."
	@find internal/config -type f \( -name "*.pb.go" -o -name "*.pb.validate.go" -o -name "*.pb.gw.go" \) -delete; true

.PHONY: config
config: config-clean
	@echo "[config] buf generate monolith config..."
	$(call run,cd $(ROOT_DIR) && $(BUF) generate --config $(BUF_MONOLITH_CONFIG) --template $(BUF_GEN_CONFIG) --path $(MONOLITH_CONFIG_PROTO) --output monolith,[config] buf generate)

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
	$(call run,go mod tidy && go mod download && go build -trimpath -ldflags "-s -w" -o $(MONOLITH_DIR)/server ./cmd/...,[build] go build)

.PHONY: build-clean
build-clean:
	@echo "[build-clean] cleaning monolith binary..."
	@rm -f $(MONOLITH_DIR)/server; true

.PHONY: clean
clean: config-clean
clean: wire-clean
clean: build-clean

.PHONY: gen
gen: config wire

.PHONY: all
all: init api tidy gen build
