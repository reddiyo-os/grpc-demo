package grpcclient

import (
	"context"
	"log"
	"time"

	genproto "github.com/reddiyo-os/grpc-demo/pkg/grpc-service/genProto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//GrpcServiceClient - main struct that is used to call the server
type GrpcServiceClient struct {
	userConn genproto.ExampleReddiyoGRPCServiceClient
	timeout  time.Duration
}

/*
ConstructClient - Constructor to return a client

Typically the constructor doesn't take any vars but this one does so that I can use the same client on different service locations.

Normally I would use an Environment Variable to pull the DNS location for the client
*/
func ConstructClient(location string) (*GrpcServiceClient, error) {
	//Precondition
	if location == "" {
		return nil, PreconditionError{Msg: "Location Not Set"}
	}
	//Dial the server - this will open the pipe to the server to be multiplexed
	//This is where you would configure your TLS || Client Load Balancing
	connection, err := grpc.Dial(location, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Could not get the connection. " + err.Error())
		return nil, ClientConstructionError{Msg: "Could not get a client"}
	}

	constructedClient := &GrpcServiceClient{
		userConn: genproto.NewExampleReddiyoGRPCServiceClient(connection),
		timeout:  time.Second,
	}
	return constructedClient, nil
}

//GrpcServiceClient - methods

/*
ReverseArray - call into the grpcService and will reverse the array
*/
func (client *GrpcServiceClient) ReverseArray(arrayToReverse *[]float32) (*[]float32, error) {
	//Precondition Check - I do a precondition check in the client to avoid making a call to the server for anything that isn't well constructed
	if arrayToReverse == nil || len(*arrayToReverse) == 0 {
		return nil, PreconditionError{Msg: "Empty Array"}
	}

	//sets up the cancel function and the context
	ctx, cancel := context.WithTimeout(context.Background(), client.timeout)
	defer cancel()

	//Construct the protobuf params
	//We don't let the protobuf structs leak out of the client or the server layer
	request := &genproto.ReverseArrayRequest{
		ArrayOfNumbers: *arrayToReverse,
	}
	//Make the call to the server
	response, err := client.userConn.ReverseArray(ctx, request)
	if err != nil {
		//Handle error Cases - thse could be GRPC, Connection, or custom
		statusCode, ok := status.FromError(err)
		if !ok {
			//This happens if we cannot get the status - will throw the generic
			return nil, InternalServerError{Msg: err.Error()}
		}
		switch statusCode.Code() {
		case codes.FailedPrecondition:
			return nil, PreconditionError{Msg: err.Error()}
		default:
			return nil, InternalServerError{Msg: err.Error()}
		}
	}
	//REturn the Reversed array - we don't allow the protobuf objects to leave the client
	reversedArray := response.GetReversedArrayOfNumbers()
	return &reversedArray, nil
}

//Error Structs -- THese are duplicated intentionally

/*
PreconditionError - Precondition error
*/
type PreconditionError struct {
	Msg string
}

func (e PreconditionError) Error() string {
	return e.Msg
}

/*
ClientConstructionError - Precondition error
*/
type ClientConstructionError struct {
	Msg string
}

func (e ClientConstructionError) Error() string {
	return e.Msg
}

/*
InternalServerError - Generic Error
*/
type InternalServerError struct {
	Msg string
}

func (e InternalServerError) Error() string {
	return e.Msg
}
