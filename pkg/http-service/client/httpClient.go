package httpclient

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
)

//ReverseArrays - simple client function that will call the httpService and reverse the arrays
func ReverseArrays(serviceName string, arrayToReverse *[]float32) (*[]float32, error) {

	//Preconidtion
	if serviceName == "" {
		return nil, ClientConstructionError{Msg: "Empty Service Name"}
	}
	if arrayToReverse == nil || serviceName == "" || len(*arrayToReverse) == 0 {
		return nil, PreconditionError{Msg: "Empty Array"}
	}
	//Marshal arraryToReverse to body param
	requestBody, err := json.Marshal(arrayToReverse)
	if err != nil {
		return nil, InternalServerError{Msg: "Error Marshalling to JSON: " + err.Error()}
	}

	url := "http://" + serviceName + "/reverse"
	response, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, InternalServerError{Msg: "Error Calling Service: " + err.Error()}
	}
	defer response.Body.Close()
	//Handle non 200
	if response.StatusCode > 299 {
		return nil, HTTPError{Msg: "Non-200 Error", Code: response.StatusCode}
	}

	//Marshal it over to a slice
	var reversedArray []float32
	json.NewDecoder(response.Body).Decode(&reversedArray)
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

/*
HTTPError - Generic Error
*/
type HTTPError struct {
	Msg  string
	Code int
}

func (e HTTPError) Error() string {
	return e.Msg + " Error Code: " + strconv.Itoa(e.Code)
}
