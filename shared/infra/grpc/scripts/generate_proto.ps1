# gateway/scripts/generate_proto.ps1
Write-Host "Generating gRPC code for Gateway..." -ForegroundColor Cyan

# Create output directory
if (!(Test-Path "gen")) {
    New-Item -ItemType Directory -Path "gen" -Force
}

# Check if protoc is installed
if (!(Get-Command protoc -ErrorAction SilentlyContinue)) {
    Write-Host "ERROR: protoc is not installed or not in PATH" -ForegroundColor Red
    Write-Host "Please run: make install-protoc-winget" -ForegroundColor Yellow
    exit 1
}

# Find all proto files
$protoFiles = Get-ChildItem -Path "proto" -Filter "*.proto" -Recurse

if ($protoFiles.Count -eq 0) {
    Write-Host "WARNING: No .proto files found in proto directory" -ForegroundColor Yellow
    Write-Host "Current directory: $(Get-Location)" -ForegroundColor Yellow
    Write-Host "Looking in: $(Resolve-Path 'proto' -ErrorAction SilentlyContinue)" -ForegroundColor Yellow
    exit 0
}

Write-Host "Found proto files:" -ForegroundColor Green
foreach ($file in $protoFiles) {
    Write-Host "  - $($file.Name)" -ForegroundColor White
}

# Convert to relative paths
$protoFilePaths = @()
foreach ($file in $protoFiles) {
    $relativePath = $file.FullName.Replace((Get-Location).Path + "\", "").Replace("\", "/")
    $protoFilePaths += $relativePath
}

try {
    & protoc `
        --proto_path=proto `
        --go_out=gen `
        --go_opt=paths=source_relative `
        --go-grpc_out=gen `
        --go-grpc_opt=paths=source_relative `
        $protoFilePaths

    Write-Host "SUCCESS: Gateway gRPC code generated successfully" -ForegroundColor Green
} catch {
    Write-Host "ERROR: Error generating protobuf code: $_" -ForegroundColor Red
    exit 1
}