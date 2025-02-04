package main

import (
	"log"
	api "mcezario/backend-challenge/api/matrix"
	"net/http"
)

func handleRequests() {
	http.Handle("/echo", http.HandlerFunc(api.Echo))
	http.Handle("/invert", http.HandlerFunc(api.Invert))
	http.Handle("/flatten", http.HandlerFunc(api.Flatten))
	http.Handle("/sum", http.HandlerFunc(api.Sum))
	http.Handle("/multiply", http.HandlerFunc(api.Multiply))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
