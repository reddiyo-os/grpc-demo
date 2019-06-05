package grpcservice

/*
ReverseArray - function that will reverse the order of the array and return it.

Note: This function intentionally doesn't touch protobuf objects to maintain decoupling
*/
func ReverseArray(arraysToReverse []float32) ([]float32, error) {

	//Put this in just to show error
	if len(arraysToReverse) == 0 {
		return nil, GrpcPreconditionError{Msg: "Empty Array"}
	}
	//Loop through all the items in the array.
	//For efficiency purposes I am swapping the indexes with its mirrored counterpart
	//If there is an odd number then that odd number will maintain its position
	for i := 0; i < len(arraysToReverse)/2; i++ {
		j := len(arraysToReverse) - i - 1
		arraysToReverse[i], arraysToReverse[j] = arraysToReverse[j], arraysToReverse[i]
	}
	return arraysToReverse, nil
}

// GrpcPreconditionError - sample error that is used to show how error handling works
type GrpcPreconditionError struct {
	Msg string
}

func (e GrpcPreconditionError) Error() string {
	return e.Msg
}
