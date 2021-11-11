package main

import (
	"net/http"
	"encoding/json"	
)

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
