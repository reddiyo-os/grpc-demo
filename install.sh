#!/bin/bash

####  ASSUMPTIONS - 
#####   gcloud installed  - https://cloud.google.com/sdk/gcloud/reference/services/enable
#####  kubectl installed

## Get the project ID and the set it as the default
read -p "Enter the ProjectID for your GCP Project: "  projectID
echo  $projectID
echo "Running the Command 'gcloud config set project $projectID'"
gcloud config set project $projectID

## Get the cluster name and create it
read -p "What do you want to call your new cluster: "  clusterName
echo "Running the Command 'gcloud container clusters create $clusterName'"
gcloud container clusters create $clusterName
## Install the grpc service
echo "Installing the GRPC Microservice #1:"
kubectl apply -f deployments/grpcDeployment/grpc-deployment-1.yaml
echo "Installing the GRPC Microservice #2:"
kubectl apply -f deployments/grpcDeployment/grpc-deployment-2.yaml
echo "Installing the GRPC Microservice #3:"
kubectl apply -f deployments/grpcDeployment/grpc-deployment-3.yaml
echo "Installing the GRPC Microservice #4:"
kubectl apply -f deployments/grpcDeployment/grpc-deployment-4.yaml
echo "Installing the GRPC Microservice #5:"
kubectl apply -f deployments/grpcDeployment/grpc-deployment-5.yaml

## Install the http services
echo "Installing the HTTP Microservice #1:"
kubectl apply -f deployments/httpDeployment/http-deployment-1.yaml
echo "Installing the HTTP Microservice #2:"
kubectl apply -f deployments/httpDeployment/http-deployment-2.yaml
echo "Installing the HTTP Microservice #3:"
kubectl apply -f deployments/httpDeployment/http-deployment-3.yaml
echo "Installing the HTTP Microservice #4:"
kubectl apply -f deployments/httpDeployment/http-deployment-4.yaml
echo "Installing the HTTP Microservice #5:"
kubectl apply -f deployments/httpDeployment/http-deployment-5.yaml
## Install the orchestrator and the load balancer
echo "Insalling the orchestrator: "
kubectl apply -f deployments/orchestrator/orchestrator-deployment.yaml
echo "ALL Done - the pods and services will take a minute or two to fully deploy"


