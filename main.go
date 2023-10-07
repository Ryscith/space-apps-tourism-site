package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}

func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
