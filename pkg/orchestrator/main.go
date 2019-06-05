package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/reddiyo-os/grpc-demo/pkg/grpc-service/client"
)

//Main function that sets up the routes
func main() {
	http.HandleFunc("/grpc", grpcRequestHandler)
	http.HandleFunc("/http", httpRequestHandler)
	http.HandleFunc("/healthz", health)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println("Error Starting Server: " + err.Error())
	}
}

var microserviceConnections = []string{"grpc-microservice-1:50051", "grpc-microservice-2:50051", "grpc-microservice-3:50051", "grpc-microservice-4:50051", "grpc-microservice-5:50051"}

//grpcClient Connections
//This is not normally an array but it made the code much less verbose
var allConnections []*grpcclient.GrpcServiceClient

//The Init will handle constructing the client
func init() {
	//Loop through all the microservice connection strings and setup connections
	for _, value := range microserviceConnections {
		//Construct the connection
		tempConnection, err := grpcclient.ConstructClient(value)
		if err != nil {
			log.Fatal("Error getting Client: " + err.Error())
		}
		allConnections = append(allConnections, tempConnection)
	}
}

//Handles the grpcCall
func grpcRequestHandler(w http.ResponseWriter, req *http.Request) {
	//Get the time we start processing
	responseStartTime := time.Now()
	//Construct the main struct
	auditTrail := &RequestAuditTrail{
		StartTime: responseStartTime,
	}

	//Generate random array of floats - only used for testing purposes
	arrayOfFloats := make([]float32, 50)
	for i := 0; i < 50; i++ {
		value := rand.Float32()
		arrayOfFloats[i] = value
	}
	//Perform the microservice calls
	//This is done by looping through all the connections in the connections Variable
	for _, conn := range allConnections {
		//Call the first microservice
		updatedAudit, err := callMicroserviceAndReturnAuditTrail(conn, &arrayOfFloats)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//Append the Audit
		auditTrail.MSAuditTrail = append(auditTrail.MSAuditTrail, updatedAudit)
		arrayOfFloats = updatedAudit.ReturnedArray
	}

	//Get the end time
	endTime := time.Now()
	differenceInMicro := endTime.Sub(responseStartTime)

	//Append the duration and the end time
	auditTrail.DurationInNanos = differenceInMicro.Nanoseconds()
	auditTrail.EndTime = endTime

	//Convert the struct to JSON
	js, err := json.Marshal(auditTrail)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func httpRequestHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}

//helper function to call a microserivce and get the results
func callMicroserviceAndReturnAuditTrail(client *grpcclient.GrpcServiceClient, floatArray *[]float32) (*MicroserviceAuditTrail, error) {
	startCallTime := time.Now()
	singleCallAuditTrail := &MicroserviceAuditTrail{}
	updatedArray, err := client.ReverseArray(*floatArray)
	if err != nil {
		return nil, err
	}
	endTime := time.Now()
	singleCallAuditTrail.DurationOfMicroserviceCall = endTime.Sub(startCallTime).Nanoseconds()
	singleCallAuditTrail.ReturnedArray = updatedArray
	return singleCallAuditTrail, nil
}

//Handles the healthcheck
func health(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}

//#########  Structs

//RequestAuditTrail - Struct to hold an audit report of the data
type RequestAuditTrail struct {
	StartTime       time.Time                 `json:"StartTime"`
	EndTime         time.Time                 `json:"EndTime"`
	DurationInNanos int64                     `json:"DurationInNanos"`
	MSAuditTrail    []*MicroserviceAuditTrail `json:"MSAuditTrail"`
}

//MicroserviceAuditTrail - struct to store the results of a specific microservice call
type MicroserviceAuditTrail struct {
	DurationOfMicroserviceCall int64     `json:"DurationOfMicroserviceCall"`
	ReturnedArray              []float32 `json:"ReturnedArray"`
}
