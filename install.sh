#!/bin/bash

## Get the project ID and the set it as the default
read -p "Enter the ProjectID for your GCP Project: "  projectID
echo  $projectID
echo "Running the Command 'gcloud config set project $projectID'"
gcloud config set project $projectID

## Get the cluster name and create it
read -p "What do you want to call your new cluster: "  clusterName
echo "Running the Command 'gcloud container clusters create $clusterName'"
gcloud container clusters create $clusterName --enable-autoscaling --max-nodes=10 --min-nodes=3
echo "Cluster Completed, now installing the services"


printf "\n"
## Install the grpc service
echo "Installing All the GRPC Microservices:"
kubectl apply -f deployments/grpcDeployment/grpc-deployment-1.yaml
echo "GRPC Microservice 1 Installed"
kubectl apply -f deployments/grpcDeployment/grpc-deployment-2.yaml
echo "GRPC Microservice 2 Installed"
kubectl apply -f deployments/grpcDeployment/grpc-deployment-3.yaml
echo "GRPC Microservice 3 Installed"
kubectl apply -f deployments/grpcDeployment/grpc-deployment-4.yaml
echo "GRPC Microservice 4 Installed"
kubectl apply -f deployments/grpcDeployment/grpc-deployment-5.yaml
echo "GRPC Microservice 5 Installed"

printf "\n"

## Install the http services
echo "Installing All the HTTP Microservices:"
kubectl apply -f deployments/httpDeployment/http-deployment-1.yaml
echo "HTTP Microservice 1 Installed"
kubectl apply -f deployments/httpDeployment/http-deployment-2.yaml
echo "HTTP Microservice 2 Installed"
kubectl apply -f deployments/httpDeployment/http-deployment-3.yaml
echo "HTTP Microservice 3 Installed"
kubectl apply -f deployments/httpDeployment/http-deployment-4.yaml
echo "HTTP Microservice 4 Installed"
kubectl apply -f deployments/httpDeployment/http-deployment-5.yaml
echo "HTTP Microservice 5 Installed"

printf "\n"

## Install the orchestrator and the load balancer
echo "Insalling the orchestrator: "
kubectl apply -f deployments/orchestrator/orchestrator-deployment.yaml
echo "ALL Done - the pods and services will take a minute or two to fully deploy"


