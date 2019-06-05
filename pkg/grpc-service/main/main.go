package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/reddiyo-os/grpc-demo/pkg/grpc-service/genProto"
	"github.com/reddiyo-os/grpc-demo/pkg/grpc-service/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// server is used to implement the GRPC Service
type server struct{}

func (s *server) ReverseArray(ctx context.Context, in *genproto.ReverseArrayRequest) (*genproto.ReverseArrayResponse, error) {

	//Get the array from the protobuf
	arrayOfNumbers := in.GetArrayOfNumbers()
	//reverse it
	reversedArray, err := grpcservice.ReverseArray(arrayOfNumbers)
	if err != nil {
		fmt.Println("Error Reversing Array: " + err.Error())
		//Need to convert from Service level errors to GRPC Errors
		switch err.(type) {
		case grpcservice.GrpcPreconditionError:
			return nil, status.Errorf(codes.FailedPrecondition, err.Error())
		default:
			return nil, status.Errorf(codes.Internal, err.Error())
		}
	}
	//Create the response object
	response := &genproto.ReverseArrayResponse{}
	response.ReversedArrayOfNumbers = reversedArray
	return response, nil
}

func main() {

	//Set up the server to listen - Puke if it cannot
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal("Not Listening: " + err.Error())
	}
	s := grpc.NewServer()
	//Load the protobuf definition
	genproto.RegisterExampleReddiyoGRPCServiceServer(s, &server{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatal("Not Listening: " + err.Error())
	}
}
