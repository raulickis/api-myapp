# Para implantar a api-myapp no Kubernetes com Persistência em disco

```
cd infra

kubectl apply -f deploy-postgres.yml
kubectl delete configmap api-myapp-config 
kubectl create configmap api-myapp-config --from-env-file=../.env.kubernetes
kubectl apply -f deploy-api-myapp.yml
```


# Prometheus

### First, we will create a Kubernetes namespace for all our monitoring components. If you don’t create a dedicated namespace, all the Prometheus kubernetes deployment objects get deployed on the default namespace.
```
kubectl create namespace monitoring
```

### Prometheus uses Kubernetes APIs to read all the available metrics from Nodes, Pods, Deployments, etc. For this reason, we need to create an RBAC policy with read access to required API groups and bind the policy to the monitoring namespace.
```
kubectl apply -f prometheus-clusterrole.yml
```

### Prometheus uses Kubernetes APIs to read all the available metrics from Nodes, Pods, Deployments, etc. For this reason, we need to create an RBAC policy with read access to required API groups and bind the policy to the monitoring namespace.
```
kubectl apply -f prometheus-configmap.yml
```

### Note: This deployment uses the latest official Prometheus image from the docker hub. Also, we are not using any persistent storage volumes for Prometheus storage as it is a basic setup. When setting up Prometheus for production uses cases, make sure you add persistent storage to the deployment.
```
kubectl apply -f prometheus-deploy.yml
```


# Kubernetes Metrics for Prometheus
Observação, no deployment mudar do seletor do Node para `kubernetes.io/os: linux`
```
kubectl apply -f kube-state-metrics/
```


# Alert Manager 
### Alert Manager reads its configuration from a config.yaml file. It contains the configuration of alert template path, email and other alert receiving configuration.
### In this setup, we are using email and slack webhook receivers. 
```
kubectl apply -f prometheus-alertmanager-configmap.yml
```

### We need alert templates for all the receivers we use (email, slack etc). Alert manager will dynamically substitute the values and delivers alerts to the receivers based on the template. You can customize these templates based on your needs.
```
kubectl apply -f prometheus-alertmanager-configmap-template.yml
```

### We need alert templates for all the receivers we use (email, slack etc). Alert manager will dynamically substitute the values and delivers alerts to the receivers based on the template. You can customize these templates based on your needs.
```
kubectl apply -f prometheus-alertmanager-deploy.yml
```


# Grafana

### Note: The following data source configuration is for Prometheus. If you have more data sources, you can add more data sources with different YAMLs under the data section.
```
kubectl apply -f prometheus-grafana-datasource-config.yml
```

### Note: This Grafana deployment does not use a persistent volume. If you restart the pod all changes will be gone. Use a persistent volume if you are deploying Grafana for your project requirements. It will persist all the configs and data that Grafana uses.
```
kubectl apply -f prometheus-grafana-deploy.yml
```


# Prometheus Node Exporter

### To get all the kubernetes node-level system metrics, you need to have a node-exporter running in all the kubernetes nodes. It collects all the Linux system metrics and exposes them via /metrics endpoint on port 9100
```
kubectl apply -f prometheus-node-exporter-daemonset.yml
```


# Observações

Para ver o dashboard do Prometheus: 
minikube service --url prometheus-service -n monitoring
Para ver como fazer queries: https://prometheus.io/docs/prometheus/latest/querying/basics/

Para ver o dashboard do Alert Manager: 
minikube service --url alertmanager -n monitoring

Para ver o dashboard do Grafana: 
minikube service --url grafana -n monitoring
Credentials: admin / admin
Import dashboard 8588 and asign it to Prometheus
Import dashboard 1860 and asign it to Prometheus


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


# Monitor Linux Servers Using Prometheus
Somente para sistemas Linux que estão na mesma rede (ou SG), mas fora do Kubernetes 
https://devopscube.com/monitor-linux-servers-prometheus-node-exporter/


# Material de apoio para o Prometheus

### https://devopscube.com/setup-prometheus-monitoring-on-kubernetes/
### https://devopscube.com/setup-kube-state-metrics/
### https://devopscube.com/setup-grafana-kubernetes/
### https://devopscube.com/node-exporter-kubernetes/
### https://devopscube.com/alert-manager-kubernetes-guide/
### https://devopscube.com/monitor-linux-servers-prometheus-node-exporter/



