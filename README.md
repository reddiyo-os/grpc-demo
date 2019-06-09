# grpc-demo
This is a basic sample repo that sets up some microservices using GRPC and also with HTTP.  All code is written in Golang and should be a decent foundation for playing with the different transport protocols.

## TLDR

This command should get you going and do all the installation for you.  Good luck and don't forget to cleanup when you are done.

```
./install.sh
```

## Overview

This is a very simple starter project that will setup a GKE Cluster with 11 Microservices.  The actual services don't do much (other than reverse arrays) but it is a good foundation for adding to and playing around with GRPC or HTTP calls with Golang.

* Orchestrator - this is the Service that handles the call from outside the cluster.   It will handle all calls to the internal microservices.
* GRPC-Serivce-# - these microservices are dedicated to procesing Service Calls over GRPC.  They each have a single function that is exposed.
* HTTP-Service-# - these microservices do the exact same thing as the GRPC Service with the exception that HTTP is the protocol that is used.

Arch:
![Alt text](docs/images/Reddiyo-OS_Example_GRPC.png?raw=true "Reddiyo-GRPC Arch")

### Terraform Alternative - Preferred Method

Really!?!?!   You want me to use a shell script to install?  Terraform works too.  You need to set up the variables for your environment.

#### Setup Your Variables

Update your variables first so that it points to your GCP Account

```
deployments/terraform/global/variables
```

#### Build Your Infrastructure

```
cd deployments/terraform/infrastructure
terraform init
terraform plan -out demoInfrastructure
terraform apply "demoInfrastructure"
```

#### Install Microservices

```
cd deployments/terraform/app
terraform init
terraform plan -out demoServiceInstall
terraform apply "demoServiceInstall"
```

#### Delete the cluster

Don't forget to clean up your cluster so that you won't be charged.  You need to run terraform destroy from both the infrastructure folder and the app folder.
 ```
 terraform destroy
 ```

### Assumptions To Install in GKE

A couple assumptions are made when running the "install.sh" command.

* gcloud installed  - https://cloud.google.com/sdk/gcloud/reference/services/enable
* Available Project in Google Cloud

### Assumptions to run and build locally

* Golang installed - https://golang.org/doc/install
* Install GRPC - https://grpc.io/docs/quickstart/go/
* docker - https://docs.docker.com/v17.12/install/
* docker-compose - https://docs.docker.com/compose/install/


### Repository Structure

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
```

### Compile Protobuf

Below is the command that is used to compile proto file into GoLang.  This will need to be run from the root.

```
protoc -I api/v1/proto --proto_path=$GOPATH api/v1/proto/demo.proto --go_out=plugins=grpc:pkg/grpc-service/genProto
```
