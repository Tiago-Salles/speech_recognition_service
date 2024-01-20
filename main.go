package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HandlerMethod func(http.ResponseWriter, *http.Request)

type Endpoint struct {
	path       string
	httpMethod string
	method     HandlerMethod
}

func hello(w http.ResponseWriter, r *http.Request) error {
	data := make(map[string]string)
	data["message"] = "Olá"

	httpResponse := HttpResponse{
		statusCode: 200,
		data:       data,
	}

	return writeResponse(w, httpResponse)

}

func writeResponse(w http.ResponseWriter, httpResponse HttpResponse) error {
	w.WriteHeader(httpResponse.statusCode)
	w.Header().Add("Content-Type", "application/json")
	response := json.NewEncoder(w).Encode(httpResponse.data)

	return response
}

func InitializeEndpoints() {
	var endpoints []Endpoint

	for _, endpoint := range endpoints {
		http.HandleFunc(endpoint.path, endpoint.method)
	}
}

func main() {
	fmt.Print("Initializing server...")
	InitializeEndpoints()
	http.ListenAndServe(":3000", nil)
	var content string = Run_speech_recognition()
	fmt.Print(content)
}
