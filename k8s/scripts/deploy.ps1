# scripts/deploy.ps1
Write-Host "🚀 Deploying to local Kubernetes..." -ForegroundColor Green

Write-Host "1️⃣ Creating secrets and configmaps..." -ForegroundColor Yellow
kubectl apply -f k8s/secrets/

Write-Host "2️⃣ Deploying databases..." -ForegroundColor Yellow
kubectl apply -f k8s/services/databases.yaml

Write-Host "3️⃣ Waiting for databases to be ready..." -ForegroundColor Yellow
Write-Host "   This might take a minute..." -ForegroundColor Gray
kubectl wait --for=condition=available deployment/user-db --timeout=180s
kubectl wait --for=condition=available deployment/product-db --timeout=180s

Write-Host "4️⃣ Deploying microservices..." -ForegroundColor Yellow
kubectl apply -f k8s/services/user-service.yaml
kubectl apply -f k8s/services/product-service.yaml

Write-Host "5️⃣ Waiting for services..." -ForegroundColor Yellow
kubectl wait --for=condition=available deployment/user-service --timeout=120s
kubectl wait --for=condition=available deployment/product-service --timeout=120s

Write-Host "6️⃣ Deploying ingress..." -ForegroundColor Yellow
kubectl apply -f k8s/ingress/

Write-Host "📊 Final status:" -ForegroundColor Cyan
kubectl get pods
Write-Host ""
kubectl get services
Write-Host ""
kubectl get ingress

Write-Host "✅ Deployment complete!" -ForegroundColor Green
Write-Host "🌐 API available at: http://api.local.dev" -ForegroundColor Cyan