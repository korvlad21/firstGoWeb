package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(Response{
		Message: "Hello from Go!",
	})
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		fmt.Fprint(w, "<div>Hello World</div>")
	})

	// JSON API
	http.HandleFunc("/hello", helloHandler)

	http.ListenAndServe(":8080", nil)

	http.ListenAndServe(":8080", nil)
}
