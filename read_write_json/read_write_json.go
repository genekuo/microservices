package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloWorldRequest struct {
	Name string `json:"name"`
}

type helloWorldResponse struct {
	Message string `json:"message"`
}

func main() {
	port := 8080

	http.HandleFunc("/helloworld", helloworldHandler)

	log.Printf("Server starting on port %v\n", 8080)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))

}

func helloworldHandler(w http.ResponseWriter, r *http.Request) {

	var request helloWorldRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	response := helloWorldResponse{Message: "Hello " + request.Name}
	encoder := json.NewEncoder(w)
	encoder.Encode(&response)
}
