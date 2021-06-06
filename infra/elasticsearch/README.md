# Elasticsearch Operator  

ECK is a new orchestration product based on the Kubernetes Operator pattern that lets users provision, manage, and operate Elasticsearch clusters on Kubernetes.

- Referência básica do ECK:
  - https://www.elastic.co/guide/en/cloud-on-k8s/current/index.html
  - https://github.com/elastic/cloud-on-k8s.git

- Pontos cobertos:
  - Instalação do ECK Operator
  - Elasticsearch para armazenar métricas
  - Kibana para controle geral
  - Beats (coletores de métricas)

- Pendencias:
  - Logs no Kafka

- PAREI AINDA NA CONFIGURACAO METRICS BEATS E HEART BEATS (EXEMPLO E TUTORIAL)


## Install Operator (Kubernetes v12+)
```
kubectl apply -f https://download.elastic.co/downloads/eck/1.6.0/all-in-one.yaml
kubectl -n elastic-system logs -f statefulset.apps/elastic-operator
```

## Install Elasticsearch & Kibana
```
# Para instalar
kubectl apply -f elasticsearch.yaml
kubectl apply -f kibana.yaml
kubectl apply -f filebeat.yaml
kubectl apply -f metricbeat.yaml
kubectl apply -f elastic-apm-server.yaml
kubectl apply -f pet-clinicic-sample-app.yaml

# Validando APM Server
kubectl get service --selector='common.k8s.elastic.co/type=apm-server'
kubectl get pods --selector='apm.k8s.elastic.co/name=elastic-apm-server'
# Token de acesso ao APM Server
kubectl get secret/elastic-apm-server-apm-token -o go-template='{{index .data "secret-token" | base64decode}}'
#  Logs do file beat
kubectl port-forward services/elasticsearch-es-http 9200:9200
curl http://localhost:9200/filebeat-*/_search


# Ver o tipo de licença
curl -X GET "localhost:9200/_license"

# Validar Aplicacao exemplo
kubectl port-forward svc/petclinic -n default 8080:8080

# Para testar
kubectl get elasticsearch
kubectl port-forward services/elasticsearch-es-http 9200:9200
kubectl port-forward services/kibana-kb-http 5601:5601
kubectl port-forward services/elastic-apm-server-apm-http 8200:8200


# Para limpar tudo
kubectl delete -f elasticsearch.yaml
kubectl delete -f kibana.yaml
kubectl delete -f filebeat.yaml
kubectl delete -f metricbeat.yaml
kubectl delete -f elastic-apm-server.yaml
kubectl delete -f pet-clinicic-sample-app.yaml
kubectl delete -f https://download.elastic.co/downloads/eck/1.6.0/all-in-one.yaml
```


kubectl create -f filebeat-setup.yml

kubectl delete -f filebeat-setup.yml

kubectl apply -f filebeat-setup.yml


## Reference Links

- Exemplos oficiais (https://www.elastic.co/guide/en/cloud-on-k8s/current/k8s-beat-configuration-examples.html):
  - Metricbeats: https://raw.githubusercontent.com/elastic/cloud-on-k8s/1.6/config/recipes/beats/metricbeat_hosts.yaml
  - FileBeat com autodiscover: https://raw.githubusercontent.com/elastic/cloud-on-k8s/1.6/config/recipes/beats/filebeat_autodiscover.yaml
  - FileBeat com autodiscover em metadata: https://raw.githubusercontent.com/elastic/cloud-on-k8s/1.6/config/recipes/beats/filebeat_autodiscover_by_metadata.yaml
  - FileBeat com Elastic e Kibana: https://raw.githubusercontent.com/elastic/cloud-on-k8s/1.6/config/recipes/beats/stack_monitoring.yaml
  - HeartBeat: https://raw.githubusercontent.com/elastic/cloud-on-k8s/1.6/config/recipes/beats/heartbeat_es_kb_health.yaml

- Introducao na versão antiga (1.5) - **possui várias dicas de setup importantes**:
  - https://izekchen.medium.com/step-by-step-installation-for-elasticsearch-operator-on-kubernetes-and-metircbeat-filebeat-and-67a6ec4931fb
  - https://izekchen.medium.com/setup-elastic-apm-with-elasticsearch-operator-and-test-9b988fdb33ec