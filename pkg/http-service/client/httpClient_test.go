package httpclient

import (
	"math/rand"
	"reflect"
	"testing"
)

//ReverseArrays - simple client function that will call the httpService and reverse the arrays
func TestReverseArrays(t *testing.T) {

	serviceName := "localhost:8888"
	arrayOfFloats := make([]float32, 100)
	for i := 0; i < 100; i++ {
		value := rand.Float32()
		arrayOfFloats[i] = value
	}

	reversedArray, err := ReverseArrays(serviceName, &arrayOfFloats)
	if err != nil {
		t.Error(err.Error())
	}

	//Re-reverse back to the orinal
	reReversedArray, err := ReverseArrays(serviceName, reversedArray)
	if err != nil {
		t.Error("Error Rereversing: " + err.Error())
	}

	//Teset that the reReversed is equal
	if !reflect.DeepEqual(arrayOfFloats, *reReversedArray) {
		t.Error("Not Equal")
	}
}
