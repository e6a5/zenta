# Makefile for zenta

# Build variables
BINARY_NAME=zenta
VERSION?=$(shell git describe --tags --always --dirty)
COMMIT=$(shell git rev-parse HEAD)
BUILD_DATE=$(shell date -u +%Y-%m-%dT%H:%M:%SZ)
LDFLAGS=-s -w -X 'github.com/e6a5/zenta/internal/version.Version=$(VERSION)' -X 'github.com/e6a5/zenta/internal/version.GitCommit=$(COMMIT)' -X 'github.com/e6a5/zenta/internal/version.BuildDate=$(BUILD_DATE)'

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod

.PHONY: all build clean test deps lint install help

## Build the binary
build:
	$(GOBUILD) -ldflags="$(LDFLAGS)" -o $(BINARY_NAME) .

## Build for all platforms
build-all:
	mkdir -p dist
	GOOS=linux GOARCH=amd64 $(GOBUILD) -ldflags="$(LDFLAGS)" -o dist/$(BINARY_NAME)-linux-amd64 .
	GOOS=linux GOARCH=arm64 $(GOBUILD) -ldflags="$(LDFLAGS)" -o dist/$(BINARY_NAME)-linux-arm64 .
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -ldflags="$(LDFLAGS)" -o dist/$(BINARY_NAME)-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 $(GOBUILD) -ldflags="$(LDFLAGS)" -o dist/$(BINARY_NAME)-darwin-arm64 .
	GOOS=windows GOARCH=amd64 $(GOBUILD) -ldflags="$(LDFLAGS)" -o dist/$(BINARY_NAME)-windows-amd64.exe .
	GOOS=freebsd GOARCH=amd64 $(GOBUILD) -ldflags="$(LDFLAGS)" -o dist/$(BINARY_NAME)-freebsd-amd64 .

## Run tests
test:
	$(GOTEST) -v -race -coverprofile=coverage.out ./...

## Run tests with coverage report
test-coverage: test
	$(GOCMD) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

## Run benchmarks
bench:
	$(GOTEST) -bench=. -benchmem ./...

## Clean build artifacts
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -rf dist/
	rm -f coverage.out coverage.html

## Download dependencies
deps:
	$(GOMOD) download
	$(GOMOD) verify

## Tidy dependencies
tidy:
	$(GOMOD) tidy

## Run linting
lint:
	golangci-lint run

## Install the binary
install: build
	mv $(BINARY_NAME) $(GOPATH)/bin/

## Install to /usr/local/bin (requires sudo)
install-system: build
	sudo mv $(BINARY_NAME) /usr/local/bin/

## Format code
fmt:
	$(GOCMD) fmt ./...

## Vet code
vet:
	$(GOCMD) vet ./...

## Run all checks (test, lint, vet)
check: test lint vet

## Run the application
run: build
	./$(BINARY_NAME)

## Display help
help:
	@echo "Available targets:"
	@awk '/^##/{c=substr($$0,3);next}c&&/^[[:alpha:]][[:alnum:]_-]+:/{print substr($$1,1,index($$1,":")-1)":"c}1{c=""}' $(MAKEFILE_LIST) | column -t -s ':'

## Default target
all: clean deps test build

.DEFAULT_GOAL := help 