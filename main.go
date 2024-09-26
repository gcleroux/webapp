package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Hello, world!")
}

func main() {
	http.HandleFunc("/", hello)

	fmt.Println("Starting server on :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed to start: ", err)
	}
}
