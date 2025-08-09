# gateway/Makefile
.PHONY: proto-gen
proto-gen:
	@echo "🔄 Generating Gateway protocol buffers..."
	@powershell -ExecutionPolicy Bypass -File scripts/generate_proto.ps1

.PHONY: proto-clean
proto-clean:
	@echo "🧹 Cleaning Gateway generated files..."
	@if exist gen rmdir /s /q gen
	@mkdir gen > nul 2>&1

.PHONY: proto-lint
proto-lint:
	@echo "🔍 Checking proto files..."
	@for /r proto %%i in (*.proto) do @echo Found: %%i

.PHONY: run
run: proto-gen
	@echo "🚀 Starting Gateway..."
	@go run cmd/main.go

.PHONY: install-protoc
install-protoc:
	@echo "📥 Installing protoc via chocolatey..."
	@choco install protoc -y
	@echo "📥 Installing Go protoc plugins..."
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest