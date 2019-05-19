# izmgprk8s

izmir gophers "Kubernetes ile Go Uygulamalarını Host Etmek" sunumu

## Run

~/go/bin/present -play

## Start

minikube start

# Expose

kubectl expose deployment app --type=NodePort -f deployment.yaml

# View

kubectl port-forward svc/app 8080:8080

kubectl expose deployment app --type=NodePort

minikube service app --url