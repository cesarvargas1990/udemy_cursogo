package main

import (
	"net/http"
	"encoding/json"
)

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
	responseMovie(w,200,movie_data)
}