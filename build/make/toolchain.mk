ifndef TOOLCHAIN_MK_INCLUDED
TOOLCHAIN_MK_INCLUDED := 1

# Fixed versions for code generators and build tools. Changes here rebuild the CI toolchain image.
.PHONY: init
init:
	@echo "[init] installing pinned development tools..."
	@set -e; \
	for tool in \
		google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.11 \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1 \
		github.com/go-kratos/kratos/cmd/kratos/v3@v3.0.0-20260626125723-668db92c2c00 \
		github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v3@v3.0.0-20260626125723-668db92c2c00 \
		github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v3@v3.0.0-20260626125723-668db92c2c00 \
		github.com/envoyproxy/protoc-gen-validate@v1.3.3 \
		github.com/google/gnostic/cmd/protoc-gen-openapi@v0.7.1 \
		github.com/google/wire/cmd/wire@v0.7.0 \
		entgo.io/ent/cmd/ent@v0.14.5 \
		github.com/bufbuild/buf/cmd/buf@v1.50.0; do \
		for attempt in 1 2 3; do \
			if go install "$$tool"; then break; fi; \
			if [ "$$attempt" -eq 3 ]; then exit 1; fi; \
			sleep "$$attempt"; \
		done; \
	done

endif
