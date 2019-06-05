package grpcservice

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

//TestReverseArrays - this test case is a service level test case and will NOT include any knowledge of the transport protocol (e.g. HTTP, Protobuf)
func TestReverseArrays(t *testing.T) {

	arrayOfFloats := make([]float32, 100)
	for i := 0; i < 100; i++ {
		value := rand.Float32()
		arrayOfFloats[i] = value
	}
	fmt.Println("Original Array: ")
	printFloatArray(arrayOfFloats)

	reversedArray, err := ReverseArray(arrayOfFloats)
	if err != nil {
		t.Error(err.Error())
	}
	//Test the length
	if len(reversedArray) != 100 {
		t.Error("Not enough results")
	}
	fmt.Println("Reversed Array: ")
	printFloatArray(reversedArray)
	//Reverse it again and ensure that it is the same as the starting point
	secondReversed, err := ReverseArray(reversedArray)
	if err != nil {
		t.Error(err.Error())
	}
	//test that it is identical
	if !reflect.DeepEqual(arrayOfFloats, secondReversed) {
		t.Error("Not a Match")
	}
	fmt.Println("Re-Reversed Array: ")
	printFloatArray(secondReversed)

	//Test the error case
	_, err = ReverseArray(make([]float32, 0))
	if err == nil {
		t.Error("DIdn't get an error when we expected it")
	}
}

func printFloatArray(floats []float32) {
	for i, value := range floats {
		fmt.Printf("%f", value)
		if i < len(floats)-1 {
			fmt.Printf(",")
		}
	}
	fmt.Print("\n")
}
