package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/grpc", grpcRequestHandler)
	http.HandleFunc("/healthz", health)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println("Error Starting Server: " + err.Error())
	}
}

//Handles the grpcCall
func grpcRequestHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}

//Handles the healthcheck
func health(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}
