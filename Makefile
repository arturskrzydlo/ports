# Install all development tools to the `bin` directory.
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

.PHONY: generate-proto
generate-proto:
	protoc --proto_path=proto/api --go_out=internal/pb --go_opt=paths=source_relative \
		--go-grpc_out=internal/pb --go-grpc_opt=paths=source_relative  proto/api/ports.proto

.PHONY: lint
lint: ## Lint the source code.
	$(GOBIN)/golangci-lint run --config $(shell pwd)/build/.golangci.yml --verbose ./...

.PHONY: clean-integration-tests
clean-integration-tests: ## clean integration test by running down docker compose
	docker-compose -f docker-compose-test.yml --project-name ports-service-test down


.PHONY: prepare-integration-tests
prepare-integration-tests:  ## Run ports server before integration tests
	docker-compose -f docker-compose-test.yml --project-name ports-service-test up --detach --remove-orphans

.PHONY: all-tests
all-tests: prepare-integration-tests ## Run all test (unit + integration)
	$(GO) test -v -tags=integration ./...
	$(MAKE) clean-integration-tests

.PHONY: tests
tests: ## Run unit tests
	$(GO) test -v  ./...

.PHONY: tidy
tidy: ## Tidy go modules
	@go mod tidy
