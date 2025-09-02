#!/bin/bash
# scripts/debug.sh
echo "🔍 Debugging Kubernetes deployment..."

echo "=== PODS STATUS ==="
kubectl get pods -o wide

echo -e "\n=== SERVICES ==="
kubectl get services

echo -e "\n=== INGRESS ==="
kubectl get ingress

echo -e "\n=== PERSISTENT VOLUMES ==="
kubectl get pv,pvc

echo -e "\n=== EVENTS (last 10 minutes) ==="
kubectl get events --sort-by=.metadata.creationTimestamp | tail -20

echo -e "\n=== POD LOGS ==="
echo "User Service logs:"
kubectl logs deployment/user-service --tail=10

echo -e "\nProduct Service logs:"
kubectl logs deployment/product-service --tail=10

echo -e "\nUser DB logs:"
kubectl logs deployment/user-db --tail=10

echo -e "\nProduct DB logs:"
kubectl logs deployment/product-db --tail=10