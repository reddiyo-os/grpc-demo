package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	httpservice "github.com/reddiyo-os/grpc-demo/pkg/http-service/service"
)

func main() {

	http.HandleFunc("/reverse", reverseRequestHandler)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println("Error Starting Server: " + err.Error())
	}
}

func reverseRequestHandler(w http.ResponseWriter, req *http.Request) {
	//Decode the request
	decoder := json.NewDecoder(req.Body)
	var arrayOfFloats []float32
	err := decoder.Decode(&arrayOfFloats)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//Reverse the arrays
	reversedArray, err := httpservice.ReverseArray(arrayOfFloats)
	if err != nil {
		switch err.(type) {
		case httpservice.HTTPPreconditionError:
			http.Error(w, err.Error(), http.StatusBadRequest)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	//Convert the struct to JSON
	js, err := json.Marshal(reversedArray)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
