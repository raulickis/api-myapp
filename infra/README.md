# Para implantar

cd infra

kubectl apply -f deploy-postgres.yml
kubectl delete configmap api-myapp-config 
kubectl create configmap api-myapp-config --from-env-file=../.env.kubernetes
kubectl apply -f deploy-api-myapp.yml