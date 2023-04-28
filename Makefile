# Install all development tools and build artifacts to the project's
# `bin` directory.
export GOBIN=$(CURDIR)/bin

# Default to the system 'go'.
GO?=$(shell which go)

$(GOBIN):
	mkdir -p $(GOBIN)

.PHONY: setup
setup: ## Setting up local env
	cd $(GOBIN)
	wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.51.2

.PHONY: clean
clean: ## Remove build artifacts.
	rm -rf $(GOBIN)

.PHONY: lint
lint: ## Lint the source code.
	$(GOBIN)/golangci-lint run --config $(shell pwd)/build/.golangci.yml --verbose ./...

.PHONY: tests
tests: ## Run unit tests
	$(GO) test -v  ./...

.PHONY: tidy
tidy: ## Tidy go modules and re-vendor
	@go mod tidy
