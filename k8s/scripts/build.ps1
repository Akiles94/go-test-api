# scripts/build.ps1
Write-Host "Building microservices..." -ForegroundColor Green

# Cambiar al directorio del proyecto
$projectDir = Split-Path $PSScriptRoot -Parent | Split-Path -Parent
Write-Host "Project directory: $projectDir" -ForegroundColor Gray

Write-Host "Building User Service..." -ForegroundColor Yellow
docker build -t user-service:local --build-arg SERVICE_PATH=services/user -f "$projectDir/Dockerfile" $projectDir

Write-Host "Building Product Service..." -ForegroundColor Yellow
docker build -t product-service:local --build-arg SERVICE_PATH=services/product -f "$projectDir/Dockerfile" $projectDir

Write-Host "Build complete!" -ForegroundColor Green