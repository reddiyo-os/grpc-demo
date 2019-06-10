package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/health/liveness", healthHandler)
	http.HandleFunc("/health/readiness", readinessHandler)
	fmt.Println("Handling GRPC HealthCheck on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(" Error:: " + err.Error())
		return
	}
}

//HTTP Health Handler
func healthHandler(w http.ResponseWriter, r *http.Request) {
	//Handle the methods that we don't support
	switch r.Method {
	case http.MethodPost:
		w.WriteHeader(405)
	case http.MethodPut:
		w.WriteHeader(405)
	case http.MethodDelete:
		w.WriteHeader(405)
	case http.MethodPatch:
		w.WriteHeader(405)
	}
	//NOTE:  This is where I would put the logic for checking the specific microservice

	w.WriteHeader(200)
}

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	//Handle the methods that we don't support
	switch r.Method {
	case http.MethodPost:
		w.WriteHeader(405)
	case http.MethodPut:
		w.WriteHeader(405)
	case http.MethodDelete:
		w.WriteHeader(405)
	case http.MethodPatch:
		w.WriteHeader(405)
	}
	//NOTE:  This is where I would put the logic for checking the specific microservice

	w.WriteHeader(200)
}
