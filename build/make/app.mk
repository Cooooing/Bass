ifndef APP_MK_INCLUDED
APP_MK_INCLUDED := 1

# 引入一次性目标和共享变量。
MAKE_DIR := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))
include $(MAKE_DIR)/common.mk

# --- 模块变量 ---
SERVER := $(notdir $(CURDIR))
APP_DIR := $(ROOT_DIR)/app/$(SERVER)
INTERNAL_DIR := $(APP_DIR)/internal
BUF_SERVICE_CONFIG := '{"version":"v2","modules":[{"path":"common/proto/app"},{"path":"app"}],"deps":["buf.build/googleapis/googleapis"]}'
LINT_BINARY := $(ROOT_DIR)/.golangci-plugin/bin/bass-golangci-lint$(shell go env GOEXE)
# 使用 Git 跟踪文件生成依赖列表，避免 Windows 的 find.exe 与 Unix find 参数不兼容。
# glob 路径限定在插件目录，防止 app 源码变更触发定制 linter 重建。
LINT_PLUGIN_SOURCES := $(addprefix $(ROOT_DIR)/,$(shell git -C $(ROOT_DIR) ls-files -- .golangci-plugin/go.mod .golangci-plugin/go.sum ':(glob).golangci-plugin/**/*.go'))

IGNORE_ERROR ?= 0
MODULE_GEN_TARGETS := cfg wire

# run 在解析阶段根据 IGNORE_ERROR 选择错误处理方式。
# 用法：$(call run,command,error-label)
# 命令默认隐藏，失败时向 stderr 输出标签。
ifeq ($(IGNORE_ERROR),1)
run = @$(1) || echo "[ERROR] $(2) failed" >&2
else
run = @$(1) || { echo "[ERROR] $(2) failed" >&2; exit 1; }
endif

# --- 模块目标 ---


.PHONY: cfg-clean
cfg-clean:
	@echo "[cfg-clean] cleaning proto generated Go files..."
	@cd $(APP_DIR) && find . -type f \( -name "*.pb.go" -o -name "*.pb.validate.go" -o -name "*.pb.gw.go" \) -delete; true

.PHONY: cfg
cfg: cfg-clean
	@echo "[cfg] buf generate..."
	$(call run,cd $(ROOT_DIR) && $(BUF) generate --config $(BUF_SERVICE_CONFIG) --template $(BUF_GEN_CONFIG) --path app/$(SERVER)/internal/config/config.proto --output app,[cfg] buf generate)

.PHONY: wire-clean
wire-clean:
	@echo "[wire-clean] cleaning wire_gen.go..."
	@find $(APP_DIR) -type f -name "wire_gen.go" -delete; true

.PHONY: wire
wire: wire-clean
	@echo "[wire] wire..."
	@find $(APP_DIR) -type f -name "wire.go" -exec dirname {} \; | sort -u | while read dir; do \
		(cd "$$dir" && wire) || exit 1; \
	done

.PHONY: build
build:
	@echo "[build] go build..."
	$(call run,cd $(APP_DIR) && go build -trimpath -ldflags "-s -w" -o $(APP_DIR)/server ./cmd/...,[build] go build)

.PHONY: build-clean
build-clean:
	@echo "[build-clean] cleaning service binary..."
	@rm -f $(APP_DIR)/server; true

.PHONY: gen-clean
gen-clean: cfg-clean
gen-clean: wire-clean

.PHONY: gen
gen: gen-clean
	@for target in $(MODULE_GEN_TARGETS); do \
		$(MAKE) $$target IGNORE_ERROR=$(IGNORE_ERROR) || exit 1; \
	done


# 插件源码、模块依赖或 linter 配置变更后，自动重建定制二进制。
.PHONY: lint-toolchain
lint-toolchain: $(LINT_BINARY)

$(LINT_BINARY): $(ROOT_DIR)/.custom-gcl.yml $(ROOT_DIR)/.golangci.yml $(LINT_PLUGIN_SOURCES)
	@echo "[lint] building bass-golangci-lint v2.14.0..."
	@go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.14.0
	@cd $(ROOT_DIR) && $(subst \,/,$(shell go env GOPATH))/bin/golangci-lint$(shell go env GOEXE) custom

.PHONY: lint
lint: lint-toolchain
	@echo "[lint] $(SERVER)..."
	@cd $(APP_DIR) && $(LINT_BINARY) run --allow-parallel-runners --disable=bassarch --fix --issues-exit-code=0 ./... || echo "[lint] fix pass did not complete; continuing strict verification"
	@cd $(APP_DIR) && $(LINT_BINARY) run --allow-parallel-runners --disable=bassarch ./...

endif
