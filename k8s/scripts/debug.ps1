# scripts/debug.ps1
Write-Host "Debugging Kubernetes deployment..." -ForegroundColor Cyan

Write-Host "`n=== PODS STATUS ===" -ForegroundColor Yellow
kubectl get pods -o wide

Write-Host "`n=== SERVICES ===" -ForegroundColor Yellow
kubectl get services

Write-Host "`n=== PERSISTENT VOLUMES ===" -ForegroundColor Yellow
kubectl get pv,pvc

Write-Host "`n=== RECENT EVENTS ===" -ForegroundColor Yellow
kubectl get events --sort-by=.metadata.creationTimestamp | Select-Object -Last 15

Write-Host "`n=== POD DESCRIPTIONS ===" -ForegroundColor Yellow
$pods = kubectl get pods --no-headers | ForEach-Object { ($_ -split '\s+')[0] }
foreach ($pod in $pods) {
    if ($pod -like "*user-service*" -or $pod -like "*product-service*") {
        Write-Host "--- $pod ---" -ForegroundColor Cyan
        kubectl describe pod $pod | Select-Object -Last 10
    }
}

Write-Host "`n=== SERVICE LOGS ===" -ForegroundColor Yellow
Write-Host "User Service logs:" -ForegroundColor Cyan
kubectl logs deployment/user-service --tail=20 2>$null

Write-Host "`nProduct Service logs:" -ForegroundColor Cyan
kubectl logs deployment/product-service --tail=20 2>$null