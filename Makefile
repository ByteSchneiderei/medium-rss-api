.PHONY: all build clean run fmt lint test vet clean_deps install_deps jenkins all_linux deploy clear_deps_cache

BINARY_NAME := medium-rss-api
BUILD_DIR=build

VERSION := $(shell cat VERSION)
LDFLAGS := -ldflags "-X main.VERSION=${VERSION} -w"

default: build

all: clean build fmt lint test

build: vet
	go build ${LDFLAGS} -o $(BUILD_DIR)/$(BINARY_NAME) -v ./cmd/$(BINARY_NAME)

all_linux: clean fmt lint test
	CGO_ENABLED=0 GOOS=linux go build -a -tags netgo ${LDFLAGS} -o $(BUILD_DIR)/$(BINARY_NAME) -v ./cmd/$(BINARY_NAME)

linux: clean
	CGO_ENABLED=0 GOOS=linux go build -a -tags netgo ${LDFLAGS} -o $(BUILD_DIR)/$(BINARY_NAME) -v ./cmd/$(BINARY_NAME)



clean:
	go clean

run: build
	./$(BUILD_DIR)/$(BINARY_NAME)

fmt:
	gofmt -s -l . | grep -v vendor | tee /dev/stderr

lint:
	golint ./... | grep -v vendor | tee /dev/stderr

test:
	go test -short $(shell go list ./... | grep -v vendor)

vet:
	go vet $(shell go list ./... | grep -v vendor)


