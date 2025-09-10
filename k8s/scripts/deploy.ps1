# scripts/deploy.ps1
Write-Host "Deploying to local Kubernetes..." -ForegroundColor Green

Write-Host "Step 1: Creating secrets and configmaps..." -ForegroundColor Yellow
kubectl apply -f secrets/

Write-Host "Step 2: Deploying databases..." -ForegroundColor Yellow
kubectl apply -f services/databases.yaml

Write-Host "Step 3: Waiting for databases to be ready..." -ForegroundColor Yellow
Write-Host "   This might take a minute..." -ForegroundColor Gray
kubectl wait --for=condition=available deployment/user-db --timeout=180s
kubectl wait --for=condition=available deployment/product-db --timeout=180s

Write-Host "Step 4: Deploying microservices..." -ForegroundColor Yellow
kubectl apply -f services/user-service.yaml
kubectl apply -f services/product-service.yaml

Write-Host "Step 5: Waiting for services..." -ForegroundColor Yellow
Write-Host "   Checking current pod status..." -ForegroundColor Gray
kubectl get pods

Write-Host "   Waiting for user-service..." -ForegroundColor Gray
$userServiceReady = kubectl wait --for=condition=available deployment/user-service --timeout=120s 2>&1
if ($LASTEXITCODE -ne 0) {
    Write-Host "User service timeout. Checking status..." -ForegroundColor Red
    kubectl describe deployment/user-service
    kubectl logs deployment/user-service --tail=10
}

Write-Host "   Waiting for product-service..." -ForegroundColor Gray
$productServiceReady = kubectl wait --for=condition=available deployment/product-service --timeout=120s 2>&1
if ($LASTEXITCODE -ne 0) {
    Write-Host "Product service timeout. Checking status..." -ForegroundColor Red
    kubectl describe deployment/product-service
    kubectl logs deployment/product-service --tail=10
}

Write-Host "Step 6: Deploying ingress..." -ForegroundColor Yellow
kubectl apply -f ingress/

Write-Host "Final status:" -ForegroundColor Cyan
kubectl get pods
Write-Host ""
kubectl get services
Write-Host ""
kubectl get ingress

if ($LASTEXITCODE -eq 0) {
    Write-Host "Deployment complete!" -ForegroundColor Green
    Write-Host "API available at: http://api.local.dev" -ForegroundColor Cyan
} else {
    Write-Host "Deployment had issues. Run 'make debug' for more info." -ForegroundColor Red
}