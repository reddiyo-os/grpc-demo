# grpc-demo
This is a basic sample repo that sets up some microservices using GRPC and also with HTTP

## TLDR

```
./install.sh
```

## Overview

##Assumptions To Install in GKE

gcloud installed  - https://cloud.google.com/sdk/gcloud/reference/services/enable
kubectl installed

##Assumptions to run and build locally

Golang installed - https://golang.org/doc/install
Install GRPC - https://grpc.io/docs/quickstart/go/
docker - https://docs.docker.com/v17.12/install/
docker-compose - https://docs.docker.com/compose/install/


## Compile Protobuf

```
protoc -I api/v1/proto --proto_path=$GOPATH api/v1/proto/demo.proto --go_out=plugins=grpc:pkg/grpc-service/genProto
```
