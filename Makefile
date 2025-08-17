.PHONY: dev-gateway dev-product build run test test-services test-shared test-coverage test-coverage-services test-coverage-shared clean install air help docs-consolidated gateway-proto-gen fmt lint

BINARY_NAME=go-test-api
BUILD_DIR=bin
GATEWAY_MAIN_PATH=./gateway/cmd/main.go
PRODUCT_MAIN_PATH=./services/product/cmd/main.go

docs-consolidated:
	@echo "🔄 Generating consolidated docs for Gateway..."
	swag init \
		--parseDependency \
		--parseInternal \
		--parseDepth 2 \
		--dir ./gateway,./services \
		--exclude ./shared,./bin,./tmp \
		-o ./gateway/docs \
		-g ./cmd/main.go
	@echo "✅ Consolidated docs generated in ./gateway/docs"

dev-gateway:
	nodemon --exec "go run $(GATEWAY_MAIN_PATH)" --ext go

dev-product:
	nodemon --exec "go run $(PRODUCT_MAIN_PATH)" --ext go

gateway-proto-gen:
	@echo "🔄 Generating protocol buffers..."
	@cd shared/infra/grpc && powershell -ExecutionPolicy Bypass -File scripts/generate_proto.ps1

build:
	go build -o $(BUILD_DIR)/$(BINARY_NAME).exe $(MAIN_PATH)

run:
	go run $(MAIN_PATH)

run-bin:
	./$(BUILD_DIR)/$(BINARY_NAME).exe

test:
	go test ./services/... ./shared/...

test-services:
	go test ./services/...

test-shared:
	go test ./shared/...

test-coverage:
	go test -cover ./services/... ./shared/...

test-coverage-services:
	go test -cover ./services/...

test-coverage-shared:
	go test -cover ./shared/...

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
	@echo "  dev-gateway  - Run gateway with hot reload"
	@echo "  dev-product  - Run product service with hot reload"
	@echo "  build        - Build application"
	@echo "  run          - Run without building"
	@echo "  run-bin      - Run compiled binary"
	@echo "  test         - Run all tests (services + shared)"
	@echo "  test-services - Run only services tests"
	@echo "  test-shared  - Run only shared tests"
	@echo "  test-coverage - Run all tests with coverage"
	@echo "  test-coverage-services - Run services tests with coverage"
	@echo "  test-coverage-shared - Run shared tests with coverage"
	@echo "  clean        - Clean up compiled files"
	@echo "  install      - Install dependencies"
	@echo "  fmt          - Format code"
	@echo "  lint         - Run linter"
	@echo "  help         - Show this help"