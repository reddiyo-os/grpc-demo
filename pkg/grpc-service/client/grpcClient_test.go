package grpcclient

import (
	"math/rand"
	"reflect"
	"testing"
)

//TestGrpcClientReverseArray - This is integration testing and puts the responsibility on the people using the client to write their tests
func TestGrpcClientReverseArray(t *testing.T) {

	arrayOfFloats := make([]float32, 100)
	for i := 0; i < 100; i++ {
		value := rand.Float32()
		arrayOfFloats[i] = value
	}
	//Construct the connection
	grpcClientConnection, err := ConstructClient("localhost:50051")
	if err != nil {
		t.Error(err.Error())
	}

	updatedArray, err := grpcClientConnection.ReverseArray(arrayOfFloats)
	if err != nil {
		t.Error("Error Reversing Array: " + err.Error())
	}

	//Reverse one more time
	finalReversed, err := grpcClientConnection.ReverseArray(updatedArray)
	if err != nil {
		t.Error("Error Getting the final array : " + err.Error())
	}

	if !reflect.DeepEqual(arrayOfFloats, finalReversed) {
		t.Error("Arrays Dont Match")
	}
}
