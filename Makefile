NAME := awsw
VERSION := $(shell git describe --tags --abbrev=0)
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -X 'main.version=$(VERSION)' \
           -X 'main.revision=$(REVISION)'

.DEFAULT_GOAL := help

.PHONY: test
test:  ## Run the tests.
	@go test ./...

.PHONY: build
build: main.go  ## Build a binary.
	GO111MODULE=on go build -ldflags "$(LDFLAGS)"

.PHONY: cross
cross: main.go  ## Build binaries for cross platform.
	mkdir -p pkg
	rm -rf pkg/*

	@# darwin
	@for arch in "amd64" "386"; do \
		GOOS=darwin GOARCH=$${arch} make build; \
		zip pkg/awsw_$(VERSION)_darwin_$${arch}.zip awsw; \
	done;
	@# linux
	@for arch in "amd64" "386" "arm64"; do \
		GOOS=linux GOARCH=$${arch} make build; \
		zip pkg/awsw_$(VERSION)_linux_$${arch}.zip awsw; \
	done;
	rm -rf awsw

.PHONY: help
help: ## Show help text
	@echo "Commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "    \033[36m%-20s\033[0m %s\n", $$1, $$2}'
