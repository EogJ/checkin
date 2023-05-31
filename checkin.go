package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", entries)
	http.ListenAndServe(":8080", nil)
}

func entries(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}
