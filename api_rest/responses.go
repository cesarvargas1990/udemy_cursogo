package main

import (
	"net/http"
	"encoding/json"	
)

type Message struct {
	Status string `json:"status"`
	Message string `json:"message"`
}

func responseMovie(w http.ResponseWriter, status int, results Movie) {
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(results)
}

func responseMovies(w http.ResponseWriter, status int, results []Movie) {
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)
}
