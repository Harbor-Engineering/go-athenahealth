all: tidy format build test

build:
	go build ./...

format:
	go fmt ./...

test:
	go test ./... -race

tidy:
	go mod tidy

# Release management targets
.PHONY: setup-hooks release

setup-hooks:
	@echo "Setting up git hooks..."
	@./scripts/setup-hooks.sh

release:
	@if [ -z "$(VERSION)" ]; then \
		echo "Error: VERSION is required"; \
		echo "Usage: make release VERSION=v0.0.4"; \
		exit 1; \
	fi
	@./scripts/release.sh $(VERSION)
