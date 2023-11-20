#!/bin/bash

go build -o /output/app ./cmd/api

# Image build

# sudo docker build -t adapter:latest -f config/adapter/Dockerfile .

# KIND
# sudo kind create cluster
# sudo kubectl cluster-info --context kind-kind

# GUI
# kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.6.1/aio/deploy/recommended.yaml
# kubectl get all -n kubernetes-dashboard


# Correr solo kubecost
# helm install kubecost cost-analyzer --repo https://kubecost.github.io/cost-analyzer/ --namespace kubecost --create-namespace --set kubecostToken="ZGFyaW9zb3RlbG8zQGdtYWlsLmNvbQ==xm343yadf98"
# kubectl port-forward --namespace kubecost deployment/kubecost-cost-analyzer 9090


# Correr todo compleot kubecost + thanos
# Original
# helm upgrade kubecost kubecost/cost-analyzer --install --namespace kubecost -f https://raw.githubusercontent.com/kubecost/cost-analyzer-helm-chart/master/cost-analyzer/values-thanos.yaml -f values-clusterName.yaml

# Funcional
# helm install kubecost cost-analyzer --repo https://kubecost.github.io/cost-analyzer/  --namespace kubecost --create-namespace --set kubecostToken="ZGFyaW9zb3RlbG8zQGdtYWlsLmNvbQ==xm343yadf98" -f values-thanos.yaml -f values-clusterName.yaml

# SI el helm esta en uso -> eleiminar
# kubectl -n kubecost get secrets | grep helm
# kubectl delete secret sh.helm.release.v1.kubecost.v1 -n kubecost
# kubectl port-forward --namespace kubecost [POD] 9090

# Ver pods (thanos)
# kubectl get pods -A

# Editar config del pod
# kubectl edit pod [POD] --namespace kubecost

# Puertos
# 9090 kubecost
# 3000 grafana
# 10902 thanos 