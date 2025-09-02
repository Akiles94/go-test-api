#!/bin/bash
# scripts/clean.sh
echo "🧹 Cleaning up Kubernetes resources..."

kubectl delete -f k8s/ingress/ --ignore-not-found=true
kubectl delete -f k8s/services/ --ignore-not-found=true
kubectl delete -f k8s/secrets/ --ignore-not-found=true

# Force delete stuck resources
kubectl delete pvc --all --force --grace-period=0 --ignore-not-found=true

echo "✅ Cleanup complete!"