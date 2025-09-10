# scripts/clean.ps1
Write-Host "Cleaning up Kubernetes resources..." -ForegroundColor Yellow

Write-Host "Deleting ingress..." -ForegroundColor Gray
kubectl delete -f ingress/ --ignore-not-found=true

Write-Host "Deleting services..." -ForegroundColor Gray
kubectl delete -f services/ --ignore-not-found=true

Write-Host "Deleting secrets and configmaps..." -ForegroundColor Gray
kubectl delete -f secrets/ --ignore-not-found=true

Write-Host "Force deleting persistent volume claims..." -ForegroundColor Gray
kubectl delete pvc --all --force --grace-period=0 --ignore-not-found=true

Write-Host "Waiting for cleanup to complete..." -ForegroundColor Gray
Start-Sleep -Seconds 5

Write-Host "Final status:" -ForegroundColor Cyan
kubectl get pods,services,pvc 2>$null

Write-Host "Cleanup complete!" -ForegroundColor Green