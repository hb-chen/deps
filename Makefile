GOPATH:=$(shell go env GOPATH)

ifeq ($(OS),Windows_NT)
    uname_S := Windows
else
    uname_S := $(shell uname -s)
endif

# Goreleaser config
ifeq ($(uname_S), Darwin)
    goreleaser_config = .goreleaser.yml
else
	goreleaser_config = .goreleaser.yml
endif

.PHONY: release
release:
	goreleaser release --config $(goreleaser_config) --skip-validate --skip-publish --rm-dist

.PHONY: snapshot
snapshot:
	goreleaser release --config $(goreleaser_config) --skip-publish --snapshot --rm-dist

.PHONY: test
test:
	go test -race -cover -v ./...

.PHONY: lint
lint:
	golangci-lint run ./...