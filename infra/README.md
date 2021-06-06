# Para implantar a api-myapp no Kubernetes com Persistência em disco

```
# diretorio /api-myapp/infra
cd infra

kubectl apply -f https://raw.githubusercontent.com/kubernetes/kops/master/addons/metrics-server/v1.8.x.yaml

kubectl delete -f deploy-postgres.yml
kubectl delete configmap api-myapp-config 
kubectl delete -f deploy-api-myapp.yml

kubectl apply -f deploy-postgres.yml
kubectl create configmap api-myapp-config --from-env-file=../.env.kubernetes
kubectl apply -f deploy-api-myapp.yml

# diretorio /web-myapp/infra
# kubectl create configmap web-myapp-config --from-env-file=../.env.kubernetes
# kubectl apply -f deploy-web-myapp.yml


# Para testar no Kubernetes
kubectl port-forward services/api-myapp 9990:9990


```



# Kubernetes Metrics Server - Somente Metricas Instantaneas 
```
kubectl apply -f https://raw.githubusercontent.com/kubernetes/kops/master/addons/metrics-server/v1.8.x.yaml
```

Exemplo de utilização no HorizontalPodAutoscaler
```
apiVersion: autoscaling/v1
kind: HorizontalPodAutoscaler
metadata:
  name: api-order
spec:
  scaleTargetRef:
    apiVersion: apps/v2beta1
    kind: Deployment
    name: api-order
  minReplicas: 1
  maxReplicas: 2
  targetCPUUtilizationPercentage: 90
apiVersion: autoscaling/v1
kind: HorizontalPodAutoscaler
metadata:
  name: api-order
spec:
  scaleTargetRef:
    apiVersion: apps/v2beta1
    kind: Deployment
    name: api-order
  minReplicas: 1
  maxReplicas: 2
  targetCPUUtilizationPercentage: 90
```  