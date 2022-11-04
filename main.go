package main

import (
	"fmt"
	"net/http"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handle Request called")
	fmt.Fprintf(w, "Hello there")
}

func main() {
	http.HandleFunc("/hello", handleRequest)

	fmt.Println("Running on port 8090")
	http.ListenAndServe(":8090", nil)
	defer fmt.Println("Server shutting down, bye!")
}
