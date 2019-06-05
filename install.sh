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
echo "Installing the first GRPC Microservice:  'kubectl apply -f deployments/grpcDeployment1/grpc-deployment-1.yaml'"
kubectl apply -f deployments/grpcDeployment1/grpc-deployment-1.yaml
## Install the orchestrator and the load balancer
echo "Insalling the orchestrator: "
kubectl apply -f deployments/orchestrator/orchestrator-deployment.yaml
## Sleep to let the load balancer get an IP
sleep 10
kubectl describe services orchestrator-service


