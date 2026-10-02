package handlers

import (
	"encoding/json"
	"net/http"
)

type Greeting struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

func GetHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}

	greeting := Greeting{
		Name:    name,
		Message: "Hello, " + name + "!",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(greeting)
}
