# fc-kubernets
Estudo de kubernets da full cycle

## Pontos importantes

- Kubernets master
    - kube-apiserver
    - kube-controllet-manager
    - kube-shceduler
- Kubernets nodes
    - kublet
    - kubproxy
- Pod
    - Conjunto de containers que fica dentro de um nó
- Deployment
    - Provisiona os Pods
    - ReplicaSets - replicar nos deploys
- Services
    - 

## Comandos
    - Listar cluster: kubectl config get-clusters
    - Posso alternar entre cluster (minikube/kind/nuvem) - kubectl config use-context <cluster_name>
## Ferramentas 
    - minikube
    - [kind](https://kind.sigs.k8s.io/docs/user/quick-start) - emula cluster com mais de um nó em containers dockers
        - kind create cluster:  kind create cluster --config=kind.yaml --name=neres-test
        - fazer apontamento do kubctl: kubectl cluster-info --context kind-neres-test

-- Parei [Aula1-concluida]
    -- Aula 2.1 - criando aplicacão exemplo e imagem