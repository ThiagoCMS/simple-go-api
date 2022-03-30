package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	http.ListenAndServe(":8080", r)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Hello World, this is a litte project I am doing to learn Golang"}
	json.NewEncoder(w).Encode(response)
}
