package httpservice

/*
ReverseArray - function that will reverse the order of the array and return it.

Note: This function intentionally doesn't touch protobuf objects to maintain decoupling
*/
func ReverseArray(arrayToReverse []float32) ([]float32, error) {

	//Put this in just to show error
	if len(arrayToReverse) == 0 {
		return nil, HTTPPreconditionError{Msg: "Empty Array"}
	}
	//Loop through all the items in the array.
	//For efficiency purposes I am swapping the indexes with its mirrored counterpart
	//If there is an odd number then that odd number will maintain its position
	for i := 0; i < len(arrayToReverse)/2; i++ {
		j := len(arrayToReverse) - i - 1
		arrayToReverse[i], arrayToReverse[j] = arrayToReverse[j], arrayToReverse[i]
	}
	return arrayToReverse, nil
}

// HTTPPreconditionError - sample error that is used to show how error handling works
type HTTPPreconditionError struct {
	Msg string
}

func (e HTTPPreconditionError) Error() string {
	return e.Msg
}
