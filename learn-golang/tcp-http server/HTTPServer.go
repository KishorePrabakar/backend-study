package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/users", userHandler)
	http.HandleFunc("/health", healthHandler)

	fmt.Println("HTTP Server listening on port :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to the API!")
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("content/type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `
			[
		{"id":1, "name":"Kishore"},
		{"id":2, "name":"Nithish"},
		{"id":3, "name":"Nikil"}
			]
		`)
	} else if r.Method == "POST" {
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, `{"id":1, "name":"Pravin"}`)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"status":"healthy"}`)
}
