package main

import (
	"net/http"
	"encoding/json"
)

var collection = getSession().DB("movies").C("movies")

func MovieAdd(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var movie_data Movie 
	err := decoder.Decode(&movie_data)
	if (err !=nil) {
		panic(err)
	}
	defer r.Body.Close()
	err = collection.Insert(movie_data)	
	if err  != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(movie_data)
}