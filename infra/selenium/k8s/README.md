# selenium hub kubernetes

https://sahajamit.medium.com/spinning-up-your-own-scalable-selenium-grid-in-kubernetes-part-1-e4017bac68f4
https://sahajamit.medium.com/spinning-up-your-own-auto-scalable-selenium-grid-in-kubernetes-part-2-15b11f228ed8

minikube start --kubernetes-version 1.17.1 --driver='virtualbox' --cpus 4 --memory 8192 --disk-size 10GB --insecure-registry="gcr.io" \
        && minikube addons enable ingress


```
cd infra/selenium/k8s

kubectl delete -f selenium-hub-deploy.yaml
kubectl apply -f selenium-hub-deploy.yaml

kubectl delete -f selenium-node-chrome-deploy.yaml
kubectl apply -f selenium-node-chrome-deploy.yaml

kubectl port-forward services/selenium-hub 4444:4444
http://localhost:4444/grid/console



# kubectl expose deployment selenium-hub-deployment --type=NodePort --port=4444
# minikube service selenium-hub-deployment --url
# http://192.168.99.100:31178/grid/console

# sudo echo $(minikube ip) "selenium-grid.raulickis.com" >> /etc/hosts 
# tail -n 10 /etc/hosts
# curl -v http://selenium-grid.raulickis.com/grid/console

```


## Connect via VNC Viewer
```
kubectl get pods | grep selenium-node-chrome
kubectl port-forward pod/selenium-node-chrome-67db74b575-w68h5 5900:5900
```
Adress:   localhost:5900 
Password: secret




