package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Received request from %s: %s %s\n", r.RemoteAddr, r.Method, r.URL.Path)
	})

	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}