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
.PHONY: setup-hooks release github-release

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

github-release:
	@if [ -z "$(VERSION)" ]; then \
		echo "Error: VERSION is required"; \
		echo "Usage: make github-release VERSION=v0.0.4"; \
		exit 1; \
	fi
	@if ! command -v gh >/dev/null 2>&1; then \
		echo "Error: GitHub CLI (gh) is not installed"; \
		echo "Install it from: https://cli.github.com/"; \
		exit 1; \
	fi
	@echo "Creating GitHub release for $(VERSION)..."
	@gh release create $(VERSION) \
		--title "Release $(VERSION)" \
		--generate-notes \
		--verify-tag
	@echo "✓ GitHub release $(VERSION) created successfully!"
	@echo "View at: https://github.com/eleanorhealth/go-athenahealth/releases/tag/$(VERSION)"
