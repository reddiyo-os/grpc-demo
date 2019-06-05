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

