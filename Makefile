.PHONY: dev build run test clean install air help

BINARY_NAME=go-test-api
BUILD_DIR=bin
MAIN_PATH=./cmd/main.go

dev:
	nodemon --exec "go run $(MAIN_PATH)" --ext go

build:
	go build -o $(BUILD_DIR)/$(BINARY_NAME).exe $(MAIN_PATH)

run:
	go run $(MAIN_PATH)

run-bin:
	./$(BUILD_DIR)/$(BINARY_NAME).exe

test:
	go test ./contexts/...

test-coverage:
	go test -cover ./contexts/...

clean:
	rm -rf $(BUILD_DIR)/
	rm -rf tmp/

install:
	go mod download
	go mod tidy

fmt:
	go fmt ./...

lint:
	golangci-lint run

help:
	@echo "Available commands:"
	@echo "  dev          - Run with hot reload (nodemon)"
	@echo "  build        - Build application"
	@echo "  run          - Run without building"
	@echo "  run-bin      - Run compiled binary"
	@echo "  test         - Run tests"
	@echo "  test-coverage - Run tests with coverage"
	@echo "  clean        - Clean up compiled files"
	@echo "  install      - Install dependencies"
	@echo "  fmt          - Format code"
	@echo "  lint         - Run linter"
	@echo "  help         - Show this help"