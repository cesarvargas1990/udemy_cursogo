package main

import (
	"net/http"
	"encoding/json"
	"log"
	"fmt"
)

func MovieList(w http.ResponseWriter, r *http.Request) {
	var results []Movie
	err := collection.Find(nil).Sort("-_id").All(&results)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Resultados", results)
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)
}