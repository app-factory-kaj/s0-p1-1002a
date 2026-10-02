package main

import (
	"log"
	"net/http"

	"greeter/internal/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello", handlers.GetHello)

	log.Println("greeter listening on :9090")
	if err := http.ListenAndServe(":9090", mux); err != nil {
		log.Fatal(err)
	}
}
