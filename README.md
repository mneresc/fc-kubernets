# fc-kubernets
Estudo de kubernets da full cycle

## Pontos importantes

### Kubernets master
* kube-apiserver
* kube-controllet-manager
* kube-shceduler

### Kubernets nodes
* kublet
* kubproxy
### Pod
* Conjunto de containers que fica dentro de um nó [pod](./kubernets/pod.yaml)
### ReplicaSet
* criar replicas do pod [replicaset](./kubernets/replicaset.yml)
* problema do replicaset: Quanto atualizo o replicaset pra pegar a versão nova tenho que matar os pods ou o replicaset 
* Resolvendo com Deployment 
### Deployment
* Provisiona os Pods
* ReplicaSets - replicar nos deploys
### Rollouts e Revisões
* Retornar deployments para a versão anterior
* kubectl rollout undo deployment goserver --to-revision=version
### Services
- Services atua como load balancer
- Expõe servicos de um deployment
- Tipos:
    ClusterIP: Criar cluster com um ip interno que só participantes do cluster veem - usado para servicos que não quer export
    NodePort: Antiga forma de bater num cluster expondo para fora
    LoadBalancer:

- 



## Comandos
- Listar cluster: kubectl config get-clusters
- Posso alternar entre cluster (minikube/kind/nuvem) - kubectl config use-context <cluster_name>
- descrever pod: kubectl describe pod <nome_pod>
- Aplicar um arquivo: kubectl apply -f kubernets/pod.yml
- Listar: kubectl get pod|replicasets
- Deletar pod: kubctl delete pod|replicasets|deployment goserver
- Histórico das versões: kubectl rollout history deployment goserver
## Ferramentas 
- minikube
- [kind](https://kind.sigs.k8s.io/docs/user/quick-start) - emula cluster com mais de um nó em containers dockers
- kind create cluster:  kind create cluster --config=kind.yaml --name=neres-test
- fazer apontamento do kubctl: kubectl cluster-info --context kind-neres-test

-- Parei [Aula1-concluida]
-- Aula 4.1 - Entendendo objstos de configuracao