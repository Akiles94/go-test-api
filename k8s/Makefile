# Makefile for Windows
.PHONY: build deploy test clean

build:
	@powershell -ExecutionPolicy Bypass -File scripts/build.ps1

deploy: build
	@powershell -ExecutionPolicy Bypass -File scripts/deploy.ps1

test:
	@powershell -ExecutionPolicy Bypass -File scripts/test.ps1

clean:
	@kubectl delete -f . --recursive --ignore-not-found=true

logs-user:
	@kubectl logs -f deployment/user-service

logs-product:
	@kubectl logs -f deployment/product-service

status:
	@kubectl get pods,services,ingress

quick-test:
	@powershell -Command "try { Invoke-WebRequest -Uri 'http://api.local.dev/health' -UseBasicParsing | Out-Host } catch { Write-Host 'API not responding' }"