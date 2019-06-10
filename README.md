# grpc-demo
This is a basic sample repo that sets up some microservices using GRPC and also with HTTP.  At the end you will have a functioning set of microservices (11 of them) running in your VPC within GKE.  It isn't production ready but is a decent starting point for GRPC (and HTTP) microservices.

Technology Used:
*  **Golang**
*  **GKE**
*  **GRPC**
*  **Protobuf**
*  **Terraform**
*  **Docker** (and docker-compose)
*  **gcloud**
*  **kubectl**

## Tech Assumptions

* [GoogleAccount](https://cloud.google.com/billing/docs/how-to/manage-billing-account) - Project to install into
* [docker](https://docs.docker.com/v17.12/install/)
* [docker-compose](https://docs.docker.com/compose/install/)
* [Kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
* [Golang](https://golang.org/doc/install) - if you want to customize
* [GRPC](https://grpc.io/docs/quickstart/go/) - if you want to customize
* [terraform](https://learn.hashicorp.com/terraform/getting-started/install.html) - Optional


## TLDR

This command should get you going and do all the installation for you.  Good luck and don't forget to cleanup when you are done.

```
./install.sh
```

If you prefer terraform then go to the terraform section of this document.

## Overview

This is a very simple starter project that will setup a GKE Cluster with 11 Microservices.  The actual services don't do much (other than reverse arrays) but it is a good foundation for adding to and playing around with GRPC or HTTP calls with Golang.

* **Orchestrator** - this is the Service that handles the call from outside the cluster.   It will handle all calls to the internal microservices.
* **GRPC-Serivce** - - these microservices are dedicated to procesing Service Calls over GRPC.  They each have a single function that is exposed.
* **HTTP-Service** - these microservices do the exact same thing as the GRPC Service with the exception that HTTP is the protocol that is used.
* **Healthcheck Sidecar** - each pod will have a separate container deployed into it that owns the healthchecking of the pod.   It actually doesn't do anything other than return a 200 but the plumbing is there to do deeper healthchecks and readiness checks.

Arch:
![Alt text](docs/images/Reddiyo-OS_Example_GRPC.png?raw=true "Reddiyo-GRPC Arch")

### Terraform Alternative - Preferred Method

Really!?!?!   You want me to use a shell script to install?  Terraform works too.  You need to set up the variables for your environment.

#### Setup Your Variables

Update your variables first so that it points to your GCP Account.  You will need to download a service account.json and build the relative path to it.  If you don't know how to create a service account go [here](https://cloud.google.com/iam/docs/creating-managing-service-account-keys).

```
deployments/terraform/global/variables
```

#### Build Your Infrastructure

```
cd deployments/terraform/infrastructure
terraform init
terraform apply 
```

#### Install Microservices

```
cd deployments/terraform/app
terraform init
terraform apply
```

#### Delete the cluster

Don't forget to clean up your cluster so that you won't be charged.  You need to run terraform destroy from both the infrastructure folder and the app folder.
 ```
 terraform destroy
 ```


## Repository Structure

```
project
|   README.md
|   LICENSE
|   install.sh
|
└─── api - folder that stores API Definitions.  In this case it is the proto file for GRPC Microservices
|
└─── deployments - all the K8s files, docker files, or terraform files needed to build and deploy into GKE
|   |
|   └─── dockerFiles 
|   |
|   └─── grpcDeployment 
|   |
|   └─── httpDeployment 
|   |
|   └─── orchestrator
|   |
|   └─── terraform  
|
└─── docs - any supporting docs (e.g. Images)
|
└─── pkg - golang public packages
|   |
|   |─── grpc-service - the grpc microservice
|   |   └─── client - the client code
|   |   └─── genProto - any generated Protobuf files
|   |   └─── main - the main server code
|   |   └─── service - the service layer and functions
|   |
|   |─── http-service - the http microservice
|   |   └─── client - the client code
|   |   └─── main - the main server code
|   |   └─── service - the service layer and functions
|   |
|   |─── orchestrator - the orchestrator service that handles making both the GRPC and HTTP Calls
|   |
|   └─── healthcheck - the healthcheck sidecar
```

## Post Build Steps

After the build has completed you will be able to test your endpoints and run them with GRPC or with HTTP.  To do this you need to get the IP Address that google gave to your load balancer.   To do this you need to go to your Google Cloud Console.

1.  Click the Google Cloud Platform Dropdown --> Kubernetes Engine --> Services  
2.  Get the IP Address From the orchestrator "Load Balancer"
3.  Run your curls or postman endpoints

Curl using GRPC:
```
curl -v http://<YOUR_IP_ADDRESS>:8888/grpc
```

Curl Using HTTP:
```
curl -v http://<YOUR_IP_ADDRESS>:8888/http
```

## Build Notes and Commands

### Compile Protobuf

Below is the command that is used to compile proto file into GoLang.  This will need to be run from the root.

```
protoc -I api/v1/proto --proto_path=$GOPATH api/v1/proto/demo.proto --go_out=plugins=grpc:pkg/grpc-service/genProto
```

### Docker Build

You can build your own docker files from the source.   It is easiest to use docker-compose.  You can use the docker-compose.yml to set your params and then pass into the docker file.  Docker Files are located at "deployments/dockerFiles".  All docker builds are "from scratch" and should only be about 10MB


Note: If you fork the repository you need to update the docker file to pull from your repository.


Command to Build and Run
```
docker-compose up
```

#### Docker Registries with the latest built

* **GRPC Microservice** - https://cloud.docker.com/repository/docker/mornindew/grpc-demo-microservice
* **HTTP Microservice** - https://cloud.docker.com/repository/docker/mornindew/grpc-http-demo-microservice
* **Orchestrator Microservice** - https://cloud.docker.com/repository/docker/mornindew/reddiyo-os-orchestrator-example
* **Healthcheck Sidecar** - https://cloud.docker.com/repository/docker/mornindew/grpc-demo-healthcheck-sidecar
* **Base GRPC Container** - https://cloud.docker.com/repository/docker/mornindew/base-demo-protobuf